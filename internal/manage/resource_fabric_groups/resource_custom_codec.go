// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_groups

import (
	"encoding/json"
	"errors"
	"fmt"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

var (
	_ resource_fabric_common.FabricModel           = (*FabricGroupsModel)(nil)
	_ resource_fabric_common.FabricJSONMarshaler   = (*FabricGroupsModel)(nil)
	_ resource_fabric_common.FabricJSONUnmarshaler = (*FabricGroupsModel)(nil)
)

// This defined type keeps group JSON overrides separate from the shared model.
type fabricManagementJSON resource_fabric_common.NDFCManagementValue

// The explicit field overrides the embedded management model's Boolean JSON
// field. The remaining management fields retain their generated JSON mappings.
type fabricGroupManagementJSON struct {
	*fabricManagementJSON
	SecurityGroupTag string `json:"securityGroupTag,omitempty"`
}

type fabricGroupJSON struct {
	Category   string                    `json:"category"`
	FabricName string                    `json:"name"`
	Management map[string]json.RawMessage `json:"management"`
}

// GetFabricName returns the fabric group name from the Terraform model.
func (v *FabricGroupsModel) GetFabricName() string {
	return v.FabricName.ValueString()
}

// GetFabricType keeps group handler dispatch separate from individual fabrics.
// The management type (vxlan, vxlanAci, etc.) remains in the generated payload.
func (*FabricGroupsModel) GetFabricType() string {
	return "fabricGroup"
}

// MarshalFabricJSON uses generated common fields while exposing the group's
// private String as management.securityGroupTag instead of the fabric Boolean.
func (*FabricGroupsModel) MarshalFabricJSON(data *resource_fabric_common.NDFCFabricCommonModel) ([]byte, error) {
	if data == nil {
		return nil, errors.New("cannot marshal fabric group from a nil common model")
	}

	groupType := data.Management.FabricType
	if data.Category != "fabricGroup" {
		return nil, fmt.Errorf("expected fabric group category %q, got %q", "fabricGroup", data.Category)
	}
	if data.FabricName == "" {
		return nil, errors.New("cannot marshal fabric group without a name")
	}
	if !validFabricGroupType(groupType) {
		return nil, fmt.Errorf("unsupported fabric group management type %q", groupType)
	}
	if groupType == "vxlanAci" && data.FabricGroupCodec.SecurityGroupTag == "off" {
		return nil, errors.New("vxlanAci supports only loose or strict securityGroupTag enforcement")
	}

	management := fabricManagementJSON(data.Management)
	encoded, err := json.Marshal(fabricGroupManagementJSON{
		fabricManagementJSON: &management,
		SecurityGroupTag:      data.FabricGroupCodec.SecurityGroupTag,
	})
	if err != nil {
		return nil, fmt.Errorf("encode fabric group management: %w", err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &fields); err != nil {
		return nil, fmt.Errorf("prepare fabric group management: %w", err)
	}
	// Generated GetModelData distinguishes an explicit empty list (non-nil)
	// from omission (nil). Preserve [] despite the common field's omitempty tag.
	if data.Management.RouteServerCollection != nil && len(data.Management.RouteServerCollection) == 0 {
		fields["routeServerCollection"] = json.RawMessage("[]")
	}
	fields = filterFabricGroupManagement(fields, groupType, false)
	for _, allocationRange := range fabricGroupAllocationRanges {
		raw, present := fields[allocationRange.apiName]
		if !present {
			continue
		}
		var value string
		if err := json.Unmarshal(raw, &value); err != nil {
			return nil, fmt.Errorf("decode fabric group %s: %w", allocationRange.apiName, err)
		}
		if err := validateFabricGroupRange(value, allocationRange.minimum, allocationRange.maximum); err != nil {
			return nil, fmt.Errorf("fabric group %s: %w", allocationRange.apiName, err)
		}
	}

	// Explicit root fields exclude unrelated common fabric objects, including
	// non-pointer empty structs that encoding/json would retain with omitempty.
	return json.Marshal(fabricGroupJSON{
		Category:   data.Category,
		FabricName: data.FabricName,
		Management: fields,
	})
}

// UnmarshalFabricJSON decodes the group String before the common Boolean
// decoder can see it, then supplies the private storage to generated SetModelData.
func (v *FabricGroupsModel) UnmarshalFabricJSON(payload []byte, data *resource_fabric_common.NDFCFabricCommonModel) error {
	if data == nil {
		return errors.New("cannot unmarshal fabric group into a nil common model")
	}

	var wire fabricGroupJSON
	if err := json.Unmarshal(payload, &wire); err != nil {
		return fmt.Errorf("decode fabric group JSON: %w", err)
	}
	if wire.Category != "fabricGroup" {
		return fmt.Errorf("expected fabric group category %q, got %q", "fabricGroup", wire.Category)
	}
	if wire.FabricName == "" {
		return errors.New("fabric group response is missing name")
	}
	if len(wire.Management["type"]) == 0 {
		return errors.New("fabric group response is missing management.type")
	}
	var groupType string
	if err := json.Unmarshal(wire.Management["type"], &groupType); err != nil {
		return fmt.Errorf("decode fabric group management.type: %w", err)
	}
	if !validFabricGroupType(groupType) {
		return fmt.Errorf("unsupported fabric group management type %q", groupType)
	}

	// Filter before typed decoding so irrelevant fields cannot hydrate state or
	// trigger type errors. The response-only securityGroupStatus remains readable.
	filtered, err := json.Marshal(filterFabricGroupManagement(wire.Management, groupType, true))
	if err != nil {
		return fmt.Errorf("prepare fabric group response management: %w", err)
	}
	var management fabricManagementJSON
	managementWire := fabricGroupManagementJSON{fabricManagementJSON: &management}
	if err := json.Unmarshal(filtered, &managementWire); err != nil {
		return fmt.Errorf("decode fabric group management: %w", err)
	}

	// Fresh storage clears omitted/unsupported fields and leaves the caller
	// unchanged on failure. Generated SetModelData converts them to Terraform null.
	decoded := resource_fabric_common.NDFCFabricCommonModel{
		Category:   wire.Category,
		FabricName: wire.FabricName,
		Management: resource_fabric_common.NDFCManagementValue(management),
	}
	decoded.FabricGroupCodec.SecurityGroupTag = managementWire.SecurityGroupTag
	// Preserve only known sensitive values actually omitted by this response.
	// An explicitly returned null still clears the value; unsupported group
	// types never inherit a key from prior state.
	if v != nil && (groupType == "vxlan" || groupType == "vxlanAci") {
		if _, present := wire.Management["multisiteInterConnectBgpKey"]; !present && !v.MultisiteInterConnectBgpKey.IsNull() && !v.MultisiteInterConnectBgpKey.IsUnknown() {
			decoded.Management.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey.ValueString()
		}
	}
	if v != nil && groupType == "vxlan" {
		if _, present := wire.Management["cloudSecKey"]; !present && !v.CloudSecKey.IsNull() && !v.CloudSecKey.IsUnknown() {
			decoded.Management.CloudSecKey = v.CloudSecKey.ValueString()
		}
	}
	*data = decoded
	return nil
}

func filterFabricGroupManagement(fields map[string]json.RawMessage, groupType string, response bool) map[string]json.RawMessage {
	filtered := map[string]json.RawMessage{"type": fields["type"]}
	for _, field := range fabricGroupManagementFields {
		if !field.supports(groupType) || (field.readOnly && !response) {
			continue
		}
		if value, present := fields[field.apiName]; present {
			filtered[field.apiName] = value
		}
	}
	return filtered
}
