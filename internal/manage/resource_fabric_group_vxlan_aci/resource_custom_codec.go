// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_group_vxlan_aci

import (
	"context"
	"encoding/json"
	"fmt"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource_fabric_common.FabricModel         = (*fabricGroupModel)(nil)
	_ resource_fabric_common.FabricPreMarshal    = (*FabricGroupVxlanAciModel)(nil)
	_ resource_fabric_common.FabricPostUnmarshal = (*fabricGroupModel)(nil)
)

const (
	fabricGroupCategory = "fabricGroup"
	fabricGroupType     = "vxlanAci"
)

func (v *FabricGroupVxlanAciModel) GetFabricName() string {
	if !v.Id.IsNull() && !v.Id.IsUnknown() && v.Id.ValueString() != "" {
		return v.Id.ValueString()
	}
	return v.FabricName.ValueString()
}

func (v *FabricGroupVxlanAciModel) GetFabricType() string {
	return fabricGroupType
}

// PreMarshal supplies hidden discriminators. The generated codec owns
// securityGroupTag conversion; no custom JSON methods are needed for that field.
func (v *FabricGroupVxlanAciModel) PreMarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	data.Category = fabricGroupCategory
	data.Management.FabricType = fabricGroupType
	data.Management.SecurityGroupStatus = ""
}

func (v *FabricGroupVxlanAciModel) validateRouteServerConversion(ctx context.Context) diag.Diagnostics {
	var diags diag.Diagnostics
	// The generated collection codec cannot return conversion diagnostics and
	// panics on invalid elements. Check conversion before calling GetModelData.
	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		var servers []RouteServerCollectionValue
		diags.Append(v.RouteServerCollection.ElementsAs(ctx, &servers, false)...)
	}
	return diags
}

// MarshalFabricPayload selects the fabric-group envelope from the shared model.
func (v *FabricGroupVxlanAciModel) MarshalFabricPayload(data *resource_fabric_common.NDFCFabricCommonModel) ([]byte, error) {
	managementJSON, err := json.Marshal(data.Management)
	if err != nil {
		return nil, fmt.Errorf("could not encode fabric-group management settings")
	}
	var management map[string]json.RawMessage
	if err := json.Unmarshal(managementJSON, &management); err != nil {
		return nil, fmt.Errorf("could not decode the generated fabric-group request")
	}
	// Unset struct-valued management fields can be serialized as empty objects.
	for name, value := range management {
		if string(value) == "{}" {
			delete(management, name)
		}
	}
	return json.Marshal(struct {
		Name       string                     `json:"name"`
		Category   string                     `json:"category"`
		Management map[string]json.RawMessage `json:"management"`
	}{Name: data.FabricName, Category: data.Category, Management: management})
}

func validateFabricGroupIdentity(expectedName, name, category, fabricType string) error {
	if category != fabricGroupCategory || fabricType != fabricGroupType {
		return fmt.Errorf("object %q must have category %s and management.type %s", expectedName, fabricGroupCategory, fabricGroupType)
	}
	if name == "" || name != expectedName {
		return fmt.Errorf("the returned fabric-group name does not match the requested name %q", expectedName)
	}
	return nil
}

// fabricGroupModel adapts the generated model to shared CRUD for one operation.
type fabricGroupModel struct {
	*FabricGroupVxlanAciModel
	importing      bool
	identityOnly   bool
	writeContext   context.Context
	createAccepted bool
}

func (v *FabricGroupVxlanAciModel) FabricAPIOptions() *ndapi.APIOptions {
	return &ndapi.APIOptions{
		DisablePayloadLog:         true,
		AllowDeployInResourceName: true,
	}
}

func (v *fabricGroupModel) OnFabricCreateAccepted() {
	v.createAccepted = true
}

func (v *fabricGroupModel) PostUnmarshal(_ context.Context, data *resource_fabric_common.NDFCFabricCommonModel) {
	management := &data.Management
	// SetModelData preserves the key from the Terraform plan or state.
	management.MultisiteInterConnectBgpKey = ""
	// Keep omitted optional values null when GET echoes disabled backup settings.
	if v.importing || v.identityOnly || (management.ScheduledBackup != nil && *management.ScheduledBackup) {
		return
	}
	if v.ScheduledBackup.IsNull() {
		management.ScheduledBackup = nil
	}
	if v.ScheduledBackupTime.IsNull() {
		management.ScheduledBackupTime = ""
	}
}

func (v *fabricGroupModel) SetModelData(data *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Invalid Fabric Group Response", "The API returned no fabric-group model.")
		return diags
	}
	if err := validateFabricGroupIdentity(v.GetFabricName(), data.FabricName, data.Category, data.Management.FabricType); err != nil {
		diags.AddError("Invalid Fabric Group Response", err.Error())
		return diags
	}
	if v.identityOnly {
		return diags
	}

	if data.Management.SecurityGroupTag != nil {
		mode, ok := data.Management.SecurityGroupTag.(string)
		if !ok || (mode != "loose" && mode != "strict") {
			diags.AddError("Invalid Fabric Group Response", "securityGroupTag must be the string loose or strict for a VXLAN-ACI fabric group.")
			return diags
		}
	}
	current := *v.FabricGroupVxlanAciModel
	diags.Append(current.SetModelData(data)...)
	if diags.HasError() {
		return diags
	}
	current.Id = current.FabricName
	// GET can omit, mask, or re-encrypt the configured key. Import leaves it null.
	current.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey
	if len(data.Management.RouteServerCollection) == 0 && !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() && len(v.RouteServerCollection.Elements()) == 0 {
		current.RouteServerCollection = v.RouteServerCollection
	}
	if v.writeContext != nil {
		diags.Append(v.preserveOmittedPlanValues(&current)...)
		if diags.HasError() {
			return diags
		}
	}
	*v.FabricGroupVxlanAciModel = current
	return diags
}

// preserveOmittedPlanValues keeps API defaults from replacing planned nulls during apply.
func (v *fabricGroupModel) preserveOmittedPlanValues(current *FabricGroupVxlanAciModel) diag.Diagnostics {
	var diags diag.Diagnostics
	ctx := v.writeContext
	var planned, returned types.Object
	objectType := FabricGroupVxlanAciResourceSchema(ctx).Type()
	diags.Append(tfsdk.ValueFrom(ctx, v.FabricGroupVxlanAciModel, objectType, &planned)...)
	diags.Append(tfsdk.ValueFrom(ctx, current, objectType, &returned)...)
	if diags.HasError() {
		return diags
	}

	attributes := returned.Attributes()
	for name, value := range planned.Attributes() {
		if value.IsNull() && name != "security_group_status" {
			attributes[name] = value
		}
	}

	normalized, normalizationDiags := types.ObjectValue(planned.AttributeTypes(ctx), attributes)
	diags.Append(normalizationDiags...)
	if !diags.HasError() {
		diags.Append(tfsdk.ValueAs(ctx, normalized, current)...)
	}
	return diags
}
