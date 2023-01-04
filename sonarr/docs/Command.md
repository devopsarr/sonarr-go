# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendUpdatesToClient** | Pointer to **bool** |  | [optional] 
**UpdateScheduledTask** | Pointer to **bool** |  | [optional] [readonly] 
**CompletionMessage** | Pointer to **NullableString** |  | [optional] [readonly] 
**RequiresDiskAccess** | Pointer to **bool** |  | [optional] [readonly] 
**IsExclusive** | Pointer to **bool** |  | [optional] [readonly] 
**IsLongRunning** | Pointer to **bool** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**LastExecutionTime** | Pointer to **NullableTime** |  | [optional] 
**Trigger** | Pointer to [**CommandTrigger**](CommandTrigger.md) |  | [optional] 
**SuppressMessages** | Pointer to **bool** |  | [optional] 
**ClientUserAgent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendUpdatesToClient

`func (o *Command) GetSendUpdatesToClient() bool`

GetSendUpdatesToClient returns the SendUpdatesToClient field if non-nil, zero value otherwise.

### GetSendUpdatesToClientOk

`func (o *Command) GetSendUpdatesToClientOk() (*bool, bool)`

GetSendUpdatesToClientOk returns a tuple with the SendUpdatesToClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUpdatesToClient

`func (o *Command) SetSendUpdatesToClient(v bool)`

SetSendUpdatesToClient sets SendUpdatesToClient field to given value.

### HasSendUpdatesToClient

`func (o *Command) HasSendUpdatesToClient() bool`

HasSendUpdatesToClient returns a boolean if a field has been set.

### GetUpdateScheduledTask

`func (o *Command) GetUpdateScheduledTask() bool`

GetUpdateScheduledTask returns the UpdateScheduledTask field if non-nil, zero value otherwise.

### GetUpdateScheduledTaskOk

`func (o *Command) GetUpdateScheduledTaskOk() (*bool, bool)`

GetUpdateScheduledTaskOk returns a tuple with the UpdateScheduledTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateScheduledTask

`func (o *Command) SetUpdateScheduledTask(v bool)`

SetUpdateScheduledTask sets UpdateScheduledTask field to given value.

### HasUpdateScheduledTask

`func (o *Command) HasUpdateScheduledTask() bool`

HasUpdateScheduledTask returns a boolean if a field has been set.

### GetCompletionMessage

`func (o *Command) GetCompletionMessage() string`

GetCompletionMessage returns the CompletionMessage field if non-nil, zero value otherwise.

### GetCompletionMessageOk

`func (o *Command) GetCompletionMessageOk() (*string, bool)`

GetCompletionMessageOk returns a tuple with the CompletionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionMessage

`func (o *Command) SetCompletionMessage(v string)`

SetCompletionMessage sets CompletionMessage field to given value.

### HasCompletionMessage

`func (o *Command) HasCompletionMessage() bool`

HasCompletionMessage returns a boolean if a field has been set.

### SetCompletionMessageNil

`func (o *Command) SetCompletionMessageNil(b bool)`

 SetCompletionMessageNil sets the value for CompletionMessage to be an explicit nil

### UnsetCompletionMessage
`func (o *Command) UnsetCompletionMessage()`

UnsetCompletionMessage ensures that no value is present for CompletionMessage, not even an explicit nil
### GetRequiresDiskAccess

`func (o *Command) GetRequiresDiskAccess() bool`

GetRequiresDiskAccess returns the RequiresDiskAccess field if non-nil, zero value otherwise.

### GetRequiresDiskAccessOk

`func (o *Command) GetRequiresDiskAccessOk() (*bool, bool)`

GetRequiresDiskAccessOk returns a tuple with the RequiresDiskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDiskAccess

`func (o *Command) SetRequiresDiskAccess(v bool)`

SetRequiresDiskAccess sets RequiresDiskAccess field to given value.

### HasRequiresDiskAccess

`func (o *Command) HasRequiresDiskAccess() bool`

HasRequiresDiskAccess returns a boolean if a field has been set.

### GetIsExclusive

`func (o *Command) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *Command) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *Command) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *Command) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### GetIsLongRunning

`func (o *Command) GetIsLongRunning() bool`

GetIsLongRunning returns the IsLongRunning field if non-nil, zero value otherwise.

### GetIsLongRunningOk

`func (o *Command) GetIsLongRunningOk() (*bool, bool)`

GetIsLongRunningOk returns a tuple with the IsLongRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLongRunning

`func (o *Command) SetIsLongRunning(v bool)`

SetIsLongRunning sets IsLongRunning field to given value.

### HasIsLongRunning

`func (o *Command) HasIsLongRunning() bool`

HasIsLongRunning returns a boolean if a field has been set.

### GetName

`func (o *Command) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Command) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Command) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Command) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Command) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Command) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLastExecutionTime

`func (o *Command) GetLastExecutionTime() time.Time`

GetLastExecutionTime returns the LastExecutionTime field if non-nil, zero value otherwise.

### GetLastExecutionTimeOk

`func (o *Command) GetLastExecutionTimeOk() (*time.Time, bool)`

GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionTime

`func (o *Command) SetLastExecutionTime(v time.Time)`

SetLastExecutionTime sets LastExecutionTime field to given value.

### HasLastExecutionTime

`func (o *Command) HasLastExecutionTime() bool`

HasLastExecutionTime returns a boolean if a field has been set.

### SetLastExecutionTimeNil

`func (o *Command) SetLastExecutionTimeNil(b bool)`

 SetLastExecutionTimeNil sets the value for LastExecutionTime to be an explicit nil

### UnsetLastExecutionTime
`func (o *Command) UnsetLastExecutionTime()`

UnsetLastExecutionTime ensures that no value is present for LastExecutionTime, not even an explicit nil
### GetTrigger

`func (o *Command) GetTrigger() CommandTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Command) GetTriggerOk() (*CommandTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Command) SetTrigger(v CommandTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Command) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetSuppressMessages

`func (o *Command) GetSuppressMessages() bool`

GetSuppressMessages returns the SuppressMessages field if non-nil, zero value otherwise.

### GetSuppressMessagesOk

`func (o *Command) GetSuppressMessagesOk() (*bool, bool)`

GetSuppressMessagesOk returns a tuple with the SuppressMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressMessages

`func (o *Command) SetSuppressMessages(v bool)`

SetSuppressMessages sets SuppressMessages field to given value.

### HasSuppressMessages

`func (o *Command) HasSuppressMessages() bool`

HasSuppressMessages returns a boolean if a field has been set.

### GetClientUserAgent

`func (o *Command) GetClientUserAgent() string`

GetClientUserAgent returns the ClientUserAgent field if non-nil, zero value otherwise.

### GetClientUserAgentOk

`func (o *Command) GetClientUserAgentOk() (*string, bool)`

GetClientUserAgentOk returns a tuple with the ClientUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserAgent

`func (o *Command) SetClientUserAgent(v string)`

SetClientUserAgent sets ClientUserAgent field to given value.

### HasClientUserAgent

`func (o *Command) HasClientUserAgent() bool`

HasClientUserAgent returns a boolean if a field has been set.

### SetClientUserAgentNil

`func (o *Command) SetClientUserAgentNil(b bool)`

 SetClientUserAgentNil sets the value for ClientUserAgent to be an explicit nil

### UnsetClientUserAgent
`func (o *Command) UnsetClientUserAgent()`

UnsetClientUserAgent ensures that no value is present for ClientUserAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


