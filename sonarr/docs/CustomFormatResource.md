# CustomFormatResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IncludeCustomFormatWhenRenaming** | Pointer to **NullableBool** |  | [optional] 
**Specifications** | Pointer to [**[]CustomFormatSpecificationSchema**](CustomFormatSpecificationSchema.md) |  | [optional] 

## Methods

### NewCustomFormatResource

`func NewCustomFormatResource() *CustomFormatResource`

NewCustomFormatResource instantiates a new CustomFormatResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormatResourceWithDefaults

`func NewCustomFormatResourceWithDefaults() *CustomFormatResource`

NewCustomFormatResourceWithDefaults instantiates a new CustomFormatResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFormatResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFormatResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFormatResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFormatResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomFormatResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFormatResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFormatResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFormatResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomFormatResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomFormatResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIncludeCustomFormatWhenRenaming

`func (o *CustomFormatResource) GetIncludeCustomFormatWhenRenaming() bool`

GetIncludeCustomFormatWhenRenaming returns the IncludeCustomFormatWhenRenaming field if non-nil, zero value otherwise.

### GetIncludeCustomFormatWhenRenamingOk

`func (o *CustomFormatResource) GetIncludeCustomFormatWhenRenamingOk() (*bool, bool)`

GetIncludeCustomFormatWhenRenamingOk returns a tuple with the IncludeCustomFormatWhenRenaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomFormatWhenRenaming

`func (o *CustomFormatResource) SetIncludeCustomFormatWhenRenaming(v bool)`

SetIncludeCustomFormatWhenRenaming sets IncludeCustomFormatWhenRenaming field to given value.

### HasIncludeCustomFormatWhenRenaming

`func (o *CustomFormatResource) HasIncludeCustomFormatWhenRenaming() bool`

HasIncludeCustomFormatWhenRenaming returns a boolean if a field has been set.

### SetIncludeCustomFormatWhenRenamingNil

`func (o *CustomFormatResource) SetIncludeCustomFormatWhenRenamingNil(b bool)`

 SetIncludeCustomFormatWhenRenamingNil sets the value for IncludeCustomFormatWhenRenaming to be an explicit nil

### UnsetIncludeCustomFormatWhenRenaming
`func (o *CustomFormatResource) UnsetIncludeCustomFormatWhenRenaming()`

UnsetIncludeCustomFormatWhenRenaming ensures that no value is present for IncludeCustomFormatWhenRenaming, not even an explicit nil
### GetSpecifications

`func (o *CustomFormatResource) GetSpecifications() []CustomFormatSpecificationSchema`

GetSpecifications returns the Specifications field if non-nil, zero value otherwise.

### GetSpecificationsOk

`func (o *CustomFormatResource) GetSpecificationsOk() (*[]CustomFormatSpecificationSchema, bool)`

GetSpecificationsOk returns a tuple with the Specifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifications

`func (o *CustomFormatResource) SetSpecifications(v []CustomFormatSpecificationSchema)`

SetSpecifications sets Specifications field to given value.

### HasSpecifications

`func (o *CustomFormatResource) HasSpecifications() bool`

HasSpecifications returns a boolean if a field has been set.

### SetSpecificationsNil

`func (o *CustomFormatResource) SetSpecificationsNil(b bool)`

 SetSpecificationsNil sets the value for Specifications to be an explicit nil

### UnsetSpecifications
`func (o *CustomFormatResource) UnsetSpecifications()`

UnsetSpecifications ensures that no value is present for Specifications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


