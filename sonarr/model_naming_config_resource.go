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

// NamingConfigResource struct for NamingConfigResource
type NamingConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	RenameEpisodes *bool `json:"renameEpisodes,omitempty"`
	ReplaceIllegalCharacters *bool `json:"replaceIllegalCharacters,omitempty"`
	MultiEpisodeStyle *int32 `json:"multiEpisodeStyle,omitempty"`
	StandardEpisodeFormat NullableString `json:"standardEpisodeFormat,omitempty"`
	DailyEpisodeFormat NullableString `json:"dailyEpisodeFormat,omitempty"`
	AnimeEpisodeFormat NullableString `json:"animeEpisodeFormat,omitempty"`
	SeriesFolderFormat NullableString `json:"seriesFolderFormat,omitempty"`
	SeasonFolderFormat NullableString `json:"seasonFolderFormat,omitempty"`
	SpecialsFolderFormat NullableString `json:"specialsFolderFormat,omitempty"`
	IncludeSeriesTitle *bool `json:"includeSeriesTitle,omitempty"`
	IncludeEpisodeTitle *bool `json:"includeEpisodeTitle,omitempty"`
	IncludeQuality *bool `json:"includeQuality,omitempty"`
	ReplaceSpaces *bool `json:"replaceSpaces,omitempty"`
	Separator NullableString `json:"separator,omitempty"`
	NumberStyle NullableString `json:"numberStyle,omitempty"`
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
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NamingConfigResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
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
	if o == nil || isNil(o.RenameEpisodes) {
		var ret bool
		return ret
	}
	return *o.RenameEpisodes
}

// GetRenameEpisodesOk returns a tuple with the RenameEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetRenameEpisodesOk() (*bool, bool) {
	if o == nil || isNil(o.RenameEpisodes) {
    return nil, false
	}
	return o.RenameEpisodes, true
}

// HasRenameEpisodes returns a boolean if a field has been set.
func (o *NamingConfigResource) HasRenameEpisodes() bool {
	if o != nil && !isNil(o.RenameEpisodes) {
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
	if o == nil || isNil(o.ReplaceIllegalCharacters) {
		var ret bool
		return ret
	}
	return *o.ReplaceIllegalCharacters
}

// GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.ReplaceIllegalCharacters) {
    return nil, false
	}
	return o.ReplaceIllegalCharacters, true
}

// HasReplaceIllegalCharacters returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool {
	if o != nil && !isNil(o.ReplaceIllegalCharacters) {
		return true
	}

	return false
}

// SetReplaceIllegalCharacters gets a reference to the given bool and assigns it to the ReplaceIllegalCharacters field.
func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool) {
	o.ReplaceIllegalCharacters = &v
}

// GetMultiEpisodeStyle returns the MultiEpisodeStyle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetMultiEpisodeStyle() int32 {
	if o == nil || isNil(o.MultiEpisodeStyle) {
		var ret int32
		return ret
	}
	return *o.MultiEpisodeStyle
}

// GetMultiEpisodeStyleOk returns a tuple with the MultiEpisodeStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetMultiEpisodeStyleOk() (*int32, bool) {
	if o == nil || isNil(o.MultiEpisodeStyle) {
    return nil, false
	}
	return o.MultiEpisodeStyle, true
}

// HasMultiEpisodeStyle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasMultiEpisodeStyle() bool {
	if o != nil && !isNil(o.MultiEpisodeStyle) {
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
	if o == nil || isNil(o.StandardEpisodeFormat.Get()) {
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
	if o == nil || isNil(o.DailyEpisodeFormat.Get()) {
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
	if o == nil || isNil(o.AnimeEpisodeFormat.Get()) {
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
	if o == nil || isNil(o.SeriesFolderFormat.Get()) {
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
	if o == nil || isNil(o.SeasonFolderFormat.Get()) {
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
	if o == nil || isNil(o.SpecialsFolderFormat.Get()) {
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

// GetIncludeSeriesTitle returns the IncludeSeriesTitle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeSeriesTitle() bool {
	if o == nil || isNil(o.IncludeSeriesTitle) {
		var ret bool
		return ret
	}
	return *o.IncludeSeriesTitle
}

// GetIncludeSeriesTitleOk returns a tuple with the IncludeSeriesTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeSeriesTitleOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeSeriesTitle) {
    return nil, false
	}
	return o.IncludeSeriesTitle, true
}

// HasIncludeSeriesTitle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeSeriesTitle() bool {
	if o != nil && !isNil(o.IncludeSeriesTitle) {
		return true
	}

	return false
}

// SetIncludeSeriesTitle gets a reference to the given bool and assigns it to the IncludeSeriesTitle field.
func (o *NamingConfigResource) SetIncludeSeriesTitle(v bool) {
	o.IncludeSeriesTitle = &v
}

// GetIncludeEpisodeTitle returns the IncludeEpisodeTitle field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeEpisodeTitle() bool {
	if o == nil || isNil(o.IncludeEpisodeTitle) {
		var ret bool
		return ret
	}
	return *o.IncludeEpisodeTitle
}

// GetIncludeEpisodeTitleOk returns a tuple with the IncludeEpisodeTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeEpisodeTitleOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeEpisodeTitle) {
    return nil, false
	}
	return o.IncludeEpisodeTitle, true
}

// HasIncludeEpisodeTitle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeEpisodeTitle() bool {
	if o != nil && !isNil(o.IncludeEpisodeTitle) {
		return true
	}

	return false
}

// SetIncludeEpisodeTitle gets a reference to the given bool and assigns it to the IncludeEpisodeTitle field.
func (o *NamingConfigResource) SetIncludeEpisodeTitle(v bool) {
	o.IncludeEpisodeTitle = &v
}

// GetIncludeQuality returns the IncludeQuality field value if set, zero value otherwise.
func (o *NamingConfigResource) GetIncludeQuality() bool {
	if o == nil || isNil(o.IncludeQuality) {
		var ret bool
		return ret
	}
	return *o.IncludeQuality
}

// GetIncludeQualityOk returns a tuple with the IncludeQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeQuality) {
    return nil, false
	}
	return o.IncludeQuality, true
}

// HasIncludeQuality returns a boolean if a field has been set.
func (o *NamingConfigResource) HasIncludeQuality() bool {
	if o != nil && !isNil(o.IncludeQuality) {
		return true
	}

	return false
}

// SetIncludeQuality gets a reference to the given bool and assigns it to the IncludeQuality field.
func (o *NamingConfigResource) SetIncludeQuality(v bool) {
	o.IncludeQuality = &v
}

// GetReplaceSpaces returns the ReplaceSpaces field value if set, zero value otherwise.
func (o *NamingConfigResource) GetReplaceSpaces() bool {
	if o == nil || isNil(o.ReplaceSpaces) {
		var ret bool
		return ret
	}
	return *o.ReplaceSpaces
}

// GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool) {
	if o == nil || isNil(o.ReplaceSpaces) {
    return nil, false
	}
	return o.ReplaceSpaces, true
}

// HasReplaceSpaces returns a boolean if a field has been set.
func (o *NamingConfigResource) HasReplaceSpaces() bool {
	if o != nil && !isNil(o.ReplaceSpaces) {
		return true
	}

	return false
}

// SetReplaceSpaces gets a reference to the given bool and assigns it to the ReplaceSpaces field.
func (o *NamingConfigResource) SetReplaceSpaces(v bool) {
	o.ReplaceSpaces = &v
}

// GetSeparator returns the Separator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetSeparator() string {
	if o == nil || isNil(o.Separator.Get()) {
		var ret string
		return ret
	}
	return *o.Separator.Get()
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetSeparatorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Separator.Get(), o.Separator.IsSet()
}

// HasSeparator returns a boolean if a field has been set.
func (o *NamingConfigResource) HasSeparator() bool {
	if o != nil && o.Separator.IsSet() {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given NullableString and assigns it to the Separator field.
func (o *NamingConfigResource) SetSeparator(v string) {
	o.Separator.Set(&v)
}
// SetSeparatorNil sets the value for Separator to be an explicit nil
func (o *NamingConfigResource) SetSeparatorNil() {
	o.Separator.Set(nil)
}

// UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
func (o *NamingConfigResource) UnsetSeparator() {
	o.Separator.Unset()
}

// GetNumberStyle returns the NumberStyle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NamingConfigResource) GetNumberStyle() string {
	if o == nil || isNil(o.NumberStyle.Get()) {
		var ret string
		return ret
	}
	return *o.NumberStyle.Get()
}

// GetNumberStyleOk returns a tuple with the NumberStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NumberStyle.Get(), o.NumberStyle.IsSet()
}

// HasNumberStyle returns a boolean if a field has been set.
func (o *NamingConfigResource) HasNumberStyle() bool {
	if o != nil && o.NumberStyle.IsSet() {
		return true
	}

	return false
}

// SetNumberStyle gets a reference to the given NullableString and assigns it to the NumberStyle field.
func (o *NamingConfigResource) SetNumberStyle(v string) {
	o.NumberStyle.Set(&v)
}
// SetNumberStyleNil sets the value for NumberStyle to be an explicit nil
func (o *NamingConfigResource) SetNumberStyleNil() {
	o.NumberStyle.Set(nil)
}

// UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil
func (o *NamingConfigResource) UnsetNumberStyle() {
	o.NumberStyle.Unset()
}

func (o NamingConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RenameEpisodes) {
		toSerialize["renameEpisodes"] = o.RenameEpisodes
	}
	if !isNil(o.ReplaceIllegalCharacters) {
		toSerialize["replaceIllegalCharacters"] = o.ReplaceIllegalCharacters
	}
	if !isNil(o.MultiEpisodeStyle) {
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
	if !isNil(o.IncludeSeriesTitle) {
		toSerialize["includeSeriesTitle"] = o.IncludeSeriesTitle
	}
	if !isNil(o.IncludeEpisodeTitle) {
		toSerialize["includeEpisodeTitle"] = o.IncludeEpisodeTitle
	}
	if !isNil(o.IncludeQuality) {
		toSerialize["includeQuality"] = o.IncludeQuality
	}
	if !isNil(o.ReplaceSpaces) {
		toSerialize["replaceSpaces"] = o.ReplaceSpaces
	}
	if o.Separator.IsSet() {
		toSerialize["separator"] = o.Separator.Get()
	}
	if o.NumberStyle.IsSet() {
		toSerialize["numberStyle"] = o.NumberStyle.Get()
	}
	return json.Marshal(toSerialize)
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


