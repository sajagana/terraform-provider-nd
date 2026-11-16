// Copyright (c) 2026 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// GenericToBool converts a generic API scalar to a Terraform Boolean value.
// Callers handle nil separately so an omitted value remains Terraform null.
func GenericToBool(value any) (bool, error) {
	switch value := value.(type) {
	case bool:
		return value, nil
	case string:
		switch value {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
	}
	return false, fmt.Errorf("expected a Boolean or the string true/false, received %T", value)
}

// GenericToString converts a generic API scalar to a Terraform string value.
// JSON numbers retain their original precision when decoded with UseNumber.
// Collections and objects require their own Terraform schema and are rejected.
func GenericToString(value any) (string, error) {
	switch value := value.(type) {
	case string:
		return value, nil
	case bool:
		return strconv.FormatBool(value), nil
	case json.Number, int, int64, float64:
		encoded, err := json.Marshal(value)
		if err != nil {
			return "", fmt.Errorf("cannot convert an invalid JSON number to string")
		}
		return string(encoded), nil
	}
	return "", fmt.Errorf("expected a string, Boolean, or number, received %T", value)
}
