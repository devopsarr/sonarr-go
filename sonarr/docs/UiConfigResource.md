# UiConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FirstDayOfWeek** | Pointer to **int32** |  | [optional] 
**CalendarWeekColumnHeader** | Pointer to **NullableString** |  | [optional] 
**ShortDateFormat** | Pointer to **NullableString** |  | [optional] 
**LongDateFormat** | Pointer to **NullableString** |  | [optional] 
**TimeFormat** | Pointer to **NullableString** |  | [optional] 
**ShowRelativeDates** | Pointer to **bool** |  | [optional] 
**EnableColorImpairedMode** | Pointer to **bool** |  | [optional] 
**Theme** | Pointer to **NullableString** |  | [optional] 
**UiLanguage** | Pointer to **int32** |  | [optional] 

## Methods

### NewUiConfigResource

`func NewUiConfigResource() *UiConfigResource`

NewUiConfigResource instantiates a new UiConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiConfigResourceWithDefaults

`func NewUiConfigResourceWithDefaults() *UiConfigResource`

NewUiConfigResourceWithDefaults instantiates a new UiConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UiConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UiConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstDayOfWeek

`func (o *UiConfigResource) GetFirstDayOfWeek() int32`

GetFirstDayOfWeek returns the FirstDayOfWeek field if non-nil, zero value otherwise.

### GetFirstDayOfWeekOk

`func (o *UiConfigResource) GetFirstDayOfWeekOk() (*int32, bool)`

GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDayOfWeek

`func (o *UiConfigResource) SetFirstDayOfWeek(v int32)`

SetFirstDayOfWeek sets FirstDayOfWeek field to given value.

### HasFirstDayOfWeek

`func (o *UiConfigResource) HasFirstDayOfWeek() bool`

HasFirstDayOfWeek returns a boolean if a field has been set.

### GetCalendarWeekColumnHeader

`func (o *UiConfigResource) GetCalendarWeekColumnHeader() string`

GetCalendarWeekColumnHeader returns the CalendarWeekColumnHeader field if non-nil, zero value otherwise.

### GetCalendarWeekColumnHeaderOk

`func (o *UiConfigResource) GetCalendarWeekColumnHeaderOk() (*string, bool)`

GetCalendarWeekColumnHeaderOk returns a tuple with the CalendarWeekColumnHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarWeekColumnHeader

`func (o *UiConfigResource) SetCalendarWeekColumnHeader(v string)`

SetCalendarWeekColumnHeader sets CalendarWeekColumnHeader field to given value.

### HasCalendarWeekColumnHeader

`func (o *UiConfigResource) HasCalendarWeekColumnHeader() bool`

HasCalendarWeekColumnHeader returns a boolean if a field has been set.

### SetCalendarWeekColumnHeaderNil

`func (o *UiConfigResource) SetCalendarWeekColumnHeaderNil(b bool)`

 SetCalendarWeekColumnHeaderNil sets the value for CalendarWeekColumnHeader to be an explicit nil

### UnsetCalendarWeekColumnHeader
`func (o *UiConfigResource) UnsetCalendarWeekColumnHeader()`

UnsetCalendarWeekColumnHeader ensures that no value is present for CalendarWeekColumnHeader, not even an explicit nil
### GetShortDateFormat

`func (o *UiConfigResource) GetShortDateFormat() string`

GetShortDateFormat returns the ShortDateFormat field if non-nil, zero value otherwise.

### GetShortDateFormatOk

`func (o *UiConfigResource) GetShortDateFormatOk() (*string, bool)`

GetShortDateFormatOk returns a tuple with the ShortDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDateFormat

`func (o *UiConfigResource) SetShortDateFormat(v string)`

SetShortDateFormat sets ShortDateFormat field to given value.

### HasShortDateFormat

`func (o *UiConfigResource) HasShortDateFormat() bool`

HasShortDateFormat returns a boolean if a field has been set.

### SetShortDateFormatNil

`func (o *UiConfigResource) SetShortDateFormatNil(b bool)`

 SetShortDateFormatNil sets the value for ShortDateFormat to be an explicit nil

### UnsetShortDateFormat
`func (o *UiConfigResource) UnsetShortDateFormat()`

UnsetShortDateFormat ensures that no value is present for ShortDateFormat, not even an explicit nil
### GetLongDateFormat

`func (o *UiConfigResource) GetLongDateFormat() string`

GetLongDateFormat returns the LongDateFormat field if non-nil, zero value otherwise.

### GetLongDateFormatOk

`func (o *UiConfigResource) GetLongDateFormatOk() (*string, bool)`

GetLongDateFormatOk returns a tuple with the LongDateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDateFormat

`func (o *UiConfigResource) SetLongDateFormat(v string)`

SetLongDateFormat sets LongDateFormat field to given value.

### HasLongDateFormat

`func (o *UiConfigResource) HasLongDateFormat() bool`

HasLongDateFormat returns a boolean if a field has been set.

### SetLongDateFormatNil

`func (o *UiConfigResource) SetLongDateFormatNil(b bool)`

 SetLongDateFormatNil sets the value for LongDateFormat to be an explicit nil

### UnsetLongDateFormat
`func (o *UiConfigResource) UnsetLongDateFormat()`

UnsetLongDateFormat ensures that no value is present for LongDateFormat, not even an explicit nil
### GetTimeFormat

`func (o *UiConfigResource) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *UiConfigResource) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *UiConfigResource) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *UiConfigResource) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### SetTimeFormatNil

`func (o *UiConfigResource) SetTimeFormatNil(b bool)`

 SetTimeFormatNil sets the value for TimeFormat to be an explicit nil

### UnsetTimeFormat
`func (o *UiConfigResource) UnsetTimeFormat()`

UnsetTimeFormat ensures that no value is present for TimeFormat, not even an explicit nil
### GetShowRelativeDates

`func (o *UiConfigResource) GetShowRelativeDates() bool`

GetShowRelativeDates returns the ShowRelativeDates field if non-nil, zero value otherwise.

### GetShowRelativeDatesOk

`func (o *UiConfigResource) GetShowRelativeDatesOk() (*bool, bool)`

GetShowRelativeDatesOk returns a tuple with the ShowRelativeDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRelativeDates

`func (o *UiConfigResource) SetShowRelativeDates(v bool)`

SetShowRelativeDates sets ShowRelativeDates field to given value.

### HasShowRelativeDates

`func (o *UiConfigResource) HasShowRelativeDates() bool`

HasShowRelativeDates returns a boolean if a field has been set.

### GetEnableColorImpairedMode

`func (o *UiConfigResource) GetEnableColorImpairedMode() bool`

GetEnableColorImpairedMode returns the EnableColorImpairedMode field if non-nil, zero value otherwise.

### GetEnableColorImpairedModeOk

`func (o *UiConfigResource) GetEnableColorImpairedModeOk() (*bool, bool)`

GetEnableColorImpairedModeOk returns a tuple with the EnableColorImpairedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableColorImpairedMode

`func (o *UiConfigResource) SetEnableColorImpairedMode(v bool)`

SetEnableColorImpairedMode sets EnableColorImpairedMode field to given value.

### HasEnableColorImpairedMode

`func (o *UiConfigResource) HasEnableColorImpairedMode() bool`

HasEnableColorImpairedMode returns a boolean if a field has been set.

### GetTheme

`func (o *UiConfigResource) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *UiConfigResource) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *UiConfigResource) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *UiConfigResource) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *UiConfigResource) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *UiConfigResource) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil
### GetUiLanguage

`func (o *UiConfigResource) GetUiLanguage() int32`

GetUiLanguage returns the UiLanguage field if non-nil, zero value otherwise.

### GetUiLanguageOk

`func (o *UiConfigResource) GetUiLanguageOk() (*int32, bool)`

GetUiLanguageOk returns a tuple with the UiLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiLanguage

`func (o *UiConfigResource) SetUiLanguage(v int32)`

SetUiLanguage sets UiLanguage field to given value.

### HasUiLanguage

`func (o *UiConfigResource) HasUiLanguage() bool`

HasUiLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


