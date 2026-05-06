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
	"testing"

	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
	helper "terraform-provider-nd/internal/provider/testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMultiClusterConnectivityResourceCRUD(t *testing.T) {
	mccCfg := helper.GetConfig("global").ND.MultiCluster

	x := &map[string]string{
		"RscType":  "nd_multi_cluster_connectivity",
		"RscName":  "mcc_test",
		"User":     helper.GetConfig("global").ND.User,
		"Password": helper.GetConfig("global").ND.Password,
		"Host":     helper.GetConfig("global").ND.URL,
		"Insecure": helper.GetConfig("global").ND.Insecure,
	}

	tfConfig := new(string)
	stepCount := new(int)
	*stepCount = 0

	mccRsc := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "global") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create a basic multi cluster connectivity (ND) resource
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.GenerateMultiClusterConnectivityObject(&mccRsc,
						mccCfg.ClusterName,
						mccCfg.Hostname,
						mccCfg.Username,
						mccCfg.Password,
						nil,
					)

					(*x)["RscName"] = "mcc_test"
					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{mccRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(
						"nd_multi_cluster_connectivity.mcc_test",
						*mccRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 2: Set explicit cluster_type and login_domain overrides
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyMultiClusterConnectivityObject(&mccRsc, map[string]interface{}{
						"cluster_type": "ND",
						"login_domain": "DefaultAuth",
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{mccRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(
						"nd_multi_cluster_connectivity.mcc_test",
						*mccRsc,
						path.Empty(),
					)...,
				),
			},
			// Step 3: Set the multi_cluster_login_domain
			{
				Config: func() string {
					*stepCount++
					tName := fmt.Sprintf("%s_%d", t.Name(), *stepCount)

					helper.ModifyMultiClusterConnectivityObject(&mccRsc, map[string]interface{}{
						"multi_cluster_login_domain": mccCfg.MultiClusterLoginDomain,
					})

					helper.GetTFConfigWithSingleResource(tName, *x,
						[]interface{}{mccRsc}, &tfConfig)

					return *tfConfig
				}(),
				Check: resource.ComposeTestCheckFunc(
					MultiClusterConnectivityModelHelperStateCheck(
						"nd_multi_cluster_connectivity.mcc_test",
						*mccRsc,
						path.Empty(),
					)...,
				),
			},
		},
	})
}
