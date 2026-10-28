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

// FabricJSONMarshaler optionally replaces JSON encoding after GetModelData and
// PreMarshal. Models without this hook retain the standard common-model encoder.
type FabricJSONMarshaler interface {
	MarshalFabricJSON(data *NDFCFabricCommonModel) ([]byte, error)
}

// FabricJSONUnmarshaler optionally replaces JSON decoding before PostUnmarshal
// and SetModelData, so incompatible API field types can be handled first.
type FabricJSONUnmarshaler interface {
	UnmarshalFabricJSON(payload []byte, data *NDFCFabricCommonModel) error
}

func marshalFabricJSON(model FabricModel, data *NDFCFabricCommonModel) ([]byte, error) {
	if codec, ok := model.(FabricJSONMarshaler); ok {
		return codec.MarshalFabricJSON(data)
	}
	return json.Marshal(data)
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

	fabricPayload, err := marshalFabricJSON(model, inData)
	if err != nil {
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, Data Marshall error: %v", err),
		)
		return
	}

	options := fabricRequestOptions(ctx, model)
	res, err := fabricAPI.Post(fabricPayload, options)
	if err != nil {
		if options != nil && options.DisablePayloadLog {
			dg.AddError("Error Creating Fabric", protectedFabricRequestError(http.MethodPost, fabricAPI.PostUrl(), err))
			return
		}
		dg.AddError(
			"Error Creating Fabric",
			fmt.Sprintf("Could not create fabric, unexpected error: %v: %s", err, res.String()),
		)
		return
	}
	if readback, ok := model.(FabricWriteReadback); ok {
		readback.ReadFabricAfterWrite(ctx, client, dg, inData)
		return
	}
	// ND BUG - license_tier is not set in the response. Try delaying the read
	time.Sleep(2 * time.Second)
	if !RscGetFabric(ctx, client, dg, model) {
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

// readDefault is the built-in read logic.
// Returns true if the resource was found and read successfully, false if the
// resource no longer exists on the remote infrastructure (HTTP 404). When false
// is returned, no error diagnostic is added — the caller should remove the
// resource from state so Terraform can plan a re-create.
func readDefault(ctx context.Context, client *nd.Client, dg *diag.Diagnostics, model FabricModel) bool {
	fabricName := model.GetFabricName()
	tflog.Debug(ctx, "Read fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	options := fabricRequestOptions(ctx, model)
	fabricAPI.RequestOptions = options
	respData, err := fabricAPI.Get()
	if err != nil {
		if options != nil && options.DisablePayloadLog {
			classified := ndapi.ClassifyRequestError(http.MethodGet, fabricAPI.GetUrl(), nil, err)
			if errors.Is(classified, ndapi.ErrNotFound) {
				tflog.Debug(ctx, "Fabric GET returned HTTP 404", map[string]interface{}{"fabric_name": fabricName})
				return false
			}
			dg.AddError("Error Reading Fabric", protectedFabricRequestError(http.MethodGet, fabricAPI.GetUrl(), err))
			return false
		}
		if strings.Contains(err.Error(), "StatusCode 404") {
			tflog.Warn(ctx, "Fabric not found on remote, marking for re-creation", map[string]interface{}{
				"fabric_name": fabricName,
			})
			return false
		}
		dg.AddError(
			"Error reading Fabric",
			fmt.Sprintf("Could not read fabric, unexpected error: %v: %s", err, string(respData)),
		)
		return false
	}
	var outData NDFCFabricCommonModel

	if codec, ok := model.(FabricJSONUnmarshaler); ok {
		err = codec.UnmarshalFabricJSON(respData, &outData)
	} else {
		err = json.Unmarshal(respData, &outData)
	}
	if err != nil {
		if options != nil && options.DisablePayloadLog {
			dg.AddError("Error Reading Fabric", "Could not decode the fabric response. Response details are omitted because they may contain secrets.")
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
	conversionDiagnostics := model.SetModelData(&outData)
	if options != nil {
		dg.Append(conversionDiagnostics...)
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

	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = inData.FabricName

	inDataBytes, err := marshalFabricJSON(model, inData)
	if err != nil {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, Data Marshall error: %v", err),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	options := fabricRequestOptions(ctx, model)
	res, err := fabricAPI.Put(inDataBytes, options)
	if err != nil {
		if options != nil && options.DisablePayloadLog {
			dg.AddError("Error Updating Fabric", protectedFabricRequestError(http.MethodPut, fabricAPI.PutUrl(), err))
			return
		}
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Could not update fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Updating Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
	if readback, ok := model.(FabricWriteReadback); ok {
		if !readback.ReadFabricAfterWrite(ctx, client, dg, inData) {
			return
		}
		log.Printf("Updated fabric %s with category %s", inData.FabricName, inData.Category)
		return
	}
	// Read the updated fabric
	if !RscGetFabric(ctx, client, dg, model) {
		dg.AddError(
			"Error Updating Fabric",
			fmt.Sprintf("Fabric %q was not found after update", inData.FabricName),
		)
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
	fabricAPI := api.NewFabricAPI(client, ndapi.DefaultFabric)
	fabricAPI.FabricName = fabricName
	options := fabricRequestOptions(ctx, model)
	fabricAPI.RequestOptions = options
	tflog.Debug(ctx, "Delete fabric", map[string]interface{}{
		"fabric_name": fabricName,
	})
	res, err := fabricAPI.Delete()
	if err != nil {
		if policy, ok := model.(FabricDeleteNotFound); ok && policy.DeleteMissingFabricIsSuccess() {
			classified := ndapi.ClassifyRequestError(http.MethodDelete, fabricAPI.DeleteUrl(), nil, err)
			if errors.Is(classified, ndapi.ErrNotFound) {
				tflog.Debug(ctx, "Fabric already absent during delete", map[string]interface{}{"fabric_name": fabricName})
				return
			}
		}
		if options != nil && options.DisablePayloadLog {
			dg.AddError("Error Deleting Fabric", protectedFabricRequestError(http.MethodDelete, fabricAPI.DeleteUrl(), err))
			return
		}
		dg.AddError(
			"Error Deleting Fabric",
			fmt.Sprintf("Could not delete fabric, unexpected error: %v: %s", err, res.String()),
		)
		tflog.Error(ctx, "Error Deleting Fabric", map[string]interface{}{"error": err.Error()})
		return
	}
}
