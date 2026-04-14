package resource_link

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TopLevelFields are the top-level link attributes common to all policy types.
// These are merged into every PolicyTemplateFields entry via init().
var TopLevelFields = map[string]bool{
	"srcSwitchId":      true,
	"dstSwitchId":      true,
	"srcFabricName":    true,
	"srcSwitchName":    true,
	"srcInterfaceName": true,
	"dstFabricName":    true,
	"dstSwitchName":    true,
	"dstInterfaceName": true,
	"srcClusterName":   true,
	"dstClusterName":   true,
	"linkType":         true,
	"policyType":       true,
}

// PolicyTemplateFields maps each policyType to the set of allowed JSON field names.
// Top-level link attributes (TopLevelFields) are automatically included in every
// policy type via init(). The per-policy entries below list only the templateInputs
// fields specific to each policyType.
//
// Derived from the "# Used in:" annotations in generator/defs/links.yaml.
var PolicyTemplateFields = map[string]map[string]bool{
	"ipv6LinkLocal": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"speed":                   true,
		"fec":                     true,
		"macsec":                  true,
	},
	"numbered": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"speed":                   true,
		"fec":                     true,
		"macsec":                  true,
		"srcIp":                   true,
		"srcIpv6":                 true,
		"dstIp":                   true,
		"dstIpv6":                 true,
		"dhcpRelayOnSrcInterface": true,
		"dhcpRelayOnDstInterface": true,
		"bfdEchoOnSrcInterface":   true,
		"bfdEchoOnDstInterface":   true,
	},
	"unnumbered": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"speed":                   true,
		"fec":                     true,
		"macsec":                  true,
		"dhcpRelayOnSrcInterface": true,
		"dhcpRelayOnDstInterface": true,
	},
	"vpcPeerKeepalive": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"srcIp":                   true,
		"srcIpv6":                 true,
		"dstIp":                   true,
		"dstIpv6":                 true,
		"interfaceVrf":            true,
	},
	"ipfmNumbered": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"speed":                   true,
		"srcIp":                   true,
		"dstIp":                   true,
		"interfaceVrf":            true,
	},
	"preprovision": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"mtu":                     true,
		"speed":                   true,
	},
	"iosXeNumbered": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"interfaceAdminState":     true,
		"mtu":                     true,
		"speed":                   true,
		"srcIp":                   true,
		"dstIp":                   true,
	},
	"ebgpVrfLite": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"linkMtu":                              true,
		"macsec":                               true,
		"srcIpAddressMask":                     true,
		"srcIpv6AddressMask":                   true,
		"dstIpAddress":                         true,
		"dstIpv6Address":                       true,
		"srcEbgpAsn":                           true,
		"dstEbgpAsn":                           true,
		"defaultVrfEbgpNeighborPassword":       true,
		"vrfNameNxPeerSwitch":                  true,
		"autoGenConfigPeer":                    true,
		"autoGenConfigDefaultVrf":              true,
		"autoGenConfigNxPeerDefaultVrf":        true,
		"templateConfigGenPeer":                true,
		"redistribEbgpRouteMapName":            true,
		"routingTag":                           true,
		"dciTracking":                          true,
		"inheritTtagFabricSetting":             true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
	},
	"ospfVrfLite": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"linkMtu":                              true,
		"macsec":                               true,
		"srcIpAddressMask":                     true,
		"srcIpv6AddressMask":                   true,
		"dstIpAddress":                         true,
		"dstIpv6Address":                       true,
		"autoGenConfigPeer":                    true,
		"templateConfigGenPeer":                true,
		"inheritTtagFabricSetting":             true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
	},
	"layer3DciVrfLite": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"linkMtu":                              true,
		"speed":                                true,
		"macsec":                               true,
		"srcIpAddressMask":                     true,
		"srcIpv6AddressMask":                   true,
		"dstIpAddressMask":                     true,
		"dstIpv6AddressMask":                   true,
		"srcVrfName":                           true,
		"dstVrfName":                           true,
		"inheritTtagFabricSetting":             true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
		"ipv4Pim":                              true,
		"ipv6Pim":                              true,
		"ipRedirects":                          true,
		"netflowOnSrcInterface":                true,
		"netflowOnDstInterface":                true,
		"srcNetflowMonitorName":                true,
		"dstNetflowMonitorName":                true,
	},
	"csrMultisiteUnderlay": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"srcIpAddressMask":        true,
		"dstIp":                   true,
		"enableEbgpPassword":      true,
	},
	"multisiteUnderlay": {
		"srcInterfaceDescription":        true,
		"srcInterfaceConfig":             true,
		"dstInterfaceDescription":        true,
		"dstInterfaceConfig":             true,
		"linkMtu":                        true,
		"speed":                          true,
		"srcIpAddressMask":               true,
		"srcIpv6AddressMask":             true,
		"dstIpAddress":                   true,
		"dstIpv6Address":                 true,
		"srcEbgpAsn":                     true,
		"dstEbgpAsn":                     true,
		"ebgpPassword":                   true,
		"ebgpAuthKeyEncryptionType":      true,
		"ebgpBfd":                        true,
		"ebgpLogNeighborChange":          true,
		"ebgpMaximumPaths":               true,
		"ebgpSendComboth":                true,
		"enableEbgpPassword":             true,
		"inheritEbgpPasswordMsdSettings": true,
		"routingTag":                     true,
		"dciTrackingEnableFlag":          true,
		"inheritTtagFabricSetting":       true,
	},
	"vxlanAciMultisiteUnderlay": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"srcIpAddressMask":        true,
		"srcIpv6AddressMask":      true,
		"dstIpAddress":            true,
		"dstIpv6Address":          true,
		"srcEbgpAsn":              true,
		"dstEbgpAsn":              true,
	},
	"csrEvpnMultisiteOverlay": {
		"ebgpMultihop":       true,
		"enableEbgpPassword": true,
	},
	"multisiteOverlay": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"macsec":                               true,
		"srcIpAddress":                         true,
		"dstIpAddress":                         true,
		"srcEbgpAsn":                           true,
		"dstEbgpAsn":                           true,
		"ebgpPassword":                         true,
		"ebgpAuthKeyEncryptionType":            true,
		"ebgpMultihop":                         true,
		"enableEbgpPassword":                   true,
		"inheritEbgpPasswordMsdSettings":       true,
		"redistributeRouteServer":              true,
		"routeServerRoutingTag":                true,
		"ipv4Trm":                              true,
		"ipv6Trm":                              true,
		"skipConfigGeneration":                 true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
	},
	"vxlanAciEvpnOverlay": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"macsec":                               true,
		"srcIpAddress":                         true,
		"dstIpAddress":                         true,
		"srcEbgpAsn":                           true,
		"dstEbgpAsn":                           true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
	},
	"layer2Dci": {
		"srcInterfaceDescription":              true,
		"srcInterfaceConfig":                   true,
		"dstInterfaceDescription":              true,
		"dstInterfaceConfig":                   true,
		"speed":                                true,
		"macsec":                               true,
		"overrideFabricMacsec":                 true,
		"macsecCipherSuite":                    true,
		"macsecPrimaryKeyString":               true,
		"macsecPrimaryCryptographicAlgorithm":  true,
		"macsecFallbackKeyString":              true,
		"macsecFallbackCryptographicAlgorithm": true,
		"qkd":                                  true,
		"ignoreCertificate":                    true,
		"srcKmeServerIp":                       true,
		"srcMacsecKeyChainPrefix":              true,
		"srcQkdProfileName":                    true,
		"srcTrustpointLabel":                   true,
		"dstKmeServerIp":                       true,
		"dstMacsecKeyChainPrefix":              true,
		"dstQkdProfileName":                    true,
		"dstTrustpointLabel":                   true,
		"mtuType":                              true,
		"bpduGuard":                            true,
		"portTypeFast":                         true,
		"nativeVlan":                           true,
		"trunkAllowedVlans":                    true,
	},
	"mplsUnderlay": {
		"srcInterfaceDescription":  true,
		"srcInterfaceConfig":       true,
		"dstInterfaceDescription":  true,
		"dstInterfaceConfig":       true,
		"linkMtu":                  true,
		"srcIpAddressMask":         true,
		"dstIpAddress":             true,
		"inheritTtagFabricSetting": true,
	},
	"mplsOverlay": {
		"dstIpAddress": true,
		"srcEbgpAsn":   true,
		"dstEbgpAsn":   true,
	},
	"routedFabric": {
		"srcInterfaceDescription": true,
		"srcInterfaceConfig":      true,
		"dstInterfaceDescription": true,
		"dstInterfaceConfig":      true,
		"linkMtu":                 true,
		"srcIpAddressMask":        true,
		"srcIpv6AddressMask":      true,
		"dstIpAddress":            true,
		"dstIpv6Address":          true,
		"srcEbgpAsn":              true,
		"dstEbgpAsn":              true,
		"ebgpMaximumPaths":        true,
	},
	"userDefined": {
		"templateName": true,
	},
	"csrBgpOverIpsec": {},
}

func init() {
	for policy := range PolicyTemplateFields {
		for field := range TopLevelFields {
			PolicyTemplateFields[policy][field] = true
		}
	}
}

// FilterByPolicyType zeroes out all templateInputs fields that are not applicable
// to the given policyType. If policyType is empty or not found in the mapping,
// no filtering is performed (safe fallback for unknown/future policy types).
func (t *NDFCConfigDataTemplateInputsValue) FilterByPolicyType(policyType string) {
	allowed, ok := PolicyTemplateFields[policyType]
	if !ok || policyType == "" {
		return
	}

	if !allowed["srcInterfaceDescription"] {
		t.SrcInterfaceDescription = ""
	}
	if !allowed["srcInterfaceConfig"] {
		t.SrcInterfaceConfig = ""
	}
	if !allowed["dstInterfaceDescription"] {
		t.DstInterfaceDescription = ""
	}
	if !allowed["dstInterfaceConfig"] {
		t.DstInterfaceConfig = ""
	}
	if !allowed["interfaceAdminState"] {
		t.InterfaceAdminState = nil
	}
	if !allowed["mtu"] {
		t.Mtu = nil
	}
	if !allowed["linkMtu"] {
		t.LinkMtu = nil
	}
	if !allowed["speed"] {
		t.Speed = ""
	}
	if !allowed["fec"] {
		t.Fec = ""
	}
	if !allowed["macsec"] {
		t.Macsec = nil
	}
	if !allowed["srcIp"] {
		t.SrcIp = ""
	}
	if !allowed["srcIpv6"] {
		t.SrcIpv6 = ""
	}
	if !allowed["srcIpAddress"] {
		t.SrcIpAddress = ""
	}
	if !allowed["srcIpAddressMask"] {
		t.SrcIpAddressMask = ""
	}
	if !allowed["srcIpv6AddressMask"] {
		t.SrcIpv6AddressMask = ""
	}
	if !allowed["dstIp"] {
		t.DstIp = ""
	}
	if !allowed["dstIpv6"] {
		t.DstIpv6 = ""
	}
	if !allowed["dstIpAddress"] {
		t.DstIpAddress = ""
	}
	if !allowed["dstIpAddressMask"] {
		t.DstIpAddressMask = ""
	}
	if !allowed["dstIpv6Address"] {
		t.DstIpv6Address = ""
	}
	if !allowed["dstIpv6AddressMask"] {
		t.DstIpv6AddressMask = ""
	}
	if !allowed["srcEbgpAsn"] {
		t.SrcEbgpAsn = ""
	}
	if !allowed["dstEbgpAsn"] {
		t.DstEbgpAsn = ""
	}
	if !allowed["ebgpPassword"] {
		t.EbgpPassword = ""
	}
	if !allowed["ebgpAuthKeyEncryptionType"] {
		t.EbgpAuthKeyEncryptionType = ""
	}
	if !allowed["ebgpBfd"] {
		t.EbgpBfd = nil
	}
	if !allowed["ebgpLogNeighborChange"] {
		t.EbgpLogNeighborChange = nil
	}
	if !allowed["ebgpMaximumPaths"] {
		t.EbgpMaximumPaths = nil
	}
	if !allowed["ebgpMultihop"] {
		t.EbgpMultihop = nil
	}
	if !allowed["ebgpSendComboth"] {
		t.EbgpSendComboth = nil
	}
	if !allowed["enableEbgpPassword"] {
		t.EnableEbgpPassword = nil
	}
	if !allowed["inheritEbgpPasswordMsdSettings"] {
		t.InheritEbgpPasswordMsdSettings = nil
	}
	if !allowed["defaultVrfEbgpNeighborPassword"] {
		t.DefaultVrfEbgpNeighborPassword = ""
	}
	if !allowed["interfaceVrf"] {
		t.InterfaceVrf = ""
	}
	if !allowed["srcVrfName"] {
		t.SrcVrfName = ""
	}
	if !allowed["dstVrfName"] {
		t.DstVrfName = ""
	}
	if !allowed["vrfNameNxPeerSwitch"] {
		t.VrfNameNxPeerSwitch = ""
	}
	if !allowed["autoGenConfigPeer"] {
		t.AutoGenConfigPeer = nil
	}
	if !allowed["autoGenConfigDefaultVrf"] {
		t.AutoGenConfigDefaultVrf = nil
	}
	if !allowed["autoGenConfigNxPeerDefaultVrf"] {
		t.AutoGenConfigNxPeerDefaultVrf = nil
	}
	if !allowed["templateConfigGenPeer"] {
		t.TemplateConfigGenPeer = ""
	}
	if !allowed["redistribEbgpRouteMapName"] {
		t.RedistribEbgpRouteMapName = ""
	}
	if !allowed["redistributeRouteServer"] {
		t.RedistributeRouteServer = nil
	}
	if !allowed["routeServerRoutingTag"] {
		t.RouteServerRoutingTag = ""
	}
	if !allowed["routingTag"] {
		t.RoutingTag = ""
	}
	if !allowed["dciTracking"] {
		t.DciTracking = nil
	}
	if !allowed["dciTrackingEnableFlag"] {
		t.DciTrackingEnableFlag = nil
	}
	if !allowed["inheritTtagFabricSetting"] {
		t.InheritTtagFabricSetting = nil
	}
	if !allowed["overrideFabricMacsec"] {
		t.OverrideFabricMacsec = nil
	}
	if !allowed["macsecCipherSuite"] {
		t.MacsecCipherSuite = ""
	}
	if !allowed["macsecPrimaryKeyString"] {
		t.MacsecPrimaryKeyString = ""
	}
	if !allowed["macsecPrimaryCryptographicAlgorithm"] {
		t.MacsecPrimaryCryptographicAlgorithm = ""
	}
	if !allowed["macsecFallbackKeyString"] {
		t.MacsecFallbackKeyString = ""
	}
	if !allowed["macsecFallbackCryptographicAlgorithm"] {
		t.MacsecFallbackCryptographicAlgorithm = ""
	}
	if !allowed["qkd"] {
		t.Qkd = nil
	}
	if !allowed["ignoreCertificate"] {
		t.IgnoreCertificate = nil
	}
	if !allowed["srcKmeServerIp"] {
		t.SrcKmeServerIp = ""
	}
	if !allowed["srcMacsecKeyChainPrefix"] {
		t.SrcMacsecKeyChainPrefix = ""
	}
	if !allowed["srcQkdProfileName"] {
		t.SrcQkdProfileName = ""
	}
	if !allowed["srcTrustpointLabel"] {
		t.SrcTrustpointLabel = ""
	}
	if !allowed["dstKmeServerIp"] {
		t.DstKmeServerIp = ""
	}
	if !allowed["dstMacsecKeyChainPrefix"] {
		t.DstMacsecKeyChainPrefix = ""
	}
	if !allowed["dstQkdProfileName"] {
		t.DstQkdProfileName = ""
	}
	if !allowed["dstTrustpointLabel"] {
		t.DstTrustpointLabel = ""
	}
	if !allowed["dhcpRelayOnSrcInterface"] {
		t.DhcpRelayOnSrcInterface = nil
	}
	if !allowed["dhcpRelayOnDstInterface"] {
		t.DhcpRelayOnDstInterface = nil
	}
	if !allowed["bfdEchoOnSrcInterface"] {
		t.BfdEchoOnSrcInterface = nil
	}
	if !allowed["bfdEchoOnDstInterface"] {
		t.BfdEchoOnDstInterface = nil
	}
	if !allowed["mtuType"] {
		t.MtuType = ""
	}
	if !allowed["ipv4Pim"] {
		t.Ipv4Pim = nil
	}
	if !allowed["ipv4Trm"] {
		t.Ipv4Trm = nil
	}
	if !allowed["ipv6Pim"] {
		t.Ipv6Pim = nil
	}
	if !allowed["ipv6Trm"] {
		t.Ipv6Trm = nil
	}
	if !allowed["ipRedirects"] {
		t.IpRedirects = nil
	}
	if !allowed["skipConfigGeneration"] {
		t.SkipConfigGeneration = nil
	}
	if !allowed["bpduGuard"] {
		t.BpduGuard = ""
	}
	if !allowed["portTypeFast"] {
		t.PortTypeFast = nil
	}
	if !allowed["nativeVlan"] {
		t.NativeVlan = nil
	}
	if !allowed["trunkAllowedVlans"] {
		t.TrunkAllowedVlans = ""
	}
	if !allowed["netflowOnSrcInterface"] {
		t.NetflowOnSrcInterface = nil
	}
	if !allowed["netflowOnDstInterface"] {
		t.NetflowOnDstInterface = nil
	}
	if !allowed["srcNetflowMonitorName"] {
		t.SrcNetflowMonitorName = ""
	}
	if !allowed["dstNetflowMonitorName"] {
		t.DstNetflowMonitorName = ""
	}
	if !allowed["templateName"] {
		t.TemplateName = ""
	}
}

// PreserveNonPolicyDefaults restores LinkModel fields that are not applicable to
// the given policyType from a previously-saved snapshot. This prevents Terraform
// "inconsistent result after apply" errors caused by SetModelData setting
// non-applicable (or empty) fields to null when the plan had schema defaults.
//
// For non-applicable templateInputs fields: always restore from saved.
// For applicable fields: restore only when the new value is null but saved was not
// (handles API returning zero values that SetModelData maps to null).
func (v *LinkModel) PreserveNonPolicyDefaults(policyType string, saved *LinkModel) {
	allowed, ok := PolicyTemplateFields[policyType]
	if !ok || policyType == "" {
		return
	}

	// --- templateInputs fields ---
	// SrcInterfaceDescription
	if !allowed["srcInterfaceDescription"] {
		v.SrcInterfaceDescription = saved.SrcInterfaceDescription
	} else if v.SrcInterfaceDescription.IsNull() && !saved.SrcInterfaceDescription.IsNull() {
		v.SrcInterfaceDescription = saved.SrcInterfaceDescription
	}
	// SrcInterfaceConfig
	if !allowed["srcInterfaceConfig"] {
		v.SrcInterfaceConfig = saved.SrcInterfaceConfig
	} else if v.SrcInterfaceConfig.IsNull() && !saved.SrcInterfaceConfig.IsNull() {
		v.SrcInterfaceConfig = saved.SrcInterfaceConfig
	}
	// DstInterfaceDescription
	if !allowed["dstInterfaceDescription"] {
		v.DstInterfaceDescription = saved.DstInterfaceDescription
	} else if v.DstInterfaceDescription.IsNull() && !saved.DstInterfaceDescription.IsNull() {
		v.DstInterfaceDescription = saved.DstInterfaceDescription
	}
	// DstInterfaceConfig
	if !allowed["dstInterfaceConfig"] {
		v.DstInterfaceConfig = saved.DstInterfaceConfig
	} else if v.DstInterfaceConfig.IsNull() && !saved.DstInterfaceConfig.IsNull() {
		v.DstInterfaceConfig = saved.DstInterfaceConfig
	}
	// InterfaceAdminState
	if !allowed["interfaceAdminState"] {
		v.InterfaceAdminState = saved.InterfaceAdminState
	} else if v.InterfaceAdminState.IsNull() && !saved.InterfaceAdminState.IsNull() {
		v.InterfaceAdminState = saved.InterfaceAdminState
	}
	// Mtu
	if !allowed["mtu"] {
		v.Mtu = saved.Mtu
	} else if v.Mtu.IsNull() && !saved.Mtu.IsNull() {
		v.Mtu = saved.Mtu
	}
	// LinkMtu
	if !allowed["linkMtu"] {
		v.LinkMtu = saved.LinkMtu
	} else if v.LinkMtu.IsNull() && !saved.LinkMtu.IsNull() {
		v.LinkMtu = saved.LinkMtu
	}
	// Speed
	if !allowed["speed"] {
		v.Speed = saved.Speed
	} else if v.Speed.IsNull() && !saved.Speed.IsNull() {
		v.Speed = saved.Speed
	}
	// Fec
	if !allowed["fec"] {
		v.Fec = saved.Fec
	} else if v.Fec.IsNull() && !saved.Fec.IsNull() {
		v.Fec = saved.Fec
	}
	// Macsec
	if !allowed["macsec"] {
		v.Macsec = saved.Macsec
	} else if v.Macsec.IsNull() && !saved.Macsec.IsNull() {
		v.Macsec = saved.Macsec
	}
	// SrcIp
	if !allowed["srcIp"] {
		v.SrcIp = saved.SrcIp
	} else if v.SrcIp.IsNull() && !saved.SrcIp.IsNull() {
		v.SrcIp = saved.SrcIp
	}
	// SrcIpv6
	if !allowed["srcIpv6"] {
		v.SrcIpv6 = saved.SrcIpv6
	} else if v.SrcIpv6.IsNull() && !saved.SrcIpv6.IsNull() {
		v.SrcIpv6 = saved.SrcIpv6
	}
	// SrcIpAddress
	if !allowed["srcIpAddress"] {
		v.SrcIpAddress = saved.SrcIpAddress
	} else if v.SrcIpAddress.IsNull() && !saved.SrcIpAddress.IsNull() {
		v.SrcIpAddress = saved.SrcIpAddress
	}
	// SrcIpAddressMask
	if !allowed["srcIpAddressMask"] {
		v.SrcIpAddressMask = saved.SrcIpAddressMask
	} else if v.SrcIpAddressMask.IsNull() && !saved.SrcIpAddressMask.IsNull() {
		v.SrcIpAddressMask = saved.SrcIpAddressMask
	}
	// SrcIpv6AddressMask
	if !allowed["srcIpv6AddressMask"] {
		v.SrcIpv6AddressMask = saved.SrcIpv6AddressMask
	} else if v.SrcIpv6AddressMask.IsNull() && !saved.SrcIpv6AddressMask.IsNull() {
		v.SrcIpv6AddressMask = saved.SrcIpv6AddressMask
	}
	// DstIp
	if !allowed["dstIp"] {
		v.DstIp = saved.DstIp
	} else if v.DstIp.IsNull() && !saved.DstIp.IsNull() {
		v.DstIp = saved.DstIp
	}
	// DstIpv6
	if !allowed["dstIpv6"] {
		v.DstIpv6 = saved.DstIpv6
	} else if v.DstIpv6.IsNull() && !saved.DstIpv6.IsNull() {
		v.DstIpv6 = saved.DstIpv6
	}
	// DstIpAddress
	if !allowed["dstIpAddress"] {
		v.DstIpAddress = saved.DstIpAddress
	} else if v.DstIpAddress.IsNull() && !saved.DstIpAddress.IsNull() {
		v.DstIpAddress = saved.DstIpAddress
	}
	// DstIpAddressMask
	if !allowed["dstIpAddressMask"] {
		v.DstIpAddressMask = saved.DstIpAddressMask
	} else if v.DstIpAddressMask.IsNull() && !saved.DstIpAddressMask.IsNull() {
		v.DstIpAddressMask = saved.DstIpAddressMask
	}
	// DstIpv6Address
	if !allowed["dstIpv6Address"] {
		v.DstIpv6Address = saved.DstIpv6Address
	} else if v.DstIpv6Address.IsNull() && !saved.DstIpv6Address.IsNull() {
		v.DstIpv6Address = saved.DstIpv6Address
	}
	// DstIpv6AddressMask
	if !allowed["dstIpv6AddressMask"] {
		v.DstIpv6AddressMask = saved.DstIpv6AddressMask
	} else if v.DstIpv6AddressMask.IsNull() && !saved.DstIpv6AddressMask.IsNull() {
		v.DstIpv6AddressMask = saved.DstIpv6AddressMask
	}
	// SrcEbgpAsn
	if !allowed["srcEbgpAsn"] {
		v.SrcEbgpAsn = saved.SrcEbgpAsn
	} else if v.SrcEbgpAsn.IsNull() && !saved.SrcEbgpAsn.IsNull() {
		v.SrcEbgpAsn = saved.SrcEbgpAsn
	}
	// DstEbgpAsn
	if !allowed["dstEbgpAsn"] {
		v.DstEbgpAsn = saved.DstEbgpAsn
	} else if v.DstEbgpAsn.IsNull() && !saved.DstEbgpAsn.IsNull() {
		v.DstEbgpAsn = saved.DstEbgpAsn
	}
	// EbgpPassword
	if !allowed["ebgpPassword"] {
		v.EbgpPassword = saved.EbgpPassword
	} else if v.EbgpPassword.IsNull() && !saved.EbgpPassword.IsNull() {
		v.EbgpPassword = saved.EbgpPassword
	}
	// EbgpAuthKeyEncryptionType
	if !allowed["ebgpAuthKeyEncryptionType"] {
		v.EbgpAuthKeyEncryptionType = saved.EbgpAuthKeyEncryptionType
	} else if v.EbgpAuthKeyEncryptionType.IsNull() && !saved.EbgpAuthKeyEncryptionType.IsNull() {
		v.EbgpAuthKeyEncryptionType = saved.EbgpAuthKeyEncryptionType
	}
	// EbgpBfd
	if !allowed["ebgpBfd"] {
		v.EbgpBfd = saved.EbgpBfd
	} else if v.EbgpBfd.IsNull() && !saved.EbgpBfd.IsNull() {
		v.EbgpBfd = saved.EbgpBfd
	}
	// EbgpLogNeighborChange
	if !allowed["ebgpLogNeighborChange"] {
		v.EbgpLogNeighborChange = saved.EbgpLogNeighborChange
	} else if v.EbgpLogNeighborChange.IsNull() && !saved.EbgpLogNeighborChange.IsNull() {
		v.EbgpLogNeighborChange = saved.EbgpLogNeighborChange
	}
	// EbgpMaximumPaths
	if !allowed["ebgpMaximumPaths"] {
		v.EbgpMaximumPaths = saved.EbgpMaximumPaths
	} else if v.EbgpMaximumPaths.IsNull() && !saved.EbgpMaximumPaths.IsNull() {
		v.EbgpMaximumPaths = saved.EbgpMaximumPaths
	}
	// EbgpMultihop
	if !allowed["ebgpMultihop"] {
		v.EbgpMultihop = saved.EbgpMultihop
	} else if v.EbgpMultihop.IsNull() && !saved.EbgpMultihop.IsNull() {
		v.EbgpMultihop = saved.EbgpMultihop
	}
	// EbgpSendComboth
	if !allowed["ebgpSendComboth"] {
		v.EbgpSendComboth = saved.EbgpSendComboth
	} else if v.EbgpSendComboth.IsNull() && !saved.EbgpSendComboth.IsNull() {
		v.EbgpSendComboth = saved.EbgpSendComboth
	}
	// EnableEbgpPassword
	if !allowed["enableEbgpPassword"] {
		v.EnableEbgpPassword = saved.EnableEbgpPassword
	} else if v.EnableEbgpPassword.IsNull() && !saved.EnableEbgpPassword.IsNull() {
		v.EnableEbgpPassword = saved.EnableEbgpPassword
	}
	// InheritEbgpPasswordMsdSettings
	if !allowed["inheritEbgpPasswordMsdSettings"] {
		v.InheritEbgpPasswordMsdSettings = saved.InheritEbgpPasswordMsdSettings
	} else if v.InheritEbgpPasswordMsdSettings.IsNull() && !saved.InheritEbgpPasswordMsdSettings.IsNull() {
		v.InheritEbgpPasswordMsdSettings = saved.InheritEbgpPasswordMsdSettings
	}
	// DefaultVrfEbgpNeighborPassword
	if !allowed["defaultVrfEbgpNeighborPassword"] {
		v.DefaultVrfEbgpNeighborPassword = saved.DefaultVrfEbgpNeighborPassword
	} else if v.DefaultVrfEbgpNeighborPassword.IsNull() && !saved.DefaultVrfEbgpNeighborPassword.IsNull() {
		v.DefaultVrfEbgpNeighborPassword = saved.DefaultVrfEbgpNeighborPassword
	}
	// InterfaceVrf
	if !allowed["interfaceVrf"] {
		v.InterfaceVrf = saved.InterfaceVrf
	} else if v.InterfaceVrf.IsNull() && !saved.InterfaceVrf.IsNull() {
		v.InterfaceVrf = saved.InterfaceVrf
	}
	// SrcVrfName
	if !allowed["srcVrfName"] {
		v.SrcVrfName = saved.SrcVrfName
	} else if v.SrcVrfName.IsNull() && !saved.SrcVrfName.IsNull() {
		v.SrcVrfName = saved.SrcVrfName
	}
	// DstVrfName
	if !allowed["dstVrfName"] {
		v.DstVrfName = saved.DstVrfName
	} else if v.DstVrfName.IsNull() && !saved.DstVrfName.IsNull() {
		v.DstVrfName = saved.DstVrfName
	}
	// VrfNameNxPeerSwitch
	if !allowed["vrfNameNxPeerSwitch"] {
		v.VrfNameNxPeerSwitch = saved.VrfNameNxPeerSwitch
	} else if v.VrfNameNxPeerSwitch.IsNull() && !saved.VrfNameNxPeerSwitch.IsNull() {
		v.VrfNameNxPeerSwitch = saved.VrfNameNxPeerSwitch
	}
	// AutoGenConfigPeer
	if !allowed["autoGenConfigPeer"] {
		v.AutoGenConfigPeer = saved.AutoGenConfigPeer
	} else if v.AutoGenConfigPeer.IsNull() && !saved.AutoGenConfigPeer.IsNull() {
		v.AutoGenConfigPeer = saved.AutoGenConfigPeer
	}
	// AutoGenConfigDefaultVrf
	if !allowed["autoGenConfigDefaultVrf"] {
		v.AutoGenConfigDefaultVrf = saved.AutoGenConfigDefaultVrf
	} else if v.AutoGenConfigDefaultVrf.IsNull() && !saved.AutoGenConfigDefaultVrf.IsNull() {
		v.AutoGenConfigDefaultVrf = saved.AutoGenConfigDefaultVrf
	}
	// AutoGenConfigNxPeerDefaultVrf
	if !allowed["autoGenConfigNxPeerDefaultVrf"] {
		v.AutoGenConfigNxPeerDefaultVrf = saved.AutoGenConfigNxPeerDefaultVrf
	} else if v.AutoGenConfigNxPeerDefaultVrf.IsNull() && !saved.AutoGenConfigNxPeerDefaultVrf.IsNull() {
		v.AutoGenConfigNxPeerDefaultVrf = saved.AutoGenConfigNxPeerDefaultVrf
	}
	// TemplateConfigGenPeer
	if !allowed["templateConfigGenPeer"] {
		v.TemplateConfigGenPeer = saved.TemplateConfigGenPeer
	} else if v.TemplateConfigGenPeer.IsNull() && !saved.TemplateConfigGenPeer.IsNull() {
		v.TemplateConfigGenPeer = saved.TemplateConfigGenPeer
	}
	// RedistribEbgpRouteMapName
	if !allowed["redistribEbgpRouteMapName"] {
		v.RedistribEbgpRouteMapName = saved.RedistribEbgpRouteMapName
	} else if v.RedistribEbgpRouteMapName.IsNull() && !saved.RedistribEbgpRouteMapName.IsNull() {
		v.RedistribEbgpRouteMapName = saved.RedistribEbgpRouteMapName
	}
	// RedistributeRouteServer
	if !allowed["redistributeRouteServer"] {
		v.RedistributeRouteServer = saved.RedistributeRouteServer
	} else if v.RedistributeRouteServer.IsNull() && !saved.RedistributeRouteServer.IsNull() {
		v.RedistributeRouteServer = saved.RedistributeRouteServer
	}
	// RouteServerRoutingTag
	if !allowed["routeServerRoutingTag"] {
		v.RouteServerRoutingTag = saved.RouteServerRoutingTag
	} else if v.RouteServerRoutingTag.IsNull() && !saved.RouteServerRoutingTag.IsNull() {
		v.RouteServerRoutingTag = saved.RouteServerRoutingTag
	}
	// RoutingTag
	if !allowed["routingTag"] {
		v.RoutingTag = saved.RoutingTag
	} else if v.RoutingTag.IsNull() && !saved.RoutingTag.IsNull() {
		v.RoutingTag = saved.RoutingTag
	}
	// DciTracking
	if !allowed["dciTracking"] {
		v.DciTracking = saved.DciTracking
	} else if v.DciTracking.IsNull() && !saved.DciTracking.IsNull() {
		v.DciTracking = saved.DciTracking
	}
	// DciTrackingEnableFlag
	if !allowed["dciTrackingEnableFlag"] {
		v.DciTrackingEnableFlag = saved.DciTrackingEnableFlag
	} else if v.DciTrackingEnableFlag.IsNull() && !saved.DciTrackingEnableFlag.IsNull() {
		v.DciTrackingEnableFlag = saved.DciTrackingEnableFlag
	}
	// InheritTtagFabricSetting
	if !allowed["inheritTtagFabricSetting"] {
		v.InheritTtagFabricSetting = saved.InheritTtagFabricSetting
	} else if v.InheritTtagFabricSetting.IsNull() && !saved.InheritTtagFabricSetting.IsNull() {
		v.InheritTtagFabricSetting = saved.InheritTtagFabricSetting
	}
	// OverrideFabricMacsec
	if !allowed["overrideFabricMacsec"] {
		v.OverrideFabricMacsec = saved.OverrideFabricMacsec
	} else if v.OverrideFabricMacsec.IsNull() && !saved.OverrideFabricMacsec.IsNull() {
		v.OverrideFabricMacsec = saved.OverrideFabricMacsec
	}
	// MacsecCipherSuite
	if !allowed["macsecCipherSuite"] {
		v.MacsecCipherSuite = saved.MacsecCipherSuite
	} else if v.MacsecCipherSuite.IsNull() && !saved.MacsecCipherSuite.IsNull() {
		v.MacsecCipherSuite = saved.MacsecCipherSuite
	}
	// MacsecPrimaryKeyString
	if !allowed["macsecPrimaryKeyString"] {
		v.MacsecPrimaryKeyString = saved.MacsecPrimaryKeyString
	} else if v.MacsecPrimaryKeyString.IsNull() && !saved.MacsecPrimaryKeyString.IsNull() {
		v.MacsecPrimaryKeyString = saved.MacsecPrimaryKeyString
	}
	// MacsecPrimaryCryptographicAlgorithm
	if !allowed["macsecPrimaryCryptographicAlgorithm"] {
		v.MacsecPrimaryCryptographicAlgorithm = saved.MacsecPrimaryCryptographicAlgorithm
	} else if v.MacsecPrimaryCryptographicAlgorithm.IsNull() && !saved.MacsecPrimaryCryptographicAlgorithm.IsNull() {
		v.MacsecPrimaryCryptographicAlgorithm = saved.MacsecPrimaryCryptographicAlgorithm
	}
	// MacsecFallbackKeyString
	if !allowed["macsecFallbackKeyString"] {
		v.MacsecFallbackKeyString = saved.MacsecFallbackKeyString
	} else if v.MacsecFallbackKeyString.IsNull() && !saved.MacsecFallbackKeyString.IsNull() {
		v.MacsecFallbackKeyString = saved.MacsecFallbackKeyString
	}
	// MacsecFallbackCryptographicAlgorithm
	if !allowed["macsecFallbackCryptographicAlgorithm"] {
		v.MacsecFallbackCryptographicAlgorithm = saved.MacsecFallbackCryptographicAlgorithm
	} else if v.MacsecFallbackCryptographicAlgorithm.IsNull() && !saved.MacsecFallbackCryptographicAlgorithm.IsNull() {
		v.MacsecFallbackCryptographicAlgorithm = saved.MacsecFallbackCryptographicAlgorithm
	}
	// Qkd
	if !allowed["qkd"] {
		v.Qkd = saved.Qkd
	} else if v.Qkd.IsNull() && !saved.Qkd.IsNull() {
		v.Qkd = saved.Qkd
	}
	// IgnoreCertificate
	if !allowed["ignoreCertificate"] {
		v.IgnoreCertificate = saved.IgnoreCertificate
	} else if v.IgnoreCertificate.IsNull() && !saved.IgnoreCertificate.IsNull() {
		v.IgnoreCertificate = saved.IgnoreCertificate
	}
	// SrcKmeServerIp
	if !allowed["srcKmeServerIp"] {
		v.SrcKmeServerIp = saved.SrcKmeServerIp
	} else if v.SrcKmeServerIp.IsNull() && !saved.SrcKmeServerIp.IsNull() {
		v.SrcKmeServerIp = saved.SrcKmeServerIp
	}
	// SrcMacsecKeyChainPrefix
	if !allowed["srcMacsecKeyChainPrefix"] {
		v.SrcMacsecKeyChainPrefix = saved.SrcMacsecKeyChainPrefix
	} else if v.SrcMacsecKeyChainPrefix.IsNull() && !saved.SrcMacsecKeyChainPrefix.IsNull() {
		v.SrcMacsecKeyChainPrefix = saved.SrcMacsecKeyChainPrefix
	}
	// SrcQkdProfileName
	if !allowed["srcQkdProfileName"] {
		v.SrcQkdProfileName = saved.SrcQkdProfileName
	} else if v.SrcQkdProfileName.IsNull() && !saved.SrcQkdProfileName.IsNull() {
		v.SrcQkdProfileName = saved.SrcQkdProfileName
	}
	// SrcTrustpointLabel
	if !allowed["srcTrustpointLabel"] {
		v.SrcTrustpointLabel = saved.SrcTrustpointLabel
	} else if v.SrcTrustpointLabel.IsNull() && !saved.SrcTrustpointLabel.IsNull() {
		v.SrcTrustpointLabel = saved.SrcTrustpointLabel
	}
	// DstKmeServerIp
	if !allowed["dstKmeServerIp"] {
		v.DstKmeServerIp = saved.DstKmeServerIp
	} else if v.DstKmeServerIp.IsNull() && !saved.DstKmeServerIp.IsNull() {
		v.DstKmeServerIp = saved.DstKmeServerIp
	}
	// DstMacsecKeyChainPrefix
	if !allowed["dstMacsecKeyChainPrefix"] {
		v.DstMacsecKeyChainPrefix = saved.DstMacsecKeyChainPrefix
	} else if v.DstMacsecKeyChainPrefix.IsNull() && !saved.DstMacsecKeyChainPrefix.IsNull() {
		v.DstMacsecKeyChainPrefix = saved.DstMacsecKeyChainPrefix
	}
	// DstQkdProfileName
	if !allowed["dstQkdProfileName"] {
		v.DstQkdProfileName = saved.DstQkdProfileName
	} else if v.DstQkdProfileName.IsNull() && !saved.DstQkdProfileName.IsNull() {
		v.DstQkdProfileName = saved.DstQkdProfileName
	}
	// DstTrustpointLabel
	if !allowed["dstTrustpointLabel"] {
		v.DstTrustpointLabel = saved.DstTrustpointLabel
	} else if v.DstTrustpointLabel.IsNull() && !saved.DstTrustpointLabel.IsNull() {
		v.DstTrustpointLabel = saved.DstTrustpointLabel
	}
	// DhcpRelayOnSrcInterface
	if !allowed["dhcpRelayOnSrcInterface"] {
		v.DhcpRelayOnSrcInterface = saved.DhcpRelayOnSrcInterface
	} else if v.DhcpRelayOnSrcInterface.IsNull() && !saved.DhcpRelayOnSrcInterface.IsNull() {
		v.DhcpRelayOnSrcInterface = saved.DhcpRelayOnSrcInterface
	}
	// DhcpRelayOnDstInterface
	if !allowed["dhcpRelayOnDstInterface"] {
		v.DhcpRelayOnDstInterface = saved.DhcpRelayOnDstInterface
	} else if v.DhcpRelayOnDstInterface.IsNull() && !saved.DhcpRelayOnDstInterface.IsNull() {
		v.DhcpRelayOnDstInterface = saved.DhcpRelayOnDstInterface
	}
	// BfdEchoOnSrcInterface
	if !allowed["bfdEchoOnSrcInterface"] {
		v.BfdEchoOnSrcInterface = saved.BfdEchoOnSrcInterface
	} else if v.BfdEchoOnSrcInterface.IsNull() && !saved.BfdEchoOnSrcInterface.IsNull() {
		v.BfdEchoOnSrcInterface = saved.BfdEchoOnSrcInterface
	}
	// BfdEchoOnDstInterface
	if !allowed["bfdEchoOnDstInterface"] {
		v.BfdEchoOnDstInterface = saved.BfdEchoOnDstInterface
	} else if v.BfdEchoOnDstInterface.IsNull() && !saved.BfdEchoOnDstInterface.IsNull() {
		v.BfdEchoOnDstInterface = saved.BfdEchoOnDstInterface
	}
	// MtuType
	if !allowed["mtuType"] {
		v.MtuType = saved.MtuType
	} else if v.MtuType.IsNull() && !saved.MtuType.IsNull() {
		v.MtuType = saved.MtuType
	}
	// Ipv4Pim
	if !allowed["ipv4Pim"] {
		v.Ipv4Pim = saved.Ipv4Pim
	} else if v.Ipv4Pim.IsNull() && !saved.Ipv4Pim.IsNull() {
		v.Ipv4Pim = saved.Ipv4Pim
	}
	// Ipv4Trm
	if !allowed["ipv4Trm"] {
		v.Ipv4Trm = saved.Ipv4Trm
	} else if v.Ipv4Trm.IsNull() && !saved.Ipv4Trm.IsNull() {
		v.Ipv4Trm = saved.Ipv4Trm
	}
	// Ipv6Pim
	if !allowed["ipv6Pim"] {
		v.Ipv6Pim = saved.Ipv6Pim
	} else if v.Ipv6Pim.IsNull() && !saved.Ipv6Pim.IsNull() {
		v.Ipv6Pim = saved.Ipv6Pim
	}
	// Ipv6Trm
	if !allowed["ipv6Trm"] {
		v.Ipv6Trm = saved.Ipv6Trm
	} else if v.Ipv6Trm.IsNull() && !saved.Ipv6Trm.IsNull() {
		v.Ipv6Trm = saved.Ipv6Trm
	}
	// IpRedirects
	if !allowed["ipRedirects"] {
		v.IpRedirects = saved.IpRedirects
	} else if v.IpRedirects.IsNull() && !saved.IpRedirects.IsNull() {
		v.IpRedirects = saved.IpRedirects
	}
	// SkipConfigGeneration
	if !allowed["skipConfigGeneration"] {
		v.SkipConfigGeneration = saved.SkipConfigGeneration
	} else if v.SkipConfigGeneration.IsNull() && !saved.SkipConfigGeneration.IsNull() {
		v.SkipConfigGeneration = saved.SkipConfigGeneration
	}
	// BpduGuard
	if !allowed["bpduGuard"] {
		v.BpduGuard = saved.BpduGuard
	} else if v.BpduGuard.IsNull() && !saved.BpduGuard.IsNull() {
		v.BpduGuard = saved.BpduGuard
	}
	// PortTypeFast
	if !allowed["portTypeFast"] {
		v.PortTypeFast = saved.PortTypeFast
	} else if v.PortTypeFast.IsNull() && !saved.PortTypeFast.IsNull() {
		v.PortTypeFast = saved.PortTypeFast
	}
	// NativeVlan
	if !allowed["nativeVlan"] {
		v.NativeVlan = saved.NativeVlan
	} else if v.NativeVlan.IsNull() && !saved.NativeVlan.IsNull() {
		v.NativeVlan = saved.NativeVlan
	}
	// TrunkAllowedVlans
	if !allowed["trunkAllowedVlans"] {
		v.TrunkAllowedVlans = saved.TrunkAllowedVlans
	} else if v.TrunkAllowedVlans.IsNull() && !saved.TrunkAllowedVlans.IsNull() {
		v.TrunkAllowedVlans = saved.TrunkAllowedVlans
	}
	// NetflowOnSrcInterface
	if !allowed["netflowOnSrcInterface"] {
		v.NetflowOnSrcInterface = saved.NetflowOnSrcInterface
	} else if v.NetflowOnSrcInterface.IsNull() && !saved.NetflowOnSrcInterface.IsNull() {
		v.NetflowOnSrcInterface = saved.NetflowOnSrcInterface
	}
	// NetflowOnDstInterface
	if !allowed["netflowOnDstInterface"] {
		v.NetflowOnDstInterface = saved.NetflowOnDstInterface
	} else if v.NetflowOnDstInterface.IsNull() && !saved.NetflowOnDstInterface.IsNull() {
		v.NetflowOnDstInterface = saved.NetflowOnDstInterface
	}
	// SrcNetflowMonitorName
	if !allowed["srcNetflowMonitorName"] {
		v.SrcNetflowMonitorName = saved.SrcNetflowMonitorName
	} else if v.SrcNetflowMonitorName.IsNull() && !saved.SrcNetflowMonitorName.IsNull() {
		v.SrcNetflowMonitorName = saved.SrcNetflowMonitorName
	}
	// DstNetflowMonitorName
	if !allowed["dstNetflowMonitorName"] {
		v.DstNetflowMonitorName = saved.DstNetflowMonitorName
	} else if v.DstNetflowMonitorName.IsNull() && !saved.DstNetflowMonitorName.IsNull() {
		v.DstNetflowMonitorName = saved.DstNetflowMonitorName
	}
	// TemplateName
	if !allowed["templateName"] {
		v.TemplateName = saved.TemplateName
	} else if v.TemplateName.IsNull() && !saved.TemplateName.IsNull() {
		v.TemplateName = saved.TemplateName
	}
}

// buildFilteredTemplateInputs builds a map with only the allowed templateInputs
// fields for the given policyType using reflection. This produces a map that,
// when marshaled, only contains the fields applicable to the policy.
func buildFilteredTemplateInputs(ti *NDFCConfigDataTemplateInputsValue, policyType string) map[string]interface{} {
	allowed, ok := PolicyTemplateFields[policyType]
	if !ok || policyType == "" {
		// Unknown policy: include all fields
		result := make(map[string]interface{})
		v := reflect.ValueOf(ti).Elem()
		typ := v.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" || jsonTag == "-" {
				continue
			}
			jsonName := strings.Split(jsonTag, ",")[0]
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr {
				if fv.IsNil() {
					result[jsonName] = nil
				} else {
					result[jsonName] = fv.Elem().Interface()
				}
			} else {
				result[jsonName] = fv.Interface()
			}
		}
		return result
	}

	result := make(map[string]interface{})
	v := reflect.ValueOf(ti).Elem()
	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		jsonName := strings.Split(jsonTag, ",")[0]
		if !allowed[jsonName] {
			continue
		}

		fv := v.Field(i)
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				result[jsonName] = nil
			} else {
				result[jsonName] = fv.Elem().Interface()
			}
		} else {
			result[jsonName] = fv.Interface()
		}
	}
	return result
}

// BuildFilteredPayload builds a JSON payload for a single link with only the
// templateInputs fields applicable to the policyType. All other fields on
// NDFCLinkModel are included as-is.
func BuildFilteredPayload(model *NDFCLinkModel) ([]byte, error) {
	tiMap := buildFilteredTemplateInputs(&model.ConfigData.TemplateInputs, model.ConfigData.PolicyType)

	payload := map[string]interface{}{
		"srcSwitchId":      model.SrcSwitchId,
		"dstSwitchId":      model.DstSwitchId,
		"srcFabricName":    model.SrcFabricName,
		"srcSwitchName":    model.SrcSwitchName,
		"srcInterfaceName": model.SrcInterfaceName,
		"dstFabricName":    model.DstFabricName,
		"dstSwitchName":    model.DstSwitchName,
		"dstInterfaceName": model.DstInterfaceName,
		"srcClusterName":   model.SrcClusterName,
		"dstClusterName":   model.DstClusterName,
		"linkType":         model.LinkType,
		"configData": map[string]interface{}{
			"policyType":     model.ConfigData.PolicyType,
			"templateInputs": tiMap,
		},
	}

	return json.Marshal(payload)
}

// BuildFilteredCreatePayload builds the full create payload wrapped in {"links": [...]}.
func BuildFilteredCreatePayload(model *NDFCLinkModel) ([]byte, error) {
	linkBytes, err := BuildFilteredPayload(model)
	if err != nil {
		return nil, err
	}

	var linkMap map[string]interface{}
	if err := json.Unmarshal(linkBytes, &linkMap); err != nil {
		return nil, err
	}

	wrapper := map[string]interface{}{
		"links": []interface{}{linkMap},
	}
	return json.Marshal(wrapper)
}
