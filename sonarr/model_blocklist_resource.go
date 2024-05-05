/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.4.1491
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"time"
)

// checks if the BlocklistResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlocklistResource{}

// BlocklistResource struct for BlocklistResource
type BlocklistResource struct {
	Id *int32 `json:"id,omitempty"`
	SeriesId *int32 `json:"seriesId,omitempty"`
	EpisodeIds []int32 `json:"episodeIds,omitempty"`
	SourceTitle NullableString `json:"sourceTitle,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	Protocol *DownloadProtocol `json:"protocol,omitempty"`
	Indexer NullableString `json:"indexer,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Series *SeriesResource `json:"series,omitempty"`
}

// NewBlocklistResource instantiates a new BlocklistResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlocklistResource() *BlocklistResource {
	this := BlocklistResource{}
	return &this
}

// NewBlocklistResourceWithDefaults instantiates a new BlocklistResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlocklistResourceWithDefaults() *BlocklistResource {
	this := BlocklistResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlocklistResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlocklistResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BlocklistResource) SetId(v int32) {
	o.Id = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *BlocklistResource) GetSeriesId() int32 {
	if o == nil || IsNil(o.SeriesId) {
		var ret int32
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetSeriesIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *BlocklistResource) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int32 and assigns it to the SeriesId field.
func (o *BlocklistResource) SetSeriesId(v int32) {
	o.SeriesId = &v
}

// GetEpisodeIds returns the EpisodeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetEpisodeIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.EpisodeIds
}

// GetEpisodeIdsOk returns a tuple with the EpisodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetEpisodeIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.EpisodeIds) {
		return nil, false
	}
	return o.EpisodeIds, true
}

// HasEpisodeIds returns a boolean if a field has been set.
func (o *BlocklistResource) HasEpisodeIds() bool {
	if o != nil && !IsNil(o.EpisodeIds) {
		return true
	}

	return false
}

// SetEpisodeIds gets a reference to the given []int32 and assigns it to the EpisodeIds field.
func (o *BlocklistResource) SetEpisodeIds(v []int32) {
	o.EpisodeIds = v
}

// GetSourceTitle returns the SourceTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetSourceTitle() string {
	if o == nil || IsNil(o.SourceTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SourceTitle.Get()
}

// GetSourceTitleOk returns a tuple with the SourceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetSourceTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceTitle.Get(), o.SourceTitle.IsSet()
}

// HasSourceTitle returns a boolean if a field has been set.
func (o *BlocklistResource) HasSourceTitle() bool {
	if o != nil && o.SourceTitle.IsSet() {
		return true
	}

	return false
}

// SetSourceTitle gets a reference to the given NullableString and assigns it to the SourceTitle field.
func (o *BlocklistResource) SetSourceTitle(v string) {
	o.SourceTitle.Set(&v)
}
// SetSourceTitleNil sets the value for SourceTitle to be an explicit nil
func (o *BlocklistResource) SetSourceTitleNil() {
	o.SourceTitle.Set(nil)
}

// UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
func (o *BlocklistResource) UnsetSourceTitle() {
	o.SourceTitle.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *BlocklistResource) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *BlocklistResource) SetLanguages(v []Language) {
	o.Languages = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *BlocklistResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *BlocklistResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *BlocklistResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *BlocklistResource) HasCustomFormats() bool {
	if o != nil && !IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *BlocklistResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BlocklistResource) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BlocklistResource) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *BlocklistResource) SetDate(v time.Time) {
	o.Date = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *BlocklistResource) GetProtocol() DownloadProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret DownloadProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *BlocklistResource) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given DownloadProtocol and assigns it to the Protocol field.
func (o *BlocklistResource) SetProtocol(v DownloadProtocol) {
	o.Protocol = &v
}

// GetIndexer returns the Indexer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetIndexer() string {
	if o == nil || IsNil(o.Indexer.Get()) {
		var ret string
		return ret
	}
	return *o.Indexer.Get()
}

// GetIndexerOk returns a tuple with the Indexer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetIndexerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indexer.Get(), o.Indexer.IsSet()
}

// HasIndexer returns a boolean if a field has been set.
func (o *BlocklistResource) HasIndexer() bool {
	if o != nil && o.Indexer.IsSet() {
		return true
	}

	return false
}

// SetIndexer gets a reference to the given NullableString and assigns it to the Indexer field.
func (o *BlocklistResource) SetIndexer(v string) {
	o.Indexer.Set(&v)
}
// SetIndexerNil sets the value for Indexer to be an explicit nil
func (o *BlocklistResource) SetIndexerNil() {
	o.Indexer.Set(nil)
}

// UnsetIndexer ensures that no value is present for Indexer, not even an explicit nil
func (o *BlocklistResource) UnsetIndexer() {
	o.Indexer.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlocklistResource) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlocklistResource) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *BlocklistResource) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *BlocklistResource) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *BlocklistResource) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *BlocklistResource) UnsetMessage() {
	o.Message.Unset()
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *BlocklistResource) GetSeries() SeriesResource {
	if o == nil || IsNil(o.Series) {
		var ret SeriesResource
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlocklistResource) GetSeriesOk() (*SeriesResource, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *BlocklistResource) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given SeriesResource and assigns it to the Series field.
func (o *BlocklistResource) SetSeries(v SeriesResource) {
	o.Series = &v
}

func (o BlocklistResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlocklistResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SeriesId) {
		toSerialize["seriesId"] = o.SeriesId
	}
	if o.EpisodeIds != nil {
		toSerialize["episodeIds"] = o.EpisodeIds
	}
	if o.SourceTitle.IsSet() {
		toSerialize["sourceTitle"] = o.SourceTitle.Get()
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Indexer.IsSet() {
		toSerialize["indexer"] = o.Indexer.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableBlocklistResource struct {
	value *BlocklistResource
	isSet bool
}

func (v NullableBlocklistResource) Get() *BlocklistResource {
	return v.value
}

func (v *NullableBlocklistResource) Set(val *BlocklistResource) {
	v.value = val
	v.isSet = true
}

func (v NullableBlocklistResource) IsSet() bool {
	return v.isSet
}

func (v *NullableBlocklistResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlocklistResource(val *BlocklistResource) *NullableBlocklistResource {
	return &NullableBlocklistResource{value: val, isSet: true}
}

func (v NullableBlocklistResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlocklistResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


