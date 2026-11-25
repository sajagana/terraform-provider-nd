// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"terraform-provider-nd/internal/common/ndapi"
	"terraform-provider-nd/internal/manage/api"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nd"
)

// FabricModel is the interface that all fabric type models must implement.
// It bridges the per-type Terraform models and the shared CRUD backend.
type FabricModel interface {
	GetModelData() *NDFCFabricCommonModel
	SetModelData(*NDFCFabricCommonModel) diag.Diagnostics
	GetFabricName() string
	GetFabricType() string
}

// FabricPreMarshal is an optional interface. Implement it on a fabric model
// to adjust the common model data before it is marshalled for POST/PUT.
// Use cases: injecting fabric type, removing irrelevant fields, setting defaults.
type FabricPreMarshal interface {
	PreMarshal(ctx context.Context, data *NDFCFabricCommonModel)
}

// FabricPostUnmarshal is an optional interface. Implement it on a fabric model
// to adjust the common model data after GET unmarshal, before SetModelData.
// Use cases: normalizing API response quirks, filtering read-only noise.
type FabricPostUnmarshal interface {
	PostUnmarshal(ctx context.Context, data *NDFCFabricCommonModel)
}

// fabricAPIOptions allows a fabric model to opt into shared request options.
type fabricAPIOptions interface {
	FabricAPIOptions() *ndapi.APIOptions
}

// fabricPayloadMarshaler projects the prepared common model into a type-specific body.
type fabricPayloadMarshaler interface {
	MarshalFabricPayload(*NDFCFabricCommonModel) ([]byte, error)
}

// fabricCreateObserver records a successful POST before the follow-up GET so
// the resource can return its Terraform ID if that GET fails.
type fabricCreateObserver interface {
	OnFabricCreateAccepted()
}

// fabricRequestOptions retrieves the resource's opt-in settings, such as
// suppressing HTTP body logs. A nil option set keeps the default behavior.
func fabricRequestOptions(model FabricModel) *ndapi.APIOptions {
	if configured, ok := model.(fabricAPIOptions); ok {
		return configured.FabricAPIOptions()
	}
	return nil
}

// marshalFabricPayload lets a resource shape the prepared common model for its
// Create and Update requests. By default, it encodes the full common model.
func marshalFabricPayload(model FabricModel, data *NDFCFabricCommonModel) ([]byte, error) {
	if marshaler, ok := model.(fabricPayloadMarshaler); ok {
		return marshaler.MarshalFabricPayload(data)
	}
	return json.Marshal(data)
}

// namedFabricAPI sets up the API for reading, updating, or deleting a fabric
// by name. For opt-in resources, it rejects blank names and escapes the name
// for use as a URL path segment.
func namedFabricAPI(client *nd.Client, name string, opts *ndapi.APIOptions, dg *diag.Diagnostics) *api.FabricAPI {
	if opts != nil {
		if strings.TrimSpace(name) == "" {
			dg.AddError("Invalid Fabric Name", "A non-empty fabric name is required.")
			return nil
		}
		name = url.PathEscape(name)
	}
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = name
	return fabricAPI
}

// fabricRequestError formats the request failure for a Terraform diagnostic.
// When payload logging is disabled, it omits the raw response body, which may
// contain sensitive data.
func fabricRequestError(err error, response string, opts *ndapi.APIOptions) string {
	if opts != nil && opts.DisablePayloadLog {
		return err.Error()
	}
	return fmt.Sprintf("%v: %s", err, response)
}

// RscCreateFabric creates a fabric resource via the NDFC API.
func RscCreateFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Create(ctx, client, dg, model)
}

// createDefault is the built-in create logic.
func createDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	inData := model.GetModelData()
	if pm, ok := model.(FabricPreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}

	log.Printf("Creating fabric %s with category %s", inData.FabricName, inData.Category)

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)

	fabricPayload, err := marshalFabricPayload(model, inData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, Data Marshall error: %v", err),
		)
		return
	}

	opts := fabricRequestOptions(model)
	res, err := fabricAPI.Post(fabricPayload, opts)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, unexpected error: %s", fabricRequestError(err, res.String(), opts)),
		)
		return
	}
	if observer, ok := model.(fabricCreateObserver); ok {
		observer.OnFabricCreateAccepted()
	}
	// ND BUG - license_tier is not set in the response. Try delaying the read
	time.Sleep(2 * time.Second)
	if !RscGetFabric(ctx, client, dg, model) && (opts == nil || !dg.HasError()) {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Fabric %q was not found after creation", model.GetFabricName()),
		)
	}
}

// RscGetFabric retrieves fabric information by name and populates the model.
// Returns true if the resource was found, false if it no longer exists (404).
func RscGetFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	return fabricHandler(model.GetFabricType()).Read(ctx, client, dg, model)
}

// RscValidateFabricType checks that a retrieved fabric belongs to one of the
// resource types supported by an import operation.
func RscValidateFabricType(dg *diag.Diagnostics, model FabricModel, allowedTypes ...string) bool {
	actualType := model.GetModelData().Management.FabricType
	for _, allowedType := range allowedTypes {
		if actualType == allowedType {
			return true
		}
	}

	dg.AddError(
		"Error Importing Fabric",
		fmt.Sprintf("Fabric %q has type %q; expected one of %v", model.GetFabricName(), actualType, allowedTypes),
	)
	return false
}

// readDefault is the built-in read logic.
// Returns true if the resource was found and read successfully. Callers must
// check error diagnostics before treating false as a missing resource (HTTP 404)
// and removing it from state.
func readDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	fabricName := model.GetFabricName()
	tflog.Debug(ctx, "Read fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	opts := fabricRequestOptions(model)
	fabricAPI := namedFabricAPI(client, fabricName, opts, dg)
	if fabricAPI == nil {
		return false
	}
	respData, err := fabricAPI.GetWithOptions(opts)
	if err != nil {
		if opts != nil && errors.Is(ndapi.ClassifyRequestError(http.MethodGet, fabricAPI.GetUrl(), nil, err), ndapi.ErrNotFound) {
			return false
		}
		if opts == nil && strings.Contains(err.Error(), "StatusCode 404") {
			tflog.Warn(ctx, "Fabric not found on remote, marking for re-creation", map[string]interface{}{
				"fabric_name": fabricName,
			})
			return false
		}
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %s", fabricRequestError(err, string(respData), opts)),
		)
		return false
	}
	var outData NDFCFabricCommonModel

	err = json.Unmarshal(respData, &outData)
	if err != nil {
		if opts != nil && opts.DisablePayloadLog {
			dg.AddError("Error reading Fabric", fmt.Sprintf("Could not decode the API response for fabric %q.", fabricName))
			return false
		}
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v %v", err, respData),
		)
		return false
	}
	outData.Id = fabricName
	if pu, ok := model.(FabricPostUnmarshal); ok {
		pu.PostUnmarshal(ctx, &outData)
	}
	modelDiags := model.SetModelData(&outData)
	// Keep the existing read behavior for fabrics that do not opt in.
	if opts != nil {
		dg.Append(modelDiags...)
		return !dg.HasError()
	}
	return true
}

// RscUpdateFabric updates a fabric resource via the NDFC API.
func RscUpdateFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Update(ctx, client, dg, model)
}

// updateDefault is the built-in update logic.
func updateDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	inData := model.GetModelData()
	if pm, ok := model.(FabricPreMarshal); ok {
		pm.PreMarshal(ctx, inData)
	}
	tflog.Debug(ctx, "Update fabric", map[string]interface{}{
		"fabric_name": inData.FabricName,
		"category":    inData.Category,
	})

	opts := fabricRequestOptions(model)
	fabricAPI := namedFabricAPI(client, inData.FabricName, opts, dg)
	if fabricAPI == nil {
		return
	}

	inDataBytes, err := marshalFabricPayload(model, inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	res, err := fabricAPI.Put(inDataBytes, opts)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, unexpected error: %s", fabricRequestError(err, res.String(), opts)),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	// Read the updated fabric
	if !RscGetFabric(ctx, client, dg, model) {
		if opts == nil || !dg.HasError() {
			dg.AddError(
				"Error Updating Fabric",
				fmt.Sprintf("Fabric %q was not found after update", inData.FabricName),
			)
		}
		return
	}
	log.Printf("Updated fabric %s with category %s", inData.FabricName, inData.Category)
}

// RscDeleteFabric deletes a fabric resource via the NDFC API.
func RscDeleteFabric(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricHandler(model.GetFabricType()).Delete(ctx, client, dg, model)
}

// deleteDefault is the built-in delete logic.
func deleteDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) {
	fabricName := model.GetFabricName()
	opts := fabricRequestOptions(model)
	fabricAPI := namedFabricAPI(client, fabricName, opts, dg)
	if fabricAPI == nil {
		return
	}
	tflog.Debug(ctx, "Delete fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	res, err := fabricAPI.Delete()
	if err != nil {
		if opts != nil && errors.Is(ndapi.ClassifyRequestError(http.MethodDelete, fabricAPI.DeleteUrl(), nil, err), ndapi.ErrNotFound) {
			return
		}
		dg.AddError(
			"Error Deleting Fabric",
			fmt.Sprintf("Could not delete fabric, unexpected error: %s", fabricRequestError(err, res.String(), opts)),
		)
		tflog.Error(ctx, "Error Deleting Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
}
