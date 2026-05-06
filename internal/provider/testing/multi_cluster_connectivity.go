// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package testing

import (
	"terraform-provider-nd/internal/infra/resource_multi_cluster_connectivity"
)

// defaultMultiClusterConnectivityValues returns sensible defaults for a
// multi cluster connectivity (ND) resource. Tests may override any of these
// via the overrides map passed to GenerateMultiClusterConnectivityObject.
func defaultMultiClusterConnectivityValues() map[string]interface{} {
	return map[string]interface{}{
		"cluster_type": "ND",
		"login_domain": "local",
	}
}

// GenerateMultiClusterConnectivityObject creates a multi cluster connectivity
// model object for testing.
//
// clusterName, hostname, username and password are mandatory identifiers /
// required attributes. overrides lets each test supply unique values for any
// other field. Any key not present in overrides gets the value from
// defaultMultiClusterConnectivityValues().
func GenerateMultiClusterConnectivityObject(
	obj **resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	clusterName string, hostname string, username string, password string,
	overrides map[string]interface{},
) {
	mcc := new(resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel)

	mcc.Spec.ClusterName = clusterName
	mcc.Spec.Hostname = hostname
	mcc.Spec.Credentials.Username = username
	mcc.Spec.Credentials.Password = password

	merged := defaultMultiClusterConnectivityValues()
	for k, v := range overrides {
		merged[k] = v
	}

	applyMultiClusterConnectivityValues(mcc, merged)

	*obj = mcc
}

// ModifyMultiClusterConnectivityObject modifies fields on an existing model.
// Uses the same key set as GenerateMultiClusterConnectivityObject overrides.
func ModifyMultiClusterConnectivityObject(
	obj **resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	values map[string]interface{},
) {
	mcc := *obj
	if mcc == nil {
		return
	}

	applyMultiClusterConnectivityValues(mcc, values)

	*obj = mcc
}

// applyMultiClusterConnectivityValues is the shared engine that sets fields on
// the model from a key-value map. Used by both Generate and Modify.
func applyMultiClusterConnectivityValues(
	mcc *resource_multi_cluster_connectivity.NDFCMultiClusterConnectivityModel,
	values map[string]interface{},
) {
	for key, val := range values {
		switch key {
		case "cluster_name":
			mcc.Spec.ClusterName = val.(string)
		case "cluster_type":
			mcc.Spec.ClusterType = val.(string)
		case "hostname":
			mcc.Spec.Hostname = val.(string)
		case "username":
			mcc.Spec.Credentials.Username = val.(string)
		case "password":
			mcc.Spec.Credentials.Password = val.(string)
		case "login_domain":
			mcc.Spec.Credentials.LoginDomain = val.(string)
		case "multi_cluster_login_domain":
			mcc.Spec.Nd.MultiClusterLoginDomain = val.(string)
		}
	}
}
