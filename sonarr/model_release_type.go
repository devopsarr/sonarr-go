/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.13.2932
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// ReleaseType the model 'ReleaseType'
type ReleaseType string

// List of ReleaseType
const (
	RELEASETYPE_UNKNOWN ReleaseType = "unknown"
	RELEASETYPE_SINGLE_EPISODE ReleaseType = "singleEpisode"
	RELEASETYPE_MULTI_EPISODE ReleaseType = "multiEpisode"
	RELEASETYPE_SEASON_PACK ReleaseType = "seasonPack"
)

// All allowed values of ReleaseType enum
var AllowedReleaseTypeEnumValues = []ReleaseType{
	"unknown",
	"singleEpisode",
	"multiEpisode",
	"seasonPack",
}

func (v *ReleaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReleaseType(value)
	for _, existing := range AllowedReleaseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReleaseType", value)
}

// NewReleaseTypeFromValue returns a pointer to a valid ReleaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReleaseTypeFromValue(v string) (*ReleaseType, error) {
	ev := ReleaseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReleaseType: valid values are %v", v, AllowedReleaseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReleaseType) IsValid() bool {
	for _, existing := range AllowedReleaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReleaseType value
func (v ReleaseType) Ptr() *ReleaseType {
	return &v
}

type NullableReleaseType struct {
	value *ReleaseType
	isSet bool
}

func (v NullableReleaseType) Get() *ReleaseType {
	return v.value
}

func (v *NullableReleaseType) Set(val *ReleaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseType(val *ReleaseType) *NullableReleaseType {
	return &NullableReleaseType{value: val, isSet: true}
}

func (v NullableReleaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

