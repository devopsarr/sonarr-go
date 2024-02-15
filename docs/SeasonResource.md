# SeasonResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeasonNumber** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**Statistics** | Pointer to [**SeasonStatisticsResource**](SeasonStatisticsResource.md) |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 

## Methods

### NewSeasonResource

`func NewSeasonResource() *SeasonResource`

NewSeasonResource instantiates a new SeasonResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonResourceWithDefaults

`func NewSeasonResourceWithDefaults() *SeasonResource`

NewSeasonResourceWithDefaults instantiates a new SeasonResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeasonNumber

`func (o *SeasonResource) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *SeasonResource) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *SeasonResource) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *SeasonResource) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetMonitored

`func (o *SeasonResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *SeasonResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *SeasonResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *SeasonResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetStatistics

`func (o *SeasonResource) GetStatistics() SeasonStatisticsResource`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *SeasonResource) GetStatisticsOk() (*SeasonStatisticsResource, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *SeasonResource) SetStatistics(v SeasonStatisticsResource)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *SeasonResource) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetImages

`func (o *SeasonResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SeasonResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SeasonResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *SeasonResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *SeasonResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *SeasonResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


