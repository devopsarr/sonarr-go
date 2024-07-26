# NamingConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RenameEpisodes** | Pointer to **bool** |  | [optional] 
**ReplaceIllegalCharacters** | Pointer to **bool** |  | [optional] 
**ColonReplacementFormat** | Pointer to **int32** |  | [optional] 
**CustomColonReplacementFormat** | Pointer to **NullableString** |  | [optional] 
**MultiEpisodeStyle** | Pointer to **int32** |  | [optional] 
**StandardEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**DailyEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**AnimeEpisodeFormat** | Pointer to **NullableString** |  | [optional] 
**SeriesFolderFormat** | Pointer to **NullableString** |  | [optional] 
**SeasonFolderFormat** | Pointer to **NullableString** |  | [optional] 
**SpecialsFolderFormat** | Pointer to **NullableString** |  | [optional] 

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

### GetColonReplacementFormat

`func (o *NamingConfigResource) GetColonReplacementFormat() int32`

GetColonReplacementFormat returns the ColonReplacementFormat field if non-nil, zero value otherwise.

### GetColonReplacementFormatOk

`func (o *NamingConfigResource) GetColonReplacementFormatOk() (*int32, bool)`

GetColonReplacementFormatOk returns a tuple with the ColonReplacementFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColonReplacementFormat

`func (o *NamingConfigResource) SetColonReplacementFormat(v int32)`

SetColonReplacementFormat sets ColonReplacementFormat field to given value.

### HasColonReplacementFormat

`func (o *NamingConfigResource) HasColonReplacementFormat() bool`

HasColonReplacementFormat returns a boolean if a field has been set.

### GetCustomColonReplacementFormat

`func (o *NamingConfigResource) GetCustomColonReplacementFormat() string`

GetCustomColonReplacementFormat returns the CustomColonReplacementFormat field if non-nil, zero value otherwise.

### GetCustomColonReplacementFormatOk

`func (o *NamingConfigResource) GetCustomColonReplacementFormatOk() (*string, bool)`

GetCustomColonReplacementFormatOk returns a tuple with the CustomColonReplacementFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomColonReplacementFormat

`func (o *NamingConfigResource) SetCustomColonReplacementFormat(v string)`

SetCustomColonReplacementFormat sets CustomColonReplacementFormat field to given value.

### HasCustomColonReplacementFormat

`func (o *NamingConfigResource) HasCustomColonReplacementFormat() bool`

HasCustomColonReplacementFormat returns a boolean if a field has been set.

### SetCustomColonReplacementFormatNil

`func (o *NamingConfigResource) SetCustomColonReplacementFormatNil(b bool)`

 SetCustomColonReplacementFormatNil sets the value for CustomColonReplacementFormat to be an explicit nil

### UnsetCustomColonReplacementFormat
`func (o *NamingConfigResource) UnsetCustomColonReplacementFormat()`

UnsetCustomColonReplacementFormat ensures that no value is present for CustomColonReplacementFormat, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


