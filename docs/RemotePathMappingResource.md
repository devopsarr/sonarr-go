# RemotePathMappingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**RemotePath** | Pointer to **NullableString** |  | [optional] 
**LocalPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRemotePathMappingResource

`func NewRemotePathMappingResource() *RemotePathMappingResource`

NewRemotePathMappingResource instantiates a new RemotePathMappingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePathMappingResourceWithDefaults

`func NewRemotePathMappingResourceWithDefaults() *RemotePathMappingResource`

NewRemotePathMappingResourceWithDefaults instantiates a new RemotePathMappingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemotePathMappingResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemotePathMappingResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemotePathMappingResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RemotePathMappingResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHost

`func (o *RemotePathMappingResource) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemotePathMappingResource) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemotePathMappingResource) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemotePathMappingResource) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *RemotePathMappingResource) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RemotePathMappingResource) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetRemotePath

`func (o *RemotePathMappingResource) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *RemotePathMappingResource) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *RemotePathMappingResource) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *RemotePathMappingResource) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### SetRemotePathNil

`func (o *RemotePathMappingResource) SetRemotePathNil(b bool)`

 SetRemotePathNil sets the value for RemotePath to be an explicit nil

### UnsetRemotePath
`func (o *RemotePathMappingResource) UnsetRemotePath()`

UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
### GetLocalPath

`func (o *RemotePathMappingResource) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *RemotePathMappingResource) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *RemotePathMappingResource) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.

### HasLocalPath

`func (o *RemotePathMappingResource) HasLocalPath() bool`

HasLocalPath returns a boolean if a field has been set.

### SetLocalPathNil

`func (o *RemotePathMappingResource) SetLocalPathNil(b bool)`

 SetLocalPathNil sets the value for LocalPath to be an explicit nil

### UnsetLocalPath
`func (o *RemotePathMappingResource) UnsetLocalPath()`

UnsetLocalPath ensures that no value is present for LocalPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


