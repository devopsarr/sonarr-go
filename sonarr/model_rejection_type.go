/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// RejectionType the model 'RejectionType'
type RejectionType string

// List of RejectionType
const (
	REJECTIONTYPE_PERMANENT RejectionType = "permanent"
	REJECTIONTYPE_TEMPORARY RejectionType = "temporary"
)

// All allowed values of RejectionType enum
var AllowedRejectionTypeEnumValues = []RejectionType{
	"permanent",
	"temporary",
}

func (v *RejectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RejectionType(value)
	for _, existing := range AllowedRejectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RejectionType", value)
}

// NewRejectionTypeFromValue returns a pointer to a valid RejectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRejectionTypeFromValue(v string) (*RejectionType, error) {
	ev := RejectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RejectionType: valid values are %v", v, AllowedRejectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RejectionType) IsValid() bool {
	for _, existing := range AllowedRejectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RejectionType value
func (v RejectionType) Ptr() *RejectionType {
	return &v
}

type NullableRejectionType struct {
	value *RejectionType
	isSet bool
}

func (v NullableRejectionType) Get() *RejectionType {
	return v.value
}

func (v *NullableRejectionType) Set(val *RejectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectionType(val *RejectionType) *NullableRejectionType {
	return &NullableRejectionType{value: val, isSet: true}
}

func (v NullableRejectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

