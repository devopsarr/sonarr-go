# AutoTaggingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RemoveTagsAutomatically** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Specifications** | Pointer to [**[]AutoTaggingSpecificationSchema**](AutoTaggingSpecificationSchema.md) |  | [optional] 

## Methods

### NewAutoTaggingResource

`func NewAutoTaggingResource() *AutoTaggingResource`

NewAutoTaggingResource instantiates a new AutoTaggingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTaggingResourceWithDefaults

`func NewAutoTaggingResourceWithDefaults() *AutoTaggingResource`

NewAutoTaggingResourceWithDefaults instantiates a new AutoTaggingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTaggingResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTaggingResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTaggingResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTaggingResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutoTaggingResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTaggingResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTaggingResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTaggingResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTaggingResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTaggingResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoveTagsAutomatically

`func (o *AutoTaggingResource) GetRemoveTagsAutomatically() bool`

GetRemoveTagsAutomatically returns the RemoveTagsAutomatically field if non-nil, zero value otherwise.

### GetRemoveTagsAutomaticallyOk

`func (o *AutoTaggingResource) GetRemoveTagsAutomaticallyOk() (*bool, bool)`

GetRemoveTagsAutomaticallyOk returns a tuple with the RemoveTagsAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTagsAutomatically

`func (o *AutoTaggingResource) SetRemoveTagsAutomatically(v bool)`

SetRemoveTagsAutomatically sets RemoveTagsAutomatically field to given value.

### HasRemoveTagsAutomatically

`func (o *AutoTaggingResource) HasRemoveTagsAutomatically() bool`

HasRemoveTagsAutomatically returns a boolean if a field has been set.

### GetTags

`func (o *AutoTaggingResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutoTaggingResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutoTaggingResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutoTaggingResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AutoTaggingResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AutoTaggingResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSpecifications

`func (o *AutoTaggingResource) GetSpecifications() []AutoTaggingSpecificationSchema`

GetSpecifications returns the Specifications field if non-nil, zero value otherwise.

### GetSpecificationsOk

`func (o *AutoTaggingResource) GetSpecificationsOk() (*[]AutoTaggingSpecificationSchema, bool)`

GetSpecificationsOk returns a tuple with the Specifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifications

`func (o *AutoTaggingResource) SetSpecifications(v []AutoTaggingSpecificationSchema)`

SetSpecifications sets Specifications field to given value.

### HasSpecifications

`func (o *AutoTaggingResource) HasSpecifications() bool`

HasSpecifications returns a boolean if a field has been set.

### SetSpecificationsNil

`func (o *AutoTaggingResource) SetSpecificationsNil(b bool)`

 SetSpecificationsNil sets the value for Specifications to be an explicit nil

### UnsetSpecifications
`func (o *AutoTaggingResource) UnsetSpecifications()`

UnsetSpecifications ensures that no value is present for Specifications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


