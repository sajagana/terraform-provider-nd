
resource "ndfc_multi_cluster_connectivity" "test_resource_multi_cluster_connectivity_1" {
  hostname                        = "198.18.133.203"
  type                            = "ND"
  username                        = "admin"
  password                        = "password"
  login_domain                    = "local"
  fabric_name                     = "apic1"
  license_tier                    = "premier"
  security_domain                 = "all"
  verify_ca                       = false
  orchestration_status            = "enabled"
  telemetry_status                = "enabled"
  telemetry_network               = "inband"
  telemetry_epg                   = "uni/tn-infra/ap-access/epg-default"
  telemetry_streaming_protocol    = "ipv4"
  multi_cluster_login_domain_name = "my-mc-domain"
  annotation                      = key1
  force_remove                    = false
  re_register                     = false
}