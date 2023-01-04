# SeriesStatisticsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeasonCount** | Pointer to **int32** |  | [optional] 
**EpisodeFileCount** | Pointer to **int32** |  | [optional] 
**EpisodeCount** | Pointer to **int32** |  | [optional] 
**TotalEpisodeCount** | Pointer to **int32** |  | [optional] 
**SizeOnDisk** | Pointer to **int64** |  | [optional] 
**ReleaseGroups** | Pointer to **[]string** |  | [optional] 
**PercentOfEpisodes** | Pointer to **float64** |  | [optional] [readonly] 

## Methods

### NewSeriesStatisticsResource

`func NewSeriesStatisticsResource() *SeriesStatisticsResource`

NewSeriesStatisticsResource instantiates a new SeriesStatisticsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesStatisticsResourceWithDefaults

`func NewSeriesStatisticsResourceWithDefaults() *SeriesStatisticsResource`

NewSeriesStatisticsResourceWithDefaults instantiates a new SeriesStatisticsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeasonCount

`func (o *SeriesStatisticsResource) GetSeasonCount() int32`

GetSeasonCount returns the SeasonCount field if non-nil, zero value otherwise.

### GetSeasonCountOk

`func (o *SeriesStatisticsResource) GetSeasonCountOk() (*int32, bool)`

GetSeasonCountOk returns a tuple with the SeasonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonCount

`func (o *SeriesStatisticsResource) SetSeasonCount(v int32)`

SetSeasonCount sets SeasonCount field to given value.

### HasSeasonCount

`func (o *SeriesStatisticsResource) HasSeasonCount() bool`

HasSeasonCount returns a boolean if a field has been set.

### GetEpisodeFileCount

`func (o *SeriesStatisticsResource) GetEpisodeFileCount() int32`

GetEpisodeFileCount returns the EpisodeFileCount field if non-nil, zero value otherwise.

### GetEpisodeFileCountOk

`func (o *SeriesStatisticsResource) GetEpisodeFileCountOk() (*int32, bool)`

GetEpisodeFileCountOk returns a tuple with the EpisodeFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFileCount

`func (o *SeriesStatisticsResource) SetEpisodeFileCount(v int32)`

SetEpisodeFileCount sets EpisodeFileCount field to given value.

### HasEpisodeFileCount

`func (o *SeriesStatisticsResource) HasEpisodeFileCount() bool`

HasEpisodeFileCount returns a boolean if a field has been set.

### GetEpisodeCount

`func (o *SeriesStatisticsResource) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *SeriesStatisticsResource) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *SeriesStatisticsResource) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *SeriesStatisticsResource) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetTotalEpisodeCount

`func (o *SeriesStatisticsResource) GetTotalEpisodeCount() int32`

GetTotalEpisodeCount returns the TotalEpisodeCount field if non-nil, zero value otherwise.

### GetTotalEpisodeCountOk

`func (o *SeriesStatisticsResource) GetTotalEpisodeCountOk() (*int32, bool)`

GetTotalEpisodeCountOk returns a tuple with the TotalEpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEpisodeCount

`func (o *SeriesStatisticsResource) SetTotalEpisodeCount(v int32)`

SetTotalEpisodeCount sets TotalEpisodeCount field to given value.

### HasTotalEpisodeCount

`func (o *SeriesStatisticsResource) HasTotalEpisodeCount() bool`

HasTotalEpisodeCount returns a boolean if a field has been set.

### GetSizeOnDisk

`func (o *SeriesStatisticsResource) GetSizeOnDisk() int64`

GetSizeOnDisk returns the SizeOnDisk field if non-nil, zero value otherwise.

### GetSizeOnDiskOk

`func (o *SeriesStatisticsResource) GetSizeOnDiskOk() (*int64, bool)`

GetSizeOnDiskOk returns a tuple with the SizeOnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOnDisk

`func (o *SeriesStatisticsResource) SetSizeOnDisk(v int64)`

SetSizeOnDisk sets SizeOnDisk field to given value.

### HasSizeOnDisk

`func (o *SeriesStatisticsResource) HasSizeOnDisk() bool`

HasSizeOnDisk returns a boolean if a field has been set.

### GetReleaseGroups

`func (o *SeriesStatisticsResource) GetReleaseGroups() []string`

GetReleaseGroups returns the ReleaseGroups field if non-nil, zero value otherwise.

### GetReleaseGroupsOk

`func (o *SeriesStatisticsResource) GetReleaseGroupsOk() (*[]string, bool)`

GetReleaseGroupsOk returns a tuple with the ReleaseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroups

`func (o *SeriesStatisticsResource) SetReleaseGroups(v []string)`

SetReleaseGroups sets ReleaseGroups field to given value.

### HasReleaseGroups

`func (o *SeriesStatisticsResource) HasReleaseGroups() bool`

HasReleaseGroups returns a boolean if a field has been set.

### SetReleaseGroupsNil

`func (o *SeriesStatisticsResource) SetReleaseGroupsNil(b bool)`

 SetReleaseGroupsNil sets the value for ReleaseGroups to be an explicit nil

### UnsetReleaseGroups
`func (o *SeriesStatisticsResource) UnsetReleaseGroups()`

UnsetReleaseGroups ensures that no value is present for ReleaseGroups, not even an explicit nil
### GetPercentOfEpisodes

`func (o *SeriesStatisticsResource) GetPercentOfEpisodes() float64`

GetPercentOfEpisodes returns the PercentOfEpisodes field if non-nil, zero value otherwise.

### GetPercentOfEpisodesOk

`func (o *SeriesStatisticsResource) GetPercentOfEpisodesOk() (*float64, bool)`

GetPercentOfEpisodesOk returns a tuple with the PercentOfEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfEpisodes

`func (o *SeriesStatisticsResource) SetPercentOfEpisodes(v float64)`

SetPercentOfEpisodes sets PercentOfEpisodes field to given value.

### HasPercentOfEpisodes

`func (o *SeriesStatisticsResource) HasPercentOfEpisodes() bool`

HasPercentOfEpisodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


