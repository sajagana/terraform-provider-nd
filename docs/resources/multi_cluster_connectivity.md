---
subcategory: "infra"
layout: "nd"
page_title: "ND: nd_multi_cluster_connectivity"
sidebar_current: "docs-nd-resource-nd_multi_cluster_connectivity"
description: |-
  
---

# nd_multi_cluster_connectivity #



## API Information ##
* Multi Cluster Connectivity Management API information is not defined in the YAML.
* API Endpoint: not defined in the YAML.

## GUI Information ##
* Location: not defined in the YAML.

## Example Usage ##

The configuration snippet below shows all possible attributes of the Multi Cluster Connectivity resource.

!> This example might not be valid configuration and is only used to show all possible attributes.

Example source: `examples/resources/nd_multi_cluster_connectivity/resource.tf`

```hcl
resource "nd_multi_cluster_connectivity" "test_resource_multi_cluster_connectivity_1" {
  cluster_type               = "ND"
  cluster_name               = "nd4x"
  hostname                   = "10.15.1.111"
  username                   = "username"
  password                   = "password"
  login_domain               = "local"
  multi_cluster_login_domain = "multi_cluster_login_domain"
}
```

All examples for the Multi Cluster Connectivity resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/nd_multi_cluster_connectivity) folder.

## Schema ##

### Required ###
* `hostname` (onboardUrl) - (String) The IP address or Hostname of the ND cluster.
* `username` (user) - (Sensitive, String) The username of the ND cluster.
* `password` (password) - (Sensitive, String) The password of the ND cluster.

### Optional ###
* `cluster_name` (name) - (String) The name of the ND cluster.
* `login_domain` (loginDomain) - (Sensitive, String) The login domain of the ND cluster. The login domain is used during the connection of clusters to specify the authentication domain for the remote clusters.
* `multi_cluster_login_domain` (multiClusterLoginDomainName) - (Sensitive, String) The multi-cluster login domain is created on the primary cluster and can be configured either during initial onboarding or manually at a later time. Each primary cluster supports only one login domain of type Multi-cluster. This domain enables the primary cluster to function as the shared authentication authority for the multi-cluster group, allowing users defined on the primary cluster to remotely log in to any connected cluster.

### Read-Only ###
* `id` (name) - (String) The unique identifier of the terraform resource.
* `cluster_type` (clusterType) - (String) The type of the cluster. The value will be auto filled as "ND" for this resource.
    * Default Value: `ND`

## Importing

An existing Multi Cluster Connectivity resource can be imported with its identifier via the following command:

```shell
terraform import nd_multi_cluster_connectivity.example {id}
```

Starting in Terraform version 1.5, an existing Multi Cluster Connectivity resource can be imported using import blocks via the following configuration:

```hcl
import {
  id = "{id}"
  to = nd_multi_cluster_connectivity.example
}
```
