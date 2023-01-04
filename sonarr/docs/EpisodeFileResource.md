# EpisodeFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SeriesId** | Pointer to **int32** |  | [optional] 
**SeasonNumber** | Pointer to **int32** |  | [optional] 
**RelativePath** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoResource**](MediaInfoResource.md) |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 

## Methods

### NewEpisodeFileResource

`func NewEpisodeFileResource() *EpisodeFileResource`

NewEpisodeFileResource instantiates a new EpisodeFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeFileResourceWithDefaults

`func NewEpisodeFileResourceWithDefaults() *EpisodeFileResource`

NewEpisodeFileResourceWithDefaults instantiates a new EpisodeFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EpisodeFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EpisodeFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EpisodeFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EpisodeFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeriesId

`func (o *EpisodeFileResource) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *EpisodeFileResource) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *EpisodeFileResource) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *EpisodeFileResource) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *EpisodeFileResource) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *EpisodeFileResource) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *EpisodeFileResource) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *EpisodeFileResource) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetRelativePath

`func (o *EpisodeFileResource) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *EpisodeFileResource) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *EpisodeFileResource) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *EpisodeFileResource) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *EpisodeFileResource) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *EpisodeFileResource) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetPath

`func (o *EpisodeFileResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EpisodeFileResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EpisodeFileResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EpisodeFileResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *EpisodeFileResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *EpisodeFileResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *EpisodeFileResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EpisodeFileResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EpisodeFileResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *EpisodeFileResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDateAdded

`func (o *EpisodeFileResource) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *EpisodeFileResource) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *EpisodeFileResource) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *EpisodeFileResource) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSceneName

`func (o *EpisodeFileResource) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *EpisodeFileResource) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *EpisodeFileResource) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *EpisodeFileResource) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *EpisodeFileResource) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *EpisodeFileResource) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetReleaseGroup

`func (o *EpisodeFileResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *EpisodeFileResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *EpisodeFileResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *EpisodeFileResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *EpisodeFileResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *EpisodeFileResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetLanguages

`func (o *EpisodeFileResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *EpisodeFileResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *EpisodeFileResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *EpisodeFileResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *EpisodeFileResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *EpisodeFileResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetQuality

`func (o *EpisodeFileResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *EpisodeFileResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *EpisodeFileResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *EpisodeFileResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *EpisodeFileResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *EpisodeFileResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *EpisodeFileResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *EpisodeFileResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *EpisodeFileResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *EpisodeFileResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetMediaInfo

`func (o *EpisodeFileResource) GetMediaInfo() MediaInfoResource`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *EpisodeFileResource) GetMediaInfoOk() (*MediaInfoResource, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *EpisodeFileResource) SetMediaInfo(v MediaInfoResource)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *EpisodeFileResource) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetQualityCutoffNotMet

`func (o *EpisodeFileResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *EpisodeFileResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *EpisodeFileResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *EpisodeFileResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


