// Code generated;  DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	ActionNone = iota
	ValuesDeeplyEqual
	RequiresReplace
	RequiresUpdate
	ControlFlagUpdate
	PortListUpdate
)

type BoolOrString string

func NewBoolOrStringFromBool(value bool) *BoolOrString {
	converted := BoolOrString(strconv.FormatBool(value))
	return &converted
}

func NewBoolOrStringFromString(value string) *BoolOrString {
	converted := BoolOrString(value)
	return &converted
}

func (i BoolOrString) String() string {
	return string(i)
}

func (i BoolOrString) AsBool() (bool, error) {
	switch string(i) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("BoolOrString value is not a canonical boolean")
	}
}

func (i *BoolOrString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = BoolOrString("")
		return nil
	}
	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err == nil {
		*i = BoolOrString(strconv.FormatBool(boolValue))
		return nil
	}
	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		*i = BoolOrString(stringValue)
		return nil
	}
	return errors.New("BoolOrString requires a JSON boolean or string")
}

func (i BoolOrString) MarshalJSON() ([]byte, error) {
	switch string(i) {
	case "true":
		return []byte("true"), nil
	case "false":
		return []byte("false"), nil
	default:
		return json.Marshal(string(i))
	}
}
