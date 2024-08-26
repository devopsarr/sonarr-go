/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.9.2244
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the UpdateChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateChanges{}

// UpdateChanges struct for UpdateChanges
type UpdateChanges struct {
	New []string `json:"new,omitempty"`
	Fixed []string `json:"fixed,omitempty"`
}

// NewUpdateChanges instantiates a new UpdateChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateChanges() *UpdateChanges {
	this := UpdateChanges{}
	return &this
}

// NewUpdateChangesWithDefaults instantiates a new UpdateChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateChangesWithDefaults() *UpdateChanges {
	this := UpdateChanges{}
	return &this
}

// GetNew returns the New field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateChanges) GetNew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateChanges) GetNewOk() ([]string, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *UpdateChanges) HasNew() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given []string and assigns it to the New field.
func (o *UpdateChanges) SetNew(v []string) {
	o.New = v
}

// GetFixed returns the Fixed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateChanges) GetFixed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateChanges) GetFixedOk() ([]string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *UpdateChanges) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given []string and assigns it to the Fixed field.
func (o *UpdateChanges) SetFixed(v []string) {
	o.Fixed = v
}

func (o UpdateChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	if o.Fixed != nil {
		toSerialize["fixed"] = o.Fixed
	}
	return toSerialize, nil
}

type NullableUpdateChanges struct {
	value *UpdateChanges
	isSet bool
}

func (v NullableUpdateChanges) Get() *UpdateChanges {
	return v.value
}

func (v *NullableUpdateChanges) Set(val *UpdateChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChanges(val *UpdateChanges) *NullableUpdateChanges {
	return &NullableUpdateChanges{value: val, isSet: true}
}

func (v NullableUpdateChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


