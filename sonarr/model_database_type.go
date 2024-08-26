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

// DatabaseType the model 'DatabaseType'
type DatabaseType string

// List of DatabaseType
const (
	DATABASETYPE_SQ_LITE DatabaseType = "sqLite"
	DATABASETYPE_POSTGRE_SQL DatabaseType = "postgreSQL"
)

// All allowed values of DatabaseType enum
var AllowedDatabaseTypeEnumValues = []DatabaseType{
	"sqLite",
	"postgreSQL",
}

func (v *DatabaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatabaseType(value)
	for _, existing := range AllowedDatabaseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatabaseType", value)
}

// NewDatabaseTypeFromValue returns a pointer to a valid DatabaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatabaseTypeFromValue(v string) (*DatabaseType, error) {
	ev := DatabaseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatabaseType: valid values are %v", v, AllowedDatabaseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatabaseType) IsValid() bool {
	for _, existing := range AllowedDatabaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatabaseType value
func (v DatabaseType) Ptr() *DatabaseType {
	return &v
}

type NullableDatabaseType struct {
	value *DatabaseType
	isSet bool
}

func (v NullableDatabaseType) Get() *DatabaseType {
	return v.value
}

func (v *NullableDatabaseType) Set(val *DatabaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseType(val *DatabaseType) *NullableDatabaseType {
	return &NullableDatabaseType{value: val, isSet: true}
}

func (v NullableDatabaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

