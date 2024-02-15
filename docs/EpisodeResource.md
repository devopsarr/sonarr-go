# EpisodeResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SeriesId** | Pointer to **int32** |  | [optional] 
**TvdbId** | Pointer to **int32** |  | [optional] 
**EpisodeFileId** | Pointer to **int32** |  | [optional] 
**SeasonNumber** | Pointer to **int32** |  | [optional] 
**EpisodeNumber** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**AirDate** | Pointer to **NullableString** |  | [optional] 
**AirDateUtc** | Pointer to **NullableTime** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**FinaleType** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**EpisodeFile** | Pointer to [**EpisodeFileResource**](EpisodeFileResource.md) |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**AbsoluteEpisodeNumber** | Pointer to **NullableInt32** |  | [optional] 
**SceneAbsoluteEpisodeNumber** | Pointer to **NullableInt32** |  | [optional] 
**SceneEpisodeNumber** | Pointer to **NullableInt32** |  | [optional] 
**SceneSeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**UnverifiedSceneNumbering** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**GrabDate** | Pointer to **NullableTime** |  | [optional] 
**SeriesTitle** | Pointer to **NullableString** |  | [optional] 
**Series** | Pointer to [**SeriesResource**](SeriesResource.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Grabbed** | Pointer to **bool** |  | [optional] 

## Methods

### NewEpisodeResource

`func NewEpisodeResource() *EpisodeResource`

NewEpisodeResource instantiates a new EpisodeResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeResourceWithDefaults

`func NewEpisodeResourceWithDefaults() *EpisodeResource`

NewEpisodeResourceWithDefaults instantiates a new EpisodeResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EpisodeResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EpisodeResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EpisodeResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EpisodeResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeriesId

`func (o *EpisodeResource) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *EpisodeResource) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *EpisodeResource) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *EpisodeResource) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetTvdbId

`func (o *EpisodeResource) GetTvdbId() int32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *EpisodeResource) GetTvdbIdOk() (*int32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *EpisodeResource) SetTvdbId(v int32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *EpisodeResource) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### GetEpisodeFileId

`func (o *EpisodeResource) GetEpisodeFileId() int32`

GetEpisodeFileId returns the EpisodeFileId field if non-nil, zero value otherwise.

### GetEpisodeFileIdOk

`func (o *EpisodeResource) GetEpisodeFileIdOk() (*int32, bool)`

GetEpisodeFileIdOk returns a tuple with the EpisodeFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFileId

`func (o *EpisodeResource) SetEpisodeFileId(v int32)`

SetEpisodeFileId sets EpisodeFileId field to given value.

### HasEpisodeFileId

`func (o *EpisodeResource) HasEpisodeFileId() bool`

HasEpisodeFileId returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *EpisodeResource) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *EpisodeResource) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *EpisodeResource) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *EpisodeResource) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetEpisodeNumber

`func (o *EpisodeResource) GetEpisodeNumber() int32`

GetEpisodeNumber returns the EpisodeNumber field if non-nil, zero value otherwise.

### GetEpisodeNumberOk

`func (o *EpisodeResource) GetEpisodeNumberOk() (*int32, bool)`

GetEpisodeNumberOk returns a tuple with the EpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeNumber

`func (o *EpisodeResource) SetEpisodeNumber(v int32)`

SetEpisodeNumber sets EpisodeNumber field to given value.

### HasEpisodeNumber

`func (o *EpisodeResource) HasEpisodeNumber() bool`

HasEpisodeNumber returns a boolean if a field has been set.

### GetTitle

`func (o *EpisodeResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EpisodeResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EpisodeResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EpisodeResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EpisodeResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EpisodeResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAirDate

`func (o *EpisodeResource) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *EpisodeResource) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *EpisodeResource) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *EpisodeResource) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### SetAirDateNil

`func (o *EpisodeResource) SetAirDateNil(b bool)`

 SetAirDateNil sets the value for AirDate to be an explicit nil

### UnsetAirDate
`func (o *EpisodeResource) UnsetAirDate()`

UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
### GetAirDateUtc

`func (o *EpisodeResource) GetAirDateUtc() time.Time`

GetAirDateUtc returns the AirDateUtc field if non-nil, zero value otherwise.

### GetAirDateUtcOk

`func (o *EpisodeResource) GetAirDateUtcOk() (*time.Time, bool)`

GetAirDateUtcOk returns a tuple with the AirDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDateUtc

`func (o *EpisodeResource) SetAirDateUtc(v time.Time)`

SetAirDateUtc sets AirDateUtc field to given value.

### HasAirDateUtc

`func (o *EpisodeResource) HasAirDateUtc() bool`

HasAirDateUtc returns a boolean if a field has been set.

### SetAirDateUtcNil

`func (o *EpisodeResource) SetAirDateUtcNil(b bool)`

 SetAirDateUtcNil sets the value for AirDateUtc to be an explicit nil

### UnsetAirDateUtc
`func (o *EpisodeResource) UnsetAirDateUtc()`

UnsetAirDateUtc ensures that no value is present for AirDateUtc, not even an explicit nil
### GetRuntime

`func (o *EpisodeResource) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *EpisodeResource) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *EpisodeResource) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *EpisodeResource) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetFinaleType

`func (o *EpisodeResource) GetFinaleType() string`

GetFinaleType returns the FinaleType field if non-nil, zero value otherwise.

### GetFinaleTypeOk

`func (o *EpisodeResource) GetFinaleTypeOk() (*string, bool)`

GetFinaleTypeOk returns a tuple with the FinaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinaleType

`func (o *EpisodeResource) SetFinaleType(v string)`

SetFinaleType sets FinaleType field to given value.

### HasFinaleType

`func (o *EpisodeResource) HasFinaleType() bool`

HasFinaleType returns a boolean if a field has been set.

### SetFinaleTypeNil

`func (o *EpisodeResource) SetFinaleTypeNil(b bool)`

 SetFinaleTypeNil sets the value for FinaleType to be an explicit nil

### UnsetFinaleType
`func (o *EpisodeResource) UnsetFinaleType()`

UnsetFinaleType ensures that no value is present for FinaleType, not even an explicit nil
### GetOverview

`func (o *EpisodeResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *EpisodeResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *EpisodeResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *EpisodeResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *EpisodeResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *EpisodeResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetEpisodeFile

`func (o *EpisodeResource) GetEpisodeFile() EpisodeFileResource`

GetEpisodeFile returns the EpisodeFile field if non-nil, zero value otherwise.

### GetEpisodeFileOk

`func (o *EpisodeResource) GetEpisodeFileOk() (*EpisodeFileResource, bool)`

GetEpisodeFileOk returns a tuple with the EpisodeFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFile

`func (o *EpisodeResource) SetEpisodeFile(v EpisodeFileResource)`

SetEpisodeFile sets EpisodeFile field to given value.

### HasEpisodeFile

`func (o *EpisodeResource) HasEpisodeFile() bool`

HasEpisodeFile returns a boolean if a field has been set.

### GetHasFile

`func (o *EpisodeResource) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *EpisodeResource) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *EpisodeResource) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *EpisodeResource) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetMonitored

`func (o *EpisodeResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *EpisodeResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *EpisodeResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *EpisodeResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetAbsoluteEpisodeNumber

`func (o *EpisodeResource) GetAbsoluteEpisodeNumber() int32`

GetAbsoluteEpisodeNumber returns the AbsoluteEpisodeNumber field if non-nil, zero value otherwise.

### GetAbsoluteEpisodeNumberOk

`func (o *EpisodeResource) GetAbsoluteEpisodeNumberOk() (*int32, bool)`

GetAbsoluteEpisodeNumberOk returns a tuple with the AbsoluteEpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteEpisodeNumber

`func (o *EpisodeResource) SetAbsoluteEpisodeNumber(v int32)`

SetAbsoluteEpisodeNumber sets AbsoluteEpisodeNumber field to given value.

### HasAbsoluteEpisodeNumber

`func (o *EpisodeResource) HasAbsoluteEpisodeNumber() bool`

HasAbsoluteEpisodeNumber returns a boolean if a field has been set.

### SetAbsoluteEpisodeNumberNil

`func (o *EpisodeResource) SetAbsoluteEpisodeNumberNil(b bool)`

 SetAbsoluteEpisodeNumberNil sets the value for AbsoluteEpisodeNumber to be an explicit nil

### UnsetAbsoluteEpisodeNumber
`func (o *EpisodeResource) UnsetAbsoluteEpisodeNumber()`

UnsetAbsoluteEpisodeNumber ensures that no value is present for AbsoluteEpisodeNumber, not even an explicit nil
### GetSceneAbsoluteEpisodeNumber

`func (o *EpisodeResource) GetSceneAbsoluteEpisodeNumber() int32`

GetSceneAbsoluteEpisodeNumber returns the SceneAbsoluteEpisodeNumber field if non-nil, zero value otherwise.

### GetSceneAbsoluteEpisodeNumberOk

`func (o *EpisodeResource) GetSceneAbsoluteEpisodeNumberOk() (*int32, bool)`

GetSceneAbsoluteEpisodeNumberOk returns a tuple with the SceneAbsoluteEpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneAbsoluteEpisodeNumber

`func (o *EpisodeResource) SetSceneAbsoluteEpisodeNumber(v int32)`

SetSceneAbsoluteEpisodeNumber sets SceneAbsoluteEpisodeNumber field to given value.

### HasSceneAbsoluteEpisodeNumber

`func (o *EpisodeResource) HasSceneAbsoluteEpisodeNumber() bool`

HasSceneAbsoluteEpisodeNumber returns a boolean if a field has been set.

### SetSceneAbsoluteEpisodeNumberNil

`func (o *EpisodeResource) SetSceneAbsoluteEpisodeNumberNil(b bool)`

 SetSceneAbsoluteEpisodeNumberNil sets the value for SceneAbsoluteEpisodeNumber to be an explicit nil

### UnsetSceneAbsoluteEpisodeNumber
`func (o *EpisodeResource) UnsetSceneAbsoluteEpisodeNumber()`

UnsetSceneAbsoluteEpisodeNumber ensures that no value is present for SceneAbsoluteEpisodeNumber, not even an explicit nil
### GetSceneEpisodeNumber

`func (o *EpisodeResource) GetSceneEpisodeNumber() int32`

GetSceneEpisodeNumber returns the SceneEpisodeNumber field if non-nil, zero value otherwise.

### GetSceneEpisodeNumberOk

`func (o *EpisodeResource) GetSceneEpisodeNumberOk() (*int32, bool)`

GetSceneEpisodeNumberOk returns a tuple with the SceneEpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneEpisodeNumber

`func (o *EpisodeResource) SetSceneEpisodeNumber(v int32)`

SetSceneEpisodeNumber sets SceneEpisodeNumber field to given value.

### HasSceneEpisodeNumber

`func (o *EpisodeResource) HasSceneEpisodeNumber() bool`

HasSceneEpisodeNumber returns a boolean if a field has been set.

### SetSceneEpisodeNumberNil

`func (o *EpisodeResource) SetSceneEpisodeNumberNil(b bool)`

 SetSceneEpisodeNumberNil sets the value for SceneEpisodeNumber to be an explicit nil

### UnsetSceneEpisodeNumber
`func (o *EpisodeResource) UnsetSceneEpisodeNumber()`

UnsetSceneEpisodeNumber ensures that no value is present for SceneEpisodeNumber, not even an explicit nil
### GetSceneSeasonNumber

`func (o *EpisodeResource) GetSceneSeasonNumber() int32`

GetSceneSeasonNumber returns the SceneSeasonNumber field if non-nil, zero value otherwise.

### GetSceneSeasonNumberOk

`func (o *EpisodeResource) GetSceneSeasonNumberOk() (*int32, bool)`

GetSceneSeasonNumberOk returns a tuple with the SceneSeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneSeasonNumber

`func (o *EpisodeResource) SetSceneSeasonNumber(v int32)`

SetSceneSeasonNumber sets SceneSeasonNumber field to given value.

### HasSceneSeasonNumber

`func (o *EpisodeResource) HasSceneSeasonNumber() bool`

HasSceneSeasonNumber returns a boolean if a field has been set.

### SetSceneSeasonNumberNil

`func (o *EpisodeResource) SetSceneSeasonNumberNil(b bool)`

 SetSceneSeasonNumberNil sets the value for SceneSeasonNumber to be an explicit nil

### UnsetSceneSeasonNumber
`func (o *EpisodeResource) UnsetSceneSeasonNumber()`

UnsetSceneSeasonNumber ensures that no value is present for SceneSeasonNumber, not even an explicit nil
### GetUnverifiedSceneNumbering

`func (o *EpisodeResource) GetUnverifiedSceneNumbering() bool`

GetUnverifiedSceneNumbering returns the UnverifiedSceneNumbering field if non-nil, zero value otherwise.

### GetUnverifiedSceneNumberingOk

`func (o *EpisodeResource) GetUnverifiedSceneNumberingOk() (*bool, bool)`

GetUnverifiedSceneNumberingOk returns a tuple with the UnverifiedSceneNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedSceneNumbering

`func (o *EpisodeResource) SetUnverifiedSceneNumbering(v bool)`

SetUnverifiedSceneNumbering sets UnverifiedSceneNumbering field to given value.

### HasUnverifiedSceneNumbering

`func (o *EpisodeResource) HasUnverifiedSceneNumbering() bool`

HasUnverifiedSceneNumbering returns a boolean if a field has been set.

### GetEndTime

`func (o *EpisodeResource) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EpisodeResource) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EpisodeResource) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EpisodeResource) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *EpisodeResource) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *EpisodeResource) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetGrabDate

`func (o *EpisodeResource) GetGrabDate() time.Time`

GetGrabDate returns the GrabDate field if non-nil, zero value otherwise.

### GetGrabDateOk

`func (o *EpisodeResource) GetGrabDateOk() (*time.Time, bool)`

GetGrabDateOk returns a tuple with the GrabDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabDate

`func (o *EpisodeResource) SetGrabDate(v time.Time)`

SetGrabDate sets GrabDate field to given value.

### HasGrabDate

`func (o *EpisodeResource) HasGrabDate() bool`

HasGrabDate returns a boolean if a field has been set.

### SetGrabDateNil

`func (o *EpisodeResource) SetGrabDateNil(b bool)`

 SetGrabDateNil sets the value for GrabDate to be an explicit nil

### UnsetGrabDate
`func (o *EpisodeResource) UnsetGrabDate()`

UnsetGrabDate ensures that no value is present for GrabDate, not even an explicit nil
### GetSeriesTitle

`func (o *EpisodeResource) GetSeriesTitle() string`

GetSeriesTitle returns the SeriesTitle field if non-nil, zero value otherwise.

### GetSeriesTitleOk

`func (o *EpisodeResource) GetSeriesTitleOk() (*string, bool)`

GetSeriesTitleOk returns a tuple with the SeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitle

`func (o *EpisodeResource) SetSeriesTitle(v string)`

SetSeriesTitle sets SeriesTitle field to given value.

### HasSeriesTitle

`func (o *EpisodeResource) HasSeriesTitle() bool`

HasSeriesTitle returns a boolean if a field has been set.

### SetSeriesTitleNil

`func (o *EpisodeResource) SetSeriesTitleNil(b bool)`

 SetSeriesTitleNil sets the value for SeriesTitle to be an explicit nil

### UnsetSeriesTitle
`func (o *EpisodeResource) UnsetSeriesTitle()`

UnsetSeriesTitle ensures that no value is present for SeriesTitle, not even an explicit nil
### GetSeries

`func (o *EpisodeResource) GetSeries() SeriesResource`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *EpisodeResource) GetSeriesOk() (*SeriesResource, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *EpisodeResource) SetSeries(v SeriesResource)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *EpisodeResource) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetImages

`func (o *EpisodeResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *EpisodeResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *EpisodeResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *EpisodeResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *EpisodeResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *EpisodeResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetGrabbed

`func (o *EpisodeResource) GetGrabbed() bool`

GetGrabbed returns the Grabbed field if non-nil, zero value otherwise.

### GetGrabbedOk

`func (o *EpisodeResource) GetGrabbedOk() (*bool, bool)`

GetGrabbedOk returns a tuple with the Grabbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabbed

`func (o *EpisodeResource) SetGrabbed(v bool)`

SetGrabbed sets Grabbed field to given value.

### HasGrabbed

`func (o *EpisodeResource) HasGrabbed() bool`

HasGrabbed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


