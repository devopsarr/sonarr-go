# MediaCover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverType** | Pointer to [**MediaCoverTypes**](MediaCoverTypes.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**RemoteUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMediaCover

`func NewMediaCover() *MediaCover`

NewMediaCover instantiates a new MediaCover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaCoverWithDefaults

`func NewMediaCoverWithDefaults() *MediaCover`

NewMediaCoverWithDefaults instantiates a new MediaCover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverType

`func (o *MediaCover) GetCoverType() MediaCoverTypes`

GetCoverType returns the CoverType field if non-nil, zero value otherwise.

### GetCoverTypeOk

`func (o *MediaCover) GetCoverTypeOk() (*MediaCoverTypes, bool)`

GetCoverTypeOk returns a tuple with the CoverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverType

`func (o *MediaCover) SetCoverType(v MediaCoverTypes)`

SetCoverType sets CoverType field to given value.

### HasCoverType

`func (o *MediaCover) HasCoverType() bool`

HasCoverType returns a boolean if a field has been set.

### GetUrl

`func (o *MediaCover) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaCover) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaCover) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MediaCover) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MediaCover) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MediaCover) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRemoteUrl

`func (o *MediaCover) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *MediaCover) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *MediaCover) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.

### HasRemoteUrl

`func (o *MediaCover) HasRemoteUrl() bool`

HasRemoteUrl returns a boolean if a field has been set.

### SetRemoteUrlNil

`func (o *MediaCover) SetRemoteUrlNil(b bool)`

 SetRemoteUrlNil sets the value for RemoteUrl to be an explicit nil

### UnsetRemoteUrl
`func (o *MediaCover) UnsetRemoteUrl()`

UnsetRemoteUrl ensures that no value is present for RemoteUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


