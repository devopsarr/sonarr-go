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

// checks if the LocalizationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalizationResource{}

// LocalizationResource struct for LocalizationResource
type LocalizationResource struct {
	Id *int32 `json:"id,omitempty"`
	Strings map[string]string `json:"strings,omitempty"`
}

// NewLocalizationResource instantiates a new LocalizationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalizationResource() *LocalizationResource {
	this := LocalizationResource{}
	return &this
}

// NewLocalizationResourceWithDefaults instantiates a new LocalizationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizationResourceWithDefaults() *LocalizationResource {
	this := LocalizationResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocalizationResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizationResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocalizationResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LocalizationResource) SetId(v int32) {
	o.Id = &v
}

// GetStrings returns the Strings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizationResource) GetStrings() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Strings
}

// GetStringsOk returns a tuple with the Strings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizationResource) GetStringsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Strings) {
		return nil, false
	}
	return &o.Strings, true
}

// HasStrings returns a boolean if a field has been set.
func (o *LocalizationResource) HasStrings() bool {
	if o != nil && !IsNil(o.Strings) {
		return true
	}

	return false
}

// SetStrings gets a reference to the given map[string]string and assigns it to the Strings field.
func (o *LocalizationResource) SetStrings(v map[string]string) {
	o.Strings = v
}

func (o LocalizationResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalizationResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Strings != nil {
		toSerialize["strings"] = o.Strings
	}
	return toSerialize, nil
}

type NullableLocalizationResource struct {
	value *LocalizationResource
	isSet bool
}

func (v NullableLocalizationResource) Get() *LocalizationResource {
	return v.value
}

func (v *NullableLocalizationResource) Set(val *LocalizationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalizationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalizationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalizationResource(val *LocalizationResource) *NullableLocalizationResource {
	return &NullableLocalizationResource{value: val, isSet: true}
}

func (v NullableLocalizationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalizationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


