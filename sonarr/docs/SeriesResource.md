# SeriesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**AlternateTitles** | Pointer to [**[]AlternateTitleResource**](AlternateTitleResource.md) |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**SeriesStatusType**](SeriesStatusType.md) |  | [optional] 
**Ended** | Pointer to **bool** |  | [optional] [readonly] 
**ProfileName** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**NextAiring** | Pointer to **NullableTime** |  | [optional] 
**PreviousAiring** | Pointer to **NullableTime** |  | [optional] 
**Network** | Pointer to **NullableString** |  | [optional] 
**AirTime** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**OriginalLanguage** | Pointer to [**Language**](Language.md) |  | [optional] 
**RemotePoster** | Pointer to **NullableString** |  | [optional] 
**Seasons** | Pointer to [**[]SeasonResource**](SeasonResource.md) |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**SeasonFolder** | Pointer to **bool** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**UseSceneNumbering** | Pointer to **bool** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**TvdbId** | Pointer to **int32** |  | [optional] 
**TvRageId** | Pointer to **int32** |  | [optional] 
**TvMazeId** | Pointer to **int32** |  | [optional] 
**FirstAired** | Pointer to **NullableTime** |  | [optional] 
**SeriesType** | Pointer to [**SeriesTypes**](SeriesTypes.md) |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Folder** | Pointer to **NullableString** |  | [optional] 
**Certification** | Pointer to **NullableString** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddSeriesOptions**](AddSeriesOptions.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Statistics** | Pointer to [**SeriesStatisticsResource**](SeriesStatisticsResource.md) |  | [optional] 
**EpisodesChanged** | Pointer to **NullableBool** |  | [optional] 
**LanguageProfileId** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewSeriesResource

`func NewSeriesResource() *SeriesResource`

NewSeriesResource instantiates a new SeriesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesResourceWithDefaults

`func NewSeriesResourceWithDefaults() *SeriesResource`

NewSeriesResourceWithDefaults instantiates a new SeriesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeriesResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeriesResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeriesResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SeriesResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *SeriesResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SeriesResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SeriesResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAlternateTitles

`func (o *SeriesResource) GetAlternateTitles() []AlternateTitleResource`

GetAlternateTitles returns the AlternateTitles field if non-nil, zero value otherwise.

### GetAlternateTitlesOk

`func (o *SeriesResource) GetAlternateTitlesOk() (*[]AlternateTitleResource, bool)`

GetAlternateTitlesOk returns a tuple with the AlternateTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateTitles

`func (o *SeriesResource) SetAlternateTitles(v []AlternateTitleResource)`

SetAlternateTitles sets AlternateTitles field to given value.

### HasAlternateTitles

`func (o *SeriesResource) HasAlternateTitles() bool`

HasAlternateTitles returns a boolean if a field has been set.

### SetAlternateTitlesNil

`func (o *SeriesResource) SetAlternateTitlesNil(b bool)`

 SetAlternateTitlesNil sets the value for AlternateTitles to be an explicit nil

### UnsetAlternateTitles
`func (o *SeriesResource) UnsetAlternateTitles()`

UnsetAlternateTitles ensures that no value is present for AlternateTitles, not even an explicit nil
### GetSortTitle

`func (o *SeriesResource) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *SeriesResource) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *SeriesResource) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *SeriesResource) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *SeriesResource) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *SeriesResource) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetStatus

`func (o *SeriesResource) GetStatus() SeriesStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SeriesResource) GetStatusOk() (*SeriesStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SeriesResource) SetStatus(v SeriesStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SeriesResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnded

`func (o *SeriesResource) GetEnded() bool`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *SeriesResource) GetEndedOk() (*bool, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *SeriesResource) SetEnded(v bool)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *SeriesResource) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetProfileName

`func (o *SeriesResource) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *SeriesResource) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *SeriesResource) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *SeriesResource) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### SetProfileNameNil

`func (o *SeriesResource) SetProfileNameNil(b bool)`

 SetProfileNameNil sets the value for ProfileName to be an explicit nil

### UnsetProfileName
`func (o *SeriesResource) UnsetProfileName()`

UnsetProfileName ensures that no value is present for ProfileName, not even an explicit nil
### GetOverview

`func (o *SeriesResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *SeriesResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *SeriesResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *SeriesResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *SeriesResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *SeriesResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetNextAiring

`func (o *SeriesResource) GetNextAiring() time.Time`

GetNextAiring returns the NextAiring field if non-nil, zero value otherwise.

### GetNextAiringOk

`func (o *SeriesResource) GetNextAiringOk() (*time.Time, bool)`

GetNextAiringOk returns a tuple with the NextAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAiring

`func (o *SeriesResource) SetNextAiring(v time.Time)`

SetNextAiring sets NextAiring field to given value.

### HasNextAiring

`func (o *SeriesResource) HasNextAiring() bool`

HasNextAiring returns a boolean if a field has been set.

### SetNextAiringNil

`func (o *SeriesResource) SetNextAiringNil(b bool)`

 SetNextAiringNil sets the value for NextAiring to be an explicit nil

### UnsetNextAiring
`func (o *SeriesResource) UnsetNextAiring()`

UnsetNextAiring ensures that no value is present for NextAiring, not even an explicit nil
### GetPreviousAiring

`func (o *SeriesResource) GetPreviousAiring() time.Time`

GetPreviousAiring returns the PreviousAiring field if non-nil, zero value otherwise.

### GetPreviousAiringOk

`func (o *SeriesResource) GetPreviousAiringOk() (*time.Time, bool)`

GetPreviousAiringOk returns a tuple with the PreviousAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAiring

`func (o *SeriesResource) SetPreviousAiring(v time.Time)`

SetPreviousAiring sets PreviousAiring field to given value.

### HasPreviousAiring

`func (o *SeriesResource) HasPreviousAiring() bool`

HasPreviousAiring returns a boolean if a field has been set.

### SetPreviousAiringNil

`func (o *SeriesResource) SetPreviousAiringNil(b bool)`

 SetPreviousAiringNil sets the value for PreviousAiring to be an explicit nil

### UnsetPreviousAiring
`func (o *SeriesResource) UnsetPreviousAiring()`

UnsetPreviousAiring ensures that no value is present for PreviousAiring, not even an explicit nil
### GetNetwork

`func (o *SeriesResource) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SeriesResource) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SeriesResource) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SeriesResource) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *SeriesResource) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *SeriesResource) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil
### GetAirTime

`func (o *SeriesResource) GetAirTime() string`

GetAirTime returns the AirTime field if non-nil, zero value otherwise.

### GetAirTimeOk

`func (o *SeriesResource) GetAirTimeOk() (*string, bool)`

GetAirTimeOk returns a tuple with the AirTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirTime

`func (o *SeriesResource) SetAirTime(v string)`

SetAirTime sets AirTime field to given value.

### HasAirTime

`func (o *SeriesResource) HasAirTime() bool`

HasAirTime returns a boolean if a field has been set.

### SetAirTimeNil

`func (o *SeriesResource) SetAirTimeNil(b bool)`

 SetAirTimeNil sets the value for AirTime to be an explicit nil

### UnsetAirTime
`func (o *SeriesResource) UnsetAirTime()`

UnsetAirTime ensures that no value is present for AirTime, not even an explicit nil
### GetImages

`func (o *SeriesResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SeriesResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SeriesResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *SeriesResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *SeriesResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *SeriesResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetOriginalLanguage

`func (o *SeriesResource) GetOriginalLanguage() Language`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *SeriesResource) GetOriginalLanguageOk() (*Language, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *SeriesResource) SetOriginalLanguage(v Language)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *SeriesResource) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetRemotePoster

`func (o *SeriesResource) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *SeriesResource) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *SeriesResource) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *SeriesResource) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### SetRemotePosterNil

`func (o *SeriesResource) SetRemotePosterNil(b bool)`

 SetRemotePosterNil sets the value for RemotePoster to be an explicit nil

### UnsetRemotePoster
`func (o *SeriesResource) UnsetRemotePoster()`

UnsetRemotePoster ensures that no value is present for RemotePoster, not even an explicit nil
### GetSeasons

`func (o *SeriesResource) GetSeasons() []SeasonResource`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *SeriesResource) GetSeasonsOk() (*[]SeasonResource, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *SeriesResource) SetSeasons(v []SeasonResource)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *SeriesResource) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### SetSeasonsNil

`func (o *SeriesResource) SetSeasonsNil(b bool)`

 SetSeasonsNil sets the value for Seasons to be an explicit nil

### UnsetSeasons
`func (o *SeriesResource) UnsetSeasons()`

UnsetSeasons ensures that no value is present for Seasons, not even an explicit nil
### GetYear

`func (o *SeriesResource) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeriesResource) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeriesResource) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeriesResource) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetPath

`func (o *SeriesResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SeriesResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SeriesResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SeriesResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *SeriesResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *SeriesResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQualityProfileId

`func (o *SeriesResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *SeriesResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *SeriesResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *SeriesResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetSeasonFolder

`func (o *SeriesResource) GetSeasonFolder() bool`

GetSeasonFolder returns the SeasonFolder field if non-nil, zero value otherwise.

### GetSeasonFolderOk

`func (o *SeriesResource) GetSeasonFolderOk() (*bool, bool)`

GetSeasonFolderOk returns a tuple with the SeasonFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonFolder

`func (o *SeriesResource) SetSeasonFolder(v bool)`

SetSeasonFolder sets SeasonFolder field to given value.

### HasSeasonFolder

`func (o *SeriesResource) HasSeasonFolder() bool`

HasSeasonFolder returns a boolean if a field has been set.

### GetMonitored

`func (o *SeriesResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *SeriesResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *SeriesResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *SeriesResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetUseSceneNumbering

`func (o *SeriesResource) GetUseSceneNumbering() bool`

GetUseSceneNumbering returns the UseSceneNumbering field if non-nil, zero value otherwise.

### GetUseSceneNumberingOk

`func (o *SeriesResource) GetUseSceneNumberingOk() (*bool, bool)`

GetUseSceneNumberingOk returns a tuple with the UseSceneNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSceneNumbering

`func (o *SeriesResource) SetUseSceneNumbering(v bool)`

SetUseSceneNumbering sets UseSceneNumbering field to given value.

### HasUseSceneNumbering

`func (o *SeriesResource) HasUseSceneNumbering() bool`

HasUseSceneNumbering returns a boolean if a field has been set.

### GetRuntime

`func (o *SeriesResource) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SeriesResource) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SeriesResource) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SeriesResource) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTvdbId

`func (o *SeriesResource) GetTvdbId() int32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *SeriesResource) GetTvdbIdOk() (*int32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *SeriesResource) SetTvdbId(v int32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *SeriesResource) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### GetTvRageId

`func (o *SeriesResource) GetTvRageId() int32`

GetTvRageId returns the TvRageId field if non-nil, zero value otherwise.

### GetTvRageIdOk

`func (o *SeriesResource) GetTvRageIdOk() (*int32, bool)`

GetTvRageIdOk returns a tuple with the TvRageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvRageId

`func (o *SeriesResource) SetTvRageId(v int32)`

SetTvRageId sets TvRageId field to given value.

### HasTvRageId

`func (o *SeriesResource) HasTvRageId() bool`

HasTvRageId returns a boolean if a field has been set.

### GetTvMazeId

`func (o *SeriesResource) GetTvMazeId() int32`

GetTvMazeId returns the TvMazeId field if non-nil, zero value otherwise.

### GetTvMazeIdOk

`func (o *SeriesResource) GetTvMazeIdOk() (*int32, bool)`

GetTvMazeIdOk returns a tuple with the TvMazeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvMazeId

`func (o *SeriesResource) SetTvMazeId(v int32)`

SetTvMazeId sets TvMazeId field to given value.

### HasTvMazeId

`func (o *SeriesResource) HasTvMazeId() bool`

HasTvMazeId returns a boolean if a field has been set.

### GetFirstAired

`func (o *SeriesResource) GetFirstAired() time.Time`

GetFirstAired returns the FirstAired field if non-nil, zero value otherwise.

### GetFirstAiredOk

`func (o *SeriesResource) GetFirstAiredOk() (*time.Time, bool)`

GetFirstAiredOk returns a tuple with the FirstAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAired

`func (o *SeriesResource) SetFirstAired(v time.Time)`

SetFirstAired sets FirstAired field to given value.

### HasFirstAired

`func (o *SeriesResource) HasFirstAired() bool`

HasFirstAired returns a boolean if a field has been set.

### SetFirstAiredNil

`func (o *SeriesResource) SetFirstAiredNil(b bool)`

 SetFirstAiredNil sets the value for FirstAired to be an explicit nil

### UnsetFirstAired
`func (o *SeriesResource) UnsetFirstAired()`

UnsetFirstAired ensures that no value is present for FirstAired, not even an explicit nil
### GetSeriesType

`func (o *SeriesResource) GetSeriesType() SeriesTypes`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *SeriesResource) GetSeriesTypeOk() (*SeriesTypes, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *SeriesResource) SetSeriesType(v SeriesTypes)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *SeriesResource) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetCleanTitle

`func (o *SeriesResource) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *SeriesResource) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *SeriesResource) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *SeriesResource) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *SeriesResource) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *SeriesResource) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetImdbId

`func (o *SeriesResource) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *SeriesResource) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *SeriesResource) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *SeriesResource) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *SeriesResource) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *SeriesResource) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTitleSlug

`func (o *SeriesResource) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *SeriesResource) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *SeriesResource) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *SeriesResource) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *SeriesResource) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *SeriesResource) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetRootFolderPath

`func (o *SeriesResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *SeriesResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *SeriesResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *SeriesResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *SeriesResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *SeriesResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetFolder

`func (o *SeriesResource) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *SeriesResource) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *SeriesResource) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *SeriesResource) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *SeriesResource) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *SeriesResource) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetCertification

`func (o *SeriesResource) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *SeriesResource) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *SeriesResource) SetCertification(v string)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *SeriesResource) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### SetCertificationNil

`func (o *SeriesResource) SetCertificationNil(b bool)`

 SetCertificationNil sets the value for Certification to be an explicit nil

### UnsetCertification
`func (o *SeriesResource) UnsetCertification()`

UnsetCertification ensures that no value is present for Certification, not even an explicit nil
### GetGenres

`func (o *SeriesResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SeriesResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SeriesResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SeriesResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *SeriesResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *SeriesResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetTags

`func (o *SeriesResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SeriesResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SeriesResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SeriesResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SeriesResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SeriesResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAdded

`func (o *SeriesResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *SeriesResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *SeriesResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *SeriesResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *SeriesResource) GetAddOptions() AddSeriesOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *SeriesResource) GetAddOptionsOk() (*AddSeriesOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *SeriesResource) SetAddOptions(v AddSeriesOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *SeriesResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRatings

`func (o *SeriesResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *SeriesResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *SeriesResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *SeriesResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetStatistics

`func (o *SeriesResource) GetStatistics() SeriesStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *SeriesResource) GetStatisticsOk() (*SeriesStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *SeriesResource) SetStatistics(v SeriesStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *SeriesResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetEpisodesChanged

`func (o *SeriesResource) GetEpisodesChanged() bool`

GetEpisodesChanged returns the EpisodesChanged field if non-nil, zero value otherwise.

### GetEpisodesChangedOk

`func (o *SeriesResource) GetEpisodesChangedOk() (*bool, bool)`

GetEpisodesChangedOk returns a tuple with the EpisodesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodesChanged

`func (o *SeriesResource) SetEpisodesChanged(v bool)`

SetEpisodesChanged sets EpisodesChanged field to given value.

### HasEpisodesChanged

`func (o *SeriesResource) HasEpisodesChanged() bool`

HasEpisodesChanged returns a boolean if a field has been set.

### SetEpisodesChangedNil

`func (o *SeriesResource) SetEpisodesChangedNil(b bool)`

 SetEpisodesChangedNil sets the value for EpisodesChanged to be an explicit nil

### UnsetEpisodesChanged
`func (o *SeriesResource) UnsetEpisodesChanged()`

UnsetEpisodesChanged ensures that no value is present for EpisodesChanged, not even an explicit nil
### GetLanguageProfileId

`func (o *SeriesResource) GetLanguageProfileId() int32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *SeriesResource) GetLanguageProfileIdOk() (*int32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *SeriesResource) SetLanguageProfileId(v int32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *SeriesResource) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


