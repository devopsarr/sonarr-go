/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
	"time"
)

// HistoryResource struct for HistoryResource
type HistoryResource struct {
	Id *int32 `json:"id,omitempty"`
	EpisodeId *int32 `json:"episodeId,omitempty"`
	SeriesId *int32 `json:"seriesId,omitempty"`
	SourceTitle NullableString `json:"sourceTitle,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	CustomFormats []*CustomFormatResource `json:"customFormats,omitempty"`
	QualityCutoffNotMet *bool `json:"qualityCutoffNotMet,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	EventType *EpisodeHistoryEventType `json:"eventType,omitempty"`
	Data map[string]string `json:"data,omitempty"`
	Episode *EpisodeResource `json:"episode,omitempty"`
	Series *SeriesResource `json:"series,omitempty"`
}

// NewHistoryResource instantiates a new HistoryResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryResource() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// NewHistoryResourceWithDefaults instantiates a new HistoryResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryResourceWithDefaults() *HistoryResource {
	this := HistoryResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HistoryResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HistoryResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *HistoryResource) SetId(v int32) {
	o.Id = &v
}

// GetEpisodeId returns the EpisodeId field value if set, zero value otherwise.
func (o *HistoryResource) GetEpisodeId() int32 {
	if o == nil || isNil(o.EpisodeId) {
		var ret int32
		return ret
	}
	return *o.EpisodeId
}

// GetEpisodeIdOk returns a tuple with the EpisodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetEpisodeIdOk() (*int32, bool) {
	if o == nil || isNil(o.EpisodeId) {
    return nil, false
	}
	return o.EpisodeId, true
}

// HasEpisodeId returns a boolean if a field has been set.
func (o *HistoryResource) HasEpisodeId() bool {
	if o != nil && !isNil(o.EpisodeId) {
		return true
	}

	return false
}

// SetEpisodeId gets a reference to the given int32 and assigns it to the EpisodeId field.
func (o *HistoryResource) SetEpisodeId(v int32) {
	o.EpisodeId = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *HistoryResource) GetSeriesId() int32 {
	if o == nil || isNil(o.SeriesId) {
		var ret int32
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetSeriesIdOk() (*int32, bool) {
	if o == nil || isNil(o.SeriesId) {
    return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *HistoryResource) HasSeriesId() bool {
	if o != nil && !isNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int32 and assigns it to the SeriesId field.
func (o *HistoryResource) SetSeriesId(v int32) {
	o.SeriesId = &v
}

// GetSourceTitle returns the SourceTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetSourceTitle() string {
	if o == nil || isNil(o.SourceTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SourceTitle.Get()
}

// GetSourceTitleOk returns a tuple with the SourceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetSourceTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SourceTitle.Get(), o.SourceTitle.IsSet()
}

// HasSourceTitle returns a boolean if a field has been set.
func (o *HistoryResource) HasSourceTitle() bool {
	if o != nil && o.SourceTitle.IsSet() {
		return true
	}

	return false
}

// SetSourceTitle gets a reference to the given NullableString and assigns it to the SourceTitle field.
func (o *HistoryResource) SetSourceTitle(v string) {
	o.SourceTitle.Set(&v)
}
// SetSourceTitleNil sets the value for SourceTitle to be an explicit nil
func (o *HistoryResource) SetSourceTitleNil() {
	o.SourceTitle.Set(nil)
}

// UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
func (o *HistoryResource) UnsetSourceTitle() {
	o.SourceTitle.Unset()
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *HistoryResource) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *HistoryResource) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *HistoryResource) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *HistoryResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *HistoryResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetCustomFormats() []*CustomFormatResource {
	if o == nil {
		var ret []*CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetCustomFormatsOk() ([]*CustomFormatResource, bool) {
	if o == nil || isNil(o.CustomFormats) {
    return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *HistoryResource) HasCustomFormats() bool {
	if o != nil && isNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *HistoryResource) SetCustomFormats(v []*CustomFormatResource) {
	o.CustomFormats = v
}

// GetQualityCutoffNotMet returns the QualityCutoffNotMet field value if set, zero value otherwise.
func (o *HistoryResource) GetQualityCutoffNotMet() bool {
	if o == nil || isNil(o.QualityCutoffNotMet) {
		var ret bool
		return ret
	}
	return *o.QualityCutoffNotMet
}

// GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetQualityCutoffNotMetOk() (*bool, bool) {
	if o == nil || isNil(o.QualityCutoffNotMet) {
    return nil, false
	}
	return o.QualityCutoffNotMet, true
}

// HasQualityCutoffNotMet returns a boolean if a field has been set.
func (o *HistoryResource) HasQualityCutoffNotMet() bool {
	if o != nil && !isNil(o.QualityCutoffNotMet) {
		return true
	}

	return false
}

// SetQualityCutoffNotMet gets a reference to the given bool and assigns it to the QualityCutoffNotMet field.
func (o *HistoryResource) SetQualityCutoffNotMet(v bool) {
	o.QualityCutoffNotMet = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HistoryResource) GetDate() time.Time {
	if o == nil || isNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HistoryResource) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *HistoryResource) SetDate(v time.Time) {
	o.Date = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetDownloadId() string {
	if o == nil || isNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *HistoryResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *HistoryResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *HistoryResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *HistoryResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *HistoryResource) GetEventType() EpisodeHistoryEventType {
	if o == nil || isNil(o.EventType) {
		var ret EpisodeHistoryEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetEventTypeOk() (*EpisodeHistoryEventType, bool) {
	if o == nil || isNil(o.EventType) {
    return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *HistoryResource) HasEventType() bool {
	if o != nil && !isNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given EpisodeHistoryEventType and assigns it to the EventType field.
func (o *HistoryResource) SetEventType(v EpisodeHistoryEventType) {
	o.EventType = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoryResource) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoryResource) GetDataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoryResource) HasData() bool {
	if o != nil && isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *HistoryResource) SetData(v map[string]string) {
	o.Data = v
}

// GetEpisode returns the Episode field value if set, zero value otherwise.
func (o *HistoryResource) GetEpisode() EpisodeResource {
	if o == nil || isNil(o.Episode) {
		var ret EpisodeResource
		return ret
	}
	return *o.Episode
}

// GetEpisodeOk returns a tuple with the Episode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetEpisodeOk() (*EpisodeResource, bool) {
	if o == nil || isNil(o.Episode) {
    return nil, false
	}
	return o.Episode, true
}

// HasEpisode returns a boolean if a field has been set.
func (o *HistoryResource) HasEpisode() bool {
	if o != nil && !isNil(o.Episode) {
		return true
	}

	return false
}

// SetEpisode gets a reference to the given EpisodeResource and assigns it to the Episode field.
func (o *HistoryResource) SetEpisode(v EpisodeResource) {
	o.Episode = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *HistoryResource) GetSeries() SeriesResource {
	if o == nil || isNil(o.Series) {
		var ret SeriesResource
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResource) GetSeriesOk() (*SeriesResource, bool) {
	if o == nil || isNil(o.Series) {
    return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *HistoryResource) HasSeries() bool {
	if o != nil && !isNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given SeriesResource and assigns it to the Series field.
func (o *HistoryResource) SetSeries(v SeriesResource) {
	o.Series = &v
}

func (o HistoryResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.EpisodeId) {
		toSerialize["episodeId"] = o.EpisodeId
	}
	if !isNil(o.SeriesId) {
		toSerialize["seriesId"] = o.SeriesId
	}
	if o.SourceTitle.IsSet() {
		toSerialize["sourceTitle"] = o.SourceTitle.Get()
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !isNil(o.QualityCutoffNotMet) {
		toSerialize["qualityCutoffNotMet"] = o.QualityCutoffNotMet
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if !isNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Episode) {
		toSerialize["episode"] = o.Episode
	}
	if !isNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryResource struct {
	value *HistoryResource
	isSet bool
}

func (v NullableHistoryResource) Get() *HistoryResource {
	return v.value
}

func (v *NullableHistoryResource) Set(val *HistoryResource) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryResource) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryResource(val *HistoryResource) *NullableHistoryResource {
	return &NullableHistoryResource{value: val, isSet: true}
}

func (v NullableHistoryResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


