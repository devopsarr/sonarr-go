# DelayProfileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EnableUsenet** | Pointer to **bool** |  | [optional] 
**EnableTorrent** | Pointer to **bool** |  | [optional] 
**PreferredProtocol** | Pointer to [**DownloadProtocol**](DownloadProtocol.md) |  | [optional] 
**UsenetDelay** | Pointer to **int32** |  | [optional] 
**TorrentDelay** | Pointer to **int32** |  | [optional] 
**BypassIfHighestQuality** | Pointer to **bool** |  | [optional] 
**BypassIfAboveCustomFormatScore** | Pointer to **bool** |  | [optional] 
**MinimumCustomFormatScore** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewDelayProfileResource

`func NewDelayProfileResource() *DelayProfileResource`

NewDelayProfileResource instantiates a new DelayProfileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayProfileResourceWithDefaults

`func NewDelayProfileResourceWithDefaults() *DelayProfileResource`

NewDelayProfileResourceWithDefaults instantiates a new DelayProfileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelayProfileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelayProfileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelayProfileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DelayProfileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnableUsenet

`func (o *DelayProfileResource) GetEnableUsenet() bool`

GetEnableUsenet returns the EnableUsenet field if non-nil, zero value otherwise.

### GetEnableUsenetOk

`func (o *DelayProfileResource) GetEnableUsenetOk() (*bool, bool)`

GetEnableUsenetOk returns a tuple with the EnableUsenet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUsenet

`func (o *DelayProfileResource) SetEnableUsenet(v bool)`

SetEnableUsenet sets EnableUsenet field to given value.

### HasEnableUsenet

`func (o *DelayProfileResource) HasEnableUsenet() bool`

HasEnableUsenet returns a boolean if a field has been set.

### GetEnableTorrent

`func (o *DelayProfileResource) GetEnableTorrent() bool`

GetEnableTorrent returns the EnableTorrent field if non-nil, zero value otherwise.

### GetEnableTorrentOk

`func (o *DelayProfileResource) GetEnableTorrentOk() (*bool, bool)`

GetEnableTorrentOk returns a tuple with the EnableTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTorrent

`func (o *DelayProfileResource) SetEnableTorrent(v bool)`

SetEnableTorrent sets EnableTorrent field to given value.

### HasEnableTorrent

`func (o *DelayProfileResource) HasEnableTorrent() bool`

HasEnableTorrent returns a boolean if a field has been set.

### GetPreferredProtocol

`func (o *DelayProfileResource) GetPreferredProtocol() DownloadProtocol`

GetPreferredProtocol returns the PreferredProtocol field if non-nil, zero value otherwise.

### GetPreferredProtocolOk

`func (o *DelayProfileResource) GetPreferredProtocolOk() (*DownloadProtocol, bool)`

GetPreferredProtocolOk returns a tuple with the PreferredProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredProtocol

`func (o *DelayProfileResource) SetPreferredProtocol(v DownloadProtocol)`

SetPreferredProtocol sets PreferredProtocol field to given value.

### HasPreferredProtocol

`func (o *DelayProfileResource) HasPreferredProtocol() bool`

HasPreferredProtocol returns a boolean if a field has been set.

### GetUsenetDelay

`func (o *DelayProfileResource) GetUsenetDelay() int32`

GetUsenetDelay returns the UsenetDelay field if non-nil, zero value otherwise.

### GetUsenetDelayOk

`func (o *DelayProfileResource) GetUsenetDelayOk() (*int32, bool)`

GetUsenetDelayOk returns a tuple with the UsenetDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsenetDelay

`func (o *DelayProfileResource) SetUsenetDelay(v int32)`

SetUsenetDelay sets UsenetDelay field to given value.

### HasUsenetDelay

`func (o *DelayProfileResource) HasUsenetDelay() bool`

HasUsenetDelay returns a boolean if a field has been set.

### GetTorrentDelay

`func (o *DelayProfileResource) GetTorrentDelay() int32`

GetTorrentDelay returns the TorrentDelay field if non-nil, zero value otherwise.

### GetTorrentDelayOk

`func (o *DelayProfileResource) GetTorrentDelayOk() (*int32, bool)`

GetTorrentDelayOk returns a tuple with the TorrentDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentDelay

`func (o *DelayProfileResource) SetTorrentDelay(v int32)`

SetTorrentDelay sets TorrentDelay field to given value.

### HasTorrentDelay

`func (o *DelayProfileResource) HasTorrentDelay() bool`

HasTorrentDelay returns a boolean if a field has been set.

### GetBypassIfHighestQuality

`func (o *DelayProfileResource) GetBypassIfHighestQuality() bool`

GetBypassIfHighestQuality returns the BypassIfHighestQuality field if non-nil, zero value otherwise.

### GetBypassIfHighestQualityOk

`func (o *DelayProfileResource) GetBypassIfHighestQualityOk() (*bool, bool)`

GetBypassIfHighestQualityOk returns a tuple with the BypassIfHighestQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassIfHighestQuality

`func (o *DelayProfileResource) SetBypassIfHighestQuality(v bool)`

SetBypassIfHighestQuality sets BypassIfHighestQuality field to given value.

### HasBypassIfHighestQuality

`func (o *DelayProfileResource) HasBypassIfHighestQuality() bool`

HasBypassIfHighestQuality returns a boolean if a field has been set.

### GetBypassIfAboveCustomFormatScore

`func (o *DelayProfileResource) GetBypassIfAboveCustomFormatScore() bool`

GetBypassIfAboveCustomFormatScore returns the BypassIfAboveCustomFormatScore field if non-nil, zero value otherwise.

### GetBypassIfAboveCustomFormatScoreOk

`func (o *DelayProfileResource) GetBypassIfAboveCustomFormatScoreOk() (*bool, bool)`

GetBypassIfAboveCustomFormatScoreOk returns a tuple with the BypassIfAboveCustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassIfAboveCustomFormatScore

`func (o *DelayProfileResource) SetBypassIfAboveCustomFormatScore(v bool)`

SetBypassIfAboveCustomFormatScore sets BypassIfAboveCustomFormatScore field to given value.

### HasBypassIfAboveCustomFormatScore

`func (o *DelayProfileResource) HasBypassIfAboveCustomFormatScore() bool`

HasBypassIfAboveCustomFormatScore returns a boolean if a field has been set.

### GetMinimumCustomFormatScore

`func (o *DelayProfileResource) GetMinimumCustomFormatScore() int32`

GetMinimumCustomFormatScore returns the MinimumCustomFormatScore field if non-nil, zero value otherwise.

### GetMinimumCustomFormatScoreOk

`func (o *DelayProfileResource) GetMinimumCustomFormatScoreOk() (*int32, bool)`

GetMinimumCustomFormatScoreOk returns a tuple with the MinimumCustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCustomFormatScore

`func (o *DelayProfileResource) SetMinimumCustomFormatScore(v int32)`

SetMinimumCustomFormatScore sets MinimumCustomFormatScore field to given value.

### HasMinimumCustomFormatScore

`func (o *DelayProfileResource) HasMinimumCustomFormatScore() bool`

HasMinimumCustomFormatScore returns a boolean if a field has been set.

### GetOrder

`func (o *DelayProfileResource) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DelayProfileResource) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DelayProfileResource) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DelayProfileResource) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTags

`func (o *DelayProfileResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DelayProfileResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DelayProfileResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DelayProfileResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DelayProfileResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DelayProfileResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


