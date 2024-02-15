# TagResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTagResource

`func NewTagResource() *TagResource`

NewTagResource instantiates a new TagResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResourceWithDefaults

`func NewTagResourceWithDefaults() *TagResource`

NewTagResourceWithDefaults instantiates a new TagResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TagResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TagResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TagResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TagResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TagResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *TagResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *TagResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


