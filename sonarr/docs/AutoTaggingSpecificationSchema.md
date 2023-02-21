# AutoTaggingSpecificationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Negate** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 

## Methods

### NewAutoTaggingSpecificationSchema

`func NewAutoTaggingSpecificationSchema() *AutoTaggingSpecificationSchema`

NewAutoTaggingSpecificationSchema instantiates a new AutoTaggingSpecificationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTaggingSpecificationSchemaWithDefaults

`func NewAutoTaggingSpecificationSchemaWithDefaults() *AutoTaggingSpecificationSchema`

NewAutoTaggingSpecificationSchemaWithDefaults instantiates a new AutoTaggingSpecificationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTaggingSpecificationSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTaggingSpecificationSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTaggingSpecificationSchema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTaggingSpecificationSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutoTaggingSpecificationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTaggingSpecificationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTaggingSpecificationSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTaggingSpecificationSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTaggingSpecificationSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTaggingSpecificationSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImplementation

`func (o *AutoTaggingSpecificationSchema) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *AutoTaggingSpecificationSchema) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *AutoTaggingSpecificationSchema) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *AutoTaggingSpecificationSchema) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *AutoTaggingSpecificationSchema) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *AutoTaggingSpecificationSchema) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetImplementationName

`func (o *AutoTaggingSpecificationSchema) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *AutoTaggingSpecificationSchema) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *AutoTaggingSpecificationSchema) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *AutoTaggingSpecificationSchema) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *AutoTaggingSpecificationSchema) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *AutoTaggingSpecificationSchema) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetNegate

`func (o *AutoTaggingSpecificationSchema) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *AutoTaggingSpecificationSchema) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *AutoTaggingSpecificationSchema) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *AutoTaggingSpecificationSchema) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetRequired

`func (o *AutoTaggingSpecificationSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AutoTaggingSpecificationSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AutoTaggingSpecificationSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AutoTaggingSpecificationSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetFields

`func (o *AutoTaggingSpecificationSchema) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AutoTaggingSpecificationSchema) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AutoTaggingSpecificationSchema) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AutoTaggingSpecificationSchema) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *AutoTaggingSpecificationSchema) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *AutoTaggingSpecificationSchema) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


