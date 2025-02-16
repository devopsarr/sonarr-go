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

// checks if the ManualImportResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualImportResource{}

// ManualImportResource struct for ManualImportResource
type ManualImportResource struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	RelativePath NullableString `json:"relativePath,omitempty"`
	FolderName NullableString `json:"folderName,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Series *SeriesResource `json:"series,omitempty"`
	SeasonNumber NullableInt32 `json:"seasonNumber,omitempty"`
	Episodes []EpisodeResource `json:"episodes,omitempty"`
	EpisodeFileId NullableInt32 `json:"episodeFileId,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	QualityWeight *int32 `json:"qualityWeight,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	IndexerFlags *int32 `json:"indexerFlags,omitempty"`
	ReleaseType *ReleaseType `json:"releaseType,omitempty"`
	Rejections []ImportRejectionResource `json:"rejections,omitempty"`
}

// NewManualImportResource instantiates a new ManualImportResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualImportResource() *ManualImportResource {
	this := ManualImportResource{}
	return &this
}

// NewManualImportResourceWithDefaults instantiates a new ManualImportResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualImportResourceWithDefaults() *ManualImportResource {
	this := ManualImportResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManualImportResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManualImportResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManualImportResource) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ManualImportResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ManualImportResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ManualImportResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ManualImportResource) UnsetPath() {
	o.Path.Unset()
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *ManualImportResource) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *ManualImportResource) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *ManualImportResource) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *ManualImportResource) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetFolderName returns the FolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetFolderName() string {
	if o == nil || IsNil(o.FolderName.Get()) {
		var ret string
		return ret
	}
	return *o.FolderName.Get()
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetFolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderName.Get(), o.FolderName.IsSet()
}

// HasFolderName returns a boolean if a field has been set.
func (o *ManualImportResource) HasFolderName() bool {
	if o != nil && o.FolderName.IsSet() {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given NullableString and assigns it to the FolderName field.
func (o *ManualImportResource) SetFolderName(v string) {
	o.FolderName.Set(&v)
}
// SetFolderNameNil sets the value for FolderName to be an explicit nil
func (o *ManualImportResource) SetFolderNameNil() {
	o.FolderName.Set(nil)
}

// UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
func (o *ManualImportResource) UnsetFolderName() {
	o.FolderName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ManualImportResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ManualImportResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ManualImportResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ManualImportResource) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ManualImportResource) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ManualImportResource) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ManualImportResource) SetSize(v int64) {
	o.Size = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *ManualImportResource) GetSeries() SeriesResource {
	if o == nil || IsNil(o.Series) {
		var ret SeriesResource
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetSeriesOk() (*SeriesResource, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *ManualImportResource) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given SeriesResource and assigns it to the Series field.
func (o *ManualImportResource) SetSeries(v SeriesResource) {
	o.Series = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetSeasonNumber() int32 {
	if o == nil || IsNil(o.SeasonNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.SeasonNumber.Get()
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetSeasonNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeasonNumber.Get(), o.SeasonNumber.IsSet()
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *ManualImportResource) HasSeasonNumber() bool {
	if o != nil && o.SeasonNumber.IsSet() {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given NullableInt32 and assigns it to the SeasonNumber field.
func (o *ManualImportResource) SetSeasonNumber(v int32) {
	o.SeasonNumber.Set(&v)
}
// SetSeasonNumberNil sets the value for SeasonNumber to be an explicit nil
func (o *ManualImportResource) SetSeasonNumberNil() {
	o.SeasonNumber.Set(nil)
}

// UnsetSeasonNumber ensures that no value is present for SeasonNumber, not even an explicit nil
func (o *ManualImportResource) UnsetSeasonNumber() {
	o.SeasonNumber.Unset()
}

// GetEpisodes returns the Episodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetEpisodes() []EpisodeResource {
	if o == nil {
		var ret []EpisodeResource
		return ret
	}
	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetEpisodesOk() ([]EpisodeResource, bool) {
	if o == nil || IsNil(o.Episodes) {
		return nil, false
	}
	return o.Episodes, true
}

// HasEpisodes returns a boolean if a field has been set.
func (o *ManualImportResource) HasEpisodes() bool {
	if o != nil && !IsNil(o.Episodes) {
		return true
	}

	return false
}

// SetEpisodes gets a reference to the given []EpisodeResource and assigns it to the Episodes field.
func (o *ManualImportResource) SetEpisodes(v []EpisodeResource) {
	o.Episodes = v
}

// GetEpisodeFileId returns the EpisodeFileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetEpisodeFileId() int32 {
	if o == nil || IsNil(o.EpisodeFileId.Get()) {
		var ret int32
		return ret
	}
	return *o.EpisodeFileId.Get()
}

// GetEpisodeFileIdOk returns a tuple with the EpisodeFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetEpisodeFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpisodeFileId.Get(), o.EpisodeFileId.IsSet()
}

// HasEpisodeFileId returns a boolean if a field has been set.
func (o *ManualImportResource) HasEpisodeFileId() bool {
	if o != nil && o.EpisodeFileId.IsSet() {
		return true
	}

	return false
}

// SetEpisodeFileId gets a reference to the given NullableInt32 and assigns it to the EpisodeFileId field.
func (o *ManualImportResource) SetEpisodeFileId(v int32) {
	o.EpisodeFileId.Set(&v)
}
// SetEpisodeFileIdNil sets the value for EpisodeFileId to be an explicit nil
func (o *ManualImportResource) SetEpisodeFileIdNil() {
	o.EpisodeFileId.Set(nil)
}

// UnsetEpisodeFileId ensures that no value is present for EpisodeFileId, not even an explicit nil
func (o *ManualImportResource) UnsetEpisodeFileId() {
	o.EpisodeFileId.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ManualImportResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ManualImportResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ManualImportResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ManualImportResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ManualImportResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ManualImportResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ManualImportResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManualImportResource) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ManualImportResource) SetLanguages(v []Language) {
	o.Languages = v
}

// GetQualityWeight returns the QualityWeight field value if set, zero value otherwise.
func (o *ManualImportResource) GetQualityWeight() int32 {
	if o == nil || IsNil(o.QualityWeight) {
		var ret int32
		return ret
	}
	return *o.QualityWeight
}

// GetQualityWeightOk returns a tuple with the QualityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetQualityWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityWeight) {
		return nil, false
	}
	return o.QualityWeight, true
}

// HasQualityWeight returns a boolean if a field has been set.
func (o *ManualImportResource) HasQualityWeight() bool {
	if o != nil && !IsNil(o.QualityWeight) {
		return true
	}

	return false
}

// SetQualityWeight gets a reference to the given int32 and assigns it to the QualityWeight field.
func (o *ManualImportResource) SetQualityWeight(v int32) {
	o.QualityWeight = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *ManualImportResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *ManualImportResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *ManualImportResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *ManualImportResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *ManualImportResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *ManualImportResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *ManualImportResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *ManualImportResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *ManualImportResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise.
func (o *ManualImportResource) GetIndexerFlags() int32 {
	if o == nil || IsNil(o.IndexerFlags) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.IndexerFlags) {
		return nil, false
	}
	return o.IndexerFlags, true
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *ManualImportResource) HasIndexerFlags() bool {
	if o != nil && !IsNil(o.IndexerFlags) {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given int32 and assigns it to the IndexerFlags field.
func (o *ManualImportResource) SetIndexerFlags(v int32) {
	o.IndexerFlags = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *ManualImportResource) GetReleaseType() ReleaseType {
	if o == nil || IsNil(o.ReleaseType) {
		var ret ReleaseType
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetReleaseTypeOk() (*ReleaseType, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *ManualImportResource) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given ReleaseType and assigns it to the ReleaseType field.
func (o *ManualImportResource) SetReleaseType(v ReleaseType) {
	o.ReleaseType = &v
}

// GetRejections returns the Rejections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetRejections() []ImportRejectionResource {
	if o == nil {
		var ret []ImportRejectionResource
		return ret
	}
	return o.Rejections
}

// GetRejectionsOk returns a tuple with the Rejections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetRejectionsOk() ([]ImportRejectionResource, bool) {
	if o == nil || IsNil(o.Rejections) {
		return nil, false
	}
	return o.Rejections, true
}

// HasRejections returns a boolean if a field has been set.
func (o *ManualImportResource) HasRejections() bool {
	if o != nil && !IsNil(o.Rejections) {
		return true
	}

	return false
}

// SetRejections gets a reference to the given []ImportRejectionResource and assigns it to the Rejections field.
func (o *ManualImportResource) SetRejections(v []ImportRejectionResource) {
	o.Rejections = v
}

func (o ManualImportResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualImportResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	if o.FolderName.IsSet() {
		toSerialize["folderName"] = o.FolderName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if o.SeasonNumber.IsSet() {
		toSerialize["seasonNumber"] = o.SeasonNumber.Get()
	}
	if o.Episodes != nil {
		toSerialize["episodes"] = o.Episodes
	}
	if o.EpisodeFileId.IsSet() {
		toSerialize["episodeFileId"] = o.EpisodeFileId.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.QualityWeight) {
		toSerialize["qualityWeight"] = o.QualityWeight
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

type NullableManualImportResource struct {
	value *ManualImportResource
	isSet bool
}

func (v NullableManualImportResource) Get() *ManualImportResource {
	return v.value
}

func (v *NullableManualImportResource) Set(val *ManualImportResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManualImportResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManualImportResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualImportResource(val *ManualImportResource) *NullableManualImportResource {
	return &NullableManualImportResource{value: val, isSet: true}
}

func (v NullableManualImportResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualImportResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


