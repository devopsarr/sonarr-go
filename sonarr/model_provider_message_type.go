/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.11.2680
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// ProviderMessageType the model 'ProviderMessageType'
type ProviderMessageType string

// List of ProviderMessageType
const (
	PROVIDERMESSAGETYPE_INFO ProviderMessageType = "info"
	PROVIDERMESSAGETYPE_WARNING ProviderMessageType = "warning"
	PROVIDERMESSAGETYPE_ERROR ProviderMessageType = "error"
)

// All allowed values of ProviderMessageType enum
var AllowedProviderMessageTypeEnumValues = []ProviderMessageType{
	"info",
	"warning",
	"error",
}

func (v *ProviderMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderMessageType(value)
	for _, existing := range AllowedProviderMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderMessageType", value)
}

// NewProviderMessageTypeFromValue returns a pointer to a valid ProviderMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderMessageTypeFromValue(v string) (*ProviderMessageType, error) {
	ev := ProviderMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderMessageType: valid values are %v", v, AllowedProviderMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderMessageType) IsValid() bool {
	for _, existing := range AllowedProviderMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderMessageType value
func (v ProviderMessageType) Ptr() *ProviderMessageType {
	return &v
}

type NullableProviderMessageType struct {
	value *ProviderMessageType
	isSet bool
}

func (v NullableProviderMessageType) Get() *ProviderMessageType {
	return v.value
}

func (v *NullableProviderMessageType) Set(val *ProviderMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderMessageType(val *ProviderMessageType) *NullableProviderMessageType {
	return &NullableProviderMessageType{value: val, isSet: true}
}

func (v NullableProviderMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

