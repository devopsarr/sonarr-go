/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.9.2244
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the DelayProfileResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelayProfileResource{}

// DelayProfileResource struct for DelayProfileResource
type DelayProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	EnableUsenet *bool `json:"enableUsenet,omitempty"`
	EnableTorrent *bool `json:"enableTorrent,omitempty"`
	PreferredProtocol *DownloadProtocol `json:"preferredProtocol,omitempty"`
	UsenetDelay *int32 `json:"usenetDelay,omitempty"`
	TorrentDelay *int32 `json:"torrentDelay,omitempty"`
	BypassIfHighestQuality *bool `json:"bypassIfHighestQuality,omitempty"`
	BypassIfAboveCustomFormatScore *bool `json:"bypassIfAboveCustomFormatScore,omitempty"`
	MinimumCustomFormatScore *int32 `json:"minimumCustomFormatScore,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
}

// NewDelayProfileResource instantiates a new DelayProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayProfileResource() *DelayProfileResource {
	this := DelayProfileResource{}
	return &this
}

// NewDelayProfileResourceWithDefaults instantiates a new DelayProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayProfileResourceWithDefaults() *DelayProfileResource {
	this := DelayProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DelayProfileResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DelayProfileResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DelayProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetEnableUsenet returns the EnableUsenet field value if set, zero value otherwise.
func (o *DelayProfileResource) GetEnableUsenet() bool {
	if o == nil || IsNil(o.EnableUsenet) {
		var ret bool
		return ret
	}
	return *o.EnableUsenet
}

// GetEnableUsenetOk returns a tuple with the EnableUsenet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetEnableUsenetOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUsenet) {
		return nil, false
	}
	return o.EnableUsenet, true
}

// HasEnableUsenet returns a boolean if a field has been set.
func (o *DelayProfileResource) HasEnableUsenet() bool {
	if o != nil && !IsNil(o.EnableUsenet) {
		return true
	}

	return false
}

// SetEnableUsenet gets a reference to the given bool and assigns it to the EnableUsenet field.
func (o *DelayProfileResource) SetEnableUsenet(v bool) {
	o.EnableUsenet = &v
}

// GetEnableTorrent returns the EnableTorrent field value if set, zero value otherwise.
func (o *DelayProfileResource) GetEnableTorrent() bool {
	if o == nil || IsNil(o.EnableTorrent) {
		var ret bool
		return ret
	}
	return *o.EnableTorrent
}

// GetEnableTorrentOk returns a tuple with the EnableTorrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetEnableTorrentOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTorrent) {
		return nil, false
	}
	return o.EnableTorrent, true
}

// HasEnableTorrent returns a boolean if a field has been set.
func (o *DelayProfileResource) HasEnableTorrent() bool {
	if o != nil && !IsNil(o.EnableTorrent) {
		return true
	}

	return false
}

// SetEnableTorrent gets a reference to the given bool and assigns it to the EnableTorrent field.
func (o *DelayProfileResource) SetEnableTorrent(v bool) {
	o.EnableTorrent = &v
}

// GetPreferredProtocol returns the PreferredProtocol field value if set, zero value otherwise.
func (o *DelayProfileResource) GetPreferredProtocol() DownloadProtocol {
	if o == nil || IsNil(o.PreferredProtocol) {
		var ret DownloadProtocol
		return ret
	}
	return *o.PreferredProtocol
}

// GetPreferredProtocolOk returns a tuple with the PreferredProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetPreferredProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || IsNil(o.PreferredProtocol) {
		return nil, false
	}
	return o.PreferredProtocol, true
}

// HasPreferredProtocol returns a boolean if a field has been set.
func (o *DelayProfileResource) HasPreferredProtocol() bool {
	if o != nil && !IsNil(o.PreferredProtocol) {
		return true
	}

	return false
}

// SetPreferredProtocol gets a reference to the given DownloadProtocol and assigns it to the PreferredProtocol field.
func (o *DelayProfileResource) SetPreferredProtocol(v DownloadProtocol) {
	o.PreferredProtocol = &v
}

// GetUsenetDelay returns the UsenetDelay field value if set, zero value otherwise.
func (o *DelayProfileResource) GetUsenetDelay() int32 {
	if o == nil || IsNil(o.UsenetDelay) {
		var ret int32
		return ret
	}
	return *o.UsenetDelay
}

// GetUsenetDelayOk returns a tuple with the UsenetDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetUsenetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.UsenetDelay) {
		return nil, false
	}
	return o.UsenetDelay, true
}

// HasUsenetDelay returns a boolean if a field has been set.
func (o *DelayProfileResource) HasUsenetDelay() bool {
	if o != nil && !IsNil(o.UsenetDelay) {
		return true
	}

	return false
}

// SetUsenetDelay gets a reference to the given int32 and assigns it to the UsenetDelay field.
func (o *DelayProfileResource) SetUsenetDelay(v int32) {
	o.UsenetDelay = &v
}

// GetTorrentDelay returns the TorrentDelay field value if set, zero value otherwise.
func (o *DelayProfileResource) GetTorrentDelay() int32 {
	if o == nil || IsNil(o.TorrentDelay) {
		var ret int32
		return ret
	}
	return *o.TorrentDelay
}

// GetTorrentDelayOk returns a tuple with the TorrentDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetTorrentDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.TorrentDelay) {
		return nil, false
	}
	return o.TorrentDelay, true
}

// HasTorrentDelay returns a boolean if a field has been set.
func (o *DelayProfileResource) HasTorrentDelay() bool {
	if o != nil && !IsNil(o.TorrentDelay) {
		return true
	}

	return false
}

// SetTorrentDelay gets a reference to the given int32 and assigns it to the TorrentDelay field.
func (o *DelayProfileResource) SetTorrentDelay(v int32) {
	o.TorrentDelay = &v
}

// GetBypassIfHighestQuality returns the BypassIfHighestQuality field value if set, zero value otherwise.
func (o *DelayProfileResource) GetBypassIfHighestQuality() bool {
	if o == nil || IsNil(o.BypassIfHighestQuality) {
		var ret bool
		return ret
	}
	return *o.BypassIfHighestQuality
}

// GetBypassIfHighestQualityOk returns a tuple with the BypassIfHighestQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetBypassIfHighestQualityOk() (*bool, bool) {
	if o == nil || IsNil(o.BypassIfHighestQuality) {
		return nil, false
	}
	return o.BypassIfHighestQuality, true
}

// HasBypassIfHighestQuality returns a boolean if a field has been set.
func (o *DelayProfileResource) HasBypassIfHighestQuality() bool {
	if o != nil && !IsNil(o.BypassIfHighestQuality) {
		return true
	}

	return false
}

// SetBypassIfHighestQuality gets a reference to the given bool and assigns it to the BypassIfHighestQuality field.
func (o *DelayProfileResource) SetBypassIfHighestQuality(v bool) {
	o.BypassIfHighestQuality = &v
}

// GetBypassIfAboveCustomFormatScore returns the BypassIfAboveCustomFormatScore field value if set, zero value otherwise.
func (o *DelayProfileResource) GetBypassIfAboveCustomFormatScore() bool {
	if o == nil || IsNil(o.BypassIfAboveCustomFormatScore) {
		var ret bool
		return ret
	}
	return *o.BypassIfAboveCustomFormatScore
}

// GetBypassIfAboveCustomFormatScoreOk returns a tuple with the BypassIfAboveCustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetBypassIfAboveCustomFormatScoreOk() (*bool, bool) {
	if o == nil || IsNil(o.BypassIfAboveCustomFormatScore) {
		return nil, false
	}
	return o.BypassIfAboveCustomFormatScore, true
}

// HasBypassIfAboveCustomFormatScore returns a boolean if a field has been set.
func (o *DelayProfileResource) HasBypassIfAboveCustomFormatScore() bool {
	if o != nil && !IsNil(o.BypassIfAboveCustomFormatScore) {
		return true
	}

	return false
}

// SetBypassIfAboveCustomFormatScore gets a reference to the given bool and assigns it to the BypassIfAboveCustomFormatScore field.
func (o *DelayProfileResource) SetBypassIfAboveCustomFormatScore(v bool) {
	o.BypassIfAboveCustomFormatScore = &v
}

// GetMinimumCustomFormatScore returns the MinimumCustomFormatScore field value if set, zero value otherwise.
func (o *DelayProfileResource) GetMinimumCustomFormatScore() int32 {
	if o == nil || IsNil(o.MinimumCustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.MinimumCustomFormatScore
}

// GetMinimumCustomFormatScoreOk returns a tuple with the MinimumCustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetMinimumCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumCustomFormatScore) {
		return nil, false
	}
	return o.MinimumCustomFormatScore, true
}

// HasMinimumCustomFormatScore returns a boolean if a field has been set.
func (o *DelayProfileResource) HasMinimumCustomFormatScore() bool {
	if o != nil && !IsNil(o.MinimumCustomFormatScore) {
		return true
	}

	return false
}

// SetMinimumCustomFormatScore gets a reference to the given int32 and assigns it to the MinimumCustomFormatScore field.
func (o *DelayProfileResource) SetMinimumCustomFormatScore(v int32) {
	o.MinimumCustomFormatScore = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DelayProfileResource) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DelayProfileResource) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *DelayProfileResource) SetOrder(v int32) {
	o.Order = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DelayProfileResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DelayProfileResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DelayProfileResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *DelayProfileResource) SetTags(v []int32) {
	o.Tags = v
}

func (o DelayProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelayProfileResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EnableUsenet) {
		toSerialize["enableUsenet"] = o.EnableUsenet
	}
	if !IsNil(o.EnableTorrent) {
		toSerialize["enableTorrent"] = o.EnableTorrent
	}
	if !IsNil(o.PreferredProtocol) {
		toSerialize["preferredProtocol"] = o.PreferredProtocol
	}
	if !IsNil(o.UsenetDelay) {
		toSerialize["usenetDelay"] = o.UsenetDelay
	}
	if !IsNil(o.TorrentDelay) {
		toSerialize["torrentDelay"] = o.TorrentDelay
	}
	if !IsNil(o.BypassIfHighestQuality) {
		toSerialize["bypassIfHighestQuality"] = o.BypassIfHighestQuality
	}
	if !IsNil(o.BypassIfAboveCustomFormatScore) {
		toSerialize["bypassIfAboveCustomFormatScore"] = o.BypassIfAboveCustomFormatScore
	}
	if !IsNil(o.MinimumCustomFormatScore) {
		toSerialize["minimumCustomFormatScore"] = o.MinimumCustomFormatScore
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDelayProfileResource struct {
	value *DelayProfileResource
	isSet bool
}

func (v NullableDelayProfileResource) Get() *DelayProfileResource {
	return v.value
}

func (v *NullableDelayProfileResource) Set(val *DelayProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayProfileResource(val *DelayProfileResource) *NullableDelayProfileResource {
	return &NullableDelayProfileResource{value: val, isSet: true}
}

func (v NullableDelayProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


