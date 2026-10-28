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
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithValidateConfig = &fabricGroupsResource{}

type fabricGroupManagementField struct {
	apiName  string
	tfName   string
	vxlanAci bool
	readOnly bool
}

// Manage v1 1.2.224 fabricGroup management union:
// https://developer.cisco.com/docs/nexus-dashboard/latest/fabricgroup/
// All listed fields apply to vxlan; vxlanAci selects its subset. The other
// three management types contain only type. Validation and both JSON directions
// share this list so filtered responses leave unsupported Terraform values null.
var fabricGroupManagementFields = []fabricGroupManagementField{
	{"anycastGatewayMac", "anycast_gateway_mac", true, false},
	{"vrfTemplate", "vrf_template", false, false},
	{"networkTemplate", "network_template", false, false},
	{"vrfExtensionTemplate", "vrf_extension_template", false, false},
	{"networkExtensionTemplate", "network_extension_template", false, false},
	{"l2VniRange", "l2_vni_range", true, false},
	{"l3VniRange", "l3_vni_range", true, false},
	{"scheduledBackup", "scheduled_backup", true, false},
	{"scheduledBackupTime", "scheduled_backup_time", true, false},
	{"underlayIpv6", "underlay_ipv6", false, false},
	{"privateVlan", "private_vlan", false, false},
	{"defaultPrivateVlanSecondaryNetworkTemplate", "default_private_vlan_secondary_network_template", false, false},
	{"securityGroupTag", "security_group_tag", true, false},
	{"securityGroupTagPrefix", "security_group_tag_prefix", true, false},
	{"securityGroupTagIdRange", "security_group_tag_id_range", true, false},
	{"securityGroupTagPreprovision", "security_group_tag_preprovision", true, false},
	{"securityGroupTagMacSegmentation", "security_group_tag_mac_segmentation", true, false},
	{"securityGroupStatus", "security_group_status", true, true},
	{"downstreamVni", "downstream_vni", false, false},
	{"downstreamL2VniRange", "downstream_l2_vni_range", false, false},
	{"downstreamL3VniRange", "downstream_l3_vni_range", false, false},
	{"multisiteOverlayInterConnectType", "multisite_overlay_inter_connect_type", true, false},
	{"routeServerCollection", "route_server_collection", true, false},
	{"routeServerRedistributeDirectRouteMap", "route_server_redistribute_direct_route_map", true, false},
	{"routeServerRoutingTag", "route_server_routing_tag", true, false},
	{"enableMsOverlayIfcBgpDesc", "enable_multisite_overlay_inter_fabric_bgp_description", true, false},
	{"autoMultisiteUnderlayInterConnect", "auto_multisite_underlay_inter_connect", true, false},
	{"bgpSendCommunity", "bgp_send_community", false, false},
	{"bgpLogNeighborChange", "bgp_log_neighbor_change", false, false},
	{"bgpBfd", "bgp_bfd", false, false},
	{"multisiteDelayRestore", "multisite_delay_restore", true, false},
	{"multisiteInterConnectBgpAuthentication", "multisite_inter_connect_bgp_authentication", true, false},
	{"multisiteInterConnectBgpAuthKeyType", "multisite_inter_connect_bgp_auth_key_type", true, false},
	{"multisiteInterConnectBgpKey", "multisite_inter_connect_bgp_key", true, false},
	{"multisiteLoopbackId", "multisite_loopback_id", true, false},
	{"borderGatewayRoutingTag", "border_gateway_routing_tag", true, false},
	{"tenantRoutedMulticastV4V6", "tenant_routed_multicast_ipv4_ipv6", false, false},
	{"autoConfigureCloudSec", "auto_configure_cloud_sec", false, false},
	{"cloudSecKey", "cloud_sec_key", false, false},
	{"cloudSecAlgorithm", "cloud_sec_algorithm", false, false},
	{"cloudSecEnforcement", "cloud_sec_enforcement", false, false},
	{"cloudSecReportTimer", "cloud_sec_report_timer", false, false},
	{"multisiteLoopbackIpRange", "multisite_loopback_ip_range", true, false},
	{"multisiteUnderlaySubnetRange", "multisite_underlay_subnet_range", true, false},
	{"multisiteUnderlaySubnetTargetMask", "multisite_underlay_subnet_target_mask", true, false},
	{"multisiteLoopbackIpv6Range", "multisite_loopback_ipv6_range", false, false},
	{"multisiteUnderlayIpv6SubnetRange", "multisite_underlay_ipv6_subnet_range", false, false},
	{"multisiteUnderlayIpv6SubnetTargetMask", "multisite_underlay_ipv6_subnet_target_mask", false, false},
}

var fabricGroupAllocationRanges = []struct {
	apiName string
	tfName  string
	minimum uint64
	maximum uint64
}{
	{"l2VniRange", "l2_vni_range", 1, 16777214},
	{"l3VniRange", "l3_vni_range", 1, 16777214},
	{"downstreamL2VniRange", "downstream_l2_vni_range", 1, 16777214},
	{"downstreamL3VniRange", "downstream_l3_vni_range", 1, 16777214},
	{"securityGroupTagIdRange", "security_group_tag_id_range", 16, 65535},
}

func validFabricGroupType(groupType string) bool {
	switch groupType {
	case "vxlan", "vxlanAci", "classic", "ipfmFabricGroup", "ai":
		return true
	default:
		return false
	}
}

func (f fabricGroupManagementField) supports(groupType string) bool {
	return groupType == "vxlan" || (groupType == "vxlanAci" && f.vxlanAci)
}

func (*fabricGroupsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	validateFabricGroupConfiguration(ctx, req.Config, types.StringNull(), &resp.Diagnostics)
}

// Inspect configuration rather than planned computed defaults: a configured
// false, zero, empty collection, or unknown expression still counts as supplied.
// Create/Update provide the resolved selector to repeat deferred type checks.
func validateFabricGroupConfiguration(ctx context.Context, config tfsdk.Config, resolvedType types.String, dg *diag.Diagnostics) {
	var object types.Object
	dg.Append(config.Get(ctx, &object)...)
	if dg.HasError() || object.IsNull() || object.IsUnknown() {
		return
	}
	attributes := object.Attributes()
	groupType := resolvedType
	if groupType.IsNull() || groupType.IsUnknown() {
		groupType, _ = attributes["fabric_type"].(types.String)
	}

	if !groupType.IsNull() && !groupType.IsUnknown() {
		selector := groupType.ValueString()
		if !validFabricGroupType(selector) {
			dg.AddAttributeError(path.Root("fabric_type"), "Invalid Fabric Group Type", "Expected vxlan, vxlanAci, classic, ipfmFabricGroup, or ai.")
			return
		}
		for _, field := range fabricGroupManagementFields {
			if field.readOnly || field.supports(selector) {
				continue
			}
			value := attributes[field.tfName]
			if value != nil && !value.IsNull() {
				dg.AddAttributeError(
					path.Root(field.tfName), "Unsupported Fabric Group Attribute",
					fmt.Sprintf("%s cannot be configured when fabric_type is %q. Remove this attribute from the configuration.", field.tfName, selector),
				)
			}
		}
		mode, ok := attributes["security_group_tag"].(types.String)
		if selector == "vxlanAci" && ok && !mode.IsNull() && !mode.IsUnknown() && mode.ValueString() == "off" {
			dg.AddAttributeError(path.Root("security_group_tag"), "Invalid Fabric Group Security Mode", "vxlanAci supports only loose or strict security_group_tag enforcement.")
		}
	}

	for _, allocationRange := range fabricGroupAllocationRanges {
		value, ok := attributes[allocationRange.tfName].(types.String)
		if !ok || value.IsNull() || value.IsUnknown() {
			continue
		}
		if err := validateFabricGroupRange(value.ValueString(), allocationRange.minimum, allocationRange.maximum); err != nil {
			dg.AddAttributeError(path.Root(allocationRange.tfName), "Invalid Fabric Group Allocation Range", err.Error())
		}
	}
}

func validateFabricGroupRange(value string, minimum, maximum uint64) error {
	_, err := parseFabricGroupRange(value, minimum, maximum)
	return err
}

func parseFabricGroupRange(value string, minimum, maximum uint64) ([2]uint64, error) {
	invalidRange := fmt.Errorf("must use start-end syntax with endpoints between %d and %d and a start no greater than the end", minimum, maximum)
	endpoints := strings.Split(value, "-")
	if len(endpoints) != 2 {
		return [2]uint64{}, invalidRange
	}
	var bounds [2]uint64
	for i, endpoint := range endpoints {
		for _, digit := range endpoint {
			if digit < '0' || digit > '9' {
				return [2]uint64{}, invalidRange
			}
		}
		bound, err := strconv.ParseUint(endpoint, 10, 64)
		if err != nil || bound < minimum || bound > maximum {
			return [2]uint64{}, invalidRange
		}
		bounds[i] = bound
	}
	if bounds[0] > bounds[1] {
		return [2]uint64{}, invalidRange
	}
	return bounds, nil
}
