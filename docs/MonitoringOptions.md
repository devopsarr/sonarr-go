# MonitoringOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreEpisodesWithFiles** | Pointer to **bool** |  | [optional] 
**IgnoreEpisodesWithoutFiles** | Pointer to **bool** |  | [optional] 
**Monitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 

## Methods

### NewMonitoringOptions

`func NewMonitoringOptions() *MonitoringOptions`

NewMonitoringOptions instantiates a new MonitoringOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringOptionsWithDefaults

`func NewMonitoringOptionsWithDefaults() *MonitoringOptions`

NewMonitoringOptionsWithDefaults instantiates a new MonitoringOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreEpisodesWithFiles

`func (o *MonitoringOptions) GetIgnoreEpisodesWithFiles() bool`

GetIgnoreEpisodesWithFiles returns the IgnoreEpisodesWithFiles field if non-nil, zero value otherwise.

### GetIgnoreEpisodesWithFilesOk

`func (o *MonitoringOptions) GetIgnoreEpisodesWithFilesOk() (*bool, bool)`

GetIgnoreEpisodesWithFilesOk returns a tuple with the IgnoreEpisodesWithFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEpisodesWithFiles

`func (o *MonitoringOptions) SetIgnoreEpisodesWithFiles(v bool)`

SetIgnoreEpisodesWithFiles sets IgnoreEpisodesWithFiles field to given value.

### HasIgnoreEpisodesWithFiles

`func (o *MonitoringOptions) HasIgnoreEpisodesWithFiles() bool`

HasIgnoreEpisodesWithFiles returns a boolean if a field has been set.

### GetIgnoreEpisodesWithoutFiles

`func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFiles() bool`

GetIgnoreEpisodesWithoutFiles returns the IgnoreEpisodesWithoutFiles field if non-nil, zero value otherwise.

### GetIgnoreEpisodesWithoutFilesOk

`func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFilesOk() (*bool, bool)`

GetIgnoreEpisodesWithoutFilesOk returns a tuple with the IgnoreEpisodesWithoutFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEpisodesWithoutFiles

`func (o *MonitoringOptions) SetIgnoreEpisodesWithoutFiles(v bool)`

SetIgnoreEpisodesWithoutFiles sets IgnoreEpisodesWithoutFiles field to given value.

### HasIgnoreEpisodesWithoutFiles

`func (o *MonitoringOptions) HasIgnoreEpisodesWithoutFiles() bool`

HasIgnoreEpisodesWithoutFiles returns a boolean if a field has been set.

### GetMonitor

`func (o *MonitoringOptions) GetMonitor() MonitorTypes`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *MonitoringOptions) GetMonitorOk() (*MonitorTypes, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *MonitoringOptions) SetMonitor(v MonitorTypes)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *MonitoringOptions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


