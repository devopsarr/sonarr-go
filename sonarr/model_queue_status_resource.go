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

// checks if the QueueStatusResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueStatusResource{}

// QueueStatusResource struct for QueueStatusResource
type QueueStatusResource struct {
	Id *int32 `json:"id,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Count *int32 `json:"count,omitempty"`
	UnknownCount *int32 `json:"unknownCount,omitempty"`
	Errors *bool `json:"errors,omitempty"`
	Warnings *bool `json:"warnings,omitempty"`
	UnknownErrors *bool `json:"unknownErrors,omitempty"`
	UnknownWarnings *bool `json:"unknownWarnings,omitempty"`
}

// NewQueueStatusResource instantiates a new QueueStatusResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueStatusResource() *QueueStatusResource {
	this := QueueStatusResource{}
	return &this
}

// NewQueueStatusResourceWithDefaults instantiates a new QueueStatusResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueStatusResourceWithDefaults() *QueueStatusResource {
	this := QueueStatusResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueueStatusResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueueStatusResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QueueStatusResource) SetId(v int32) {
	o.Id = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QueueStatusResource) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QueueStatusResource) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *QueueStatusResource) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueueStatusResource) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueueStatusResource) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *QueueStatusResource) SetCount(v int32) {
	o.Count = &v
}

// GetUnknownCount returns the UnknownCount field value if set, zero value otherwise.
func (o *QueueStatusResource) GetUnknownCount() int32 {
	if o == nil || IsNil(o.UnknownCount) {
		var ret int32
		return ret
	}
	return *o.UnknownCount
}

// GetUnknownCountOk returns a tuple with the UnknownCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetUnknownCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnknownCount) {
		return nil, false
	}
	return o.UnknownCount, true
}

// HasUnknownCount returns a boolean if a field has been set.
func (o *QueueStatusResource) HasUnknownCount() bool {
	if o != nil && !IsNil(o.UnknownCount) {
		return true
	}

	return false
}

// SetUnknownCount gets a reference to the given int32 and assigns it to the UnknownCount field.
func (o *QueueStatusResource) SetUnknownCount(v int32) {
	o.UnknownCount = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *QueueStatusResource) GetErrors() bool {
	if o == nil || IsNil(o.Errors) {
		var ret bool
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetErrorsOk() (*bool, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *QueueStatusResource) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given bool and assigns it to the Errors field.
func (o *QueueStatusResource) SetErrors(v bool) {
	o.Errors = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *QueueStatusResource) GetWarnings() bool {
	if o == nil || IsNil(o.Warnings) {
		var ret bool
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetWarningsOk() (*bool, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *QueueStatusResource) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given bool and assigns it to the Warnings field.
func (o *QueueStatusResource) SetWarnings(v bool) {
	o.Warnings = &v
}

// GetUnknownErrors returns the UnknownErrors field value if set, zero value otherwise.
func (o *QueueStatusResource) GetUnknownErrors() bool {
	if o == nil || IsNil(o.UnknownErrors) {
		var ret bool
		return ret
	}
	return *o.UnknownErrors
}

// GetUnknownErrorsOk returns a tuple with the UnknownErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetUnknownErrorsOk() (*bool, bool) {
	if o == nil || IsNil(o.UnknownErrors) {
		return nil, false
	}
	return o.UnknownErrors, true
}

// HasUnknownErrors returns a boolean if a field has been set.
func (o *QueueStatusResource) HasUnknownErrors() bool {
	if o != nil && !IsNil(o.UnknownErrors) {
		return true
	}

	return false
}

// SetUnknownErrors gets a reference to the given bool and assigns it to the UnknownErrors field.
func (o *QueueStatusResource) SetUnknownErrors(v bool) {
	o.UnknownErrors = &v
}

// GetUnknownWarnings returns the UnknownWarnings field value if set, zero value otherwise.
func (o *QueueStatusResource) GetUnknownWarnings() bool {
	if o == nil || IsNil(o.UnknownWarnings) {
		var ret bool
		return ret
	}
	return *o.UnknownWarnings
}

// GetUnknownWarningsOk returns a tuple with the UnknownWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatusResource) GetUnknownWarningsOk() (*bool, bool) {
	if o == nil || IsNil(o.UnknownWarnings) {
		return nil, false
	}
	return o.UnknownWarnings, true
}

// HasUnknownWarnings returns a boolean if a field has been set.
func (o *QueueStatusResource) HasUnknownWarnings() bool {
	if o != nil && !IsNil(o.UnknownWarnings) {
		return true
	}

	return false
}

// SetUnknownWarnings gets a reference to the given bool and assigns it to the UnknownWarnings field.
func (o *QueueStatusResource) SetUnknownWarnings(v bool) {
	o.UnknownWarnings = &v
}

func (o QueueStatusResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueStatusResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.UnknownCount) {
		toSerialize["unknownCount"] = o.UnknownCount
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.UnknownErrors) {
		toSerialize["unknownErrors"] = o.UnknownErrors
	}
	if !IsNil(o.UnknownWarnings) {
		toSerialize["unknownWarnings"] = o.UnknownWarnings
	}
	return toSerialize, nil
}

type NullableQueueStatusResource struct {
	value *QueueStatusResource
	isSet bool
}

func (v NullableQueueStatusResource) Get() *QueueStatusResource {
	return v.value
}

func (v *NullableQueueStatusResource) Set(val *QueueStatusResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueStatusResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueStatusResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueStatusResource(val *QueueStatusResource) *NullableQueueStatusResource {
	return &NullableQueueStatusResource{value: val, isSet: true}
}

func (v NullableQueueStatusResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueStatusResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


