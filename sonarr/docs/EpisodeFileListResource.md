# EpisodeFileListResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpisodeFileIds** | Pointer to **[]int32** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEpisodeFileListResource

`func NewEpisodeFileListResource() *EpisodeFileListResource`

NewEpisodeFileListResource instantiates a new EpisodeFileListResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeFileListResourceWithDefaults

`func NewEpisodeFileListResourceWithDefaults() *EpisodeFileListResource`

NewEpisodeFileListResourceWithDefaults instantiates a new EpisodeFileListResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpisodeFileIds

`func (o *EpisodeFileListResource) GetEpisodeFileIds() []int32`

GetEpisodeFileIds returns the EpisodeFileIds field if non-nil, zero value otherwise.

### GetEpisodeFileIdsOk

`func (o *EpisodeFileListResource) GetEpisodeFileIdsOk() (*[]int32, bool)`

GetEpisodeFileIdsOk returns a tuple with the EpisodeFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeFileIds

`func (o *EpisodeFileListResource) SetEpisodeFileIds(v []int32)`

SetEpisodeFileIds sets EpisodeFileIds field to given value.

### HasEpisodeFileIds

`func (o *EpisodeFileListResource) HasEpisodeFileIds() bool`

HasEpisodeFileIds returns a boolean if a field has been set.

### SetEpisodeFileIdsNil

`func (o *EpisodeFileListResource) SetEpisodeFileIdsNil(b bool)`

 SetEpisodeFileIdsNil sets the value for EpisodeFileIds to be an explicit nil

### UnsetEpisodeFileIds
`func (o *EpisodeFileListResource) UnsetEpisodeFileIds()`

UnsetEpisodeFileIds ensures that no value is present for EpisodeFileIds, not even an explicit nil
### GetLanguages

`func (o *EpisodeFileListResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *EpisodeFileListResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *EpisodeFileListResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *EpisodeFileListResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *EpisodeFileListResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *EpisodeFileListResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetQuality

`func (o *EpisodeFileListResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *EpisodeFileListResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *EpisodeFileListResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *EpisodeFileListResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetSceneName

`func (o *EpisodeFileListResource) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *EpisodeFileListResource) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *EpisodeFileListResource) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *EpisodeFileListResource) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *EpisodeFileListResource) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *EpisodeFileListResource) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetReleaseGroup

`func (o *EpisodeFileListResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *EpisodeFileListResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *EpisodeFileListResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *EpisodeFileListResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *EpisodeFileListResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *EpisodeFileListResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


