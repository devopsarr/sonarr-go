# MediaInfoResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AudioBitrate** | Pointer to **int64** |  | [optional] 
**AudioChannels** | Pointer to **float64** |  | [optional] 
**AudioCodec** | Pointer to **NullableString** |  | [optional] 
**AudioLanguages** | Pointer to **NullableString** |  | [optional] 
**AudioStreamCount** | Pointer to **int32** |  | [optional] 
**VideoBitDepth** | Pointer to **int32** |  | [optional] 
**VideoBitrate** | Pointer to **int64** |  | [optional] 
**VideoCodec** | Pointer to **NullableString** |  | [optional] 
**VideoFps** | Pointer to **float64** |  | [optional] 
**VideoDynamicRange** | Pointer to **NullableString** |  | [optional] 
**VideoDynamicRangeType** | Pointer to **NullableString** |  | [optional] 
**Resolution** | Pointer to **NullableString** |  | [optional] 
**RunTime** | Pointer to **NullableString** |  | [optional] 
**ScanType** | Pointer to **NullableString** |  | [optional] 
**Subtitles** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMediaInfoResource

`func NewMediaInfoResource() *MediaInfoResource`

NewMediaInfoResource instantiates a new MediaInfoResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaInfoResourceWithDefaults

`func NewMediaInfoResourceWithDefaults() *MediaInfoResource`

NewMediaInfoResourceWithDefaults instantiates a new MediaInfoResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaInfoResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaInfoResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaInfoResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MediaInfoResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAudioBitrate

`func (o *MediaInfoResource) GetAudioBitrate() int64`

GetAudioBitrate returns the AudioBitrate field if non-nil, zero value otherwise.

### GetAudioBitrateOk

`func (o *MediaInfoResource) GetAudioBitrateOk() (*int64, bool)`

GetAudioBitrateOk returns a tuple with the AudioBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitrate

`func (o *MediaInfoResource) SetAudioBitrate(v int64)`

SetAudioBitrate sets AudioBitrate field to given value.

### HasAudioBitrate

`func (o *MediaInfoResource) HasAudioBitrate() bool`

HasAudioBitrate returns a boolean if a field has been set.

### GetAudioChannels

`func (o *MediaInfoResource) GetAudioChannels() float64`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *MediaInfoResource) GetAudioChannelsOk() (*float64, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *MediaInfoResource) SetAudioChannels(v float64)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *MediaInfoResource) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### GetAudioCodec

`func (o *MediaInfoResource) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *MediaInfoResource) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *MediaInfoResource) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *MediaInfoResource) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *MediaInfoResource) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *MediaInfoResource) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetAudioLanguages

`func (o *MediaInfoResource) GetAudioLanguages() string`

GetAudioLanguages returns the AudioLanguages field if non-nil, zero value otherwise.

### GetAudioLanguagesOk

`func (o *MediaInfoResource) GetAudioLanguagesOk() (*string, bool)`

GetAudioLanguagesOk returns a tuple with the AudioLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguages

`func (o *MediaInfoResource) SetAudioLanguages(v string)`

SetAudioLanguages sets AudioLanguages field to given value.

### HasAudioLanguages

`func (o *MediaInfoResource) HasAudioLanguages() bool`

HasAudioLanguages returns a boolean if a field has been set.

### SetAudioLanguagesNil

`func (o *MediaInfoResource) SetAudioLanguagesNil(b bool)`

 SetAudioLanguagesNil sets the value for AudioLanguages to be an explicit nil

### UnsetAudioLanguages
`func (o *MediaInfoResource) UnsetAudioLanguages()`

UnsetAudioLanguages ensures that no value is present for AudioLanguages, not even an explicit nil
### GetAudioStreamCount

`func (o *MediaInfoResource) GetAudioStreamCount() int32`

GetAudioStreamCount returns the AudioStreamCount field if non-nil, zero value otherwise.

### GetAudioStreamCountOk

`func (o *MediaInfoResource) GetAudioStreamCountOk() (*int32, bool)`

GetAudioStreamCountOk returns a tuple with the AudioStreamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamCount

`func (o *MediaInfoResource) SetAudioStreamCount(v int32)`

SetAudioStreamCount sets AudioStreamCount field to given value.

### HasAudioStreamCount

`func (o *MediaInfoResource) HasAudioStreamCount() bool`

HasAudioStreamCount returns a boolean if a field has been set.

### GetVideoBitDepth

`func (o *MediaInfoResource) GetVideoBitDepth() int32`

GetVideoBitDepth returns the VideoBitDepth field if non-nil, zero value otherwise.

### GetVideoBitDepthOk

`func (o *MediaInfoResource) GetVideoBitDepthOk() (*int32, bool)`

GetVideoBitDepthOk returns a tuple with the VideoBitDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitDepth

`func (o *MediaInfoResource) SetVideoBitDepth(v int32)`

SetVideoBitDepth sets VideoBitDepth field to given value.

### HasVideoBitDepth

`func (o *MediaInfoResource) HasVideoBitDepth() bool`

HasVideoBitDepth returns a boolean if a field has been set.

### GetVideoBitrate

`func (o *MediaInfoResource) GetVideoBitrate() int64`

GetVideoBitrate returns the VideoBitrate field if non-nil, zero value otherwise.

### GetVideoBitrateOk

`func (o *MediaInfoResource) GetVideoBitrateOk() (*int64, bool)`

GetVideoBitrateOk returns a tuple with the VideoBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitrate

`func (o *MediaInfoResource) SetVideoBitrate(v int64)`

SetVideoBitrate sets VideoBitrate field to given value.

### HasVideoBitrate

`func (o *MediaInfoResource) HasVideoBitrate() bool`

HasVideoBitrate returns a boolean if a field has been set.

### GetVideoCodec

`func (o *MediaInfoResource) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *MediaInfoResource) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *MediaInfoResource) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *MediaInfoResource) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *MediaInfoResource) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *MediaInfoResource) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetVideoFps

`func (o *MediaInfoResource) GetVideoFps() float64`

GetVideoFps returns the VideoFps field if non-nil, zero value otherwise.

### GetVideoFpsOk

`func (o *MediaInfoResource) GetVideoFpsOk() (*float64, bool)`

GetVideoFpsOk returns a tuple with the VideoFps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFps

`func (o *MediaInfoResource) SetVideoFps(v float64)`

SetVideoFps sets VideoFps field to given value.

### HasVideoFps

`func (o *MediaInfoResource) HasVideoFps() bool`

HasVideoFps returns a boolean if a field has been set.

### GetVideoDynamicRange

`func (o *MediaInfoResource) GetVideoDynamicRange() string`

GetVideoDynamicRange returns the VideoDynamicRange field if non-nil, zero value otherwise.

### GetVideoDynamicRangeOk

`func (o *MediaInfoResource) GetVideoDynamicRangeOk() (*string, bool)`

GetVideoDynamicRangeOk returns a tuple with the VideoDynamicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDynamicRange

`func (o *MediaInfoResource) SetVideoDynamicRange(v string)`

SetVideoDynamicRange sets VideoDynamicRange field to given value.

### HasVideoDynamicRange

`func (o *MediaInfoResource) HasVideoDynamicRange() bool`

HasVideoDynamicRange returns a boolean if a field has been set.

### SetVideoDynamicRangeNil

`func (o *MediaInfoResource) SetVideoDynamicRangeNil(b bool)`

 SetVideoDynamicRangeNil sets the value for VideoDynamicRange to be an explicit nil

### UnsetVideoDynamicRange
`func (o *MediaInfoResource) UnsetVideoDynamicRange()`

UnsetVideoDynamicRange ensures that no value is present for VideoDynamicRange, not even an explicit nil
### GetVideoDynamicRangeType

`func (o *MediaInfoResource) GetVideoDynamicRangeType() string`

GetVideoDynamicRangeType returns the VideoDynamicRangeType field if non-nil, zero value otherwise.

### GetVideoDynamicRangeTypeOk

`func (o *MediaInfoResource) GetVideoDynamicRangeTypeOk() (*string, bool)`

GetVideoDynamicRangeTypeOk returns a tuple with the VideoDynamicRangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDynamicRangeType

`func (o *MediaInfoResource) SetVideoDynamicRangeType(v string)`

SetVideoDynamicRangeType sets VideoDynamicRangeType field to given value.

### HasVideoDynamicRangeType

`func (o *MediaInfoResource) HasVideoDynamicRangeType() bool`

HasVideoDynamicRangeType returns a boolean if a field has been set.

### SetVideoDynamicRangeTypeNil

`func (o *MediaInfoResource) SetVideoDynamicRangeTypeNil(b bool)`

 SetVideoDynamicRangeTypeNil sets the value for VideoDynamicRangeType to be an explicit nil

### UnsetVideoDynamicRangeType
`func (o *MediaInfoResource) UnsetVideoDynamicRangeType()`

UnsetVideoDynamicRangeType ensures that no value is present for VideoDynamicRangeType, not even an explicit nil
### GetResolution

`func (o *MediaInfoResource) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *MediaInfoResource) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *MediaInfoResource) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *MediaInfoResource) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *MediaInfoResource) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *MediaInfoResource) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetRunTime

`func (o *MediaInfoResource) GetRunTime() string`

GetRunTime returns the RunTime field if non-nil, zero value otherwise.

### GetRunTimeOk

`func (o *MediaInfoResource) GetRunTimeOk() (*string, bool)`

GetRunTimeOk returns a tuple with the RunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTime

`func (o *MediaInfoResource) SetRunTime(v string)`

SetRunTime sets RunTime field to given value.

### HasRunTime

`func (o *MediaInfoResource) HasRunTime() bool`

HasRunTime returns a boolean if a field has been set.

### SetRunTimeNil

`func (o *MediaInfoResource) SetRunTimeNil(b bool)`

 SetRunTimeNil sets the value for RunTime to be an explicit nil

### UnsetRunTime
`func (o *MediaInfoResource) UnsetRunTime()`

UnsetRunTime ensures that no value is present for RunTime, not even an explicit nil
### GetScanType

`func (o *MediaInfoResource) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *MediaInfoResource) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *MediaInfoResource) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *MediaInfoResource) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanTypeNil

`func (o *MediaInfoResource) SetScanTypeNil(b bool)`

 SetScanTypeNil sets the value for ScanType to be an explicit nil

### UnsetScanType
`func (o *MediaInfoResource) UnsetScanType()`

UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
### GetSubtitles

`func (o *MediaInfoResource) GetSubtitles() string`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *MediaInfoResource) GetSubtitlesOk() (*string, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *MediaInfoResource) SetSubtitles(v string)`

SetSubtitles sets Subtitles field to given value.

### HasSubtitles

`func (o *MediaInfoResource) HasSubtitles() bool`

HasSubtitles returns a boolean if a field has been set.

### SetSubtitlesNil

`func (o *MediaInfoResource) SetSubtitlesNil(b bool)`

 SetSubtitlesNil sets the value for Subtitles to be an explicit nil

### UnsetSubtitles
`func (o *MediaInfoResource) UnsetSubtitles()`

UnsetSubtitles ensures that no value is present for Subtitles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


