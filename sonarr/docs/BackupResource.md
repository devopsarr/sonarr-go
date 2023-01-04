# BackupResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**BackupType**](BackupType.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBackupResource

`func NewBackupResource() *BackupResource`

NewBackupResource instantiates a new BackupResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupResourceWithDefaults

`func NewBackupResourceWithDefaults() *BackupResource`

NewBackupResourceWithDefaults instantiates a new BackupResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BackupResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BackupResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BackupResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *BackupResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BackupResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BackupResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BackupResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BackupResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BackupResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetType

`func (o *BackupResource) GetType() BackupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupResource) GetTypeOk() (*BackupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupResource) SetType(v BackupType)`

SetType sets Type field to given value.

### HasType

`func (o *BackupResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *BackupResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BackupResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BackupResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BackupResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTime

`func (o *BackupResource) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *BackupResource) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *BackupResource) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *BackupResource) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


