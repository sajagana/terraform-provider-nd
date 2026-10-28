
resource "nd_fabric_groups" "test_resource_fabric_groups_1" {
  fabric_name                                     = "example_fabric_group"
  fabric_type                                     = "vxlan"
  anycast_gateway_mac                             = "2020.0000.00aa"
  category                                        = "fabricGroup"
  vrf_template                                    = "Default_VRF_Universal"
  network_template                                = "Default_Network_Universal"
  vrf_extension_template                          = "Default_VRF_Extension_Universal"
  network_extension_template                      = "Default_Network_Extension_Universal"
  l2_vni_range                                    = "30000-49000"
  l3_vni_range                                    = "50000-59000"
  scheduled_backup                                = false
  scheduled_backup_time                           = "02:00"
  underlay_ipv6                                   = false
  private_vlan                                    = false
  default_private_vlan_secondary_network_template = "Pvlan_Secondary_Network"
  security_group_tag                              = "off"
  security_group_tag_prefix                       = "SG_"
  security_group_tag_id_range                     = "10000-14000"
  security_group_tag_preprovision                 = false
  security_group_tag_mac_segmentation             = false
  downstream_vni                                  = false
  downstream_l2_vni_range                         = "60000-69000"
  downstream_l3_vni_range                         = "80000-89000"
  multisite_overlay_inter_connect_type            = "manual"
  route_server_collection = [
    {
      route_server_ip  = "192.0.2.10"
      route_server_asn = "65001"
    }
  ]
  route_server_redistribute_direct_route_map            = false
  route_server_routing_tag                              = 54321
  enable_multisite_overlay_inter_fabric_bgp_description = true
  auto_multisite_underlay_inter_connect                 = false
  bgp_send_community                                    = false
  bgp_log_neighbor_change                               = false
  bgp_bfd                                               = false
  multisite_delay_restore                               = 300
  multisite_inter_connect_bgp_authentication            = false
  multisite_inter_connect_bgp_auth_key_type             = "3des"
  multisite_inter_connect_bgp_key                       = "exampleEncryptedBgpKey"
  multisite_loopback_id                                 = 100
  border_gateway_routing_tag                            = 54321
  tenant_routed_multicast_ipv4_ipv6                     = false
  auto_configure_cloud_sec                              = false
  cloud_sec_key                                         = "070C285F4D06"
  cloud_sec_algorithm                                   = "AES_128_CMAC"
  cloud_sec_enforcement                                 = "strict"
  cloud_sec_report_timer                                = 5
  multisite_loopback_ip_range                           = "10.10.0.0/24"
  multisite_underlay_subnet_range                       = "10.10.1.0/24"
  multisite_underlay_subnet_target_mask                 = 30
  multisite_loopback_ipv6_range                         = "fd00::a10:0/120"
  multisite_underlay_ipv6_subnet_range                  = "fd00::a11:0/120"
  multisite_underlay_ipv6_subnet_target_mask            = 126
}