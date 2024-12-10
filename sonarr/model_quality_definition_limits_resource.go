/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.11.2680
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the QualityDefinitionLimitsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QualityDefinitionLimitsResource{}

// QualityDefinitionLimitsResource struct for QualityDefinitionLimitsResource
type QualityDefinitionLimitsResource struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewQualityDefinitionLimitsResource instantiates a new QualityDefinitionLimitsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityDefinitionLimitsResource() *QualityDefinitionLimitsResource {
	this := QualityDefinitionLimitsResource{}
	return &this
}

// NewQualityDefinitionLimitsResourceWithDefaults instantiates a new QualityDefinitionLimitsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityDefinitionLimitsResourceWithDefaults() *QualityDefinitionLimitsResource {
	this := QualityDefinitionLimitsResource{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *QualityDefinitionLimitsResource) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityDefinitionLimitsResource) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *QualityDefinitionLimitsResource) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *QualityDefinitionLimitsResource) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *QualityDefinitionLimitsResource) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityDefinitionLimitsResource) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *QualityDefinitionLimitsResource) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *QualityDefinitionLimitsResource) SetMax(v int32) {
	o.Max = &v
}

func (o QualityDefinitionLimitsResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QualityDefinitionLimitsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableQualityDefinitionLimitsResource struct {
	value *QualityDefinitionLimitsResource
	isSet bool
}

func (v NullableQualityDefinitionLimitsResource) Get() *QualityDefinitionLimitsResource {
	return v.value
}

func (v *NullableQualityDefinitionLimitsResource) Set(val *QualityDefinitionLimitsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityDefinitionLimitsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityDefinitionLimitsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityDefinitionLimitsResource(val *QualityDefinitionLimitsResource) *NullableQualityDefinitionLimitsResource {
	return &NullableQualityDefinitionLimitsResource{value: val, isSet: true}
}

func (v NullableQualityDefinitionLimitsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityDefinitionLimitsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

