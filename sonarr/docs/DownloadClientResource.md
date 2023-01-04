# DownloadClientResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ConfigContract** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to [**ProviderMessage**](ProviderMessage.md) |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presets** | Pointer to [**[]DownloadClientResource**](DownloadClientResource.md) |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RemoveCompletedDownloads** | Pointer to **bool** |  | [optional] 
**RemoveFailedDownloads** | Pointer to **bool** |  | [optional] 

## Methods

### NewDownloadClientResource

`func NewDownloadClientResource() *DownloadClientResource`

NewDownloadClientResource instantiates a new DownloadClientResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadClientResourceWithDefaults

`func NewDownloadClientResourceWithDefaults() *DownloadClientResource`

NewDownloadClientResourceWithDefaults instantiates a new DownloadClientResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DownloadClientResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DownloadClientResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DownloadClientResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DownloadClientResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DownloadClientResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DownloadClientResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DownloadClientResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DownloadClientResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DownloadClientResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DownloadClientResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *DownloadClientResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DownloadClientResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DownloadClientResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DownloadClientResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *DownloadClientResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *DownloadClientResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *DownloadClientResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *DownloadClientResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *DownloadClientResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *DownloadClientResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *DownloadClientResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *DownloadClientResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *DownloadClientResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *DownloadClientResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *DownloadClientResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *DownloadClientResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *DownloadClientResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *DownloadClientResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *DownloadClientResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *DownloadClientResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *DownloadClientResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *DownloadClientResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *DownloadClientResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *DownloadClientResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *DownloadClientResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *DownloadClientResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *DownloadClientResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *DownloadClientResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *DownloadClientResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *DownloadClientResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *DownloadClientResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DownloadClientResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DownloadClientResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DownloadClientResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *DownloadClientResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DownloadClientResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DownloadClientResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DownloadClientResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DownloadClientResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DownloadClientResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *DownloadClientResource) GetPresets() []DownloadClientResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *DownloadClientResource) GetPresetsOk() (*[]DownloadClientResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *DownloadClientResource) SetPresets(v []DownloadClientResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *DownloadClientResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *DownloadClientResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *DownloadClientResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetEnable

`func (o *DownloadClientResource) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DownloadClientResource) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DownloadClientResource) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DownloadClientResource) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetProtocol

`func (o *DownloadClientResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DownloadClientResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DownloadClientResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DownloadClientResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPriority

`func (o *DownloadClientResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DownloadClientResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DownloadClientResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DownloadClientResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRemoveCompletedDownloads

`func (o *DownloadClientResource) GetRemoveCompletedDownloads() bool`

GetRemoveCompletedDownloads returns the RemoveCompletedDownloads field if non-nil, zero value otherwise.

### GetRemoveCompletedDownloadsOk

`func (o *DownloadClientResource) GetRemoveCompletedDownloadsOk() (*bool, bool)`

GetRemoveCompletedDownloadsOk returns a tuple with the RemoveCompletedDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveCompletedDownloads

`func (o *DownloadClientResource) SetRemoveCompletedDownloads(v bool)`

SetRemoveCompletedDownloads sets RemoveCompletedDownloads field to given value.

### HasRemoveCompletedDownloads

`func (o *DownloadClientResource) HasRemoveCompletedDownloads() bool`

HasRemoveCompletedDownloads returns a boolean if a field has been set.

### GetRemoveFailedDownloads

`func (o *DownloadClientResource) GetRemoveFailedDownloads() bool`

GetRemoveFailedDownloads returns the RemoveFailedDownloads field if non-nil, zero value otherwise.

### GetRemoveFailedDownloadsOk

`func (o *DownloadClientResource) GetRemoveFailedDownloadsOk() (*bool, bool)`

GetRemoveFailedDownloadsOk returns a tuple with the RemoveFailedDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFailedDownloads

`func (o *DownloadClientResource) SetRemoveFailedDownloads(v bool)`

SetRemoveFailedDownloads sets RemoveFailedDownloads field to given value.

### HasRemoveFailedDownloads

`func (o *DownloadClientResource) HasRemoveFailedDownloads() bool`

HasRemoveFailedDownloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


