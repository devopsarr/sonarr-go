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

// BackupType the model 'BackupType'
type BackupType string

// List of BackupType
const (
	BACKUPTYPE_SCHEDULED BackupType = "scheduled"
	BACKUPTYPE_MANUAL BackupType = "manual"
	BACKUPTYPE_UPDATE BackupType = "update"
)

// All allowed values of BackupType enum
var AllowedBackupTypeEnumValues = []BackupType{
	"scheduled",
	"manual",
	"update",
}

func (v *BackupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupType(value)
	for _, existing := range AllowedBackupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupType", value)
}

// NewBackupTypeFromValue returns a pointer to a valid BackupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupTypeFromValue(v string) (*BackupType, error) {
	ev := BackupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupType: valid values are %v", v, AllowedBackupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupType) IsValid() bool {
	for _, existing := range AllowedBackupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupType value
func (v BackupType) Ptr() *BackupType {
	return &v
}

type NullableBackupType struct {
	value *BackupType
	isSet bool
}

func (v NullableBackupType) Get() *BackupType {
	return v.value
}

func (v *NullableBackupType) Set(val *BackupType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupType(val *BackupType) *NullableBackupType {
	return &NullableBackupType{value: val, isSet: true}
}

func (v NullableBackupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

