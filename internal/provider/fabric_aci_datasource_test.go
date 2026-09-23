// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"terraform-provider-nd/internal/manage/resource_fabric_aci"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFabricAciDataSource(t *testing.T) {
	const (
		resourceName   = "nd_fabric_aci.fabric_test"
		dataSourceName = "data.nd_fabric_aci.fabric_test"
		hostname       = "1.1.1.1"
		username       = "admin"
		password       = "**********"
	)

	cfg := helper.GetConfig("global")
	fabricName := "tf-" + acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	missingFabricName := fabricName + "-missing"

	x := &map[string]string{
		"RscType":  "nd_fabric_aci",
		"RscName":  "fabric_test",
		"User":     cfg.ND.User,
		"Password": cfg.ND.Password,
		"Host":     cfg.ND.URL,
		"Insecure": cfg.ND.Insecure,
	}

	tfConfig := new(string)
	fabric := new(resource_fabric_aci.NDFCFabricAciModel)
	helper.GenerateFabricAciObject(
		&fabric,
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
			"orchestration_status": "disabled",
		},
	)

	fabricDataSource := &helper.FabricAciDataSourceTestData{
		RscName:   "fabric_test",
		FabricName: fabricName,
		DependsOn: resourceName,
	}
	missingFabricDataSource := &helper.FabricAciDataSourceTestData{
		RscName:   "missing_fabric",
		FabricName: missingFabricName,
	}

	redactConfig := func(config string) string {
		config = strings.ReplaceAll(config, password, "<redacted>")
		if cfg.ND.Password != "" {
			config = strings.ReplaceAll(config, cfg.ND.Password, "<redacted>")
		}
		return config
	}

	resourceChecks := append(
		FabricAciModelHelperStateCheck(resourceName, *fabric, path.Empty()),
		resource.TestCheckResourceAttr(resourceName, "id", fabricName),
		resource.TestCheckResourceAttr(resourceName, "latitude", strconv.FormatFloat(*fabric.Spec.Location.Latitude, 'f', -1, 64)),
		resource.TestCheckResourceAttr(resourceName, "longitude", strconv.FormatFloat(*fabric.Spec.Location.Longitude, 'f', -1, 64)),
		resource.TestCheckResourceAttrSet(resourceName, "state"),
		resource.TestCheckResourceAttrSet(resourceName, "last_update_message"),
	)

	matchingChecks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "fabric_name", resourceName, "fabric_name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "hostname", resourceName, "hostname"),
		resource.TestCheckResourceAttrPair(dataSourceName, "latitude", resourceName, "latitude"),
		resource.TestCheckResourceAttrPair(dataSourceName, "longitude", resourceName, "longitude"),
		resource.TestCheckResourceAttrPair(dataSourceName, "license_tier", resourceName, "license_tier"),
		resource.TestCheckResourceAttrPair(dataSourceName, "orchestration_status", resourceName, "orchestration_status"),
		resource.TestCheckResourceAttrPair(dataSourceName, "security_domain", resourceName, "security_domain"),
		resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
		resource.TestCheckResourceAttrPair(dataSourceName, "last_update_message", resourceName, "last_update_message"),
		resource.TestCheckResourceAttrPair(dataSourceName, "telemetry.status", resourceName, "telemetry.status"),
		resource.TestCheckResourceAttrPair(dataSourceName, "telemetry.network", resourceName, "telemetry.network"),
		resource.TestCheckResourceAttrPair(dataSourceName, "telemetry.epg", resourceName, "telemetry.epg"),
		resource.TestCheckResourceAttrPair(dataSourceName, "telemetry.streaming_protocol", resourceName, "telemetry.streaming_protocol"),
	}

	s1 := &helper.StepInfo{}
	s2 := &helper.StepInfo{}
	s3 := &helper.StepInfo{}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s1.Index = 1
					s1.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Create the Fabric ACI used by datasource lookups",
					)

					helper.GetTFConfigWithSingleResource(
						s1.Name,
						*x,
						[]interface{}{fabric},
						&tfConfig,
					)
					s1.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s1.Index, s1.Name, redactConfig(s1.Cfg)) },
				Check:     resource.ComposeTestCheckFunc(resourceChecks...),
			},
			{
				Config: func() string {
					s2.Index = 2
					s2.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Read the existing Fabric ACI through the datasource",
					)

					helper.GetTFConfigWithSingleResource(
						s2.Name,
						*x,
						[]interface{}{fabric, fabricDataSource},
						&tfConfig,
					)
					s2.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s2.Index, s2.Name, redactConfig(s2.Cfg)) },
				Check:     resource.ComposeTestCheckFunc(matchingChecks...),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: func() string {
					s3.Index = 3
					s3.Name = fmt.Sprintf(
						"%s - %s",
						t.Name(),
						"Reject a datasource lookup for a missing Fabric ACI",
					)

					helper.GetTFConfigWithSingleResource(
						s3.Name,
						*x,
						[]interface{}{missingFabricDataSource},
						&tfConfig,
					)
					s3.Cfg = *tfConfig
					return *tfConfig
				}(),
				PreConfig: func() { helper.LogStep(t, s3.Index, s3.Name, redactConfig(s3.Cfg)) },
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(
						`Could not read nd_fabric_aci with fabric_name\s+%q:\s+resource\s+not\s+found`,
						missingFabricName,
					),
				),
			},
		},
	})
}
