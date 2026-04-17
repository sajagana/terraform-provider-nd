// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"
	"log"
	"sync"

	"github.com/netascode/go-nd"
)

// api/v1/infra/clusters
// https://10.15.0.125/api/v1/infra/clusters

// {
//     "spec": {
//         "clusterType": "ND",
//         "onboardUrl": "10.15.0.128",
//         "credentials": {
//             "user": "admin",
//             "password": "cisco.123",
//             "logindomain": "local"
//         }
//     }
// }

// Cluster API endpoints
const (
	UrlCluster             = "/infra/clusters"
	UrlClusterByName       = "/infra/clusters/%s"
	UrlClusterRemoveByName = "/infra/clusters/%s/remove"
)

// ClusterRemoveCredentials represents credentials for the cluster remove payload
type ClusterRemoveCredentials struct {
	LoginDomain string `json:"loginDomain,omitempty"`
	Password    string `json:"password,omitempty"`
	User        string `json:"user,omitempty"`
}

// ClusterRemovePayload represents the JSON body for the POST-based cluster remove endpoint
type ClusterRemovePayload struct {
	Credentials ClusterRemoveCredentials `json:"credentials"`
	Force       bool                     `json:"force"`
}

type ClusterAPI struct {
	NDInfraAPICommon
	mutex       *sync.Mutex
	ClusterName string
}

func NewClusterAPI(lock *sync.Mutex, client *nd.Client) *ClusterAPI {
	papi := new(ClusterAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDInfraAPI = papi
	return papi
}

func (c *ClusterAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *ClusterAPI) GetUrl() string {
	if c.ClusterName != "" {
		// Get a specific cluster
		return fmt.Sprintf(UrlClusterByName, c.ClusterName)
	}
	// Get all clusters
	return UrlCluster
}

// https://10.15.0.125/api/v1/infra/clusters

//	{
//	    "spec": {
//	        "clusterType": "APIC",
//	        "onboardUrl": "10.15.38.241",
//	        "location": {
//	            "latitude": 37.3,
//	            "longitude": -121.8
//	        },
//	        "credentials": {
//	            "user": "admin",
//	            "password": "cisco.123",
//	            "logindomain": ""
//	        },
//	        "aci": {
//	            "name": "apic1",
//	            "licenseTier": "essentials",
//	            "telemetry": {
//	                "status": "disabled"
//	            },
//	            "securityDomain": "",
//	            "verifyCA": false
//	        }
//	    }
//	}
func (c *ClusterAPI) PostUrl() string {
	return UrlCluster
}

func (c *ClusterAPI) PutUrl() string {
	return fmt.Sprintf(UrlClusterByName, c.ClusterName)
}

// https://10.15.0.125/api/v1/infra/clusters/{cluster-name}/remove
//
//	{
//	  "credentials": {
//	    "loginDomain": "DefaultAuth",
//	    "password": "Q2lzY29pbnMzOTY1IQ==",
//	    "user": "admin"
//	  },
//	  "force": true
//	}

// https://10.15.0.125/api/v1/infra/clusters/{cluster-name}/remove
//
// {
//     "credentials": {
//         "user": "admin",
//         "password": "cisco.123",
//         "loginDomain": ""
//     }
// }

func (c *ClusterAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName)
}

// PostDelete sends an HTTP POST to the cluster remove endpoint with the given payload.
// The /remove endpoint requires POST with credentials and force flag in the body.
func (c *ClusterAPI) PostDelete(payload []byte) (nd.Res, error) {
	lock := c.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	url := c.DeleteUrl()
	log.Printf("PostDelete URL: %s\n", url)
	res, err := c.Client.Post(url, string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

// func (c *ClusterAPI) GetDeleteQP() []string {
// 	return nil
// }

func (c *ClusterAPI) RscName() string {
	return "multi_cluster_connectivity"
}

// 2026-04-20T21:06:47.088+0530 [ERROR] vertex "nd_multi_cluster_connectivity.test_resource_multi_cluster_connectivity_1" error: Provider produced inconsistent result after apply
// 2026-04-20T21:06:47.088+0530 [ERROR] vertex "nd_multi_cluster_connectivity.test_resource_multi_cluster_connectivity_1" error: Provider produced inconsistent result after apply
// ╷
// │ Error: Provider produced inconsistent result after apply
// │
// │ When applying changes to nd_multi_cluster_connectivity.test_resource_multi_cluster_connectivity_1, provider "provider[\"registry.terraform.io/ciscodevnet/nd\"]" produced an unexpected new value: .username: was
// │ cty.StringVal("admin"), but now null.
// │
// │ This is a bug in the provider, which should be reported in the provider's own issue tracker.
// ╵
// ╷
// │ Error: Provider produced inconsistent result after apply
// │
// │ When applying changes to nd_multi_cluster_connectivity.test_resource_multi_cluster_connectivity_1, provider "provider[\"registry.terraform.io/ciscodevnet/nd\"]" produced an unexpected new value: .password: was
// │ cty.StringVal("cisco.123"), but now null.
// │
// │ This is a bug in the provider, which should be reported in the provider's own issue tracker.
