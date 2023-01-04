/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// CustomFilterResource struct for CustomFilterResource
type CustomFilterResource struct {
	Id *int32 `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Label NullableString `json:"label,omitempty"`
	Filters []map[string]interface{} `json:"filters,omitempty"`
}

// NewCustomFilterResource instantiates a new CustomFilterResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFilterResource() *CustomFilterResource {
	this := CustomFilterResource{}
	return &this
}

// NewCustomFilterResourceWithDefaults instantiates a new CustomFilterResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFilterResourceWithDefaults() *CustomFilterResource {
	this := CustomFilterResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFilterResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFilterResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFilterResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFilterResource) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFilterResource) GetType() string {
	if o == nil || isNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFilterResource) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CustomFilterResource) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CustomFilterResource) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CustomFilterResource) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CustomFilterResource) UnsetType() {
	o.Type.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFilterResource) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFilterResource) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFilterResource) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *CustomFilterResource) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *CustomFilterResource) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *CustomFilterResource) UnsetLabel() {
	o.Label.Unset()
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFilterResource) GetFilters() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFilterResource) GetFiltersOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CustomFilterResource) HasFilters() bool {
	if o != nil && isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []map[string]interface{} and assigns it to the Filters field.
func (o *CustomFilterResource) SetFilters(v []map[string]interface{}) {
	o.Filters = v
}

func (o CustomFilterResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableCustomFilterResource struct {
	value *CustomFilterResource
	isSet bool
}

func (v NullableCustomFilterResource) Get() *CustomFilterResource {
	return v.value
}

func (v *NullableCustomFilterResource) Set(val *CustomFilterResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFilterResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFilterResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFilterResource(val *CustomFilterResource) *NullableCustomFilterResource {
	return &NullableCustomFilterResource{value: val, isSet: true}
}

func (v NullableCustomFilterResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFilterResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


