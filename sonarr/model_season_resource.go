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

// SeasonResource struct for SeasonResource
type SeasonResource struct {
	SeasonNumber *int32 `json:"seasonNumber,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	Statistics *SeasonStatisticsResource `json:"statistics,omitempty"`
	Images []*MediaCover `json:"images,omitempty"`
}

// NewSeasonResource instantiates a new SeasonResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeasonResource() *SeasonResource {
	this := SeasonResource{}
	return &this
}

// NewSeasonResourceWithDefaults instantiates a new SeasonResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeasonResourceWithDefaults() *SeasonResource {
	this := SeasonResource{}
	return &this
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *SeasonResource) GetSeasonNumber() int32 {
	if o == nil || isNil(o.SeasonNumber) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonResource) GetSeasonNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SeasonNumber) {
    return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *SeasonResource) HasSeasonNumber() bool {
	if o != nil && !isNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given int32 and assigns it to the SeasonNumber field.
func (o *SeasonResource) SetSeasonNumber(v int32) {
	o.SeasonNumber = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *SeasonResource) GetMonitored() bool {
	if o == nil || isNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonResource) GetMonitoredOk() (*bool, bool) {
	if o == nil || isNil(o.Monitored) {
    return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *SeasonResource) HasMonitored() bool {
	if o != nil && !isNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *SeasonResource) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *SeasonResource) GetStatistics() SeasonStatisticsResource {
	if o == nil || isNil(o.Statistics) {
		var ret SeasonStatisticsResource
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonResource) GetStatisticsOk() (*SeasonStatisticsResource, bool) {
	if o == nil || isNil(o.Statistics) {
    return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *SeasonResource) HasStatistics() bool {
	if o != nil && !isNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given SeasonStatisticsResource and assigns it to the Statistics field.
func (o *SeasonResource) SetStatistics(v SeasonStatisticsResource) {
	o.Statistics = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeasonResource) GetImages() []*MediaCover {
	if o == nil {
		var ret []*MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeasonResource) GetImagesOk() ([]*MediaCover, bool) {
	if o == nil || isNil(o.Images) {
    return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *SeasonResource) HasImages() bool {
	if o != nil && isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *SeasonResource) SetImages(v []*MediaCover) {
	o.Images = v
}

func (o SeasonResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !isNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !isNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableSeasonResource struct {
	value *SeasonResource
	isSet bool
}

func (v NullableSeasonResource) Get() *SeasonResource {
	return v.value
}

func (v *NullableSeasonResource) Set(val *SeasonResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeasonResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeasonResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeasonResource(val *SeasonResource) *NullableSeasonResource {
	return &NullableSeasonResource{value: val, isSet: true}
}

func (v NullableSeasonResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeasonResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


