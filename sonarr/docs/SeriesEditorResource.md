# SeriesEditorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesIds** | Pointer to **[]int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 
**SeriesType** | Pointer to [**SeriesTypes**](SeriesTypes.md) |  | [optional] 
**SeasonFolder** | Pointer to **NullableBool** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**MoveFiles** | Pointer to **bool** |  | [optional] 
**DeleteFiles** | Pointer to **bool** |  | [optional] 
**AddImportListExclusion** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeriesEditorResource

`func NewSeriesEditorResource() *SeriesEditorResource`

NewSeriesEditorResource instantiates a new SeriesEditorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesEditorResourceWithDefaults

`func NewSeriesEditorResourceWithDefaults() *SeriesEditorResource`

NewSeriesEditorResourceWithDefaults instantiates a new SeriesEditorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesIds

`func (o *SeriesEditorResource) GetSeriesIds() []int32`

GetSeriesIds returns the SeriesIds field if non-nil, zero value otherwise.

### GetSeriesIdsOk

`func (o *SeriesEditorResource) GetSeriesIdsOk() (*[]int32, bool)`

GetSeriesIdsOk returns a tuple with the SeriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesIds

`func (o *SeriesEditorResource) SetSeriesIds(v []int32)`

SetSeriesIds sets SeriesIds field to given value.

### HasSeriesIds

`func (o *SeriesEditorResource) HasSeriesIds() bool`

HasSeriesIds returns a boolean if a field has been set.

### SetSeriesIdsNil

`func (o *SeriesEditorResource) SetSeriesIdsNil(b bool)`

 SetSeriesIdsNil sets the value for SeriesIds to be an explicit nil

### UnsetSeriesIds
`func (o *SeriesEditorResource) UnsetSeriesIds()`

UnsetSeriesIds ensures that no value is present for SeriesIds, not even an explicit nil
### GetMonitored

`func (o *SeriesEditorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *SeriesEditorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *SeriesEditorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *SeriesEditorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *SeriesEditorResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *SeriesEditorResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetQualityProfileId

`func (o *SeriesEditorResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *SeriesEditorResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *SeriesEditorResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *SeriesEditorResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *SeriesEditorResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *SeriesEditorResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
### GetSeriesType

`func (o *SeriesEditorResource) GetSeriesType() SeriesTypes`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *SeriesEditorResource) GetSeriesTypeOk() (*SeriesTypes, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *SeriesEditorResource) SetSeriesType(v SeriesTypes)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *SeriesEditorResource) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetSeasonFolder

`func (o *SeriesEditorResource) GetSeasonFolder() bool`

GetSeasonFolder returns the SeasonFolder field if non-nil, zero value otherwise.

### GetSeasonFolderOk

`func (o *SeriesEditorResource) GetSeasonFolderOk() (*bool, bool)`

GetSeasonFolderOk returns a tuple with the SeasonFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonFolder

`func (o *SeriesEditorResource) SetSeasonFolder(v bool)`

SetSeasonFolder sets SeasonFolder field to given value.

### HasSeasonFolder

`func (o *SeriesEditorResource) HasSeasonFolder() bool`

HasSeasonFolder returns a boolean if a field has been set.

### SetSeasonFolderNil

`func (o *SeriesEditorResource) SetSeasonFolderNil(b bool)`

 SetSeasonFolderNil sets the value for SeasonFolder to be an explicit nil

### UnsetSeasonFolder
`func (o *SeriesEditorResource) UnsetSeasonFolder()`

UnsetSeasonFolder ensures that no value is present for SeasonFolder, not even an explicit nil
### GetRootFolderPath

`func (o *SeriesEditorResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *SeriesEditorResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *SeriesEditorResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *SeriesEditorResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *SeriesEditorResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *SeriesEditorResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetTags

`func (o *SeriesEditorResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SeriesEditorResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SeriesEditorResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SeriesEditorResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SeriesEditorResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SeriesEditorResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *SeriesEditorResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *SeriesEditorResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *SeriesEditorResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *SeriesEditorResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetMoveFiles

`func (o *SeriesEditorResource) GetMoveFiles() bool`

GetMoveFiles returns the MoveFiles field if non-nil, zero value otherwise.

### GetMoveFilesOk

`func (o *SeriesEditorResource) GetMoveFilesOk() (*bool, bool)`

GetMoveFilesOk returns a tuple with the MoveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFiles

`func (o *SeriesEditorResource) SetMoveFiles(v bool)`

SetMoveFiles sets MoveFiles field to given value.

### HasMoveFiles

`func (o *SeriesEditorResource) HasMoveFiles() bool`

HasMoveFiles returns a boolean if a field has been set.

### GetDeleteFiles

`func (o *SeriesEditorResource) GetDeleteFiles() bool`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *SeriesEditorResource) GetDeleteFilesOk() (*bool, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *SeriesEditorResource) SetDeleteFiles(v bool)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *SeriesEditorResource) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetAddImportListExclusion

`func (o *SeriesEditorResource) GetAddImportListExclusion() bool`

GetAddImportListExclusion returns the AddImportListExclusion field if non-nil, zero value otherwise.

### GetAddImportListExclusionOk

`func (o *SeriesEditorResource) GetAddImportListExclusionOk() (*bool, bool)`

GetAddImportListExclusionOk returns a tuple with the AddImportListExclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddImportListExclusion

`func (o *SeriesEditorResource) SetAddImportListExclusion(v bool)`

SetAddImportListExclusion sets AddImportListExclusion field to given value.

### HasAddImportListExclusion

`func (o *SeriesEditorResource) HasAddImportListExclusion() bool`

HasAddImportListExclusion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


