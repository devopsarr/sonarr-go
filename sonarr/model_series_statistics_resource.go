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

// checks if the SeriesStatisticsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesStatisticsResource{}

// SeriesStatisticsResource struct for SeriesStatisticsResource
type SeriesStatisticsResource struct {
	SeasonCount *int32 `json:"seasonCount,omitempty"`
	EpisodeFileCount *int32 `json:"episodeFileCount,omitempty"`
	EpisodeCount *int32 `json:"episodeCount,omitempty"`
	TotalEpisodeCount *int32 `json:"totalEpisodeCount,omitempty"`
	SizeOnDisk *int64 `json:"sizeOnDisk,omitempty"`
	ReleaseGroups []string `json:"releaseGroups,omitempty"`
	PercentOfEpisodes *float64 `json:"percentOfEpisodes,omitempty"`
}

// NewSeriesStatisticsResource instantiates a new SeriesStatisticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesStatisticsResource() *SeriesStatisticsResource {
	this := SeriesStatisticsResource{}
	return &this
}

// NewSeriesStatisticsResourceWithDefaults instantiates a new SeriesStatisticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesStatisticsResourceWithDefaults() *SeriesStatisticsResource {
	this := SeriesStatisticsResource{}
	return &this
}

// GetSeasonCount returns the SeasonCount field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetSeasonCount() int32 {
	if o == nil || IsNil(o.SeasonCount) {
		var ret int32
		return ret
	}
	return *o.SeasonCount
}

// GetSeasonCountOk returns a tuple with the SeasonCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetSeasonCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeasonCount) {
		return nil, false
	}
	return o.SeasonCount, true
}

// HasSeasonCount returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasSeasonCount() bool {
	if o != nil && !IsNil(o.SeasonCount) {
		return true
	}

	return false
}

// SetSeasonCount gets a reference to the given int32 and assigns it to the SeasonCount field.
func (o *SeriesStatisticsResource) SetSeasonCount(v int32) {
	o.SeasonCount = &v
}

// GetEpisodeFileCount returns the EpisodeFileCount field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetEpisodeFileCount() int32 {
	if o == nil || IsNil(o.EpisodeFileCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeFileCount
}

// GetEpisodeFileCountOk returns a tuple with the EpisodeFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetEpisodeFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EpisodeFileCount) {
		return nil, false
	}
	return o.EpisodeFileCount, true
}

// HasEpisodeFileCount returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasEpisodeFileCount() bool {
	if o != nil && !IsNil(o.EpisodeFileCount) {
		return true
	}

	return false
}

// SetEpisodeFileCount gets a reference to the given int32 and assigns it to the EpisodeFileCount field.
func (o *SeriesStatisticsResource) SetEpisodeFileCount(v int32) {
	o.EpisodeFileCount = &v
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetEpisodeCount() int32 {
	if o == nil || IsNil(o.EpisodeCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetEpisodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EpisodeCount) {
		return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasEpisodeCount() bool {
	if o != nil && !IsNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given int32 and assigns it to the EpisodeCount field.
func (o *SeriesStatisticsResource) SetEpisodeCount(v int32) {
	o.EpisodeCount = &v
}

// GetTotalEpisodeCount returns the TotalEpisodeCount field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetTotalEpisodeCount() int32 {
	if o == nil || IsNil(o.TotalEpisodeCount) {
		var ret int32
		return ret
	}
	return *o.TotalEpisodeCount
}

// GetTotalEpisodeCountOk returns a tuple with the TotalEpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetTotalEpisodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalEpisodeCount) {
		return nil, false
	}
	return o.TotalEpisodeCount, true
}

// HasTotalEpisodeCount returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasTotalEpisodeCount() bool {
	if o != nil && !IsNil(o.TotalEpisodeCount) {
		return true
	}

	return false
}

// SetTotalEpisodeCount gets a reference to the given int32 and assigns it to the TotalEpisodeCount field.
func (o *SeriesStatisticsResource) SetTotalEpisodeCount(v int32) {
	o.TotalEpisodeCount = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetSizeOnDisk() int64 {
	if o == nil || IsNil(o.SizeOnDisk) {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || IsNil(o.SizeOnDisk) {
		return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasSizeOnDisk() bool {
	if o != nil && !IsNil(o.SizeOnDisk) {
		return true
	}

	return false
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *SeriesStatisticsResource) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetReleaseGroups returns the ReleaseGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesStatisticsResource) GetReleaseGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ReleaseGroups
}

// GetReleaseGroupsOk returns a tuple with the ReleaseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesStatisticsResource) GetReleaseGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReleaseGroups) {
		return nil, false
	}
	return o.ReleaseGroups, true
}

// HasReleaseGroups returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasReleaseGroups() bool {
	if o != nil && !IsNil(o.ReleaseGroups) {
		return true
	}

	return false
}

// SetReleaseGroups gets a reference to the given []string and assigns it to the ReleaseGroups field.
func (o *SeriesStatisticsResource) SetReleaseGroups(v []string) {
	o.ReleaseGroups = v
}

// GetPercentOfEpisodes returns the PercentOfEpisodes field value if set, zero value otherwise.
func (o *SeriesStatisticsResource) GetPercentOfEpisodes() float64 {
	if o == nil || IsNil(o.PercentOfEpisodes) {
		var ret float64
		return ret
	}
	return *o.PercentOfEpisodes
}

// GetPercentOfEpisodesOk returns a tuple with the PercentOfEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesStatisticsResource) GetPercentOfEpisodesOk() (*float64, bool) {
	if o == nil || IsNil(o.PercentOfEpisodes) {
		return nil, false
	}
	return o.PercentOfEpisodes, true
}

// HasPercentOfEpisodes returns a boolean if a field has been set.
func (o *SeriesStatisticsResource) HasPercentOfEpisodes() bool {
	if o != nil && !IsNil(o.PercentOfEpisodes) {
		return true
	}

	return false
}

// SetPercentOfEpisodes gets a reference to the given float64 and assigns it to the PercentOfEpisodes field.
func (o *SeriesStatisticsResource) SetPercentOfEpisodes(v float64) {
	o.PercentOfEpisodes = &v
}

func (o SeriesStatisticsResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesStatisticsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SeasonCount) {
		toSerialize["seasonCount"] = o.SeasonCount
	}
	if !IsNil(o.EpisodeFileCount) {
		toSerialize["episodeFileCount"] = o.EpisodeFileCount
	}
	if !IsNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	if !IsNil(o.TotalEpisodeCount) {
		toSerialize["totalEpisodeCount"] = o.TotalEpisodeCount
	}
	if !IsNil(o.SizeOnDisk) {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if o.ReleaseGroups != nil {
		toSerialize["releaseGroups"] = o.ReleaseGroups
	}
	if !IsNil(o.PercentOfEpisodes) {
		toSerialize["percentOfEpisodes"] = o.PercentOfEpisodes
	}
	return toSerialize, nil
}

type NullableSeriesStatisticsResource struct {
	value *SeriesStatisticsResource
	isSet bool
}

func (v NullableSeriesStatisticsResource) Get() *SeriesStatisticsResource {
	return v.value
}

func (v *NullableSeriesStatisticsResource) Set(val *SeriesStatisticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesStatisticsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesStatisticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesStatisticsResource(val *SeriesStatisticsResource) *NullableSeriesStatisticsResource {
	return &NullableSeriesStatisticsResource{value: val, isSet: true}
}

func (v NullableSeriesStatisticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesStatisticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


