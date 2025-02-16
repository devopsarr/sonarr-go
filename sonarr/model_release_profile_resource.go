/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.13.2932
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the ReleaseProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseProfileResource{}

// ReleaseProfileResource struct for ReleaseProfileResource
type ReleaseProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Required []string `json:"required,omitempty"`
	Ignored []string `json:"ignored,omitempty"`
	IndexerId *int32 `json:"indexerId,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
}

// NewReleaseProfileResource instantiates a new ReleaseProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseProfileResource() *ReleaseProfileResource {
	this := ReleaseProfileResource{}
	return &this
}

// NewReleaseProfileResourceWithDefaults instantiates a new ReleaseProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseProfileResourceWithDefaults() *ReleaseProfileResource {
	this := ReleaseProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReleaseProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReleaseProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseProfileResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseProfileResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ReleaseProfileResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ReleaseProfileResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ReleaseProfileResource) UnsetName() {
	o.Name.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ReleaseProfileResource) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseProfileResource) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ReleaseProfileResource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseProfileResource) GetRequired() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseProfileResource) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *ReleaseProfileResource) SetRequired(v []string) {
	o.Required = v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseProfileResource) GetIgnored() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseProfileResource) GetIgnoredOk() ([]string, bool) {
	if o == nil || IsNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasIgnored() bool {
	if o != nil && !IsNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given []string and assigns it to the Ignored field.
func (o *ReleaseProfileResource) SetIgnored(v []string) {
	o.Ignored = v
}

// GetIndexerId returns the IndexerId field value if set, zero value otherwise.
func (o *ReleaseProfileResource) GetIndexerId() int32 {
	if o == nil || IsNil(o.IndexerId) {
		var ret int32
		return ret
	}
	return *o.IndexerId
}

// GetIndexerIdOk returns a tuple with the IndexerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseProfileResource) GetIndexerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerId) {
		return nil, false
	}
	return o.IndexerId, true
}

// HasIndexerId returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasIndexerId() bool {
	if o != nil && !IsNil(o.IndexerId) {
		return true
	}

	return false
}

// SetIndexerId gets a reference to the given int32 and assigns it to the IndexerId field.
func (o *ReleaseProfileResource) SetIndexerId(v int32) {
	o.IndexerId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseProfileResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseProfileResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReleaseProfileResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ReleaseProfileResource) SetTags(v []int32) {
	o.Tags = v
}

func (o ReleaseProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Ignored != nil {
		toSerialize["ignored"] = o.Ignored
	}
	if !IsNil(o.IndexerId) {
		toSerialize["indexerId"] = o.IndexerId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableReleaseProfileResource struct {
	value *ReleaseProfileResource
	isSet bool
}

func (v NullableReleaseProfileResource) Get() *ReleaseProfileResource {
	return v.value
}

func (v *NullableReleaseProfileResource) Set(val *ReleaseProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseProfileResource(val *ReleaseProfileResource) *NullableReleaseProfileResource {
	return &NullableReleaseProfileResource{value: val, isSet: true}
}

func (v NullableReleaseProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


