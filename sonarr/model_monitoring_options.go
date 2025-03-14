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

// checks if the MonitoringOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringOptions{}

// MonitoringOptions struct for MonitoringOptions
type MonitoringOptions struct {
	IgnoreEpisodesWithFiles *bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles *bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
	Monitor *MonitorTypes `json:"monitor,omitempty"`
}

// NewMonitoringOptions instantiates a new MonitoringOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringOptions() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// NewMonitoringOptionsWithDefaults instantiates a new MonitoringOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringOptionsWithDefaults() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// GetIgnoreEpisodesWithFiles returns the IgnoreEpisodesWithFiles field value if set, zero value otherwise.
func (o *MonitoringOptions) GetIgnoreEpisodesWithFiles() bool {
	if o == nil || IsNil(o.IgnoreEpisodesWithFiles) {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithFiles
}

// GetIgnoreEpisodesWithFilesOk returns a tuple with the IgnoreEpisodesWithFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetIgnoreEpisodesWithFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreEpisodesWithFiles) {
		return nil, false
	}
	return o.IgnoreEpisodesWithFiles, true
}

// HasIgnoreEpisodesWithFiles returns a boolean if a field has been set.
func (o *MonitoringOptions) HasIgnoreEpisodesWithFiles() bool {
	if o != nil && !IsNil(o.IgnoreEpisodesWithFiles) {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithFiles field.
func (o *MonitoringOptions) SetIgnoreEpisodesWithFiles(v bool) {
	o.IgnoreEpisodesWithFiles = &v
}

// GetIgnoreEpisodesWithoutFiles returns the IgnoreEpisodesWithoutFiles field value if set, zero value otherwise.
func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFiles() bool {
	if o == nil || IsNil(o.IgnoreEpisodesWithoutFiles) {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithoutFiles
}

// GetIgnoreEpisodesWithoutFilesOk returns a tuple with the IgnoreEpisodesWithoutFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreEpisodesWithoutFiles) {
		return nil, false
	}
	return o.IgnoreEpisodesWithoutFiles, true
}

// HasIgnoreEpisodesWithoutFiles returns a boolean if a field has been set.
func (o *MonitoringOptions) HasIgnoreEpisodesWithoutFiles() bool {
	if o != nil && !IsNil(o.IgnoreEpisodesWithoutFiles) {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithoutFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithoutFiles field.
func (o *MonitoringOptions) SetIgnoreEpisodesWithoutFiles(v bool) {
	o.IgnoreEpisodesWithoutFiles = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MonitoringOptions) GetMonitor() MonitorTypes {
	if o == nil || IsNil(o.Monitor) {
		var ret MonitorTypes
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetMonitorOk() (*MonitorTypes, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MonitoringOptions) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorTypes and assigns it to the Monitor field.
func (o *MonitoringOptions) SetMonitor(v MonitorTypes) {
	o.Monitor = &v
}

func (o MonitoringOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoreEpisodesWithFiles) {
		toSerialize["ignoreEpisodesWithFiles"] = o.IgnoreEpisodesWithFiles
	}
	if !IsNil(o.IgnoreEpisodesWithoutFiles) {
		toSerialize["ignoreEpisodesWithoutFiles"] = o.IgnoreEpisodesWithoutFiles
	}
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	return toSerialize, nil
}

type NullableMonitoringOptions struct {
	value *MonitoringOptions
	isSet bool
}

func (v NullableMonitoringOptions) Get() *MonitoringOptions {
	return v.value
}

func (v *NullableMonitoringOptions) Set(val *MonitoringOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringOptions(val *MonitoringOptions) *NullableMonitoringOptions {
	return &NullableMonitoringOptions{value: val, isSet: true}
}

func (v NullableMonitoringOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


