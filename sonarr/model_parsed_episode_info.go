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

// ParsedEpisodeInfo struct for ParsedEpisodeInfo
type ParsedEpisodeInfo struct {
	ReleaseTitle NullableString `json:"releaseTitle,omitempty"`
	SeriesTitle NullableString `json:"seriesTitle,omitempty"`
	SeriesTitleInfo *SeriesTitleInfo `json:"seriesTitleInfo,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	SeasonNumber *int32 `json:"seasonNumber,omitempty"`
	EpisodeNumbers []*int32 `json:"episodeNumbers,omitempty"`
	AbsoluteEpisodeNumbers []*int32 `json:"absoluteEpisodeNumbers,omitempty"`
	SpecialAbsoluteEpisodeNumbers []*float64 `json:"specialAbsoluteEpisodeNumbers,omitempty"`
	AirDate NullableString `json:"airDate,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	FullSeason *bool `json:"fullSeason,omitempty"`
	IsPartialSeason *bool `json:"isPartialSeason,omitempty"`
	IsMultiSeason *bool `json:"isMultiSeason,omitempty"`
	IsSeasonExtra *bool `json:"isSeasonExtra,omitempty"`
	Special *bool `json:"special,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	ReleaseHash NullableString `json:"releaseHash,omitempty"`
	SeasonPart *int32 `json:"seasonPart,omitempty"`
	ReleaseTokens NullableString `json:"releaseTokens,omitempty"`
	DailyPart NullableInt32 `json:"dailyPart,omitempty"`
	IsDaily *bool `json:"isDaily,omitempty"`
	IsAbsoluteNumbering *bool `json:"isAbsoluteNumbering,omitempty"`
	IsPossibleSpecialEpisode *bool `json:"isPossibleSpecialEpisode,omitempty"`
	IsPossibleSceneSeasonSpecial *bool `json:"isPossibleSceneSeasonSpecial,omitempty"`
}

// NewParsedEpisodeInfo instantiates a new ParsedEpisodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsedEpisodeInfo() *ParsedEpisodeInfo {
	this := ParsedEpisodeInfo{}
	return &this
}

// NewParsedEpisodeInfoWithDefaults instantiates a new ParsedEpisodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsedEpisodeInfoWithDefaults() *ParsedEpisodeInfo {
	this := ParsedEpisodeInfo{}
	return &this
}

// GetReleaseTitle returns the ReleaseTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetReleaseTitle() string {
	if o == nil || isNil(o.ReleaseTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseTitle.Get()
}

// GetReleaseTitleOk returns a tuple with the ReleaseTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetReleaseTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseTitle.Get(), o.ReleaseTitle.IsSet()
}

// HasReleaseTitle returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasReleaseTitle() bool {
	if o != nil && o.ReleaseTitle.IsSet() {
		return true
	}

	return false
}

// SetReleaseTitle gets a reference to the given NullableString and assigns it to the ReleaseTitle field.
func (o *ParsedEpisodeInfo) SetReleaseTitle(v string) {
	o.ReleaseTitle.Set(&v)
}
// SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil
func (o *ParsedEpisodeInfo) SetReleaseTitleNil() {
	o.ReleaseTitle.Set(nil)
}

// UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetReleaseTitle() {
	o.ReleaseTitle.Unset()
}

// GetSeriesTitle returns the SeriesTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetSeriesTitle() string {
	if o == nil || isNil(o.SeriesTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SeriesTitle.Get()
}

// GetSeriesTitleOk returns a tuple with the SeriesTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetSeriesTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SeriesTitle.Get(), o.SeriesTitle.IsSet()
}

// HasSeriesTitle returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSeriesTitle() bool {
	if o != nil && o.SeriesTitle.IsSet() {
		return true
	}

	return false
}

// SetSeriesTitle gets a reference to the given NullableString and assigns it to the SeriesTitle field.
func (o *ParsedEpisodeInfo) SetSeriesTitle(v string) {
	o.SeriesTitle.Set(&v)
}
// SetSeriesTitleNil sets the value for SeriesTitle to be an explicit nil
func (o *ParsedEpisodeInfo) SetSeriesTitleNil() {
	o.SeriesTitle.Set(nil)
}

// UnsetSeriesTitle ensures that no value is present for SeriesTitle, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetSeriesTitle() {
	o.SeriesTitle.Unset()
}

// GetSeriesTitleInfo returns the SeriesTitleInfo field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetSeriesTitleInfo() SeriesTitleInfo {
	if o == nil || isNil(o.SeriesTitleInfo) {
		var ret SeriesTitleInfo
		return ret
	}
	return *o.SeriesTitleInfo
}

// GetSeriesTitleInfoOk returns a tuple with the SeriesTitleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetSeriesTitleInfoOk() (*SeriesTitleInfo, bool) {
	if o == nil || isNil(o.SeriesTitleInfo) {
    return nil, false
	}
	return o.SeriesTitleInfo, true
}

// HasSeriesTitleInfo returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSeriesTitleInfo() bool {
	if o != nil && !isNil(o.SeriesTitleInfo) {
		return true
	}

	return false
}

// SetSeriesTitleInfo gets a reference to the given SeriesTitleInfo and assigns it to the SeriesTitleInfo field.
func (o *ParsedEpisodeInfo) SetSeriesTitleInfo(v SeriesTitleInfo) {
	o.SeriesTitleInfo = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ParsedEpisodeInfo) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetSeasonNumber() int32 {
	if o == nil || isNil(o.SeasonNumber) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetSeasonNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SeasonNumber) {
    return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSeasonNumber() bool {
	if o != nil && !isNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given int32 and assigns it to the SeasonNumber field.
func (o *ParsedEpisodeInfo) SetSeasonNumber(v int32) {
	o.SeasonNumber = &v
}

// GetEpisodeNumbers returns the EpisodeNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetEpisodeNumbers() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.EpisodeNumbers
}

// GetEpisodeNumbersOk returns a tuple with the EpisodeNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetEpisodeNumbersOk() ([]*int32, bool) {
	if o == nil || isNil(o.EpisodeNumbers) {
    return nil, false
	}
	return o.EpisodeNumbers, true
}

// HasEpisodeNumbers returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasEpisodeNumbers() bool {
	if o != nil && isNil(o.EpisodeNumbers) {
		return true
	}

	return false
}

// SetEpisodeNumbers gets a reference to the given []int32 and assigns it to the EpisodeNumbers field.
func (o *ParsedEpisodeInfo) SetEpisodeNumbers(v []*int32) {
	o.EpisodeNumbers = v
}

// GetAbsoluteEpisodeNumbers returns the AbsoluteEpisodeNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetAbsoluteEpisodeNumbers() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.AbsoluteEpisodeNumbers
}

// GetAbsoluteEpisodeNumbersOk returns a tuple with the AbsoluteEpisodeNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetAbsoluteEpisodeNumbersOk() ([]*int32, bool) {
	if o == nil || isNil(o.AbsoluteEpisodeNumbers) {
    return nil, false
	}
	return o.AbsoluteEpisodeNumbers, true
}

// HasAbsoluteEpisodeNumbers returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasAbsoluteEpisodeNumbers() bool {
	if o != nil && isNil(o.AbsoluteEpisodeNumbers) {
		return true
	}

	return false
}

// SetAbsoluteEpisodeNumbers gets a reference to the given []int32 and assigns it to the AbsoluteEpisodeNumbers field.
func (o *ParsedEpisodeInfo) SetAbsoluteEpisodeNumbers(v []*int32) {
	o.AbsoluteEpisodeNumbers = v
}

// GetSpecialAbsoluteEpisodeNumbers returns the SpecialAbsoluteEpisodeNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetSpecialAbsoluteEpisodeNumbers() []*float64 {
	if o == nil {
		var ret []*float64
		return ret
	}
	return o.SpecialAbsoluteEpisodeNumbers
}

// GetSpecialAbsoluteEpisodeNumbersOk returns a tuple with the SpecialAbsoluteEpisodeNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetSpecialAbsoluteEpisodeNumbersOk() ([]*float64, bool) {
	if o == nil || isNil(o.SpecialAbsoluteEpisodeNumbers) {
    return nil, false
	}
	return o.SpecialAbsoluteEpisodeNumbers, true
}

// HasSpecialAbsoluteEpisodeNumbers returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSpecialAbsoluteEpisodeNumbers() bool {
	if o != nil && isNil(o.SpecialAbsoluteEpisodeNumbers) {
		return true
	}

	return false
}

// SetSpecialAbsoluteEpisodeNumbers gets a reference to the given []float64 and assigns it to the SpecialAbsoluteEpisodeNumbers field.
func (o *ParsedEpisodeInfo) SetSpecialAbsoluteEpisodeNumbers(v []*float64) {
	o.SpecialAbsoluteEpisodeNumbers = v
}

// GetAirDate returns the AirDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetAirDate() string {
	if o == nil || isNil(o.AirDate.Get()) {
		var ret string
		return ret
	}
	return *o.AirDate.Get()
}

// GetAirDateOk returns a tuple with the AirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetAirDateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AirDate.Get(), o.AirDate.IsSet()
}

// HasAirDate returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasAirDate() bool {
	if o != nil && o.AirDate.IsSet() {
		return true
	}

	return false
}

// SetAirDate gets a reference to the given NullableString and assigns it to the AirDate field.
func (o *ParsedEpisodeInfo) SetAirDate(v string) {
	o.AirDate.Set(&v)
}
// SetAirDateNil sets the value for AirDate to be an explicit nil
func (o *ParsedEpisodeInfo) SetAirDateNil() {
	o.AirDate.Set(nil)
}

// UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetAirDate() {
	o.AirDate.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ParsedEpisodeInfo) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetFullSeason returns the FullSeason field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetFullSeason() bool {
	if o == nil || isNil(o.FullSeason) {
		var ret bool
		return ret
	}
	return *o.FullSeason
}

// GetFullSeasonOk returns a tuple with the FullSeason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetFullSeasonOk() (*bool, bool) {
	if o == nil || isNil(o.FullSeason) {
    return nil, false
	}
	return o.FullSeason, true
}

// HasFullSeason returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasFullSeason() bool {
	if o != nil && !isNil(o.FullSeason) {
		return true
	}

	return false
}

// SetFullSeason gets a reference to the given bool and assigns it to the FullSeason field.
func (o *ParsedEpisodeInfo) SetFullSeason(v bool) {
	o.FullSeason = &v
}

// GetIsPartialSeason returns the IsPartialSeason field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsPartialSeason() bool {
	if o == nil || isNil(o.IsPartialSeason) {
		var ret bool
		return ret
	}
	return *o.IsPartialSeason
}

// GetIsPartialSeasonOk returns a tuple with the IsPartialSeason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsPartialSeasonOk() (*bool, bool) {
	if o == nil || isNil(o.IsPartialSeason) {
    return nil, false
	}
	return o.IsPartialSeason, true
}

// HasIsPartialSeason returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsPartialSeason() bool {
	if o != nil && !isNil(o.IsPartialSeason) {
		return true
	}

	return false
}

// SetIsPartialSeason gets a reference to the given bool and assigns it to the IsPartialSeason field.
func (o *ParsedEpisodeInfo) SetIsPartialSeason(v bool) {
	o.IsPartialSeason = &v
}

// GetIsMultiSeason returns the IsMultiSeason field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsMultiSeason() bool {
	if o == nil || isNil(o.IsMultiSeason) {
		var ret bool
		return ret
	}
	return *o.IsMultiSeason
}

// GetIsMultiSeasonOk returns a tuple with the IsMultiSeason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsMultiSeasonOk() (*bool, bool) {
	if o == nil || isNil(o.IsMultiSeason) {
    return nil, false
	}
	return o.IsMultiSeason, true
}

// HasIsMultiSeason returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsMultiSeason() bool {
	if o != nil && !isNil(o.IsMultiSeason) {
		return true
	}

	return false
}

// SetIsMultiSeason gets a reference to the given bool and assigns it to the IsMultiSeason field.
func (o *ParsedEpisodeInfo) SetIsMultiSeason(v bool) {
	o.IsMultiSeason = &v
}

// GetIsSeasonExtra returns the IsSeasonExtra field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsSeasonExtra() bool {
	if o == nil || isNil(o.IsSeasonExtra) {
		var ret bool
		return ret
	}
	return *o.IsSeasonExtra
}

// GetIsSeasonExtraOk returns a tuple with the IsSeasonExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsSeasonExtraOk() (*bool, bool) {
	if o == nil || isNil(o.IsSeasonExtra) {
    return nil, false
	}
	return o.IsSeasonExtra, true
}

// HasIsSeasonExtra returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsSeasonExtra() bool {
	if o != nil && !isNil(o.IsSeasonExtra) {
		return true
	}

	return false
}

// SetIsSeasonExtra gets a reference to the given bool and assigns it to the IsSeasonExtra field.
func (o *ParsedEpisodeInfo) SetIsSeasonExtra(v bool) {
	o.IsSeasonExtra = &v
}

// GetSpecial returns the Special field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetSpecial() bool {
	if o == nil || isNil(o.Special) {
		var ret bool
		return ret
	}
	return *o.Special
}

// GetSpecialOk returns a tuple with the Special field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetSpecialOk() (*bool, bool) {
	if o == nil || isNil(o.Special) {
    return nil, false
	}
	return o.Special, true
}

// HasSpecial returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSpecial() bool {
	if o != nil && !isNil(o.Special) {
		return true
	}

	return false
}

// SetSpecial gets a reference to the given bool and assigns it to the Special field.
func (o *ParsedEpisodeInfo) SetSpecial(v bool) {
	o.Special = &v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ParsedEpisodeInfo) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ParsedEpisodeInfo) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetReleaseHash returns the ReleaseHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetReleaseHash() string {
	if o == nil || isNil(o.ReleaseHash.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseHash.Get()
}

// GetReleaseHashOk returns a tuple with the ReleaseHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetReleaseHashOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseHash.Get(), o.ReleaseHash.IsSet()
}

// HasReleaseHash returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasReleaseHash() bool {
	if o != nil && o.ReleaseHash.IsSet() {
		return true
	}

	return false
}

// SetReleaseHash gets a reference to the given NullableString and assigns it to the ReleaseHash field.
func (o *ParsedEpisodeInfo) SetReleaseHash(v string) {
	o.ReleaseHash.Set(&v)
}
// SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil
func (o *ParsedEpisodeInfo) SetReleaseHashNil() {
	o.ReleaseHash.Set(nil)
}

// UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetReleaseHash() {
	o.ReleaseHash.Unset()
}

// GetSeasonPart returns the SeasonPart field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetSeasonPart() int32 {
	if o == nil || isNil(o.SeasonPart) {
		var ret int32
		return ret
	}
	return *o.SeasonPart
}

// GetSeasonPartOk returns a tuple with the SeasonPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetSeasonPartOk() (*int32, bool) {
	if o == nil || isNil(o.SeasonPart) {
    return nil, false
	}
	return o.SeasonPart, true
}

// HasSeasonPart returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasSeasonPart() bool {
	if o != nil && !isNil(o.SeasonPart) {
		return true
	}

	return false
}

// SetSeasonPart gets a reference to the given int32 and assigns it to the SeasonPart field.
func (o *ParsedEpisodeInfo) SetSeasonPart(v int32) {
	o.SeasonPart = &v
}

// GetReleaseTokens returns the ReleaseTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetReleaseTokens() string {
	if o == nil || isNil(o.ReleaseTokens.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseTokens.Get()
}

// GetReleaseTokensOk returns a tuple with the ReleaseTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetReleaseTokensOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseTokens.Get(), o.ReleaseTokens.IsSet()
}

// HasReleaseTokens returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasReleaseTokens() bool {
	if o != nil && o.ReleaseTokens.IsSet() {
		return true
	}

	return false
}

// SetReleaseTokens gets a reference to the given NullableString and assigns it to the ReleaseTokens field.
func (o *ParsedEpisodeInfo) SetReleaseTokens(v string) {
	o.ReleaseTokens.Set(&v)
}
// SetReleaseTokensNil sets the value for ReleaseTokens to be an explicit nil
func (o *ParsedEpisodeInfo) SetReleaseTokensNil() {
	o.ReleaseTokens.Set(nil)
}

// UnsetReleaseTokens ensures that no value is present for ReleaseTokens, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetReleaseTokens() {
	o.ReleaseTokens.Unset()
}

// GetDailyPart returns the DailyPart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedEpisodeInfo) GetDailyPart() int32 {
	if o == nil || isNil(o.DailyPart.Get()) {
		var ret int32
		return ret
	}
	return *o.DailyPart.Get()
}

// GetDailyPartOk returns a tuple with the DailyPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedEpisodeInfo) GetDailyPartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.DailyPart.Get(), o.DailyPart.IsSet()
}

// HasDailyPart returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasDailyPart() bool {
	if o != nil && o.DailyPart.IsSet() {
		return true
	}

	return false
}

// SetDailyPart gets a reference to the given NullableInt32 and assigns it to the DailyPart field.
func (o *ParsedEpisodeInfo) SetDailyPart(v int32) {
	o.DailyPart.Set(&v)
}
// SetDailyPartNil sets the value for DailyPart to be an explicit nil
func (o *ParsedEpisodeInfo) SetDailyPartNil() {
	o.DailyPart.Set(nil)
}

// UnsetDailyPart ensures that no value is present for DailyPart, not even an explicit nil
func (o *ParsedEpisodeInfo) UnsetDailyPart() {
	o.DailyPart.Unset()
}

// GetIsDaily returns the IsDaily field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsDaily() bool {
	if o == nil || isNil(o.IsDaily) {
		var ret bool
		return ret
	}
	return *o.IsDaily
}

// GetIsDailyOk returns a tuple with the IsDaily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsDailyOk() (*bool, bool) {
	if o == nil || isNil(o.IsDaily) {
    return nil, false
	}
	return o.IsDaily, true
}

// HasIsDaily returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsDaily() bool {
	if o != nil && !isNil(o.IsDaily) {
		return true
	}

	return false
}

// SetIsDaily gets a reference to the given bool and assigns it to the IsDaily field.
func (o *ParsedEpisodeInfo) SetIsDaily(v bool) {
	o.IsDaily = &v
}

// GetIsAbsoluteNumbering returns the IsAbsoluteNumbering field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsAbsoluteNumbering() bool {
	if o == nil || isNil(o.IsAbsoluteNumbering) {
		var ret bool
		return ret
	}
	return *o.IsAbsoluteNumbering
}

// GetIsAbsoluteNumberingOk returns a tuple with the IsAbsoluteNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsAbsoluteNumberingOk() (*bool, bool) {
	if o == nil || isNil(o.IsAbsoluteNumbering) {
    return nil, false
	}
	return o.IsAbsoluteNumbering, true
}

// HasIsAbsoluteNumbering returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsAbsoluteNumbering() bool {
	if o != nil && !isNil(o.IsAbsoluteNumbering) {
		return true
	}

	return false
}

// SetIsAbsoluteNumbering gets a reference to the given bool and assigns it to the IsAbsoluteNumbering field.
func (o *ParsedEpisodeInfo) SetIsAbsoluteNumbering(v bool) {
	o.IsAbsoluteNumbering = &v
}

// GetIsPossibleSpecialEpisode returns the IsPossibleSpecialEpisode field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsPossibleSpecialEpisode() bool {
	if o == nil || isNil(o.IsPossibleSpecialEpisode) {
		var ret bool
		return ret
	}
	return *o.IsPossibleSpecialEpisode
}

// GetIsPossibleSpecialEpisodeOk returns a tuple with the IsPossibleSpecialEpisode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsPossibleSpecialEpisodeOk() (*bool, bool) {
	if o == nil || isNil(o.IsPossibleSpecialEpisode) {
    return nil, false
	}
	return o.IsPossibleSpecialEpisode, true
}

// HasIsPossibleSpecialEpisode returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsPossibleSpecialEpisode() bool {
	if o != nil && !isNil(o.IsPossibleSpecialEpisode) {
		return true
	}

	return false
}

// SetIsPossibleSpecialEpisode gets a reference to the given bool and assigns it to the IsPossibleSpecialEpisode field.
func (o *ParsedEpisodeInfo) SetIsPossibleSpecialEpisode(v bool) {
	o.IsPossibleSpecialEpisode = &v
}

// GetIsPossibleSceneSeasonSpecial returns the IsPossibleSceneSeasonSpecial field value if set, zero value otherwise.
func (o *ParsedEpisodeInfo) GetIsPossibleSceneSeasonSpecial() bool {
	if o == nil || isNil(o.IsPossibleSceneSeasonSpecial) {
		var ret bool
		return ret
	}
	return *o.IsPossibleSceneSeasonSpecial
}

// GetIsPossibleSceneSeasonSpecialOk returns a tuple with the IsPossibleSceneSeasonSpecial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedEpisodeInfo) GetIsPossibleSceneSeasonSpecialOk() (*bool, bool) {
	if o == nil || isNil(o.IsPossibleSceneSeasonSpecial) {
    return nil, false
	}
	return o.IsPossibleSceneSeasonSpecial, true
}

// HasIsPossibleSceneSeasonSpecial returns a boolean if a field has been set.
func (o *ParsedEpisodeInfo) HasIsPossibleSceneSeasonSpecial() bool {
	if o != nil && !isNil(o.IsPossibleSceneSeasonSpecial) {
		return true
	}

	return false
}

// SetIsPossibleSceneSeasonSpecial gets a reference to the given bool and assigns it to the IsPossibleSceneSeasonSpecial field.
func (o *ParsedEpisodeInfo) SetIsPossibleSceneSeasonSpecial(v bool) {
	o.IsPossibleSceneSeasonSpecial = &v
}

func (o ParsedEpisodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReleaseTitle.IsSet() {
		toSerialize["releaseTitle"] = o.ReleaseTitle.Get()
	}
	if o.SeriesTitle.IsSet() {
		toSerialize["seriesTitle"] = o.SeriesTitle.Get()
	}
	if !isNil(o.SeriesTitleInfo) {
		toSerialize["seriesTitleInfo"] = o.SeriesTitleInfo
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if o.EpisodeNumbers != nil {
		toSerialize["episodeNumbers"] = o.EpisodeNumbers
	}
	if o.AbsoluteEpisodeNumbers != nil {
		toSerialize["absoluteEpisodeNumbers"] = o.AbsoluteEpisodeNumbers
	}
	if o.SpecialAbsoluteEpisodeNumbers != nil {
		toSerialize["specialAbsoluteEpisodeNumbers"] = o.SpecialAbsoluteEpisodeNumbers
	}
	if o.AirDate.IsSet() {
		toSerialize["airDate"] = o.AirDate.Get()
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !isNil(o.FullSeason) {
		toSerialize["fullSeason"] = o.FullSeason
	}
	if !isNil(o.IsPartialSeason) {
		toSerialize["isPartialSeason"] = o.IsPartialSeason
	}
	if !isNil(o.IsMultiSeason) {
		toSerialize["isMultiSeason"] = o.IsMultiSeason
	}
	if !isNil(o.IsSeasonExtra) {
		toSerialize["isSeasonExtra"] = o.IsSeasonExtra
	}
	if !isNil(o.Special) {
		toSerialize["special"] = o.Special
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.ReleaseHash.IsSet() {
		toSerialize["releaseHash"] = o.ReleaseHash.Get()
	}
	if !isNil(o.SeasonPart) {
		toSerialize["seasonPart"] = o.SeasonPart
	}
	if o.ReleaseTokens.IsSet() {
		toSerialize["releaseTokens"] = o.ReleaseTokens.Get()
	}
	if o.DailyPart.IsSet() {
		toSerialize["dailyPart"] = o.DailyPart.Get()
	}
	if !isNil(o.IsDaily) {
		toSerialize["isDaily"] = o.IsDaily
	}
	if !isNil(o.IsAbsoluteNumbering) {
		toSerialize["isAbsoluteNumbering"] = o.IsAbsoluteNumbering
	}
	if !isNil(o.IsPossibleSpecialEpisode) {
		toSerialize["isPossibleSpecialEpisode"] = o.IsPossibleSpecialEpisode
	}
	if !isNil(o.IsPossibleSceneSeasonSpecial) {
		toSerialize["isPossibleSceneSeasonSpecial"] = o.IsPossibleSceneSeasonSpecial
	}
	return json.Marshal(toSerialize)
}

type NullableParsedEpisodeInfo struct {
	value *ParsedEpisodeInfo
	isSet bool
}

func (v NullableParsedEpisodeInfo) Get() *ParsedEpisodeInfo {
	return v.value
}

func (v *NullableParsedEpisodeInfo) Set(val *ParsedEpisodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableParsedEpisodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableParsedEpisodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsedEpisodeInfo(val *ParsedEpisodeInfo) *NullableParsedEpisodeInfo {
	return &NullableParsedEpisodeInfo{value: val, isSet: true}
}

func (v NullableParsedEpisodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsedEpisodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


