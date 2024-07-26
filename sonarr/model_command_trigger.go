/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.8.1874
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// CommandTrigger the model 'CommandTrigger'
type CommandTrigger string

// List of CommandTrigger
const (
	COMMANDTRIGGER_UNSPECIFIED CommandTrigger = "unspecified"
	COMMANDTRIGGER_MANUAL CommandTrigger = "manual"
	COMMANDTRIGGER_SCHEDULED CommandTrigger = "scheduled"
)

// All allowed values of CommandTrigger enum
var AllowedCommandTriggerEnumValues = []CommandTrigger{
	"unspecified",
	"manual",
	"scheduled",
}

func (v *CommandTrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommandTrigger(value)
	for _, existing := range AllowedCommandTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommandTrigger", value)
}

// NewCommandTriggerFromValue returns a pointer to a valid CommandTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommandTriggerFromValue(v string) (*CommandTrigger, error) {
	ev := CommandTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommandTrigger: valid values are %v", v, AllowedCommandTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandTrigger) IsValid() bool {
	for _, existing := range AllowedCommandTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommandTrigger value
func (v CommandTrigger) Ptr() *CommandTrigger {
	return &v
}

type NullableCommandTrigger struct {
	value *CommandTrigger
	isSet bool
}

func (v NullableCommandTrigger) Get() *CommandTrigger {
	return v.value
}

func (v *NullableCommandTrigger) Set(val *CommandTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandTrigger(val *CommandTrigger) *NullableCommandTrigger {
	return &NullableCommandTrigger{value: val, isSet: true}
}

func (v NullableCommandTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

