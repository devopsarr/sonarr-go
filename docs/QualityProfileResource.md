# QualityProfileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**UpgradeAllowed** | Pointer to **bool** |  | [optional] 
**Cutoff** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]QualityProfileQualityItemResource**](QualityProfileQualityItemResource.md) |  | [optional] 
**MinFormatScore** | Pointer to **int32** |  | [optional] 
**CutoffFormatScore** | Pointer to **int32** |  | [optional] 
**MinUpgradeFormatScore** | Pointer to **int32** |  | [optional] 
**FormatItems** | Pointer to [**[]ProfileFormatItemResource**](ProfileFormatItemResource.md) |  | [optional] 

## Methods

### NewQualityProfileResource

`func NewQualityProfileResource() *QualityProfileResource`

NewQualityProfileResource instantiates a new QualityProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityProfileResourceWithDefaults

`func NewQualityProfileResourceWithDefaults() *QualityProfileResource`

NewQualityProfileResourceWithDefaults instantiates a new QualityProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QualityProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QualityProfileResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QualityProfileResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QualityProfileResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QualityProfileResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QualityProfileResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QualityProfileResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUpgradeAllowed

`func (o *QualityProfileResource) GetUpgradeAllowed() bool`

GetUpgradeAllowed returns the UpgradeAllowed field if non-nil, zero value otherwise.

### GetUpgradeAllowedOk

`func (o *QualityProfileResource) GetUpgradeAllowedOk() (*bool, bool)`

GetUpgradeAllowedOk returns a tuple with the UpgradeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeAllowed

`func (o *QualityProfileResource) SetUpgradeAllowed(v bool)`

SetUpgradeAllowed sets UpgradeAllowed field to given value.

### HasUpgradeAllowed

`func (o *QualityProfileResource) HasUpgradeAllowed() bool`

HasUpgradeAllowed returns a boolean if a field has been set.

### GetCutoff

`func (o *QualityProfileResource) GetCutoff() int32`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *QualityProfileResource) GetCutoffOk() (*int32, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *QualityProfileResource) SetCutoff(v int32)`

SetCutoff sets Cutoff field to given value.

### HasCutoff

`func (o *QualityProfileResource) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### GetItems

`func (o *QualityProfileResource) GetItems() []QualityProfileQualityItemResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QualityProfileResource) GetItemsOk() (*[]QualityProfileQualityItemResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QualityProfileResource) SetItems(v []QualityProfileQualityItemResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *QualityProfileResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *QualityProfileResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *QualityProfileResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetMinFormatScore

`func (o *QualityProfileResource) GetMinFormatScore() int32`

GetMinFormatScore returns the MinFormatScore field if non-nil, zero value otherwise.

### GetMinFormatScoreOk

`func (o *QualityProfileResource) GetMinFormatScoreOk() (*int32, bool)`

GetMinFormatScoreOk returns a tuple with the MinFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFormatScore

`func (o *QualityProfileResource) SetMinFormatScore(v int32)`

SetMinFormatScore sets MinFormatScore field to given value.

### HasMinFormatScore

`func (o *QualityProfileResource) HasMinFormatScore() bool`

HasMinFormatScore returns a boolean if a field has been set.

### GetCutoffFormatScore

`func (o *QualityProfileResource) GetCutoffFormatScore() int32`

GetCutoffFormatScore returns the CutoffFormatScore field if non-nil, zero value otherwise.

### GetCutoffFormatScoreOk

`func (o *QualityProfileResource) GetCutoffFormatScoreOk() (*int32, bool)`

GetCutoffFormatScoreOk returns a tuple with the CutoffFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoffFormatScore

`func (o *QualityProfileResource) SetCutoffFormatScore(v int32)`

SetCutoffFormatScore sets CutoffFormatScore field to given value.

### HasCutoffFormatScore

`func (o *QualityProfileResource) HasCutoffFormatScore() bool`

HasCutoffFormatScore returns a boolean if a field has been set.

### GetMinUpgradeFormatScore

`func (o *QualityProfileResource) GetMinUpgradeFormatScore() int32`

GetMinUpgradeFormatScore returns the MinUpgradeFormatScore field if non-nil, zero value otherwise.

### GetMinUpgradeFormatScoreOk

`func (o *QualityProfileResource) GetMinUpgradeFormatScoreOk() (*int32, bool)`

GetMinUpgradeFormatScoreOk returns a tuple with the MinUpgradeFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpgradeFormatScore

`func (o *QualityProfileResource) SetMinUpgradeFormatScore(v int32)`

SetMinUpgradeFormatScore sets MinUpgradeFormatScore field to given value.

### HasMinUpgradeFormatScore

`func (o *QualityProfileResource) HasMinUpgradeFormatScore() bool`

HasMinUpgradeFormatScore returns a boolean if a field has been set.

### GetFormatItems

`func (o *QualityProfileResource) GetFormatItems() []ProfileFormatItemResource`

GetFormatItems returns the FormatItems field if non-nil, zero value otherwise.

### GetFormatItemsOk

`func (o *QualityProfileResource) GetFormatItemsOk() (*[]ProfileFormatItemResource, bool)`

GetFormatItemsOk returns a tuple with the FormatItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatItems

`func (o *QualityProfileResource) SetFormatItems(v []ProfileFormatItemResource)`

SetFormatItems sets FormatItems field to given value.

### HasFormatItems

`func (o *QualityProfileResource) HasFormatItems() bool`

HasFormatItems returns a boolean if a field has been set.

### SetFormatItemsNil

`func (o *QualityProfileResource) SetFormatItemsNil(b bool)`

 SetFormatItemsNil sets the value for FormatItems to be an explicit nil

### UnsetFormatItems
`func (o *QualityProfileResource) UnsetFormatItems()`

UnsetFormatItems ensures that no value is present for FormatItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


