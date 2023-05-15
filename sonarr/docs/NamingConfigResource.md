# NamingConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RenameEpisodes** | Pointer to **bool** |  | [optional] 
**ReplaceIllegalCharacters** | Pointer to **bool** |  | [optional] 
**MultiEpisodeStyle** | Pointer to **int32** |  | [optional] 
**StandardEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**DailyEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**AnimeEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**SeriesFolderFormat** | Pointer to **NullableString** |  | [optional] 
**SeasonFolderFormat** | Pointer to **NullableString** |  | [optional] 
**SpecialsFolderFormat** | Pointer to **NullableString** |  | [optional] 
**IncludeSeriesTitle** | Pointer to **bool** |  | [optional] 
**IncludeEpisodeTitle** | Pointer to **bool** |  | [optional] 
**IncludeQuality** | Pointer to **bool** |  | [optional] 
**ReplaceSpaces** | Pointer to **bool** |  | [optional] 
**Separator** | Pointer to **NullableString** |  | [optional] 
**NumberStyle** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNamingConfigResource

`func NewNamingConfigResource() *NamingConfigResource`

NewNamingConfigResource instantiates a new NamingConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingConfigResourceWithDefaults

`func NewNamingConfigResourceWithDefaults() *NamingConfigResource`

NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamingConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRenameEpisodes

`func (o *NamingConfigResource) GetRenameEpisodes() bool`

GetRenameEpisodes returns the RenameEpisodes field if non-nil, zero value otherwise.

### GetRenameEpisodesOk

`func (o *NamingConfigResource) GetRenameEpisodesOk() (*bool, bool)`

GetRenameEpisodesOk returns a tuple with the RenameEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameEpisodes

`func (o *NamingConfigResource) SetRenameEpisodes(v bool)`

SetRenameEpisodes sets RenameEpisodes field to given value.

### HasRenameEpisodes

`func (o *NamingConfigResource) HasRenameEpisodes() bool`

HasRenameEpisodes returns a boolean if a field has been set.

### GetReplaceIllegalCharacters

`func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool`

GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field if non-nil, zero value otherwise.

### GetReplaceIllegalCharactersOk

`func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool)`

GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceIllegalCharacters

`func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool)`

SetReplaceIllegalCharacters sets ReplaceIllegalCharacters field to given value.

### HasReplaceIllegalCharacters

`func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool`

HasReplaceIllegalCharacters returns a boolean if a field has been set.

### GetMultiEpisodeStyle

`func (o *NamingConfigResource) GetMultiEpisodeStyle() int32`

GetMultiEpisodeStyle returns the MultiEpisodeStyle field if non-nil, zero value otherwise.

### GetMultiEpisodeStyleOk

`func (o *NamingConfigResource) GetMultiEpisodeStyleOk() (*int32, bool)`

GetMultiEpisodeStyleOk returns a tuple with the MultiEpisodeStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiEpisodeStyle

`func (o *NamingConfigResource) SetMultiEpisodeStyle(v int32)`

SetMultiEpisodeStyle sets MultiEpisodeStyle field to given value.

### HasMultiEpisodeStyle

`func (o *NamingConfigResource) HasMultiEpisodeStyle() bool`

HasMultiEpisodeStyle returns a boolean if a field has been set.

### GetStandardEpisodeFormat

`func (o *NamingConfigResource) GetStandardEpisodeFormat() string`

GetStandardEpisodeFormat returns the StandardEpisodeFormat field if non-nil, zero value otherwise.

### GetStandardEpisodeFormatOk

`func (o *NamingConfigResource) GetStandardEpisodeFormatOk() (*string, bool)`

GetStandardEpisodeFormatOk returns a tuple with the StandardEpisodeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardEpisodeFormat

`func (o *NamingConfigResource) SetStandardEpisodeFormat(v string)`

SetStandardEpisodeFormat sets StandardEpisodeFormat field to given value.

### HasStandardEpisodeFormat

`func (o *NamingConfigResource) HasStandardEpisodeFormat() bool`

HasStandardEpisodeFormat returns a boolean if a field has been set.

### SetStandardEpisodeFormatNil

`func (o *NamingConfigResource) SetStandardEpisodeFormatNil(b bool)`

 SetStandardEpisodeFormatNil sets the value for StandardEpisodeFormat to be an explicit nil

### UnsetStandardEpisodeFormat
`func (o *NamingConfigResource) UnsetStandardEpisodeFormat()`

UnsetStandardEpisodeFormat ensures that no value is present for StandardEpisodeFormat, not even an explicit nil
### GetDailyEpisodeFormat

`func (o *NamingConfigResource) GetDailyEpisodeFormat() string`

GetDailyEpisodeFormat returns the DailyEpisodeFormat field if non-nil, zero value otherwise.

### GetDailyEpisodeFormatOk

`func (o *NamingConfigResource) GetDailyEpisodeFormatOk() (*string, bool)`

GetDailyEpisodeFormatOk returns a tuple with the DailyEpisodeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyEpisodeFormat

`func (o *NamingConfigResource) SetDailyEpisodeFormat(v string)`

SetDailyEpisodeFormat sets DailyEpisodeFormat field to given value.

### HasDailyEpisodeFormat

`func (o *NamingConfigResource) HasDailyEpisodeFormat() bool`

HasDailyEpisodeFormat returns a boolean if a field has been set.

### SetDailyEpisodeFormatNil

`func (o *NamingConfigResource) SetDailyEpisodeFormatNil(b bool)`

 SetDailyEpisodeFormatNil sets the value for DailyEpisodeFormat to be an explicit nil

### UnsetDailyEpisodeFormat
`func (o *NamingConfigResource) UnsetDailyEpisodeFormat()`

UnsetDailyEpisodeFormat ensures that no value is present for DailyEpisodeFormat, not even an explicit nil
### GetAnimeEpisodeFormat

`func (o *NamingConfigResource) GetAnimeEpisodeFormat() string`

GetAnimeEpisodeFormat returns the AnimeEpisodeFormat field if non-nil, zero value otherwise.

### GetAnimeEpisodeFormatOk

`func (o *NamingConfigResource) GetAnimeEpisodeFormatOk() (*string, bool)`

GetAnimeEpisodeFormatOk returns a tuple with the AnimeEpisodeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimeEpisodeFormat

`func (o *NamingConfigResource) SetAnimeEpisodeFormat(v string)`

SetAnimeEpisodeFormat sets AnimeEpisodeFormat field to given value.

### HasAnimeEpisodeFormat

`func (o *NamingConfigResource) HasAnimeEpisodeFormat() bool`

HasAnimeEpisodeFormat returns a boolean if a field has been set.

### SetAnimeEpisodeFormatNil

`func (o *NamingConfigResource) SetAnimeEpisodeFormatNil(b bool)`

 SetAnimeEpisodeFormatNil sets the value for AnimeEpisodeFormat to be an explicit nil

### UnsetAnimeEpisodeFormat
`func (o *NamingConfigResource) UnsetAnimeEpisodeFormat()`

UnsetAnimeEpisodeFormat ensures that no value is present for AnimeEpisodeFormat, not even an explicit nil
### GetSeriesFolderFormat

`func (o *NamingConfigResource) GetSeriesFolderFormat() string`

GetSeriesFolderFormat returns the SeriesFolderFormat field if non-nil, zero value otherwise.

### GetSeriesFolderFormatOk

`func (o *NamingConfigResource) GetSeriesFolderFormatOk() (*string, bool)`

GetSeriesFolderFormatOk returns a tuple with the SeriesFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesFolderFormat

`func (o *NamingConfigResource) SetSeriesFolderFormat(v string)`

SetSeriesFolderFormat sets SeriesFolderFormat field to given value.

### HasSeriesFolderFormat

`func (o *NamingConfigResource) HasSeriesFolderFormat() bool`

HasSeriesFolderFormat returns a boolean if a field has been set.

### SetSeriesFolderFormatNil

`func (o *NamingConfigResource) SetSeriesFolderFormatNil(b bool)`

 SetSeriesFolderFormatNil sets the value for SeriesFolderFormat to be an explicit nil

### UnsetSeriesFolderFormat
`func (o *NamingConfigResource) UnsetSeriesFolderFormat()`

UnsetSeriesFolderFormat ensures that no value is present for SeriesFolderFormat, not even an explicit nil
### GetSeasonFolderFormat

`func (o *NamingConfigResource) GetSeasonFolderFormat() string`

GetSeasonFolderFormat returns the SeasonFolderFormat field if non-nil, zero value otherwise.

### GetSeasonFolderFormatOk

`func (o *NamingConfigResource) GetSeasonFolderFormatOk() (*string, bool)`

GetSeasonFolderFormatOk returns a tuple with the SeasonFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonFolderFormat

`func (o *NamingConfigResource) SetSeasonFolderFormat(v string)`

SetSeasonFolderFormat sets SeasonFolderFormat field to given value.

### HasSeasonFolderFormat

`func (o *NamingConfigResource) HasSeasonFolderFormat() bool`

HasSeasonFolderFormat returns a boolean if a field has been set.

### SetSeasonFolderFormatNil

`func (o *NamingConfigResource) SetSeasonFolderFormatNil(b bool)`

 SetSeasonFolderFormatNil sets the value for SeasonFolderFormat to be an explicit nil

### UnsetSeasonFolderFormat
`func (o *NamingConfigResource) UnsetSeasonFolderFormat()`

UnsetSeasonFolderFormat ensures that no value is present for SeasonFolderFormat, not even an explicit nil
### GetSpecialsFolderFormat

`func (o *NamingConfigResource) GetSpecialsFolderFormat() string`

GetSpecialsFolderFormat returns the SpecialsFolderFormat field if non-nil, zero value otherwise.

### GetSpecialsFolderFormatOk

`func (o *NamingConfigResource) GetSpecialsFolderFormatOk() (*string, bool)`

GetSpecialsFolderFormatOk returns a tuple with the SpecialsFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialsFolderFormat

`func (o *NamingConfigResource) SetSpecialsFolderFormat(v string)`

SetSpecialsFolderFormat sets SpecialsFolderFormat field to given value.

### HasSpecialsFolderFormat

`func (o *NamingConfigResource) HasSpecialsFolderFormat() bool`

HasSpecialsFolderFormat returns a boolean if a field has been set.

### SetSpecialsFolderFormatNil

`func (o *NamingConfigResource) SetSpecialsFolderFormatNil(b bool)`

 SetSpecialsFolderFormatNil sets the value for SpecialsFolderFormat to be an explicit nil

### UnsetSpecialsFolderFormat
`func (o *NamingConfigResource) UnsetSpecialsFolderFormat()`

UnsetSpecialsFolderFormat ensures that no value is present for SpecialsFolderFormat, not even an explicit nil
### GetIncludeSeriesTitle

`func (o *NamingConfigResource) GetIncludeSeriesTitle() bool`

GetIncludeSeriesTitle returns the IncludeSeriesTitle field if non-nil, zero value otherwise.

### GetIncludeSeriesTitleOk

`func (o *NamingConfigResource) GetIncludeSeriesTitleOk() (*bool, bool)`

GetIncludeSeriesTitleOk returns a tuple with the IncludeSeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSeriesTitle

`func (o *NamingConfigResource) SetIncludeSeriesTitle(v bool)`

SetIncludeSeriesTitle sets IncludeSeriesTitle field to given value.

### HasIncludeSeriesTitle

`func (o *NamingConfigResource) HasIncludeSeriesTitle() bool`

HasIncludeSeriesTitle returns a boolean if a field has been set.

### GetIncludeEpisodeTitle

`func (o *NamingConfigResource) GetIncludeEpisodeTitle() bool`

GetIncludeEpisodeTitle returns the IncludeEpisodeTitle field if non-nil, zero value otherwise.

### GetIncludeEpisodeTitleOk

`func (o *NamingConfigResource) GetIncludeEpisodeTitleOk() (*bool, bool)`

GetIncludeEpisodeTitleOk returns a tuple with the IncludeEpisodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEpisodeTitle

`func (o *NamingConfigResource) SetIncludeEpisodeTitle(v bool)`

SetIncludeEpisodeTitle sets IncludeEpisodeTitle field to given value.

### HasIncludeEpisodeTitle

`func (o *NamingConfigResource) HasIncludeEpisodeTitle() bool`

HasIncludeEpisodeTitle returns a boolean if a field has been set.

### GetIncludeQuality

`func (o *NamingConfigResource) GetIncludeQuality() bool`

GetIncludeQuality returns the IncludeQuality field if non-nil, zero value otherwise.

### GetIncludeQualityOk

`func (o *NamingConfigResource) GetIncludeQualityOk() (*bool, bool)`

GetIncludeQualityOk returns a tuple with the IncludeQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQuality

`func (o *NamingConfigResource) SetIncludeQuality(v bool)`

SetIncludeQuality sets IncludeQuality field to given value.

### HasIncludeQuality

`func (o *NamingConfigResource) HasIncludeQuality() bool`

HasIncludeQuality returns a boolean if a field has been set.

### GetReplaceSpaces

`func (o *NamingConfigResource) GetReplaceSpaces() bool`

GetReplaceSpaces returns the ReplaceSpaces field if non-nil, zero value otherwise.

### GetReplaceSpacesOk

`func (o *NamingConfigResource) GetReplaceSpacesOk() (*bool, bool)`

GetReplaceSpacesOk returns a tuple with the ReplaceSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceSpaces

`func (o *NamingConfigResource) SetReplaceSpaces(v bool)`

SetReplaceSpaces sets ReplaceSpaces field to given value.

### HasReplaceSpaces

`func (o *NamingConfigResource) HasReplaceSpaces() bool`

HasReplaceSpaces returns a boolean if a field has been set.

### GetSeparator

`func (o *NamingConfigResource) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *NamingConfigResource) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *NamingConfigResource) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *NamingConfigResource) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### SetSeparatorNil

`func (o *NamingConfigResource) SetSeparatorNil(b bool)`

 SetSeparatorNil sets the value for Separator to be an explicit nil

### UnsetSeparator
`func (o *NamingConfigResource) UnsetSeparator()`

UnsetSeparator ensures that no value is present for Separator, not even an explicit nil
### GetNumberStyle

`func (o *NamingConfigResource) GetNumberStyle() string`

GetNumberStyle returns the NumberStyle field if non-nil, zero value otherwise.

### GetNumberStyleOk

`func (o *NamingConfigResource) GetNumberStyleOk() (*string, bool)`

GetNumberStyleOk returns a tuple with the NumberStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberStyle

`func (o *NamingConfigResource) SetNumberStyle(v string)`

SetNumberStyle sets NumberStyle field to given value.

### HasNumberStyle

`func (o *NamingConfigResource) HasNumberStyle() bool`

HasNumberStyle returns a boolean if a field has been set.

### SetNumberStyleNil

`func (o *NamingConfigResource) SetNumberStyleNil(b bool)`

 SetNumberStyleNil sets the value for NumberStyle to be an explicit nil

### UnsetNumberStyle
`func (o *NamingConfigResource) UnsetNumberStyle()`

UnsetNumberStyle ensures that no value is present for NumberStyle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


