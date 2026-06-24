---
subcategory: "Users and Security"
layout: "nd"
page_title: "ND: nd_local_user"
sidebar_current: "docs-nd-resource-nd_local_user"
description: |-
  Manages local user for Nexus Dashboard
---

# nd_local_user #

Manages local user for Nexus Dashboard

## API Information ##
* Local User Management API information is not defined in the YAML.
* API Endpoint: not defined in the YAML.

## GUI Information ##
* Location: not defined in the YAML.

## Example Usage ##

The configuration snippet below shows all possible attributes of the Local User resource.

!> This example might not be valid configuration and is only used to show all possible attributes.

Example source: `examples/resources/nd_local_user/resource.tf`

```hcl
resource "nd_local_user" "test_resource_local_user_1" {
  login_id                  = "local_user_123"
  user_password             = "local_user_123"
  email                     = "local_user_123@mail.com"
  first_name                = "first_name"
  last_name                 = "last_name"
  remote_id_claim           = "tf_remote_id_claim"
  remote_user_authorization = false
  tenant_domain             = "all-tenants-domain"
  security_domains = {
    "all" = {
      roles = ["approver", "designer"]
    }
  }

}
```

All examples for the Local User resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/nd_local_user) folder.

## Schema ##

### Required ###
* `login_id` (loginID) - (String) The User ID of the local user.
* `user_password` (password) - (Sensitive, String) The password of the local user. User accounts are provisioned with temporary credentials and require a mandatory password reset at first login to complete onboarding.
* `security_domains` (domains) - (Map) The security domains of the local user. At least one security domain must be provided when creating the user.
    * `roles` (roles) - (List:String) The list of Nexus Dashboard Roles of the local user.
        * Valid Values: `approver`, or `designer`, or `fabric-admin`, or `observer`, or `super-admin`, or `support-engineer`.

### Optional ###
* `email` (email) - (String) The email address of the local user.
* `first_name` (firstName) - (String) The first name of the local user.
* `last_name` (lastName) - (String) The last name of the local user.
* `remote_id_claim` (remoteIDClaim) - (String) The Remote ID claim of the local user. This is required when the remote user authorization option is enabled for the local user.
* `remote_user_authorization` (xLaunch) - (Bool) The Remote user authorization is used for signing into Nexus Dashboard when using identity providers that cannot provide authorization claims. Once this attribute is enabled, the local user ID cannot be used to directly login to Nexus Dashboard.
    * Default Value: `false`
* `tenant_domain` (tenantDomain) - (String) The name of the tenant domain of the local user.

### Read-Only ###
* `id` (id) - (String) The unique identifier of the local user.

## Importing

An existing Local User resource can be imported with its identifier via the following command:

```shell
terraform import nd_local_user.example {id}
```

Starting in Terraform version 1.5, an existing Local User resource can be imported using import blocks via the following configuration:

```hcl
import {
  id = "{id}"
  to = nd_local_user.example
}
```
