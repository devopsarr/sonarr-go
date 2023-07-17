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

// TagDetailsResource struct for TagDetailsResource
type TagDetailsResource struct {
	Id *int32 `json:"id,omitempty"`
	Label NullableString `json:"label,omitempty"`
	DelayProfileIds []*int32 `json:"delayProfileIds,omitempty"`
	ImportListIds []*int32 `json:"importListIds,omitempty"`
	NotificationIds []*int32 `json:"notificationIds,omitempty"`
	RestrictionIds []*int32 `json:"restrictionIds,omitempty"`
	IndexerIds []*int32 `json:"indexerIds,omitempty"`
	DownloadClientIds []*int32 `json:"downloadClientIds,omitempty"`
	AutoTagIds []*int32 `json:"autoTagIds,omitempty"`
	SeriesIds []*int32 `json:"seriesIds,omitempty"`
}

// NewTagDetailsResource instantiates a new TagDetailsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDetailsResource() *TagDetailsResource {
	this := TagDetailsResource{}
	return &this
}

// NewTagDetailsResourceWithDefaults instantiates a new TagDetailsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDetailsResourceWithDefaults() *TagDetailsResource {
	this := TagDetailsResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagDetailsResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagDetailsResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagDetailsResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TagDetailsResource) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetLabel() string {
	if o == nil || isNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetLabelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *TagDetailsResource) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *TagDetailsResource) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *TagDetailsResource) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *TagDetailsResource) UnsetLabel() {
	o.Label.Unset()
}

// GetDelayProfileIds returns the DelayProfileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetDelayProfileIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.DelayProfileIds
}

// GetDelayProfileIdsOk returns a tuple with the DelayProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetDelayProfileIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.DelayProfileIds) {
    return nil, false
	}
	return o.DelayProfileIds, true
}

// HasDelayProfileIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasDelayProfileIds() bool {
	if o != nil && isNil(o.DelayProfileIds) {
		return true
	}

	return false
}

// SetDelayProfileIds gets a reference to the given []int32 and assigns it to the DelayProfileIds field.
func (o *TagDetailsResource) SetDelayProfileIds(v []*int32) {
	o.DelayProfileIds = v
}

// GetImportListIds returns the ImportListIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetImportListIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.ImportListIds
}

// GetImportListIdsOk returns a tuple with the ImportListIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetImportListIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.ImportListIds) {
    return nil, false
	}
	return o.ImportListIds, true
}

// HasImportListIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasImportListIds() bool {
	if o != nil && isNil(o.ImportListIds) {
		return true
	}

	return false
}

// SetImportListIds gets a reference to the given []int32 and assigns it to the ImportListIds field.
func (o *TagDetailsResource) SetImportListIds(v []*int32) {
	o.ImportListIds = v
}

// GetNotificationIds returns the NotificationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetNotificationIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.NotificationIds
}

// GetNotificationIdsOk returns a tuple with the NotificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetNotificationIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.NotificationIds) {
    return nil, false
	}
	return o.NotificationIds, true
}

// HasNotificationIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasNotificationIds() bool {
	if o != nil && isNil(o.NotificationIds) {
		return true
	}

	return false
}

// SetNotificationIds gets a reference to the given []int32 and assigns it to the NotificationIds field.
func (o *TagDetailsResource) SetNotificationIds(v []*int32) {
	o.NotificationIds = v
}

// GetRestrictionIds returns the RestrictionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetRestrictionIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.RestrictionIds
}

// GetRestrictionIdsOk returns a tuple with the RestrictionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetRestrictionIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.RestrictionIds) {
    return nil, false
	}
	return o.RestrictionIds, true
}

// HasRestrictionIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasRestrictionIds() bool {
	if o != nil && isNil(o.RestrictionIds) {
		return true
	}

	return false
}

// SetRestrictionIds gets a reference to the given []int32 and assigns it to the RestrictionIds field.
func (o *TagDetailsResource) SetRestrictionIds(v []*int32) {
	o.RestrictionIds = v
}

// GetIndexerIds returns the IndexerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetIndexerIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.IndexerIds
}

// GetIndexerIdsOk returns a tuple with the IndexerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetIndexerIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.IndexerIds) {
    return nil, false
	}
	return o.IndexerIds, true
}

// HasIndexerIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasIndexerIds() bool {
	if o != nil && isNil(o.IndexerIds) {
		return true
	}

	return false
}

// SetIndexerIds gets a reference to the given []int32 and assigns it to the IndexerIds field.
func (o *TagDetailsResource) SetIndexerIds(v []*int32) {
	o.IndexerIds = v
}

// GetDownloadClientIds returns the DownloadClientIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetDownloadClientIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.DownloadClientIds
}

// GetDownloadClientIdsOk returns a tuple with the DownloadClientIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetDownloadClientIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.DownloadClientIds) {
    return nil, false
	}
	return o.DownloadClientIds, true
}

// HasDownloadClientIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasDownloadClientIds() bool {
	if o != nil && isNil(o.DownloadClientIds) {
		return true
	}

	return false
}

// SetDownloadClientIds gets a reference to the given []int32 and assigns it to the DownloadClientIds field.
func (o *TagDetailsResource) SetDownloadClientIds(v []*int32) {
	o.DownloadClientIds = v
}

// GetAutoTagIds returns the AutoTagIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetAutoTagIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.AutoTagIds
}

// GetAutoTagIdsOk returns a tuple with the AutoTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetAutoTagIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.AutoTagIds) {
    return nil, false
	}
	return o.AutoTagIds, true
}

// HasAutoTagIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasAutoTagIds() bool {
	if o != nil && isNil(o.AutoTagIds) {
		return true
	}

	return false
}

// SetAutoTagIds gets a reference to the given []int32 and assigns it to the AutoTagIds field.
func (o *TagDetailsResource) SetAutoTagIds(v []*int32) {
	o.AutoTagIds = v
}

// GetSeriesIds returns the SeriesIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagDetailsResource) GetSeriesIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.SeriesIds
}

// GetSeriesIdsOk returns a tuple with the SeriesIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagDetailsResource) GetSeriesIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.SeriesIds) {
    return nil, false
	}
	return o.SeriesIds, true
}

// HasSeriesIds returns a boolean if a field has been set.
func (o *TagDetailsResource) HasSeriesIds() bool {
	if o != nil && isNil(o.SeriesIds) {
		return true
	}

	return false
}

// SetSeriesIds gets a reference to the given []int32 and assigns it to the SeriesIds field.
func (o *TagDetailsResource) SetSeriesIds(v []*int32) {
	o.SeriesIds = v
}

func (o TagDetailsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.DelayProfileIds != nil {
		toSerialize["delayProfileIds"] = o.DelayProfileIds
	}
	if o.ImportListIds != nil {
		toSerialize["importListIds"] = o.ImportListIds
	}
	if o.NotificationIds != nil {
		toSerialize["notificationIds"] = o.NotificationIds
	}
	if o.RestrictionIds != nil {
		toSerialize["restrictionIds"] = o.RestrictionIds
	}
	if o.IndexerIds != nil {
		toSerialize["indexerIds"] = o.IndexerIds
	}
	if o.DownloadClientIds != nil {
		toSerialize["downloadClientIds"] = o.DownloadClientIds
	}
	if o.AutoTagIds != nil {
		toSerialize["autoTagIds"] = o.AutoTagIds
	}
	if o.SeriesIds != nil {
		toSerialize["seriesIds"] = o.SeriesIds
	}
	return json.Marshal(toSerialize)
}

type NullableTagDetailsResource struct {
	value *TagDetailsResource
	isSet bool
}

func (v NullableTagDetailsResource) Get() *TagDetailsResource {
	return v.value
}

func (v *NullableTagDetailsResource) Set(val *TagDetailsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDetailsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDetailsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDetailsResource(val *TagDetailsResource) *NullableTagDetailsResource {
	return &NullableTagDetailsResource{value: val, isSet: true}
}

func (v NullableTagDetailsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDetailsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


