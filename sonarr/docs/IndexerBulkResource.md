# IndexerBulkResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**EnableRss** | Pointer to **NullableBool** |  | [optional] 
**EnableAutomaticSearch** | Pointer to **NullableBool** |  | [optional] 
**EnableInteractiveSearch** | Pointer to **NullableBool** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIndexerBulkResource

`func NewIndexerBulkResource() *IndexerBulkResource`

NewIndexerBulkResource instantiates a new IndexerBulkResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerBulkResourceWithDefaults

`func NewIndexerBulkResourceWithDefaults() *IndexerBulkResource`

NewIndexerBulkResourceWithDefaults instantiates a new IndexerBulkResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *IndexerBulkResource) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *IndexerBulkResource) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *IndexerBulkResource) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *IndexerBulkResource) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *IndexerBulkResource) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *IndexerBulkResource) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetTags

`func (o *IndexerBulkResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndexerBulkResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndexerBulkResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IndexerBulkResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IndexerBulkResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IndexerBulkResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *IndexerBulkResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *IndexerBulkResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *IndexerBulkResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *IndexerBulkResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetEnableRss

`func (o *IndexerBulkResource) GetEnableRss() bool`

GetEnableRss returns the EnableRss field if non-nil, zero value otherwise.

### GetEnableRssOk

`func (o *IndexerBulkResource) GetEnableRssOk() (*bool, bool)`

GetEnableRssOk returns a tuple with the EnableRss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRss

`func (o *IndexerBulkResource) SetEnableRss(v bool)`

SetEnableRss sets EnableRss field to given value.

### HasEnableRss

`func (o *IndexerBulkResource) HasEnableRss() bool`

HasEnableRss returns a boolean if a field has been set.

### SetEnableRssNil

`func (o *IndexerBulkResource) SetEnableRssNil(b bool)`

 SetEnableRssNil sets the value for EnableRss to be an explicit nil

### UnsetEnableRss
`func (o *IndexerBulkResource) UnsetEnableRss()`

UnsetEnableRss ensures that no value is present for EnableRss, not even an explicit nil
### GetEnableAutomaticSearch

`func (o *IndexerBulkResource) GetEnableAutomaticSearch() bool`

GetEnableAutomaticSearch returns the EnableAutomaticSearch field if non-nil, zero value otherwise.

### GetEnableAutomaticSearchOk

`func (o *IndexerBulkResource) GetEnableAutomaticSearchOk() (*bool, bool)`

GetEnableAutomaticSearchOk returns a tuple with the EnableAutomaticSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSearch

`func (o *IndexerBulkResource) SetEnableAutomaticSearch(v bool)`

SetEnableAutomaticSearch sets EnableAutomaticSearch field to given value.

### HasEnableAutomaticSearch

`func (o *IndexerBulkResource) HasEnableAutomaticSearch() bool`

HasEnableAutomaticSearch returns a boolean if a field has been set.

### SetEnableAutomaticSearchNil

`func (o *IndexerBulkResource) SetEnableAutomaticSearchNil(b bool)`

 SetEnableAutomaticSearchNil sets the value for EnableAutomaticSearch to be an explicit nil

### UnsetEnableAutomaticSearch
`func (o *IndexerBulkResource) UnsetEnableAutomaticSearch()`

UnsetEnableAutomaticSearch ensures that no value is present for EnableAutomaticSearch, not even an explicit nil
### GetEnableInteractiveSearch

`func (o *IndexerBulkResource) GetEnableInteractiveSearch() bool`

GetEnableInteractiveSearch returns the EnableInteractiveSearch field if non-nil, zero value otherwise.

### GetEnableInteractiveSearchOk

`func (o *IndexerBulkResource) GetEnableInteractiveSearchOk() (*bool, bool)`

GetEnableInteractiveSearchOk returns a tuple with the EnableInteractiveSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInteractiveSearch

`func (o *IndexerBulkResource) SetEnableInteractiveSearch(v bool)`

SetEnableInteractiveSearch sets EnableInteractiveSearch field to given value.

### HasEnableInteractiveSearch

`func (o *IndexerBulkResource) HasEnableInteractiveSearch() bool`

HasEnableInteractiveSearch returns a boolean if a field has been set.

### SetEnableInteractiveSearchNil

`func (o *IndexerBulkResource) SetEnableInteractiveSearchNil(b bool)`

 SetEnableInteractiveSearchNil sets the value for EnableInteractiveSearch to be an explicit nil

### UnsetEnableInteractiveSearch
`func (o *IndexerBulkResource) UnsetEnableInteractiveSearch()`

UnsetEnableInteractiveSearch ensures that no value is present for EnableInteractiveSearch, not even an explicit nil
### GetPriority

`func (o *IndexerBulkResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IndexerBulkResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IndexerBulkResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IndexerBulkResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *IndexerBulkResource) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *IndexerBulkResource) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


