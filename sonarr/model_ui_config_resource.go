/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// UiConfigResource struct for UiConfigResource
type UiConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	FirstDayOfWeek *int32 `json:"firstDayOfWeek,omitempty"`
	CalendarWeekColumnHeader NullableString `json:"calendarWeekColumnHeader,omitempty"`
	ShortDateFormat NullableString `json:"shortDateFormat,omitempty"`
	LongDateFormat NullableString `json:"longDateFormat,omitempty"`
	TimeFormat NullableString `json:"timeFormat,omitempty"`
	ShowRelativeDates *bool `json:"showRelativeDates,omitempty"`
	EnableColorImpairedMode *bool `json:"enableColorImpairedMode,omitempty"`
	Theme NullableString `json:"theme,omitempty"`
	UiLanguage *int32 `json:"uiLanguage,omitempty"`
}

// NewUiConfigResource instantiates a new UiConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiConfigResource() *UiConfigResource {
	this := UiConfigResource{}
	return &this
}

// NewUiConfigResourceWithDefaults instantiates a new UiConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiConfigResourceWithDefaults() *UiConfigResource {
	this := UiConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UiConfigResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UiConfigResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UiConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetFirstDayOfWeek returns the FirstDayOfWeek field value if set, zero value otherwise.
func (o *UiConfigResource) GetFirstDayOfWeek() int32 {
	if o == nil || isNil(o.FirstDayOfWeek) {
		var ret int32
		return ret
	}
	return *o.FirstDayOfWeek
}

// GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetFirstDayOfWeekOk() (*int32, bool) {
	if o == nil || isNil(o.FirstDayOfWeek) {
    return nil, false
	}
	return o.FirstDayOfWeek, true
}

// HasFirstDayOfWeek returns a boolean if a field has been set.
func (o *UiConfigResource) HasFirstDayOfWeek() bool {
	if o != nil && !isNil(o.FirstDayOfWeek) {
		return true
	}

	return false
}

// SetFirstDayOfWeek gets a reference to the given int32 and assigns it to the FirstDayOfWeek field.
func (o *UiConfigResource) SetFirstDayOfWeek(v int32) {
	o.FirstDayOfWeek = &v
}

// GetCalendarWeekColumnHeader returns the CalendarWeekColumnHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetCalendarWeekColumnHeader() string {
	if o == nil || isNil(o.CalendarWeekColumnHeader.Get()) {
		var ret string
		return ret
	}
	return *o.CalendarWeekColumnHeader.Get()
}

// GetCalendarWeekColumnHeaderOk returns a tuple with the CalendarWeekColumnHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetCalendarWeekColumnHeaderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CalendarWeekColumnHeader.Get(), o.CalendarWeekColumnHeader.IsSet()
}

// HasCalendarWeekColumnHeader returns a boolean if a field has been set.
func (o *UiConfigResource) HasCalendarWeekColumnHeader() bool {
	if o != nil && o.CalendarWeekColumnHeader.IsSet() {
		return true
	}

	return false
}

// SetCalendarWeekColumnHeader gets a reference to the given NullableString and assigns it to the CalendarWeekColumnHeader field.
func (o *UiConfigResource) SetCalendarWeekColumnHeader(v string) {
	o.CalendarWeekColumnHeader.Set(&v)
}
// SetCalendarWeekColumnHeaderNil sets the value for CalendarWeekColumnHeader to be an explicit nil
func (o *UiConfigResource) SetCalendarWeekColumnHeaderNil() {
	o.CalendarWeekColumnHeader.Set(nil)
}

// UnsetCalendarWeekColumnHeader ensures that no value is present for CalendarWeekColumnHeader, not even an explicit nil
func (o *UiConfigResource) UnsetCalendarWeekColumnHeader() {
	o.CalendarWeekColumnHeader.Unset()
}

// GetShortDateFormat returns the ShortDateFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetShortDateFormat() string {
	if o == nil || isNil(o.ShortDateFormat.Get()) {
		var ret string
		return ret
	}
	return *o.ShortDateFormat.Get()
}

// GetShortDateFormatOk returns a tuple with the ShortDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetShortDateFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ShortDateFormat.Get(), o.ShortDateFormat.IsSet()
}

// HasShortDateFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasShortDateFormat() bool {
	if o != nil && o.ShortDateFormat.IsSet() {
		return true
	}

	return false
}

// SetShortDateFormat gets a reference to the given NullableString and assigns it to the ShortDateFormat field.
func (o *UiConfigResource) SetShortDateFormat(v string) {
	o.ShortDateFormat.Set(&v)
}
// SetShortDateFormatNil sets the value for ShortDateFormat to be an explicit nil
func (o *UiConfigResource) SetShortDateFormatNil() {
	o.ShortDateFormat.Set(nil)
}

// UnsetShortDateFormat ensures that no value is present for ShortDateFormat, not even an explicit nil
func (o *UiConfigResource) UnsetShortDateFormat() {
	o.ShortDateFormat.Unset()
}

// GetLongDateFormat returns the LongDateFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetLongDateFormat() string {
	if o == nil || isNil(o.LongDateFormat.Get()) {
		var ret string
		return ret
	}
	return *o.LongDateFormat.Get()
}

// GetLongDateFormatOk returns a tuple with the LongDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetLongDateFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.LongDateFormat.Get(), o.LongDateFormat.IsSet()
}

// HasLongDateFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasLongDateFormat() bool {
	if o != nil && o.LongDateFormat.IsSet() {
		return true
	}

	return false
}

// SetLongDateFormat gets a reference to the given NullableString and assigns it to the LongDateFormat field.
func (o *UiConfigResource) SetLongDateFormat(v string) {
	o.LongDateFormat.Set(&v)
}
// SetLongDateFormatNil sets the value for LongDateFormat to be an explicit nil
func (o *UiConfigResource) SetLongDateFormatNil() {
	o.LongDateFormat.Set(nil)
}

// UnsetLongDateFormat ensures that no value is present for LongDateFormat, not even an explicit nil
func (o *UiConfigResource) UnsetLongDateFormat() {
	o.LongDateFormat.Unset()
}

// GetTimeFormat returns the TimeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetTimeFormat() string {
	if o == nil || isNil(o.TimeFormat.Get()) {
		var ret string
		return ret
	}
	return *o.TimeFormat.Get()
}

// GetTimeFormatOk returns a tuple with the TimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetTimeFormatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TimeFormat.Get(), o.TimeFormat.IsSet()
}

// HasTimeFormat returns a boolean if a field has been set.
func (o *UiConfigResource) HasTimeFormat() bool {
	if o != nil && o.TimeFormat.IsSet() {
		return true
	}

	return false
}

// SetTimeFormat gets a reference to the given NullableString and assigns it to the TimeFormat field.
func (o *UiConfigResource) SetTimeFormat(v string) {
	o.TimeFormat.Set(&v)
}
// SetTimeFormatNil sets the value for TimeFormat to be an explicit nil
func (o *UiConfigResource) SetTimeFormatNil() {
	o.TimeFormat.Set(nil)
}

// UnsetTimeFormat ensures that no value is present for TimeFormat, not even an explicit nil
func (o *UiConfigResource) UnsetTimeFormat() {
	o.TimeFormat.Unset()
}

// GetShowRelativeDates returns the ShowRelativeDates field value if set, zero value otherwise.
func (o *UiConfigResource) GetShowRelativeDates() bool {
	if o == nil || isNil(o.ShowRelativeDates) {
		var ret bool
		return ret
	}
	return *o.ShowRelativeDates
}

// GetShowRelativeDatesOk returns a tuple with the ShowRelativeDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetShowRelativeDatesOk() (*bool, bool) {
	if o == nil || isNil(o.ShowRelativeDates) {
    return nil, false
	}
	return o.ShowRelativeDates, true
}

// HasShowRelativeDates returns a boolean if a field has been set.
func (o *UiConfigResource) HasShowRelativeDates() bool {
	if o != nil && !isNil(o.ShowRelativeDates) {
		return true
	}

	return false
}

// SetShowRelativeDates gets a reference to the given bool and assigns it to the ShowRelativeDates field.
func (o *UiConfigResource) SetShowRelativeDates(v bool) {
	o.ShowRelativeDates = &v
}

// GetEnableColorImpairedMode returns the EnableColorImpairedMode field value if set, zero value otherwise.
func (o *UiConfigResource) GetEnableColorImpairedMode() bool {
	if o == nil || isNil(o.EnableColorImpairedMode) {
		var ret bool
		return ret
	}
	return *o.EnableColorImpairedMode
}

// GetEnableColorImpairedModeOk returns a tuple with the EnableColorImpairedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetEnableColorImpairedModeOk() (*bool, bool) {
	if o == nil || isNil(o.EnableColorImpairedMode) {
    return nil, false
	}
	return o.EnableColorImpairedMode, true
}

// HasEnableColorImpairedMode returns a boolean if a field has been set.
func (o *UiConfigResource) HasEnableColorImpairedMode() bool {
	if o != nil && !isNil(o.EnableColorImpairedMode) {
		return true
	}

	return false
}

// SetEnableColorImpairedMode gets a reference to the given bool and assigns it to the EnableColorImpairedMode field.
func (o *UiConfigResource) SetEnableColorImpairedMode(v bool) {
	o.EnableColorImpairedMode = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiConfigResource) GetTheme() string {
	if o == nil || isNil(o.Theme.Get()) {
		var ret string
		return ret
	}
	return *o.Theme.Get()
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiConfigResource) GetThemeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Theme.Get(), o.Theme.IsSet()
}

// HasTheme returns a boolean if a field has been set.
func (o *UiConfigResource) HasTheme() bool {
	if o != nil && o.Theme.IsSet() {
		return true
	}

	return false
}

// SetTheme gets a reference to the given NullableString and assigns it to the Theme field.
func (o *UiConfigResource) SetTheme(v string) {
	o.Theme.Set(&v)
}
// SetThemeNil sets the value for Theme to be an explicit nil
func (o *UiConfigResource) SetThemeNil() {
	o.Theme.Set(nil)
}

// UnsetTheme ensures that no value is present for Theme, not even an explicit nil
func (o *UiConfigResource) UnsetTheme() {
	o.Theme.Unset()
}

// GetUiLanguage returns the UiLanguage field value if set, zero value otherwise.
func (o *UiConfigResource) GetUiLanguage() int32 {
	if o == nil || isNil(o.UiLanguage) {
		var ret int32
		return ret
	}
	return *o.UiLanguage
}

// GetUiLanguageOk returns a tuple with the UiLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfigResource) GetUiLanguageOk() (*int32, bool) {
	if o == nil || isNil(o.UiLanguage) {
    return nil, false
	}
	return o.UiLanguage, true
}

// HasUiLanguage returns a boolean if a field has been set.
func (o *UiConfigResource) HasUiLanguage() bool {
	if o != nil && !isNil(o.UiLanguage) {
		return true
	}

	return false
}

// SetUiLanguage gets a reference to the given int32 and assigns it to the UiLanguage field.
func (o *UiConfigResource) SetUiLanguage(v int32) {
	o.UiLanguage = &v
}

func (o UiConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.FirstDayOfWeek) {
		toSerialize["firstDayOfWeek"] = o.FirstDayOfWeek
	}
	if o.CalendarWeekColumnHeader.IsSet() {
		toSerialize["calendarWeekColumnHeader"] = o.CalendarWeekColumnHeader.Get()
	}
	if o.ShortDateFormat.IsSet() {
		toSerialize["shortDateFormat"] = o.ShortDateFormat.Get()
	}
	if o.LongDateFormat.IsSet() {
		toSerialize["longDateFormat"] = o.LongDateFormat.Get()
	}
	if o.TimeFormat.IsSet() {
		toSerialize["timeFormat"] = o.TimeFormat.Get()
	}
	if !isNil(o.ShowRelativeDates) {
		toSerialize["showRelativeDates"] = o.ShowRelativeDates
	}
	if !isNil(o.EnableColorImpairedMode) {
		toSerialize["enableColorImpairedMode"] = o.EnableColorImpairedMode
	}
	if o.Theme.IsSet() {
		toSerialize["theme"] = o.Theme.Get()
	}
	if !isNil(o.UiLanguage) {
		toSerialize["uiLanguage"] = o.UiLanguage
	}
	return json.Marshal(toSerialize)
}

type NullableUiConfigResource struct {
	value *UiConfigResource
	isSet bool
}

func (v NullableUiConfigResource) Get() *UiConfigResource {
	return v.value
}

func (v *NullableUiConfigResource) Set(val *UiConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableUiConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableUiConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiConfigResource(val *UiConfigResource) *NullableUiConfigResource {
	return &NullableUiConfigResource{value: val, isSet: true}
}

func (v NullableUiConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


