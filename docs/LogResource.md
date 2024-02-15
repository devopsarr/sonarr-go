# LogResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Exception** | Pointer to **NullableString** |  | [optional] 
**ExceptionType** | Pointer to **NullableString** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] 
**Logger** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogResource

`func NewLogResource() *LogResource`

NewLogResource instantiates a new LogResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResourceWithDefaults

`func NewLogResourceWithDefaults() *LogResource`

NewLogResourceWithDefaults instantiates a new LogResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LogResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *LogResource) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogResource) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LogResource) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *LogResource) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetException

`func (o *LogResource) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *LogResource) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *LogResource) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *LogResource) HasException() bool`

HasException returns a boolean if a field has been set.

### SetExceptionNil

`func (o *LogResource) SetExceptionNil(b bool)`

 SetExceptionNil sets the value for Exception to be an explicit nil

### UnsetException
`func (o *LogResource) UnsetException()`

UnsetException ensures that no value is present for Exception, not even an explicit nil
### GetExceptionType

`func (o *LogResource) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *LogResource) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *LogResource) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *LogResource) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### SetExceptionTypeNil

`func (o *LogResource) SetExceptionTypeNil(b bool)`

 SetExceptionTypeNil sets the value for ExceptionType to be an explicit nil

### UnsetExceptionType
`func (o *LogResource) UnsetExceptionType()`

UnsetExceptionType ensures that no value is present for ExceptionType, not even an explicit nil
### GetLevel

`func (o *LogResource) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogResource) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogResource) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LogResource) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *LogResource) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *LogResource) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetLogger

`func (o *LogResource) GetLogger() string`

GetLogger returns the Logger field if non-nil, zero value otherwise.

### GetLoggerOk

`func (o *LogResource) GetLoggerOk() (*string, bool)`

GetLoggerOk returns a tuple with the Logger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogger

`func (o *LogResource) SetLogger(v string)`

SetLogger sets Logger field to given value.

### HasLogger

`func (o *LogResource) HasLogger() bool`

HasLogger returns a boolean if a field has been set.

### SetLoggerNil

`func (o *LogResource) SetLoggerNil(b bool)`

 SetLoggerNil sets the value for Logger to be an explicit nil

### UnsetLogger
`func (o *LogResource) UnsetLogger()`

UnsetLogger ensures that no value is present for Logger, not even an explicit nil
### GetMessage

`func (o *LogResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogResource) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *LogResource) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *LogResource) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMethod

`func (o *LogResource) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogResource) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogResource) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogResource) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *LogResource) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *LogResource) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


