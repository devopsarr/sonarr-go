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

// QualityProfileQualityItemResource struct for QualityProfileQualityItemResource
type QualityProfileQualityItemResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Quality *Quality `json:"quality,omitempty"`
	Items []*QualityProfileQualityItemResource `json:"items,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewQualityProfileQualityItemResource instantiates a new QualityProfileQualityItemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityProfileQualityItemResource() *QualityProfileQualityItemResource {
	this := QualityProfileQualityItemResource{}
	return &this
}

// NewQualityProfileQualityItemResourceWithDefaults instantiates a new QualityProfileQualityItemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityProfileQualityItemResourceWithDefaults() *QualityProfileQualityItemResource {
	this := QualityProfileQualityItemResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QualityProfileQualityItemResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItemResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QualityProfileQualityItemResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QualityProfileQualityItemResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileQualityItemResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileQualityItemResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *QualityProfileQualityItemResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *QualityProfileQualityItemResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *QualityProfileQualityItemResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *QualityProfileQualityItemResource) UnsetName() {
	o.Name.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *QualityProfileQualityItemResource) GetQuality() Quality {
	if o == nil || isNil(o.Quality) {
		var ret Quality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItemResource) GetQualityOk() (*Quality, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *QualityProfileQualityItemResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given Quality and assigns it to the Quality field.
func (o *QualityProfileQualityItemResource) SetQuality(v Quality) {
	o.Quality = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileQualityItemResource) GetItems() []*QualityProfileQualityItemResource {
	if o == nil {
		var ret []*QualityProfileQualityItemResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileQualityItemResource) GetItemsOk() ([]*QualityProfileQualityItemResource, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QualityProfileQualityItemResource) HasItems() bool {
	if o != nil && isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QualityProfileQualityItemResource and assigns it to the Items field.
func (o *QualityProfileQualityItemResource) SetItems(v []*QualityProfileQualityItemResource) {
	o.Items = v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *QualityProfileQualityItemResource) GetAllowed() bool {
	if o == nil || isNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItemResource) GetAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.Allowed) {
    return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *QualityProfileQualityItemResource) HasAllowed() bool {
	if o != nil && !isNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *QualityProfileQualityItemResource) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o QualityProfileQualityItemResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return json.Marshal(toSerialize)
}

type NullableQualityProfileQualityItemResource struct {
	value *QualityProfileQualityItemResource
	isSet bool
}

func (v NullableQualityProfileQualityItemResource) Get() *QualityProfileQualityItemResource {
	return v.value
}

func (v *NullableQualityProfileQualityItemResource) Set(val *QualityProfileQualityItemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityProfileQualityItemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityProfileQualityItemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityProfileQualityItemResource(val *QualityProfileQualityItemResource) *NullableQualityProfileQualityItemResource {
	return &NullableQualityProfileQualityItemResource{value: val, isSet: true}
}

func (v NullableQualityProfileQualityItemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityProfileQualityItemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


