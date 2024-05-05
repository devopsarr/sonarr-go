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

// EpisodeTitleRequiredType the model 'EpisodeTitleRequiredType'
type EpisodeTitleRequiredType string

// List of EpisodeTitleRequiredType
const (
	EPISODETITLEREQUIREDTYPE_ALWAYS EpisodeTitleRequiredType = "always"
	EPISODETITLEREQUIREDTYPE_BULK_SEASON_RELEASES EpisodeTitleRequiredType = "bulkSeasonReleases"
	EPISODETITLEREQUIREDTYPE_NEVER EpisodeTitleRequiredType = "never"
)

// All allowed values of EpisodeTitleRequiredType enum
var AllowedEpisodeTitleRequiredTypeEnumValues = []EpisodeTitleRequiredType{
	"always",
	"bulkSeasonReleases",
	"never",
}

func (v *EpisodeTitleRequiredType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EpisodeTitleRequiredType(value)
	for _, existing := range AllowedEpisodeTitleRequiredTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EpisodeTitleRequiredType", value)
}

// NewEpisodeTitleRequiredTypeFromValue returns a pointer to a valid EpisodeTitleRequiredType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEpisodeTitleRequiredTypeFromValue(v string) (*EpisodeTitleRequiredType, error) {
	ev := EpisodeTitleRequiredType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EpisodeTitleRequiredType: valid values are %v", v, AllowedEpisodeTitleRequiredTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EpisodeTitleRequiredType) IsValid() bool {
	for _, existing := range AllowedEpisodeTitleRequiredTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EpisodeTitleRequiredType value
func (v EpisodeTitleRequiredType) Ptr() *EpisodeTitleRequiredType {
	return &v
}

type NullableEpisodeTitleRequiredType struct {
	value *EpisodeTitleRequiredType
	isSet bool
}

func (v NullableEpisodeTitleRequiredType) Get() *EpisodeTitleRequiredType {
	return v.value
}

func (v *NullableEpisodeTitleRequiredType) Set(val *EpisodeTitleRequiredType) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeTitleRequiredType) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeTitleRequiredType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeTitleRequiredType(val *EpisodeTitleRequiredType) *NullableEpisodeTitleRequiredType {
	return &NullableEpisodeTitleRequiredType{value: val, isSet: true}
}

func (v NullableEpisodeTitleRequiredType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeTitleRequiredType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

