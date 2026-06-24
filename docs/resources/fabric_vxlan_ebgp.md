---
subcategory: "manage"
layout: "nd"
page_title: "ND: nd_fabric_vxlan_ebgp"
sidebar_current: "docs-nd-resource-nd_fabric_vxlan_ebgp"
description: |-
  
---

# nd_fabric_vxlan_ebgp #



## API Information ##
* Fabric Vxlan Ebgp Management API information is not defined in the YAML.
* API Endpoint: not defined in the YAML.

## GUI Information ##
* Location: not defined in the YAML.

## Example Usage ##

The configuration snippet below shows all possible attributes of the Fabric Vxlan Ebgp resource.

!> This example might not be valid configuration and is only used to show all possible attributes.

Example source: `examples/resources/nd_fabric_vxlan_ebgp/resource.tf`

```hcl
resource "nd_fabric_vxlan_ebgp" "test_resource_fabric_vxlan_ebgp_1" {
  fabric_name                                     = "my_fabric"
  license_tier                                    = "premier"
  controller_status                               = "enabled"
  telemetry_status                                = "enabled"
  orchestration_status                            = "enabled"
  trap_forwarder_status                           = "enabled"
  telemetry_collection                            = false
  telemetry_collection_type                       = "outOfBand"
  telemetry_streaming_protocol                    = "ipv4"
  telemetry_source_interface                      = "eth0"
  telemetry_source_vrf                            = "vrf1"
  security_domain                                 = "all"
  bgp_asn                                         = "55000"
  super_spine_bgp_as                              = "65000"
  leaf_bgp_as                                     = "65001"
  border_bgp_as                                   = "65002"
  bgp_as_mode                                     = "multiAS"
  bgp_asn_auto_allocation                         = true
  bgp_asn_range                                   = "65000-65535"
  bgp_allow_as_in_num                             = 1
  bgp_max_path                                    = 4
  bgp_underlay_failure_protect                    = false
  auto_configure_ebgp_evpn_peering                = true
  allow_leaf_same_as                              = false
  assign_ipv4_to_loopback0                        = true
  target_subnet_mask                              = 30
  anycast_gateway_mac                             = "2020.0000.00aa"
  performance_monitoring                          = false
  replication_mode                                = "multicast"
  multicast_group_subnet                          = "239.1.1.0/25"
  tenant_routed_multicast                         = false
  rendezvous_point_count                          = 2
  category                                        = "fabric"
  alert_suspend                                   = "disabled"
  rendezvous_point_loopback_id                    = 254
  vpc_peer_link_vlan                              = "3600"
  vpc_peer_link_enable_native_vlan                = false
  vpc_peer_keep_alive_option                      = "management"
  vpc_auto_recovery_timer                         = 360
  vpc_delay_restore_timer                         = 150
  vpc_peer_link_port_channel_id                   = "500"
  vpc_ipv6_neighbor_discovery_sync                = true
  advertise_physical_ip                           = false
  vpc_domain_id_range                             = "1-1000"
  bgp_loopback_id                                 = 0
  nve_loopback_id                                 = 1
  vrf_template                                    = "Default_VRF_Universal"
  network_template                                = "Default_Network_Universal"
  vrf_extension_template                          = "Default_VRF_Extension_Universal"
  network_extension_template                      = "Default_Network_Extension_Universal"
  l3_vni_no_vlan_default_option                   = false
  site_id                                         = "65001"
  fabric_mtu                                      = 9216
  l2_host_interface_mtu                           = 9216
  tenant_dhcp                                     = true
  nxapi                                           = false
  nxapi_https_port                                = 443
  nxapi_http                                      = true
  nxapi_http_port                                 = 80
  snmp_trap                                       = true
  anycast_border_gateway_advertise_physical_ip    = false
  greenfield_debug_flag                           = "disable"
  tcam_allocation                                 = true
  real_time_interface_statistics_collection       = false
  interface_statistics_load_interval              = 10
  bgp_loopback_ip_range                           = "10.2.0.0/22"
  nve_loopback_ip_range                           = "10.3.0.0/22"
  anycast_rendezvous_point_ip_range               = "10.254.254.0/24"
  intra_fabric_subnet_range                       = "10.4.0.0/16"
  l2_vni_range                                    = "30000-49000"
  l3_vni_range                                    = "50000-59000"
  network_vlan_range                              = "2300-2999"
  vrf_vlan_range                                  = "2000-2299"
  sub_interface_dot1q_range                       = "2-511"
  vrf_lite_auto_config                            = "manual"
  vrf_lite_subnet_range                           = "10.33.0.0/16"
  vrf_lite_subnet_target_mask                     = 30
  vrf_lite_ipv6_subnet_range                      = "2001::10.33.0.0/16"
  vrf_lite_ipv6_subnet_target_mask                = 126
  auto_unique_vrf_lite_ip_prefix                  = false
  per_vrf_loopback_auto_provision                 = false
  per_vrf_loopback_ip_range                       = "10.5.0.0/22"
  per_vrf_loopback_auto_provision_ipv6            = false
  per_vrf_loopback_ipv6_range                     = "fd00::a05:0/112"
  banner                                          = "my_banner"
  day0_bootstrap                                  = false
  local_dhcp_server                               = false
  dhcp_protocol_version                           = "dhcpv4"
  dhcp_start_address                              = "192.168.1.1"
  dhcp_end_address                                = "192.168.1.1"
  management_gateway                              = "192.168.1.1"
  management_ipv4_prefix                          = 24
  management_ipv6_prefix                          = 64
  extra_config_nxos_bootstrap                     = "aaa authentication login default group radius"
  real_time_backup                                = true
  scheduled_backup                                = true
  scheduled_backup_time                           = "22:00"
  underlay_ipv6                                   = false
  ipv6_multicast_group_subnet                     = "ff1e::/121"
  tenant_routed_multicast_ipv6                    = false
  mvpn_vrf_route_import_id                        = true
  mvpn_vrf_route_import_id_range                  = "1-65535"
  vrf_route_import_id_reallocation                = false
  l3vni_multicast_group                           = "239.1.1.0"
  l3_vni_ipv6_multicast_group                     = "ff1e::"
  rendezvous_point_mode                           = "asm"
  auto_generate_multicast_group_address           = false
  phantom_rendezvous_point_loopback_id1           = 2
  phantom_rendezvous_point_loopback_id2           = 3
  phantom_rendezvous_point_loopback_id3           = 4
  phantom_rendezvous_point_loopback_id4           = 5
  advertise_physical_ip_on_border                 = true
  fabric_vpc_domain_id                            = false
  shared_vpc_domain_id                            = 1
  vpc_layer3_peer_router                          = true
  fabric_vpc_qos                                  = false
  fabric_vpc_qos_policy_name                      = "spine_qos_for_fabric_vpc_peering"
  anycast_loopback_id                             = 10
  bgp_authentication                              = false
  bgp_authentication_key_type                     = "3des"
  bgp_authentication_key                          = "BGP_KEY"
  pim_hello_authentication                        = false
  pim_hello_authentication_key                    = "PIM_KEY"
  bfd                                             = false
  bfd_ibgp                                        = false
  bfd_authentication                              = false
  bfd_authentication_key_id                       = 100
  bfd_authentication_key                          = "BFD_KEY"
  macsec                                          = false
  macsec_cipher_suite                             = "GCM-AES-XPN-256"
  macsec_key_string                               = "0123456789ABCDEF"
  macsec_algorithm                                = "AES_128_CMAC"
  macsec_fallback_key_string                      = "0123456789ABCDEF"
  macsec_fallback_algorithm                       = "AES_128_CMAC"
  macsec_report_timer                             = 5
  overlay_mode                                    = "cli"
  private_vlan                                    = false
  default_private_vlan_secondary_network_template = "Pvlan_Secondary_Network"
  power_redundancy_mode                           = "redundant"
  copp_policy                                     = "strict"
  nve_hold_down_timer                             = 180
  cdp                                             = false
  next_generation_oam                             = true
  ngoam_south_bound_loop_detect                   = false
  ngoam_south_bound_loop_detect_probe_interval    = 300
  ngoam_south_bound_loop_detect_recovery_interval = 600
  strict_config_compliance_mode                   = false
  advanced_ssh_option                             = false
  ptp                                             = false
  ptp_loopback_id                                 = 0
  ptp_domain_id                                   = 0
  default_queuing_policy                          = false
  default_queuing_policy_cloudscale               = "queuing_policy_default_8q_cloudscale"
  default_queuing_policy_r_series                 = "queuing_policy_default_r_series"
  default_queuing_policy_other                    = "queuing_policy_default_other"
  aiml_qos                                        = false
  aiml_qos_policy                                 = "400G"
  priority_flow_control_watch_interval            = 101
  static_underlay_ip_allocation                   = false
  bgp_loopback_ipv6_range                         = "fd00::a02:0/119"
  nve_loopback_ipv6_range                         = "fd00::a03:0/118"
  extra_config_aaa                                = "radius-server host 10.1.1.1"
  aaa                                             = false
  ipv6_link_local                                 = true
  fabric_interface_type                           = "p2p"
  ipv6_subnet_target_mask                         = 126
  link_state_routing_protocol                     = "ospf"
  route_reflector_count                           = 2
  vpc_tor_delay_restore_timer                     = 30
  leaf_tor_id_range                               = false
  leaf_tor_vpc_port_channel_id_range              = "1-499"
  link_state_routing_tag                          = "UNDERLAY"
  ospf_area_id                                    = "0.0.0.0"
  ospf_authentication                             = false
  ospf_authentication_key_id                      = 127
  ospf_authentication_key                         = "OSPF_KEY"
  isis_level                                      = "level-2"
  isis_area_number                                = "0001"
  isis_point_to_point                             = true
  isis_authentication                             = false
  isis_authentication_keychain_name               = "CiscoIsisAuth"
  isis_authentication_keychain_key_id             = 127
  isis_authentication_key                         = "ISIS_KEY"
  isis_overload                                   = true
  isis_overload_elapse_time                       = 60
  bfd_ospf                                        = false
  bfd_isis                                        = false
  bfd_pim                                         = false
  auto_bgp_neighbor_description                   = true
  ibgp_peer_template                              = " template peer RR"
  leafibgp_peer_template                          = " template peer LEAF"
  security_group_tag                              = false
  security_group_tag_id_range                     = "10000-14000"
  security_group_tag_preprovision                 = false
  security_group_status                           = "disabled"
  vrf_lite_macsec                                 = false
  quantum_key_distribution                        = false
  vrf_lite_macsec_cipher_suite                    = "GCM-AES-XPN-256"
  vrf_lite_macsec_key_string                      = "0123456789ABCDEF"
  vrf_lite_macsec_algorithm                       = "AES_128_CMAC"
  vrf_lite_macsec_fallback_key_string             = "0123456789ABCDEF"
  vrf_lite_macsec_fallback_algorithm              = "AES_128_CMAC"
  quantum_key_distribution_profile_name           = "QKDProfile"
  key_management_entity_server_ip                 = "192.168.1.1"
  key_management_entity_server_port               = 42
  trustpoint_label                                = "tp_label"
  skip_certificate_verification                   = false
  host_interface_admin_state                      = true
  brownfield_network_name_format                  = "Auto_Net_VNI$$VNI$$_VLAN$$VLAN_ID$$"
  brownfield_skip_overlay_network_attachments     = false
  policy_based_routing                            = false
  ptp_vlan_id                                     = 2
  mpls_handoff                                    = false
  mpls_loopback_identifier                        = 101
  mpls_isis_area_number                           = "0001"
  stp_root_option                                 = "unmanaged"
  stp_vlan_range                                  = "1,3-5,7,9-11"
  mst_instance_range                              = "0-3,5,7-9"
  stp_bridge_priority                             = 0
  allow_vlan_on_leaf_tor_pairing                  = "none"
  pre_interface_config_leaf                       = "speed 40000"
  pre_interface_config_spine                      = "speed 40000"
  pre_interface_config_tor                        = "speed 40000"
  extra_config_leaf                               = "no shutdown"
  extra_config_spine                              = "no shutdown"
  extra_config_tor                                = "no shutdown"
  extra_config_intra_fabric_links                 = "no shutdown"
  mpls_loopback_ip_range                          = "10.101.0.0/25"
  ipv6_subnet_range                               = "fd00::a04:0/112"
  router_id_range                                 = "10.2.0.0/23"
  auto_symmetric_vrf_lite                         = false
  auto_vrf_lite_default_vrf                       = false
  auto_symmetric_default_vrf                      = false
  default_vrf_redistribution_bgp_route_map        = "extcon-rmap-filter"
  ip_service_level_agreement_id_range             = "10000-19999"
  object_tracking_number_range                    = "100-299"
  service_network_vlan_range                      = "3000-3199"
  route_map_sequence_number_range                 = "1-65534"
  inband_management                               = false
  seed_switch_core_interfaces                     = "Ethernet1/1"
  spine_switch_core_interfaces                    = "Ethernet1/1"
  inband_dhcp_servers                             = "10.0.0.1"
  un_numbered_bootstrap_lb_id                     = 253
  un_numbered_dhcp_start_address                  = "192.168.1.1"
  un_numbered_dhcp_end_address                    = "192.168.1.1"
  heartbeat_interval                              = 190
  allow_smart_switch_onboarding                   = false
  enable_dpu_pinning                              = false
  connectivity_domain_name                        = "my_domain"
  hypershield_connectivity_proxy_server           = "10.0.0.1"
  hypershield_connectivity_proxy_server_port      = 8080
  hypershield_connectivity_source_intf            = "loopback0"
  netflow_enable                                  = false
  traffic_analytics                               = "enabled"
  net_flow                                        = false
  s_flow                                          = false
  flow_telemetry                                  = false
  traffic_analytics_rules_enabled                 = true
  traffic_analytics_mode                          = "full"
  udp_categorization                              = "supported"
  traffic_analytics_filter_rules                  = "supported"
  operating_mode                                  = "flowTelemetry"
  udp_categorization_support                      = "supported"
  microburst                                      = false
  sensitivity                                     = "low"
  analysis_settings_is_enabled                    = false
  server                                          = "stream-server-1"
  export_type                                     = "full"
  export_format                                   = "json"
  syslog_servers                                  = ["example_servers_item"]
  syslog_facility                                 = "local7"
  syslog_anomalies                                = ["critical"]
}
```

All examples for the Fabric Vxlan Ebgp resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/nd_fabric_vxlan_ebgp) folder.

## Schema ##

### Required ###
* `fabric_name` (name) - (String) The name of the fabric_vxlan_ebgp resource
* `bgp_asn` (bgpAsn) - (String) Autonomous system number 1-4294967295 | 1-65535[.0-65535]

### Optional ###
* `license_tier` (licenseTier) - (String) License Tier value of a fabric
* `telemetry_collection` (telemetryCollection) - (Bool) Enable telemetry collection
* `allowed_actions` (allowedActions) - (List:String) Actions allowed on the listed fabrics
* `super_spine_bgp_as` (superSpineBgpAs) - (String) BGP AS number for super spine switches 1-4294967295 | 1-65535[.0-65535]
* `leaf_bgp_as` (leafBgpAs) - (String) BGP AS number for leaf switches 1-4294967295 | 1-65535[.0-65535]
* `border_bgp_as` (borderBgpAs) - (String) BGP AS number for border switches 1-4294967295 | 1-65535[.0-65535]
* `bgp_as_mode` (bgpAsMode) - (String) Multi-AS Unique ASN per Leaf/Border/Border Gateway (Borders and border gateways are allowed to share ASN). Same-Tier-AS Leafs share one ASN, Borders/border gateways share one ASN
* `bgp_asn_range` (bgpAsnRange) - (String) BGP ASN range for automatic ASN allocation (e.g. 65000-65535)
* `bgp_allow_as_in_num` (bgpAllowAsInNum) - (Int64) Number of occurrences of the local AS number allowed in the BGP AS-path
* `bgp_max_path` (bgpMaxPath) - (Int64) Maximum number of BGP equal-cost paths
* `assign_ipv4_to_loopback0` (assignIpv4ToLoopback0) - (Bool) In an IPv6 routed fabric or VXLAN EVPN fabric with IPv6 underlay, assign IPv4 address used for BGP Router ID to the routing loopback interface
* `target_subnet_mask` (targetSubnetMask) - (Int64) Mask for underlay subnet IP range
* `anycast_gateway_mac` (anycastGatewayMac) - (String) Shared anycast gateway MAC address for all VTEPs (xxxx.xxxx.xxxx)
* `replication_mode` (replicationMode) - (String) Replication Mode for BUM Traffic
* `multicast_group_subnet` (multicastGroupSubnet) - (String) Multicast pool prefix between 8 to 30. A multicast group ipv4 from this pool is used for BUM traffic for each overlay network
* `rendezvous_point_count` (rendezvousPointCount) - (Int64) Number of spines acting as Rendezvous-Points (RPs)
* `location` (location) - (Struct) Location to access the resource
    * `latitude` (latitude) - (Float64) Latitude in decimal degrees
    * `longitude` (longitude) - (Float64) Longitude in decimal degrees
* `site_id` (siteId) - (String) For EVPN Multi-Site Support. Defaults to Fabric ASN
* `anycast_rendezvous_point_ip_range` (anycastRendezvousPointIpRange) - (String) Anycast or Phantom RP IP Address Range
* `intra_fabric_subnet_range` (intraFabricSubnetRange) - (String) Address range to assign numbered and peer link SVI IPs
* `vrf_lite_ipv6_subnet_range` (vrfLiteIpv6SubnetRange) - (String) IPv6 address range to assign P2P Aggregation-Core connections, and peering between vPC Aggregation switches
* `bootstrap_multi_subnet` (bootstrapMultiSubnet) - (String) Enter One Subnet Scope per line. Start_IP, End_IP, Gateway, Prefix e.g. 10.6.0.2, 10.6.0.9, 10.6.0.1, 24
* `auto_generate_multicast_group_address` (autoGenerateMulticastGroupAddress) - (Bool) Auto generate multicast group address
* `bgp_authentication_key_type` (bgpAuthenticationKeyType) - (String) 
* `macsec` (macsec) - (Bool) Enable MACsec on this
* `aiml_qos_policy` (aimlQosPolicy) - (String) Queuing Policy based on predominant fabric link speed 800G / 400G / 100G / 25G
* `ipv6_link_local` (ipv6LinkLocal) - (Bool) Enables IPv6 link-local Option under VRF SVI. Not applicable to L3VNI without VLAN config. NX-OS Specific
* `fabric_interface_type` (fabricInterfaceType) - (String) Numbered(Point-to-Point) or unNumbered
* `ipv6_subnet_target_mask` (ipv6SubnetTargetMask) - (Int64) Mask for Underlay Subnet IPv6 Range
* `link_state_routing_protocol` (linkStateRoutingProtocol) - (String) Underlay Routing Protocol. Used for spine-leaf connectivity
* `route_reflector_count` (routeReflectorCount) - (Int64) Number of spines acting as Route-Reflectors
* `link_state_routing_tag` (linkStateRoutingTag) - (String) Underlay routing protocol process tag
* `ospf_area_id` (ospfAreaId) - (String) OSPF Area Id in IP address format. Required if an OSPF process tag is specified
* `ospf_authentication` (ospfAuthentication) - (Bool) Whether to enable OSPF authentication
* `isis_level` (isisLevel) - (String) IS-IS level
* `isis_area_number` (isisAreaNumber) - (String) NET in form of XX.<4-hex-digit Custom Area Number>.XXXX.XXXX.XXXX.00, default Area Number is 0001. If area number in existing NETs matches the previous area number set in fabric settings and is different from the current area number, these NETs will be updated by Recalculate and Deploy
* `isis_authentication` (isisAuthentication) - (Bool) Enable IS-IS authentication
* `bfd_ospf` (bfdOspf) - (Bool) Enable BFD For OSPF
* `bfd_isis` (bfdIsis) - (Bool) Enable BFD For ISIS
* `bfd_pim` (bfdPim) - (Bool) Enable BFD For PIM
* `auto_bgp_neighbor_description` (autoBgpNeighborDescription) - (Bool) Generate BGP EVPN Neighbor Description
* `security_group_tag` (securityGroupTag) - (Bool) If set to strict, only security groups enabled child fabrics will be allowed
* `vrf_lite_macsec` (vrfLiteMacsec) - (Bool) Enable MACsec on DCI links. DCI MACsec fabric parameters are used for configuring MACsec on a DCI link if 'Use Link MACsec Setting' is disabled on the link
* `host_interface_admin_state` (hostInterfaceAdminState) - (Bool) Unshut Host Interfaces by Default
* `brownfield_network_name_format` (brownfieldNetworkNameFormat) - (String) Generated network name should be less than 64 characters
* `brownfield_skip_overlay_network_attachments` (brownfieldSkipOverlayNetworkAttachments) - (Bool) Skip Overlay Network Interface Attachments for Brownfield and Host Port Resync cases
* `policy_based_routing` (policyBasedRouting) - (Bool) Enable feature pbr, sla sender, epbr, or enable feature pbr, based on the L4-L7 Services use case
* `mpls_handoff` (mplsHandoff) - (Bool) Enable MPLS Handoff
* `mpls_isis_area_number` (mplsIsisAreaNumber) - (String) NET in form of XX.<4-hex-digit Custom Area Number>.XXXX.XXXX.XXXX.00, default Area Number is 0001, used only if routing protocol on DCI MPLS link is is-is
* `stp_root_option` (stpRootOption) - (String) Protocol to be used for configuring Root Bridge: rpvst+: Rapid Per-VLAN Spanning Tree, mst: Multiple Spanning Tree, unmanaged: STP Root not managed by ND. Note: Spanning Tree Settings and Bridge Configs are applicable at Aggregation layer only
* `auto_symmetric_vrf_lite` (autoSymmetricVrfLite) - (Bool) Flag that controls auto generation of VRF Lite sub-interface and peering configuration on Aggregation and Core/Edge switches. If set, auto created VRF Lite links will have 'Auto Generate Flag' enabled
* `auto_vrf_lite_default_vrf` (autoVrfLiteDefaultVrf) - (Bool) Whether to auto generate Default VRF interface and BGP peering configuration on VRF LITE IFC auto deployment. If set, auto created VRF Lite IFC links will have 'Auto Deploy Default VRF' enabled
* `auto_symmetric_default_vrf` (autoSymmetricDefaultVrf) - (Bool) Whether to auto generate Default VRF interface and BGP peering configuration on managed neighbor devices. If set, auto created VRF Lite IFC links will have 'Auto Deploy Default VRF for Peer' enabled
* `ip_service_level_agreement_id_range` (ipServiceLevelAgreementIdRange) - (String) Service Level Agreement (SLA) ID Range (minimum: 1, maximum: 655214748364735). Per switch SLA ID Range
* `object_tracking_number_range` (objectTrackingNumberRange) - (String) Tracked Object ID Range (minimum: 1, maximum: 512) Per switch tracked object ID Range
* `service_network_vlan_range` (serviceNetworkVlanRange) - (String) Service Network VLAN Range (minimum: 2, maximum: 4094). Per Switch Overlay Service Network VLAN Range
* `route_map_sequence_number_range` (routeMapSequenceNumberRange) - (String) Route Map Sequence Number Range (minimum: 1, maximum: 65534)
* `inband_management` (inbandManagement) - (Bool) Import switches with inband connectivity
* `heartbeat_interval` (heartbeatInterval) - (Int64) Heartbeat Interval in seconds
* `enable_dpu_pinning` (enableDpuPinning) - (Bool) Enable pinning of VRFs and networks to specific DPUs on smart switches
* `dns_collection` (dnsCollection) - (List:String) List of IPv4 and IPv6 DNS addresses
* `dns_vrf_collection` (dnsVrfCollection) - (List:String) DNS Server VRFs. One VRF for all DNS servers or a list of VRFs, one per DNS server
* `ntp_server_collection` (ntpServerCollection) - (List:String) List of NTP server IPv4/IPv6 addresses and/or hostnames
* `ntp_server_vrf_collection` (ntpServerVrfCollection) - (List:String) NTP Server VRFs. One VRF for all NTP servers or a list of VRFs, one per NTP server
* `syslog_server_collection` (syslogServerCollection) - (List:String) List of Syslog server IPv4/IPv6 addresses and/or hostnames
* `syslog_severity_collection` (syslogSeverityCollection) - (List:Int64) List of Syslog severity values, one per Syslog server
* `syslog_server_vrf_collection` (syslogServerVrfCollection) - (List:String) Syslog Server VRFs. One VRF for all Syslog servers or a list of VRFs, one per Syslog server
* `microburst` (microburst) - (Bool) Enable microburst detection
* `sensitivity` (sensitivity) - (String) Microburst sensitivity level
* `cost` (cost) - (Float64) Energy cost in USD/kWh
* `syslog_servers` (servers) - (List:String) Syslog servers to which alerts are sent
* `syslog_anomalies` (anomalies) - (List:String) Level of anomalies to be collected

### Read-Only ###
* `ipv6_anycast_rendezvous_point_ip_range` (ipv6AnycastRendezvousPointIpRange) - (String) Anycast RP IPv6 Address Range

## Importing

An existing Fabric Vxlan Ebgp resource can be imported with its identifier via the following command:

```shell
terraform import nd_fabric_vxlan_ebgp.example {id}
```

Starting in Terraform version 1.5, an existing Fabric Vxlan Ebgp resource can be imported using import blocks via the following configuration:

```hcl
import {
  id = "{id}"
  to = nd_fabric_vxlan_ebgp.example
}
```
