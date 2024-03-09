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

// checks if the RootFolderResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootFolderResource{}

// RootFolderResource struct for RootFolderResource
type RootFolderResource struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Accessible *bool `json:"accessible,omitempty"`
	FreeSpace NullableInt64 `json:"freeSpace,omitempty"`
	UnmappedFolders []UnmappedFolder `json:"unmappedFolders,omitempty"`
}

// NewRootFolderResource instantiates a new RootFolderResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootFolderResource() *RootFolderResource {
	this := RootFolderResource{}
	return &this
}

// NewRootFolderResourceWithDefaults instantiates a new RootFolderResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootFolderResourceWithDefaults() *RootFolderResource {
	this := RootFolderResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RootFolderResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RootFolderResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RootFolderResource) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *RootFolderResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *RootFolderResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *RootFolderResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *RootFolderResource) UnsetPath() {
	o.Path.Unset()
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *RootFolderResource) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootFolderResource) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *RootFolderResource) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *RootFolderResource) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetFreeSpace() int64 {
	if o == nil || IsNil(o.FreeSpace.Get()) {
		var ret int64
		return ret
	}
	return *o.FreeSpace.Get()
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetFreeSpaceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreeSpace.Get(), o.FreeSpace.IsSet()
}

// HasFreeSpace returns a boolean if a field has been set.
func (o *RootFolderResource) HasFreeSpace() bool {
	if o != nil && o.FreeSpace.IsSet() {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given NullableInt64 and assigns it to the FreeSpace field.
func (o *RootFolderResource) SetFreeSpace(v int64) {
	o.FreeSpace.Set(&v)
}
// SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil
func (o *RootFolderResource) SetFreeSpaceNil() {
	o.FreeSpace.Set(nil)
}

// UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
func (o *RootFolderResource) UnsetFreeSpace() {
	o.FreeSpace.Unset()
}

// GetUnmappedFolders returns the UnmappedFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RootFolderResource) GetUnmappedFolders() []UnmappedFolder {
	if o == nil {
		var ret []UnmappedFolder
		return ret
	}
	return o.UnmappedFolders
}

// GetUnmappedFoldersOk returns a tuple with the UnmappedFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RootFolderResource) GetUnmappedFoldersOk() ([]UnmappedFolder, bool) {
	if o == nil || IsNil(o.UnmappedFolders) {
		return nil, false
	}
	return o.UnmappedFolders, true
}

// HasUnmappedFolders returns a boolean if a field has been set.
func (o *RootFolderResource) HasUnmappedFolders() bool {
	if o != nil && IsNil(o.UnmappedFolders) {
		return true
	}

	return false
}

// SetUnmappedFolders gets a reference to the given []UnmappedFolder and assigns it to the UnmappedFolders field.
func (o *RootFolderResource) SetUnmappedFolders(v []UnmappedFolder) {
	o.UnmappedFolders = v
}

func (o RootFolderResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootFolderResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.Accessible) {
		toSerialize["accessible"] = o.Accessible
	}
	if o.FreeSpace.IsSet() {
		toSerialize["freeSpace"] = o.FreeSpace.Get()
	}
	if o.UnmappedFolders != nil {
		toSerialize["unmappedFolders"] = o.UnmappedFolders
	}
	return toSerialize, nil
}

type NullableRootFolderResource struct {
	value *RootFolderResource
	isSet bool
}

func (v NullableRootFolderResource) Get() *RootFolderResource {
	return v.value
}

func (v *NullableRootFolderResource) Set(val *RootFolderResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRootFolderResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRootFolderResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootFolderResource(val *RootFolderResource) *NullableRootFolderResource {
	return &NullableRootFolderResource{value: val, isSet: true}
}

func (v NullableRootFolderResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootFolderResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


