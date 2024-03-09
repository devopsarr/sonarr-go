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

// checks if the SeasonPassResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeasonPassResource{}

// SeasonPassResource struct for SeasonPassResource
type SeasonPassResource struct {
	Series []SeasonPassSeriesResource `json:"series,omitempty"`
	MonitoringOptions *MonitoringOptions `json:"monitoringOptions,omitempty"`
}

// NewSeasonPassResource instantiates a new SeasonPassResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeasonPassResource() *SeasonPassResource {
	this := SeasonPassResource{}
	return &this
}

// NewSeasonPassResourceWithDefaults instantiates a new SeasonPassResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeasonPassResourceWithDefaults() *SeasonPassResource {
	this := SeasonPassResource{}
	return &this
}

// GetSeries returns the Series field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeasonPassResource) GetSeries() []SeasonPassSeriesResource {
	if o == nil {
		var ret []SeasonPassSeriesResource
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeasonPassResource) GetSeriesOk() ([]SeasonPassSeriesResource, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *SeasonPassResource) HasSeries() bool {
	if o != nil && IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []SeasonPassSeriesResource and assigns it to the Series field.
func (o *SeasonPassResource) SetSeries(v []SeasonPassSeriesResource) {
	o.Series = v
}

// GetMonitoringOptions returns the MonitoringOptions field value if set, zero value otherwise.
func (o *SeasonPassResource) GetMonitoringOptions() MonitoringOptions {
	if o == nil || IsNil(o.MonitoringOptions) {
		var ret MonitoringOptions
		return ret
	}
	return *o.MonitoringOptions
}

// GetMonitoringOptionsOk returns a tuple with the MonitoringOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeasonPassResource) GetMonitoringOptionsOk() (*MonitoringOptions, bool) {
	if o == nil || IsNil(o.MonitoringOptions) {
		return nil, false
	}
	return o.MonitoringOptions, true
}

// HasMonitoringOptions returns a boolean if a field has been set.
func (o *SeasonPassResource) HasMonitoringOptions() bool {
	if o != nil && !IsNil(o.MonitoringOptions) {
		return true
	}

	return false
}

// SetMonitoringOptions gets a reference to the given MonitoringOptions and assigns it to the MonitoringOptions field.
func (o *SeasonPassResource) SetMonitoringOptions(v MonitoringOptions) {
	o.MonitoringOptions = &v
}

func (o SeasonPassResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeasonPassResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.MonitoringOptions) {
		toSerialize["monitoringOptions"] = o.MonitoringOptions
	}
	return toSerialize, nil
}

type NullableSeasonPassResource struct {
	value *SeasonPassResource
	isSet bool
}

func (v NullableSeasonPassResource) Get() *SeasonPassResource {
	return v.value
}

func (v *NullableSeasonPassResource) Set(val *SeasonPassResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeasonPassResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeasonPassResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeasonPassResource(val *SeasonPassResource) *NullableSeasonPassResource {
	return &NullableSeasonPassResource{value: val, isSet: true}
}

func (v NullableSeasonPassResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeasonPassResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


