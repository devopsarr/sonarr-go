# ManualImportReprocessResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**SeriesId** | Pointer to **int32** |  | [optional] 
**SeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**Episodes** | Pointer to [**[]EpisodeResource**](EpisodeResource.md) |  | [optional] 
**EpisodeIds** | Pointer to **[]int32** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 
**IndexerFlags** | Pointer to **int32** |  | [optional] 
**ReleaseType** | Pointer to [**ReleaseType**](ReleaseType.md) |  | [optional] 
**Rejections** | Pointer to [**[]ImportRejectionResource**](ImportRejectionResource.md) |  | [optional] 

## Methods

### NewManualImportReprocessResource

`func NewManualImportReprocessResource() *ManualImportReprocessResource`

NewManualImportReprocessResource instantiates a new ManualImportReprocessResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportReprocessResourceWithDefaults

`func NewManualImportReprocessResourceWithDefaults() *ManualImportReprocessResource`

NewManualImportReprocessResourceWithDefaults instantiates a new ManualImportReprocessResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportReprocessResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportReprocessResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportReprocessResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportReprocessResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportReprocessResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportReprocessResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportReprocessResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportReprocessResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportReprocessResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportReprocessResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSeriesId

`func (o *ManualImportReprocessResource) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *ManualImportReprocessResource) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *ManualImportReprocessResource) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *ManualImportReprocessResource) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *ManualImportReprocessResource) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *ManualImportReprocessResource) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *ManualImportReprocessResource) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *ManualImportReprocessResource) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### SetSeasonNumberNil

`func (o *ManualImportReprocessResource) SetSeasonNumberNil(b bool)`

 SetSeasonNumberNil sets the value for SeasonNumber to be an explicit nil

### UnsetSeasonNumber
`func (o *ManualImportReprocessResource) UnsetSeasonNumber()`

UnsetSeasonNumber ensures that no value is present for SeasonNumber, not even an explicit nil
### GetEpisodes

`func (o *ManualImportReprocessResource) GetEpisodes() []EpisodeResource`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *ManualImportReprocessResource) GetEpisodesOk() (*[]EpisodeResource, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *ManualImportReprocessResource) SetEpisodes(v []EpisodeResource)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *ManualImportReprocessResource) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.

### SetEpisodesNil

`func (o *ManualImportReprocessResource) SetEpisodesNil(b bool)`

 SetEpisodesNil sets the value for Episodes to be an explicit nil

### UnsetEpisodes
`func (o *ManualImportReprocessResource) UnsetEpisodes()`

UnsetEpisodes ensures that no value is present for Episodes, not even an explicit nil
### GetEpisodeIds

`func (o *ManualImportReprocessResource) GetEpisodeIds() []int32`

GetEpisodeIds returns the EpisodeIds field if non-nil, zero value otherwise.

### GetEpisodeIdsOk

`func (o *ManualImportReprocessResource) GetEpisodeIdsOk() (*[]int32, bool)`

GetEpisodeIdsOk returns a tuple with the EpisodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeIds

`func (o *ManualImportReprocessResource) SetEpisodeIds(v []int32)`

SetEpisodeIds sets EpisodeIds field to given value.

### HasEpisodeIds

`func (o *ManualImportReprocessResource) HasEpisodeIds() bool`

HasEpisodeIds returns a boolean if a field has been set.

### SetEpisodeIdsNil

`func (o *ManualImportReprocessResource) SetEpisodeIdsNil(b bool)`

 SetEpisodeIdsNil sets the value for EpisodeIds to be an explicit nil

### UnsetEpisodeIds
`func (o *ManualImportReprocessResource) UnsetEpisodeIds()`

UnsetEpisodeIds ensures that no value is present for EpisodeIds, not even an explicit nil
### GetQuality

`func (o *ManualImportReprocessResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportReprocessResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportReprocessResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportReprocessResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetLanguages

`func (o *ManualImportReprocessResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ManualImportReprocessResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ManualImportReprocessResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ManualImportReprocessResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *ManualImportReprocessResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *ManualImportReprocessResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetReleaseGroup

`func (o *ManualImportReprocessResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ManualImportReprocessResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ManualImportReprocessResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ManualImportReprocessResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ManualImportReprocessResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ManualImportReprocessResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetDownloadId

`func (o *ManualImportReprocessResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportReprocessResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportReprocessResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportReprocessResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportReprocessResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportReprocessResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetCustomFormats

`func (o *ManualImportReprocessResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *ManualImportReprocessResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *ManualImportReprocessResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *ManualImportReprocessResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *ManualImportReprocessResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *ManualImportReprocessResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *ManualImportReprocessResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *ManualImportReprocessResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *ManualImportReprocessResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *ManualImportReprocessResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.

### GetIndexerFlags

`func (o *ManualImportReprocessResource) GetIndexerFlags() int32`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *ManualImportReprocessResource) GetIndexerFlagsOk() (*int32, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *ManualImportReprocessResource) SetIndexerFlags(v int32)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *ManualImportReprocessResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### GetReleaseType

`func (o *ManualImportReprocessResource) GetReleaseType() ReleaseType`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *ManualImportReprocessResource) GetReleaseTypeOk() (*ReleaseType, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *ManualImportReprocessResource) SetReleaseType(v ReleaseType)`

SetReleaseType sets ReleaseType field to given value.

### HasReleaseType

`func (o *ManualImportReprocessResource) HasReleaseType() bool`

HasReleaseType returns a boolean if a field has been set.

### GetRejections

`func (o *ManualImportReprocessResource) GetRejections() []ImportRejectionResource`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportReprocessResource) GetRejectionsOk() (*[]ImportRejectionResource, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportReprocessResource) SetRejections(v []ImportRejectionResource)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportReprocessResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportReprocessResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportReprocessResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


