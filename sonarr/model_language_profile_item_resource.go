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

// checks if the LanguageProfileItemResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageProfileItemResource{}

// LanguageProfileItemResource struct for LanguageProfileItemResource
type LanguageProfileItemResource struct {
	Id *int32 `json:"id,omitempty"`
	Language *Language `json:"language,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewLanguageProfileItemResource instantiates a new LanguageProfileItemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageProfileItemResource() *LanguageProfileItemResource {
	this := LanguageProfileItemResource{}
	return &this
}

// NewLanguageProfileItemResourceWithDefaults instantiates a new LanguageProfileItemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageProfileItemResourceWithDefaults() *LanguageProfileItemResource {
	this := LanguageProfileItemResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LanguageProfileItemResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileItemResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LanguageProfileItemResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LanguageProfileItemResource) SetId(v int32) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *LanguageProfileItemResource) GetLanguage() Language {
	if o == nil || IsNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileItemResource) GetLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *LanguageProfileItemResource) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *LanguageProfileItemResource) SetLanguage(v Language) {
	o.Language = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *LanguageProfileItemResource) GetAllowed() bool {
	if o == nil || IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileItemResource) GetAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *LanguageProfileItemResource) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *LanguageProfileItemResource) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o LanguageProfileItemResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageProfileItemResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableLanguageProfileItemResource struct {
	value *LanguageProfileItemResource
	isSet bool
}

func (v NullableLanguageProfileItemResource) Get() *LanguageProfileItemResource {
	return v.value
}

func (v *NullableLanguageProfileItemResource) Set(val *LanguageProfileItemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageProfileItemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageProfileItemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageProfileItemResource(val *LanguageProfileItemResource) *NullableLanguageProfileItemResource {
	return &NullableLanguageProfileItemResource{value: val, isSet: true}
}

func (v NullableLanguageProfileItemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageProfileItemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


