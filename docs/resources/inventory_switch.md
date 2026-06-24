---
subcategory: "manage"
layout: "nd"
page_title: "ND: nd_inventory_switch"
sidebar_current: "docs-nd-resource-nd_inventory_switch"
description: |-
  
---

# nd_inventory_switch #



## API Information ##
* Inventory Switch Management API information is not defined in the YAML.
* API Endpoint: not defined in the YAML.

## GUI Information ##
* Location: not defined in the YAML.

## Example Usage ##

The configuration snippet below shows all possible attributes of the Inventory Switch resource.

!> This example might not be valid configuration and is only used to show all possible attributes.

Example source: `examples/resources/nd_inventory_switch/resource.tf`

```hcl
resource "nd_inventory_switch" "test_resource_inventory_switch_1" {
  fabric_name             = "my_fabric"
  mode                    = "discovery"
  preserve_config         = true
  username                = "admin"
  password                = "mysecret"
  snmp_v3_auth_protocol   = "MD5"
  remote_credential_store = "local"
  max_hop                 = 2
  recalculate             = true
  deploy                  = true
}
```

All examples for the Inventory Switch resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/nd_inventory_switch) folder.

## Schema ##

### Required ###
* `fabric_name` (fabricName) - (String) Name of the fabric to add switches to
* `switches` (switches) - (Map) Map of switches to manage, keyed by serial number
    * `hostname` (hostname) - (String) Switch hostname
    * `serial_number` (serialNumber) - (String) Switch serial number
    * `ip_address` (ip) - (String) Switch management IP address
    * `model` (model) - (String) Switch hardware model
    * `software_version` (softwareVersion) - (String) Switch software version
    * `switch_role` (switchRole) - (String) Role of switch in fabric.
        * Default Value: `leaf`
    * `status` (status) - (String) Current switch status
    * `status_reason` (statusReason) - (String) Reason for current switch status (e.g., failure details)
    * `gateway_ip_mask` (gatewayIpMask) - (String) Gateway IP with mask for POAP (e.g., 10.1.1.1/24)
    * `discovery_auth_protocol` (discoveryAuthProtocol) - (String) Authentication protocol for POAP discovery
    * `poap_password` (poapPassword) - (Sensitive, String) Password for POAP bootstrap
    * `vdc_id` (vdcId) - (Int64) VDC ID for N7K switches
    * `vdc_mac` (vdcMac) - (String) VDC MAC for N7K switches

### Optional ###
* `mode` (mode) - (String) Operation mode - discovery (brownfield/greenfield) or bootstrap (POAP)
    * Default Value: `discovery`
* `username` (username) - (Sensitive, String) Switch login username
* `password` (password) - (Sensitive, String) Switch login password
* `snmp_v3_auth_protocol` (snmpV3AuthProtocol) - (String) SNMP v3 authentication protocol
    * Default Value: `MD5`
* `remote_credential_store` (remoteCredentialStore) - (String) Credential store type (local, cyberark)
    * Default Value: `local`
* `remote_credential_store_key` (remoteCredentialStoreKey) - (String) Key for remote credential store
* `max_hop` (maxHop) - (Int64) Maximum hops for CDP/LLDP discovery
    * Default Value: `0`

### Read-Only ###
* `platform_type` (platformType) - (String) Platform type (nx-os, ios-xe)
* `recalculate` (recalculate) - (Bool) Recalculate (config-save) fabric configuration after switch operations
    * Default Value: `false`
* `deploy` (deploy) - (Bool) Deploy configuration to switches after operations
    * Default Value: `false`

## Importing

An existing Inventory Switch resource can be imported with its identifier via the following command:

```shell
terraform import nd_inventory_switch.example {id}
```

Starting in Terraform version 1.5, an existing Inventory Switch resource can be imported using import blocks via the following configuration:

```hcl
import {
  id = "{id}"
  to = nd_inventory_switch.example
}
```
