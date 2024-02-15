# UpdateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Installed** | Pointer to **bool** |  | [optional] 
**InstalledOn** | Pointer to **NullableTime** |  | [optional] 
**Installable** | Pointer to **bool** |  | [optional] 
**Latest** | Pointer to **bool** |  | [optional] 
**Changes** | Pointer to [**UpdateChanges**](UpdateChanges.md) |  | [optional] 
**Hash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateResource

`func NewUpdateResource() *UpdateResource`

NewUpdateResource instantiates a new UpdateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResourceWithDefaults

`func NewUpdateResourceWithDefaults() *UpdateResource`

NewUpdateResourceWithDefaults instantiates a new UpdateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateResource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateResource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateResource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateResource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBranch

`func (o *UpdateResource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *UpdateResource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *UpdateResource) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *UpdateResource) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *UpdateResource) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *UpdateResource) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetReleaseDate

`func (o *UpdateResource) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *UpdateResource) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *UpdateResource) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *UpdateResource) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetFileName

`func (o *UpdateResource) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateResource) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateResource) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateResource) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *UpdateResource) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *UpdateResource) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetUrl

`func (o *UpdateResource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateResource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateResource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateResource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpdateResource) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpdateResource) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetInstalled

`func (o *UpdateResource) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *UpdateResource) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *UpdateResource) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *UpdateResource) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetInstalledOn

`func (o *UpdateResource) GetInstalledOn() time.Time`

GetInstalledOn returns the InstalledOn field if non-nil, zero value otherwise.

### GetInstalledOnOk

`func (o *UpdateResource) GetInstalledOnOk() (*time.Time, bool)`

GetInstalledOnOk returns a tuple with the InstalledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledOn

`func (o *UpdateResource) SetInstalledOn(v time.Time)`

SetInstalledOn sets InstalledOn field to given value.

### HasInstalledOn

`func (o *UpdateResource) HasInstalledOn() bool`

HasInstalledOn returns a boolean if a field has been set.

### SetInstalledOnNil

`func (o *UpdateResource) SetInstalledOnNil(b bool)`

 SetInstalledOnNil sets the value for InstalledOn to be an explicit nil

### UnsetInstalledOn
`func (o *UpdateResource) UnsetInstalledOn()`

UnsetInstalledOn ensures that no value is present for InstalledOn, not even an explicit nil
### GetInstallable

`func (o *UpdateResource) GetInstallable() bool`

GetInstallable returns the Installable field if non-nil, zero value otherwise.

### GetInstallableOk

`func (o *UpdateResource) GetInstallableOk() (*bool, bool)`

GetInstallableOk returns a tuple with the Installable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallable

`func (o *UpdateResource) SetInstallable(v bool)`

SetInstallable sets Installable field to given value.

### HasInstallable

`func (o *UpdateResource) HasInstallable() bool`

HasInstallable returns a boolean if a field has been set.

### GetLatest

`func (o *UpdateResource) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *UpdateResource) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *UpdateResource) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *UpdateResource) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetChanges

`func (o *UpdateResource) GetChanges() UpdateChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *UpdateResource) GetChangesOk() (*UpdateChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *UpdateResource) SetChanges(v UpdateChanges)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *UpdateResource) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetHash

`func (o *UpdateResource) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UpdateResource) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UpdateResource) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *UpdateResource) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *UpdateResource) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *UpdateResource) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


