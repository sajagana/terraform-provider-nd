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
	"strings"

	"terraform-provider-nd/internal/manage"
	"terraform-provider-nd/internal/manage/resource_fabric_common"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &fabricGroupVxlanAciResource{}
	_ resource.ResourceWithConfigure   = &fabricGroupVxlanAciResource{}
	_ resource.ResourceWithImportState = &fabricGroupVxlanAciResource{}
)

// NewFabricGroupVxlanAciResource returns the dedicated VXLAN-ACI fabric group resource.
func NewFabricGroupVxlanAciResource() resource.Resource {
	return &fabricGroupVxlanAciResource{}
}

type fabricGroupVxlanAciResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *fabricGroupVxlanAciResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_group_vxlan_aci"
}

func (r *fabricGroupVxlanAciResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricGroupVxlanAciResourceSchema(ctx)
}

func (r *fabricGroupVxlanAciResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}

	module := client.GetModule(ModuleKey)
	if module == nil {
		resp.Diagnostics.AddError("Manage Module Not Found", "The manage module was not registered with the provider.")
		return
	}
	manageClient, ok := module.(*manage.NexusDashboardManage)
	if !ok || manageClient == nil || manageClient.ApiClient == nil {
		resp.Diagnostics.AddError("Invalid Manage Module", "The manage module does not contain a configured Nexus Dashboard API client.")
		return
	}
	r.manageClient = manageClient
}

func (r *fabricGroupVxlanAciResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics, plan.FabricName.ValueString()) {
		return
	}
	plan.Id = plan.FabricName
	resource_fabric_common.RscCreateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, "Created VXLAN-ACI fabric group", map[string]interface{}{"id": plan.Id.ValueString()})
	}
}

func (r *fabricGroupVxlanAciResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics, state.GetFabricName()) {
		return
	}
	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		tflog.Warn(ctx, "VXLAN-ACI fabric group returned HTTP 404; removing resource from state", map[string]interface{}{"id": state.Id.ValueString()})
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *fabricGroupVxlanAciResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics, plan.FabricName.ValueString()) {
		return
	}
	if !plan.FabricName.Equal(state.FabricName) {
		resp.Diagnostics.AddError("Invalid Fabric Group Update", "Changing fabric_name requires replacing the resource.")
		return
	}
	plan.Id = state.Id
	resource_fabric_common.RscUpdateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		tflog.Debug(ctx, "Updated VXLAN-ACI fabric group", map[string]interface{}{"id": plan.Id.ValueString()})
	}
}

func (r *fabricGroupVxlanAciResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricGroupVxlanAciModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics, state.GetFabricName()) {
		return
	}
	resource_fabric_common.RscDeleteFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
}

func (r *fabricGroupVxlanAciResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if !r.configured(&resp.Diagnostics, req.ID) {
		return
	}
	state := FabricGroupVxlanAciModel{
		Id:         types.StringValue(req.ID),
		FabricName: types.StringValue(req.ID),
	}
	// Shared read invokes the model hook to validate category and management.type.
	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Error Importing Fabric Group VXLAN-ACI", fmt.Sprintf("Could not import nd fabric group vxlan aci with id %q: resource not found", req.ID))
		return
	}
	// Import cannot recover the API-omitted BGP authentication key.
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *fabricGroupVxlanAciResource) configured(dg *diag.Diagnostics, name string) bool {
	if r.manageClient == nil || r.manageClient.ApiClient == nil {
		dg.AddError("Unconfigured Fabric Group Resource", "The manage API client must be configured before operating on a fabric group.")
		return false
	}
	if strings.TrimSpace(name) == "" {
		dg.AddError("Invalid Fabric Group Identifier", "The fabric-group name must not be empty.")
		return false
	}
	return true
}
