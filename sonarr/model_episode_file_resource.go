/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.12.2823
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"time"
)

// checks if the EpisodeFileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpisodeFileResource{}

// EpisodeFileResource struct for EpisodeFileResource
type EpisodeFileResource struct {
	Id *int32 `json:"id,omitempty"`
	SeriesId *int32 `json:"seriesId,omitempty"`
	SeasonNumber *int32 `json:"seasonNumber,omitempty"`
	RelativePath NullableString `json:"relativePath,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Size *int64 `json:"size,omitempty"`
	DateAdded *time.Time `json:"dateAdded,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	IndexerFlags NullableInt32 `json:"indexerFlags,omitempty"`
	ReleaseType *ReleaseType `json:"releaseType,omitempty"`
	MediaInfo *MediaInfoResource `json:"mediaInfo,omitempty"`
	QualityCutoffNotMet *bool `json:"qualityCutoffNotMet,omitempty"`
}

// NewEpisodeFileResource instantiates a new EpisodeFileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisodeFileResource() *EpisodeFileResource {
	this := EpisodeFileResource{}
	return &this
}

// NewEpisodeFileResourceWithDefaults instantiates a new EpisodeFileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodeFileResourceWithDefaults() *EpisodeFileResource {
	this := EpisodeFileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EpisodeFileResource) SetId(v int32) {
	o.Id = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetSeriesId() int32 {
	if o == nil || IsNil(o.SeriesId) {
		var ret int32
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetSeriesIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int32 and assigns it to the SeriesId field.
func (o *EpisodeFileResource) SetSeriesId(v int32) {
	o.SeriesId = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetSeasonNumber() int32 {
	if o == nil || IsNil(o.SeasonNumber) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetSeasonNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SeasonNumber) {
		return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasSeasonNumber() bool {
	if o != nil && !IsNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given int32 and assigns it to the SeasonNumber field.
func (o *EpisodeFileResource) SetSeasonNumber(v int32) {
	o.SeasonNumber = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *EpisodeFileResource) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *EpisodeFileResource) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *EpisodeFileResource) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *EpisodeFileResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *EpisodeFileResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *EpisodeFileResource) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *EpisodeFileResource) SetSize(v int64) {
	o.Size = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetDateAdded() time.Time {
	if o == nil || IsNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *EpisodeFileResource) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetSceneName() string {
	if o == nil || IsNil(o.SceneName.Get()) {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetSceneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *EpisodeFileResource) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *EpisodeFileResource) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *EpisodeFileResource) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *EpisodeFileResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *EpisodeFileResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *EpisodeFileResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *EpisodeFileResource) SetLanguages(v []Language) {
	o.Languages = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *EpisodeFileResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *EpisodeFileResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *EpisodeFileResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileResource) GetIndexerFlags() int32 {
	if o == nil || IsNil(o.IndexerFlags.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags.Get()
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexerFlags.Get(), o.IndexerFlags.IsSet()
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasIndexerFlags() bool {
	if o != nil && o.IndexerFlags.IsSet() {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given NullableInt32 and assigns it to the IndexerFlags field.
func (o *EpisodeFileResource) SetIndexerFlags(v int32) {
	o.IndexerFlags.Set(&v)
}
// SetIndexerFlagsNil sets the value for IndexerFlags to be an explicit nil
func (o *EpisodeFileResource) SetIndexerFlagsNil() {
	o.IndexerFlags.Set(nil)
}

// UnsetIndexerFlags ensures that no value is present for IndexerFlags, not even an explicit nil
func (o *EpisodeFileResource) UnsetIndexerFlags() {
	o.IndexerFlags.Unset()
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetReleaseType() ReleaseType {
	if o == nil || IsNil(o.ReleaseType) {
		var ret ReleaseType
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetReleaseTypeOk() (*ReleaseType, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given ReleaseType and assigns it to the ReleaseType field.
func (o *EpisodeFileResource) SetReleaseType(v ReleaseType) {
	o.ReleaseType = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetMediaInfo() MediaInfoResource {
	if o == nil || IsNil(o.MediaInfo) {
		var ret MediaInfoResource
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetMediaInfoOk() (*MediaInfoResource, bool) {
	if o == nil || IsNil(o.MediaInfo) {
		return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasMediaInfo() bool {
	if o != nil && !IsNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfoResource and assigns it to the MediaInfo field.
func (o *EpisodeFileResource) SetMediaInfo(v MediaInfoResource) {
	o.MediaInfo = &v
}

// GetQualityCutoffNotMet returns the QualityCutoffNotMet field value if set, zero value otherwise.
func (o *EpisodeFileResource) GetQualityCutoffNotMet() bool {
	if o == nil || IsNil(o.QualityCutoffNotMet) {
		var ret bool
		return ret
	}
	return *o.QualityCutoffNotMet
}

// GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileResource) GetQualityCutoffNotMetOk() (*bool, bool) {
	if o == nil || IsNil(o.QualityCutoffNotMet) {
		return nil, false
	}
	return o.QualityCutoffNotMet, true
}

// HasQualityCutoffNotMet returns a boolean if a field has been set.
func (o *EpisodeFileResource) HasQualityCutoffNotMet() bool {
	if o != nil && !IsNil(o.QualityCutoffNotMet) {
		return true
	}

	return false
}

// SetQualityCutoffNotMet gets a reference to the given bool and assigns it to the QualityCutoffNotMet field.
func (o *EpisodeFileResource) SetQualityCutoffNotMet(v bool) {
	o.QualityCutoffNotMet = &v
}

func (o EpisodeFileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpisodeFileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SeriesId) {
		toSerialize["seriesId"] = o.SeriesId
	}
	if !IsNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.DateAdded) {
		toSerialize["dateAdded"] = o.DateAdded
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	if o.IndexerFlags.IsSet() {
		toSerialize["indexerFlags"] = o.IndexerFlags.Get()
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !IsNil(o.QualityCutoffNotMet) {
		toSerialize["qualityCutoffNotMet"] = o.QualityCutoffNotMet
	}
	return toSerialize, nil
}

type NullableEpisodeFileResource struct {
	value *EpisodeFileResource
	isSet bool
}

func (v NullableEpisodeFileResource) Get() *EpisodeFileResource {
	return v.value
}

func (v *NullableEpisodeFileResource) Set(val *EpisodeFileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeFileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeFileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeFileResource(val *EpisodeFileResource) *NullableEpisodeFileResource {
	return &NullableEpisodeFileResource{value: val, isSet: true}
}

func (v NullableEpisodeFileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeFileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


