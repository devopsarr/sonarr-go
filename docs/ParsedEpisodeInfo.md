# ParsedEpisodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseTitle** | Pointer to **NullableString** |  | [optional] 
**SeriesTitle** | Pointer to **NullableString** |  | [optional] 
**SeriesTitleInfo** | Pointer to [**SeriesTitleInfo**](SeriesTitleInfo.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**SeasonNumber** | Pointer to **int32** |  | [optional] 
**EpisodeNumbers** | Pointer to **[]int32** |  | [optional] 
**AbsoluteEpisodeNumbers** | Pointer to **[]int32** |  | [optional] 
**SpecialAbsoluteEpisodeNumbers** | Pointer to **[]float64** |  | [optional] 
**AirDate** | Pointer to **NullableString** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**FullSeason** | Pointer to **bool** |  | [optional] 
**IsPartialSeason** | Pointer to **bool** |  | [optional] 
**IsMultiSeason** | Pointer to **bool** |  | [optional] 
**IsSeasonExtra** | Pointer to **bool** |  | [optional] 
**Special** | Pointer to **bool** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**SeasonPart** | Pointer to **int32** |  | [optional] 
**ReleaseTokens** | Pointer to **NullableString** |  | [optional] 
**DailyPart** | Pointer to **NullableInt32** |  | [optional] 
**IsDaily** | Pointer to **bool** |  | [optional] [readonly] 
**IsAbsoluteNumbering** | Pointer to **bool** |  | [optional] [readonly] 
**IsPossibleSpecialEpisode** | Pointer to **bool** |  | [optional] [readonly] 
**IsPossibleSceneSeasonSpecial** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewParsedEpisodeInfo

`func NewParsedEpisodeInfo() *ParsedEpisodeInfo`

NewParsedEpisodeInfo instantiates a new ParsedEpisodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedEpisodeInfoWithDefaults

`func NewParsedEpisodeInfoWithDefaults() *ParsedEpisodeInfo`

NewParsedEpisodeInfoWithDefaults instantiates a new ParsedEpisodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseTitle

`func (o *ParsedEpisodeInfo) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *ParsedEpisodeInfo) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *ParsedEpisodeInfo) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *ParsedEpisodeInfo) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### SetReleaseTitleNil

`func (o *ParsedEpisodeInfo) SetReleaseTitleNil(b bool)`

 SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil

### UnsetReleaseTitle
`func (o *ParsedEpisodeInfo) UnsetReleaseTitle()`

UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
### GetSeriesTitle

`func (o *ParsedEpisodeInfo) GetSeriesTitle() string`

GetSeriesTitle returns the SeriesTitle field if non-nil, zero value otherwise.

### GetSeriesTitleOk

`func (o *ParsedEpisodeInfo) GetSeriesTitleOk() (*string, bool)`

GetSeriesTitleOk returns a tuple with the SeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitle

`func (o *ParsedEpisodeInfo) SetSeriesTitle(v string)`

SetSeriesTitle sets SeriesTitle field to given value.

### HasSeriesTitle

`func (o *ParsedEpisodeInfo) HasSeriesTitle() bool`

HasSeriesTitle returns a boolean if a field has been set.

### SetSeriesTitleNil

`func (o *ParsedEpisodeInfo) SetSeriesTitleNil(b bool)`

 SetSeriesTitleNil sets the value for SeriesTitle to be an explicit nil

### UnsetSeriesTitle
`func (o *ParsedEpisodeInfo) UnsetSeriesTitle()`

UnsetSeriesTitle ensures that no value is present for SeriesTitle, not even an explicit nil
### GetSeriesTitleInfo

`func (o *ParsedEpisodeInfo) GetSeriesTitleInfo() SeriesTitleInfo`

GetSeriesTitleInfo returns the SeriesTitleInfo field if non-nil, zero value otherwise.

### GetSeriesTitleInfoOk

`func (o *ParsedEpisodeInfo) GetSeriesTitleInfoOk() (*SeriesTitleInfo, bool)`

GetSeriesTitleInfoOk returns a tuple with the SeriesTitleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitleInfo

`func (o *ParsedEpisodeInfo) SetSeriesTitleInfo(v SeriesTitleInfo)`

SetSeriesTitleInfo sets SeriesTitleInfo field to given value.

### HasSeriesTitleInfo

`func (o *ParsedEpisodeInfo) HasSeriesTitleInfo() bool`

HasSeriesTitleInfo returns a boolean if a field has been set.

### GetQuality

`func (o *ParsedEpisodeInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedEpisodeInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedEpisodeInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedEpisodeInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *ParsedEpisodeInfo) GetSeasonNumber() int32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *ParsedEpisodeInfo) GetSeasonNumberOk() (*int32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *ParsedEpisodeInfo) SetSeasonNumber(v int32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *ParsedEpisodeInfo) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetEpisodeNumbers

`func (o *ParsedEpisodeInfo) GetEpisodeNumbers() []int32`

GetEpisodeNumbers returns the EpisodeNumbers field if non-nil, zero value otherwise.

### GetEpisodeNumbersOk

`func (o *ParsedEpisodeInfo) GetEpisodeNumbersOk() (*[]int32, bool)`

GetEpisodeNumbersOk returns a tuple with the EpisodeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeNumbers

`func (o *ParsedEpisodeInfo) SetEpisodeNumbers(v []int32)`

SetEpisodeNumbers sets EpisodeNumbers field to given value.

### HasEpisodeNumbers

`func (o *ParsedEpisodeInfo) HasEpisodeNumbers() bool`

HasEpisodeNumbers returns a boolean if a field has been set.

### SetEpisodeNumbersNil

`func (o *ParsedEpisodeInfo) SetEpisodeNumbersNil(b bool)`

 SetEpisodeNumbersNil sets the value for EpisodeNumbers to be an explicit nil

### UnsetEpisodeNumbers
`func (o *ParsedEpisodeInfo) UnsetEpisodeNumbers()`

UnsetEpisodeNumbers ensures that no value is present for EpisodeNumbers, not even an explicit nil
### GetAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) GetAbsoluteEpisodeNumbers() []int32`

GetAbsoluteEpisodeNumbers returns the AbsoluteEpisodeNumbers field if non-nil, zero value otherwise.

### GetAbsoluteEpisodeNumbersOk

`func (o *ParsedEpisodeInfo) GetAbsoluteEpisodeNumbersOk() (*[]int32, bool)`

GetAbsoluteEpisodeNumbersOk returns a tuple with the AbsoluteEpisodeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) SetAbsoluteEpisodeNumbers(v []int32)`

SetAbsoluteEpisodeNumbers sets AbsoluteEpisodeNumbers field to given value.

### HasAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) HasAbsoluteEpisodeNumbers() bool`

HasAbsoluteEpisodeNumbers returns a boolean if a field has been set.

### SetAbsoluteEpisodeNumbersNil

`func (o *ParsedEpisodeInfo) SetAbsoluteEpisodeNumbersNil(b bool)`

 SetAbsoluteEpisodeNumbersNil sets the value for AbsoluteEpisodeNumbers to be an explicit nil

### UnsetAbsoluteEpisodeNumbers
`func (o *ParsedEpisodeInfo) UnsetAbsoluteEpisodeNumbers()`

UnsetAbsoluteEpisodeNumbers ensures that no value is present for AbsoluteEpisodeNumbers, not even an explicit nil
### GetSpecialAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) GetSpecialAbsoluteEpisodeNumbers() []float64`

GetSpecialAbsoluteEpisodeNumbers returns the SpecialAbsoluteEpisodeNumbers field if non-nil, zero value otherwise.

### GetSpecialAbsoluteEpisodeNumbersOk

`func (o *ParsedEpisodeInfo) GetSpecialAbsoluteEpisodeNumbersOk() (*[]float64, bool)`

GetSpecialAbsoluteEpisodeNumbersOk returns a tuple with the SpecialAbsoluteEpisodeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) SetSpecialAbsoluteEpisodeNumbers(v []float64)`

SetSpecialAbsoluteEpisodeNumbers sets SpecialAbsoluteEpisodeNumbers field to given value.

### HasSpecialAbsoluteEpisodeNumbers

`func (o *ParsedEpisodeInfo) HasSpecialAbsoluteEpisodeNumbers() bool`

HasSpecialAbsoluteEpisodeNumbers returns a boolean if a field has been set.

### SetSpecialAbsoluteEpisodeNumbersNil

`func (o *ParsedEpisodeInfo) SetSpecialAbsoluteEpisodeNumbersNil(b bool)`

 SetSpecialAbsoluteEpisodeNumbersNil sets the value for SpecialAbsoluteEpisodeNumbers to be an explicit nil

### UnsetSpecialAbsoluteEpisodeNumbers
`func (o *ParsedEpisodeInfo) UnsetSpecialAbsoluteEpisodeNumbers()`

UnsetSpecialAbsoluteEpisodeNumbers ensures that no value is present for SpecialAbsoluteEpisodeNumbers, not even an explicit nil
### GetAirDate

`func (o *ParsedEpisodeInfo) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *ParsedEpisodeInfo) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *ParsedEpisodeInfo) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *ParsedEpisodeInfo) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### SetAirDateNil

`func (o *ParsedEpisodeInfo) SetAirDateNil(b bool)`

 SetAirDateNil sets the value for AirDate to be an explicit nil

### UnsetAirDate
`func (o *ParsedEpisodeInfo) UnsetAirDate()`

UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
### GetLanguages

`func (o *ParsedEpisodeInfo) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ParsedEpisodeInfo) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ParsedEpisodeInfo) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ParsedEpisodeInfo) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *ParsedEpisodeInfo) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *ParsedEpisodeInfo) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetFullSeason

`func (o *ParsedEpisodeInfo) GetFullSeason() bool`

GetFullSeason returns the FullSeason field if non-nil, zero value otherwise.

### GetFullSeasonOk

`func (o *ParsedEpisodeInfo) GetFullSeasonOk() (*bool, bool)`

GetFullSeasonOk returns a tuple with the FullSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSeason

`func (o *ParsedEpisodeInfo) SetFullSeason(v bool)`

SetFullSeason sets FullSeason field to given value.

### HasFullSeason

`func (o *ParsedEpisodeInfo) HasFullSeason() bool`

HasFullSeason returns a boolean if a field has been set.

### GetIsPartialSeason

`func (o *ParsedEpisodeInfo) GetIsPartialSeason() bool`

GetIsPartialSeason returns the IsPartialSeason field if non-nil, zero value otherwise.

### GetIsPartialSeasonOk

`func (o *ParsedEpisodeInfo) GetIsPartialSeasonOk() (*bool, bool)`

GetIsPartialSeasonOk returns a tuple with the IsPartialSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialSeason

`func (o *ParsedEpisodeInfo) SetIsPartialSeason(v bool)`

SetIsPartialSeason sets IsPartialSeason field to given value.

### HasIsPartialSeason

`func (o *ParsedEpisodeInfo) HasIsPartialSeason() bool`

HasIsPartialSeason returns a boolean if a field has been set.

### GetIsMultiSeason

`func (o *ParsedEpisodeInfo) GetIsMultiSeason() bool`

GetIsMultiSeason returns the IsMultiSeason field if non-nil, zero value otherwise.

### GetIsMultiSeasonOk

`func (o *ParsedEpisodeInfo) GetIsMultiSeasonOk() (*bool, bool)`

GetIsMultiSeasonOk returns a tuple with the IsMultiSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiSeason

`func (o *ParsedEpisodeInfo) SetIsMultiSeason(v bool)`

SetIsMultiSeason sets IsMultiSeason field to given value.

### HasIsMultiSeason

`func (o *ParsedEpisodeInfo) HasIsMultiSeason() bool`

HasIsMultiSeason returns a boolean if a field has been set.

### GetIsSeasonExtra

`func (o *ParsedEpisodeInfo) GetIsSeasonExtra() bool`

GetIsSeasonExtra returns the IsSeasonExtra field if non-nil, zero value otherwise.

### GetIsSeasonExtraOk

`func (o *ParsedEpisodeInfo) GetIsSeasonExtraOk() (*bool, bool)`

GetIsSeasonExtraOk returns a tuple with the IsSeasonExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeasonExtra

`func (o *ParsedEpisodeInfo) SetIsSeasonExtra(v bool)`

SetIsSeasonExtra sets IsSeasonExtra field to given value.

### HasIsSeasonExtra

`func (o *ParsedEpisodeInfo) HasIsSeasonExtra() bool`

HasIsSeasonExtra returns a boolean if a field has been set.

### GetSpecial

`func (o *ParsedEpisodeInfo) GetSpecial() bool`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *ParsedEpisodeInfo) GetSpecialOk() (*bool, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *ParsedEpisodeInfo) SetSpecial(v bool)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *ParsedEpisodeInfo) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### GetReleaseGroup

`func (o *ParsedEpisodeInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedEpisodeInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedEpisodeInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedEpisodeInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedEpisodeInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedEpisodeInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedEpisodeInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedEpisodeInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedEpisodeInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedEpisodeInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedEpisodeInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedEpisodeInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetSeasonPart

`func (o *ParsedEpisodeInfo) GetSeasonPart() int32`

GetSeasonPart returns the SeasonPart field if non-nil, zero value otherwise.

### GetSeasonPartOk

`func (o *ParsedEpisodeInfo) GetSeasonPartOk() (*int32, bool)`

GetSeasonPartOk returns a tuple with the SeasonPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonPart

`func (o *ParsedEpisodeInfo) SetSeasonPart(v int32)`

SetSeasonPart sets SeasonPart field to given value.

### HasSeasonPart

`func (o *ParsedEpisodeInfo) HasSeasonPart() bool`

HasSeasonPart returns a boolean if a field has been set.

### GetReleaseTokens

`func (o *ParsedEpisodeInfo) GetReleaseTokens() string`

GetReleaseTokens returns the ReleaseTokens field if non-nil, zero value otherwise.

### GetReleaseTokensOk

`func (o *ParsedEpisodeInfo) GetReleaseTokensOk() (*string, bool)`

GetReleaseTokensOk returns a tuple with the ReleaseTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTokens

`func (o *ParsedEpisodeInfo) SetReleaseTokens(v string)`

SetReleaseTokens sets ReleaseTokens field to given value.

### HasReleaseTokens

`func (o *ParsedEpisodeInfo) HasReleaseTokens() bool`

HasReleaseTokens returns a boolean if a field has been set.

### SetReleaseTokensNil

`func (o *ParsedEpisodeInfo) SetReleaseTokensNil(b bool)`

 SetReleaseTokensNil sets the value for ReleaseTokens to be an explicit nil

### UnsetReleaseTokens
`func (o *ParsedEpisodeInfo) UnsetReleaseTokens()`

UnsetReleaseTokens ensures that no value is present for ReleaseTokens, not even an explicit nil
### GetDailyPart

`func (o *ParsedEpisodeInfo) GetDailyPart() int32`

GetDailyPart returns the DailyPart field if non-nil, zero value otherwise.

### GetDailyPartOk

`func (o *ParsedEpisodeInfo) GetDailyPartOk() (*int32, bool)`

GetDailyPartOk returns a tuple with the DailyPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyPart

`func (o *ParsedEpisodeInfo) SetDailyPart(v int32)`

SetDailyPart sets DailyPart field to given value.

### HasDailyPart

`func (o *ParsedEpisodeInfo) HasDailyPart() bool`

HasDailyPart returns a boolean if a field has been set.

### SetDailyPartNil

`func (o *ParsedEpisodeInfo) SetDailyPartNil(b bool)`

 SetDailyPartNil sets the value for DailyPart to be an explicit nil

### UnsetDailyPart
`func (o *ParsedEpisodeInfo) UnsetDailyPart()`

UnsetDailyPart ensures that no value is present for DailyPart, not even an explicit nil
### GetIsDaily

`func (o *ParsedEpisodeInfo) GetIsDaily() bool`

GetIsDaily returns the IsDaily field if non-nil, zero value otherwise.

### GetIsDailyOk

`func (o *ParsedEpisodeInfo) GetIsDailyOk() (*bool, bool)`

GetIsDailyOk returns a tuple with the IsDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDaily

`func (o *ParsedEpisodeInfo) SetIsDaily(v bool)`

SetIsDaily sets IsDaily field to given value.

### HasIsDaily

`func (o *ParsedEpisodeInfo) HasIsDaily() bool`

HasIsDaily returns a boolean if a field has been set.

### GetIsAbsoluteNumbering

`func (o *ParsedEpisodeInfo) GetIsAbsoluteNumbering() bool`

GetIsAbsoluteNumbering returns the IsAbsoluteNumbering field if non-nil, zero value otherwise.

### GetIsAbsoluteNumberingOk

`func (o *ParsedEpisodeInfo) GetIsAbsoluteNumberingOk() (*bool, bool)`

GetIsAbsoluteNumberingOk returns a tuple with the IsAbsoluteNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbsoluteNumbering

`func (o *ParsedEpisodeInfo) SetIsAbsoluteNumbering(v bool)`

SetIsAbsoluteNumbering sets IsAbsoluteNumbering field to given value.

### HasIsAbsoluteNumbering

`func (o *ParsedEpisodeInfo) HasIsAbsoluteNumbering() bool`

HasIsAbsoluteNumbering returns a boolean if a field has been set.

### GetIsPossibleSpecialEpisode

`func (o *ParsedEpisodeInfo) GetIsPossibleSpecialEpisode() bool`

GetIsPossibleSpecialEpisode returns the IsPossibleSpecialEpisode field if non-nil, zero value otherwise.

### GetIsPossibleSpecialEpisodeOk

`func (o *ParsedEpisodeInfo) GetIsPossibleSpecialEpisodeOk() (*bool, bool)`

GetIsPossibleSpecialEpisodeOk returns a tuple with the IsPossibleSpecialEpisode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPossibleSpecialEpisode

`func (o *ParsedEpisodeInfo) SetIsPossibleSpecialEpisode(v bool)`

SetIsPossibleSpecialEpisode sets IsPossibleSpecialEpisode field to given value.

### HasIsPossibleSpecialEpisode

`func (o *ParsedEpisodeInfo) HasIsPossibleSpecialEpisode() bool`

HasIsPossibleSpecialEpisode returns a boolean if a field has been set.

### GetIsPossibleSceneSeasonSpecial

`func (o *ParsedEpisodeInfo) GetIsPossibleSceneSeasonSpecial() bool`

GetIsPossibleSceneSeasonSpecial returns the IsPossibleSceneSeasonSpecial field if non-nil, zero value otherwise.

### GetIsPossibleSceneSeasonSpecialOk

`func (o *ParsedEpisodeInfo) GetIsPossibleSceneSeasonSpecialOk() (*bool, bool)`

GetIsPossibleSceneSeasonSpecialOk returns a tuple with the IsPossibleSceneSeasonSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPossibleSceneSeasonSpecial

`func (o *ParsedEpisodeInfo) SetIsPossibleSceneSeasonSpecial(v bool)`

SetIsPossibleSceneSeasonSpecial sets IsPossibleSceneSeasonSpecial field to given value.

### HasIsPossibleSceneSeasonSpecial

`func (o *ParsedEpisodeInfo) HasIsPossibleSceneSeasonSpecial() bool`

HasIsPossibleSceneSeasonSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


