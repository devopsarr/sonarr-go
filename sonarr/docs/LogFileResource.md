# LogFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Filename** | Pointer to **NullableString** |  | [optional] 
**LastWriteTime** | Pointer to **time.Time** |  | [optional] 
**ContentsUrl** | Pointer to **NullableString** |  | [optional] 
**DownloadUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogFileResource

`func NewLogFileResource() *LogFileResource`

NewLogFileResource instantiates a new LogFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFileResourceWithDefaults

`func NewLogFileResourceWithDefaults() *LogFileResource`

NewLogFileResourceWithDefaults instantiates a new LogFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LogFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFilename

`func (o *LogFileResource) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *LogFileResource) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *LogFileResource) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *LogFileResource) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *LogFileResource) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *LogFileResource) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetLastWriteTime

`func (o *LogFileResource) GetLastWriteTime() time.Time`

GetLastWriteTime returns the LastWriteTime field if non-nil, zero value otherwise.

### GetLastWriteTimeOk

`func (o *LogFileResource) GetLastWriteTimeOk() (*time.Time, bool)`

GetLastWriteTimeOk returns a tuple with the LastWriteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWriteTime

`func (o *LogFileResource) SetLastWriteTime(v time.Time)`

SetLastWriteTime sets LastWriteTime field to given value.

### HasLastWriteTime

`func (o *LogFileResource) HasLastWriteTime() bool`

HasLastWriteTime returns a boolean if a field has been set.

### GetContentsUrl

`func (o *LogFileResource) GetContentsUrl() string`

GetContentsUrl returns the ContentsUrl field if non-nil, zero value otherwise.

### GetContentsUrlOk

`func (o *LogFileResource) GetContentsUrlOk() (*string, bool)`

GetContentsUrlOk returns a tuple with the ContentsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsUrl

`func (o *LogFileResource) SetContentsUrl(v string)`

SetContentsUrl sets ContentsUrl field to given value.

### HasContentsUrl

`func (o *LogFileResource) HasContentsUrl() bool`

HasContentsUrl returns a boolean if a field has been set.

### SetContentsUrlNil

`func (o *LogFileResource) SetContentsUrlNil(b bool)`

 SetContentsUrlNil sets the value for ContentsUrl to be an explicit nil

### UnsetContentsUrl
`func (o *LogFileResource) UnsetContentsUrl()`

UnsetContentsUrl ensures that no value is present for ContentsUrl, not even an explicit nil
### GetDownloadUrl

`func (o *LogFileResource) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *LogFileResource) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *LogFileResource) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *LogFileResource) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *LogFileResource) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *LogFileResource) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


