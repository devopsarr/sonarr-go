# SeasonStatisticsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextAiring** | Pointer to **NullableTime** |  | [optional] 
**PreviousAiring** | Pointer to **NullableTime** |  | [optional] 
**EpisodeFileCount** | Pointer to **int32** |  | [optional] 
**EpisodeCount** | Pointer to **int32** |  | [optional] 
**TotalEpisodeCount** | Pointer to **int32** |  | [optional] 
**SizeOnDisk** | Pointer to **int64** |  | [optional] 
**ReleaseGroups** | Pointer to **[]string** |  | [optional] 
**PercentOfEpisodes** | Pointer to **float64** |  | [optional] [readonly] 

## Methods

### NewSeasonStatisticsResource

`func NewSeasonStatisticsResource() *SeasonStatisticsResource`

NewSeasonStatisticsResource instantiates a new SeasonStatisticsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonStatisticsResourceWithDefaults

`func NewSeasonStatisticsResourceWithDefaults() *SeasonStatisticsResource`

NewSeasonStatisticsResourceWithDefaults instantiates a new SeasonStatisticsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextAiring

`func (o *SeasonStatisticsResource) GetNextAiring() time.Time`

GetNextAiring returns the NextAiring field if non-nil, zero value otherwise.

### GetNextAiringOk

`func (o *SeasonStatisticsResource) GetNextAiringOk() (*time.Time, bool)`

GetNextAiringOk returns a tuple with the NextAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAiring

`func (o *SeasonStatisticsResource) SetNextAiring(v time.Time)`

SetNextAiring sets NextAiring field to given value.

### HasNextAiring

`func (o *SeasonStatisticsResource) HasNextAiring() bool`

HasNextAiring returns a boolean if a field has been set.

### SetNextAiringNil

`func (o *SeasonStatisticsResource) SetNextAiringNil(b bool)`

 SetNextAiringNil sets the value for NextAiring to be an explicit nil

### UnsetNextAiring
`func (o *SeasonStatisticsResource) UnsetNextAiring()`

UnsetNextAiring ensures that no value is present for NextAiring, not even an explicit nil
### GetPreviousAiring

`func (o *SeasonStatisticsResource) GetPreviousAiring() time.Time`

GetPreviousAiring returns the PreviousAiring field if non-nil, zero value otherwise.

### GetPreviousAiringOk

`func (o *SeasonStatisticsResource) GetPreviousAiringOk() (*time.Time, bool)`

GetPreviousAiringOk returns a tuple with the PreviousAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAiring

`func (o *SeasonStatisticsResource) SetPreviousAiring(v time.Time)`

SetPreviousAiring sets PreviousAiring field to given value.

### HasPreviousAiring

`func (o *SeasonStatisticsResource) HasPreviousAiring() bool`

HasPreviousAiring returns a boolean if a field has been set.

### SetPreviousAiringNil

`func (o *SeasonStatisticsResource) SetPreviousAiringNil(b bool)`

 SetPreviousAiringNil sets the value for PreviousAiring to be an explicit nil

### UnsetPreviousAiring
`func (o *SeasonStatisticsResource) UnsetPreviousAiring()`

UnsetPreviousAiring ensures that no value is present for PreviousAiring, not even an explicit nil
### GetEpisodeFileCount

`func (o *SeasonStatisticsResource) GetEpisodeFileCount() int32`

GetEpisodeFileCount returns the EpisodeFileCount field if non-nil, zero value otherwise.

### GetEpisodeFileCountOk

`func (o *SeasonStatisticsResource) GetEpisodeFileCountOk() (*int32, bool)`

GetEpisodeFileCountOk returns a tuple with the EpisodeFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFileCount

`func (o *SeasonStatisticsResource) SetEpisodeFileCount(v int32)`

SetEpisodeFileCount sets EpisodeFileCount field to given value.

### HasEpisodeFileCount

`func (o *SeasonStatisticsResource) HasEpisodeFileCount() bool`

HasEpisodeFileCount returns a boolean if a field has been set.

### GetEpisodeCount

`func (o *SeasonStatisticsResource) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *SeasonStatisticsResource) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *SeasonStatisticsResource) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *SeasonStatisticsResource) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetTotalEpisodeCount

`func (o *SeasonStatisticsResource) GetTotalEpisodeCount() int32`

GetTotalEpisodeCount returns the TotalEpisodeCount field if non-nil, zero value otherwise.

### GetTotalEpisodeCountOk

`func (o *SeasonStatisticsResource) GetTotalEpisodeCountOk() (*int32, bool)`

GetTotalEpisodeCountOk returns a tuple with the TotalEpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEpisodeCount

`func (o *SeasonStatisticsResource) SetTotalEpisodeCount(v int32)`

SetTotalEpisodeCount sets TotalEpisodeCount field to given value.

### HasTotalEpisodeCount

`func (o *SeasonStatisticsResource) HasTotalEpisodeCount() bool`

HasTotalEpisodeCount returns a boolean if a field has been set.

### GetSizeOnDisk

`func (o *SeasonStatisticsResource) GetSizeOnDisk() int64`

GetSizeOnDisk returns the SizeOnDisk field if non-nil, zero value otherwise.

### GetSizeOnDiskOk

`func (o *SeasonStatisticsResource) GetSizeOnDiskOk() (*int64, bool)`

GetSizeOnDiskOk returns a tuple with the SizeOnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOnDisk

`func (o *SeasonStatisticsResource) SetSizeOnDisk(v int64)`

SetSizeOnDisk sets SizeOnDisk field to given value.

### HasSizeOnDisk

`func (o *SeasonStatisticsResource) HasSizeOnDisk() bool`

HasSizeOnDisk returns a boolean if a field has been set.

### GetReleaseGroups

`func (o *SeasonStatisticsResource) GetReleaseGroups() []string`

GetReleaseGroups returns the ReleaseGroups field if non-nil, zero value otherwise.

### GetReleaseGroupsOk

`func (o *SeasonStatisticsResource) GetReleaseGroupsOk() (*[]string, bool)`

GetReleaseGroupsOk returns a tuple with the ReleaseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroups

`func (o *SeasonStatisticsResource) SetReleaseGroups(v []string)`

SetReleaseGroups sets ReleaseGroups field to given value.

### HasReleaseGroups

`func (o *SeasonStatisticsResource) HasReleaseGroups() bool`

HasReleaseGroups returns a boolean if a field has been set.

### SetReleaseGroupsNil

`func (o *SeasonStatisticsResource) SetReleaseGroupsNil(b bool)`

 SetReleaseGroupsNil sets the value for ReleaseGroups to be an explicit nil

### UnsetReleaseGroups
`func (o *SeasonStatisticsResource) UnsetReleaseGroups()`

UnsetReleaseGroups ensures that no value is present for ReleaseGroups, not even an explicit nil
### GetPercentOfEpisodes

`func (o *SeasonStatisticsResource) GetPercentOfEpisodes() float64`

GetPercentOfEpisodes returns the PercentOfEpisodes field if non-nil, zero value otherwise.

### GetPercentOfEpisodesOk

`func (o *SeasonStatisticsResource) GetPercentOfEpisodesOk() (*float64, bool)`

GetPercentOfEpisodesOk returns a tuple with the PercentOfEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfEpisodes

`func (o *SeasonStatisticsResource) SetPercentOfEpisodes(v float64)`

SetPercentOfEpisodes sets PercentOfEpisodes field to given value.

### HasPercentOfEpisodes

`func (o *SeasonStatisticsResource) HasPercentOfEpisodes() bool`

HasPercentOfEpisodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


