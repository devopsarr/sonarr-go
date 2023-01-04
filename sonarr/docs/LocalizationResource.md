# LocalizationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Strings** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLocalizationResource

`func NewLocalizationResource() *LocalizationResource`

NewLocalizationResource instantiates a new LocalizationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizationResourceWithDefaults

`func NewLocalizationResourceWithDefaults() *LocalizationResource`

NewLocalizationResourceWithDefaults instantiates a new LocalizationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalizationResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalizationResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalizationResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LocalizationResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStrings

`func (o *LocalizationResource) GetStrings() map[string]string`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *LocalizationResource) GetStringsOk() (*map[string]string, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *LocalizationResource) SetStrings(v map[string]string)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *LocalizationResource) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### SetStringsNil

`func (o *LocalizationResource) SetStringsNil(b bool)`

 SetStringsNil sets the value for Strings to be an explicit nil

### UnsetStrings
`func (o *LocalizationResource) UnsetStrings()`

UnsetStrings ensures that no value is present for Strings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


