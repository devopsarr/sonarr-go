/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.12.2823
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"time"
)

// checks if the UpdateResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateResource{}

// UpdateResource struct for UpdateResource
type UpdateResource struct {
	Id *int32 `json:"id,omitempty"`
	Version NullableString `json:"version,omitempty"`
	Branch NullableString `json:"branch,omitempty"`
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
	FileName NullableString `json:"fileName,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Installed *bool `json:"installed,omitempty"`
	InstalledOn NullableTime `json:"installedOn,omitempty"`
	Installable *bool `json:"installable,omitempty"`
	Latest *bool `json:"latest,omitempty"`
	Changes *UpdateChanges `json:"changes,omitempty"`
	Hash NullableString `json:"hash,omitempty"`
}

// NewUpdateResource instantiates a new UpdateResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateResource() *UpdateResource {
	this := UpdateResource{}
	return &this
}

// NewUpdateResourceWithDefaults instantiates a new UpdateResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateResourceWithDefaults() *UpdateResource {
	this := UpdateResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateResource) SetId(v int32) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *UpdateResource) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *UpdateResource) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *UpdateResource) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *UpdateResource) UnsetVersion() {
	o.Version.Unset()
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetBranch() string {
	if o == nil || IsNil(o.Branch.Get()) {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *UpdateResource) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *UpdateResource) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *UpdateResource) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *UpdateResource) UnsetBranch() {
	o.Branch.Unset()
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *UpdateResource) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *UpdateResource) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *UpdateResource) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *UpdateResource) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *UpdateResource) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *UpdateResource) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *UpdateResource) UnsetFileName() {
	o.FileName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateResource) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *UpdateResource) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *UpdateResource) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *UpdateResource) UnsetUrl() {
	o.Url.Unset()
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *UpdateResource) GetInstalled() bool {
	if o == nil || IsNil(o.Installed) {
		var ret bool
		return ret
	}
	return *o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetInstalledOk() (*bool, bool) {
	if o == nil || IsNil(o.Installed) {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *UpdateResource) HasInstalled() bool {
	if o != nil && !IsNil(o.Installed) {
		return true
	}

	return false
}

// SetInstalled gets a reference to the given bool and assigns it to the Installed field.
func (o *UpdateResource) SetInstalled(v bool) {
	o.Installed = &v
}

// GetInstalledOn returns the InstalledOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetInstalledOn() time.Time {
	if o == nil || IsNil(o.InstalledOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.InstalledOn.Get()
}

// GetInstalledOnOk returns a tuple with the InstalledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetInstalledOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledOn.Get(), o.InstalledOn.IsSet()
}

// HasInstalledOn returns a boolean if a field has been set.
func (o *UpdateResource) HasInstalledOn() bool {
	if o != nil && o.InstalledOn.IsSet() {
		return true
	}

	return false
}

// SetInstalledOn gets a reference to the given NullableTime and assigns it to the InstalledOn field.
func (o *UpdateResource) SetInstalledOn(v time.Time) {
	o.InstalledOn.Set(&v)
}
// SetInstalledOnNil sets the value for InstalledOn to be an explicit nil
func (o *UpdateResource) SetInstalledOnNil() {
	o.InstalledOn.Set(nil)
}

// UnsetInstalledOn ensures that no value is present for InstalledOn, not even an explicit nil
func (o *UpdateResource) UnsetInstalledOn() {
	o.InstalledOn.Unset()
}

// GetInstallable returns the Installable field value if set, zero value otherwise.
func (o *UpdateResource) GetInstallable() bool {
	if o == nil || IsNil(o.Installable) {
		var ret bool
		return ret
	}
	return *o.Installable
}

// GetInstallableOk returns a tuple with the Installable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetInstallableOk() (*bool, bool) {
	if o == nil || IsNil(o.Installable) {
		return nil, false
	}
	return o.Installable, true
}

// HasInstallable returns a boolean if a field has been set.
func (o *UpdateResource) HasInstallable() bool {
	if o != nil && !IsNil(o.Installable) {
		return true
	}

	return false
}

// SetInstallable gets a reference to the given bool and assigns it to the Installable field.
func (o *UpdateResource) SetInstallable(v bool) {
	o.Installable = &v
}

// GetLatest returns the Latest field value if set, zero value otherwise.
func (o *UpdateResource) GetLatest() bool {
	if o == nil || IsNil(o.Latest) {
		var ret bool
		return ret
	}
	return *o.Latest
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.Latest) {
		return nil, false
	}
	return o.Latest, true
}

// HasLatest returns a boolean if a field has been set.
func (o *UpdateResource) HasLatest() bool {
	if o != nil && !IsNil(o.Latest) {
		return true
	}

	return false
}

// SetLatest gets a reference to the given bool and assigns it to the Latest field.
func (o *UpdateResource) SetLatest(v bool) {
	o.Latest = &v
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *UpdateResource) GetChanges() UpdateChanges {
	if o == nil || IsNil(o.Changes) {
		var ret UpdateChanges
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateResource) GetChangesOk() (*UpdateChanges, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *UpdateResource) HasChanges() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given UpdateChanges and assigns it to the Changes field.
func (o *UpdateResource) SetChanges(v UpdateChanges) {
	o.Changes = &v
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateResource) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateResource) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *UpdateResource) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *UpdateResource) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *UpdateResource) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *UpdateResource) UnsetHash() {
	o.Hash.Unset()
}

func (o UpdateResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Installed) {
		toSerialize["installed"] = o.Installed
	}
	if o.InstalledOn.IsSet() {
		toSerialize["installedOn"] = o.InstalledOn.Get()
	}
	if !IsNil(o.Installable) {
		toSerialize["installable"] = o.Installable
	}
	if !IsNil(o.Latest) {
		toSerialize["latest"] = o.Latest
	}
	if !IsNil(o.Changes) {
		toSerialize["changes"] = o.Changes
	}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	return toSerialize, nil
}

type NullableUpdateResource struct {
	value *UpdateResource
	isSet bool
}

func (v NullableUpdateResource) Get() *UpdateResource {
	return v.value
}

func (v *NullableUpdateResource) Set(val *UpdateResource) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateResource) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateResource(val *UpdateResource) *NullableUpdateResource {
	return &NullableUpdateResource{value: val, isSet: true}
}

func (v NullableUpdateResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


