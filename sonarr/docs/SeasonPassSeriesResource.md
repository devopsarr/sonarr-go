# SeasonPassSeriesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**Seasons** | Pointer to [**[]SeasonResource**](SeasonResource.md) |  | [optional] 

## Methods

### NewSeasonPassSeriesResource

`func NewSeasonPassSeriesResource() *SeasonPassSeriesResource`

NewSeasonPassSeriesResource instantiates a new SeasonPassSeriesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonPassSeriesResourceWithDefaults

`func NewSeasonPassSeriesResourceWithDefaults() *SeasonPassSeriesResource`

NewSeasonPassSeriesResourceWithDefaults instantiates a new SeasonPassSeriesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeasonPassSeriesResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeasonPassSeriesResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeasonPassSeriesResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SeasonPassSeriesResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonitored

`func (o *SeasonPassSeriesResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *SeasonPassSeriesResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *SeasonPassSeriesResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *SeasonPassSeriesResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *SeasonPassSeriesResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *SeasonPassSeriesResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetSeasons

`func (o *SeasonPassSeriesResource) GetSeasons() []SeasonResource`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *SeasonPassSeriesResource) GetSeasonsOk() (*[]SeasonResource, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *SeasonPassSeriesResource) SetSeasons(v []SeasonResource)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *SeasonPassSeriesResource) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### SetSeasonsNil

`func (o *SeasonPassSeriesResource) SetSeasonsNil(b bool)`

 SetSeasonsNil sets the value for Seasons to be an explicit nil

### UnsetSeasons
`func (o *SeasonPassSeriesResource) UnsetSeasons()`

UnsetSeasons ensures that no value is present for Seasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


