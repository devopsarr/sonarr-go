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

// EpisodeHistoryEventType the model 'EpisodeHistoryEventType'
type EpisodeHistoryEventType string

// List of EpisodeHistoryEventType
const (
	EPISODEHISTORYEVENTTYPE_UNKNOWN EpisodeHistoryEventType = "unknown"
	EPISODEHISTORYEVENTTYPE_GRABBED EpisodeHistoryEventType = "grabbed"
	EPISODEHISTORYEVENTTYPE_SERIES_FOLDER_IMPORTED EpisodeHistoryEventType = "seriesFolderImported"
	EPISODEHISTORYEVENTTYPE_DOWNLOAD_FOLDER_IMPORTED EpisodeHistoryEventType = "downloadFolderImported"
	EPISODEHISTORYEVENTTYPE_DOWNLOAD_FAILED EpisodeHistoryEventType = "downloadFailed"
	EPISODEHISTORYEVENTTYPE_EPISODE_FILE_DELETED EpisodeHistoryEventType = "episodeFileDeleted"
	EPISODEHISTORYEVENTTYPE_EPISODE_FILE_RENAMED EpisodeHistoryEventType = "episodeFileRenamed"
	EPISODEHISTORYEVENTTYPE_DOWNLOAD_IGNORED EpisodeHistoryEventType = "downloadIgnored"
)

// All allowed values of EpisodeHistoryEventType enum
var AllowedEpisodeHistoryEventTypeEnumValues = []EpisodeHistoryEventType{
	"unknown",
	"grabbed",
	"seriesFolderImported",
	"downloadFolderImported",
	"downloadFailed",
	"episodeFileDeleted",
	"episodeFileRenamed",
	"downloadIgnored",
}

func (v *EpisodeHistoryEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EpisodeHistoryEventType(value)
	for _, existing := range AllowedEpisodeHistoryEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EpisodeHistoryEventType", value)
}

// NewEpisodeHistoryEventTypeFromValue returns a pointer to a valid EpisodeHistoryEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEpisodeHistoryEventTypeFromValue(v string) (*EpisodeHistoryEventType, error) {
	ev := EpisodeHistoryEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EpisodeHistoryEventType: valid values are %v", v, AllowedEpisodeHistoryEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EpisodeHistoryEventType) IsValid() bool {
	for _, existing := range AllowedEpisodeHistoryEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EpisodeHistoryEventType value
func (v EpisodeHistoryEventType) Ptr() *EpisodeHistoryEventType {
	return &v
}

type NullableEpisodeHistoryEventType struct {
	value *EpisodeHistoryEventType
	isSet bool
}

func (v NullableEpisodeHistoryEventType) Get() *EpisodeHistoryEventType {
	return v.value
}

func (v *NullableEpisodeHistoryEventType) Set(val *EpisodeHistoryEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeHistoryEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeHistoryEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeHistoryEventType(val *EpisodeHistoryEventType) *NullableEpisodeHistoryEventType {
	return &NullableEpisodeHistoryEventType{value: val, isSet: true}
}

func (v NullableEpisodeHistoryEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeHistoryEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

