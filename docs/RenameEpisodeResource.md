# RenameEpisodeResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SeriesId** | Pointer to **int32** |  | [optional] 
**SeasonNumber** | Pointer to **int32** |  | [optional] 
**EpisodeNumbers** | Pointer to **[]int32** |  | [optional] 
**EpisodeFileId** | Pointer to **int32** |  | [optional] 
**ExistingPath** | Pointer to **NullableString** |  | [optional] 
**NewPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRenameEpisodeResource

`func NewRenameEpisodeResource() *RenameEpisodeResource`

NewRenameEpisodeResource instantiates a new RenameEpisodeResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenameEpisodeResourceWithDefaults

`func NewRenameEpisodeResourceWithDefaults() *RenameEpisodeResource`

NewRenameEpisodeResourceWithDefaults instantiates a new RenameEpisodeResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RenameEpisodeResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenameEpisodeResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenameEpisodeResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RenameEpisodeResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeriesId

`func (o *RenameEpisodeResource) GetSeriesId() int32`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *RenameEpisodeResource) GetSeriesIdOk() (*int32, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *RenameEpisodeResource) SetSeriesId(v int32)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *RenameEpisodeResource) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *RenameEpisodeResource) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *RenameEpisodeResource) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *RenameEpisodeResource) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *RenameEpisodeResource) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetEpisodeNumbers

`func (o *RenameEpisodeResource) GetEpisodeNumbers() []int32`

GetEpisodeNumbers returns the EpisodeNumbers field if non-nil, zero value otherwise.

### GetEpisodeNumbersOk

`func (o *RenameEpisodeResource) GetEpisodeNumbersOk() (*[]int32, bool)`

GetEpisodeNumbersOk returns a tuple with the EpisodeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeNumbers

`func (o *RenameEpisodeResource) SetEpisodeNumbers(v []int32)`

SetEpisodeNumbers sets EpisodeNumbers field to given value.

### HasEpisodeNumbers

`func (o *RenameEpisodeResource) HasEpisodeNumbers() bool`

HasEpisodeNumbers returns a boolean if a field has been set.

### SetEpisodeNumbersNil

`func (o *RenameEpisodeResource) SetEpisodeNumbersNil(b bool)`

 SetEpisodeNumbersNil sets the value for EpisodeNumbers to be an explicit nil

### UnsetEpisodeNumbers
`func (o *RenameEpisodeResource) UnsetEpisodeNumbers()`

UnsetEpisodeNumbers ensures that no value is present for EpisodeNumbers, not even an explicit nil
### GetEpisodeFileId

`func (o *RenameEpisodeResource) GetEpisodeFileId() int32`

GetEpisodeFileId returns the EpisodeFileId field if non-nil, zero value otherwise.

### GetEpisodeFileIdOk

`func (o *RenameEpisodeResource) GetEpisodeFileIdOk() (*int32, bool)`

GetEpisodeFileIdOk returns a tuple with the EpisodeFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFileId

`func (o *RenameEpisodeResource) SetEpisodeFileId(v int32)`

SetEpisodeFileId sets EpisodeFileId field to given value.

### HasEpisodeFileId

`func (o *RenameEpisodeResource) HasEpisodeFileId() bool`

HasEpisodeFileId returns a boolean if a field has been set.

### GetExistingPath

`func (o *RenameEpisodeResource) GetExistingPath() string`

GetExistingPath returns the ExistingPath field if non-nil, zero value otherwise.

### GetExistingPathOk

`func (o *RenameEpisodeResource) GetExistingPathOk() (*string, bool)`

GetExistingPathOk returns a tuple with the ExistingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingPath

`func (o *RenameEpisodeResource) SetExistingPath(v string)`

SetExistingPath sets ExistingPath field to given value.

### HasExistingPath

`func (o *RenameEpisodeResource) HasExistingPath() bool`

HasExistingPath returns a boolean if a field has been set.

### SetExistingPathNil

`func (o *RenameEpisodeResource) SetExistingPathNil(b bool)`

 SetExistingPathNil sets the value for ExistingPath to be an explicit nil

### UnsetExistingPath
`func (o *RenameEpisodeResource) UnsetExistingPath()`

UnsetExistingPath ensures that no value is present for ExistingPath, not even an explicit nil
### GetNewPath

`func (o *RenameEpisodeResource) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *RenameEpisodeResource) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *RenameEpisodeResource) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *RenameEpisodeResource) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### SetNewPathNil

`func (o *RenameEpisodeResource) SetNewPathNil(b bool)`

 SetNewPathNil sets the value for NewPath to be an explicit nil

### UnsetNewPath
`func (o *RenameEpisodeResource) UnsetNewPath()`

UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


