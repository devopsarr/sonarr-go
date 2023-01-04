# IndexerResource

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
**Presets** | Pointer to [**[]IndexerResource**](IndexerResource.md) |  | [optional] 
**EnableRss** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSearch** | Pointer to **bool** |  | [optional] 
**EnableInteractiveSearch** | Pointer to **bool** |  | [optional] 
**SupportsRss** | Pointer to **bool** |  | [optional] 
**SupportsSearch** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**SeasonSearchMaximumSingleEpisodeAge** | Pointer to **int32** |  | [optional] 
**DownloadClientId** | Pointer to **int32** |  | [optional] 

## Methods

### NewIndexerResource

`func NewIndexerResource() *IndexerResource`

NewIndexerResource instantiates a new IndexerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerResourceWithDefaults

`func NewIndexerResourceWithDefaults() *IndexerResource`

NewIndexerResourceWithDefaults instantiates a new IndexerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IndexerResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexerResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexerResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexerResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IndexerResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IndexerResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *IndexerResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexerResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexerResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IndexerResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *IndexerResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *IndexerResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *IndexerResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *IndexerResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *IndexerResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *IndexerResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *IndexerResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *IndexerResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *IndexerResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *IndexerResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *IndexerResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *IndexerResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *IndexerResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *IndexerResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *IndexerResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *IndexerResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *IndexerResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *IndexerResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *IndexerResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *IndexerResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *IndexerResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *IndexerResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *IndexerResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *IndexerResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *IndexerResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *IndexerResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *IndexerResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IndexerResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IndexerResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IndexerResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *IndexerResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IndexerResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IndexerResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IndexerResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IndexerResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IndexerResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *IndexerResource) GetPresets() []IndexerResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *IndexerResource) GetPresetsOk() (*[]IndexerResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *IndexerResource) SetPresets(v []IndexerResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *IndexerResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *IndexerResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *IndexerResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetEnableRss

`func (o *IndexerResource) GetEnableRss() bool`

GetEnableRss returns the EnableRss field if non-nil, zero value otherwise.

### GetEnableRssOk

`func (o *IndexerResource) GetEnableRssOk() (*bool, bool)`

GetEnableRssOk returns a tuple with the EnableRss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRss

`func (o *IndexerResource) SetEnableRss(v bool)`

SetEnableRss sets EnableRss field to given value.

### HasEnableRss

`func (o *IndexerResource) HasEnableRss() bool`

HasEnableRss returns a boolean if a field has been set.

### GetEnableAutomaticSearch

`func (o *IndexerResource) GetEnableAutomaticSearch() bool`

GetEnableAutomaticSearch returns the EnableAutomaticSearch field if non-nil, zero value otherwise.

### GetEnableAutomaticSearchOk

`func (o *IndexerResource) GetEnableAutomaticSearchOk() (*bool, bool)`

GetEnableAutomaticSearchOk returns a tuple with the EnableAutomaticSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSearch

`func (o *IndexerResource) SetEnableAutomaticSearch(v bool)`

SetEnableAutomaticSearch sets EnableAutomaticSearch field to given value.

### HasEnableAutomaticSearch

`func (o *IndexerResource) HasEnableAutomaticSearch() bool`

HasEnableAutomaticSearch returns a boolean if a field has been set.

### GetEnableInteractiveSearch

`func (o *IndexerResource) GetEnableInteractiveSearch() bool`

GetEnableInteractiveSearch returns the EnableInteractiveSearch field if non-nil, zero value otherwise.

### GetEnableInteractiveSearchOk

`func (o *IndexerResource) GetEnableInteractiveSearchOk() (*bool, bool)`

GetEnableInteractiveSearchOk returns a tuple with the EnableInteractiveSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInteractiveSearch

`func (o *IndexerResource) SetEnableInteractiveSearch(v bool)`

SetEnableInteractiveSearch sets EnableInteractiveSearch field to given value.

### HasEnableInteractiveSearch

`func (o *IndexerResource) HasEnableInteractiveSearch() bool`

HasEnableInteractiveSearch returns a boolean if a field has been set.

### GetSupportsRss

`func (o *IndexerResource) GetSupportsRss() bool`

GetSupportsRss returns the SupportsRss field if non-nil, zero value otherwise.

### GetSupportsRssOk

`func (o *IndexerResource) GetSupportsRssOk() (*bool, bool)`

GetSupportsRssOk returns a tuple with the SupportsRss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRss

`func (o *IndexerResource) SetSupportsRss(v bool)`

SetSupportsRss sets SupportsRss field to given value.

### HasSupportsRss

`func (o *IndexerResource) HasSupportsRss() bool`

HasSupportsRss returns a boolean if a field has been set.

### GetSupportsSearch

`func (o *IndexerResource) GetSupportsSearch() bool`

GetSupportsSearch returns the SupportsSearch field if non-nil, zero value otherwise.

### GetSupportsSearchOk

`func (o *IndexerResource) GetSupportsSearchOk() (*bool, bool)`

GetSupportsSearchOk returns a tuple with the SupportsSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSearch

`func (o *IndexerResource) SetSupportsSearch(v bool)`

SetSupportsSearch sets SupportsSearch field to given value.

### HasSupportsSearch

`func (o *IndexerResource) HasSupportsSearch() bool`

HasSupportsSearch returns a boolean if a field has been set.

### GetProtocol

`func (o *IndexerResource) GetProtocol() DownloadProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IndexerResource) GetProtocolOk() (*DownloadProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IndexerResource) SetProtocol(v DownloadProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IndexerResource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPriority

`func (o *IndexerResource) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IndexerResource) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IndexerResource) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IndexerResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSeasonSearchMaximumSingleEpisodeAge

`func (o *IndexerResource) GetSeasonSearchMaximumSingleEpisodeAge() int32`

GetSeasonSearchMaximumSingleEpisodeAge returns the SeasonSearchMaximumSingleEpisodeAge field if non-nil, zero value otherwise.

### GetSeasonSearchMaximumSingleEpisodeAgeOk

`func (o *IndexerResource) GetSeasonSearchMaximumSingleEpisodeAgeOk() (*int32, bool)`

GetSeasonSearchMaximumSingleEpisodeAgeOk returns a tuple with the SeasonSearchMaximumSingleEpisodeAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonSearchMaximumSingleEpisodeAge

`func (o *IndexerResource) SetSeasonSearchMaximumSingleEpisodeAge(v int32)`

SetSeasonSearchMaximumSingleEpisodeAge sets SeasonSearchMaximumSingleEpisodeAge field to given value.

### HasSeasonSearchMaximumSingleEpisodeAge

`func (o *IndexerResource) HasSeasonSearchMaximumSingleEpisodeAge() bool`

HasSeasonSearchMaximumSingleEpisodeAge returns a boolean if a field has been set.

### GetDownloadClientId

`func (o *IndexerResource) GetDownloadClientId() int32`

GetDownloadClientId returns the DownloadClientId field if non-nil, zero value otherwise.

### GetDownloadClientIdOk

`func (o *IndexerResource) GetDownloadClientIdOk() (*int32, bool)`

GetDownloadClientIdOk returns a tuple with the DownloadClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadClientId

`func (o *IndexerResource) SetDownloadClientId(v int32)`

SetDownloadClientId sets DownloadClientId field to given value.

### HasDownloadClientId

`func (o *IndexerResource) HasDownloadClientId() bool`

HasDownloadClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


