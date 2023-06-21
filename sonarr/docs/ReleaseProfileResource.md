# ReleaseProfileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **[]string** |  | [optional] 
**Ignored** | Pointer to **[]string** |  | [optional] 
**IndexerId** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewReleaseProfileResource

`func NewReleaseProfileResource() *ReleaseProfileResource`

NewReleaseProfileResource instantiates a new ReleaseProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseProfileResourceWithDefaults

`func NewReleaseProfileResourceWithDefaults() *ReleaseProfileResource`

NewReleaseProfileResourceWithDefaults instantiates a new ReleaseProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReleaseProfileResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReleaseProfileResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReleaseProfileResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReleaseProfileResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReleaseProfileResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReleaseProfileResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *ReleaseProfileResource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReleaseProfileResource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReleaseProfileResource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ReleaseProfileResource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequired

`func (o *ReleaseProfileResource) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ReleaseProfileResource) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ReleaseProfileResource) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ReleaseProfileResource) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ReleaseProfileResource) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ReleaseProfileResource) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetIgnored

`func (o *ReleaseProfileResource) GetIgnored() []string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ReleaseProfileResource) GetIgnoredOk() (*[]string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ReleaseProfileResource) SetIgnored(v []string)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ReleaseProfileResource) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### SetIgnoredNil

`func (o *ReleaseProfileResource) SetIgnoredNil(b bool)`

 SetIgnoredNil sets the value for Ignored to be an explicit nil

### UnsetIgnored
`func (o *ReleaseProfileResource) UnsetIgnored()`

UnsetIgnored ensures that no value is present for Ignored, not even an explicit nil
### GetIndexerId

`func (o *ReleaseProfileResource) GetIndexerId() int32`

GetIndexerId returns the IndexerId field if non-nil, zero value otherwise.

### GetIndexerIdOk

`func (o *ReleaseProfileResource) GetIndexerIdOk() (*int32, bool)`

GetIndexerIdOk returns a tuple with the IndexerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerId

`func (o *ReleaseProfileResource) SetIndexerId(v int32)`

SetIndexerId sets IndexerId field to given value.

### HasIndexerId

`func (o *ReleaseProfileResource) HasIndexerId() bool`

HasIndexerId returns a boolean if a field has been set.

### GetTags

`func (o *ReleaseProfileResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReleaseProfileResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReleaseProfileResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReleaseProfileResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ReleaseProfileResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ReleaseProfileResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


