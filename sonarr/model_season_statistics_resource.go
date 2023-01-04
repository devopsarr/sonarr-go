/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"time"
)

// SeasonStatisticsResource struct for SeasonStatisticsResource
type SeasonStatisticsResource struct {
	NextAiring NullableTime `json:"nextAiring,omitempty"`
	PreviousAiring NullableTime `json:"previousAiring,omitempty"`
	EpisodeFileCount *int32 `json:"episodeFileCount,omitempty"`
	EpisodeCount *int32 `json:"episodeCount,omitempty"`
	TotalEpisodeCount *int32 `json:"totalEpisodeCount,omitempty"`
	SizeOnDisk *int64 `json:"sizeOnDisk,omitempty"`
	ReleaseGroups []string `json:"releaseGroups,omitempty"`
	PercentOfEpisodes *float64 `json:"percentOfEpisodes,omitempty"`
}

// NewSeasonStatisticsResource instantiates a new SeasonStatisticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeasonStatisticsResource() *SeasonStatisticsResource {
	this := SeasonStatisticsResource{}
	return &this
}

// NewSeasonStatisticsResourceWithDefaults instantiates a new SeasonStatisticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeasonStatisticsResourceWithDefaults() *SeasonStatisticsResource {
	this := SeasonStatisticsResource{}
	return &this
}

// GetNextAiring returns the NextAiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeasonStatisticsResource) GetNextAiring() time.Time {
	if o == nil || isNil(o.NextAiring.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextAiring.Get()
}

// GetNextAiringOk returns a tuple with the NextAiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeasonStatisticsResource) GetNextAiringOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextAiring.Get(), o.NextAiring.IsSet()
}

// HasNextAiring returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasNextAiring() bool {
	if o != nil && o.NextAiring.IsSet() {
		return true
	}

	return false
}

// SetNextAiring gets a reference to the given NullableTime and assigns it to the NextAiring field.
func (o *SeasonStatisticsResource) SetNextAiring(v time.Time) {
	o.NextAiring.Set(&v)
}
// SetNextAiringNil sets the value for NextAiring to be an explicit nil
func (o *SeasonStatisticsResource) SetNextAiringNil() {
	o.NextAiring.Set(nil)
}

// UnsetNextAiring ensures that no value is present for NextAiring, not even an explicit nil
func (o *SeasonStatisticsResource) UnsetNextAiring() {
	o.NextAiring.Unset()
}

// GetPreviousAiring returns the PreviousAiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeasonStatisticsResource) GetPreviousAiring() time.Time {
	if o == nil || isNil(o.PreviousAiring.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PreviousAiring.Get()
}

// GetPreviousAiringOk returns a tuple with the PreviousAiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeasonStatisticsResource) GetPreviousAiringOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.PreviousAiring.Get(), o.PreviousAiring.IsSet()
}

// HasPreviousAiring returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasPreviousAiring() bool {
	if o != nil && o.PreviousAiring.IsSet() {
		return true
	}

	return false
}

// SetPreviousAiring gets a reference to the given NullableTime and assigns it to the PreviousAiring field.
func (o *SeasonStatisticsResource) SetPreviousAiring(v time.Time) {
	o.PreviousAiring.Set(&v)
}
// SetPreviousAiringNil sets the value for PreviousAiring to be an explicit nil
func (o *SeasonStatisticsResource) SetPreviousAiringNil() {
	o.PreviousAiring.Set(nil)
}

// UnsetPreviousAiring ensures that no value is present for PreviousAiring, not even an explicit nil
func (o *SeasonStatisticsResource) UnsetPreviousAiring() {
	o.PreviousAiring.Unset()
}

// GetEpisodeFileCount returns the EpisodeFileCount field value if set, zero value otherwise.
func (o *SeasonStatisticsResource) GetEpisodeFileCount() int32 {
	if o == nil || isNil(o.EpisodeFileCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeFileCount
}

// GetEpisodeFileCountOk returns a tuple with the EpisodeFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonStatisticsResource) GetEpisodeFileCountOk() (*int32, bool) {
	if o == nil || isNil(o.EpisodeFileCount) {
    return nil, false
	}
	return o.EpisodeFileCount, true
}

// HasEpisodeFileCount returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasEpisodeFileCount() bool {
	if o != nil && !isNil(o.EpisodeFileCount) {
		return true
	}

	return false
}

// SetEpisodeFileCount gets a reference to the given int32 and assigns it to the EpisodeFileCount field.
func (o *SeasonStatisticsResource) SetEpisodeFileCount(v int32) {
	o.EpisodeFileCount = &v
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *SeasonStatisticsResource) GetEpisodeCount() int32 {
	if o == nil || isNil(o.EpisodeCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonStatisticsResource) GetEpisodeCountOk() (*int32, bool) {
	if o == nil || isNil(o.EpisodeCount) {
    return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasEpisodeCount() bool {
	if o != nil && !isNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given int32 and assigns it to the EpisodeCount field.
func (o *SeasonStatisticsResource) SetEpisodeCount(v int32) {
	o.EpisodeCount = &v
}

// GetTotalEpisodeCount returns the TotalEpisodeCount field value if set, zero value otherwise.
func (o *SeasonStatisticsResource) GetTotalEpisodeCount() int32 {
	if o == nil || isNil(o.TotalEpisodeCount) {
		var ret int32
		return ret
	}
	return *o.TotalEpisodeCount
}

// GetTotalEpisodeCountOk returns a tuple with the TotalEpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonStatisticsResource) GetTotalEpisodeCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalEpisodeCount) {
    return nil, false
	}
	return o.TotalEpisodeCount, true
}

// HasTotalEpisodeCount returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasTotalEpisodeCount() bool {
	if o != nil && !isNil(o.TotalEpisodeCount) {
		return true
	}

	return false
}

// SetTotalEpisodeCount gets a reference to the given int32 and assigns it to the TotalEpisodeCount field.
func (o *SeasonStatisticsResource) SetTotalEpisodeCount(v int32) {
	o.TotalEpisodeCount = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *SeasonStatisticsResource) GetSizeOnDisk() int64 {
	if o == nil || isNil(o.SizeOnDisk) {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonStatisticsResource) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || isNil(o.SizeOnDisk) {
    return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasSizeOnDisk() bool {
	if o != nil && !isNil(o.SizeOnDisk) {
		return true
	}

	return false
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *SeasonStatisticsResource) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetReleaseGroups returns the ReleaseGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeasonStatisticsResource) GetReleaseGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ReleaseGroups
}

// GetReleaseGroupsOk returns a tuple with the ReleaseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeasonStatisticsResource) GetReleaseGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.ReleaseGroups) {
    return nil, false
	}
	return o.ReleaseGroups, true
}

// HasReleaseGroups returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasReleaseGroups() bool {
	if o != nil && isNil(o.ReleaseGroups) {
		return true
	}

	return false
}

// SetReleaseGroups gets a reference to the given []string and assigns it to the ReleaseGroups field.
func (o *SeasonStatisticsResource) SetReleaseGroups(v []string) {
	o.ReleaseGroups = v
}

// GetPercentOfEpisodes returns the PercentOfEpisodes field value if set, zero value otherwise.
func (o *SeasonStatisticsResource) GetPercentOfEpisodes() float64 {
	if o == nil || isNil(o.PercentOfEpisodes) {
		var ret float64
		return ret
	}
	return *o.PercentOfEpisodes
}

// GetPercentOfEpisodesOk returns a tuple with the PercentOfEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonStatisticsResource) GetPercentOfEpisodesOk() (*float64, bool) {
	if o == nil || isNil(o.PercentOfEpisodes) {
    return nil, false
	}
	return o.PercentOfEpisodes, true
}

// HasPercentOfEpisodes returns a boolean if a field has been set.
func (o *SeasonStatisticsResource) HasPercentOfEpisodes() bool {
	if o != nil && !isNil(o.PercentOfEpisodes) {
		return true
	}

	return false
}

// SetPercentOfEpisodes gets a reference to the given float64 and assigns it to the PercentOfEpisodes field.
func (o *SeasonStatisticsResource) SetPercentOfEpisodes(v float64) {
	o.PercentOfEpisodes = &v
}

func (o SeasonStatisticsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextAiring.IsSet() {
		toSerialize["nextAiring"] = o.NextAiring.Get()
	}
	if o.PreviousAiring.IsSet() {
		toSerialize["previousAiring"] = o.PreviousAiring.Get()
	}
	if !isNil(o.EpisodeFileCount) {
		toSerialize["episodeFileCount"] = o.EpisodeFileCount
	}
	if !isNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	if !isNil(o.TotalEpisodeCount) {
		toSerialize["totalEpisodeCount"] = o.TotalEpisodeCount
	}
	if !isNil(o.SizeOnDisk) {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if o.ReleaseGroups != nil {
		toSerialize["releaseGroups"] = o.ReleaseGroups
	}
	if !isNil(o.PercentOfEpisodes) {
		toSerialize["percentOfEpisodes"] = o.PercentOfEpisodes
	}
	return json.Marshal(toSerialize)
}

type NullableSeasonStatisticsResource struct {
	value *SeasonStatisticsResource
	isSet bool
}

func (v NullableSeasonStatisticsResource) Get() *SeasonStatisticsResource {
	return v.value
}

func (v *NullableSeasonStatisticsResource) Set(val *SeasonStatisticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeasonStatisticsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeasonStatisticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeasonStatisticsResource(val *SeasonStatisticsResource) *NullableSeasonStatisticsResource {
	return &NullableSeasonStatisticsResource{value: val, isSet: true}
}

func (v NullableSeasonStatisticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeasonStatisticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


