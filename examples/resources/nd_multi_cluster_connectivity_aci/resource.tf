
resource "nd_multi_cluster_connectivity_aci" "test_resource_multi_cluster_connectivity_aci_1" {
  cluster_type              = "APIC"
  cluster_name              = "ins15-apic1"
  host_name                 = "10.195.219.153:443"
  user_name                 = "admin"
  password                  = "Q2lzY29pbnMzOTY1IQ=="
  login_domain              = "DefaultAuth"
  fabric_name               = "ins15-apic1"
  license_tier              = "premier"
  security_domain           = "all"
  validate_peer_certificate = false
  enable_orchestration      = "enabled"
  enable_telemetry          = "enabled"
  telemetry_collection      = "inband"
  telemetry_streaming       = "ipv4"
  telemetry_epg             = "uni/tn-mgmt/mgmtp-default/inb-default"
}