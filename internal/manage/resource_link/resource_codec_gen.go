// Code generated;  DO NOT EDIT.

package resource_link

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NDFCLinkModel struct {
	LinkId           string              `json:"-"`
	SrcSwitchId      string              `json:"srcSwitchId"`
	DstSwitchId      string              `json:"dstSwitchId"`
	SrcFabricName    string              `json:"srcFabricName"`
	SrcSwitchName    string              `json:"srcSwitchName"`
	SrcInterfaceName string              `json:"srcInterfaceName"`
	DstFabricName    string              `json:"dstFabricName"`
	DstSwitchName    string              `json:"dstSwitchName"`
	DstInterfaceName string              `json:"dstInterfaceName"`
	SrcClusterName   string              `json:"srcClusterName"`
	DstClusterName   string              `json:"dstClusterName"`
	LinkType         string              `json:"linkType"`
	ConfigData       NDFCConfigDataValue `json:"configData,omitempty"`
}

type NDFCConfigDataValue struct {
	PolicyType     string                            `json:"policyType"`
	TemplateInputs NDFCConfigDataTemplateInputsValue `json:"templateInputs,omitempty"`
}

type NDFCConfigDataTemplateInputsValue struct {
	SrcInterfaceDescription              string `json:"srcInterfaceDescription"`
	SrcInterfaceConfig                   string `json:"srcInterfaceConfig"`
	DstInterfaceDescription              string `json:"dstInterfaceDescription"`
	DstInterfaceConfig                   string `json:"dstInterfaceConfig"`
	InterfaceAdminState                  *bool  `json:"interfaceAdminState"`
	Mtu                                  *int64 `json:"mtu"`
	LinkMtu                              *int64 `json:"linkMtu"`
	Speed                                string `json:"speed"`
	Fec                                  string `json:"fec"`
	Macsec                               *bool  `json:"macsec"`
	SrcIp                                string `json:"srcIp"`
	SrcIpv6                              string `json:"srcIpv6"`
	SrcIpAddress                         string `json:"srcIpAddress"`
	SrcIpAddressMask                     string `json:"srcIpAddressMask"`
	SrcIpv6AddressMask                   string `json:"srcIpv6AddressMask"`
	DstIp                                string `json:"dstIp"`
	DstIpv6                              string `json:"dstIpv6"`
	DstIpAddress                         string `json:"dstIpAddress"`
	DstIpAddressMask                     string `json:"dstIpAddressMask"`
	DstIpv6Address                       string `json:"dstIpv6Address"`
	DstIpv6AddressMask                   string `json:"dstIpv6AddressMask"`
	SrcEbgpAsn                           string `json:"srcEbgpAsn"`
	DstEbgpAsn                           string `json:"dstEbgpAsn"`
	EbgpPassword                         string `json:"ebgpPassword"`
	EbgpAuthKeyEncryptionType            string `json:"ebgpAuthKeyEncryptionType"`
	EbgpBfd                              *bool  `json:"ebgpBfd"`
	EbgpLogNeighborChange                *bool  `json:"ebgpLogNeighborChange"`
	EbgpMaximumPaths                     *int64 `json:"ebgpMaximumPaths"`
	EbgpMultihop                         *int64 `json:"ebgpMultihop"`
	EbgpSendComboth                      *bool  `json:"ebgpSendComboth"`
	EnableEbgpPassword                   *bool  `json:"enableEbgpPassword"`
	InheritEbgpPasswordMsdSettings       *bool  `json:"inheritEbgpPasswordMsdSettings"`
	DefaultVrfEbgpNeighborPassword       string `json:"defaultVrfEbgpNeighborPassword"`
	InterfaceVrf                         string `json:"interfaceVrf"`
	SrcVrfName                           string `json:"srcVrfName"`
	DstVrfName                           string `json:"dstVrfName"`
	VrfNameNxPeerSwitch                  string `json:"vrfNameNxPeerSwitch"`
	AutoGenConfigPeer                    *bool  `json:"autoGenConfigPeer"`
	AutoGenConfigDefaultVrf              *bool  `json:"autoGenConfigDefaultVrf"`
	AutoGenConfigNxPeerDefaultVrf        *bool  `json:"autoGenConfigNxPeerDefaultVrf"`
	TemplateConfigGenPeer                string `json:"templateConfigGenPeer"`
	RedistribEbgpRouteMapName            string `json:"redistribEbgpRouteMapName"`
	RedistributeRouteServer              *bool  `json:"redistributeRouteServer"`
	RouteServerRoutingTag                string `json:"routeServerRoutingTag"`
	RoutingTag                           string `json:"routingTag"`
	DciTracking                          *bool  `json:"dciTracking"`
	DciTrackingEnableFlag                *bool  `json:"dciTrackingEnableFlag"`
	InheritTtagFabricSetting             *bool  `json:"inheritTtagFabricSetting"`
	OverrideFabricMacsec                 *bool  `json:"overrideFabricMacsec"`
	MacsecCipherSuite                    string `json:"macsecCipherSuite"`
	MacsecPrimaryKeyString               string `json:"macsecPrimaryKeyString"`
	MacsecPrimaryCryptographicAlgorithm  string `json:"macsecPrimaryCryptographicAlgorithm"`
	MacsecFallbackKeyString              string `json:"macsecFallbackKeyString"`
	MacsecFallbackCryptographicAlgorithm string `json:"macsecFallbackCryptographicAlgorithm"`
	Qkd                                  *bool  `json:"qkd"`
	IgnoreCertificate                    *bool  `json:"ignoreCertificate"`
	SrcKmeServerIp                       string `json:"srcKmeServerIp"`
	SrcMacsecKeyChainPrefix              string `json:"srcMacsecKeyChainPrefix"`
	SrcQkdProfileName                    string `json:"srcQkdProfileName"`
	SrcTrustpointLabel                   string `json:"srcTrustpointLabel"`
	DstKmeServerIp                       string `json:"dstKmeServerIp"`
	DstMacsecKeyChainPrefix              string `json:"dstMacsecKeyChainPrefix"`
	DstQkdProfileName                    string `json:"dstQkdProfileName"`
	DstTrustpointLabel                   string `json:"dstTrustpointLabel"`
	DhcpRelayOnSrcInterface              *bool  `json:"dhcpRelayOnSrcInterface"`
	DhcpRelayOnDstInterface              *bool  `json:"dhcpRelayOnDstInterface"`
	BfdEchoOnSrcInterface                *bool  `json:"bfdEchoOnSrcInterface"`
	BfdEchoOnDstInterface                *bool  `json:"bfdEchoOnDstInterface"`
	MtuType                              string `json:"mtuType"`
	Ipv4Pim                              *bool  `json:"ipv4Pim"`
	Ipv4Trm                              *bool  `json:"ipv4Trm"`
	Ipv6Pim                              *bool  `json:"ipv6Pim"`
	Ipv6Trm                              *bool  `json:"ipv6Trm"`
	IpRedirects                          *bool  `json:"ipRedirects"`
	SkipConfigGeneration                 *bool  `json:"skipConfigGeneration"`
	BpduGuard                            string `json:"bpduGuard"`
	PortTypeFast                         *bool  `json:"portTypeFast"`
	NativeVlan                           *int64 `json:"nativeVlan"`
	TrunkAllowedVlans                    string `json:"trunkAllowedVlans"`
	NetflowOnSrcInterface                *bool  `json:"netflowOnSrcInterface"`
	NetflowOnDstInterface                *bool  `json:"netflowOnDstInterface"`
	SrcNetflowMonitorName                string `json:"srcNetflowMonitorName"`
	DstNetflowMonitorName                string `json:"dstNetflowMonitorName"`
	TemplateName                         string `json:"templateName"`
}

func (v *LinkModel) SetModelData(jsonData *NDFCLinkModel) diag.Diagnostics {
	var err diag.Diagnostics
	err = nil

	if jsonData.LinkId != "" {
		v.LinkId = types.StringValue(jsonData.LinkId)
	} else {
		v.LinkId = types.StringNull()
	}

	if jsonData.SrcSwitchId != "" {
		v.SrcSwitchId = types.StringValue(jsonData.SrcSwitchId)
	} else {
		v.SrcSwitchId = types.StringNull()
	}

	if jsonData.DstSwitchId != "" {
		v.DstSwitchId = types.StringValue(jsonData.DstSwitchId)
	} else {
		v.DstSwitchId = types.StringNull()
	}

	if jsonData.SrcFabricName != "" {
		v.SrcFabricName = types.StringValue(jsonData.SrcFabricName)
	} else {
		v.SrcFabricName = types.StringNull()
	}

	if jsonData.SrcSwitchName != "" {
		v.SrcSwitchName = types.StringValue(jsonData.SrcSwitchName)
	} else {
		v.SrcSwitchName = types.StringNull()
	}

	if jsonData.SrcInterfaceName != "" {
		v.SrcInterfaceName = types.StringValue(jsonData.SrcInterfaceName)
	} else {
		v.SrcInterfaceName = types.StringNull()
	}

	if jsonData.DstFabricName != "" {
		v.DstFabricName = types.StringValue(jsonData.DstFabricName)
	} else {
		v.DstFabricName = types.StringNull()
	}

	if jsonData.DstSwitchName != "" {
		v.DstSwitchName = types.StringValue(jsonData.DstSwitchName)
	} else {
		v.DstSwitchName = types.StringNull()
	}

	if jsonData.DstInterfaceName != "" {
		v.DstInterfaceName = types.StringValue(jsonData.DstInterfaceName)
	} else {
		v.DstInterfaceName = types.StringNull()
	}

	if jsonData.SrcClusterName != "" {
		v.SrcClusterName = types.StringValue(jsonData.SrcClusterName)
	} else {
		v.SrcClusterName = types.StringNull()
	}

	if jsonData.DstClusterName != "" {
		v.DstClusterName = types.StringValue(jsonData.DstClusterName)
	} else {
		v.DstClusterName = types.StringNull()
	}

	if jsonData.LinkType != "" {
		v.LinkType = types.StringValue(jsonData.LinkType)
	} else {
		v.LinkType = types.StringNull()
	}

	if jsonData.ConfigData.PolicyType != "" {
		v.PolicyType = types.StringValue(jsonData.ConfigData.PolicyType)

	} else {
		v.PolicyType = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcInterfaceDescription != "" {
		v.SrcInterfaceDescription = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcInterfaceDescription)

	} else {
		v.SrcInterfaceDescription = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcInterfaceConfig != "" {
		v.SrcInterfaceConfig = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcInterfaceConfig)

	} else {
		v.SrcInterfaceConfig = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstInterfaceDescription != "" {
		v.DstInterfaceDescription = types.StringValue(jsonData.ConfigData.TemplateInputs.DstInterfaceDescription)

	} else {
		v.DstInterfaceDescription = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstInterfaceConfig != "" {
		v.DstInterfaceConfig = types.StringValue(jsonData.ConfigData.TemplateInputs.DstInterfaceConfig)

	} else {
		v.DstInterfaceConfig = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.InterfaceAdminState != nil {
		v.InterfaceAdminState = types.BoolValue(*jsonData.ConfigData.TemplateInputs.InterfaceAdminState)

	} else {
		v.InterfaceAdminState = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.Mtu != nil {
		v.Mtu = types.Int64Value(*jsonData.ConfigData.TemplateInputs.Mtu)

	} else {
		v.Mtu = types.Int64Null()
	}

	if jsonData.ConfigData.TemplateInputs.LinkMtu != nil {
		v.LinkMtu = types.Int64Value(*jsonData.ConfigData.TemplateInputs.LinkMtu)

	} else {
		v.LinkMtu = types.Int64Null()
	}

	if jsonData.ConfigData.TemplateInputs.Speed != "" {
		v.Speed = types.StringValue(jsonData.ConfigData.TemplateInputs.Speed)

	} else {
		v.Speed = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.Fec != "" {
		v.Fec = types.StringValue(jsonData.ConfigData.TemplateInputs.Fec)

	} else {
		v.Fec = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.Macsec != nil {
		v.Macsec = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Macsec)

	} else {
		v.Macsec = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcIp != "" {
		v.SrcIp = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcIp)

	} else {
		v.SrcIp = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcIpv6 != "" {
		v.SrcIpv6 = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcIpv6)

	} else {
		v.SrcIpv6 = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcIpAddress != "" {
		v.SrcIpAddress = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcIpAddress)

	} else {
		v.SrcIpAddress = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcIpAddressMask != "" {
		v.SrcIpAddressMask = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcIpAddressMask)

	} else {
		v.SrcIpAddressMask = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcIpv6AddressMask != "" {
		v.SrcIpv6AddressMask = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcIpv6AddressMask)

	} else {
		v.SrcIpv6AddressMask = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIp != "" {
		v.DstIp = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIp)

	} else {
		v.DstIp = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIpv6 != "" {
		v.DstIpv6 = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIpv6)

	} else {
		v.DstIpv6 = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIpAddress != "" {
		v.DstIpAddress = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIpAddress)

	} else {
		v.DstIpAddress = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIpAddressMask != "" {
		v.DstIpAddressMask = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIpAddressMask)

	} else {
		v.DstIpAddressMask = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIpv6Address != "" {
		v.DstIpv6Address = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIpv6Address)

	} else {
		v.DstIpv6Address = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstIpv6AddressMask != "" {
		v.DstIpv6AddressMask = types.StringValue(jsonData.ConfigData.TemplateInputs.DstIpv6AddressMask)

	} else {
		v.DstIpv6AddressMask = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcEbgpAsn != "" {
		v.SrcEbgpAsn = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcEbgpAsn)

	} else {
		v.SrcEbgpAsn = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstEbgpAsn != "" {
		v.DstEbgpAsn = types.StringValue(jsonData.ConfigData.TemplateInputs.DstEbgpAsn)

	} else {
		v.DstEbgpAsn = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpPassword != "" {
		v.EbgpPassword = types.StringValue(jsonData.ConfigData.TemplateInputs.EbgpPassword)

	} else {
		v.EbgpPassword = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpAuthKeyEncryptionType != "" {
		v.EbgpAuthKeyEncryptionType = types.StringValue(jsonData.ConfigData.TemplateInputs.EbgpAuthKeyEncryptionType)

	} else {
		v.EbgpAuthKeyEncryptionType = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpBfd != nil {
		v.EbgpBfd = types.BoolValue(*jsonData.ConfigData.TemplateInputs.EbgpBfd)

	} else {
		v.EbgpBfd = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpLogNeighborChange != nil {
		v.EbgpLogNeighborChange = types.BoolValue(*jsonData.ConfigData.TemplateInputs.EbgpLogNeighborChange)

	} else {
		v.EbgpLogNeighborChange = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpMaximumPaths != nil {
		v.EbgpMaximumPaths = types.Int64Value(*jsonData.ConfigData.TemplateInputs.EbgpMaximumPaths)

	} else {
		v.EbgpMaximumPaths = types.Int64Null()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpMultihop != nil {
		v.EbgpMultihop = types.Int64Value(*jsonData.ConfigData.TemplateInputs.EbgpMultihop)

	} else {
		v.EbgpMultihop = types.Int64Null()
	}

	if jsonData.ConfigData.TemplateInputs.EbgpSendComboth != nil {
		v.EbgpSendComboth = types.BoolValue(*jsonData.ConfigData.TemplateInputs.EbgpSendComboth)

	} else {
		v.EbgpSendComboth = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.EnableEbgpPassword != nil {
		v.EnableEbgpPassword = types.BoolValue(*jsonData.ConfigData.TemplateInputs.EnableEbgpPassword)

	} else {
		v.EnableEbgpPassword = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.InheritEbgpPasswordMsdSettings != nil {
		v.InheritEbgpPasswordMsdSettings = types.BoolValue(*jsonData.ConfigData.TemplateInputs.InheritEbgpPasswordMsdSettings)

	} else {
		v.InheritEbgpPasswordMsdSettings = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.DefaultVrfEbgpNeighborPassword != "" {
		v.DefaultVrfEbgpNeighborPassword = types.StringValue(jsonData.ConfigData.TemplateInputs.DefaultVrfEbgpNeighborPassword)

	} else {
		v.DefaultVrfEbgpNeighborPassword = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.InterfaceVrf != "" {
		v.InterfaceVrf = types.StringValue(jsonData.ConfigData.TemplateInputs.InterfaceVrf)

	} else {
		v.InterfaceVrf = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcVrfName != "" {
		v.SrcVrfName = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcVrfName)

	} else {
		v.SrcVrfName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstVrfName != "" {
		v.DstVrfName = types.StringValue(jsonData.ConfigData.TemplateInputs.DstVrfName)

	} else {
		v.DstVrfName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.VrfNameNxPeerSwitch != "" {
		v.VrfNameNxPeerSwitch = types.StringValue(jsonData.ConfigData.TemplateInputs.VrfNameNxPeerSwitch)

	} else {
		v.VrfNameNxPeerSwitch = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.AutoGenConfigPeer != nil {
		v.AutoGenConfigPeer = types.BoolValue(*jsonData.ConfigData.TemplateInputs.AutoGenConfigPeer)

	} else {
		v.AutoGenConfigPeer = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.AutoGenConfigDefaultVrf != nil {
		v.AutoGenConfigDefaultVrf = types.BoolValue(*jsonData.ConfigData.TemplateInputs.AutoGenConfigDefaultVrf)

	} else {
		v.AutoGenConfigDefaultVrf = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.AutoGenConfigNxPeerDefaultVrf != nil {
		v.AutoGenConfigNxPeerDefaultVrf = types.BoolValue(*jsonData.ConfigData.TemplateInputs.AutoGenConfigNxPeerDefaultVrf)

	} else {
		v.AutoGenConfigNxPeerDefaultVrf = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.TemplateConfigGenPeer != "" {
		v.TemplateConfigGenPeer = types.StringValue(jsonData.ConfigData.TemplateInputs.TemplateConfigGenPeer)

	} else {
		v.TemplateConfigGenPeer = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.RedistribEbgpRouteMapName != "" {
		v.RedistribEbgpRouteMapName = types.StringValue(jsonData.ConfigData.TemplateInputs.RedistribEbgpRouteMapName)

	} else {
		v.RedistribEbgpRouteMapName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.RedistributeRouteServer != nil {
		v.RedistributeRouteServer = types.BoolValue(*jsonData.ConfigData.TemplateInputs.RedistributeRouteServer)

	} else {
		v.RedistributeRouteServer = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.RouteServerRoutingTag != "" {
		v.RouteServerRoutingTag = types.StringValue(jsonData.ConfigData.TemplateInputs.RouteServerRoutingTag)

	} else {
		v.RouteServerRoutingTag = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.RoutingTag != "" {
		v.RoutingTag = types.StringValue(jsonData.ConfigData.TemplateInputs.RoutingTag)

	} else {
		v.RoutingTag = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DciTracking != nil {
		v.DciTracking = types.BoolValue(*jsonData.ConfigData.TemplateInputs.DciTracking)

	} else {
		v.DciTracking = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.DciTrackingEnableFlag != nil {
		v.DciTrackingEnableFlag = types.BoolValue(*jsonData.ConfigData.TemplateInputs.DciTrackingEnableFlag)

	} else {
		v.DciTrackingEnableFlag = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.InheritTtagFabricSetting != nil {
		v.InheritTtagFabricSetting = types.BoolValue(*jsonData.ConfigData.TemplateInputs.InheritTtagFabricSetting)

	} else {
		v.InheritTtagFabricSetting = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.OverrideFabricMacsec != nil {
		v.OverrideFabricMacsec = types.BoolValue(*jsonData.ConfigData.TemplateInputs.OverrideFabricMacsec)

	} else {
		v.OverrideFabricMacsec = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.MacsecCipherSuite != "" {
		v.MacsecCipherSuite = types.StringValue(jsonData.ConfigData.TemplateInputs.MacsecCipherSuite)

	} else {
		v.MacsecCipherSuite = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.MacsecPrimaryKeyString != "" {
		v.MacsecPrimaryKeyString = types.StringValue(jsonData.ConfigData.TemplateInputs.MacsecPrimaryKeyString)

	} else {
		v.MacsecPrimaryKeyString = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.MacsecPrimaryCryptographicAlgorithm != "" {
		v.MacsecPrimaryCryptographicAlgorithm = types.StringValue(jsonData.ConfigData.TemplateInputs.MacsecPrimaryCryptographicAlgorithm)

	} else {
		v.MacsecPrimaryCryptographicAlgorithm = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.MacsecFallbackKeyString != "" {
		v.MacsecFallbackKeyString = types.StringValue(jsonData.ConfigData.TemplateInputs.MacsecFallbackKeyString)

	} else {
		v.MacsecFallbackKeyString = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.MacsecFallbackCryptographicAlgorithm != "" {
		v.MacsecFallbackCryptographicAlgorithm = types.StringValue(jsonData.ConfigData.TemplateInputs.MacsecFallbackCryptographicAlgorithm)

	} else {
		v.MacsecFallbackCryptographicAlgorithm = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.Qkd != nil {
		v.Qkd = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Qkd)

	} else {
		v.Qkd = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.IgnoreCertificate != nil {
		v.IgnoreCertificate = types.BoolValue(*jsonData.ConfigData.TemplateInputs.IgnoreCertificate)

	} else {
		v.IgnoreCertificate = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcKmeServerIp != "" {
		v.SrcKmeServerIp = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcKmeServerIp)

	} else {
		v.SrcKmeServerIp = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcMacsecKeyChainPrefix != "" {
		v.SrcMacsecKeyChainPrefix = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcMacsecKeyChainPrefix)

	} else {
		v.SrcMacsecKeyChainPrefix = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcQkdProfileName != "" {
		v.SrcQkdProfileName = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcQkdProfileName)

	} else {
		v.SrcQkdProfileName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcTrustpointLabel != "" {
		v.SrcTrustpointLabel = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcTrustpointLabel)

	} else {
		v.SrcTrustpointLabel = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstKmeServerIp != "" {
		v.DstKmeServerIp = types.StringValue(jsonData.ConfigData.TemplateInputs.DstKmeServerIp)

	} else {
		v.DstKmeServerIp = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstMacsecKeyChainPrefix != "" {
		v.DstMacsecKeyChainPrefix = types.StringValue(jsonData.ConfigData.TemplateInputs.DstMacsecKeyChainPrefix)

	} else {
		v.DstMacsecKeyChainPrefix = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstQkdProfileName != "" {
		v.DstQkdProfileName = types.StringValue(jsonData.ConfigData.TemplateInputs.DstQkdProfileName)

	} else {
		v.DstQkdProfileName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstTrustpointLabel != "" {
		v.DstTrustpointLabel = types.StringValue(jsonData.ConfigData.TemplateInputs.DstTrustpointLabel)

	} else {
		v.DstTrustpointLabel = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DhcpRelayOnSrcInterface != nil {
		v.DhcpRelayOnSrcInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.DhcpRelayOnSrcInterface)

	} else {
		v.DhcpRelayOnSrcInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.DhcpRelayOnDstInterface != nil {
		v.DhcpRelayOnDstInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.DhcpRelayOnDstInterface)

	} else {
		v.DhcpRelayOnDstInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.BfdEchoOnSrcInterface != nil {
		v.BfdEchoOnSrcInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.BfdEchoOnSrcInterface)

	} else {
		v.BfdEchoOnSrcInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.BfdEchoOnDstInterface != nil {
		v.BfdEchoOnDstInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.BfdEchoOnDstInterface)

	} else {
		v.BfdEchoOnDstInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.MtuType != "" {
		v.MtuType = types.StringValue(jsonData.ConfigData.TemplateInputs.MtuType)

	} else {
		v.MtuType = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.Ipv4Pim != nil {
		v.Ipv4Pim = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Ipv4Pim)

	} else {
		v.Ipv4Pim = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.Ipv4Trm != nil {
		v.Ipv4Trm = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Ipv4Trm)

	} else {
		v.Ipv4Trm = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.Ipv6Pim != nil {
		v.Ipv6Pim = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Ipv6Pim)

	} else {
		v.Ipv6Pim = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.Ipv6Trm != nil {
		v.Ipv6Trm = types.BoolValue(*jsonData.ConfigData.TemplateInputs.Ipv6Trm)

	} else {
		v.Ipv6Trm = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.IpRedirects != nil {
		v.IpRedirects = types.BoolValue(*jsonData.ConfigData.TemplateInputs.IpRedirects)

	} else {
		v.IpRedirects = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.SkipConfigGeneration != nil {
		v.SkipConfigGeneration = types.BoolValue(*jsonData.ConfigData.TemplateInputs.SkipConfigGeneration)

	} else {
		v.SkipConfigGeneration = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.BpduGuard != "" {
		v.BpduGuard = types.StringValue(jsonData.ConfigData.TemplateInputs.BpduGuard)

	} else {
		v.BpduGuard = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.PortTypeFast != nil {
		v.PortTypeFast = types.BoolValue(*jsonData.ConfigData.TemplateInputs.PortTypeFast)

	} else {
		v.PortTypeFast = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.NativeVlan != nil {
		v.NativeVlan = types.Int64Value(*jsonData.ConfigData.TemplateInputs.NativeVlan)

	} else {
		v.NativeVlan = types.Int64Null()
	}

	if jsonData.ConfigData.TemplateInputs.TrunkAllowedVlans != "" {
		v.TrunkAllowedVlans = types.StringValue(jsonData.ConfigData.TemplateInputs.TrunkAllowedVlans)

	} else {
		v.TrunkAllowedVlans = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.NetflowOnSrcInterface != nil {
		v.NetflowOnSrcInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.NetflowOnSrcInterface)

	} else {
		v.NetflowOnSrcInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.NetflowOnDstInterface != nil {
		v.NetflowOnDstInterface = types.BoolValue(*jsonData.ConfigData.TemplateInputs.NetflowOnDstInterface)

	} else {
		v.NetflowOnDstInterface = types.BoolNull()
	}

	if jsonData.ConfigData.TemplateInputs.SrcNetflowMonitorName != "" {
		v.SrcNetflowMonitorName = types.StringValue(jsonData.ConfigData.TemplateInputs.SrcNetflowMonitorName)

	} else {
		v.SrcNetflowMonitorName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.DstNetflowMonitorName != "" {
		v.DstNetflowMonitorName = types.StringValue(jsonData.ConfigData.TemplateInputs.DstNetflowMonitorName)

	} else {
		v.DstNetflowMonitorName = types.StringNull()
	}

	if jsonData.ConfigData.TemplateInputs.TemplateName != "" {
		v.TemplateName = types.StringValue(jsonData.ConfigData.TemplateInputs.TemplateName)

	} else {
		v.TemplateName = types.StringNull()
	}

	return err
}

func (v LinkModel) GetModelData() *NDFCLinkModel {
	var data = new(NDFCLinkModel)

	//MARSHAL_BODY

	if !v.SrcSwitchId.IsNull() && !v.SrcSwitchId.IsUnknown() {
		data.SrcSwitchId = v.SrcSwitchId.ValueString()
	} else {
		data.SrcSwitchId = ""
	}

	if !v.DstSwitchId.IsNull() && !v.DstSwitchId.IsUnknown() {
		data.DstSwitchId = v.DstSwitchId.ValueString()
	} else {
		data.DstSwitchId = ""
	}

	if !v.SrcFabricName.IsNull() && !v.SrcFabricName.IsUnknown() {
		data.SrcFabricName = v.SrcFabricName.ValueString()
	} else {
		data.SrcFabricName = ""
	}

	if !v.SrcSwitchName.IsNull() && !v.SrcSwitchName.IsUnknown() {
		data.SrcSwitchName = v.SrcSwitchName.ValueString()
	} else {
		data.SrcSwitchName = ""
	}

	if !v.SrcInterfaceName.IsNull() && !v.SrcInterfaceName.IsUnknown() {
		data.SrcInterfaceName = v.SrcInterfaceName.ValueString()
	} else {
		data.SrcInterfaceName = ""
	}

	if !v.DstFabricName.IsNull() && !v.DstFabricName.IsUnknown() {
		data.DstFabricName = v.DstFabricName.ValueString()
	} else {
		data.DstFabricName = ""
	}

	if !v.DstSwitchName.IsNull() && !v.DstSwitchName.IsUnknown() {
		data.DstSwitchName = v.DstSwitchName.ValueString()
	} else {
		data.DstSwitchName = ""
	}

	if !v.DstInterfaceName.IsNull() && !v.DstInterfaceName.IsUnknown() {
		data.DstInterfaceName = v.DstInterfaceName.ValueString()
	} else {
		data.DstInterfaceName = ""
	}

	if !v.SrcClusterName.IsNull() && !v.SrcClusterName.IsUnknown() {
		data.SrcClusterName = v.SrcClusterName.ValueString()
	} else {
		data.SrcClusterName = ""
	}

	if !v.DstClusterName.IsNull() && !v.DstClusterName.IsUnknown() {
		data.DstClusterName = v.DstClusterName.ValueString()
	} else {
		data.DstClusterName = ""
	}

	if !v.PolicyType.IsNull() && !v.PolicyType.IsUnknown() {
		data.ConfigData.PolicyType = v.PolicyType.ValueString()
	} else {
		data.ConfigData.PolicyType = ""
	}

	if !v.SrcInterfaceDescription.IsNull() && !v.SrcInterfaceDescription.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcInterfaceDescription = v.SrcInterfaceDescription.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcInterfaceDescription = ""
	}

	if !v.SrcInterfaceConfig.IsNull() && !v.SrcInterfaceConfig.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcInterfaceConfig = v.SrcInterfaceConfig.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcInterfaceConfig = ""
	}

	if !v.DstInterfaceDescription.IsNull() && !v.DstInterfaceDescription.IsUnknown() {
		data.ConfigData.TemplateInputs.DstInterfaceDescription = v.DstInterfaceDescription.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstInterfaceDescription = ""
	}

	if !v.DstInterfaceConfig.IsNull() && !v.DstInterfaceConfig.IsUnknown() {
		data.ConfigData.TemplateInputs.DstInterfaceConfig = v.DstInterfaceConfig.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstInterfaceConfig = ""
	}

	if !v.InterfaceAdminState.IsNull() && !v.InterfaceAdminState.IsUnknown() {
		data.ConfigData.TemplateInputs.InterfaceAdminState = new(bool)
		*data.ConfigData.TemplateInputs.InterfaceAdminState = v.InterfaceAdminState.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.InterfaceAdminState = nil
	}

	if !v.Mtu.IsNull() && !v.Mtu.IsUnknown() {
		data.ConfigData.TemplateInputs.Mtu = new(int64)
		*data.ConfigData.TemplateInputs.Mtu = v.Mtu.ValueInt64()

	} else {
		data.ConfigData.TemplateInputs.Mtu = nil
	}

	if !v.LinkMtu.IsNull() && !v.LinkMtu.IsUnknown() {
		data.ConfigData.TemplateInputs.LinkMtu = new(int64)
		*data.ConfigData.TemplateInputs.LinkMtu = v.LinkMtu.ValueInt64()

	} else {
		data.ConfigData.TemplateInputs.LinkMtu = nil
	}

	if !v.Speed.IsNull() && !v.Speed.IsUnknown() {
		data.ConfigData.TemplateInputs.Speed = v.Speed.ValueString()
	} else {
		data.ConfigData.TemplateInputs.Speed = ""
	}

	if !v.Fec.IsNull() && !v.Fec.IsUnknown() {
		data.ConfigData.TemplateInputs.Fec = v.Fec.ValueString()
	} else {
		data.ConfigData.TemplateInputs.Fec = ""
	}

	if !v.Macsec.IsNull() && !v.Macsec.IsUnknown() {
		data.ConfigData.TemplateInputs.Macsec = new(bool)
		*data.ConfigData.TemplateInputs.Macsec = v.Macsec.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Macsec = nil
	}

	if !v.SrcIp.IsNull() && !v.SrcIp.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcIp = v.SrcIp.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcIp = ""
	}

	if !v.SrcIpv6.IsNull() && !v.SrcIpv6.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcIpv6 = v.SrcIpv6.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcIpv6 = ""
	}

	if !v.SrcIpAddress.IsNull() && !v.SrcIpAddress.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcIpAddress = v.SrcIpAddress.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcIpAddress = ""
	}

	if !v.SrcIpAddressMask.IsNull() && !v.SrcIpAddressMask.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcIpAddressMask = v.SrcIpAddressMask.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcIpAddressMask = ""
	}

	if !v.SrcIpv6AddressMask.IsNull() && !v.SrcIpv6AddressMask.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcIpv6AddressMask = v.SrcIpv6AddressMask.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcIpv6AddressMask = ""
	}

	if !v.DstIp.IsNull() && !v.DstIp.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIp = v.DstIp.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIp = ""
	}

	if !v.DstIpv6.IsNull() && !v.DstIpv6.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIpv6 = v.DstIpv6.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIpv6 = ""
	}

	if !v.DstIpAddress.IsNull() && !v.DstIpAddress.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIpAddress = v.DstIpAddress.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIpAddress = ""
	}

	if !v.DstIpAddressMask.IsNull() && !v.DstIpAddressMask.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIpAddressMask = v.DstIpAddressMask.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIpAddressMask = ""
	}

	if !v.DstIpv6Address.IsNull() && !v.DstIpv6Address.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIpv6Address = v.DstIpv6Address.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIpv6Address = ""
	}

	if !v.DstIpv6AddressMask.IsNull() && !v.DstIpv6AddressMask.IsUnknown() {
		data.ConfigData.TemplateInputs.DstIpv6AddressMask = v.DstIpv6AddressMask.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstIpv6AddressMask = ""
	}

	if !v.SrcEbgpAsn.IsNull() && !v.SrcEbgpAsn.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcEbgpAsn = v.SrcEbgpAsn.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcEbgpAsn = ""
	}

	if !v.DstEbgpAsn.IsNull() && !v.DstEbgpAsn.IsUnknown() {
		data.ConfigData.TemplateInputs.DstEbgpAsn = v.DstEbgpAsn.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstEbgpAsn = ""
	}

	if !v.EbgpPassword.IsNull() && !v.EbgpPassword.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpPassword = v.EbgpPassword.ValueString()
	} else {
		data.ConfigData.TemplateInputs.EbgpPassword = ""
	}

	if !v.EbgpAuthKeyEncryptionType.IsNull() && !v.EbgpAuthKeyEncryptionType.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpAuthKeyEncryptionType = v.EbgpAuthKeyEncryptionType.ValueString()
	} else {
		data.ConfigData.TemplateInputs.EbgpAuthKeyEncryptionType = ""
	}

	if !v.EbgpBfd.IsNull() && !v.EbgpBfd.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpBfd = new(bool)
		*data.ConfigData.TemplateInputs.EbgpBfd = v.EbgpBfd.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.EbgpBfd = nil
	}

	if !v.EbgpLogNeighborChange.IsNull() && !v.EbgpLogNeighborChange.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpLogNeighborChange = new(bool)
		*data.ConfigData.TemplateInputs.EbgpLogNeighborChange = v.EbgpLogNeighborChange.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.EbgpLogNeighborChange = nil
	}

	if !v.EbgpMaximumPaths.IsNull() && !v.EbgpMaximumPaths.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpMaximumPaths = new(int64)
		*data.ConfigData.TemplateInputs.EbgpMaximumPaths = v.EbgpMaximumPaths.ValueInt64()

	} else {
		data.ConfigData.TemplateInputs.EbgpMaximumPaths = nil
	}

	if !v.EbgpMultihop.IsNull() && !v.EbgpMultihop.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpMultihop = new(int64)
		*data.ConfigData.TemplateInputs.EbgpMultihop = v.EbgpMultihop.ValueInt64()

	} else {
		data.ConfigData.TemplateInputs.EbgpMultihop = nil
	}

	if !v.EbgpSendComboth.IsNull() && !v.EbgpSendComboth.IsUnknown() {
		data.ConfigData.TemplateInputs.EbgpSendComboth = new(bool)
		*data.ConfigData.TemplateInputs.EbgpSendComboth = v.EbgpSendComboth.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.EbgpSendComboth = nil
	}

	if !v.EnableEbgpPassword.IsNull() && !v.EnableEbgpPassword.IsUnknown() {
		data.ConfigData.TemplateInputs.EnableEbgpPassword = new(bool)
		*data.ConfigData.TemplateInputs.EnableEbgpPassword = v.EnableEbgpPassword.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.EnableEbgpPassword = nil
	}

	if !v.InheritEbgpPasswordMsdSettings.IsNull() && !v.InheritEbgpPasswordMsdSettings.IsUnknown() {
		data.ConfigData.TemplateInputs.InheritEbgpPasswordMsdSettings = new(bool)
		*data.ConfigData.TemplateInputs.InheritEbgpPasswordMsdSettings = v.InheritEbgpPasswordMsdSettings.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.InheritEbgpPasswordMsdSettings = nil
	}

	if !v.DefaultVrfEbgpNeighborPassword.IsNull() && !v.DefaultVrfEbgpNeighborPassword.IsUnknown() {
		data.ConfigData.TemplateInputs.DefaultVrfEbgpNeighborPassword = v.DefaultVrfEbgpNeighborPassword.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DefaultVrfEbgpNeighborPassword = ""
	}

	if !v.InterfaceVrf.IsNull() && !v.InterfaceVrf.IsUnknown() {
		data.ConfigData.TemplateInputs.InterfaceVrf = v.InterfaceVrf.ValueString()
	} else {
		data.ConfigData.TemplateInputs.InterfaceVrf = ""
	}

	if !v.SrcVrfName.IsNull() && !v.SrcVrfName.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcVrfName = v.SrcVrfName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcVrfName = ""
	}

	if !v.DstVrfName.IsNull() && !v.DstVrfName.IsUnknown() {
		data.ConfigData.TemplateInputs.DstVrfName = v.DstVrfName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstVrfName = ""
	}

	if !v.VrfNameNxPeerSwitch.IsNull() && !v.VrfNameNxPeerSwitch.IsUnknown() {
		data.ConfigData.TemplateInputs.VrfNameNxPeerSwitch = v.VrfNameNxPeerSwitch.ValueString()
	} else {
		data.ConfigData.TemplateInputs.VrfNameNxPeerSwitch = ""
	}

	if !v.AutoGenConfigPeer.IsNull() && !v.AutoGenConfigPeer.IsUnknown() {
		data.ConfigData.TemplateInputs.AutoGenConfigPeer = new(bool)
		*data.ConfigData.TemplateInputs.AutoGenConfigPeer = v.AutoGenConfigPeer.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.AutoGenConfigPeer = nil
	}

	if !v.AutoGenConfigDefaultVrf.IsNull() && !v.AutoGenConfigDefaultVrf.IsUnknown() {
		data.ConfigData.TemplateInputs.AutoGenConfigDefaultVrf = new(bool)
		*data.ConfigData.TemplateInputs.AutoGenConfigDefaultVrf = v.AutoGenConfigDefaultVrf.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.AutoGenConfigDefaultVrf = nil
	}

	if !v.AutoGenConfigNxPeerDefaultVrf.IsNull() && !v.AutoGenConfigNxPeerDefaultVrf.IsUnknown() {
		data.ConfigData.TemplateInputs.AutoGenConfigNxPeerDefaultVrf = new(bool)
		*data.ConfigData.TemplateInputs.AutoGenConfigNxPeerDefaultVrf = v.AutoGenConfigNxPeerDefaultVrf.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.AutoGenConfigNxPeerDefaultVrf = nil
	}

	if !v.TemplateConfigGenPeer.IsNull() && !v.TemplateConfigGenPeer.IsUnknown() {
		data.ConfigData.TemplateInputs.TemplateConfigGenPeer = v.TemplateConfigGenPeer.ValueString()
	} else {
		data.ConfigData.TemplateInputs.TemplateConfigGenPeer = ""
	}

	if !v.RedistribEbgpRouteMapName.IsNull() && !v.RedistribEbgpRouteMapName.IsUnknown() {
		data.ConfigData.TemplateInputs.RedistribEbgpRouteMapName = v.RedistribEbgpRouteMapName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.RedistribEbgpRouteMapName = ""
	}

	if !v.RedistributeRouteServer.IsNull() && !v.RedistributeRouteServer.IsUnknown() {
		data.ConfigData.TemplateInputs.RedistributeRouteServer = new(bool)
		*data.ConfigData.TemplateInputs.RedistributeRouteServer = v.RedistributeRouteServer.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.RedistributeRouteServer = nil
	}

	if !v.RouteServerRoutingTag.IsNull() && !v.RouteServerRoutingTag.IsUnknown() {
		data.ConfigData.TemplateInputs.RouteServerRoutingTag = v.RouteServerRoutingTag.ValueString()
	} else {
		data.ConfigData.TemplateInputs.RouteServerRoutingTag = ""
	}

	if !v.RoutingTag.IsNull() && !v.RoutingTag.IsUnknown() {
		data.ConfigData.TemplateInputs.RoutingTag = v.RoutingTag.ValueString()
	} else {
		data.ConfigData.TemplateInputs.RoutingTag = ""
	}

	if !v.DciTracking.IsNull() && !v.DciTracking.IsUnknown() {
		data.ConfigData.TemplateInputs.DciTracking = new(bool)
		*data.ConfigData.TemplateInputs.DciTracking = v.DciTracking.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.DciTracking = nil
	}

	if !v.DciTrackingEnableFlag.IsNull() && !v.DciTrackingEnableFlag.IsUnknown() {
		data.ConfigData.TemplateInputs.DciTrackingEnableFlag = new(bool)
		*data.ConfigData.TemplateInputs.DciTrackingEnableFlag = v.DciTrackingEnableFlag.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.DciTrackingEnableFlag = nil
	}

	if !v.InheritTtagFabricSetting.IsNull() && !v.InheritTtagFabricSetting.IsUnknown() {
		data.ConfigData.TemplateInputs.InheritTtagFabricSetting = new(bool)
		*data.ConfigData.TemplateInputs.InheritTtagFabricSetting = v.InheritTtagFabricSetting.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.InheritTtagFabricSetting = nil
	}

	if !v.OverrideFabricMacsec.IsNull() && !v.OverrideFabricMacsec.IsUnknown() {
		data.ConfigData.TemplateInputs.OverrideFabricMacsec = new(bool)
		*data.ConfigData.TemplateInputs.OverrideFabricMacsec = v.OverrideFabricMacsec.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.OverrideFabricMacsec = nil
	}

	if !v.MacsecCipherSuite.IsNull() && !v.MacsecCipherSuite.IsUnknown() {
		data.ConfigData.TemplateInputs.MacsecCipherSuite = v.MacsecCipherSuite.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MacsecCipherSuite = ""
	}

	if !v.MacsecPrimaryKeyString.IsNull() && !v.MacsecPrimaryKeyString.IsUnknown() {
		data.ConfigData.TemplateInputs.MacsecPrimaryKeyString = v.MacsecPrimaryKeyString.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MacsecPrimaryKeyString = ""
	}

	if !v.MacsecPrimaryCryptographicAlgorithm.IsNull() && !v.MacsecPrimaryCryptographicAlgorithm.IsUnknown() {
		data.ConfigData.TemplateInputs.MacsecPrimaryCryptographicAlgorithm = v.MacsecPrimaryCryptographicAlgorithm.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MacsecPrimaryCryptographicAlgorithm = ""
	}

	if !v.MacsecFallbackKeyString.IsNull() && !v.MacsecFallbackKeyString.IsUnknown() {
		data.ConfigData.TemplateInputs.MacsecFallbackKeyString = v.MacsecFallbackKeyString.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MacsecFallbackKeyString = ""
	}

	if !v.MacsecFallbackCryptographicAlgorithm.IsNull() && !v.MacsecFallbackCryptographicAlgorithm.IsUnknown() {
		data.ConfigData.TemplateInputs.MacsecFallbackCryptographicAlgorithm = v.MacsecFallbackCryptographicAlgorithm.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MacsecFallbackCryptographicAlgorithm = ""
	}

	if !v.Qkd.IsNull() && !v.Qkd.IsUnknown() {
		data.ConfigData.TemplateInputs.Qkd = new(bool)
		*data.ConfigData.TemplateInputs.Qkd = v.Qkd.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Qkd = nil
	}

	if !v.IgnoreCertificate.IsNull() && !v.IgnoreCertificate.IsUnknown() {
		data.ConfigData.TemplateInputs.IgnoreCertificate = new(bool)
		*data.ConfigData.TemplateInputs.IgnoreCertificate = v.IgnoreCertificate.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.IgnoreCertificate = nil
	}

	if !v.SrcKmeServerIp.IsNull() && !v.SrcKmeServerIp.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcKmeServerIp = v.SrcKmeServerIp.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcKmeServerIp = ""
	}

	if !v.SrcMacsecKeyChainPrefix.IsNull() && !v.SrcMacsecKeyChainPrefix.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcMacsecKeyChainPrefix = v.SrcMacsecKeyChainPrefix.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcMacsecKeyChainPrefix = ""
	}

	if !v.SrcQkdProfileName.IsNull() && !v.SrcQkdProfileName.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcQkdProfileName = v.SrcQkdProfileName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcQkdProfileName = ""
	}

	if !v.SrcTrustpointLabel.IsNull() && !v.SrcTrustpointLabel.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcTrustpointLabel = v.SrcTrustpointLabel.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcTrustpointLabel = ""
	}

	if !v.DstKmeServerIp.IsNull() && !v.DstKmeServerIp.IsUnknown() {
		data.ConfigData.TemplateInputs.DstKmeServerIp = v.DstKmeServerIp.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstKmeServerIp = ""
	}

	if !v.DstMacsecKeyChainPrefix.IsNull() && !v.DstMacsecKeyChainPrefix.IsUnknown() {
		data.ConfigData.TemplateInputs.DstMacsecKeyChainPrefix = v.DstMacsecKeyChainPrefix.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstMacsecKeyChainPrefix = ""
	}

	if !v.DstQkdProfileName.IsNull() && !v.DstQkdProfileName.IsUnknown() {
		data.ConfigData.TemplateInputs.DstQkdProfileName = v.DstQkdProfileName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstQkdProfileName = ""
	}

	if !v.DstTrustpointLabel.IsNull() && !v.DstTrustpointLabel.IsUnknown() {
		data.ConfigData.TemplateInputs.DstTrustpointLabel = v.DstTrustpointLabel.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstTrustpointLabel = ""
	}

	if !v.DhcpRelayOnSrcInterface.IsNull() && !v.DhcpRelayOnSrcInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.DhcpRelayOnSrcInterface = new(bool)
		*data.ConfigData.TemplateInputs.DhcpRelayOnSrcInterface = v.DhcpRelayOnSrcInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.DhcpRelayOnSrcInterface = nil
	}

	if !v.DhcpRelayOnDstInterface.IsNull() && !v.DhcpRelayOnDstInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.DhcpRelayOnDstInterface = new(bool)
		*data.ConfigData.TemplateInputs.DhcpRelayOnDstInterface = v.DhcpRelayOnDstInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.DhcpRelayOnDstInterface = nil
	}

	if !v.BfdEchoOnSrcInterface.IsNull() && !v.BfdEchoOnSrcInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.BfdEchoOnSrcInterface = new(bool)
		*data.ConfigData.TemplateInputs.BfdEchoOnSrcInterface = v.BfdEchoOnSrcInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.BfdEchoOnSrcInterface = nil
	}

	if !v.BfdEchoOnDstInterface.IsNull() && !v.BfdEchoOnDstInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.BfdEchoOnDstInterface = new(bool)
		*data.ConfigData.TemplateInputs.BfdEchoOnDstInterface = v.BfdEchoOnDstInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.BfdEchoOnDstInterface = nil
	}

	if !v.MtuType.IsNull() && !v.MtuType.IsUnknown() {
		data.ConfigData.TemplateInputs.MtuType = v.MtuType.ValueString()
	} else {
		data.ConfigData.TemplateInputs.MtuType = ""
	}

	if !v.Ipv4Pim.IsNull() && !v.Ipv4Pim.IsUnknown() {
		data.ConfigData.TemplateInputs.Ipv4Pim = new(bool)
		*data.ConfigData.TemplateInputs.Ipv4Pim = v.Ipv4Pim.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Ipv4Pim = nil
	}

	if !v.Ipv4Trm.IsNull() && !v.Ipv4Trm.IsUnknown() {
		data.ConfigData.TemplateInputs.Ipv4Trm = new(bool)
		*data.ConfigData.TemplateInputs.Ipv4Trm = v.Ipv4Trm.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Ipv4Trm = nil
	}

	if !v.Ipv6Pim.IsNull() && !v.Ipv6Pim.IsUnknown() {
		data.ConfigData.TemplateInputs.Ipv6Pim = new(bool)
		*data.ConfigData.TemplateInputs.Ipv6Pim = v.Ipv6Pim.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Ipv6Pim = nil
	}

	if !v.Ipv6Trm.IsNull() && !v.Ipv6Trm.IsUnknown() {
		data.ConfigData.TemplateInputs.Ipv6Trm = new(bool)
		*data.ConfigData.TemplateInputs.Ipv6Trm = v.Ipv6Trm.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.Ipv6Trm = nil
	}

	if !v.IpRedirects.IsNull() && !v.IpRedirects.IsUnknown() {
		data.ConfigData.TemplateInputs.IpRedirects = new(bool)
		*data.ConfigData.TemplateInputs.IpRedirects = v.IpRedirects.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.IpRedirects = nil
	}

	if !v.SkipConfigGeneration.IsNull() && !v.SkipConfigGeneration.IsUnknown() {
		data.ConfigData.TemplateInputs.SkipConfigGeneration = new(bool)
		*data.ConfigData.TemplateInputs.SkipConfigGeneration = v.SkipConfigGeneration.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.SkipConfigGeneration = nil
	}

	if !v.BpduGuard.IsNull() && !v.BpduGuard.IsUnknown() {
		data.ConfigData.TemplateInputs.BpduGuard = v.BpduGuard.ValueString()
	} else {
		data.ConfigData.TemplateInputs.BpduGuard = ""
	}

	if !v.PortTypeFast.IsNull() && !v.PortTypeFast.IsUnknown() {
		data.ConfigData.TemplateInputs.PortTypeFast = new(bool)
		*data.ConfigData.TemplateInputs.PortTypeFast = v.PortTypeFast.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.PortTypeFast = nil
	}

	if !v.NativeVlan.IsNull() && !v.NativeVlan.IsUnknown() {
		data.ConfigData.TemplateInputs.NativeVlan = new(int64)
		*data.ConfigData.TemplateInputs.NativeVlan = v.NativeVlan.ValueInt64()

	} else {
		data.ConfigData.TemplateInputs.NativeVlan = nil
	}

	if !v.TrunkAllowedVlans.IsNull() && !v.TrunkAllowedVlans.IsUnknown() {
		data.ConfigData.TemplateInputs.TrunkAllowedVlans = v.TrunkAllowedVlans.ValueString()
	} else {
		data.ConfigData.TemplateInputs.TrunkAllowedVlans = ""
	}

	if !v.NetflowOnSrcInterface.IsNull() && !v.NetflowOnSrcInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.NetflowOnSrcInterface = new(bool)
		*data.ConfigData.TemplateInputs.NetflowOnSrcInterface = v.NetflowOnSrcInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.NetflowOnSrcInterface = nil
	}

	if !v.NetflowOnDstInterface.IsNull() && !v.NetflowOnDstInterface.IsUnknown() {
		data.ConfigData.TemplateInputs.NetflowOnDstInterface = new(bool)
		*data.ConfigData.TemplateInputs.NetflowOnDstInterface = v.NetflowOnDstInterface.ValueBool()
	} else {
		data.ConfigData.TemplateInputs.NetflowOnDstInterface = nil
	}

	if !v.SrcNetflowMonitorName.IsNull() && !v.SrcNetflowMonitorName.IsUnknown() {
		data.ConfigData.TemplateInputs.SrcNetflowMonitorName = v.SrcNetflowMonitorName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.SrcNetflowMonitorName = ""
	}

	if !v.DstNetflowMonitorName.IsNull() && !v.DstNetflowMonitorName.IsUnknown() {
		data.ConfigData.TemplateInputs.DstNetflowMonitorName = v.DstNetflowMonitorName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.DstNetflowMonitorName = ""
	}

	if !v.TemplateName.IsNull() && !v.TemplateName.IsUnknown() {
		data.ConfigData.TemplateInputs.TemplateName = v.TemplateName.ValueString()
	} else {
		data.ConfigData.TemplateInputs.TemplateName = ""
	}

	return data
}
