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

// checks if the EpisodesMonitoredResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpisodesMonitoredResource{}

// EpisodesMonitoredResource struct for EpisodesMonitoredResource
type EpisodesMonitoredResource struct {
	EpisodeIds []int32 `json:"episodeIds,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
}

// NewEpisodesMonitoredResource instantiates a new EpisodesMonitoredResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisodesMonitoredResource() *EpisodesMonitoredResource {
	this := EpisodesMonitoredResource{}
	return &this
}

// NewEpisodesMonitoredResourceWithDefaults instantiates a new EpisodesMonitoredResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodesMonitoredResourceWithDefaults() *EpisodesMonitoredResource {
	this := EpisodesMonitoredResource{}
	return &this
}

// GetEpisodeIds returns the EpisodeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodesMonitoredResource) GetEpisodeIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.EpisodeIds
}

// GetEpisodeIdsOk returns a tuple with the EpisodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodesMonitoredResource) GetEpisodeIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.EpisodeIds) {
		return nil, false
	}
	return o.EpisodeIds, true
}

// HasEpisodeIds returns a boolean if a field has been set.
func (o *EpisodesMonitoredResource) HasEpisodeIds() bool {
	if o != nil && !IsNil(o.EpisodeIds) {
		return true
	}

	return false
}

// SetEpisodeIds gets a reference to the given []int32 and assigns it to the EpisodeIds field.
func (o *EpisodesMonitoredResource) SetEpisodeIds(v []int32) {
	o.EpisodeIds = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *EpisodesMonitoredResource) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodesMonitoredResource) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *EpisodesMonitoredResource) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *EpisodesMonitoredResource) SetMonitored(v bool) {
	o.Monitored = &v
}

func (o EpisodesMonitoredResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpisodesMonitoredResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EpisodeIds != nil {
		toSerialize["episodeIds"] = o.EpisodeIds
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	return toSerialize, nil
}

type NullableEpisodesMonitoredResource struct {
	value *EpisodesMonitoredResource
	isSet bool
}

func (v NullableEpisodesMonitoredResource) Get() *EpisodesMonitoredResource {
	return v.value
}

func (v *NullableEpisodesMonitoredResource) Set(val *EpisodesMonitoredResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodesMonitoredResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodesMonitoredResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodesMonitoredResource(val *EpisodesMonitoredResource) *NullableEpisodesMonitoredResource {
	return &NullableEpisodesMonitoredResource{value: val, isSet: true}
}

func (v NullableEpisodesMonitoredResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodesMonitoredResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


