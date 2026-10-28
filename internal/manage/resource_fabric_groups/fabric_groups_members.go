// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_groups

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const fabricGroupMemberPageSize = 100

// Read-only Manage endpoints used for preflight checks. Embedding the existing
// API helper retains its CRUD locking and per-request logging/context options.
type fabricGroupCollectionAPI struct {
	*api.FabricAPI
	collection string
	offset     int
}

func (a *fabricGroupCollectionAPI) GetUrl() string {
	return fmt.Sprintf(api.UrlFabricByName, url.PathEscape(a.FabricName)) + "/" + a.collection
}

func (a *fabricGroupCollectionAPI) GetQueryParams() []string {
	return []string{"max=" + strconv.Itoa(fabricGroupMemberPageSize), "offset=" + strconv.Itoa(a.offset)}
}

func (r *fabricGroupsResource) visitFabricGroupCollection(ctx context.Context, fabricName, collection, responseKey string, visit func(json.RawMessage) error) error {
	collectionAPI := &fabricGroupCollectionAPI{
		FabricAPI:  api.NewFabricAPI(r.manageClient.ApiClient, ndapi.DefaultFabric),
		collection: collection,
	}
	collectionAPI.FabricName = fabricName
	collectionAPI.NexusDashboardAPI = collectionAPI
	collectionAPI.RequestOptions = (&FabricGroupsModel{}).FabricAPIOptions(ctx)
	seenPages := make(map[[32]byte]struct{})

	for {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("reading %s for fabric %q: %w", collection, fabricName, err)
		}
		payload, err := collectionAPI.Get()
		if err != nil {
			classified := ndapi.ClassifyRequestError(http.MethodGet, collectionAPI.GetUrl(), nil, err)
			var requestErr *ndapi.RequestError
			if errors.As(classified, &requestErr) {
				return fmt.Errorf("reading %s for fabric %q: %w", collection, fabricName, classified)
			}
			if ctx.Err() != nil {
				return fmt.Errorf("reading %s for fabric %q: %w", collection, fabricName, ctx.Err())
			}
			return fmt.Errorf("could not read %s for fabric %q; response details are omitted because they may contain secrets", collection, fabricName)
		}
		var page map[string]json.RawMessage
		if err := json.Unmarshal(payload, &page); err != nil {
			return fmt.Errorf("invalid %s response for fabric %q", collection, fabricName)
		}
		rawItems, present := page[responseKey]
		if !present {
			return fmt.Errorf("%s response for fabric %q is missing %s", collection, fabricName, responseKey)
		}
		var items []json.RawMessage
		if err := json.Unmarshal(rawItems, &items); err != nil {
			return fmt.Errorf("invalid %s list for fabric %q", responseKey, fabricName)
		}
		fingerprint := sha256.Sum256(rawItems)
		if _, repeated := seenPages[fingerprint]; repeated && len(items) > 0 {
			return fmt.Errorf("%s pagination for fabric %q repeated a page", collection, fabricName)
		}
		seenPages[fingerprint] = struct{}{}
		for _, item := range items {
			if err := visit(item); err != nil {
				return err
			}
		}

		// The members schema uses meta.remaining; network/VRF examples use
		// meta.counts.remaining and may also report meta.links.next.
		var meta struct {
			Remaining *int `json:"remaining"`
			Counts struct {
				Remaining *int `json:"remaining"`
			} `json:"counts"`
			Links struct {
				Next string `json:"next"`
			} `json:"links"`
		}
		if rawMeta, present := page["meta"]; present {
			if err := json.Unmarshal(rawMeta, &meta); err != nil {
				return fmt.Errorf("invalid %s pagination metadata for fabric %q", collection, fabricName)
			}
		}
		remaining := meta.Remaining
		if meta.Counts.Remaining != nil {
			remaining = meta.Counts.Remaining
		}
		more := len(items) == fabricGroupMemberPageSize
		if remaining != nil {
			if *remaining < 0 {
				return fmt.Errorf("invalid remaining count in %s response for fabric %q", collection, fabricName)
			}
			more = *remaining > 0
		}
		more = more || meta.Links.Next != ""
		if !more {
			return nil
		}
		if len(items) == 0 {
			return fmt.Errorf("%s pagination for fabric %q did not advance", collection, fabricName)
		}
		// Offset/max are documented on all three collection endpoints. Never
		// follow a response-provided URL to another host or service.
		collectionAPI.offset += len(items)
	}
}

func (r *fabricGroupsResource) fabricGroupMembers(ctx context.Context, name string) ([]string, error) {
	var members []string
	err := r.visitFabricGroupCollection(ctx, name, "members", "fabrics", func(item json.RawMessage) error {
		var member struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(item, &member); err != nil || member.Name == "" {
			return fmt.Errorf("invalid member identity returned for fabric group %q", name)
		}
		members = append(members, member.Name)
		return nil
	})
	return members, err
}

type fabricGroupVniConflict struct {
	attribute string
	member    string
	vni       uint64
}

func (e *fabricGroupVniConflict) Error() string {
	return fmt.Sprintf("%s overlaps VNI %d already allocated in member fabric %q", e.attribute, e.vni, e.member)
}

func (r *fabricGroupsResource) validateMemberVniOverlap(ctx context.Context, plan, previous *FabricGroupsModel, dg *diag.Diagnostics) {
	if plan == nil || previous == nil {
		dg.AddError("Invalid Input", "Member VNI validation requires planned and prior fabric group state.")
		return
	}
	if plan.FabricType.ValueString() != "vxlan" {
		return
	}
	enabled := plan.DownstreamVni
	if enabled.IsNull() || enabled.IsUnknown() {
		enabled = previous.DownstreamVni
	}
	if enabled.IsNull() || enabled.IsUnknown() || !enabled.ValueBool() {
		return
	}
	previouslyEnabled := !previous.DownstreamVni.IsNull() && !previous.DownstreamVni.IsUnknown() && previous.DownstreamVni.ValueBool()

	// Creating this resource creates an empty group: membership is managed by
	// separate API actions. Update checks actual allocated member VNIs, rather
	// than rejecting entire allocation pools that may contain unused VNIs.
	ranges := []struct {
		attribute string
		planned   types.String
		prior     types.String
		fallback  string
		bounds    [2]uint64
		oldBounds [2]uint64
	}{
		{attribute: "downstream_l2_vni_range", planned: plan.DownstreamL2VniRange, prior: previous.DownstreamL2VniRange, fallback: "60000-69000"},
		{attribute: "downstream_l3_vni_range", planned: plan.DownstreamL3VniRange, prior: previous.DownstreamL3VniRange, fallback: "80000-89000"},
	}
	checkNeeded := !previouslyEnabled
	for i := range ranges {
		value := ranges[i].planned
		if value.IsNull() || value.IsUnknown() {
			value = ranges[i].prior
		}
		text := value.ValueString()
		if value.IsNull() || value.IsUnknown() {
			// Documented API defaults, used only for validation when neither
			// configuration nor prior readback supplies an effective range.
			text = ranges[i].fallback
		}
		bounds, err := parseFabricGroupRange(text, 1, 16777214)
		if err != nil {
			dg.AddAttributeError(path.Root(ranges[i].attribute), "Invalid Effective Downstream VNI Range", err.Error())
			return
		}
		ranges[i].bounds = bounds
		priorText := ranges[i].prior.ValueString()
		if ranges[i].prior.IsNull() || ranges[i].prior.IsUnknown() {
			priorText = ranges[i].fallback
		}
		oldBounds, err := parseFabricGroupRange(priorText, 1, 16777214)
		if err != nil {
			dg.AddAttributeError(path.Root(ranges[i].attribute), "Invalid Prior Downstream VNI Range", err.Error())
			return
		}
		ranges[i].oldBounds = oldBounds
		checkNeeded = checkNeeded || bounds != oldBounds
	}
	if !checkNeeded {
		return
	}

	checkCtx, cancel := context.WithTimeout(ctx, fabricGroupReadbackTimeout)
	defer cancel()
	members, err := r.fabricGroupMembers(checkCtx, plan.FabricName.ValueString())
	if err == nil {
		for _, member := range members {
			for _, collection := range []string{"networks", "vrfs"} {
				err = r.visitFabricGroupCollection(checkCtx, member, collection, collection, func(item json.RawMessage) error {
					var allocated struct {
						NetworkID *uint64 `json:"networkId"`
						VrfID     *uint64 `json:"vrfId"`
					}
					if err := json.Unmarshal(item, &allocated); err != nil {
						return fmt.Errorf("invalid allocated VNI in %s for member fabric %q", collection, member)
					}
					id := allocated.NetworkID
					if collection == "vrfs" {
						id = allocated.VrfID
					}
					if id == nil {
						return nil
					}
					for _, allocationRange := range ranges {
						if *id >= allocationRange.bounds[0] && *id <= allocationRange.bounds[1] {
							// Existing allocations inside the same previously enabled
							// pool are not newly introduced conflicts. Validate newly
							// covered space, or the whole pool when enabling downstream VNIs.
							if previouslyEnabled && *id >= allocationRange.oldBounds[0] && *id <= allocationRange.oldBounds[1] {
								continue
							}
							return &fabricGroupVniConflict{attribute: allocationRange.attribute, member: member, vni: *id}
						}
					}
					return nil
				})
				if err != nil {
					break
				}
			}
			if err != nil {
				break
			}
		}
	}
	if err != nil {
		var conflict *fabricGroupVniConflict
		if errors.As(err, &conflict) {
			dg.AddAttributeError(path.Root(conflict.attribute), "Downstream VNI Range Overlap", conflict.Error())
			return
		}
		dg.AddError("Error Checking Fabric Group Member VNIs", err.Error())
	}
}
