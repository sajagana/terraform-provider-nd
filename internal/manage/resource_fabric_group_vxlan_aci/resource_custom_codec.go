// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_group_vxlan_aci

import (
	"context"
	"fmt"
	"time"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource_fabric_common.FabricModel               = (*FabricGroupVxlanAciModel)(nil)
	_ resource_fabric_common.FabricPreMarshal          = (*FabricGroupVxlanAciModel)(nil)
	_ resource_fabric_common.FabricCRUDOptionsProvider = (*FabricGroupVxlanAciModel)(nil)
)

func (v *FabricGroupVxlanAciModel) GetFabricName() string {
	if !v.Id.IsNull() && !v.Id.IsUnknown() && v.Id.ValueString() != "" {
		return v.Id.ValueString()
	}
	return v.FabricName.ValueString()
}

func (v *FabricGroupVxlanAciModel) GetFabricType() string {
	return "vxlanAci"
}

// CRUDOptions enables shared lifecycle extensions only for this resource.
func (v *FabricGroupVxlanAciModel) CRUDOptions() resource_fabric_common.FabricCRUDOptions {
	return resource_fabric_common.FabricCRUDOptions{
		APIOptions:               &ndapi.APIOptions{DisablePayloadLog: true},
		ReadbackTimeout:          2 * time.Minute,
		DeleteNotFoundOK:         true,
		EscapeFabricName:         true,
		PruneEmptyObjects:        true,
		SetModelData:             v.setModelDataPreservingState,
		ModifyPayload:            v.modifyPayload,
		ReadbackState:            v.readbackState,
	}
}

// PreMarshal supplies hidden discriminators and excludes read-only status.
// All attribute conversion remains in the generated common-model codec.
func (v *FabricGroupVxlanAciModel) PreMarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	data.Category = "fabricGroup"
	data.Management.FabricType = "vxlanAci"
	data.Management.SecurityGroupStatus = ""
}

func (v *FabricGroupVxlanAciModel) setModelDataPreservingState(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if data.Category != "fabricGroup" || data.Management.FabricType != "vxlanAci" {
		diags.AddError("Invalid Fabric Group Type", fmt.Sprintf("Object %q has category %q and management.type %q; expected fabricGroup and vxlanAci.", v.GetFabricName(), data.Category, data.Management.FabricType))
		return diags
	}
	if data.FabricName == "" || data.FabricName != v.GetFabricName() {
		diags.AddError("Invalid Fabric Group Response", fmt.Sprintf("Response name %q does not match requested fabric-group name %q.", data.FabricName, v.GetFabricName()))
		return diags
	}
	current := *v
	diags.Append(current.SetModelData(data)...)
	if diags.HasError() {
		return diags
	}
	// GET cannot restore the configured key. Retain it rather than accepting a
	// masked or re-encrypted value; import leaves this write-only attribute null.
	current.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey
	if len(data.Management.RouteServerCollection) == 0 && !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0 {
		// The generated codec normalizes no route servers to null. Retain an
		// explicitly configured empty list so readback agrees with the plan.
		current.RouteServerCollection = v.RouteServerCollection
	}
	*v = current
	return diags
}

func (v *FabricGroupVxlanAciModel) modifyPayload(_ context.Context, payload map[string]interface{}, create bool) error {
	if (!create && v.RouteServerCollection.IsNull()) || (!v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0) {
		management, ok := payload["management"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("fabric-group payload is missing its management object")
		}
		// omitempty drops empty slices; send [] explicitly to clear route servers.
		management["routeServerCollection"] = []interface{}{}
	}
	return nil
}

func (v *FabricGroupVxlanAciModel) readbackState(ctx context.Context) (types.Object, diag.Diagnostics) {
	var state types.Object
	var diags diag.Diagnostics
	// Validate nested conversion before generated GetModelData, which cannot
	// return diagnostics for an invalid list value.
	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		var servers []RouteServerCollectionValue
		diags.Append(v.RouteServerCollection.ElementsAs(ctx, &servers, false)...)
		if diags.HasError() {
			return state, diags
		}
	}
	// Computed status can change independently of the requested settings.
	snapshot := *v
	snapshot.SecurityGroupStatus = types.StringUnknown()
	diags.Append(tfsdk.ValueFrom(ctx, &snapshot, FabricGroupVxlanAciResourceSchema(ctx).Type(), &state)...)
	return state, diags
}
