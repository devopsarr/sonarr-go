# Rejection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**RejectionType**](RejectionType.md) |  | [optional] 

## Methods

### NewRejection

`func NewRejection() *Rejection`

NewRejection instantiates a new Rejection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectionWithDefaults

`func NewRejectionWithDefaults() *Rejection`

NewRejectionWithDefaults instantiates a new Rejection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Rejection) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Rejection) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Rejection) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Rejection) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *Rejection) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Rejection) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetType

`func (o *Rejection) GetType() RejectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rejection) GetTypeOk() (*RejectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rejection) SetType(v RejectionType)`

SetType sets Type field to given value.

### HasType

`func (o *Rejection) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


