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

// SeriesStatusType the model 'SeriesStatusType'
type SeriesStatusType string

// List of SeriesStatusType
const (
	SERIESSTATUSTYPE_CONTINUING SeriesStatusType = "continuing"
	SERIESSTATUSTYPE_ENDED SeriesStatusType = "ended"
	SERIESSTATUSTYPE_UPCOMING SeriesStatusType = "upcoming"
	SERIESSTATUSTYPE_DELETED SeriesStatusType = "deleted"
)

// All allowed values of SeriesStatusType enum
var AllowedSeriesStatusTypeEnumValues = []SeriesStatusType{
	"continuing",
	"ended",
	"upcoming",
	"deleted",
}

func (v *SeriesStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeriesStatusType(value)
	for _, existing := range AllowedSeriesStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeriesStatusType", value)
}

// NewSeriesStatusTypeFromValue returns a pointer to a valid SeriesStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeriesStatusTypeFromValue(v string) (*SeriesStatusType, error) {
	ev := SeriesStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeriesStatusType: valid values are %v", v, AllowedSeriesStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeriesStatusType) IsValid() bool {
	for _, existing := range AllowedSeriesStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeriesStatusType value
func (v SeriesStatusType) Ptr() *SeriesStatusType {
	return &v
}

type NullableSeriesStatusType struct {
	value *SeriesStatusType
	isSet bool
}

func (v NullableSeriesStatusType) Get() *SeriesStatusType {
	return v.value
}

func (v *NullableSeriesStatusType) Set(val *SeriesStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesStatusType(val *SeriesStatusType) *NullableSeriesStatusType {
	return &NullableSeriesStatusType{value: val, isSet: true}
}

func (v NullableSeriesStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

