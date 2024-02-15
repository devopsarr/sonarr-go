/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.1.929
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the LanguageProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageProfileResource{}

// LanguageProfileResource struct for LanguageProfileResource
type LanguageProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	UpgradeAllowed *bool `json:"upgradeAllowed,omitempty"`
	Cutoff *Language `json:"cutoff,omitempty"`
	Languages []LanguageProfileItemResource `json:"languages,omitempty"`
}

// NewLanguageProfileResource instantiates a new LanguageProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageProfileResource() *LanguageProfileResource {
	this := LanguageProfileResource{}
	return &this
}

// NewLanguageProfileResourceWithDefaults instantiates a new LanguageProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageProfileResourceWithDefaults() *LanguageProfileResource {
	this := LanguageProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LanguageProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LanguageProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LanguageProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LanguageProfileResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanguageProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LanguageProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LanguageProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LanguageProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LanguageProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetUpgradeAllowed returns the UpgradeAllowed field value if set, zero value otherwise.
func (o *LanguageProfileResource) GetUpgradeAllowed() bool {
	if o == nil || IsNil(o.UpgradeAllowed) {
		var ret bool
		return ret
	}
	return *o.UpgradeAllowed
}

// GetUpgradeAllowedOk returns a tuple with the UpgradeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileResource) GetUpgradeAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.UpgradeAllowed) {
		return nil, false
	}
	return o.UpgradeAllowed, true
}

// HasUpgradeAllowed returns a boolean if a field has been set.
func (o *LanguageProfileResource) HasUpgradeAllowed() bool {
	if o != nil && !IsNil(o.UpgradeAllowed) {
		return true
	}

	return false
}

// SetUpgradeAllowed gets a reference to the given bool and assigns it to the UpgradeAllowed field.
func (o *LanguageProfileResource) SetUpgradeAllowed(v bool) {
	o.UpgradeAllowed = &v
}

// GetCutoff returns the Cutoff field value if set, zero value otherwise.
func (o *LanguageProfileResource) GetCutoff() Language {
	if o == nil || IsNil(o.Cutoff) {
		var ret Language
		return ret
	}
	return *o.Cutoff
}

// GetCutoffOk returns a tuple with the Cutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageProfileResource) GetCutoffOk() (*Language, bool) {
	if o == nil || IsNil(o.Cutoff) {
		return nil, false
	}
	return o.Cutoff, true
}

// HasCutoff returns a boolean if a field has been set.
func (o *LanguageProfileResource) HasCutoff() bool {
	if o != nil && !IsNil(o.Cutoff) {
		return true
	}

	return false
}

// SetCutoff gets a reference to the given Language and assigns it to the Cutoff field.
func (o *LanguageProfileResource) SetCutoff(v Language) {
	o.Cutoff = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LanguageProfileResource) GetLanguages() []LanguageProfileItemResource {
	if o == nil {
		var ret []LanguageProfileItemResource
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanguageProfileResource) GetLanguagesOk() ([]LanguageProfileItemResource, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *LanguageProfileResource) HasLanguages() bool {
	if o != nil && IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []LanguageProfileItemResource and assigns it to the Languages field.
func (o *LanguageProfileResource) SetLanguages(v []LanguageProfileItemResource) {
	o.Languages = v
}

func (o LanguageProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.UpgradeAllowed) {
		toSerialize["upgradeAllowed"] = o.UpgradeAllowed
	}
	if !IsNil(o.Cutoff) {
		toSerialize["cutoff"] = o.Cutoff
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableLanguageProfileResource struct {
	value *LanguageProfileResource
	isSet bool
}

func (v NullableLanguageProfileResource) Get() *LanguageProfileResource {
	return v.value
}

func (v *NullableLanguageProfileResource) Set(val *LanguageProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageProfileResource(val *LanguageProfileResource) *NullableLanguageProfileResource {
	return &NullableLanguageProfileResource{value: val, isSet: true}
}

func (v NullableLanguageProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


