# ImportListResource

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
**Presets** | Pointer to [**[]ImportListResource**](ImportListResource.md) |  | [optional] 
**EnableAutomaticAdd** | Pointer to **bool** |  | [optional] 
**ShouldMonitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**SeriesType** | Pointer to [**SeriesTypes**](SeriesTypes.md) |  | [optional] 
**SeasonFolder** | Pointer to **bool** |  | [optional] 
**ListType** | Pointer to [**ImportListType**](ImportListType.md) |  | [optional] 
**ListOrder** | Pointer to **int32** |  | [optional] 
**MinRefreshInterval** | Pointer to [**TimeSpan**](TimeSpan.md) |  | [optional] 

## Methods

### NewImportListResource

`func NewImportListResource() *ImportListResource`

NewImportListResource instantiates a new ImportListResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportListResourceWithDefaults

`func NewImportListResourceWithDefaults() *ImportListResource`

NewImportListResourceWithDefaults instantiates a new ImportListResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportListResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportListResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportListResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImportListResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImportListResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportListResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportListResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportListResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImportListResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImportListResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *ImportListResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ImportListResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ImportListResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ImportListResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ImportListResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ImportListResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *ImportListResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *ImportListResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *ImportListResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *ImportListResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *ImportListResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *ImportListResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *ImportListResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *ImportListResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *ImportListResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *ImportListResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *ImportListResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *ImportListResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *ImportListResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *ImportListResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *ImportListResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *ImportListResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *ImportListResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *ImportListResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *ImportListResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *ImportListResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *ImportListResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *ImportListResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *ImportListResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *ImportListResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *ImportListResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportListResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportListResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImportListResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *ImportListResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportListResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportListResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportListResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ImportListResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ImportListResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *ImportListResource) GetPresets() []ImportListResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *ImportListResource) GetPresetsOk() (*[]ImportListResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *ImportListResource) SetPresets(v []ImportListResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *ImportListResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *ImportListResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *ImportListResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetEnableAutomaticAdd

`func (o *ImportListResource) GetEnableAutomaticAdd() bool`

GetEnableAutomaticAdd returns the EnableAutomaticAdd field if non-nil, zero value otherwise.

### GetEnableAutomaticAddOk

`func (o *ImportListResource) GetEnableAutomaticAddOk() (*bool, bool)`

GetEnableAutomaticAddOk returns a tuple with the EnableAutomaticAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticAdd

`func (o *ImportListResource) SetEnableAutomaticAdd(v bool)`

SetEnableAutomaticAdd sets EnableAutomaticAdd field to given value.

### HasEnableAutomaticAdd

`func (o *ImportListResource) HasEnableAutomaticAdd() bool`

HasEnableAutomaticAdd returns a boolean if a field has been set.

### GetShouldMonitor

`func (o *ImportListResource) GetShouldMonitor() MonitorTypes`

GetShouldMonitor returns the ShouldMonitor field if non-nil, zero value otherwise.

### GetShouldMonitorOk

`func (o *ImportListResource) GetShouldMonitorOk() (*MonitorTypes, bool)`

GetShouldMonitorOk returns a tuple with the ShouldMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldMonitor

`func (o *ImportListResource) SetShouldMonitor(v MonitorTypes)`

SetShouldMonitor sets ShouldMonitor field to given value.

### HasShouldMonitor

`func (o *ImportListResource) HasShouldMonitor() bool`

HasShouldMonitor returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *ImportListResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *ImportListResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *ImportListResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *ImportListResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *ImportListResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *ImportListResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetQualityProfileId

`func (o *ImportListResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *ImportListResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *ImportListResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *ImportListResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetSeriesType

`func (o *ImportListResource) GetSeriesType() SeriesTypes`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *ImportListResource) GetSeriesTypeOk() (*SeriesTypes, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *ImportListResource) SetSeriesType(v SeriesTypes)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *ImportListResource) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetSeasonFolder

`func (o *ImportListResource) GetSeasonFolder() bool`

GetSeasonFolder returns the SeasonFolder field if non-nil, zero value otherwise.

### GetSeasonFolderOk

`func (o *ImportListResource) GetSeasonFolderOk() (*bool, bool)`

GetSeasonFolderOk returns a tuple with the SeasonFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonFolder

`func (o *ImportListResource) SetSeasonFolder(v bool)`

SetSeasonFolder sets SeasonFolder field to given value.

### HasSeasonFolder

`func (o *ImportListResource) HasSeasonFolder() bool`

HasSeasonFolder returns a boolean if a field has been set.

### GetListType

`func (o *ImportListResource) GetListType() ImportListType`

GetListType returns the ListType field if non-nil, zero value otherwise.

### GetListTypeOk

`func (o *ImportListResource) GetListTypeOk() (*ImportListType, bool)`

GetListTypeOk returns a tuple with the ListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListType

`func (o *ImportListResource) SetListType(v ImportListType)`

SetListType sets ListType field to given value.

### HasListType

`func (o *ImportListResource) HasListType() bool`

HasListType returns a boolean if a field has been set.

### GetListOrder

`func (o *ImportListResource) GetListOrder() int32`

GetListOrder returns the ListOrder field if non-nil, zero value otherwise.

### GetListOrderOk

`func (o *ImportListResource) GetListOrderOk() (*int32, bool)`

GetListOrderOk returns a tuple with the ListOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrder

`func (o *ImportListResource) SetListOrder(v int32)`

SetListOrder sets ListOrder field to given value.

### HasListOrder

`func (o *ImportListResource) HasListOrder() bool`

HasListOrder returns a boolean if a field has been set.

### GetMinRefreshInterval

`func (o *ImportListResource) GetMinRefreshInterval() TimeSpan`

GetMinRefreshInterval returns the MinRefreshInterval field if non-nil, zero value otherwise.

### GetMinRefreshIntervalOk

`func (o *ImportListResource) GetMinRefreshIntervalOk() (*TimeSpan, bool)`

GetMinRefreshIntervalOk returns a tuple with the MinRefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRefreshInterval

`func (o *ImportListResource) SetMinRefreshInterval(v TimeSpan)`

SetMinRefreshInterval sets MinRefreshInterval field to given value.

### HasMinRefreshInterval

`func (o *ImportListResource) HasMinRefreshInterval() bool`

HasMinRefreshInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


