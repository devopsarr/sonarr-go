# QualityProfileQualityItemResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] 
**Items** | Pointer to [**[]QualityProfileQualityItemResource**](QualityProfileQualityItemResource.md) |  | [optional] 
**Allowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewQualityProfileQualityItemResource

`func NewQualityProfileQualityItemResource() *QualityProfileQualityItemResource`

NewQualityProfileQualityItemResource instantiates a new QualityProfileQualityItemResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityProfileQualityItemResourceWithDefaults

`func NewQualityProfileQualityItemResourceWithDefaults() *QualityProfileQualityItemResource`

NewQualityProfileQualityItemResourceWithDefaults instantiates a new QualityProfileQualityItemResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QualityProfileQualityItemResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QualityProfileQualityItemResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QualityProfileQualityItemResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QualityProfileQualityItemResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *QualityProfileQualityItemResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QualityProfileQualityItemResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QualityProfileQualityItemResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QualityProfileQualityItemResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QualityProfileQualityItemResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QualityProfileQualityItemResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuality

`func (o *QualityProfileQualityItemResource) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *QualityProfileQualityItemResource) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *QualityProfileQualityItemResource) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *QualityProfileQualityItemResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetItems

`func (o *QualityProfileQualityItemResource) GetItems() []QualityProfileQualityItemResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QualityProfileQualityItemResource) GetItemsOk() (*[]QualityProfileQualityItemResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QualityProfileQualityItemResource) SetItems(v []QualityProfileQualityItemResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *QualityProfileQualityItemResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *QualityProfileQualityItemResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *QualityProfileQualityItemResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetAllowed

`func (o *QualityProfileQualityItemResource) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *QualityProfileQualityItemResource) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *QualityProfileQualityItemResource) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *QualityProfileQualityItemResource) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


