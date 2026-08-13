// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/manage/resource_fabric_aci"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	nd "github.com/netascode/go-nd"
)

func fabricAciImportStateCheck(fabricName, username, password string) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		if len(states) != 1 {
			return fmt.Errorf("expected one imported Fabric ACI state, got %d", len(states))
		}

		state := states[0]
		if state.ID != fabricName {
			return fmt.Errorf("expected imported Fabric ACI id %q, got %q", fabricName, state.ID)
		}
		if got := state.Attributes["fabric_name"]; got != fabricName {
			return fmt.Errorf("expected imported fabric_name %q, got %q", fabricName, got)
		}
		if got := state.Attributes["username"]; got != username {
			return fmt.Errorf("expected imported username from environment variable to be %q, got %q", username, got)
		}
		if got := state.Attributes["password"]; got != password {
			return fmt.Errorf("expected imported password to match the environment variable")
		}

		return nil
	}
}

func deleteFabricAciOutsideTerraform(t *testing.T, fabricName, username, password string) {
	t.Helper()

	cfg := helper.GetConfig("global")
	client, err := nd.NewClient(
		cfg.ND.URL,
		"/api/v1",
		cfg.ND.User,
		cfg.ND.Password,
		"",
		cfg.ND.Insecure == "true",
		nd.MaxRetries(3),
	)
	if err != nil {
		t.Fatalf("failed to create ND client for out-of-band Fabric ACI delete: %v", err)
	}

	payload, err := json.Marshal(map[string]interface{}{
		"credentials": map[string]string{
			"user":     username,
			"password": password,
		},
		"force": false,
	})
	if err != nil {
		t.Fatalf("failed to marshal out-of-band Fabric ACI delete payload: %v", err)
	}

	fabricAPI := api.NewFabricAciAPI(&client, ndapi.DefaultFabric)
	fabricAPI.ClusterName = fabricName
	fabricAPI.Delete = true

	res, err := fabricAPI.Post(payload, &ndapi.APIOptions{DisablePayloadLog: true})
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			t.Logf("Fabric ACI %q was already absent before Terraform destroy", fabricName)
			return
		}
		t.Fatalf("failed to delete Fabric ACI %q outside Terraform: %v %v", fabricName, err, res)
	}

	deadline := time.Now().Add(2 * time.Minute)
	for {
		_, err := fabricAPI.Get()
		if err != nil {
			if strings.Contains(err.Error(), "StatusCode 404") {
				return
			}
			t.Fatalf("failed to verify out-of-band deletion of Fabric ACI %q: %v", fabricName, err)
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for Fabric ACI %q to be deleted outside Terraform", fabricName)
		}
		time.Sleep(5 * time.Second)
	}
}

func TestAccFabricAciResourceCRUD(t *testing.T) {
	const (
		rscAddr  = "nd_fabric_aci.fabric_test"
		hostname = "1.1.1.1"
		username = "admin"
		password = "**********"
	)

	cfg := helper.GetConfig("global")
	// fabricName := "tf_" + acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	fabricName := "tf-kyzef"
	environmentVariablePrefix := strings.ToUpper(strings.ReplaceAll(fabricName, "-", "_"))
	forceEnvironmentVariable := environmentVariablePrefix + "_FORCE"
	usernameEnvironmentVariable := environmentVariablePrefix + "_USERNAME"
	passwordEnvironmentVariable := environmentVariablePrefix + "_PASSWORD"
	terraformPlanArgs, terraformPlanArgsSet := os.LookupEnv("TF_CLI_ARGS_plan")
	restoreTerraformPlanArgs := func() {
		if terraformPlanArgsSet {
			_ = os.Setenv("TF_CLI_ARGS_plan", terraformPlanArgs)
			return
		}
		_ = os.Unsetenv("TF_CLI_ARGS_plan")
	}
	t.Cleanup(func() {
		_ = os.Unsetenv(forceEnvironmentVariable)
		_ = os.Unsetenv(usernameEnvironmentVariable)
		_ = os.Unsetenv(passwordEnvironmentVariable)
		restoreTerraformPlanArgs()
	})

	testCases := []map[string]string{
		{"name": "create_required", "purpose": "Create Fabric ACI with only required attributes."},
		{"name": "update_optional", "purpose": "Configure location, license, security, certificate verification, and enable orchestration."},
		{"name": "re_register_and_update", "purpose": "Re-register while updating the license tier and keeping orchestration enabled."},
		{"name": "empty_plan", "purpose": "Reapply the unchanged re-registration configuration and require an empty plan."},
		{"name": "destroy_in_use", "purpose": "Verify normal destroy fails while orchestration is enabled."},
		{"name": "disable_before_force", "purpose": "Disable orchestration before testing forced removal."},
		{"name": "import_with_credentials", "purpose": "Import the existing fabric with credentials supplied by fabric-scoped environment variables."},
		{"name": "force_destroy", "purpose": "Verify the fabric-scoped force environment variable permits destroy."},
		{"name": "import_missing", "purpose": "Verify importing the removed fabric reports resource not found."},
		{"name": "recreate_enabled", "purpose": "Recreate the fabric with orchestration enabled."},
		{"name": "disable_before_normal_destroy", "purpose": "Disable orchestration before de-registration and normal removal."},
		{"name": "de_register_re_register_conflict", "purpose": "Reject configuring de-register and re-register together before an API action."},
		{"name": "de_register_with_update", "purpose": "Reject de-registration combined with an ordinary attribute update."},
		{"name": "de_register_only", "purpose": "De-register only after all local fabric features are disabled."},
		{"name": "de_register_reset", "purpose": "Reset the de-register action flag without another backend action."},
		{"name": "re_register_only", "purpose": "Restore registration before continuing the resource lifecycle."},
		{"name": "re_register_reset", "purpose": "Reset the re-register action flag without another backend action."},
		{"name": "destroy_after_external_delete", "purpose": "Delete outside Terraform and verify provider Delete treats HTTP 404 as success."},
	}
	for _, testCase := range testCases {
		t.Logf("Coverage: %s - %s", testCase["name"], testCase["purpose"])
	}

	x := &map[string]string{
		"RscType":  "nd_fabric_aci",
		"RscName":  "fabric_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	fabricRsc := new(resource_fabric_aci.NDFCFabricAciModel)

	redactConfig := func(config string) string {
		config = strings.ReplaceAll(config, password, "<redacted>")
		if cfg.ND.Password != "" {
			config = strings.ReplaceAll(config, cfg.ND.Password, "<redacted>")
		}
		return config
	}

	fabricAciStateCheck := func(model resource_fabric_aci.NDFCFabricAciModel) resource.TestCheckFunc {
		checks := append(
			FabricAciModelHelperStateCheck(rscAddr, model, path.Empty()),
			resource.TestCheckResourceAttr(rscAddr, "id", fabricName),
		)
		if model.Spec.Location.Latitude != nil {
			checks = append(checks, resource.TestCheckResourceAttr(
				rscAddr,
				"latitude",
				strconv.FormatFloat(*model.Spec.Location.Latitude, 'f', -1, 64),
			))
		}
		if model.Spec.Location.Longitude != nil {
			checks = append(checks, resource.TestCheckResourceAttr(
				rscAddr,
				"longitude",
				strconv.FormatFloat(*model.Spec.Location.Longitude, 'f', -1, 64),
			))
		}
		return resource.ComposeTestCheckFunc(checks...)
	}

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}
	s4 := &helper.StepInfo{}
	s5 := &helper.StepInfo{}
	s6 := &helper.StepInfo{}
	s7 := &helper.StepInfo{}
	s8 := &helper.StepInfo{}
	s9 := &helper.StepInfo{}
	s10 := &helper.StepInfo{}
	s11 := &helper.StepInfo{}
	s12 := &helper.StepInfo{}
	s13 := &helper.StepInfo{}
	s14 := &helper.StepInfo{}
	s15 := &helper.StepInfo{}
	s16 := &helper.StepInfo{}
	s17 := &helper.StepInfo{}
	s18 := &helper.StepInfo{}
	completedStep := 0
	requirePreviousStep := func(step int) {
		expectedPreviousStep := step - 1
		if completedStep != expectedPreviousStep {
			t.Fatalf(
				"Step %d cannot run because step %d did not complete successfully; last completed step: %d",
				step,
				expectedPreviousStep,
				completedStep,
			)
		}
	}
	completeStep := func(step int) {
		requirePreviousStep(step)
		completedStep = step
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with required attributes only.
			{
				Config: func() string {
					*stepCount++
					s1.Index = *stepCount
					s1.Name = fmt.Sprintf("%s_%d_create_required", t.Name(), s1.Index)

					helper.GenerateFabricAciObject(
						&fabricRsc,
						fabricName,
						hostname,
						username,
						password,
						nil,
					)
					helper.GetTFConfigWithSingleResource(s1.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(1)
					helper.LogStep(t, s1.Index, s1.Name, redactConfig(s1.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(1) },
			},
			// Step 2: Configure optional attributes and enable orchestration.
			{
				Config: func() string {
					*stepCount++
					s2.Index = *stepCount
					s2.Name = fmt.Sprintf("%s_%d_update_optional_enable_orchestration", t.Name(), s2.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"latitude":             float64(9),
						"longitude":            float64(-128),
						"verify_ca":            false,
						"security_domain":      "all",
						"license_tier":         "advantage",
						"orchestration_status": "enabled",
					})
					helper.GetTFConfigWithSingleResource(s2.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(2)
					helper.LogStep(t, s2.Index, s2.Name, redactConfig(s2.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(2) },
			},
			// Step 3: Re-register, update the license tier, and keep orchestration enabled.
			{
				Config: func() string {
					*stepCount++
					s3.Index = *stepCount
					s3.Name = fmt.Sprintf("%s_%d_re_register_update_optional", t.Name(), s3.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"license_tier": "premier",
						"re_register":  true,
					})
					helper.GetTFConfigWithSingleResource(s3.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(3)
					helper.LogStep(t, s3.Index, s3.Name, redactConfig(s3.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(3) },
			},
			// Step 4: Reapply the same configuration and require an empty plan.
			{
				Config: func() string {
					*stepCount++
					s4.Index = *stepCount
					s4.Name = fmt.Sprintf("%s_%d_reapply_expect_empty_plan", t.Name(), s4.Index)

					helper.GetTFConfigWithSingleResource(s4.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s4.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(4)
					helper.LogStep(t, s4.Index, s4.Name, redactConfig(s4.Cfg))
				},
				PlanOnly:      true,
				PostApplyFunc: func() { completeStep(4) },
			},
			// Step 5: Normal destroy must fail while orchestration is enabled.
			{
				Config: func() string {
					*stepCount++
					s5.Index = *stepCount
					s5.Name = fmt.Sprintf("%s_%d_destroy_expect_orchestration_error", t.Name(), s5.Index)

					helper.GetTFConfigWithSingleResource(s5.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s5.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(5)
					helper.LogStep(t, s5.Index, s5.Name, redactConfig(s5.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before normal destroy: %v", forceEnvironmentVariable, err)
					}
				},
				Destroy: true,
				ExpectError: regexp.MustCompile(
					`(?s)Error Deleting Fabric ACI.*StatusCode 400.*Cannot\s+disconnect\s+fabric\..*enabled\s+features:\s+Orchestration\..*disable\s+these\s+features\s+first`,
				),
			},
			// Step 6: Disable orchestration before testing forced removal.
			{
				Config: func() string {
					*stepCount++
					s6.Index = *stepCount
					s6.Name = fmt.Sprintf("%s_%d_disable_orchestration_before_force", t.Name(), s6.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"orchestration_status": "disabled",
						"re_register":          false,
					})

					helper.GetTFConfigWithSingleResource(s6.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s6.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					// Reaching step 6 means the framework accepted step 5's expected error.
					completeStep(5)
					requirePreviousStep(6)
					helper.LogStep(t, s6.Index, s6.Name, redactConfig(s6.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before disabling orchestration: %v", forceEnvironmentVariable, err)
					}
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(6) },
			},
			// Step 7: Import the existing fabric with credentials supplied through
			// fabric-scoped environment variables in the isolated import state.
			{
				Config: func() string {
					*stepCount++
					s7.Index = *stepCount
					s7.Name = fmt.Sprintf("%s_%d_import_with_environment_credentials", t.Name(), s7.Index)

					helper.GetTFConfigWithSingleResource(s7.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s7.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(7)
					helper.LogStep(t, s7.Index, s7.Name, redactConfig(s7.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before import: %v", forceEnvironmentVariable, err)
					}
					if err := os.Setenv(usernameEnvironmentVariable, username); err != nil {
						t.Fatalf("Failed to set %s for import: %v", usernameEnvironmentVariable, err)
					}
					if err := os.Setenv(passwordEnvironmentVariable, password); err != nil {
						t.Fatalf("Failed to set %s for import: %v", passwordEnvironmentVariable, err)
					}
				},
				ResourceName:  rscAddr,
				ImportState:   true,
				ImportStateId: fabricName,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if err := fabricAciImportStateCheck(fabricName, username, password)(states); err != nil {
						return err
					}
					completeStep(7)
					return nil
				},
			},
			// Step 8: Remove the import environment variables, set the
			// fabric-scoped force variable, and destroy successfully.
			{
				Config: func() string {
					*stepCount++
					s8.Index = *stepCount
					s8.Name = fmt.Sprintf("%s_%d_force_destroy_imported_fabric", t.Name(), s8.Index)

					helper.GetTFConfigWithSingleResource(s8.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s8.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(8)
					helper.LogStep(t, s8.Index, s8.Name, redactConfig(s8.Cfg))
					if err := os.Unsetenv(usernameEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before forced destroy: %v", usernameEnvironmentVariable, err)
					}
					if err := os.Unsetenv(passwordEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before forced destroy: %v", passwordEnvironmentVariable, err)
					}
					if err := os.Setenv(forceEnvironmentVariable, "true"); err != nil {
						t.Fatalf("Failed to set %s for forced destroy: %v", forceEnvironmentVariable, err)
					}
				},
				Destroy: true,
				PostApplyFunc: func() {
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s after forced destroy: %v", forceEnvironmentVariable, err)
					}
					completeStep(8)
				},
			},
			// Step 9: Import the removed fabric and verify the import Read path
			// reports resource not found.
			{
				Config: func() string {
					*stepCount++
					s9.Index = *stepCount
					s9.Name = fmt.Sprintf("%s_%d_import_missing_fabric", t.Name(), s9.Index)

					helper.GetTFConfigWithSingleResource(s9.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s9.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(9)
					helper.LogStep(t, s9.Index, s9.Name, redactConfig(s9.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before missing import: %v", forceEnvironmentVariable, err)
					}
				},
				ResourceName:  rscAddr,
				ImportState:   true,
				ImportStateId: fabricName,
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`(?s)Error Importing Fabric ACI.*Could not import nd_fabric_aci with id\s+%q:\s+resource\s+not\s+found`, fabricName),
				),
			},
			// Step 10: Recreate with orchestration enabled.
			{
				Config: func() string {
					*stepCount++
					s10.Index = *stepCount
					s10.Name = fmt.Sprintf("%s_%d_recreate_with_orchestration_enabled", t.Name(), s10.Index)

					helper.GenerateFabricAciObject(
						&fabricRsc,
						fabricName,
						hostname,
						username,
						password,
						map[string]interface{}{
							"latitude":             float64(9),
							"longitude":            float64(-128),
							"verify_ca":            false,
							"security_domain":      "all",
							"license_tier":         "advantage",
							"orchestration_status": "enabled",
						},
					)
					helper.GetTFConfigWithSingleResource(s10.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s10.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					// Reaching step 10 means the framework accepted step 9's expected error.
					completeStep(9)
					requirePreviousStep(10)
					helper.LogStep(t, s10.Index, s10.Name, redactConfig(s10.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before recreate: %v", forceEnvironmentVariable, err)
					}
					// Allow Nexus Dashboard to completely clear the previous configuration.
					time.Sleep(10 * time.Second)
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(10) },
			},
			// Step 11: Disable orchestration before de-registration and out-of-band removal.
			{
				Config: func() string {
					*stepCount++
					s11.Index = *stepCount
					s11.Name = fmt.Sprintf("%s_%d_disable_orchestration_before_external_delete", t.Name(), s11.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"orchestration_status": "disabled",
					})
					helper.GetTFConfigWithSingleResource(s11.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s11.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(11)
					helper.LogStep(t, s11.Index, s11.Name, redactConfig(s11.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(11) },
			},
			// Step 12: Reject configuring de-register and re-register together.
			{
				Config: func() string {
					*stepCount++
					s12.Index = *stepCount
					s12.Name = fmt.Sprintf("%s_%d_de_register_re_register_conflict", t.Name(), s12.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"de_register": true,
						"re_register": true,
					})
					helper.GetTFConfigWithSingleResource(s12.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s12.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(12)
					helper.LogStep(t, s12.Index, s12.Name, redactConfig(s12.Cfg))
				},
				ExpectError: regexp.MustCompile(
					`(?s)Invalid Attribute Combination.*cannot be specified when.*specified`,
				),
			},
			// Step 13: Reject de-registration combined with a normal attribute update.
			{
				Config: func() string {
					*stepCount++
					s13.Index = *stepCount
					s13.Name = fmt.Sprintf("%s_%d_de_register_with_normal_update", t.Name(), s13.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"de_register":  true,
						"re_register":  false,
						"license_tier": "premier",
					})
					helper.GetTFConfigWithSingleResource(s13.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s13.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					// Reaching step 13 means the framework accepted step 12's expected error.
					completeStep(12)
					requirePreviousStep(13)
					helper.LogStep(t, s13.Index, s13.Name, redactConfig(s13.Cfg))
				},
				ExpectError: regexp.MustCompile(
					`(?s)Invalid Fabric ACI De-registration Update.*de_register action cannot be combined with changes to other nd_fabric_aci attributes`,
				),
			},
			// Step 14: Revert the ordinary attribute change and de-register only.
			{
				Config: func() string {
					*stepCount++
					s14.Index = *stepCount
					s14.Name = fmt.Sprintf("%s_%d_de_register_only", t.Name(), s14.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"license_tier": "advantage",
						"de_register":  true,
					})
					helper.GetTFConfigWithSingleResource(s14.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s14.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					// Reaching step 14 means the framework accepted step 13's expected error.
					completeStep(13)
					requirePreviousStep(14)
					helper.LogStep(t, s14.Index, s14.Name, redactConfig(s14.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(14) },
			},
			// Step 15: Reset de-register without invoking another backend action.
			{
				Config: func() string {
					*stepCount++
					s15.Index = *stepCount
					s15.Name = fmt.Sprintf("%s_%d_reset_de_register", t.Name(), s15.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"de_register": false,
					})
					helper.GetTFConfigWithSingleResource(s15.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s15.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(15)
					helper.LogStep(t, s15.Index, s15.Name, redactConfig(s15.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(15) },
			},
			// Step 16: Re-register only so the fabric is reachable before cleanup.
			{
				Config: func() string {
					*stepCount++
					s16.Index = *stepCount
					s16.Name = fmt.Sprintf("%s_%d_re_register_only", t.Name(), s16.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"re_register": true,
					})
					helper.GetTFConfigWithSingleResource(s16.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s16.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(16)
					helper.LogStep(t, s16.Index, s16.Name, redactConfig(s16.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(16) },
			},
			// Step 17: Reset re-register without invoking another backend action.
			{
				Config: func() string {
					*stepCount++
					s17.Index = *stepCount
					s17.Name = fmt.Sprintf("%s_%d_reset_re_register", t.Name(), s17.Index)

					helper.ModifyFabricAciObject(&fabricRsc, map[string]interface{}{
						"re_register": false,
					})
					helper.GetTFConfigWithSingleResource(s17.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s17.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(17)
					helper.LogStep(t, s17.Index, s17.Name, redactConfig(s17.Cfg))
				},
				Check:         fabricAciStateCheck(*fabricRsc),
				PostApplyFunc: func() { completeStep(17) },
			},
			// Step 18: Delete outside Terraform, then disable plan refresh so the
			// provider Delete receives the stale state and handles HTTP 404.
			{
				Config: func() string {
					*stepCount++
					s18.Index = *stepCount
					s18.Name = fmt.Sprintf("%s_%d_destroy_after_external_delete", t.Name(), s18.Index)

					helper.GetTFConfigWithSingleResource(s18.Name, *x,
						[]interface{}{fabricRsc}, &tfConfig)

					s18.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() {
					requirePreviousStep(18)
					helper.LogStep(t, s18.Index, s18.Name, redactConfig(s18.Cfg))
					if err := os.Unsetenv(forceEnvironmentVariable); err != nil {
						t.Fatalf("Failed to unset %s before final destroy: %v", forceEnvironmentVariable, err)
					}
					deleteFabricAciOutsideTerraform(t, fabricName, username, password)

					noRefreshPlanArgs := "-refresh=false"
					if terraformPlanArgsSet && strings.TrimSpace(terraformPlanArgs) != "" {
						noRefreshPlanArgs = strings.TrimSpace(terraformPlanArgs) + " -refresh=false"
					}
					if err := os.Setenv("TF_CLI_ARGS_plan", noRefreshPlanArgs); err != nil {
						t.Fatalf("Failed to disable Terraform refresh before final destroy: %v", err)
					}
				},
				Destroy: true,
				PostApplyFunc: func() {
					restoreTerraformPlanArgs()
					completeStep(18)
				},
			},
		},
	})
}
