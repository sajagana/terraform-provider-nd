// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package resource_fabric_aci

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"terraform-provider-nd/internal/manage"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	nd "github.com/netascode/go-nd"
)

type fabricAciUpdateRequest struct {
	Method string
	Path   string
	Body   string
}

func fabricAciUpdateModel() FabricAciModel {
	return FabricAciModel{
		DeRegister:                 types.BoolValue(false),
		FabricName:                 types.StringValue("fabric1"),
		Hostname:                   types.StringValue("1.1.1.1"),
		Id:                         types.StringValue("fabric1"),
		Latitude:                   types.Float64Null(),
		LicenseTier:                types.StringValue("advantage"),
		LoginDomain:                types.StringValue("DefaultAuth"),
		Longitude:                  types.Float64Null(),
		OrchestrationStatus:        types.StringValue("disabled"),
		Password:                   types.StringValue("test-password"),
		ReRegister:                 types.BoolValue(false),
		SecurityDomain:             types.StringValue("all"),
		TelemetryEpg:               types.StringNull(),
		TelemetryNetwork:           types.StringNull(),
		TelemetryStatus:            types.StringValue("disabled"),
		TelemetryStreamingProtocol: types.StringNull(),
		Username:                   types.StringValue("admin"),
		VerifyCa:                   types.BoolValue(false),
	}
}

func TestFabricAciUpdateActionOrder(t *testing.T) {
	testCases := []struct {
		name             string
		modify           func(plan *FabricAciModel, state *FabricAciModel)
		expectedSummary  string
		expectedRequests []fabricAciUpdateRequest
		expectedBody     map[string]interface{}
	}{
		{
			name: "de-register only returns after the action",
			modify: func(plan *FabricAciModel, _ *FabricAciModel) {
				plan.DeRegister = types.BoolValue(true)
			},
			expectedRequests: []fabricAciUpdateRequest{
				{Method: http.MethodPost, Path: "/api/v1/infra/clusters/fabric1/deregister"},
			},
			expectedBody: map[string]interface{}{
				"credentials": map[string]interface{}{
					"user":        "admin",
					"password":    "test-password",
					"loginDomain": "DefaultAuth",
				},
			},
		},
		{
			name: "de-register with a normal update is rejected before a request",
			modify: func(plan *FabricAciModel, _ *FabricAciModel) {
				plan.DeRegister = types.BoolValue(true)
				plan.LicenseTier = types.StringValue("premier")
			},
			expectedSummary: "Invalid Fabric ACI De-registration Update",
		},
		{
			name: "resetting de-register does not send a request",
			modify: func(plan *FabricAciModel, state *FabricAciModel) {
				plan.DeRegister = types.BoolValue(false)
				state.DeRegister = types.BoolValue(true)
			},
		},
		{
			name: "re-register precedes the normal update",
			modify: func(plan *FabricAciModel, _ *FabricAciModel) {
				plan.ReRegister = types.BoolValue(true)
				plan.LicenseTier = types.StringValue("premier")
			},
			expectedRequests: []fabricAciUpdateRequest{
				{Method: http.MethodPost, Path: "/api/v1/manage/fabrics/fabric1/actions/reRegister"},
				{Method: http.MethodGet, Path: "/api/v1/infra/clusters/fabric1"},
				{Method: http.MethodPut, Path: "/api/v1/manage/fabrics/fabric1"},
				{Method: http.MethodGet, Path: "/api/v1/infra/clusters/fabric1"},
			},
		},
		{
			name: "normal update runs without an action",
			modify: func(plan *FabricAciModel, _ *FabricAciModel) {
				plan.LicenseTier = types.StringValue("premier")
			},
			expectedRequests: []fabricAciUpdateRequest{
				{Method: http.MethodPut, Path: "/api/v1/manage/fabrics/fabric1"},
				{Method: http.MethodGet, Path: "/api/v1/infra/clusters/fabric1"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			requests := make([]fabricAciUpdateRequest, 0, len(testCase.expectedRequests))
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					t.Errorf("failed to read request body: %v", err)
				}
				requests = append(requests, fabricAciUpdateRequest{
					Method: req.Method,
					Path:   req.URL.Path,
					Body:   string(body),
				})

				w.Header().Set("Content-Type", "application/json")
				if req.Method == http.MethodGet {
					_, _ = w.Write([]byte(`{
						"spec": {
							"clusterType": "APIC",
							"onboardUrl": "1.1.1.1",
							"name": "fabric1",
							"aci": {
								"licenseTier": "premier",
								"securityDomain": "all",
								"verifyCA": false,
								"telemetry": {"status": "disabled"},
								"orchestration": {"status": "disabled"}
							}
						},
						"status": {"lastUpdate": {"message": "Fabric reachable"}}
					}`))
					return
				}
				_, _ = w.Write([]byte(`{}`))
			}))
			defer server.Close()

			client, err := nd.NewClient(server.URL, "/api/v1", "", "", "", false, nd.MaxRetries(0))
			if err != nil {
				t.Fatalf("failed to create test ND client: %v", err)
			}
			client.Token = "test-token"
			client.AuthTimeStamp = time.Now()
			client.AuthTokenTimeout = time.Hour

			resource := &fabricAciResource{
				manageClient: &manage.NexusDashboardManage{ApiClient: &client},
			}
			plan := fabricAciUpdateModel()
			state := fabricAciUpdateModel()
			testCase.modify(&plan, &state)

			var diagnostics diag.Diagnostics
			resource.rscUpdateFabricAci(context.Background(), &diagnostics, &plan, &state)

			if testCase.expectedSummary == "" && diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diagnostics)
			}
			if testCase.expectedSummary != "" {
				if !diagnostics.HasError() {
					t.Fatalf("expected diagnostic %q, got none", testCase.expectedSummary)
				}
				found := false
				for _, diagnostic := range diagnostics {
					if diagnostic.Summary() == testCase.expectedSummary {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("expected diagnostic %q, got: %v", testCase.expectedSummary, diagnostics)
				}
			}

			if len(requests) != len(testCase.expectedRequests) {
				t.Fatalf("expected %d requests, got %d: %#v", len(testCase.expectedRequests), len(requests), requests)
			}
			for i, expected := range testCase.expectedRequests {
				if requests[i].Method != expected.Method || requests[i].Path != expected.Path {
					t.Fatalf(
						"request %d: expected %s %s, got %s %s",
						i+1,
						expected.Method,
						expected.Path,
						requests[i].Method,
						requests[i].Path,
					)
				}
			}

			if testCase.expectedBody != nil {
				var actualBody map[string]interface{}
				if err := json.Unmarshal([]byte(requests[0].Body), &actualBody); err != nil {
					t.Fatalf("failed to unmarshal request body: %v", err)
				}
				if !reflect.DeepEqual(actualBody, testCase.expectedBody) {
					t.Fatalf("unexpected deregistration body: got %#v, want %#v", actualBody, testCase.expectedBody)
				}
			}
		})
	}
}
