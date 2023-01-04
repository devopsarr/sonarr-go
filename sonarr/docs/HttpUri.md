# HttpUri

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullUri** | Pointer to **NullableString** |  | [optional] [readonly] 
**Scheme** | Pointer to **NullableString** |  | [optional] [readonly] 
**Host** | Pointer to **NullableString** |  | [optional] [readonly] 
**Port** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Path** | Pointer to **NullableString** |  | [optional] [readonly] 
**Query** | Pointer to **NullableString** |  | [optional] [readonly] 
**Fragment** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewHttpUri

`func NewHttpUri() *HttpUri`

NewHttpUri instantiates a new HttpUri object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpUriWithDefaults

`func NewHttpUriWithDefaults() *HttpUri`

NewHttpUriWithDefaults instantiates a new HttpUri object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullUri

`func (o *HttpUri) GetFullUri() string`

GetFullUri returns the FullUri field if non-nil, zero value otherwise.

### GetFullUriOk

`func (o *HttpUri) GetFullUriOk() (*string, bool)`

GetFullUriOk returns a tuple with the FullUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullUri

`func (o *HttpUri) SetFullUri(v string)`

SetFullUri sets FullUri field to given value.

### HasFullUri

`func (o *HttpUri) HasFullUri() bool`

HasFullUri returns a boolean if a field has been set.

### SetFullUriNil

`func (o *HttpUri) SetFullUriNil(b bool)`

 SetFullUriNil sets the value for FullUri to be an explicit nil

### UnsetFullUri
`func (o *HttpUri) UnsetFullUri()`

UnsetFullUri ensures that no value is present for FullUri, not even an explicit nil
### GetScheme

`func (o *HttpUri) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *HttpUri) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *HttpUri) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *HttpUri) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *HttpUri) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *HttpUri) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetHost

`func (o *HttpUri) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpUri) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpUri) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HttpUri) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HttpUri) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HttpUri) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HttpUri) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpUri) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpUri) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HttpUri) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HttpUri) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HttpUri) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPath

`func (o *HttpUri) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HttpUri) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HttpUri) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HttpUri) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *HttpUri) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *HttpUri) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQuery

`func (o *HttpUri) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HttpUri) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HttpUri) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *HttpUri) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *HttpUri) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *HttpUri) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetFragment

`func (o *HttpUri) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *HttpUri) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *HttpUri) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *HttpUri) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### SetFragmentNil

`func (o *HttpUri) SetFragmentNil(b bool)`

 SetFragmentNil sets the value for Fragment to be an explicit nil

### UnsetFragment
`func (o *HttpUri) UnsetFragment()`

UnsetFragment ensures that no value is present for Fragment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


