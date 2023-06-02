# DownloadClientBulkResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**Enable** | Pointer to **NullableBool** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**RemoveCompletedDownloads** | Pointer to **NullableBool** |  | [optional] 
**RemoveFailedDownloads** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDownloadClientBulkResource

`func NewDownloadClientBulkResource() *DownloadClientBulkResource`

NewDownloadClientBulkResource instantiates a new DownloadClientBulkResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadClientBulkResourceWithDefaults

`func NewDownloadClientBulkResourceWithDefaults() *DownloadClientBulkResource`

NewDownloadClientBulkResourceWithDefaults instantiates a new DownloadClientBulkResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DownloadClientBulkResource) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DownloadClientBulkResource) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DownloadClientBulkResource) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DownloadClientBulkResource) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *DownloadClientBulkResource) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *DownloadClientBulkResource) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetTags

`func (o *DownloadClientBulkResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DownloadClientBulkResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DownloadClientBulkResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DownloadClientBulkResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DownloadClientBulkResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DownloadClientBulkResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *DownloadClientBulkResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *DownloadClientBulkResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *DownloadClientBulkResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *DownloadClientBulkResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetEnable

`func (o *DownloadClientBulkResource) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DownloadClientBulkResource) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DownloadClientBulkResource) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DownloadClientBulkResource) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *DownloadClientBulkResource) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *DownloadClientBulkResource) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetPriority

`func (o *DownloadClientBulkResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DownloadClientBulkResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DownloadClientBulkResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DownloadClientBulkResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *DownloadClientBulkResource) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *DownloadClientBulkResource) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetRemoveCompletedDownloads

`func (o *DownloadClientBulkResource) GetRemoveCompletedDownloads() bool`

GetRemoveCompletedDownloads returns the RemoveCompletedDownloads field if non-nil, zero value otherwise.

### GetRemoveCompletedDownloadsOk

`func (o *DownloadClientBulkResource) GetRemoveCompletedDownloadsOk() (*bool, bool)`

GetRemoveCompletedDownloadsOk returns a tuple with the RemoveCompletedDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveCompletedDownloads

`func (o *DownloadClientBulkResource) SetRemoveCompletedDownloads(v bool)`

SetRemoveCompletedDownloads sets RemoveCompletedDownloads field to given value.

### HasRemoveCompletedDownloads

`func (o *DownloadClientBulkResource) HasRemoveCompletedDownloads() bool`

HasRemoveCompletedDownloads returns a boolean if a field has been set.

### SetRemoveCompletedDownloadsNil

`func (o *DownloadClientBulkResource) SetRemoveCompletedDownloadsNil(b bool)`

 SetRemoveCompletedDownloadsNil sets the value for RemoveCompletedDownloads to be an explicit nil

### UnsetRemoveCompletedDownloads
`func (o *DownloadClientBulkResource) UnsetRemoveCompletedDownloads()`

UnsetRemoveCompletedDownloads ensures that no value is present for RemoveCompletedDownloads, not even an explicit nil
### GetRemoveFailedDownloads

`func (o *DownloadClientBulkResource) GetRemoveFailedDownloads() bool`

GetRemoveFailedDownloads returns the RemoveFailedDownloads field if non-nil, zero value otherwise.

### GetRemoveFailedDownloadsOk

`func (o *DownloadClientBulkResource) GetRemoveFailedDownloadsOk() (*bool, bool)`

GetRemoveFailedDownloadsOk returns a tuple with the RemoveFailedDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFailedDownloads

`func (o *DownloadClientBulkResource) SetRemoveFailedDownloads(v bool)`

SetRemoveFailedDownloads sets RemoveFailedDownloads field to given value.

### HasRemoveFailedDownloads

`func (o *DownloadClientBulkResource) HasRemoveFailedDownloads() bool`

HasRemoveFailedDownloads returns a boolean if a field has been set.

### SetRemoveFailedDownloadsNil

`func (o *DownloadClientBulkResource) SetRemoveFailedDownloadsNil(b bool)`

 SetRemoveFailedDownloadsNil sets the value for RemoveFailedDownloads to be an explicit nil

### UnsetRemoveFailedDownloads
`func (o *DownloadClientBulkResource) UnsetRemoveFailedDownloads()`

UnsetRemoveFailedDownloads ensures that no value is present for RemoveFailedDownloads, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


