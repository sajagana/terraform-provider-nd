// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"errors"
	"testing"
)

func TestNormalizeErrorResponse(t *testing.T) {
	err := errors.New("Status Code 404")
	response := []byte("{\n  \"code\": 404,\n  \"message\": \"Local user not found\"\n}")

	text, compactText := normalizeErrorResponse(err, response)

	expectedText := "status code 404 {\n  \"code\": 404,\n  \"message\": \"local user not found\"\n}"
	if text != expectedText {
		t.Fatalf("expected normalized text %q, got %q", expectedText, text)
	}

	expectedCompactText := `statuscode404{"code":404,"message":"localusernotfound"}`
	if compactText != expectedCompactText {
		t.Fatalf("expected compact text %q, got %q", expectedCompactText, compactText)
	}
}

func TestIsNotFoundError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		response []byte
		options  []NotFoundOptions
		expected bool
	}{
		{
			name:     "json code 404 with not found message",
			response: []byte(`{"code": 404, "description": "", "message": "local user not found"}`),
			expected: true,
		},
		{
			name:     "json status 404 with not found message",
			response: []byte(`{"status": 404, "message": "local user not found"}`),
			expected: true,
		},
		{
			name:     "error contains status code 404 and body contains not found",
			err:      errors.New("unexpected response: status code 404"),
			response: []byte(`{"message": "local user not found"}`),
			expected: true,
		},
		{
			name:     "error contains 404 not found",
			err:      errors.New("404 not found"),
			expected: true,
		},
		{
			name:     "custom message matches",
			response: []byte(`{"message": "local user not found"}`),
			options: []NotFoundOptions{
				{Messages: []string{"local user not found"}},
			},
			expected: true,
		},
		{
			name:     "custom status code matches",
			response: []byte(`{"code": 404, "message": "local user not found"}`),
			options: []NotFoundOptions{
				{StatusCodes: []int{404}},
			},
			expected: true,
		},
		{
			name:     "custom option disables default matching when not matched",
			response: []byte(`{"code": 404, "message": "local user not found"}`),
			options: []NotFoundOptions{
				{Messages: []string{"different not found message"}},
			},
			expected: false,
		},
		{
			name:     "404 without not found message does not match default",
			response: []byte(`{"code": 404, "message": "request failed"}`),
			expected: false,
		},
		{
			name:     "not found without 404 does not match default",
			response: []byte(`{"code": 500, "message": "local user not found"}`),
			expected: false,
		},
		{
			name:     "empty input does not match",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsNotFoundError(test.err, test.response, test.options...)
			if actual != test.expected {
				t.Fatalf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
