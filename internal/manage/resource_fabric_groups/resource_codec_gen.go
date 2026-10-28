// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated;  DO NOT EDIT.

package resource_fabric_groups

import (
	"context"
	"log"
	"terraform-provider-nd/internal/manage/resource_fabric_common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (v *FabricGroupsModel) SetModelData(jsonData *resource_fabric_common.NDFCFabricCommonModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.FabricName != "" {
		v.FabricName = types.StringValue(jsonData.FabricName)
	} else {
		v.FabricName = types.StringNull()
	}

	if jsonData.Management.FabricType != "" {
		v.FabricType = types.StringValue(jsonData.Management.FabricType)
	} else {
		v.FabricType = types.StringNull()
	}

	if jsonData.Management.AnycastGatewayMac != "" {
		v.AnycastGatewayMac = types.StringValue(jsonData.Management.AnycastGatewayMac)
	} else {
		v.AnycastGatewayMac = types.StringNull()
	}

	if jsonData.Category != "" {
		v.Category = types.StringValue(jsonData.Category)
	} else {
		v.Category = types.StringNull()
	}

	if jsonData.Management.VrfTemplate != "" {
		v.VrfTemplate = types.StringValue(jsonData.Management.VrfTemplate)
	} else {
		v.VrfTemplate = types.StringNull()
	}

	if jsonData.Management.NetworkTemplate != "" {
		v.NetworkTemplate = types.StringValue(jsonData.Management.NetworkTemplate)
	} else {
		v.NetworkTemplate = types.StringNull()
	}

	if jsonData.Management.VrfExtensionTemplate != "" {
		v.VrfExtensionTemplate = types.StringValue(jsonData.Management.VrfExtensionTemplate)
	} else {
		v.VrfExtensionTemplate = types.StringNull()
	}

	if jsonData.Management.NetworkExtensionTemplate != "" {
		v.NetworkExtensionTemplate = types.StringValue(jsonData.Management.NetworkExtensionTemplate)
	} else {
		v.NetworkExtensionTemplate = types.StringNull()
	}

	if jsonData.Management.L2VniRange != "" {
		v.L2VniRange = types.StringValue(jsonData.Management.L2VniRange)
	} else {
		v.L2VniRange = types.StringNull()
	}

	if jsonData.Management.L3VniRange != "" {
		v.L3VniRange = types.StringValue(jsonData.Management.L3VniRange)
	} else {
		v.L3VniRange = types.StringNull()
	}

	if jsonData.Management.ScheduledBackup != nil {
		v.ScheduledBackup = types.BoolValue(*jsonData.Management.ScheduledBackup)
	} else {
		v.ScheduledBackup = types.BoolNull()
	}

	if jsonData.Management.ScheduledBackupTime != "" {
		v.ScheduledBackupTime = types.StringValue(jsonData.Management.ScheduledBackupTime)
	} else {
		v.ScheduledBackupTime = types.StringNull()
	}

	if jsonData.Management.UnderlayIpv6 != nil {
		v.UnderlayIpv6 = types.BoolValue(*jsonData.Management.UnderlayIpv6)
	} else {
		v.UnderlayIpv6 = types.BoolNull()
	}

	if jsonData.Management.PrivateVlan != nil {
		v.PrivateVlan = types.BoolValue(*jsonData.Management.PrivateVlan)
	} else {
		v.PrivateVlan = types.BoolNull()
	}

	if jsonData.Management.DefaultPrivateVlanSecondaryNetworkTemplate != "" {
		v.DefaultPrivateVlanSecondaryNetworkTemplate = types.StringValue(jsonData.Management.DefaultPrivateVlanSecondaryNetworkTemplate)
	} else {
		v.DefaultPrivateVlanSecondaryNetworkTemplate = types.StringNull()
	}

	if jsonData.FabricGroupCodec.SecurityGroupTag != "" {
		v.SecurityGroupTag = types.StringValue(jsonData.FabricGroupCodec.SecurityGroupTag)
	} else {
		v.SecurityGroupTag = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagPrefix != "" {
		v.SecurityGroupTagPrefix = types.StringValue(jsonData.Management.SecurityGroupTagPrefix)
	} else {
		v.SecurityGroupTagPrefix = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagIdRange != "" {
		v.SecurityGroupTagIdRange = types.StringValue(jsonData.Management.SecurityGroupTagIdRange)
	} else {
		v.SecurityGroupTagIdRange = types.StringNull()
	}

	if jsonData.Management.SecurityGroupTagPreprovision != nil {
		v.SecurityGroupTagPreprovision = types.BoolValue(*jsonData.Management.SecurityGroupTagPreprovision)
	} else {
		v.SecurityGroupTagPreprovision = types.BoolNull()
	}

	if jsonData.Management.SecurityGroupTagMacSegmentation != nil {
		v.SecurityGroupTagMacSegmentation = types.BoolValue(*jsonData.Management.SecurityGroupTagMacSegmentation)
	} else {
		v.SecurityGroupTagMacSegmentation = types.BoolNull()
	}

	if jsonData.Management.SecurityGroupStatus != "" {
		v.SecurityGroupStatus = types.StringValue(jsonData.Management.SecurityGroupStatus)
	} else {
		v.SecurityGroupStatus = types.StringNull()
	}

	if jsonData.Management.DownstreamVni != nil {
		v.DownstreamVni = types.BoolValue(*jsonData.Management.DownstreamVni)
	} else {
		v.DownstreamVni = types.BoolNull()
	}

	if jsonData.Management.DownstreamL2VniRange != "" {
		v.DownstreamL2VniRange = types.StringValue(jsonData.Management.DownstreamL2VniRange)
	} else {
		v.DownstreamL2VniRange = types.StringNull()
	}

	if jsonData.Management.DownstreamL3VniRange != "" {
		v.DownstreamL3VniRange = types.StringValue(jsonData.Management.DownstreamL3VniRange)
	} else {
		v.DownstreamL3VniRange = types.StringNull()
	}

	if jsonData.Management.MultisiteOverlayInterConnectType != "" {
		v.MultisiteOverlayInterConnectType = types.StringValue(jsonData.Management.MultisiteOverlayInterConnectType)
	} else {
		v.MultisiteOverlayInterConnectType = types.StringNull()
	}

	if len(jsonData.Management.RouteServerCollection) == 0 {
		log.Printf("v.RouteServerCollection is empty")
		v.RouteServerCollection = types.ListNull(RouteServerCollectionValue{}.Type(context.Background()))
	} else {
		listData := make([]RouteServerCollectionValue, len(jsonData.Management.RouteServerCollection))
		for i, item := range jsonData.Management.RouteServerCollection {
			err = listData[i].SetValue(&item)
			if err != nil {
				return err
			}
			listData[i].state = attr.ValueStateKnown
		}
		v.RouteServerCollection, err = types.ListValueFrom(context.Background(), RouteServerCollectionValue{}.Type(context.Background()), listData)
		if err != nil {
			return err
		}
	}
	if jsonData.Management.RouteServerRedistributeDirectRouteMap != nil {
		v.RouteServerRedistributeDirectRouteMap = types.BoolValue(*jsonData.Management.RouteServerRedistributeDirectRouteMap)
	} else {
		v.RouteServerRedistributeDirectRouteMap = types.BoolNull()
	}

	if jsonData.Management.RouteServerRoutingTag != nil {
		v.RouteServerRoutingTag = types.Int64Value(*jsonData.Management.RouteServerRoutingTag)
	} else {
		v.RouteServerRoutingTag = types.Int64Null()
	}

	if jsonData.Management.EnableMultisiteOverlayInterFabricBgpDescription != nil {
		v.EnableMultisiteOverlayInterFabricBgpDescription = types.BoolValue(*jsonData.Management.EnableMultisiteOverlayInterFabricBgpDescription)
	} else {
		v.EnableMultisiteOverlayInterFabricBgpDescription = types.BoolNull()
	}

	if jsonData.Management.AutoMultisiteUnderlayInterConnect != nil {
		v.AutoMultisiteUnderlayInterConnect = types.BoolValue(*jsonData.Management.AutoMultisiteUnderlayInterConnect)
	} else {
		v.AutoMultisiteUnderlayInterConnect = types.BoolNull()
	}

	if jsonData.Management.BgpSendCommunity != nil {
		v.BgpSendCommunity = types.BoolValue(*jsonData.Management.BgpSendCommunity)
	} else {
		v.BgpSendCommunity = types.BoolNull()
	}

	if jsonData.Management.BgpLogNeighborChange != nil {
		v.BgpLogNeighborChange = types.BoolValue(*jsonData.Management.BgpLogNeighborChange)
	} else {
		v.BgpLogNeighborChange = types.BoolNull()
	}

	if jsonData.Management.BgpBfd != nil {
		v.BgpBfd = types.BoolValue(*jsonData.Management.BgpBfd)
	} else {
		v.BgpBfd = types.BoolNull()
	}

	if jsonData.Management.MultisiteDelayRestore != nil {
		v.MultisiteDelayRestore = types.Int64Value(*jsonData.Management.MultisiteDelayRestore)
	} else {
		v.MultisiteDelayRestore = types.Int64Null()
	}

	if jsonData.Management.MultisiteInterConnectBgpAuthentication != nil {
		v.MultisiteInterConnectBgpAuthentication = types.BoolValue(*jsonData.Management.MultisiteInterConnectBgpAuthentication)
	} else {
		v.MultisiteInterConnectBgpAuthentication = types.BoolNull()
	}

	if jsonData.Management.MultisiteInterConnectBgpAuthKeyType != "" {
		v.MultisiteInterConnectBgpAuthKeyType = types.StringValue(jsonData.Management.MultisiteInterConnectBgpAuthKeyType)
	} else {
		v.MultisiteInterConnectBgpAuthKeyType = types.StringNull()
	}

	if jsonData.Management.MultisiteInterConnectBgpKey != "" {
		v.MultisiteInterConnectBgpKey = types.StringValue(jsonData.Management.MultisiteInterConnectBgpKey)
	} else {
		v.MultisiteInterConnectBgpKey = types.StringNull()
	}

	if jsonData.Management.MultisiteLoopbackId != nil {
		v.MultisiteLoopbackId = types.Int64Value(*jsonData.Management.MultisiteLoopbackId)
	} else {
		v.MultisiteLoopbackId = types.Int64Null()
	}

	if jsonData.Management.BorderGatewayRoutingTag != nil {
		v.BorderGatewayRoutingTag = types.Int64Value(*jsonData.Management.BorderGatewayRoutingTag)
	} else {
		v.BorderGatewayRoutingTag = types.Int64Null()
	}

	if jsonData.Management.TenantRoutedMulticastIpv4Ipv6 != nil {
		v.TenantRoutedMulticastIpv4Ipv6 = types.BoolValue(*jsonData.Management.TenantRoutedMulticastIpv4Ipv6)
	} else {
		v.TenantRoutedMulticastIpv4Ipv6 = types.BoolNull()
	}

	if jsonData.Management.AutoConfigureCloudSec != nil {
		v.AutoConfigureCloudSec = types.BoolValue(*jsonData.Management.AutoConfigureCloudSec)
	} else {
		v.AutoConfigureCloudSec = types.BoolNull()
	}

	if jsonData.Management.CloudSecKey != "" {
		v.CloudSecKey = types.StringValue(jsonData.Management.CloudSecKey)
	} else {
		v.CloudSecKey = types.StringNull()
	}

	if jsonData.Management.CloudSecAlgorithm != "" {
		v.CloudSecAlgorithm = types.StringValue(jsonData.Management.CloudSecAlgorithm)
	} else {
		v.CloudSecAlgorithm = types.StringNull()
	}

	if jsonData.Management.CloudSecEnforcement != "" {
		v.CloudSecEnforcement = types.StringValue(jsonData.Management.CloudSecEnforcement)
	} else {
		v.CloudSecEnforcement = types.StringNull()
	}

	if jsonData.Management.CloudSecReportTimer != nil {
		v.CloudSecReportTimer = types.Int64Value(*jsonData.Management.CloudSecReportTimer)
	} else {
		v.CloudSecReportTimer = types.Int64Null()
	}

	if jsonData.Management.MultisiteLoopbackIpRange != "" {
		v.MultisiteLoopbackIpRange = types.StringValue(jsonData.Management.MultisiteLoopbackIpRange)
	} else {
		v.MultisiteLoopbackIpRange = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlaySubnetRange != "" {
		v.MultisiteUnderlaySubnetRange = types.StringValue(jsonData.Management.MultisiteUnderlaySubnetRange)
	} else {
		v.MultisiteUnderlaySubnetRange = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlaySubnetTargetMask != nil {
		v.MultisiteUnderlaySubnetTargetMask = types.Int64Value(*jsonData.Management.MultisiteUnderlaySubnetTargetMask)
	} else {
		v.MultisiteUnderlaySubnetTargetMask = types.Int64Null()
	}

	if jsonData.Management.MultisiteLoopbackIpv6Range != "" {
		v.MultisiteLoopbackIpv6Range = types.StringValue(jsonData.Management.MultisiteLoopbackIpv6Range)
	} else {
		v.MultisiteLoopbackIpv6Range = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlayIpv6SubnetRange != "" {
		v.MultisiteUnderlayIpv6SubnetRange = types.StringValue(jsonData.Management.MultisiteUnderlayIpv6SubnetRange)
	} else {
		v.MultisiteUnderlayIpv6SubnetRange = types.StringNull()
	}

	if jsonData.Management.MultisiteUnderlayIpv6SubnetTargetMask != nil {
		v.MultisiteUnderlayIpv6SubnetTargetMask = types.Int64Value(*jsonData.Management.MultisiteUnderlayIpv6SubnetTargetMask)
	} else {
		v.MultisiteUnderlayIpv6SubnetTargetMask = types.Int64Null()
	}

	return err
}

func (v *RouteServerCollectionValue) SetValue(jsonData *resource_fabric_common.NDFCRouteServerCollectionValue) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.RouteServerIp != "" {
		v.RouteServerIp = types.StringValue(jsonData.RouteServerIp)
	} else {
		v.RouteServerIp = types.StringNull()
	}

	if jsonData.RouteServerAsn != "" {
		v.RouteServerAsn = types.StringValue(jsonData.RouteServerAsn)
	} else {
		v.RouteServerAsn = types.StringNull()
	}

	return err
}

func (v FabricGroupsModel) GetModelData() *resource_fabric_common.NDFCFabricCommonModel {
	var data = new(resource_fabric_common.NDFCFabricCommonModel)

	//MARSHAL_BODY

	if !v.FabricName.IsNull() && !v.FabricName.IsUnknown() {
		data.FabricName = v.FabricName.ValueString()
	} else {
		data.FabricName = ""
	}

	if !v.FabricType.IsNull() && !v.FabricType.IsUnknown() {
		data.Management.FabricType = v.FabricType.ValueString()
	} else {
		data.Management.FabricType = ""
	}

	if !v.AnycastGatewayMac.IsNull() && !v.AnycastGatewayMac.IsUnknown() {
		data.Management.AnycastGatewayMac = v.AnycastGatewayMac.ValueString()
	} else {
		data.Management.AnycastGatewayMac = ""
	}

	if !v.Category.IsNull() && !v.Category.IsUnknown() {
		data.Category = v.Category.ValueString()
	} else {
		data.Category = ""
	}

	if !v.VrfTemplate.IsNull() && !v.VrfTemplate.IsUnknown() {
		data.Management.VrfTemplate = v.VrfTemplate.ValueString()
	} else {
		data.Management.VrfTemplate = ""
	}

	if !v.NetworkTemplate.IsNull() && !v.NetworkTemplate.IsUnknown() {
		data.Management.NetworkTemplate = v.NetworkTemplate.ValueString()
	} else {
		data.Management.NetworkTemplate = ""
	}

	if !v.VrfExtensionTemplate.IsNull() && !v.VrfExtensionTemplate.IsUnknown() {
		data.Management.VrfExtensionTemplate = v.VrfExtensionTemplate.ValueString()
	} else {
		data.Management.VrfExtensionTemplate = ""
	}

	if !v.NetworkExtensionTemplate.IsNull() && !v.NetworkExtensionTemplate.IsUnknown() {
		data.Management.NetworkExtensionTemplate = v.NetworkExtensionTemplate.ValueString()
	} else {
		data.Management.NetworkExtensionTemplate = ""
	}

	if !v.L2VniRange.IsNull() && !v.L2VniRange.IsUnknown() {
		data.Management.L2VniRange = v.L2VniRange.ValueString()
	} else {
		data.Management.L2VniRange = ""
	}

	if !v.L3VniRange.IsNull() && !v.L3VniRange.IsUnknown() {
		data.Management.L3VniRange = v.L3VniRange.ValueString()
	} else {
		data.Management.L3VniRange = ""
	}

	if !v.ScheduledBackup.IsNull() && !v.ScheduledBackup.IsUnknown() {
		data.Management.ScheduledBackup = new(bool)
		*data.Management.ScheduledBackup = v.ScheduledBackup.ValueBool()
	} else {
		data.Management.ScheduledBackup = nil
	}

	if !v.ScheduledBackupTime.IsNull() && !v.ScheduledBackupTime.IsUnknown() {
		data.Management.ScheduledBackupTime = v.ScheduledBackupTime.ValueString()
	} else {
		data.Management.ScheduledBackupTime = ""
	}

	if !v.UnderlayIpv6.IsNull() && !v.UnderlayIpv6.IsUnknown() {
		data.Management.UnderlayIpv6 = new(bool)
		*data.Management.UnderlayIpv6 = v.UnderlayIpv6.ValueBool()
	} else {
		data.Management.UnderlayIpv6 = nil
	}

	if !v.PrivateVlan.IsNull() && !v.PrivateVlan.IsUnknown() {
		data.Management.PrivateVlan = new(bool)
		*data.Management.PrivateVlan = v.PrivateVlan.ValueBool()
	} else {
		data.Management.PrivateVlan = nil
	}

	if !v.DefaultPrivateVlanSecondaryNetworkTemplate.IsNull() && !v.DefaultPrivateVlanSecondaryNetworkTemplate.IsUnknown() {
		data.Management.DefaultPrivateVlanSecondaryNetworkTemplate = v.DefaultPrivateVlanSecondaryNetworkTemplate.ValueString()
	} else {
		data.Management.DefaultPrivateVlanSecondaryNetworkTemplate = ""
	}

	if !v.SecurityGroupTag.IsNull() && !v.SecurityGroupTag.IsUnknown() {
		data.FabricGroupCodec.SecurityGroupTag = v.SecurityGroupTag.ValueString()
	} else {
		data.FabricGroupCodec.SecurityGroupTag = ""
	}

	if !v.SecurityGroupTagPrefix.IsNull() && !v.SecurityGroupTagPrefix.IsUnknown() {
		data.Management.SecurityGroupTagPrefix = v.SecurityGroupTagPrefix.ValueString()
	} else {
		data.Management.SecurityGroupTagPrefix = ""
	}

	if !v.SecurityGroupTagIdRange.IsNull() && !v.SecurityGroupTagIdRange.IsUnknown() {
		data.Management.SecurityGroupTagIdRange = v.SecurityGroupTagIdRange.ValueString()
	} else {
		data.Management.SecurityGroupTagIdRange = ""
	}

	if !v.SecurityGroupTagPreprovision.IsNull() && !v.SecurityGroupTagPreprovision.IsUnknown() {
		data.Management.SecurityGroupTagPreprovision = new(bool)
		*data.Management.SecurityGroupTagPreprovision = v.SecurityGroupTagPreprovision.ValueBool()
	} else {
		data.Management.SecurityGroupTagPreprovision = nil
	}

	if !v.SecurityGroupTagMacSegmentation.IsNull() && !v.SecurityGroupTagMacSegmentation.IsUnknown() {
		data.Management.SecurityGroupTagMacSegmentation = new(bool)
		*data.Management.SecurityGroupTagMacSegmentation = v.SecurityGroupTagMacSegmentation.ValueBool()
	} else {
		data.Management.SecurityGroupTagMacSegmentation = nil
	}

	if !v.DownstreamVni.IsNull() && !v.DownstreamVni.IsUnknown() {
		data.Management.DownstreamVni = new(bool)
		*data.Management.DownstreamVni = v.DownstreamVni.ValueBool()
	} else {
		data.Management.DownstreamVni = nil
	}

	if !v.DownstreamL2VniRange.IsNull() && !v.DownstreamL2VniRange.IsUnknown() {
		data.Management.DownstreamL2VniRange = v.DownstreamL2VniRange.ValueString()
	} else {
		data.Management.DownstreamL2VniRange = ""
	}

	if !v.DownstreamL3VniRange.IsNull() && !v.DownstreamL3VniRange.IsUnknown() {
		data.Management.DownstreamL3VniRange = v.DownstreamL3VniRange.ValueString()
	} else {
		data.Management.DownstreamL3VniRange = ""
	}

	if !v.MultisiteOverlayInterConnectType.IsNull() && !v.MultisiteOverlayInterConnectType.IsUnknown() {
		data.Management.MultisiteOverlayInterConnectType = v.MultisiteOverlayInterConnectType.ValueString()
	} else {
		data.Management.MultisiteOverlayInterConnectType = ""
	}

	if !v.RouteServerRedistributeDirectRouteMap.IsNull() && !v.RouteServerRedistributeDirectRouteMap.IsUnknown() {
		data.Management.RouteServerRedistributeDirectRouteMap = new(bool)
		*data.Management.RouteServerRedistributeDirectRouteMap = v.RouteServerRedistributeDirectRouteMap.ValueBool()
	} else {
		data.Management.RouteServerRedistributeDirectRouteMap = nil
	}

	if !v.RouteServerRoutingTag.IsNull() && !v.RouteServerRoutingTag.IsUnknown() {
		data.Management.RouteServerRoutingTag = new(int64)
		*data.Management.RouteServerRoutingTag = v.RouteServerRoutingTag.ValueInt64()

	} else {
		data.Management.RouteServerRoutingTag = nil
	}

	if !v.EnableMultisiteOverlayInterFabricBgpDescription.IsNull() && !v.EnableMultisiteOverlayInterFabricBgpDescription.IsUnknown() {
		data.Management.EnableMultisiteOverlayInterFabricBgpDescription = new(bool)
		*data.Management.EnableMultisiteOverlayInterFabricBgpDescription = v.EnableMultisiteOverlayInterFabricBgpDescription.ValueBool()
	} else {
		data.Management.EnableMultisiteOverlayInterFabricBgpDescription = nil
	}

	if !v.AutoMultisiteUnderlayInterConnect.IsNull() && !v.AutoMultisiteUnderlayInterConnect.IsUnknown() {
		data.Management.AutoMultisiteUnderlayInterConnect = new(bool)
		*data.Management.AutoMultisiteUnderlayInterConnect = v.AutoMultisiteUnderlayInterConnect.ValueBool()
	} else {
		data.Management.AutoMultisiteUnderlayInterConnect = nil
	}

	if !v.BgpSendCommunity.IsNull() && !v.BgpSendCommunity.IsUnknown() {
		data.Management.BgpSendCommunity = new(bool)
		*data.Management.BgpSendCommunity = v.BgpSendCommunity.ValueBool()
	} else {
		data.Management.BgpSendCommunity = nil
	}

	if !v.BgpLogNeighborChange.IsNull() && !v.BgpLogNeighborChange.IsUnknown() {
		data.Management.BgpLogNeighborChange = new(bool)
		*data.Management.BgpLogNeighborChange = v.BgpLogNeighborChange.ValueBool()
	} else {
		data.Management.BgpLogNeighborChange = nil
	}

	if !v.BgpBfd.IsNull() && !v.BgpBfd.IsUnknown() {
		data.Management.BgpBfd = new(bool)
		*data.Management.BgpBfd = v.BgpBfd.ValueBool()
	} else {
		data.Management.BgpBfd = nil
	}

	if !v.MultisiteDelayRestore.IsNull() && !v.MultisiteDelayRestore.IsUnknown() {
		data.Management.MultisiteDelayRestore = new(int64)
		*data.Management.MultisiteDelayRestore = v.MultisiteDelayRestore.ValueInt64()

	} else {
		data.Management.MultisiteDelayRestore = nil
	}

	if !v.MultisiteInterConnectBgpAuthentication.IsNull() && !v.MultisiteInterConnectBgpAuthentication.IsUnknown() {
		data.Management.MultisiteInterConnectBgpAuthentication = new(bool)
		*data.Management.MultisiteInterConnectBgpAuthentication = v.MultisiteInterConnectBgpAuthentication.ValueBool()
	} else {
		data.Management.MultisiteInterConnectBgpAuthentication = nil
	}

	if !v.MultisiteInterConnectBgpAuthKeyType.IsNull() && !v.MultisiteInterConnectBgpAuthKeyType.IsUnknown() {
		data.Management.MultisiteInterConnectBgpAuthKeyType = v.MultisiteInterConnectBgpAuthKeyType.ValueString()
	} else {
		data.Management.MultisiteInterConnectBgpAuthKeyType = ""
	}

	if !v.MultisiteInterConnectBgpKey.IsNull() && !v.MultisiteInterConnectBgpKey.IsUnknown() {
		data.Management.MultisiteInterConnectBgpKey = v.MultisiteInterConnectBgpKey.ValueString()
	} else {
		data.Management.MultisiteInterConnectBgpKey = ""
	}

	if !v.MultisiteLoopbackId.IsNull() && !v.MultisiteLoopbackId.IsUnknown() {
		data.Management.MultisiteLoopbackId = new(int64)
		*data.Management.MultisiteLoopbackId = v.MultisiteLoopbackId.ValueInt64()

	} else {
		data.Management.MultisiteLoopbackId = nil
	}

	if !v.BorderGatewayRoutingTag.IsNull() && !v.BorderGatewayRoutingTag.IsUnknown() {
		data.Management.BorderGatewayRoutingTag = new(int64)
		*data.Management.BorderGatewayRoutingTag = v.BorderGatewayRoutingTag.ValueInt64()

	} else {
		data.Management.BorderGatewayRoutingTag = nil
	}

	if !v.TenantRoutedMulticastIpv4Ipv6.IsNull() && !v.TenantRoutedMulticastIpv4Ipv6.IsUnknown() {
		data.Management.TenantRoutedMulticastIpv4Ipv6 = new(bool)
		*data.Management.TenantRoutedMulticastIpv4Ipv6 = v.TenantRoutedMulticastIpv4Ipv6.ValueBool()
	} else {
		data.Management.TenantRoutedMulticastIpv4Ipv6 = nil
	}

	if !v.AutoConfigureCloudSec.IsNull() && !v.AutoConfigureCloudSec.IsUnknown() {
		data.Management.AutoConfigureCloudSec = new(bool)
		*data.Management.AutoConfigureCloudSec = v.AutoConfigureCloudSec.ValueBool()
	} else {
		data.Management.AutoConfigureCloudSec = nil
	}

	if !v.CloudSecKey.IsNull() && !v.CloudSecKey.IsUnknown() {
		data.Management.CloudSecKey = v.CloudSecKey.ValueString()
	} else {
		data.Management.CloudSecKey = ""
	}

	if !v.CloudSecAlgorithm.IsNull() && !v.CloudSecAlgorithm.IsUnknown() {
		data.Management.CloudSecAlgorithm = v.CloudSecAlgorithm.ValueString()
	} else {
		data.Management.CloudSecAlgorithm = ""
	}

	if !v.CloudSecEnforcement.IsNull() && !v.CloudSecEnforcement.IsUnknown() {
		data.Management.CloudSecEnforcement = v.CloudSecEnforcement.ValueString()
	} else {
		data.Management.CloudSecEnforcement = ""
	}

	if !v.CloudSecReportTimer.IsNull() && !v.CloudSecReportTimer.IsUnknown() {
		data.Management.CloudSecReportTimer = new(int64)
		*data.Management.CloudSecReportTimer = v.CloudSecReportTimer.ValueInt64()

	} else {
		data.Management.CloudSecReportTimer = nil
	}

	if !v.MultisiteLoopbackIpRange.IsNull() && !v.MultisiteLoopbackIpRange.IsUnknown() {
		data.Management.MultisiteLoopbackIpRange = v.MultisiteLoopbackIpRange.ValueString()
	} else {
		data.Management.MultisiteLoopbackIpRange = ""
	}

	if !v.MultisiteUnderlaySubnetRange.IsNull() && !v.MultisiteUnderlaySubnetRange.IsUnknown() {
		data.Management.MultisiteUnderlaySubnetRange = v.MultisiteUnderlaySubnetRange.ValueString()
	} else {
		data.Management.MultisiteUnderlaySubnetRange = ""
	}

	if !v.MultisiteUnderlaySubnetTargetMask.IsNull() && !v.MultisiteUnderlaySubnetTargetMask.IsUnknown() {
		data.Management.MultisiteUnderlaySubnetTargetMask = new(int64)
		*data.Management.MultisiteUnderlaySubnetTargetMask = v.MultisiteUnderlaySubnetTargetMask.ValueInt64()

	} else {
		data.Management.MultisiteUnderlaySubnetTargetMask = nil
	}

	if !v.MultisiteLoopbackIpv6Range.IsNull() && !v.MultisiteLoopbackIpv6Range.IsUnknown() {
		data.Management.MultisiteLoopbackIpv6Range = v.MultisiteLoopbackIpv6Range.ValueString()
	} else {
		data.Management.MultisiteLoopbackIpv6Range = ""
	}

	if !v.MultisiteUnderlayIpv6SubnetRange.IsNull() && !v.MultisiteUnderlayIpv6SubnetRange.IsUnknown() {
		data.Management.MultisiteUnderlayIpv6SubnetRange = v.MultisiteUnderlayIpv6SubnetRange.ValueString()
	} else {
		data.Management.MultisiteUnderlayIpv6SubnetRange = ""
	}

	if !v.MultisiteUnderlayIpv6SubnetTargetMask.IsNull() && !v.MultisiteUnderlayIpv6SubnetTargetMask.IsUnknown() {
		data.Management.MultisiteUnderlayIpv6SubnetTargetMask = new(int64)
		*data.Management.MultisiteUnderlayIpv6SubnetTargetMask = v.MultisiteUnderlayIpv6SubnetTargetMask.ValueInt64()

	} else {
		data.Management.MultisiteUnderlayIpv6SubnetTargetMask = nil
	}

	//MARSHALL_LIST  BEGIN RouteServerCollection[i1]

	if !v.RouteServerCollection.IsNull() && !v.RouteServerCollection.IsUnknown() {
		elements := make([]RouteServerCollectionValue, len(v.RouteServerCollection.Elements()))
		// Not augmenting

		data.Management.RouteServerCollection = make([]resource_fabric_common.NDFCRouteServerCollectionValue, len(v.RouteServerCollection.Elements()))

		// -- Set here 1 |.Management.RouteServerCollection[i1]|RouteServerCollection[i1]| --

		diag := v.RouteServerCollection.ElementsAs(context.Background(), &elements, false)
		if diag != nil {
			panic(diag)
		}
		// .Management.RouteServerCollection[i1] RouteServerCollection[i1]
		for i1, ele1 := range elements {
			if !ele1.RouteServerIp.IsNull() && !ele1.RouteServerIp.IsUnknown() {

				data.Management.RouteServerCollection[i1].RouteServerIp = ele1.RouteServerIp.ValueString()
			} else {
				data.Management.RouteServerCollection[i1].RouteServerIp = ""
			}

			if !ele1.RouteServerAsn.IsNull() && !ele1.RouteServerAsn.IsUnknown() {

				data.Management.RouteServerCollection[i1].RouteServerAsn = ele1.RouteServerAsn.ValueString()
			} else {
				data.Management.RouteServerCollection[i1].RouteServerAsn = ""
			}

		} /* for loop */
	} /* isNull if check */

	return data
}
