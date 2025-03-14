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

// checks if the SeriesTitleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesTitleInfo{}

// SeriesTitleInfo struct for SeriesTitleInfo
type SeriesTitleInfo struct {
	Title NullableString `json:"title,omitempty"`
	TitleWithoutYear NullableString `json:"titleWithoutYear,omitempty"`
	Year *int32 `json:"year,omitempty"`
	AllTitles []string `json:"allTitles,omitempty"`
}

// NewSeriesTitleInfo instantiates a new SeriesTitleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesTitleInfo() *SeriesTitleInfo {
	this := SeriesTitleInfo{}
	return &this
}

// NewSeriesTitleInfoWithDefaults instantiates a new SeriesTitleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesTitleInfoWithDefaults() *SeriesTitleInfo {
	this := SeriesTitleInfo{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesTitleInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesTitleInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SeriesTitleInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SeriesTitleInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SeriesTitleInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SeriesTitleInfo) UnsetTitle() {
	o.Title.Unset()
}

// GetTitleWithoutYear returns the TitleWithoutYear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesTitleInfo) GetTitleWithoutYear() string {
	if o == nil || IsNil(o.TitleWithoutYear.Get()) {
		var ret string
		return ret
	}
	return *o.TitleWithoutYear.Get()
}

// GetTitleWithoutYearOk returns a tuple with the TitleWithoutYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesTitleInfo) GetTitleWithoutYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleWithoutYear.Get(), o.TitleWithoutYear.IsSet()
}

// HasTitleWithoutYear returns a boolean if a field has been set.
func (o *SeriesTitleInfo) HasTitleWithoutYear() bool {
	if o != nil && o.TitleWithoutYear.IsSet() {
		return true
	}

	return false
}

// SetTitleWithoutYear gets a reference to the given NullableString and assigns it to the TitleWithoutYear field.
func (o *SeriesTitleInfo) SetTitleWithoutYear(v string) {
	o.TitleWithoutYear.Set(&v)
}
// SetTitleWithoutYearNil sets the value for TitleWithoutYear to be an explicit nil
func (o *SeriesTitleInfo) SetTitleWithoutYearNil() {
	o.TitleWithoutYear.Set(nil)
}

// UnsetTitleWithoutYear ensures that no value is present for TitleWithoutYear, not even an explicit nil
func (o *SeriesTitleInfo) UnsetTitleWithoutYear() {
	o.TitleWithoutYear.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SeriesTitleInfo) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesTitleInfo) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SeriesTitleInfo) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *SeriesTitleInfo) SetYear(v int32) {
	o.Year = &v
}

// GetAllTitles returns the AllTitles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesTitleInfo) GetAllTitles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllTitles
}

// GetAllTitlesOk returns a tuple with the AllTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesTitleInfo) GetAllTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllTitles) {
		return nil, false
	}
	return o.AllTitles, true
}

// HasAllTitles returns a boolean if a field has been set.
func (o *SeriesTitleInfo) HasAllTitles() bool {
	if o != nil && !IsNil(o.AllTitles) {
		return true
	}

	return false
}

// SetAllTitles gets a reference to the given []string and assigns it to the AllTitles field.
func (o *SeriesTitleInfo) SetAllTitles(v []string) {
	o.AllTitles = v
}

func (o SeriesTitleInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesTitleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.TitleWithoutYear.IsSet() {
		toSerialize["titleWithoutYear"] = o.TitleWithoutYear.Get()
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if o.AllTitles != nil {
		toSerialize["allTitles"] = o.AllTitles
	}
	return toSerialize, nil
}

type NullableSeriesTitleInfo struct {
	value *SeriesTitleInfo
	isSet bool
}

func (v NullableSeriesTitleInfo) Get() *SeriesTitleInfo {
	return v.value
}

func (v *NullableSeriesTitleInfo) Set(val *SeriesTitleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTitleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTitleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTitleInfo(val *SeriesTitleInfo) *NullableSeriesTitleInfo {
	return &NullableSeriesTitleInfo{value: val, isSet: true}
}

func (v NullableSeriesTitleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTitleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


