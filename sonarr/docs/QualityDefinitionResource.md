# QualityDefinitionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**MinSize** | Pointer to **NullableFloat64** |  | [optional] 
**MaxSize** | Pointer to **NullableFloat64** |  | [optional] 
**PreferredSize** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewQualityDefinitionResource

`func NewQualityDefinitionResource() *QualityDefinitionResource`

NewQualityDefinitionResource instantiates a new QualityDefinitionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityDefinitionResourceWithDefaults

`func NewQualityDefinitionResourceWithDefaults() *QualityDefinitionResource`

NewQualityDefinitionResourceWithDefaults instantiates a new QualityDefinitionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityDefinitionResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityDefinitionResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityDefinitionResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QualityDefinitionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuality

`func (o *QualityDefinitionResource) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *QualityDefinitionResource) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *QualityDefinitionResource) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *QualityDefinitionResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetTitle

`func (o *QualityDefinitionResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QualityDefinitionResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QualityDefinitionResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QualityDefinitionResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QualityDefinitionResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QualityDefinitionResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWeight

`func (o *QualityDefinitionResource) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QualityDefinitionResource) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QualityDefinitionResource) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QualityDefinitionResource) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMinSize

`func (o *QualityDefinitionResource) GetMinSize() float64`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *QualityDefinitionResource) GetMinSizeOk() (*float64, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *QualityDefinitionResource) SetMinSize(v float64)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *QualityDefinitionResource) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### SetMinSizeNil

`func (o *QualityDefinitionResource) SetMinSizeNil(b bool)`

 SetMinSizeNil sets the value for MinSize to be an explicit nil

### UnsetMinSize
`func (o *QualityDefinitionResource) UnsetMinSize()`

UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
### GetMaxSize

`func (o *QualityDefinitionResource) GetMaxSize() float64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *QualityDefinitionResource) GetMaxSizeOk() (*float64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *QualityDefinitionResource) SetMaxSize(v float64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *QualityDefinitionResource) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *QualityDefinitionResource) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *QualityDefinitionResource) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetPreferredSize

`func (o *QualityDefinitionResource) GetPreferredSize() float64`

GetPreferredSize returns the PreferredSize field if non-nil, zero value otherwise.

### GetPreferredSizeOk

`func (o *QualityDefinitionResource) GetPreferredSizeOk() (*float64, bool)`

GetPreferredSizeOk returns a tuple with the PreferredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSize

`func (o *QualityDefinitionResource) SetPreferredSize(v float64)`

SetPreferredSize sets PreferredSize field to given value.

### HasPreferredSize

`func (o *QualityDefinitionResource) HasPreferredSize() bool`

HasPreferredSize returns a boolean if a field has been set.

### SetPreferredSizeNil

`func (o *QualityDefinitionResource) SetPreferredSizeNil(b bool)`

 SetPreferredSizeNil sets the value for PreferredSize to be an explicit nil

### UnsetPreferredSize
`func (o *QualityDefinitionResource) UnsetPreferredSize()`

UnsetPreferredSize ensures that no value is present for PreferredSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


