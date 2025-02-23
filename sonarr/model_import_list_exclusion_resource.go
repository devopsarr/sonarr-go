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

// checks if the ImportListExclusionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportListExclusionResource{}

// ImportListExclusionResource struct for ImportListExclusionResource
type ImportListExclusionResource struct {
	Id *int32 `json:"id,omitempty"`
	TvdbId *int32 `json:"tvdbId,omitempty"`
	Title NullableString `json:"title,omitempty"`
}

// NewImportListExclusionResource instantiates a new ImportListExclusionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListExclusionResource() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// NewImportListExclusionResourceWithDefaults instantiates a new ImportListExclusionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListExclusionResourceWithDefaults() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportListExclusionResource) SetId(v int32) {
	o.Id = &v
}

// GetTvdbId returns the TvdbId field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetTvdbId() int32 {
	if o == nil || IsNil(o.TvdbId) {
		var ret int32
		return ret
	}
	return *o.TvdbId
}

// GetTvdbIdOk returns a tuple with the TvdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetTvdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TvdbId) {
		return nil, false
	}
	return o.TvdbId, true
}

// HasTvdbId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasTvdbId() bool {
	if o != nil && !IsNil(o.TvdbId) {
		return true
	}

	return false
}

// SetTvdbId gets a reference to the given int32 and assigns it to the TvdbId field.
func (o *ImportListExclusionResource) SetTvdbId(v int32) {
	o.TvdbId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ImportListExclusionResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ImportListExclusionResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ImportListExclusionResource) UnsetTitle() {
	o.Title.Unset()
}

func (o ImportListExclusionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportListExclusionResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TvdbId) {
		toSerialize["tvdbId"] = o.TvdbId
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableImportListExclusionResource struct {
	value *ImportListExclusionResource
	isSet bool
}

func (v NullableImportListExclusionResource) Get() *ImportListExclusionResource {
	return v.value
}

func (v *NullableImportListExclusionResource) Set(val *ImportListExclusionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListExclusionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListExclusionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListExclusionResource(val *ImportListExclusionResource) *NullableImportListExclusionResource {
	return &NullableImportListExclusionResource{value: val, isSet: true}
}

func (v NullableImportListExclusionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListExclusionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


