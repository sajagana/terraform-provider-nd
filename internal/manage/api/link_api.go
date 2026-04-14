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

// List the links
// https://10.15.0.125/api/v1/manage/links?fabricName=test_sab_fabric

// Get a specific link
// https://10.15.0.125/api/v1/manage/links/LINK-UUID-19000

// LinkAPI is the API client for link resources
const (
	// Standard Link API endpoints
	UrlLinks         = "/manage/links"
	UrlLinksByFabric = "/manage/links?fabricName=%s"
	UrlLinkById      = "/manage/links/%s"
)

type LinkAPI struct {
	NDManageAPICommon
	mutex      *sync.Mutex
	LinkId     string
	FabricName string
}

func NewLinkAPI(lock *sync.Mutex, client *nd.Client) *LinkAPI {
	papi := new(LinkAPI)
	papi.mutex = lock
	papi.Client = client
	papi.NDManageAPI = papi
	return papi
}

func (c *LinkAPI) GetLock() *sync.Mutex {
	return c.mutex
}

func (c *LinkAPI) GetUrl() string {
	if c.LinkId != "" {
		// Get a specific link
		return fmt.Sprintf(UrlLinkById, c.LinkId)
	}
	// Get all links
	return fmt.Sprintf(UrlLinksByFabric, c.FabricName)
}

func (c *LinkAPI) PostUrl() string {
	return UrlLinks
}

func (c *LinkAPI) PutUrl() string {
	return fmt.Sprintf(UrlLinkById, c.LinkId)
}

func (c *LinkAPI) DeleteUrl() string {
	return fmt.Sprintf(UrlLinkById, c.LinkId)
}

func (c *LinkAPI) GetDeleteQP() []string {
	return nil
}

func (c *LinkAPI) RscName() string {
	return "link"
}
