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

// ReleaseEpisodeResource struct for ReleaseEpisodeResource
type ReleaseEpisodeResource struct {
	Id *int32 `json:"id,omitempty"`
	SeasonNumber *int32 `json:"seasonNumber,omitempty"`
	EpisodeNumber *int32 `json:"episodeNumber,omitempty"`
	AbsoluteEpisodeNumber NullableInt32 `json:"absoluteEpisodeNumber,omitempty"`
	Title NullableString `json:"title,omitempty"`
}

// NewReleaseEpisodeResource instantiates a new ReleaseEpisodeResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseEpisodeResource() *ReleaseEpisodeResource {
	this := ReleaseEpisodeResource{}
	return &this
}

// NewReleaseEpisodeResourceWithDefaults instantiates a new ReleaseEpisodeResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseEpisodeResourceWithDefaults() *ReleaseEpisodeResource {
	this := ReleaseEpisodeResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReleaseEpisodeResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseEpisodeResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReleaseEpisodeResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ReleaseEpisodeResource) SetId(v int32) {
	o.Id = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *ReleaseEpisodeResource) GetSeasonNumber() int32 {
	if o == nil || isNil(o.SeasonNumber) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseEpisodeResource) GetSeasonNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SeasonNumber) {
    return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *ReleaseEpisodeResource) HasSeasonNumber() bool {
	if o != nil && !isNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given int32 and assigns it to the SeasonNumber field.
func (o *ReleaseEpisodeResource) SetSeasonNumber(v int32) {
	o.SeasonNumber = &v
}

// GetEpisodeNumber returns the EpisodeNumber field value if set, zero value otherwise.
func (o *ReleaseEpisodeResource) GetEpisodeNumber() int32 {
	if o == nil || isNil(o.EpisodeNumber) {
		var ret int32
		return ret
	}
	return *o.EpisodeNumber
}

// GetEpisodeNumberOk returns a tuple with the EpisodeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseEpisodeResource) GetEpisodeNumberOk() (*int32, bool) {
	if o == nil || isNil(o.EpisodeNumber) {
    return nil, false
	}
	return o.EpisodeNumber, true
}

// HasEpisodeNumber returns a boolean if a field has been set.
func (o *ReleaseEpisodeResource) HasEpisodeNumber() bool {
	if o != nil && !isNil(o.EpisodeNumber) {
		return true
	}

	return false
}

// SetEpisodeNumber gets a reference to the given int32 and assigns it to the EpisodeNumber field.
func (o *ReleaseEpisodeResource) SetEpisodeNumber(v int32) {
	o.EpisodeNumber = &v
}

// GetAbsoluteEpisodeNumber returns the AbsoluteEpisodeNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseEpisodeResource) GetAbsoluteEpisodeNumber() int32 {
	if o == nil || isNil(o.AbsoluteEpisodeNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.AbsoluteEpisodeNumber.Get()
}

// GetAbsoluteEpisodeNumberOk returns a tuple with the AbsoluteEpisodeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseEpisodeResource) GetAbsoluteEpisodeNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AbsoluteEpisodeNumber.Get(), o.AbsoluteEpisodeNumber.IsSet()
}

// HasAbsoluteEpisodeNumber returns a boolean if a field has been set.
func (o *ReleaseEpisodeResource) HasAbsoluteEpisodeNumber() bool {
	if o != nil && o.AbsoluteEpisodeNumber.IsSet() {
		return true
	}

	return false
}

// SetAbsoluteEpisodeNumber gets a reference to the given NullableInt32 and assigns it to the AbsoluteEpisodeNumber field.
func (o *ReleaseEpisodeResource) SetAbsoluteEpisodeNumber(v int32) {
	o.AbsoluteEpisodeNumber.Set(&v)
}
// SetAbsoluteEpisodeNumberNil sets the value for AbsoluteEpisodeNumber to be an explicit nil
func (o *ReleaseEpisodeResource) SetAbsoluteEpisodeNumberNil() {
	o.AbsoluteEpisodeNumber.Set(nil)
}

// UnsetAbsoluteEpisodeNumber ensures that no value is present for AbsoluteEpisodeNumber, not even an explicit nil
func (o *ReleaseEpisodeResource) UnsetAbsoluteEpisodeNumber() {
	o.AbsoluteEpisodeNumber.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReleaseEpisodeResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseEpisodeResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ReleaseEpisodeResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ReleaseEpisodeResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ReleaseEpisodeResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ReleaseEpisodeResource) UnsetTitle() {
	o.Title.Unset()
}

func (o ReleaseEpisodeResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !isNil(o.EpisodeNumber) {
		toSerialize["episodeNumber"] = o.EpisodeNumber
	}
	if o.AbsoluteEpisodeNumber.IsSet() {
		toSerialize["absoluteEpisodeNumber"] = o.AbsoluteEpisodeNumber.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReleaseEpisodeResource struct {
	value *ReleaseEpisodeResource
	isSet bool
}

func (v NullableReleaseEpisodeResource) Get() *ReleaseEpisodeResource {
	return v.value
}

func (v *NullableReleaseEpisodeResource) Set(val *ReleaseEpisodeResource) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseEpisodeResource) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseEpisodeResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseEpisodeResource(val *ReleaseEpisodeResource) *NullableReleaseEpisodeResource {
	return &NullableReleaseEpisodeResource{value: val, isSet: true}
}

func (v NullableReleaseEpisodeResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseEpisodeResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


