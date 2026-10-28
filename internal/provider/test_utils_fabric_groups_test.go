// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package provider

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

func FabricGroupsModelHelperStateCheck(RscName string, c resource_fabric_common.NDFCFabricCommonModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.FabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_name").String(), c.FabricName))
	}
	if c.Management.FabricType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fabric_type").String(), c.Management.FabricType))
	}
	if c.Management.AnycastGatewayMac != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("anycast_gateway_mac").String(), c.Management.AnycastGatewayMac))
	}
	if c.Category != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), c.Category))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("category").String(), "fabricGroup"))
	}
	if c.Management.VrfTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_template").String(), c.Management.VrfTemplate))
	}
	if c.Management.NetworkTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_template").String(), c.Management.NetworkTemplate))
	}
	if c.Management.VrfExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_extension_template").String(), c.Management.VrfExtensionTemplate))
	}
	if c.Management.NetworkExtensionTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("network_extension_template").String(), c.Management.NetworkExtensionTemplate))
	}
	if c.Management.L2VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l2_vni_range").String(), c.Management.L2VniRange))
	}
	if c.Management.L3VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("l3_vni_range").String(), c.Management.L3VniRange))
	}
	if c.Management.ScheduledBackup != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup").String(), strconv.FormatBool(*c.Management.ScheduledBackup)))
	}
	if c.Management.ScheduledBackupTime != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("scheduled_backup_time").String(), c.Management.ScheduledBackupTime))
	}
	if c.Management.UnderlayIpv6 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("underlay_ipv6").String(), strconv.FormatBool(*c.Management.UnderlayIpv6)))
	}
	if c.Management.PrivateVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("private_vlan").String(), strconv.FormatBool(*c.Management.PrivateVlan)))
	}
	if c.Management.DefaultPrivateVlanSecondaryNetworkTemplate != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_private_vlan_secondary_network_template").String(), c.Management.DefaultPrivateVlanSecondaryNetworkTemplate))
	}
	if c.FabricGroupCodec.SecurityGroupTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag").String(), c.FabricGroupCodec.SecurityGroupTag))
	}
	if c.Management.SecurityGroupTagPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_prefix").String(), c.Management.SecurityGroupTagPrefix))
	}
	if c.Management.SecurityGroupTagIdRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_id_range").String(), c.Management.SecurityGroupTagIdRange))
	}
	if c.Management.SecurityGroupTagPreprovision != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_preprovision").String(), strconv.FormatBool(*c.Management.SecurityGroupTagPreprovision)))
	}
	if c.Management.SecurityGroupTagMacSegmentation != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_tag_mac_segmentation").String(), strconv.FormatBool(*c.Management.SecurityGroupTagMacSegmentation)))
	}
	if c.Management.SecurityGroupStatus != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("security_group_status").String(), c.Management.SecurityGroupStatus))
	}
	if c.Management.DownstreamVni != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("downstream_vni").String(), strconv.FormatBool(*c.Management.DownstreamVni)))
	}
	if c.Management.DownstreamL2VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("downstream_l2_vni_range").String(), c.Management.DownstreamL2VniRange))
	}
	if c.Management.DownstreamL3VniRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("downstream_l3_vni_range").String(), c.Management.DownstreamL3VniRange))
	}
	if c.Management.MultisiteOverlayInterConnectType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_overlay_inter_connect_type").String(), c.Management.MultisiteOverlayInterConnectType))
	}
	if c.Management.RouteServerRedistributeDirectRouteMap != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_redistribute_direct_route_map").String(), strconv.FormatBool(*c.Management.RouteServerRedistributeDirectRouteMap)))
	}
	if c.Management.RouteServerRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_routing_tag").String(), strconv.Itoa(int(*c.Management.RouteServerRoutingTag))))
	}
	if c.Management.EnableMultisiteOverlayInterFabricBgpDescription != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_multisite_overlay_inter_fabric_bgp_description").String(), strconv.FormatBool(*c.Management.EnableMultisiteOverlayInterFabricBgpDescription)))
	}
	if c.Management.AutoMultisiteUnderlayInterConnect != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_multisite_underlay_inter_connect").String(), strconv.FormatBool(*c.Management.AutoMultisiteUnderlayInterConnect)))
	}
	if c.Management.BgpSendCommunity != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_send_community").String(), strconv.FormatBool(*c.Management.BgpSendCommunity)))
	}
	if c.Management.BgpLogNeighborChange != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_log_neighbor_change").String(), strconv.FormatBool(*c.Management.BgpLogNeighborChange)))
	}
	if c.Management.BgpBfd != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bgp_bfd").String(), strconv.FormatBool(*c.Management.BgpBfd)))
	}
	if c.Management.MultisiteDelayRestore != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_delay_restore").String(), strconv.Itoa(int(*c.Management.MultisiteDelayRestore))))
	}
	if c.Management.MultisiteInterConnectBgpAuthentication != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_authentication").String(), strconv.FormatBool(*c.Management.MultisiteInterConnectBgpAuthentication)))
	}
	if c.Management.MultisiteInterConnectBgpAuthKeyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_auth_key_type").String(), c.Management.MultisiteInterConnectBgpAuthKeyType))
	}
	if c.Management.MultisiteInterConnectBgpKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_inter_connect_bgp_key").String(), c.Management.MultisiteInterConnectBgpKey))
	}
	if c.Management.MultisiteLoopbackId != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_loopback_id").String(), strconv.Itoa(int(*c.Management.MultisiteLoopbackId))))
	}
	if c.Management.BorderGatewayRoutingTag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("border_gateway_routing_tag").String(), strconv.Itoa(int(*c.Management.BorderGatewayRoutingTag))))
	}
	if c.Management.TenantRoutedMulticastIpv4Ipv6 != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("tenant_routed_multicast_ipv4_ipv6").String(), strconv.FormatBool(*c.Management.TenantRoutedMulticastIpv4Ipv6)))
	}
	if c.Management.AutoConfigureCloudSec != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_configure_cloud_sec").String(), strconv.FormatBool(*c.Management.AutoConfigureCloudSec)))
	}
	if c.Management.CloudSecKey != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloud_sec_key").String(), c.Management.CloudSecKey))
	}
	if c.Management.CloudSecAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloud_sec_algorithm").String(), c.Management.CloudSecAlgorithm))
	}
	if c.Management.CloudSecEnforcement != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloud_sec_enforcement").String(), c.Management.CloudSecEnforcement))
	}
	if c.Management.CloudSecReportTimer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("cloud_sec_report_timer").String(), strconv.Itoa(int(*c.Management.CloudSecReportTimer))))
	}
	if c.Management.MultisiteLoopbackIpRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_loopback_ip_range").String(), c.Management.MultisiteLoopbackIpRange))
	}
	if c.Management.MultisiteUnderlaySubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_subnet_range").String(), c.Management.MultisiteUnderlaySubnetRange))
	}
	if c.Management.MultisiteUnderlaySubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.MultisiteUnderlaySubnetTargetMask))))
	}
	if c.Management.MultisiteLoopbackIpv6Range != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_loopback_ipv6_range").String(), c.Management.MultisiteLoopbackIpv6Range))
	}
	if c.Management.MultisiteUnderlayIpv6SubnetRange != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_ipv6_subnet_range").String(), c.Management.MultisiteUnderlayIpv6SubnetRange))
	}
	if c.Management.MultisiteUnderlayIpv6SubnetTargetMask != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("multisite_underlay_ipv6_subnet_target_mask").String(), strconv.Itoa(int(*c.Management.MultisiteUnderlayIpv6SubnetTargetMask))))
	}
	return ret
}

func FabricGroupsRouteServerCollectionValueHelperStateCheck(RscName string, c resource_fabric_common.NDFCRouteServerCollectionValue, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.RouteServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_ip").String(), c.RouteServerIp))
	}
	if c.RouteServerAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_asn").String(), c.RouteServerAsn))
	}
	return ret
}
