# HealthResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**HealthCheckResult**](HealthCheckResult.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**WikiUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthResource

`func NewHealthResource() *HealthResource`

NewHealthResource instantiates a new HealthResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResourceWithDefaults

`func NewHealthResourceWithDefaults() *HealthResource`

NewHealthResourceWithDefaults instantiates a new HealthResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HealthResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HealthResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *HealthResource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HealthResource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HealthResource) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HealthResource) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *HealthResource) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *HealthResource) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetType

`func (o *HealthResource) GetType() HealthCheckResult`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthResource) GetTypeOk() (*HealthCheckResult, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthResource) SetType(v HealthCheckResult)`

SetType sets Type field to given value.

### HasType

`func (o *HealthResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *HealthResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *HealthResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *HealthResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetWikiUrl

`func (o *HealthResource) GetWikiUrl() string`

GetWikiUrl returns the WikiUrl field if non-nil, zero value otherwise.

### GetWikiUrlOk

`func (o *HealthResource) GetWikiUrlOk() (*string, bool)`

GetWikiUrlOk returns a tuple with the WikiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiUrl

`func (o *HealthResource) SetWikiUrl(v string)`

SetWikiUrl sets WikiUrl field to given value.

### HasWikiUrl

`func (o *HealthResource) HasWikiUrl() bool`

HasWikiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


