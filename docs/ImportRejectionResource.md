# ImportRejectionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**RejectionType**](RejectionType.md) |  | [optional] 

## Methods

### NewImportRejectionResource

`func NewImportRejectionResource() *ImportRejectionResource`

NewImportRejectionResource instantiates a new ImportRejectionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRejectionResourceWithDefaults

`func NewImportRejectionResourceWithDefaults() *ImportRejectionResource`

NewImportRejectionResourceWithDefaults instantiates a new ImportRejectionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ImportRejectionResource) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ImportRejectionResource) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ImportRejectionResource) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ImportRejectionResource) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ImportRejectionResource) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ImportRejectionResource) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetType

`func (o *ImportRejectionResource) GetType() RejectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportRejectionResource) GetTypeOk() (*RejectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportRejectionResource) SetType(v RejectionType)`

SetType sets Type field to given value.

### HasType

`func (o *ImportRejectionResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


