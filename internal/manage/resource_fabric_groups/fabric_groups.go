// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_groups

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

const ModuleKey = "manage"

var (
	_ resource.Resource                = &fabricGroupsResource{}
	_ resource.ResourceWithConfigure   = &fabricGroupsResource{}
	_ resource.ResourceWithImportState = &fabricGroupsResource{}
)

func NewFabricGroupsResource() resource.Resource {
	return &fabricGroupsResource{}
}

type fabricGroupsResource struct {
	manageClient *manage.NexusDashboardManage
}

func (r *fabricGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_groups"
}

func (r *fabricGroupsResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = FabricGroupsResourceSchema(ctx)
}

func (r *fabricGroupsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.manageClient = nil
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	manageModule := client.GetModule(ModuleKey)
	if manageModule == nil {
		resp.Diagnostics.AddError("Manage Module Not Found", "The manage module was not registered with the provider.")
		return
	}

	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok || manageClient == nil {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule),
		)
		return
	}
	if manageClient.ApiClient == nil {
		resp.Diagnostics.AddError("Manage Client Not Configured", "The manage module does not contain an API client.")
		return
	}

	r.manageClient = manageClient
}

func (r *fabricGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FabricGroupsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	validateFabricGroupConfiguration(ctx, req.Config, plan.FabricType, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating fabric group", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
		"fabric_type": plan.FabricType.ValueString(),
	})

	resource_fabric_common.RscCreateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *fabricGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FabricGroupsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	state.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading fabric group", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	previous := state
	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	state.normalizeOptionalState(previous)
	state.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *fabricGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FabricGroupsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	validateFabricGroupConfiguration(ctx, req.Config, plan.FabricType, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating fabric group", map[string]interface{}{
		"fabric_name": plan.FabricName.ValueString(),
		"fabric_type": plan.FabricType.ValueString(),
	})

	resource_fabric_common.RscUpdateFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &plan)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *fabricGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FabricGroupsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || !r.configured(&resp.Diagnostics) {
		return
	}
	state.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting fabric group", map[string]interface{}{
		"fabric_name": state.FabricName.ValueString(),
	})

	resource_fabric_common.RscDeleteFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
}

func (r *fabricGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if strings.TrimSpace(req.ID) == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "The import ID must be a fabric group name.")
		return
	}
	if !r.configured(&resp.Diagnostics) {
		return
	}

	state := FabricGroupsModel{
		Id:         types.StringValue(req.ID),
		FabricName: types.StringValue(req.ID),
	}

	tflog.Debug(ctx, "Importing fabric group", map[string]interface{}{
		"fabric_name": req.ID,
	})

	found := resource_fabric_common.RscGetFabric(ctx, r.manageClient.ApiClient, &resp.Diagnostics, &state)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError(
			"Error Importing Fabric Group",
			fmt.Sprintf("Could not import nd_fabric_groups with id %q: resource not found", req.ID),
		)
		return
	}

	state.setIDFromFabricName(&resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *fabricGroupsResource) configured(dg *diag.Diagnostics) bool {
	if r.manageClient == nil || r.manageClient.ApiClient == nil {
		dg.AddError("Manage Client Not Configured", "The fabric group resource requires a configured manage API client.")
		return false
	}
	return true
}

// The generated converter omits the Terraform-only id. After shared CRUD reads
// the API response, FabricName contains the returned name used to hydrate id.
func (v *FabricGroupsModel) setIDFromFabricName(dg *diag.Diagnostics) {
	if v == nil || v.FabricName.IsNull() || v.FabricName.IsUnknown() || v.FabricName.ValueString() == "" {
		dg.AddError("Invalid Fabric Group Name", "A non-empty fabric group name is required to populate the resource ID.")
		return
	}
	v.Id = v.FabricName
}
