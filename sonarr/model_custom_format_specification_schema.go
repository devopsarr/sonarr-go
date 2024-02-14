/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.1.929
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// CustomFormatSpecificationSchema struct for CustomFormatSpecificationSchema
type CustomFormatSpecificationSchema struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Negate *bool `json:"negate,omitempty"`
	Required *bool `json:"required,omitempty"`
	Fields []*Field `json:"fields,omitempty"`
	Presets []*CustomFormatSpecificationSchema `json:"presets,omitempty"`
}

// NewCustomFormatSpecificationSchema instantiates a new CustomFormatSpecificationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormatSpecificationSchema() *CustomFormatSpecificationSchema {
	this := CustomFormatSpecificationSchema{}
	return &this
}

// NewCustomFormatSpecificationSchemaWithDefaults instantiates a new CustomFormatSpecificationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormatSpecificationSchemaWithDefaults() *CustomFormatSpecificationSchema {
	this := CustomFormatSpecificationSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFormatSpecificationSchema) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormatSpecificationSchema) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFormatSpecificationSchema) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomFormatSpecificationSchema) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomFormatSpecificationSchema) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomFormatSpecificationSchema) UnsetName() {
	o.Name.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetImplementation() string {
	if o == nil || isNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetImplementationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *CustomFormatSpecificationSchema) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *CustomFormatSpecificationSchema) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *CustomFormatSpecificationSchema) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetImplementationName() string {
	if o == nil || isNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetImplementationNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *CustomFormatSpecificationSchema) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *CustomFormatSpecificationSchema) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *CustomFormatSpecificationSchema) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetInfoLink() string {
	if o == nil || isNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetInfoLinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *CustomFormatSpecificationSchema) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *CustomFormatSpecificationSchema) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *CustomFormatSpecificationSchema) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *CustomFormatSpecificationSchema) GetNegate() bool {
	if o == nil || isNil(o.Negate) {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormatSpecificationSchema) GetNegateOk() (*bool, bool) {
	if o == nil || isNil(o.Negate) {
    return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasNegate() bool {
	if o != nil && !isNil(o.Negate) {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *CustomFormatSpecificationSchema) SetNegate(v bool) {
	o.Negate = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *CustomFormatSpecificationSchema) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormatSpecificationSchema) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
    return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *CustomFormatSpecificationSchema) SetRequired(v bool) {
	o.Required = &v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetFields() []*Field {
	if o == nil {
		var ret []*Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetFieldsOk() ([]*Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *CustomFormatSpecificationSchema) SetFields(v []*Field) {
	o.Fields = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatSpecificationSchema) GetPresets() []*CustomFormatSpecificationSchema {
	if o == nil {
		var ret []*CustomFormatSpecificationSchema
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatSpecificationSchema) GetPresetsOk() ([]*CustomFormatSpecificationSchema, bool) {
	if o == nil || isNil(o.Presets) {
    return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *CustomFormatSpecificationSchema) HasPresets() bool {
	if o != nil && isNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []CustomFormatSpecificationSchema and assigns it to the Presets field.
func (o *CustomFormatSpecificationSchema) SetPresets(v []*CustomFormatSpecificationSchema) {
	o.Presets = v
}

func (o CustomFormatSpecificationSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if o.InfoLink.IsSet() {
		toSerialize["infoLink"] = o.InfoLink.Get()
	}
	if !isNil(o.Negate) {
		toSerialize["negate"] = o.Negate
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	return json.Marshal(toSerialize)
}

type NullableCustomFormatSpecificationSchema struct {
	value *CustomFormatSpecificationSchema
	isSet bool
}

func (v NullableCustomFormatSpecificationSchema) Get() *CustomFormatSpecificationSchema {
	return v.value
}

func (v *NullableCustomFormatSpecificationSchema) Set(val *CustomFormatSpecificationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormatSpecificationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormatSpecificationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormatSpecificationSchema(val *CustomFormatSpecificationSchema) *NullableCustomFormatSpecificationSchema {
	return &NullableCustomFormatSpecificationSchema{value: val, isSet: true}
}

func (v NullableCustomFormatSpecificationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormatSpecificationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


