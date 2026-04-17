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
	"sync"

	"github.com/netascode/go-nd"
)

// Onboard ND Cluster
// https://10.15.0.125/api/v1/infra/clusters

// {
//     "spec": {
//         "clusterType": "ND",
//         "onboardUrl": "10.15.0.128",
//         "credentials": {
//             "user": "admin",
//             "password": "cisco.123",
//             "logindomain": "local"
//         },
//         "nd": {
//             "multiClusterLoginDomainName": "test_sab_nd_mcl_domain"
//         }
//     }
// }

// Onboard ACI Cluster
// https://10.15.0.125/api/v1/infra/clusters

// {
//     "spec": {
//         "clusterType": "APIC",
//         "onboardUrl": "10.15.38.241",
//         "location": {
//             "latitude": 37.3,
//             "longitude": -121.8
//         },
//         "credentials": {
//             "user": "admin",
//             "password": "cisco.123",
//             "logindomain": "test_sab_aci_login_domain"
//         },
//         "aci": {
//             "name": "apic1",
//             "licenseTier": "essentials",
//             "telemetry": {
//                 "epg": "epg_dn",
//                 "network": ["inband", "outOfBand"],
//                 "streamingProtocol": [ "ipv4", "ipv6" ]
//                 "status": "disabled"
//             },
//             "orchestration": { "status": "disabled" },
//             "securityDomain": "test_sab_nd_security_domain",
//             "verifyCA": true
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

// PostDelete sends an HTTP POST to the cluster remove endpoint with the given payload.
// The /remove endpoint requires POST with credentials and force flag in the body.
func (c *ClusterAPI) PostDelete(payload []byte) (nd.Res, error) {
	lock := c.GetLock()
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}

	res, err := c.Client.Post(fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *ClusterAPI) RscName() string {
	return "multi_cluster_connectivity"
}
