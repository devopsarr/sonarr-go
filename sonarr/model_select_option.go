/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// SelectOption struct for SelectOption
type SelectOption struct {
	Value *int32 `json:"value,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Hint NullableString `json:"hint,omitempty"`
}

// NewSelectOption instantiates a new SelectOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectOption() *SelectOption {
	this := SelectOption{}
	return &this
}

// NewSelectOptionWithDefaults instantiates a new SelectOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectOptionWithDefaults() *SelectOption {
	this := SelectOption{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SelectOption) GetValue() int32 {
	if o == nil || isNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectOption) GetValueOk() (*int32, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SelectOption) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *SelectOption) SetValue(v int32) {
	o.Value = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SelectOption) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SelectOption) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SelectOption) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SelectOption) UnsetName() {
	o.Name.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SelectOption) GetOrder() int32 {
	if o == nil || isNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectOption) GetOrderOk() (*int32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SelectOption) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SelectOption) SetOrder(v int32) {
	o.Order = &v
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetHint() string {
	if o == nil || isNil(o.Hint.Get()) {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetHintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *SelectOption) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *SelectOption) SetHint(v string) {
	o.Hint.Set(&v)
}
// SetHintNil sets the value for Hint to be an explicit nil
func (o *SelectOption) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *SelectOption) UnsetHint() {
	o.Hint.Unset()
}

func (o SelectOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Hint.IsSet() {
		toSerialize["hint"] = o.Hint.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSelectOption struct {
	value *SelectOption
	isSet bool
}

func (v NullableSelectOption) Get() *SelectOption {
	return v.value
}

func (v *NullableSelectOption) Set(val *SelectOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectOption(val *SelectOption) *NullableSelectOption {
	return &NullableSelectOption{value: val, isSet: true}
}

func (v NullableSelectOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


