# ParseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ParsedEpisodeInfo** | Pointer to [**ParsedEpisodeInfo**](ParsedEpisodeInfo.md) |  | [optional] 
**Series** | Pointer to [**SeriesResource**](SeriesResource.md) |  | [optional] 
**Episodes** | Pointer to [**[]EpisodeResource**](EpisodeResource.md) |  | [optional] 

## Methods

### NewParseResource

`func NewParseResource() *ParseResource`

NewParseResource instantiates a new ParseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseResourceWithDefaults

`func NewParseResourceWithDefaults() *ParseResource`

NewParseResourceWithDefaults instantiates a new ParseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ParseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetParsedEpisodeInfo

`func (o *ParseResource) GetParsedEpisodeInfo() ParsedEpisodeInfo`

GetParsedEpisodeInfo returns the ParsedEpisodeInfo field if non-nil, zero value otherwise.

### GetParsedEpisodeInfoOk

`func (o *ParseResource) GetParsedEpisodeInfoOk() (*ParsedEpisodeInfo, bool)`

GetParsedEpisodeInfoOk returns a tuple with the ParsedEpisodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedEpisodeInfo

`func (o *ParseResource) SetParsedEpisodeInfo(v ParsedEpisodeInfo)`

SetParsedEpisodeInfo sets ParsedEpisodeInfo field to given value.

### HasParsedEpisodeInfo

`func (o *ParseResource) HasParsedEpisodeInfo() bool`

HasParsedEpisodeInfo returns a boolean if a field has been set.

### GetSeries

`func (o *ParseResource) GetSeries() SeriesResource`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ParseResource) GetSeriesOk() (*SeriesResource, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ParseResource) SetSeries(v SeriesResource)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ParseResource) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetEpisodes

`func (o *ParseResource) GetEpisodes() []EpisodeResource`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *ParseResource) GetEpisodesOk() (*[]EpisodeResource, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *ParseResource) SetEpisodes(v []EpisodeResource)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *ParseResource) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.

### SetEpisodesNil

`func (o *ParseResource) SetEpisodesNil(b bool)`

 SetEpisodesNil sets the value for Episodes to be an explicit nil

### UnsetEpisodes
`func (o *ParseResource) UnsetEpisodes()`

UnsetEpisodes ensures that no value is present for Episodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


