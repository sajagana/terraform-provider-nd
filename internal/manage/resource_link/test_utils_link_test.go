// Code generated;  DO NOT EDIT.

package resource_link

import (
	"strconv"
	"terraform-provider-nd/internal/manage/resource_link"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func LinkModelHelperStateCheck(RscName string, c resource_link.NDFCLinkModel, attrPath path.Path) []resource.TestCheckFunc {
	ret := []resource.TestCheckFunc{}

	if c.LinkId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_id").String(), c.LinkId))
	}
	if c.SrcSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_switch_id").String(), c.SrcSwitchId))
	}
	if c.DstSwitchId != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_switch_id").String(), c.DstSwitchId))
	}
	if c.SrcFabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_fabric_name").String(), c.SrcFabricName))
	}
	if c.SrcSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_switch_name").String(), c.SrcSwitchName))
	}
	if c.SrcInterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_interface_name").String(), c.SrcInterfaceName))
	}
	if c.DstFabricName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_fabric_name").String(), c.DstFabricName))
	}
	if c.DstSwitchName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_switch_name").String(), c.DstSwitchName))
	}
	if c.DstInterfaceName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_interface_name").String(), c.DstInterfaceName))
	}
	if c.SrcClusterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_cluster_name").String(), c.SrcClusterName))
	}
	if c.DstClusterName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_cluster_name").String(), c.DstClusterName))
	}
	if c.LinkType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_type").String(), c.LinkType))
	}
	if c.ConfigData.PolicyType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("policy_type").String(), c.ConfigData.PolicyType))
	}
	if c.ConfigData.SrcInterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_interface_description").String(), c.ConfigData.SrcInterfaceDescription))
	}
	if c.ConfigData.SrcInterfaceConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_interface_config").String(), c.ConfigData.SrcInterfaceConfig))
	}
	if c.ConfigData.DstInterfaceDescription != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_interface_description").String(), c.ConfigData.DstInterfaceDescription))
	}
	if c.ConfigData.DstInterfaceConfig != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_interface_config").String(), c.ConfigData.DstInterfaceConfig))
	}
	if c.InterfaceAdminState != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_admin_state").String(), strconv.FormatBool(*c.InterfaceAdminState)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_admin_state").String(), "true"))
	}
	if c.ConfigData.Mtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), strconv.Itoa(int(*c.ConfigData.Mtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu").String(), "9216"))
	}
	if c.ConfigData.LinkMtu != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_mtu").String(), strconv.Itoa(int(*c.ConfigData.LinkMtu))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("link_mtu").String(), "9216"))
	}
	if c.ConfigData.Speed != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("speed").String(), c.ConfigData.Speed))
	}
	if c.ConfigData.Fec != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fec").String(), c.ConfigData.Fec))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("fec").String(), "auto"))
	}
	if c.Macsec != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec").String(), strconv.FormatBool(*c.Macsec)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec").String(), "false"))
	}
	if c.ConfigData.SrcIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ip").String(), c.ConfigData.SrcIp))
	}
	if c.ConfigData.SrcIpv6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ipv6").String(), c.ConfigData.SrcIpv6))
	}
	if c.ConfigData.SrcIpAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ip_address").String(), c.ConfigData.SrcIpAddress))
	}
	if c.ConfigData.SrcIpAddressMask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ip_address_mask").String(), c.ConfigData.SrcIpAddressMask))
	}
	if c.ConfigData.SrcIpv6AddressMask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ipv6_address_mask").String(), c.ConfigData.SrcIpv6AddressMask))
	}
	if c.ConfigData.DstIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ip").String(), c.ConfigData.DstIp))
	}
	if c.ConfigData.DstIpv6 != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ipv6").String(), c.ConfigData.DstIpv6))
	}
	if c.ConfigData.DstIpAddress != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ip_address").String(), c.ConfigData.DstIpAddress))
	}
	if c.ConfigData.DstIpAddressMask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ip_address_mask").String(), c.ConfigData.DstIpAddressMask))
	}
	if c.ConfigData.DstIpv6Address != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ipv6_address").String(), c.ConfigData.DstIpv6Address))
	}
	if c.ConfigData.DstIpv6AddressMask != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ipv6_address_mask").String(), c.ConfigData.DstIpv6AddressMask))
	}
	if c.ConfigData.SrcEbgpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_ebgp_asn").String(), c.ConfigData.SrcEbgpAsn))
	}
	if c.ConfigData.DstEbgpAsn != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_ebgp_asn").String(), c.ConfigData.DstEbgpAsn))
	}
	if c.ConfigData.EbgpPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_password").String(), c.ConfigData.EbgpPassword))
	}
	if c.ConfigData.EbgpAuthKeyEncryptionType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_auth_key_encryption_type").String(), c.ConfigData.EbgpAuthKeyEncryptionType))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_auth_key_encryption_type").String(), "3des"))
	}
	if c.EbgpBfd != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_bfd").String(), strconv.FormatBool(*c.EbgpBfd)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_bfd").String(), "false"))
	}
	if c.EbgpLogNeighborChange != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_log_neighbor_change").String(), strconv.FormatBool(*c.EbgpLogNeighborChange)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_log_neighbor_change").String(), "false"))
	}
	if c.ConfigData.EbgpMaximumPaths != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_maximum_paths").String(), strconv.Itoa(int(*c.ConfigData.EbgpMaximumPaths))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_maximum_paths").String(), "64"))
	}
	if c.ConfigData.EbgpMultihop != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_multihop").String(), strconv.Itoa(int(*c.ConfigData.EbgpMultihop))))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_multihop").String(), "5"))
	}
	if c.EbgpSendComboth != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_send_comboth").String(), strconv.FormatBool(*c.EbgpSendComboth)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ebgp_send_comboth").String(), "false"))
	}
	if c.EnableEbgpPassword != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_ebgp_password").String(), strconv.FormatBool(*c.EnableEbgpPassword)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("enable_ebgp_password").String(), "false"))
	}
	if c.InheritEbgpPasswordMsdSettings != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inherit_ebgp_password_msd_settings").String(), strconv.FormatBool(*c.InheritEbgpPasswordMsdSettings)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inherit_ebgp_password_msd_settings").String(), "true"))
	}
	if c.ConfigData.DefaultVrfEbgpNeighborPassword != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("default_vrf_ebgp_neighbor_password").String(), c.ConfigData.DefaultVrfEbgpNeighborPassword))
	}
	if c.ConfigData.InterfaceVrf != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("interface_vrf").String(), c.ConfigData.InterfaceVrf))
	}
	if c.ConfigData.SrcVrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_vrf_name").String(), c.ConfigData.SrcVrfName))
	}
	if c.ConfigData.DstVrfName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_vrf_name").String(), c.ConfigData.DstVrfName))
	}
	if c.ConfigData.VrfNameNxPeerSwitch != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("vrf_name_nx_peer_switch").String(), c.ConfigData.VrfNameNxPeerSwitch))
	}
	if c.AutoGenConfigPeer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_peer").String(), strconv.FormatBool(*c.AutoGenConfigPeer)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_peer").String(), "false"))
	}
	if c.AutoGenConfigDefaultVrf != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_default_vrf").String(), strconv.FormatBool(*c.AutoGenConfigDefaultVrf)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_default_vrf").String(), "false"))
	}
	if c.AutoGenConfigNxPeerDefaultVrf != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_nx_peer_default_vrf").String(), strconv.FormatBool(*c.AutoGenConfigNxPeerDefaultVrf)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("auto_gen_config_nx_peer_default_vrf").String(), "false"))
	}
	if c.ConfigData.TemplateConfigGenPeer != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_config_gen_peer").String(), c.ConfigData.TemplateConfigGenPeer))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_config_gen_peer").String(), "Ext_VRF_Lite_Jython"))
	}
	if c.ConfigData.RedistribEbgpRouteMapName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("redistrib_ebgp_route_map_name").String(), c.ConfigData.RedistribEbgpRouteMapName))
	}
	if c.RedistributeRouteServer != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("redistribute_route_server").String(), strconv.FormatBool(*c.RedistributeRouteServer)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("redistribute_route_server").String(), "false"))
	}
	if c.ConfigData.RouteServerRoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("route_server_routing_tag").String(), c.ConfigData.RouteServerRoutingTag))
	}
	if c.ConfigData.RoutingTag != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("routing_tag").String(), c.ConfigData.RoutingTag))
	}
	if c.DciTracking != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_tracking").String(), strconv.FormatBool(*c.DciTracking)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_tracking").String(), "false"))
	}
	if c.DciTrackingEnableFlag != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_tracking_enable_flag").String(), strconv.FormatBool(*c.DciTrackingEnableFlag)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dci_tracking_enable_flag").String(), "false"))
	}
	if c.InheritTtagFabricSetting != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inherit_ttag_fabric_setting").String(), strconv.FormatBool(*c.InheritTtagFabricSetting)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("inherit_ttag_fabric_setting").String(), "true"))
	}
	if c.OverrideFabricMacsec != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("override_fabric_macsec").String(), strconv.FormatBool(*c.OverrideFabricMacsec)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("override_fabric_macsec").String(), "false"))
	}
	if c.ConfigData.MacsecCipherSuite != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_cipher_suite").String(), c.ConfigData.MacsecCipherSuite))
	}
	if c.ConfigData.MacsecPrimaryKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_primary_key_string").String(), c.ConfigData.MacsecPrimaryKeyString))
	}
	if c.ConfigData.MacsecPrimaryCryptographicAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_primary_cryptographic_algorithm").String(), c.ConfigData.MacsecPrimaryCryptographicAlgorithm))
	}
	if c.ConfigData.MacsecFallbackKeyString != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_key_string").String(), c.ConfigData.MacsecFallbackKeyString))
	}
	if c.ConfigData.MacsecFallbackCryptographicAlgorithm != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("macsec_fallback_cryptographic_algorithm").String(), c.ConfigData.MacsecFallbackCryptographicAlgorithm))
	}
	if c.Qkd != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("qkd").String(), strconv.FormatBool(*c.Qkd)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("qkd").String(), "false"))
	}
	if c.IgnoreCertificate != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ignore_certificate").String(), strconv.FormatBool(*c.IgnoreCertificate)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ignore_certificate").String(), "false"))
	}
	if c.ConfigData.SrcKmeServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_kme_server_ip").String(), c.ConfigData.SrcKmeServerIp))
	}
	if c.ConfigData.SrcMacsecKeyChainPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_macsec_key_chain_prefix").String(), c.ConfigData.SrcMacsecKeyChainPrefix))
	}
	if c.ConfigData.SrcQkdProfileName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_qkd_profile_name").String(), c.ConfigData.SrcQkdProfileName))
	}
	if c.ConfigData.SrcTrustpointLabel != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_trustpoint_label").String(), c.ConfigData.SrcTrustpointLabel))
	}
	if c.ConfigData.DstKmeServerIp != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_kme_server_ip").String(), c.ConfigData.DstKmeServerIp))
	}
	if c.ConfigData.DstMacsecKeyChainPrefix != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_macsec_key_chain_prefix").String(), c.ConfigData.DstMacsecKeyChainPrefix))
	}
	if c.ConfigData.DstQkdProfileName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_qkd_profile_name").String(), c.ConfigData.DstQkdProfileName))
	}
	if c.ConfigData.DstTrustpointLabel != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_trustpoint_label").String(), c.ConfigData.DstTrustpointLabel))
	}
	if c.DhcpRelayOnSrcInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_relay_on_src_interface").String(), strconv.FormatBool(*c.DhcpRelayOnSrcInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_relay_on_src_interface").String(), "false"))
	}
	if c.DhcpRelayOnDstInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_relay_on_dst_interface").String(), strconv.FormatBool(*c.DhcpRelayOnDstInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dhcp_relay_on_dst_interface").String(), "false"))
	}
	if c.BfdEchoOnSrcInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_echo_on_src_interface").String(), strconv.FormatBool(*c.BfdEchoOnSrcInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_echo_on_src_interface").String(), "false"))
	}
	if c.BfdEchoOnDstInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_echo_on_dst_interface").String(), strconv.FormatBool(*c.BfdEchoOnDstInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bfd_echo_on_dst_interface").String(), "false"))
	}
	if c.ConfigData.MtuType != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("mtu_type").String(), c.ConfigData.MtuType))
	}
	if c.Ipv4Pim != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_pim").String(), strconv.FormatBool(*c.Ipv4Pim)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_pim").String(), "false"))
	}
	if c.Ipv4Trm != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_trm").String(), strconv.FormatBool(*c.Ipv4Trm)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv4_trm").String(), "false"))
	}
	if c.Ipv6Pim != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_pim").String(), strconv.FormatBool(*c.Ipv6Pim)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_pim").String(), "false"))
	}
	if c.Ipv6Trm != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_trm").String(), strconv.FormatBool(*c.Ipv6Trm)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ipv6_trm").String(), "false"))
	}
	if c.IpRedirects != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip_redirects").String(), strconv.FormatBool(*c.IpRedirects)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("ip_redirects").String(), "false"))
	}
	if c.SkipConfigGeneration != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("skip_config_generation").String(), strconv.FormatBool(*c.SkipConfigGeneration)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("skip_config_generation").String(), "false"))
	}
	if c.ConfigData.BpduGuard != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("bpdu_guard").String(), c.ConfigData.BpduGuard))
	}
	if c.PortTypeFast != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_type_fast").String(), strconv.FormatBool(*c.PortTypeFast)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("port_type_fast").String(), "true"))
	}
	if c.ConfigData.NativeVlan != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("native_vlan").String(), strconv.Itoa(int(*c.ConfigData.NativeVlan))))
	}
	if c.ConfigData.TrunkAllowedVlans != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("trunk_allowed_vlans").String(), c.ConfigData.TrunkAllowedVlans))
	}
	if c.NetflowOnSrcInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_on_src_interface").String(), strconv.FormatBool(*c.NetflowOnSrcInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_on_src_interface").String(), "false"))
	}
	if c.NetflowOnDstInterface != nil {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_on_dst_interface").String(), strconv.FormatBool(*c.NetflowOnDstInterface)))
	} else {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("netflow_on_dst_interface").String(), "false"))
	}
	if c.ConfigData.SrcNetflowMonitorName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("src_netflow_monitor_name").String(), c.ConfigData.SrcNetflowMonitorName))
	}
	if c.ConfigData.DstNetflowMonitorName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("dst_netflow_monitor_name").String(), c.ConfigData.DstNetflowMonitorName))
	}
	if c.ConfigData.TemplateName != "" {
		ret = append(ret, resource.TestCheckResourceAttr(RscName, attrPath.AtName("template_name").String(), c.ConfigData.TemplateName))
	}
	return ret
}
