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

// checks if the NamingConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamingConfigResource{}

// NamingConfigResource struct for NamingConfigResource
type NamingConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	RenameEpisodes *bool `json:"renameEpisodes,omitempty"`
	ReplaceIllegalCharacters *bool `json:"replaceIllegalCharacters,omitempty"`
	ColonReplacementFormat *int32 `json:"colonReplacementFormat,omitempty"`
	CustomColonReplacementFormat NullableString `json:"customColonReplacementFormat,omitempty"`
	MultiEpisodeStyle *int32 `json:"multiEpisodeStyle,omitempty"`
	StandardEpisodeFormat NullableString `json:"standardEpisodeFormat,omitempty"`
	DailyEpisodeFormat NullableString `json:"dailyEpisodeFormat,omitempty"`
	AnimeEpisodeFormat NullableString `json:"animeEpisodeFormat,omitempty"`
	SeriesFolderFormat NullableString `json:"seriesFolderFormat,omitempty"`
	SeasonFolderFormat NullableString `json:"seasonFolderFormat,omitempty"`
	SpecialsFolderFormat NullableString `json:"specialsFolderFormat,omitempty"`
}

// NewNamingConfigResource instantiates a new NamingConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamingConfigResource() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamingConfigResourceWithDefaults() *NamingConfigResource {
	this := NamingConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NamingConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NamingConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetRenameEpisodes returns the RenameEpisodes field value if set, zero value otherwise.
func (o *NamingConfigResource) GetRenameEpisodes() bool {
	if o == nil || IsNil(o.RenameEpisodes) {
		var ret bool
		return ret
	}
	return *o.RenameEpisodes
}

// GetRenameEpisodesOk returns a tuple with the RenameEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetRenameEpisodesOk() (*bool, bool) {
	if o == nil || IsNil(o.RenameEpisodes) {
		return nil, false
	}
	return o.RenameEpisodes, true
}

// HasRenameEpisodes returns a boolean if a field has been set.
func (o *NamingConfigResource) HasRenameEpisodes() bool {
	if o != nil && !IsNil(o.RenameEpisodes) {
		return true
	}

	return false
}

// SetRenameEpisodes gets a reference to the given bool and assigns it to the RenameEpisodes field.
func (o *NamingConfigResource) SetRenameEpisodes(v bool) {
	o.RenameEpisodes = &v
}

// GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool {
	if o == nil || IsNil(o.ReplaceIllegalCharacters) {
		var ret bool
		return ret
	}
	return *o.ReplaceIllegalCharacters
}

// GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceIllegalCharacters) {
		return nil, false
	}
	return o.ReplaceIllegalCharacters, true
}

// HasReplaceIllegalCharacters returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool {
	if o != nil && !IsNil(o.ReplaceIllegalCharacters) {
		return true
	}

	return false
}

// SetReplaceIllegalCharacters gets a reference to the given bool and assigns it to the ReplaceIllegalCharacters field.
func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool) {
	o.ReplaceIllegalCharacters = &v
}

// GetColonReplacementFormat returns the ColonReplacementFormat field value if set, zero value otherwise.
func (o *NamingConfigResource) GetColonReplacementFormat() int32 {
	if o == nil || IsNil(o.ColonReplacementFormat) {
		var ret int32
		return ret
	}
	return *o.ColonReplacementFormat
}

// GetColonReplacementFormatOk returns a tuple with the ColonReplacementFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetColonReplacementFormatOk() (*int32, bool) {
	if o == nil || IsNil(o.ColonReplacementFormat) {
		return nil, false
	}
	return o.ColonReplacementFormat, true
}

// HasColonReplacementFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasColonReplacementFormat() bool {
	if o != nil && !IsNil(o.ColonReplacementFormat) {
		return true
	}

	return false
}

// SetColonReplacementFormat gets a reference to the given int32 and assigns it to the ColonReplacementFormat field.
func (o *NamingConfigResource) SetColonReplacementFormat(v int32) {
	o.ColonReplacementFormat = &v
}

// GetCustomColonReplacementFormat returns the CustomColonReplacementFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetCustomColonReplacementFormat() string {
	if o == nil || IsNil(o.CustomColonReplacementFormat.Get()) {
		var ret string
		return ret
	}
	return *o.CustomColonReplacementFormat.Get()
}

// GetCustomColonReplacementFormatOk returns a tuple with the CustomColonReplacementFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetCustomColonReplacementFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomColonReplacementFormat.Get(), o.CustomColonReplacementFormat.IsSet()
}

// HasCustomColonReplacementFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasCustomColonReplacementFormat() bool {
	if o != nil && o.CustomColonReplacementFormat.IsSet() {
		return true
	}

	return false
}

// SetCustomColonReplacementFormat gets a reference to the given NullableString and assigns it to the CustomColonReplacementFormat field.
func (o *NamingConfigResource) SetCustomColonReplacementFormat(v string) {
	o.CustomColonReplacementFormat.Set(&v)
}
// SetCustomColonReplacementFormatNil sets the value for CustomColonReplacementFormat to be an explicit nil
func (o *NamingConfigResource) SetCustomColonReplacementFormatNil() {
	o.CustomColonReplacementFormat.Set(nil)
}

// UnsetCustomColonReplacementFormat ensures that no value is present for CustomColonReplacementFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetCustomColonReplacementFormat() {
	o.CustomColonReplacementFormat.Unset()
}

// GetMultiEpisodeStyle returns the MultiEpisodeStyle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetMultiEpisodeStyle() int32 {
	if o == nil || IsNil(o.MultiEpisodeStyle) {
		var ret int32
		return ret
	}
	return *o.MultiEpisodeStyle
}

// GetMultiEpisodeStyleOk returns a tuple with the MultiEpisodeStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetMultiEpisodeStyleOk() (*int32, bool) {
	if o == nil || IsNil(o.MultiEpisodeStyle) {
		return nil, false
	}
	return o.MultiEpisodeStyle, true
}

// HasMultiEpisodeStyle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasMultiEpisodeStyle() bool {
	if o != nil && !IsNil(o.MultiEpisodeStyle) {
		return true
	}

	return false
}

// SetMultiEpisodeStyle gets a reference to the given int32 and assigns it to the MultiEpisodeStyle field.
func (o *NamingConfigResource) SetMultiEpisodeStyle(v int32) {
	o.MultiEpisodeStyle = &v
}

// GetStandardEpisodeFormat returns the StandardEpisodeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetStandardEpisodeFormat() string {
	if o == nil || IsNil(o.StandardEpisodeFormat.Get()) {
		var ret string
		return ret
	}
	return *o.StandardEpisodeFormat.Get()
}

// GetStandardEpisodeFormatOk returns a tuple with the StandardEpisodeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetStandardEpisodeFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StandardEpisodeFormat.Get(), o.StandardEpisodeFormat.IsSet()
}

// HasStandardEpisodeFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasStandardEpisodeFormat() bool {
	if o != nil && o.StandardEpisodeFormat.IsSet() {
		return true
	}

	return false
}

// SetStandardEpisodeFormat gets a reference to the given NullableString and assigns it to the StandardEpisodeFormat field.
func (o *NamingConfigResource) SetStandardEpisodeFormat(v string) {
	o.StandardEpisodeFormat.Set(&v)
}
// SetStandardEpisodeFormatNil sets the value for StandardEpisodeFormat to be an explicit nil
func (o *NamingConfigResource) SetStandardEpisodeFormatNil() {
	o.StandardEpisodeFormat.Set(nil)
}

// UnsetStandardEpisodeFormat ensures that no value is present for StandardEpisodeFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetStandardEpisodeFormat() {
	o.StandardEpisodeFormat.Unset()
}

// GetDailyEpisodeFormat returns the DailyEpisodeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetDailyEpisodeFormat() string {
	if o == nil || IsNil(o.DailyEpisodeFormat.Get()) {
		var ret string
		return ret
	}
	return *o.DailyEpisodeFormat.Get()
}

// GetDailyEpisodeFormatOk returns a tuple with the DailyEpisodeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetDailyEpisodeFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DailyEpisodeFormat.Get(), o.DailyEpisodeFormat.IsSet()
}

// HasDailyEpisodeFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasDailyEpisodeFormat() bool {
	if o != nil && o.DailyEpisodeFormat.IsSet() {
		return true
	}

	return false
}

// SetDailyEpisodeFormat gets a reference to the given NullableString and assigns it to the DailyEpisodeFormat field.
func (o *NamingConfigResource) SetDailyEpisodeFormat(v string) {
	o.DailyEpisodeFormat.Set(&v)
}
// SetDailyEpisodeFormatNil sets the value for DailyEpisodeFormat to be an explicit nil
func (o *NamingConfigResource) SetDailyEpisodeFormatNil() {
	o.DailyEpisodeFormat.Set(nil)
}

// UnsetDailyEpisodeFormat ensures that no value is present for DailyEpisodeFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetDailyEpisodeFormat() {
	o.DailyEpisodeFormat.Unset()
}

// GetAnimeEpisodeFormat returns the AnimeEpisodeFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetAnimeEpisodeFormat() string {
	if o == nil || IsNil(o.AnimeEpisodeFormat.Get()) {
		var ret string
		return ret
	}
	return *o.AnimeEpisodeFormat.Get()
}

// GetAnimeEpisodeFormatOk returns a tuple with the AnimeEpisodeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetAnimeEpisodeFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnimeEpisodeFormat.Get(), o.AnimeEpisodeFormat.IsSet()
}

// HasAnimeEpisodeFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasAnimeEpisodeFormat() bool {
	if o != nil && o.AnimeEpisodeFormat.IsSet() {
		return true
	}

	return false
}

// SetAnimeEpisodeFormat gets a reference to the given NullableString and assigns it to the AnimeEpisodeFormat field.
func (o *NamingConfigResource) SetAnimeEpisodeFormat(v string) {
	o.AnimeEpisodeFormat.Set(&v)
}
// SetAnimeEpisodeFormatNil sets the value for AnimeEpisodeFormat to be an explicit nil
func (o *NamingConfigResource) SetAnimeEpisodeFormatNil() {
	o.AnimeEpisodeFormat.Set(nil)
}

// UnsetAnimeEpisodeFormat ensures that no value is present for AnimeEpisodeFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetAnimeEpisodeFormat() {
	o.AnimeEpisodeFormat.Unset()
}

// GetSeriesFolderFormat returns the SeriesFolderFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSeriesFolderFormat() string {
	if o == nil || IsNil(o.SeriesFolderFormat.Get()) {
		var ret string
		return ret
	}
	return *o.SeriesFolderFormat.Get()
}

// GetSeriesFolderFormatOk returns a tuple with the SeriesFolderFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSeriesFolderFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeriesFolderFormat.Get(), o.SeriesFolderFormat.IsSet()
}

// HasSeriesFolderFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSeriesFolderFormat() bool {
	if o != nil && o.SeriesFolderFormat.IsSet() {
		return true
	}

	return false
}

// SetSeriesFolderFormat gets a reference to the given NullableString and assigns it to the SeriesFolderFormat field.
func (o *NamingConfigResource) SetSeriesFolderFormat(v string) {
	o.SeriesFolderFormat.Set(&v)
}
// SetSeriesFolderFormatNil sets the value for SeriesFolderFormat to be an explicit nil
func (o *NamingConfigResource) SetSeriesFolderFormatNil() {
	o.SeriesFolderFormat.Set(nil)
}

// UnsetSeriesFolderFormat ensures that no value is present for SeriesFolderFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetSeriesFolderFormat() {
	o.SeriesFolderFormat.Unset()
}

// GetSeasonFolderFormat returns the SeasonFolderFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSeasonFolderFormat() string {
	if o == nil || IsNil(o.SeasonFolderFormat.Get()) {
		var ret string
		return ret
	}
	return *o.SeasonFolderFormat.Get()
}

// GetSeasonFolderFormatOk returns a tuple with the SeasonFolderFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSeasonFolderFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeasonFolderFormat.Get(), o.SeasonFolderFormat.IsSet()
}

// HasSeasonFolderFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSeasonFolderFormat() bool {
	if o != nil && o.SeasonFolderFormat.IsSet() {
		return true
	}

	return false
}

// SetSeasonFolderFormat gets a reference to the given NullableString and assigns it to the SeasonFolderFormat field.
func (o *NamingConfigResource) SetSeasonFolderFormat(v string) {
	o.SeasonFolderFormat.Set(&v)
}
// SetSeasonFolderFormatNil sets the value for SeasonFolderFormat to be an explicit nil
func (o *NamingConfigResource) SetSeasonFolderFormatNil() {
	o.SeasonFolderFormat.Set(nil)
}

// UnsetSeasonFolderFormat ensures that no value is present for SeasonFolderFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetSeasonFolderFormat() {
	o.SeasonFolderFormat.Unset()
}

// GetSpecialsFolderFormat returns the SpecialsFolderFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSpecialsFolderFormat() string {
	if o == nil || IsNil(o.SpecialsFolderFormat.Get()) {
		var ret string
		return ret
	}
	return *o.SpecialsFolderFormat.Get()
}

// GetSpecialsFolderFormatOk returns a tuple with the SpecialsFolderFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSpecialsFolderFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecialsFolderFormat.Get(), o.SpecialsFolderFormat.IsSet()
}

// HasSpecialsFolderFormat returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSpecialsFolderFormat() bool {
	if o != nil && o.SpecialsFolderFormat.IsSet() {
		return true
	}

	return false
}

// SetSpecialsFolderFormat gets a reference to the given NullableString and assigns it to the SpecialsFolderFormat field.
func (o *NamingConfigResource) SetSpecialsFolderFormat(v string) {
	o.SpecialsFolderFormat.Set(&v)
}
// SetSpecialsFolderFormatNil sets the value for SpecialsFolderFormat to be an explicit nil
func (o *NamingConfigResource) SetSpecialsFolderFormatNil() {
	o.SpecialsFolderFormat.Set(nil)
}

// UnsetSpecialsFolderFormat ensures that no value is present for SpecialsFolderFormat, not even an explicit nil
func (o *NamingConfigResource) UnsetSpecialsFolderFormat() {
	o.SpecialsFolderFormat.Unset()
}

func (o NamingConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamingConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RenameEpisodes) {
		toSerialize["renameEpisodes"] = o.RenameEpisodes
	}
	if !IsNil(o.ReplaceIllegalCharacters) {
		toSerialize["replaceIllegalCharacters"] = o.ReplaceIllegalCharacters
	}
	if !IsNil(o.ColonReplacementFormat) {
		toSerialize["colonReplacementFormat"] = o.ColonReplacementFormat
	}
	if o.CustomColonReplacementFormat.IsSet() {
		toSerialize["customColonReplacementFormat"] = o.CustomColonReplacementFormat.Get()
	}
	if !IsNil(o.MultiEpisodeStyle) {
		toSerialize["multiEpisodeStyle"] = o.MultiEpisodeStyle
	}
	if o.StandardEpisodeFormat.IsSet() {
		toSerialize["standardEpisodeFormat"] = o.StandardEpisodeFormat.Get()
	}
	if o.DailyEpisodeFormat.IsSet() {
		toSerialize["dailyEpisodeFormat"] = o.DailyEpisodeFormat.Get()
	}
	if o.AnimeEpisodeFormat.IsSet() {
		toSerialize["animeEpisodeFormat"] = o.AnimeEpisodeFormat.Get()
	}
	if o.SeriesFolderFormat.IsSet() {
		toSerialize["seriesFolderFormat"] = o.SeriesFolderFormat.Get()
	}
	if o.SeasonFolderFormat.IsSet() {
		toSerialize["seasonFolderFormat"] = o.SeasonFolderFormat.Get()
	}
	if o.SpecialsFolderFormat.IsSet() {
		toSerialize["specialsFolderFormat"] = o.SpecialsFolderFormat.Get()
	}
	return toSerialize, nil
}

type NullableNamingConfigResource struct {
	value *NamingConfigResource
	isSet bool
}

func (v NullableNamingConfigResource) Get() *NamingConfigResource {
	return v.value
}

func (v *NullableNamingConfigResource) Set(val *NamingConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableNamingConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableNamingConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamingConfigResource(val *NamingConfigResource) *NullableNamingConfigResource {
	return &NullableNamingConfigResource{value: val, isSet: true}
}

func (v NullableNamingConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamingConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


