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

// ClusterACIAPI provides API access for onboarding ACI (APIC) clusters.
// It uses the same /infra/clusters endpoints as ClusterAPI but is a distinct
// type so the multi_cluster_connectivity_aci resource can be wired separately.
type ClusterACIAPI struct {
	NDInfraAPICommon
	mutex       *sync.Mutex
	ClusterName string
}

func NewClusterACIAPI(lock *sync.Mutex, client *nd.Client) *ClusterACIAPI {
	papi := new(ClusterACIAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDInfraAPI = papi
	return papi
}

func (c *ClusterACIAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *ClusterACIAPI) GetUrl() string {
	if c.ClusterName != "" {
		return fmt.Sprintf(UrlClusterByName, c.ClusterName)
	}
	return UrlCluster
}

func (c *ClusterACIAPI) PostUrl() string {
	return UrlCluster
}

func (c *ClusterACIAPI) PutUrl() string {
	return fmt.Sprintf(UrlClusterByName, c.ClusterName)
}

func (c *ClusterACIAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlClusterRemoveByName, c.ClusterName)
}

// Multi Cluster Delete API does not support query params so GetDeleteQP is not implemented for now, but keeping the logic in place in case it's needed in the future
func (c *ClusterACIAPI) GetDeleteQP() []string {
	return nil
}

func (c *ClusterACIAPI) PostDelete(payload []byte) (nd.Res, error) {
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

func (c *ClusterACIAPI) RscName() string {
	return "multi_cluster_connectivity_aci"
}
