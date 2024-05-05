/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.4.1491
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// CommandPriority the model 'CommandPriority'
type CommandPriority string

// List of CommandPriority
const (
	COMMANDPRIORITY_NORMAL CommandPriority = "normal"
	COMMANDPRIORITY_HIGH CommandPriority = "high"
	COMMANDPRIORITY_LOW CommandPriority = "low"
)

// All allowed values of CommandPriority enum
var AllowedCommandPriorityEnumValues = []CommandPriority{
	"normal",
	"high",
	"low",
}

func (v *CommandPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommandPriority(value)
	for _, existing := range AllowedCommandPriorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommandPriority", value)
}

// NewCommandPriorityFromValue returns a pointer to a valid CommandPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommandPriorityFromValue(v string) (*CommandPriority, error) {
	ev := CommandPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommandPriority: valid values are %v", v, AllowedCommandPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandPriority) IsValid() bool {
	for _, existing := range AllowedCommandPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommandPriority value
func (v CommandPriority) Ptr() *CommandPriority {
	return &v
}

type NullableCommandPriority struct {
	value *CommandPriority
	isSet bool
}

func (v NullableCommandPriority) Get() *CommandPriority {
	return v.value
}

func (v *NullableCommandPriority) Set(val *CommandPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandPriority(val *CommandPriority) *NullableCommandPriority {
	return &NullableCommandPriority{value: val, isSet: true}
}

func (v NullableCommandPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

