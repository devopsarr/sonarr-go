/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.2.1183
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// CommandResult the model 'CommandResult'
type CommandResult string

// List of CommandResult
const (
	COMMANDRESULT_UNKNOWN CommandResult = "unknown"
	COMMANDRESULT_SUCCESSFUL CommandResult = "successful"
	COMMANDRESULT_UNSUCCESSFUL CommandResult = "unsuccessful"
)

// All allowed values of CommandResult enum
var AllowedCommandResultEnumValues = []CommandResult{
	"unknown",
	"successful",
	"unsuccessful",
}

func (v *CommandResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommandResult(value)
	for _, existing := range AllowedCommandResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommandResult", value)
}

// NewCommandResultFromValue returns a pointer to a valid CommandResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommandResultFromValue(v string) (*CommandResult, error) {
	ev := CommandResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommandResult: valid values are %v", v, AllowedCommandResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandResult) IsValid() bool {
	for _, existing := range AllowedCommandResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommandResult value
func (v CommandResult) Ptr() *CommandResult {
	return &v
}

type NullableCommandResult struct {
	value *CommandResult
	isSet bool
}

func (v NullableCommandResult) Get() *CommandResult {
	return v.value
}

func (v *NullableCommandResult) Set(val *CommandResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResult(val *CommandResult) *NullableCommandResult {
	return &NullableCommandResult{value: val, isSet: true}
}

func (v NullableCommandResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

