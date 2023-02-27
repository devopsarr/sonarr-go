# CommandResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CommandName** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to [**Command**](Command.md) |  | [optional] 
**Priority** | Pointer to [**CommandPriority**](CommandPriority.md) |  | [optional] 
**Status** | Pointer to [**CommandStatus**](CommandStatus.md) |  | [optional] 
**Queued** | Pointer to **time.Time** |  | [optional] 
**Started** | Pointer to **NullableTime** |  | [optional] 
**Ended** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **NullableString** |  | [optional] 
**Trigger** | Pointer to [**CommandTrigger**](CommandTrigger.md) |  | [optional] 
**ClientUserAgent** | Pointer to **NullableString** |  | [optional] 
**StateChangeTime** | Pointer to **NullableTime** |  | [optional] 
**SendUpdatesToClient** | Pointer to **bool** |  | [optional] 
**UpdateScheduledTask** | Pointer to **bool** |  | [optional] 
**LastExecutionTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCommandResource

`func NewCommandResource() *CommandResource`

NewCommandResource instantiates a new CommandResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandResourceWithDefaults

`func NewCommandResourceWithDefaults() *CommandResource`

NewCommandResourceWithDefaults instantiates a new CommandResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommandResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommandResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CommandResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommandResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommandResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommandResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommandResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommandResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCommandName

`func (o *CommandResource) GetCommandName() string`

GetCommandName returns the CommandName field if non-nil, zero value otherwise.

### GetCommandNameOk

`func (o *CommandResource) GetCommandNameOk() (*string, bool)`

GetCommandNameOk returns a tuple with the CommandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandName

`func (o *CommandResource) SetCommandName(v string)`

SetCommandName sets CommandName field to given value.

### HasCommandName

`func (o *CommandResource) HasCommandName() bool`

HasCommandName returns a boolean if a field has been set.

### SetCommandNameNil

`func (o *CommandResource) SetCommandNameNil(b bool)`

 SetCommandNameNil sets the value for CommandName to be an explicit nil

### UnsetCommandName
`func (o *CommandResource) UnsetCommandName()`

UnsetCommandName ensures that no value is present for CommandName, not even an explicit nil
### GetMessage

`func (o *CommandResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommandResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommandResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommandResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CommandResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CommandResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetBody

`func (o *CommandResource) GetBody() Command`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommandResource) GetBodyOk() (*Command, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommandResource) SetBody(v Command)`

SetBody sets Body field to given value.

### HasBody

`func (o *CommandResource) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPriority

`func (o *CommandResource) GetPriority() CommandPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CommandResource) GetPriorityOk() (*CommandPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CommandResource) SetPriority(v CommandPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CommandResource) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *CommandResource) GetStatus() CommandStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommandResource) GetStatusOk() (*CommandStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommandResource) SetStatus(v CommandStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommandResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQueued

`func (o *CommandResource) GetQueued() time.Time`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *CommandResource) GetQueuedOk() (*time.Time, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *CommandResource) SetQueued(v time.Time)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *CommandResource) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetStarted

`func (o *CommandResource) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *CommandResource) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *CommandResource) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *CommandResource) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### SetStartedNil

`func (o *CommandResource) SetStartedNil(b bool)`

 SetStartedNil sets the value for Started to be an explicit nil

### UnsetStarted
`func (o *CommandResource) UnsetStarted()`

UnsetStarted ensures that no value is present for Started, not even an explicit nil
### GetEnded

`func (o *CommandResource) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *CommandResource) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *CommandResource) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *CommandResource) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *CommandResource) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *CommandResource) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetDuration

`func (o *CommandResource) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CommandResource) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CommandResource) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CommandResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetException

`func (o *CommandResource) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *CommandResource) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *CommandResource) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *CommandResource) HasException() bool`

HasException returns a boolean if a field has been set.

### SetExceptionNil

`func (o *CommandResource) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *CommandResource) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil
### GetTrigger

`func (o *CommandResource) GetTrigger() CommandTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CommandResource) GetTriggerOk() (*CommandTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CommandResource) SetTrigger(v CommandTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CommandResource) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetClientUserAgent

`func (o *CommandResource) GetClientUserAgent() string`

GetClientUserAgent returns the ClientUserAgent field if non-nil, zero value otherwise.

### GetClientUserAgentOk

`func (o *CommandResource) GetClientUserAgentOk() (*string, bool)`

GetClientUserAgentOk returns a tuple with the ClientUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserAgent

`func (o *CommandResource) SetClientUserAgent(v string)`

SetClientUserAgent sets ClientUserAgent field to given value.

### HasClientUserAgent

`func (o *CommandResource) HasClientUserAgent() bool`

HasClientUserAgent returns a boolean if a field has been set.

### SetClientUserAgentNil

`func (o *CommandResource) SetClientUserAgentNil(b bool)`

 SetClientUserAgentNil sets the value for ClientUserAgent to be an explicit nil

### UnsetClientUserAgent
`func (o *CommandResource) UnsetClientUserAgent()`

UnsetClientUserAgent ensures that no value is present for ClientUserAgent, not even an explicit nil
### GetStateChangeTime

`func (o *CommandResource) GetStateChangeTime() time.Time`

GetStateChangeTime returns the StateChangeTime field if non-nil, zero value otherwise.

### GetStateChangeTimeOk

`func (o *CommandResource) GetStateChangeTimeOk() (*time.Time, bool)`

GetStateChangeTimeOk returns a tuple with the StateChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeTime

`func (o *CommandResource) SetStateChangeTime(v time.Time)`

SetStateChangeTime sets StateChangeTime field to given value.

### HasStateChangeTime

`func (o *CommandResource) HasStateChangeTime() bool`

HasStateChangeTime returns a boolean if a field has been set.

### SetStateChangeTimeNil

`func (o *CommandResource) SetStateChangeTimeNil(b bool)`

 SetStateChangeTimeNil sets the value for StateChangeTime to be an explicit nil

### UnsetStateChangeTime
`func (o *CommandResource) UnsetStateChangeTime()`

UnsetStateChangeTime ensures that no value is present for StateChangeTime, not even an explicit nil
### GetSendUpdatesToClient

`func (o *CommandResource) GetSendUpdatesToClient() bool`

GetSendUpdatesToClient returns the SendUpdatesToClient field if non-nil, zero value otherwise.

### GetSendUpdatesToClientOk

`func (o *CommandResource) GetSendUpdatesToClientOk() (*bool, bool)`

GetSendUpdatesToClientOk returns a tuple with the SendUpdatesToClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendUpdatesToClient

`func (o *CommandResource) SetSendUpdatesToClient(v bool)`

SetSendUpdatesToClient sets SendUpdatesToClient field to given value.

### HasSendUpdatesToClient

`func (o *CommandResource) HasSendUpdatesToClient() bool`

HasSendUpdatesToClient returns a boolean if a field has been set.

### GetUpdateScheduledTask

`func (o *CommandResource) GetUpdateScheduledTask() bool`

GetUpdateScheduledTask returns the UpdateScheduledTask field if non-nil, zero value otherwise.

### GetUpdateScheduledTaskOk

`func (o *CommandResource) GetUpdateScheduledTaskOk() (*bool, bool)`

GetUpdateScheduledTaskOk returns a tuple with the UpdateScheduledTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateScheduledTask

`func (o *CommandResource) SetUpdateScheduledTask(v bool)`

SetUpdateScheduledTask sets UpdateScheduledTask field to given value.

### HasUpdateScheduledTask

`func (o *CommandResource) HasUpdateScheduledTask() bool`

HasUpdateScheduledTask returns a boolean if a field has been set.

### GetLastExecutionTime

`func (o *CommandResource) GetLastExecutionTime() time.Time`

GetLastExecutionTime returns the LastExecutionTime field if non-nil, zero value otherwise.

### GetLastExecutionTimeOk

`func (o *CommandResource) GetLastExecutionTimeOk() (*time.Time, bool)`

GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionTime

`func (o *CommandResource) SetLastExecutionTime(v time.Time)`

SetLastExecutionTime sets LastExecutionTime field to given value.

### HasLastExecutionTime

`func (o *CommandResource) HasLastExecutionTime() bool`

HasLastExecutionTime returns a boolean if a field has been set.

### SetLastExecutionTimeNil

`func (o *CommandResource) SetLastExecutionTimeNil(b bool)`

 SetLastExecutionTimeNil sets the value for LastExecutionTime to be an explicit nil

### UnsetLastExecutionTime
`func (o *CommandResource) UnsetLastExecutionTime()`

UnsetLastExecutionTime ensures that no value is present for LastExecutionTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


