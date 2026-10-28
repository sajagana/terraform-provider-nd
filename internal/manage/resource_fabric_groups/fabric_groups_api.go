// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_groups

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// Provider polling limits, not API-advertised completion guarantees.
const (
	fabricGroupReadbackTimeout  = 5 * time.Minute
	fabricGroupReadbackInterval = 2 * time.Second
)

var (
	_ resource_fabric_common.FabricRequestOptions = (*FabricGroupsModel)(nil)
	_ resource_fabric_common.FabricWriteReadback  = (*FabricGroupsModel)(nil)
	_ resource_fabric_common.FabricDeleteNotFound = (*FabricGroupsModel)(nil)
	_ resource_fabric_common.FabricModel          = (*fabricGroupCreateModel)(nil)
)

// Track accepted creates without adding handwritten fields or conversion
// methods to the generated Terraform model. All generated methods are promoted.
type fabricGroupCreateModel struct {
	*FabricGroupsModel
	accepted bool
}

func (v *fabricGroupCreateModel) ReadFabricAfterWrite(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, expected *resource_fabric_common.NDFCFabricCommonModel) bool {
	v.accepted = true
	return v.FabricGroupsModel.ReadFabricAfterWrite(ctx, client, dg, expected)
}

// Disable both request and response body logging: GET can return secret keys.
// Options are local to each API call; no shared client settings are modified.
func (*FabricGroupsModel) FabricAPIOptions(ctx context.Context) *ndapi.APIOptions {
	return &ndapi.APIOptions{DisablePayloadLog: true, Context: ctx}
}

func (*FabricGroupsModel) DeleteMissingFabricIsSuccess() bool {
	return true
}

// POST/PUT may return 202 with a statusUrl whose response schema is not defined
// in the canonical Manage spec. Poll the documented GET instead, requiring
// visible requested settings and no reported pending security-group transition.
// Only an exact 404 or an unconverged successful GET is retried; other errors
// are returned immediately. Ordinary Read/Import keep their normal 404 behavior.
func (v *FabricGroupsModel) ReadFabricAfterWrite(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, expected *resource_fabric_common.NDFCFabricCommonModel) bool {
	if v == nil || client == nil || expected == nil {
		dg.AddError("Invalid Input", "Fabric group readback requires a model, API client, and expected payload.")
		return false
	}
	expectedJSON, err := v.MarshalFabricJSON(expected)
	if err != nil {
		dg.AddError("Error Preparing Fabric Group Readback", "Could not prepare the expected fabric group settings.")
		return false
	}
	var desired fabricGroupJSON
	if err := json.Unmarshal(expectedJSON, &desired); err != nil {
		dg.AddError("Error Preparing Fabric Group Readback", "Could not decode the expected fabric group settings.")
		return false
	}

	pollCtx, cancel := context.WithTimeout(ctx, fabricGroupReadbackTimeout)
	defer cancel()
	ticker := time.NewTicker(fabricGroupReadbackInterval)
	defer ticker.Stop()

	for pollCtx.Err() == nil {
		// Failed or unconverged responses must not replace the planned values.
		candidate := *v
		var attemptDiagnostics diag.Diagnostics
		found := resource_fabric_common.RscGetFabric(pollCtx, client, &attemptDiagnostics, &candidate)
		if pollCtx.Err() != nil {
			break
		}
		if attemptDiagnostics.HasError() {
			dg.Append(attemptDiagnostics...)
			return false
		}
		if found {
			candidate.normalizeOptionalState(*v)
			currentJSON, err := candidate.MarshalFabricJSON(candidate.GetModelData())
			if err != nil {
				dg.AddError("Error Reading Fabric Group", "Could not prepare the fabric group settings returned by the API.")
				return false
			}
			var current fabricGroupJSON
			if err := json.Unmarshal(currentJSON, &current); err != nil {
				dg.AddError("Error Reading Fabric Group", "Could not decode the fabric group settings returned by the API.")
				return false
			}
			pending := false
			switch candidate.SecurityGroupStatus.ValueString() {
			case "enablePending", "enablePendingStrict", "enablePendingLoose", "disablePending":
				pending = true
			}
			if !pending && fabricGroupReadbackMatches(desired, current) {
				dg.Append(attemptDiagnostics...)
				*v = candidate
				return true
			}
		}

		tflog.Debug(pollCtx, "Waiting for fabric group readback", map[string]interface{}{
			"fabric_name": expected.FabricName,
		})
		select {
		case <-pollCtx.Done():
		case <-ticker.C:
		}
	}

	dg.AddError(
		"Error Waiting for Fabric Group",
		fmt.Sprintf("The write for fabric group %q was accepted, but its settings did not converge before readback stopped: %s. Refresh the resource before retrying.", expected.FabricName, pollCtx.Err().Error()),
	)
	return false
}

func fabricGroupReadbackMatches(expected, current fabricGroupJSON) bool {
	if expected.Category != current.Category || expected.FabricName != current.FabricName {
		return false
	}
	for name, value := range expected.Management {
		// Secret keys need not be echoed to prove readback. When a response omits
		// one, the group decoder retains its prior known Terraform value.
		if name == "multisiteInterConnectBgpKey" || name == "cloudSecKey" {
			continue
		}
		if !bytes.Equal(value, current.Management[name]) {
			return false
		}
	}
	return true
}
