# RootFolderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Accessible** | Pointer to **bool** |  | [optional] 
**FreeSpace** | Pointer to **NullableInt64** |  | [optional] 
**UnmappedFolders** | Pointer to [**[]UnmappedFolder**](UnmappedFolder.md) |  | [optional] 

## Methods

### NewRootFolderResource

`func NewRootFolderResource() *RootFolderResource`

NewRootFolderResource instantiates a new RootFolderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootFolderResourceWithDefaults

`func NewRootFolderResourceWithDefaults() *RootFolderResource`

NewRootFolderResourceWithDefaults instantiates a new RootFolderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RootFolderResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RootFolderResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RootFolderResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RootFolderResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *RootFolderResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RootFolderResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RootFolderResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RootFolderResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *RootFolderResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *RootFolderResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetAccessible

`func (o *RootFolderResource) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *RootFolderResource) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *RootFolderResource) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *RootFolderResource) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### GetFreeSpace

`func (o *RootFolderResource) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *RootFolderResource) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *RootFolderResource) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *RootFolderResource) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *RootFolderResource) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *RootFolderResource) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetUnmappedFolders

`func (o *RootFolderResource) GetUnmappedFolders() []UnmappedFolder`

GetUnmappedFolders returns the UnmappedFolders field if non-nil, zero value otherwise.

### GetUnmappedFoldersOk

`func (o *RootFolderResource) GetUnmappedFoldersOk() (*[]UnmappedFolder, bool)`

GetUnmappedFoldersOk returns a tuple with the UnmappedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmappedFolders

`func (o *RootFolderResource) SetUnmappedFolders(v []UnmappedFolder)`

SetUnmappedFolders sets UnmappedFolders field to given value.

### HasUnmappedFolders

`func (o *RootFolderResource) HasUnmappedFolders() bool`

HasUnmappedFolders returns a boolean if a field has been set.

### SetUnmappedFoldersNil

`func (o *RootFolderResource) SetUnmappedFoldersNil(b bool)`

 SetUnmappedFoldersNil sets the value for UnmappedFolders to be an explicit nil

### UnsetUnmappedFolders
`func (o *RootFolderResource) UnsetUnmappedFolders()`

UnsetUnmappedFolders ensures that no value is present for UnmappedFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


