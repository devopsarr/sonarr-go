# TaskResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TaskName** | Pointer to **NullableString** |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**LastExecution** | Pointer to **time.Time** |  | [optional] 
**LastStartTime** | Pointer to **time.Time** |  | [optional] 
**NextExecution** | Pointer to **time.Time** |  | [optional] 
**LastDuration** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskResource

`func NewTaskResource() *TaskResource`

NewTaskResource instantiates a new TaskResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResourceWithDefaults

`func NewTaskResourceWithDefaults() *TaskResource`

NewTaskResourceWithDefaults instantiates a new TaskResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaskResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaskResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaskName

`func (o *TaskResource) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskResource) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskResource) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *TaskResource) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *TaskResource) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *TaskResource) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetInterval

`func (o *TaskResource) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TaskResource) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TaskResource) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TaskResource) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLastExecution

`func (o *TaskResource) GetLastExecution() time.Time`

GetLastExecution returns the LastExecution field if non-nil, zero value otherwise.

### GetLastExecutionOk

`func (o *TaskResource) GetLastExecutionOk() (*time.Time, bool)`

GetLastExecutionOk returns a tuple with the LastExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecution

`func (o *TaskResource) SetLastExecution(v time.Time)`

SetLastExecution sets LastExecution field to given value.

### HasLastExecution

`func (o *TaskResource) HasLastExecution() bool`

HasLastExecution returns a boolean if a field has been set.

### GetLastStartTime

`func (o *TaskResource) GetLastStartTime() time.Time`

GetLastStartTime returns the LastStartTime field if non-nil, zero value otherwise.

### GetLastStartTimeOk

`func (o *TaskResource) GetLastStartTimeOk() (*time.Time, bool)`

GetLastStartTimeOk returns a tuple with the LastStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStartTime

`func (o *TaskResource) SetLastStartTime(v time.Time)`

SetLastStartTime sets LastStartTime field to given value.

### HasLastStartTime

`func (o *TaskResource) HasLastStartTime() bool`

HasLastStartTime returns a boolean if a field has been set.

### GetNextExecution

`func (o *TaskResource) GetNextExecution() time.Time`

GetNextExecution returns the NextExecution field if non-nil, zero value otherwise.

### GetNextExecutionOk

`func (o *TaskResource) GetNextExecutionOk() (*time.Time, bool)`

GetNextExecutionOk returns a tuple with the NextExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecution

`func (o *TaskResource) SetNextExecution(v time.Time)`

SetNextExecution sets NextExecution field to given value.

### HasNextExecution

`func (o *TaskResource) HasNextExecution() bool`

HasNextExecution returns a boolean if a field has been set.

### GetLastDuration

`func (o *TaskResource) GetLastDuration() string`

GetLastDuration returns the LastDuration field if non-nil, zero value otherwise.

### GetLastDurationOk

`func (o *TaskResource) GetLastDurationOk() (*string, bool)`

GetLastDurationOk returns a tuple with the LastDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDuration

`func (o *TaskResource) SetLastDuration(v string)`

SetLastDuration sets LastDuration field to given value.

### HasLastDuration

`func (o *TaskResource) HasLastDuration() bool`

HasLastDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


