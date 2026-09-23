// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package datasource_fabric_aci

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage"
	manageapi "terraform-provider-nd/internal/manage/api"
	"terraform-provider-nd/internal/registry"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// ModuleKey is the key used to get the manage module from the provider.
const ModuleKey = "manage"

var (
	_ datasource.DataSource              = &fabricAciDataSource{}
	_ datasource.DataSourceWithConfigure = &fabricAciDataSource{}
)

// NewFabricAciDataSource returns a Fabric ACI datasource.
func NewFabricAciDataSource() datasource.DataSource {
	return &fabricAciDataSource{}
}

type fabricAciDataSource struct {
	manageClient *manage.NexusDashboardManage
}

// fabricAciDataSourceResponse represents the response returned by the
// /infra/clusters/{name} endpoint. The API returns the fabric name at
// spec.name, while the generated Terraform model stores it at spec.aci.name.
type fabricAciDataSourceResponse struct {
	Spec   fabricAciDataSourceResponseSpec `json:"spec,omitempty"`
	Status NDFCStatusValue                 `json:"status,omitempty"`
}

type fabricAciDataSourceResponseSpec struct {
	NDFCSpecValue
	FabricName string `json:"name,omitempty"`
}

// Metadata returns the datasource type name.
func (d *fabricAciDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fabric_aci"
}

// Schema defines the schema for the datasource.
func (d *fabricAciDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = FabricAciDataSourceSchema(ctx)
}

// Configure adds the provider-configured manage client to the datasource.
func (d *fabricAciDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(registry.ClientProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected registry.ClientProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	manageModule := client.GetModule(ModuleKey)
	if manageModule == nil {
		resp.Diagnostics.AddError(
			"Manage Module Not Found",
			"The manage module was not registered with the provider.",
		)
		return
	}

	manageClient, ok := manageModule.(*manage.NexusDashboardManage)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Manage Module Type",
			fmt.Sprintf("Expected *manage.NexusDashboardManage, got: %T. Please report this issue to the provider developers.", manageModule),
		)
		return
	}

	d.manageClient = manageClient
}

// Read retrieves an APIC fabric by name and saves it in Terraform state.
func (d *fabricAciDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	log.Printf("[DEBUG] Start read of datasource: nd_fabric_aci")

	var data FabricAciModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// A required attribute can still be unknown when its value comes from
	// another resource. Do not let an unknown value become an empty cluster
	// name, which would change the request from a by-name GET to a list GET.
	if data.FabricName.IsNull() || data.FabricName.IsUnknown() {
		resp.Diagnostics.AddError(
			"Fabric Name Required",
			"The fabric_name attribute must contain a known value to read an APIC fabric.",
		)
		return
	}

	fabricName := data.FabricName.ValueString()
	log.Printf("[DEBUG] Reading Fabric ACI: fabric_name=%s", fabricName)

	fabricAPI := manageapi.NewFabricAciAPI(d.manageClient.ApiClient, ndapi.DefaultFabric)
	fabricAPI.ClusterName = fabricName

	respData, err := fabricAPI.Get()
	if err != nil {
		if strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError(
				"Error Reading Fabric ACI",
				fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q: resource not found", fabricName),
			)
			return
		}

		resp.Diagnostics.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q, unexpected error: %s %s", fabricName, err.Error(), string(respData)),
		)
		return
	}

	if respData == nil {
		resp.Diagnostics.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not read nd_fabric_aci with fabric_name %q: resource not found", fabricName),
		)
		return
	}

	var fabricResp fabricAciDataSourceResponse
	if err := json.Unmarshal(respData, &fabricResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Fabric ACI",
			fmt.Sprintf("Could not unmarshal nd_fabric_aci response with fabric_name %q, unexpected error: %s", fabricName, err.Error()),
		)
		return
	}

	modelData := NDFCFabricAciModel{
		Spec:   fabricResp.Spec.NDFCSpecValue,
		Status: fabricResp.Status,
	}
	if modelData.Spec.Aci.FabricName == "" {
		modelData.Spec.Aci.FabricName = fabricResp.Spec.FabricName
	}
	if modelData.Spec.Aci.FabricName == "" {
		modelData.Spec.Aci.FabricName = fabricName
	}
	modelData.Spec.Aci.Telemetry.Network = normalizeTelemetryNetworkState(modelData.Spec.Aci.Telemetry.Network)

	resp.Diagnostics.Append(data.SetModelData(&modelData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("[DEBUG] End read of datasource nd_fabric_aci with fabric_name=%s", fabricName)
}

func normalizeTelemetryNetworkState(network string) string {
	switch network {
	case "inband", "inBand":
		return "inband"
	case "outband", "outOfBand":
		return "outband"
	default:
		return network
	}
}
