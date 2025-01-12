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

// TrackedDownloadState the model 'TrackedDownloadState'
type TrackedDownloadState string

// List of TrackedDownloadState
const (
	TRACKEDDOWNLOADSTATE_DOWNLOADING TrackedDownloadState = "downloading"
	TRACKEDDOWNLOADSTATE_IMPORT_BLOCKED TrackedDownloadState = "importBlocked"
	TRACKEDDOWNLOADSTATE_IMPORT_PENDING TrackedDownloadState = "importPending"
	TRACKEDDOWNLOADSTATE_IMPORTING TrackedDownloadState = "importing"
	TRACKEDDOWNLOADSTATE_IMPORTED TrackedDownloadState = "imported"
	TRACKEDDOWNLOADSTATE_FAILED_PENDING TrackedDownloadState = "failedPending"
	TRACKEDDOWNLOADSTATE_FAILED TrackedDownloadState = "failed"
	TRACKEDDOWNLOADSTATE_IGNORED TrackedDownloadState = "ignored"
)

// All allowed values of TrackedDownloadState enum
var AllowedTrackedDownloadStateEnumValues = []TrackedDownloadState{
	"downloading",
	"importBlocked",
	"importPending",
	"importing",
	"imported",
	"failedPending",
	"failed",
	"ignored",
}

func (v *TrackedDownloadState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackedDownloadState(value)
	for _, existing := range AllowedTrackedDownloadStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackedDownloadState", value)
}

// NewTrackedDownloadStateFromValue returns a pointer to a valid TrackedDownloadState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackedDownloadStateFromValue(v string) (*TrackedDownloadState, error) {
	ev := TrackedDownloadState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackedDownloadState: valid values are %v", v, AllowedTrackedDownloadStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackedDownloadState) IsValid() bool {
	for _, existing := range AllowedTrackedDownloadStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrackedDownloadState value
func (v TrackedDownloadState) Ptr() *TrackedDownloadState {
	return &v
}

type NullableTrackedDownloadState struct {
	value *TrackedDownloadState
	isSet bool
}

func (v NullableTrackedDownloadState) Get() *TrackedDownloadState {
	return v.value
}

func (v *NullableTrackedDownloadState) Set(val *TrackedDownloadState) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedDownloadState) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedDownloadState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedDownloadState(val *TrackedDownloadState) *NullableTrackedDownloadState {
	return &NullableTrackedDownloadState{value: val, isSet: true}
}

func (v NullableTrackedDownloadState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedDownloadState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

