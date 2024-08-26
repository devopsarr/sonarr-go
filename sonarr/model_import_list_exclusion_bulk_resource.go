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

// checks if the ImportListExclusionBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportListExclusionBulkResource{}

// ImportListExclusionBulkResource struct for ImportListExclusionBulkResource
type ImportListExclusionBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
}

// NewImportListExclusionBulkResource instantiates a new ImportListExclusionBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListExclusionBulkResource() *ImportListExclusionBulkResource {
	this := ImportListExclusionBulkResource{}
	return &this
}

// NewImportListExclusionBulkResourceWithDefaults instantiates a new ImportListExclusionBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListExclusionBulkResourceWithDefaults() *ImportListExclusionBulkResource {
	this := ImportListExclusionBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ImportListExclusionBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *ImportListExclusionBulkResource) SetIds(v []int32) {
	o.Ids = v
}

func (o ImportListExclusionBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportListExclusionBulkResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableImportListExclusionBulkResource struct {
	value *ImportListExclusionBulkResource
	isSet bool
}

func (v NullableImportListExclusionBulkResource) Get() *ImportListExclusionBulkResource {
	return v.value
}

func (v *NullableImportListExclusionBulkResource) Set(val *ImportListExclusionBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListExclusionBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListExclusionBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListExclusionBulkResource(val *ImportListExclusionBulkResource) *NullableImportListExclusionBulkResource {
	return &NullableImportListExclusionBulkResource{value: val, isSet: true}
}

func (v NullableImportListExclusionBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListExclusionBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


