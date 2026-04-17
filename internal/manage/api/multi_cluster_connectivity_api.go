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
	UrlCluster             = " /clusters"
	UrlClusterByName       = " /clusters/%s"
	UrlClusterRemoveByName = " /clusters/%s/remove"
)

type ClusterAPI struct {
	NDManageAPICommon
	mutex       *sync.Mutex
	ClusterName string
}

func NewClusterAPI(lock *sync.Mutex, client *nd.Client) *ClusterAPI {
	papi := new(ClusterAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDManageAPI = papi
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

func (c *ClusterAPI) PostUrl() string {
	return UrlCluster
}

func (c *ClusterAPI) PutUrl() string {
	return fmt.Sprintf(UrlClusterByName, c.ClusterName)
}

func (c *ClusterAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName)
}

// func (c *ClusterAPI) GetDeleteQP() []string {
// 	return nil
// }

func (c *ClusterAPI) RscName() string {
	return "multi_cluster_connectivity"
}
