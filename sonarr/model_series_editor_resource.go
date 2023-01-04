/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// SeriesEditorResource struct for SeriesEditorResource
type SeriesEditorResource struct {
	SeriesIds []*int32 `json:"seriesIds,omitempty"`
	Monitored NullableBool `json:"monitored,omitempty"`
	QualityProfileId NullableInt32 `json:"qualityProfileId,omitempty"`
	SeriesType *SeriesTypes `json:"seriesType,omitempty"`
	SeasonFolder NullableBool `json:"seasonFolder,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	MoveFiles *bool `json:"moveFiles,omitempty"`
	DeleteFiles *bool `json:"deleteFiles,omitempty"`
	AddImportListExclusion *bool `json:"addImportListExclusion,omitempty"`
}

// NewSeriesEditorResource instantiates a new SeriesEditorResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesEditorResource() *SeriesEditorResource {
	this := SeriesEditorResource{}
	return &this
}

// NewSeriesEditorResourceWithDefaults instantiates a new SeriesEditorResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesEditorResourceWithDefaults() *SeriesEditorResource {
	this := SeriesEditorResource{}
	return &this
}

// GetSeriesIds returns the SeriesIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetSeriesIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.SeriesIds
}

// GetSeriesIdsOk returns a tuple with the SeriesIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetSeriesIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.SeriesIds) {
    return nil, false
	}
	return o.SeriesIds, true
}

// HasSeriesIds returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasSeriesIds() bool {
	if o != nil && isNil(o.SeriesIds) {
		return true
	}

	return false
}

// SetSeriesIds gets a reference to the given []int32 and assigns it to the SeriesIds field.
func (o *SeriesEditorResource) SetSeriesIds(v []*int32) {
	o.SeriesIds = v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetMonitored() bool {
	if o == nil || isNil(o.Monitored.Get()) {
		var ret bool
		return ret
	}
	return *o.Monitored.Get()
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetMonitoredOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.Monitored.Get(), o.Monitored.IsSet()
}

// HasMonitored returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasMonitored() bool {
	if o != nil && o.Monitored.IsSet() {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given NullableBool and assigns it to the Monitored field.
func (o *SeriesEditorResource) SetMonitored(v bool) {
	o.Monitored.Set(&v)
}
// SetMonitoredNil sets the value for Monitored to be an explicit nil
func (o *SeriesEditorResource) SetMonitoredNil() {
	o.Monitored.Set(nil)
}

// UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
func (o *SeriesEditorResource) UnsetMonitored() {
	o.Monitored.Unset()
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetQualityProfileId() int32 {
	if o == nil || isNil(o.QualityProfileId.Get()) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId.Get()
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.QualityProfileId.Get(), o.QualityProfileId.IsSet()
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasQualityProfileId() bool {
	if o != nil && o.QualityProfileId.IsSet() {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given NullableInt32 and assigns it to the QualityProfileId field.
func (o *SeriesEditorResource) SetQualityProfileId(v int32) {
	o.QualityProfileId.Set(&v)
}
// SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil
func (o *SeriesEditorResource) SetQualityProfileIdNil() {
	o.QualityProfileId.Set(nil)
}

// UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
func (o *SeriesEditorResource) UnsetQualityProfileId() {
	o.QualityProfileId.Unset()
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *SeriesEditorResource) GetSeriesType() SeriesTypes {
	if o == nil || isNil(o.SeriesType) {
		var ret SeriesTypes
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesEditorResource) GetSeriesTypeOk() (*SeriesTypes, bool) {
	if o == nil || isNil(o.SeriesType) {
    return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasSeriesType() bool {
	if o != nil && !isNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given SeriesTypes and assigns it to the SeriesType field.
func (o *SeriesEditorResource) SetSeriesType(v SeriesTypes) {
	o.SeriesType = &v
}

// GetSeasonFolder returns the SeasonFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetSeasonFolder() bool {
	if o == nil || isNil(o.SeasonFolder.Get()) {
		var ret bool
		return ret
	}
	return *o.SeasonFolder.Get()
}

// GetSeasonFolderOk returns a tuple with the SeasonFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetSeasonFolderOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SeasonFolder.Get(), o.SeasonFolder.IsSet()
}

// HasSeasonFolder returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasSeasonFolder() bool {
	if o != nil && o.SeasonFolder.IsSet() {
		return true
	}

	return false
}

// SetSeasonFolder gets a reference to the given NullableBool and assigns it to the SeasonFolder field.
func (o *SeriesEditorResource) SetSeasonFolder(v bool) {
	o.SeasonFolder.Set(&v)
}
// SetSeasonFolderNil sets the value for SeasonFolder to be an explicit nil
func (o *SeriesEditorResource) SetSeasonFolderNil() {
	o.SeasonFolder.Set(nil)
}

// UnsetSeasonFolder ensures that no value is present for SeasonFolder, not even an explicit nil
func (o *SeriesEditorResource) UnsetSeasonFolder() {
	o.SeasonFolder.Unset()
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetRootFolderPath() string {
	if o == nil || isNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *SeriesEditorResource) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *SeriesEditorResource) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *SeriesEditorResource) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesEditorResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesEditorResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *SeriesEditorResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *SeriesEditorResource) GetApplyTags() ApplyTags {
	if o == nil || isNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesEditorResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || isNil(o.ApplyTags) {
    return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasApplyTags() bool {
	if o != nil && !isNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *SeriesEditorResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetMoveFiles returns the MoveFiles field value if set, zero value otherwise.
func (o *SeriesEditorResource) GetMoveFiles() bool {
	if o == nil || isNil(o.MoveFiles) {
		var ret bool
		return ret
	}
	return *o.MoveFiles
}

// GetMoveFilesOk returns a tuple with the MoveFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesEditorResource) GetMoveFilesOk() (*bool, bool) {
	if o == nil || isNil(o.MoveFiles) {
    return nil, false
	}
	return o.MoveFiles, true
}

// HasMoveFiles returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasMoveFiles() bool {
	if o != nil && !isNil(o.MoveFiles) {
		return true
	}

	return false
}

// SetMoveFiles gets a reference to the given bool and assigns it to the MoveFiles field.
func (o *SeriesEditorResource) SetMoveFiles(v bool) {
	o.MoveFiles = &v
}

// GetDeleteFiles returns the DeleteFiles field value if set, zero value otherwise.
func (o *SeriesEditorResource) GetDeleteFiles() bool {
	if o == nil || isNil(o.DeleteFiles) {
		var ret bool
		return ret
	}
	return *o.DeleteFiles
}

// GetDeleteFilesOk returns a tuple with the DeleteFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesEditorResource) GetDeleteFilesOk() (*bool, bool) {
	if o == nil || isNil(o.DeleteFiles) {
    return nil, false
	}
	return o.DeleteFiles, true
}

// HasDeleteFiles returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasDeleteFiles() bool {
	if o != nil && !isNil(o.DeleteFiles) {
		return true
	}

	return false
}

// SetDeleteFiles gets a reference to the given bool and assigns it to the DeleteFiles field.
func (o *SeriesEditorResource) SetDeleteFiles(v bool) {
	o.DeleteFiles = &v
}

// GetAddImportListExclusion returns the AddImportListExclusion field value if set, zero value otherwise.
func (o *SeriesEditorResource) GetAddImportListExclusion() bool {
	if o == nil || isNil(o.AddImportListExclusion) {
		var ret bool
		return ret
	}
	return *o.AddImportListExclusion
}

// GetAddImportListExclusionOk returns a tuple with the AddImportListExclusion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesEditorResource) GetAddImportListExclusionOk() (*bool, bool) {
	if o == nil || isNil(o.AddImportListExclusion) {
    return nil, false
	}
	return o.AddImportListExclusion, true
}

// HasAddImportListExclusion returns a boolean if a field has been set.
func (o *SeriesEditorResource) HasAddImportListExclusion() bool {
	if o != nil && !isNil(o.AddImportListExclusion) {
		return true
	}

	return false
}

// SetAddImportListExclusion gets a reference to the given bool and assigns it to the AddImportListExclusion field.
func (o *SeriesEditorResource) SetAddImportListExclusion(v bool) {
	o.AddImportListExclusion = &v
}

func (o SeriesEditorResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SeriesIds != nil {
		toSerialize["seriesIds"] = o.SeriesIds
	}
	if o.Monitored.IsSet() {
		toSerialize["monitored"] = o.Monitored.Get()
	}
	if o.QualityProfileId.IsSet() {
		toSerialize["qualityProfileId"] = o.QualityProfileId.Get()
	}
	if !isNil(o.SeriesType) {
		toSerialize["seriesType"] = o.SeriesType
	}
	if o.SeasonFolder.IsSet() {
		toSerialize["seasonFolder"] = o.SeasonFolder.Get()
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if !isNil(o.MoveFiles) {
		toSerialize["moveFiles"] = o.MoveFiles
	}
	if !isNil(o.DeleteFiles) {
		toSerialize["deleteFiles"] = o.DeleteFiles
	}
	if !isNil(o.AddImportListExclusion) {
		toSerialize["addImportListExclusion"] = o.AddImportListExclusion
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesEditorResource struct {
	value *SeriesEditorResource
	isSet bool
}

func (v NullableSeriesEditorResource) Get() *SeriesEditorResource {
	return v.value
}

func (v *NullableSeriesEditorResource) Set(val *SeriesEditorResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesEditorResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesEditorResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesEditorResource(val *SeriesEditorResource) *NullableSeriesEditorResource {
	return &NullableSeriesEditorResource{value: val, isSet: true}
}

func (v NullableSeriesEditorResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesEditorResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


