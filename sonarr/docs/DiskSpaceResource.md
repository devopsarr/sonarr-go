# DiskSpaceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**TotalSpace** | Pointer to **int64** |  | [optional] 

## Methods

### NewDiskSpaceResource

`func NewDiskSpaceResource() *DiskSpaceResource`

NewDiskSpaceResource instantiates a new DiskSpaceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSpaceResourceWithDefaults

`func NewDiskSpaceResourceWithDefaults() *DiskSpaceResource`

NewDiskSpaceResourceWithDefaults instantiates a new DiskSpaceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskSpaceResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskSpaceResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskSpaceResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DiskSpaceResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DiskSpaceResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DiskSpaceResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DiskSpaceResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DiskSpaceResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *DiskSpaceResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *DiskSpaceResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetLabel

`func (o *DiskSpaceResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DiskSpaceResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DiskSpaceResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DiskSpaceResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *DiskSpaceResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *DiskSpaceResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetFreeSpace

`func (o *DiskSpaceResource) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *DiskSpaceResource) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *DiskSpaceResource) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *DiskSpaceResource) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetTotalSpace

`func (o *DiskSpaceResource) GetTotalSpace() int64`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *DiskSpaceResource) GetTotalSpaceOk() (*int64, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *DiskSpaceResource) SetTotalSpace(v int64)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *DiskSpaceResource) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


