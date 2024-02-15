# ImportListBulkResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**EnableAutomaticAdd** | Pointer to **NullableBool** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewImportListBulkResource

`func NewImportListBulkResource() *ImportListBulkResource`

NewImportListBulkResource instantiates a new ImportListBulkResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportListBulkResourceWithDefaults

`func NewImportListBulkResourceWithDefaults() *ImportListBulkResource`

NewImportListBulkResourceWithDefaults instantiates a new ImportListBulkResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ImportListBulkResource) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ImportListBulkResource) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ImportListBulkResource) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ImportListBulkResource) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *ImportListBulkResource) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *ImportListBulkResource) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetTags

`func (o *ImportListBulkResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportListBulkResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportListBulkResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportListBulkResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ImportListBulkResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ImportListBulkResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *ImportListBulkResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *ImportListBulkResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *ImportListBulkResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *ImportListBulkResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetEnableAutomaticAdd

`func (o *ImportListBulkResource) GetEnableAutomaticAdd() bool`

GetEnableAutomaticAdd returns the EnableAutomaticAdd field if non-nil, zero value otherwise.

### GetEnableAutomaticAddOk

`func (o *ImportListBulkResource) GetEnableAutomaticAddOk() (*bool, bool)`

GetEnableAutomaticAddOk returns a tuple with the EnableAutomaticAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticAdd

`func (o *ImportListBulkResource) SetEnableAutomaticAdd(v bool)`

SetEnableAutomaticAdd sets EnableAutomaticAdd field to given value.

### HasEnableAutomaticAdd

`func (o *ImportListBulkResource) HasEnableAutomaticAdd() bool`

HasEnableAutomaticAdd returns a boolean if a field has been set.

### SetEnableAutomaticAddNil

`func (o *ImportListBulkResource) SetEnableAutomaticAddNil(b bool)`

 SetEnableAutomaticAddNil sets the value for EnableAutomaticAdd to be an explicit nil

### UnsetEnableAutomaticAdd
`func (o *ImportListBulkResource) UnsetEnableAutomaticAdd()`

UnsetEnableAutomaticAdd ensures that no value is present for EnableAutomaticAdd, not even an explicit nil
### GetRootFolderPath

`func (o *ImportListBulkResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *ImportListBulkResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *ImportListBulkResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *ImportListBulkResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *ImportListBulkResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *ImportListBulkResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetQualityProfileId

`func (o *ImportListBulkResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *ImportListBulkResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *ImportListBulkResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *ImportListBulkResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *ImportListBulkResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *ImportListBulkResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


