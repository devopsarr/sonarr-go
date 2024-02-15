# CustomFormatSpecificationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Negate** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**Presets** | Pointer to [**[]CustomFormatSpecificationSchema**](CustomFormatSpecificationSchema.md) |  | [optional] 

## Methods

### NewCustomFormatSpecificationSchema

`func NewCustomFormatSpecificationSchema() *CustomFormatSpecificationSchema`

NewCustomFormatSpecificationSchema instantiates a new CustomFormatSpecificationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormatSpecificationSchemaWithDefaults

`func NewCustomFormatSpecificationSchemaWithDefaults() *CustomFormatSpecificationSchema`

NewCustomFormatSpecificationSchemaWithDefaults instantiates a new CustomFormatSpecificationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFormatSpecificationSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFormatSpecificationSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFormatSpecificationSchema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFormatSpecificationSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomFormatSpecificationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFormatSpecificationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFormatSpecificationSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFormatSpecificationSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomFormatSpecificationSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomFormatSpecificationSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImplementation

`func (o *CustomFormatSpecificationSchema) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *CustomFormatSpecificationSchema) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *CustomFormatSpecificationSchema) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *CustomFormatSpecificationSchema) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *CustomFormatSpecificationSchema) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *CustomFormatSpecificationSchema) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetImplementationName

`func (o *CustomFormatSpecificationSchema) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *CustomFormatSpecificationSchema) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *CustomFormatSpecificationSchema) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *CustomFormatSpecificationSchema) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *CustomFormatSpecificationSchema) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *CustomFormatSpecificationSchema) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetInfoLink

`func (o *CustomFormatSpecificationSchema) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *CustomFormatSpecificationSchema) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *CustomFormatSpecificationSchema) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *CustomFormatSpecificationSchema) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *CustomFormatSpecificationSchema) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *CustomFormatSpecificationSchema) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetNegate

`func (o *CustomFormatSpecificationSchema) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *CustomFormatSpecificationSchema) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *CustomFormatSpecificationSchema) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *CustomFormatSpecificationSchema) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetRequired

`func (o *CustomFormatSpecificationSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomFormatSpecificationSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomFormatSpecificationSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomFormatSpecificationSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetFields

`func (o *CustomFormatSpecificationSchema) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomFormatSpecificationSchema) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomFormatSpecificationSchema) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomFormatSpecificationSchema) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *CustomFormatSpecificationSchema) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *CustomFormatSpecificationSchema) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPresets

`func (o *CustomFormatSpecificationSchema) GetPresets() []CustomFormatSpecificationSchema`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *CustomFormatSpecificationSchema) GetPresetsOk() (*[]CustomFormatSpecificationSchema, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *CustomFormatSpecificationSchema) SetPresets(v []CustomFormatSpecificationSchema)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *CustomFormatSpecificationSchema) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *CustomFormatSpecificationSchema) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *CustomFormatSpecificationSchema) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


