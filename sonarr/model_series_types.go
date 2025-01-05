/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.12.2823
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// SeriesTypes the model 'SeriesTypes'
type SeriesTypes string

// List of SeriesTypes
const (
	SERIESTYPES_STANDARD SeriesTypes = "standard"
	SERIESTYPES_DAILY SeriesTypes = "daily"
	SERIESTYPES_ANIME SeriesTypes = "anime"
)

// All allowed values of SeriesTypes enum
var AllowedSeriesTypesEnumValues = []SeriesTypes{
	"standard",
	"daily",
	"anime",
}

func (v *SeriesTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeriesTypes(value)
	for _, existing := range AllowedSeriesTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeriesTypes", value)
}

// NewSeriesTypesFromValue returns a pointer to a valid SeriesTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeriesTypesFromValue(v string) (*SeriesTypes, error) {
	ev := SeriesTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeriesTypes: valid values are %v", v, AllowedSeriesTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeriesTypes) IsValid() bool {
	for _, existing := range AllowedSeriesTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeriesTypes value
func (v SeriesTypes) Ptr() *SeriesTypes {
	return &v
}

type NullableSeriesTypes struct {
	value *SeriesTypes
	isSet bool
}

func (v NullableSeriesTypes) Get() *SeriesTypes {
	return v.value
}

func (v *NullableSeriesTypes) Set(val *SeriesTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTypes(val *SeriesTypes) *NullableSeriesTypes {
	return &NullableSeriesTypes{value: val, isSet: true}
}

func (v NullableSeriesTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

