# LanguageProfileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**UpgradeAllowed** | Pointer to **bool** |  | [optional] 
**Cutoff** | Pointer to [**Language**](Language.md) |  | [optional] 
**Languages** | Pointer to [**[]LanguageProfileItemResource**](LanguageProfileItemResource.md) |  | [optional] 

## Methods

### NewLanguageProfileResource

`func NewLanguageProfileResource() *LanguageProfileResource`

NewLanguageProfileResource instantiates a new LanguageProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageProfileResourceWithDefaults

`func NewLanguageProfileResourceWithDefaults() *LanguageProfileResource`

NewLanguageProfileResourceWithDefaults instantiates a new LanguageProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanguageProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanguageProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanguageProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LanguageProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LanguageProfileResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanguageProfileResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanguageProfileResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanguageProfileResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LanguageProfileResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LanguageProfileResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUpgradeAllowed

`func (o *LanguageProfileResource) GetUpgradeAllowed() bool`

GetUpgradeAllowed returns the UpgradeAllowed field if non-nil, zero value otherwise.

### GetUpgradeAllowedOk

`func (o *LanguageProfileResource) GetUpgradeAllowedOk() (*bool, bool)`

GetUpgradeAllowedOk returns a tuple with the UpgradeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeAllowed

`func (o *LanguageProfileResource) SetUpgradeAllowed(v bool)`

SetUpgradeAllowed sets UpgradeAllowed field to given value.

### HasUpgradeAllowed

`func (o *LanguageProfileResource) HasUpgradeAllowed() bool`

HasUpgradeAllowed returns a boolean if a field has been set.

### GetCutoff

`func (o *LanguageProfileResource) GetCutoff() Language`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *LanguageProfileResource) GetCutoffOk() (*Language, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *LanguageProfileResource) SetCutoff(v Language)`

SetCutoff sets Cutoff field to given value.

### HasCutoff

`func (o *LanguageProfileResource) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### GetLanguages

`func (o *LanguageProfileResource) GetLanguages() []LanguageProfileItemResource`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *LanguageProfileResource) GetLanguagesOk() (*[]LanguageProfileItemResource, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *LanguageProfileResource) SetLanguages(v []LanguageProfileItemResource)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *LanguageProfileResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *LanguageProfileResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *LanguageProfileResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


