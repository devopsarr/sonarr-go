# SeasonPassResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]SeasonPassSeriesResource**](SeasonPassSeriesResource.md) |  | [optional] 
**MonitoringOptions** | Pointer to [**MonitoringOptions**](MonitoringOptions.md) |  | [optional] 

## Methods

### NewSeasonPassResource

`func NewSeasonPassResource() *SeasonPassResource`

NewSeasonPassResource instantiates a new SeasonPassResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonPassResourceWithDefaults

`func NewSeasonPassResourceWithDefaults() *SeasonPassResource`

NewSeasonPassResourceWithDefaults instantiates a new SeasonPassResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *SeasonPassResource) GetSeries() []SeasonPassSeriesResource`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SeasonPassResource) GetSeriesOk() (*[]SeasonPassSeriesResource, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SeasonPassResource) SetSeries(v []SeasonPassSeriesResource)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *SeasonPassResource) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *SeasonPassResource) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *SeasonPassResource) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetMonitoringOptions

`func (o *SeasonPassResource) GetMonitoringOptions() MonitoringOptions`

GetMonitoringOptions returns the MonitoringOptions field if non-nil, zero value otherwise.

### GetMonitoringOptionsOk

`func (o *SeasonPassResource) GetMonitoringOptionsOk() (*MonitoringOptions, bool)`

GetMonitoringOptionsOk returns a tuple with the MonitoringOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringOptions

`func (o *SeasonPassResource) SetMonitoringOptions(v MonitoringOptions)`

SetMonitoringOptions sets MonitoringOptions field to given value.

### HasMonitoringOptions

`func (o *SeasonPassResource) HasMonitoringOptions() bool`

HasMonitoringOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


