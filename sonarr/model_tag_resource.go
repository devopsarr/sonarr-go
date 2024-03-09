/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.2.1183
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the TagResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagResource{}

// TagResource struct for TagResource
type TagResource struct {
	Id *int32 `json:"id,omitempty"`
	Label NullableString `json:"label,omitempty"`
}

// NewTagResource instantiates a new TagResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResource() *TagResource {
	this := TagResource{}
	return &this
}

// NewTagResourceWithDefaults instantiates a new TagResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResourceWithDefaults() *TagResource {
	this := TagResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TagResource) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagResource) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagResource) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *TagResource) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *TagResource) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *TagResource) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *TagResource) UnsetLabel() {
	o.Label.Unset()
}

func (o TagResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	return toSerialize, nil
}

type NullableTagResource struct {
	value *TagResource
	isSet bool
}

func (v NullableTagResource) Get() *TagResource {
	return v.value
}

func (v *NullableTagResource) Set(val *TagResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResource(val *TagResource) *NullableTagResource {
	return &NullableTagResource{value: val, isSet: true}
}

func (v NullableTagResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


