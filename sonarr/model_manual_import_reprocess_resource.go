/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.4.1491
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the ManualImportReprocessResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualImportReprocessResource{}

// ManualImportReprocessResource struct for ManualImportReprocessResource
type ManualImportReprocessResource struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	SeriesId *int32 `json:"seriesId,omitempty"`
	SeasonNumber NullableInt32 `json:"seasonNumber,omitempty"`
	Episodes []EpisodeResource `json:"episodes,omitempty"`
	EpisodeIds []int32 `json:"episodeIds,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	IndexerFlags *int32 `json:"indexerFlags,omitempty"`
	ReleaseType *ReleaseType `json:"releaseType,omitempty"`
	Rejections []Rejection `json:"rejections,omitempty"`
}

// NewManualImportReprocessResource instantiates a new ManualImportReprocessResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualImportReprocessResource() *ManualImportReprocessResource {
	this := ManualImportReprocessResource{}
	return &this
}

// NewManualImportReprocessResourceWithDefaults instantiates a new ManualImportReprocessResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualImportReprocessResourceWithDefaults() *ManualImportReprocessResource {
	this := ManualImportReprocessResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManualImportReprocessResource) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ManualImportReprocessResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ManualImportReprocessResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ManualImportReprocessResource) UnsetPath() {
	o.Path.Unset()
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetSeriesId() int32 {
	if o == nil || IsNil(o.SeriesId) {
		var ret int32
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetSeriesIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int32 and assigns it to the SeriesId field.
func (o *ManualImportReprocessResource) SetSeriesId(v int32) {
	o.SeriesId = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetSeasonNumber() int32 {
	if o == nil || IsNil(o.SeasonNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber.Get()
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetSeasonNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeasonNumber.Get(), o.SeasonNumber.IsSet()
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasSeasonNumber() bool {
	if o != nil && o.SeasonNumber.IsSet() {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given NullableInt32 and assigns it to the SeasonNumber field.
func (o *ManualImportReprocessResource) SetSeasonNumber(v int32) {
	o.SeasonNumber.Set(&v)
}
// SetSeasonNumberNil sets the value for SeasonNumber to be an explicit nil
func (o *ManualImportReprocessResource) SetSeasonNumberNil() {
	o.SeasonNumber.Set(nil)
}

// UnsetSeasonNumber ensures that no value is present for SeasonNumber, not even an explicit nil
func (o *ManualImportReprocessResource) UnsetSeasonNumber() {
	o.SeasonNumber.Unset()
}

// GetEpisodes returns the Episodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetEpisodes() []EpisodeResource {
	if o == nil {
		var ret []EpisodeResource
		return ret
	}
	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetEpisodesOk() ([]EpisodeResource, bool) {
	if o == nil || IsNil(o.Episodes) {
		return nil, false
	}
	return o.Episodes, true
}

// HasEpisodes returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasEpisodes() bool {
	if o != nil && !IsNil(o.Episodes) {
		return true
	}

	return false
}

// SetEpisodes gets a reference to the given []EpisodeResource and assigns it to the Episodes field.
func (o *ManualImportReprocessResource) SetEpisodes(v []EpisodeResource) {
	o.Episodes = v
}

// GetEpisodeIds returns the EpisodeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetEpisodeIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.EpisodeIds
}

// GetEpisodeIdsOk returns a tuple with the EpisodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetEpisodeIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.EpisodeIds) {
		return nil, false
	}
	return o.EpisodeIds, true
}

// HasEpisodeIds returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasEpisodeIds() bool {
	if o != nil && !IsNil(o.EpisodeIds) {
		return true
	}

	return false
}

// SetEpisodeIds gets a reference to the given []int32 and assigns it to the EpisodeIds field.
func (o *ManualImportReprocessResource) SetEpisodeIds(v []int32) {
	o.EpisodeIds = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ManualImportReprocessResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ManualImportReprocessResource) SetLanguages(v []Language) {
	o.Languages = v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ManualImportReprocessResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ManualImportReprocessResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ManualImportReprocessResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *ManualImportReprocessResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *ManualImportReprocessResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *ManualImportReprocessResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *ManualImportReprocessResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *ManualImportReprocessResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetIndexerFlags() int32 {
	if o == nil || IsNil(o.IndexerFlags) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerFlags) {
		return nil, false
	}
	return o.IndexerFlags, true
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasIndexerFlags() bool {
	if o != nil && !IsNil(o.IndexerFlags) {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given int32 and assigns it to the IndexerFlags field.
func (o *ManualImportReprocessResource) SetIndexerFlags(v int32) {
	o.IndexerFlags = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *ManualImportReprocessResource) GetReleaseType() ReleaseType {
	if o == nil || IsNil(o.ReleaseType) {
		var ret ReleaseType
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportReprocessResource) GetReleaseTypeOk() (*ReleaseType, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given ReleaseType and assigns it to the ReleaseType field.
func (o *ManualImportReprocessResource) SetReleaseType(v ReleaseType) {
	o.ReleaseType = &v
}

// GetRejections returns the Rejections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportReprocessResource) GetRejections() []Rejection {
	if o == nil {
		var ret []Rejection
		return ret
	}
	return o.Rejections
}

// GetRejectionsOk returns a tuple with the Rejections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportReprocessResource) GetRejectionsOk() ([]Rejection, bool) {
	if o == nil || IsNil(o.Rejections) {
		return nil, false
	}
	return o.Rejections, true
}

// HasRejections returns a boolean if a field has been set.
func (o *ManualImportReprocessResource) HasRejections() bool {
	if o != nil && !IsNil(o.Rejections) {
		return true
	}

	return false
}

// SetRejections gets a reference to the given []Rejection and assigns it to the Rejections field.
func (o *ManualImportReprocessResource) SetRejections(v []Rejection) {
	o.Rejections = v
}

func (o ManualImportReprocessResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualImportReprocessResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.SeriesId) {
		toSerialize["seriesId"] = o.SeriesId
	}
	if o.SeasonNumber.IsSet() {
		toSerialize["seasonNumber"] = o.SeasonNumber.Get()
	}
	if o.Episodes != nil {
		toSerialize["episodes"] = o.Episodes
	}
	if o.EpisodeIds != nil {
		toSerialize["episodeIds"] = o.EpisodeIds
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	if !IsNil(o.IndexerFlags) {
		toSerialize["indexerFlags"] = o.IndexerFlags
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if o.Rejections != nil {
		toSerialize["rejections"] = o.Rejections
	}
	return toSerialize, nil
}

type NullableManualImportReprocessResource struct {
	value *ManualImportReprocessResource
	isSet bool
}

func (v NullableManualImportReprocessResource) Get() *ManualImportReprocessResource {
	return v.value
}

func (v *NullableManualImportReprocessResource) Set(val *ManualImportReprocessResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManualImportReprocessResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManualImportReprocessResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualImportReprocessResource(val *ManualImportReprocessResource) *NullableManualImportReprocessResource {
	return &NullableManualImportReprocessResource{value: val, isSet: true}
}

func (v NullableManualImportReprocessResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualImportReprocessResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


