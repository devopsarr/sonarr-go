# ProviderMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ProviderMessageType**](ProviderMessageType.md) |  | [optional] 

## Methods

### NewProviderMessage

`func NewProviderMessage() *ProviderMessage`

NewProviderMessage instantiates a new ProviderMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderMessageWithDefaults

`func NewProviderMessageWithDefaults() *ProviderMessage`

NewProviderMessageWithDefaults instantiates a new ProviderMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProviderMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProviderMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProviderMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProviderMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProviderMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProviderMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetType

`func (o *ProviderMessage) GetType() ProviderMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderMessage) GetTypeOk() (*ProviderMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderMessage) SetType(v ProviderMessageType)`

SetType sets Type field to given value.

### HasType

`func (o *ProviderMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


