# CustomFilterResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewCustomFilterResource

`func NewCustomFilterResource() *CustomFilterResource`

NewCustomFilterResource instantiates a new CustomFilterResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterResourceWithDefaults

`func NewCustomFilterResourceWithDefaults() *CustomFilterResource`

NewCustomFilterResourceWithDefaults instantiates a new CustomFilterResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFilterResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFilterResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFilterResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFilterResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomFilterResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFilterResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFilterResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomFilterResource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CustomFilterResource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomFilterResource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetLabel

`func (o *CustomFilterResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFilterResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFilterResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFilterResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CustomFilterResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CustomFilterResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetFilters

`func (o *CustomFilterResource) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CustomFilterResource) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CustomFilterResource) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CustomFilterResource) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *CustomFilterResource) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *CustomFilterResource) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


