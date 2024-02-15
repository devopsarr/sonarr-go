# TrackedDownloadStatusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Messages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTrackedDownloadStatusMessage

`func NewTrackedDownloadStatusMessage() *TrackedDownloadStatusMessage`

NewTrackedDownloadStatusMessage instantiates a new TrackedDownloadStatusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackedDownloadStatusMessageWithDefaults

`func NewTrackedDownloadStatusMessageWithDefaults() *TrackedDownloadStatusMessage`

NewTrackedDownloadStatusMessageWithDefaults instantiates a new TrackedDownloadStatusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TrackedDownloadStatusMessage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrackedDownloadStatusMessage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrackedDownloadStatusMessage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrackedDownloadStatusMessage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TrackedDownloadStatusMessage) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TrackedDownloadStatusMessage) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessages

`func (o *TrackedDownloadStatusMessage) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *TrackedDownloadStatusMessage) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *TrackedDownloadStatusMessage) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *TrackedDownloadStatusMessage) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *TrackedDownloadStatusMessage) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *TrackedDownloadStatusMessage) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


