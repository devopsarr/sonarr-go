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

// QualityModel struct for QualityModel
type QualityModel struct {
	Quality *Quality `json:"quality,omitempty"`
	Revision *Revision `json:"revision,omitempty"`
}

// NewQualityModel instantiates a new QualityModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityModel() *QualityModel {
	this := QualityModel{}
	return &this
}

// NewQualityModelWithDefaults instantiates a new QualityModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityModelWithDefaults() *QualityModel {
	this := QualityModel{}
	return &this
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *QualityModel) GetQuality() Quality {
	if o == nil || isNil(o.Quality) {
		var ret Quality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityModel) GetQualityOk() (*Quality, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *QualityModel) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given Quality and assigns it to the Quality field.
func (o *QualityModel) SetQuality(v Quality) {
	o.Quality = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *QualityModel) GetRevision() Revision {
	if o == nil || isNil(o.Revision) {
		var ret Revision
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityModel) GetRevisionOk() (*Revision, bool) {
	if o == nil || isNil(o.Revision) {
    return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *QualityModel) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given Revision and assigns it to the Revision field.
func (o *QualityModel) SetRevision(v Revision) {
	o.Revision = &v
}

func (o QualityModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableQualityModel struct {
	value *QualityModel
	isSet bool
}

func (v NullableQualityModel) Get() *QualityModel {
	return v.value
}

func (v *NullableQualityModel) Set(val *QualityModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityModel(val *QualityModel) *NullableQualityModel {
	return &NullableQualityModel{value: val, isSet: true}
}

func (v NullableQualityModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


