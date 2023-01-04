# SelectOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSelectOption

`func NewSelectOption() *SelectOption`

NewSelectOption instantiates a new SelectOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectOptionWithDefaults

`func NewSelectOptionWithDefaults() *SelectOption`

NewSelectOptionWithDefaults instantiates a new SelectOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SelectOption) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SelectOption) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SelectOption) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SelectOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *SelectOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectOption) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SelectOption) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SelectOption) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *SelectOption) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SelectOption) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SelectOption) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SelectOption) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetHint

`func (o *SelectOption) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *SelectOption) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *SelectOption) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *SelectOption) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *SelectOption) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *SelectOption) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


