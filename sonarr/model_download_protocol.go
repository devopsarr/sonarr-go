/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.9.2244
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"fmt"
)

// DownloadProtocol the model 'DownloadProtocol'
type DownloadProtocol string

// List of DownloadProtocol
const (
	DOWNLOADPROTOCOL_UNKNOWN DownloadProtocol = "unknown"
	DOWNLOADPROTOCOL_USENET DownloadProtocol = "usenet"
	DOWNLOADPROTOCOL_TORRENT DownloadProtocol = "torrent"
)

// All allowed values of DownloadProtocol enum
var AllowedDownloadProtocolEnumValues = []DownloadProtocol{
	"unknown",
	"usenet",
	"torrent",
}

func (v *DownloadProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DownloadProtocol(value)
	for _, existing := range AllowedDownloadProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DownloadProtocol", value)
}

// NewDownloadProtocolFromValue returns a pointer to a valid DownloadProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDownloadProtocolFromValue(v string) (*DownloadProtocol, error) {
	ev := DownloadProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DownloadProtocol: valid values are %v", v, AllowedDownloadProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DownloadProtocol) IsValid() bool {
	for _, existing := range AllowedDownloadProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DownloadProtocol value
func (v DownloadProtocol) Ptr() *DownloadProtocol {
	return &v
}

type NullableDownloadProtocol struct {
	value *DownloadProtocol
	isSet bool
}

func (v NullableDownloadProtocol) Get() *DownloadProtocol {
	return v.value
}

func (v *NullableDownloadProtocol) Set(val *DownloadProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadProtocol(val *DownloadProtocol) *NullableDownloadProtocol {
	return &NullableDownloadProtocol{value: val, isSet: true}
}

func (v NullableDownloadProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

