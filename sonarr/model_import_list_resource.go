/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.11.2680
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"encoding/json"
)

// checks if the ImportListResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportListResource{}

// ImportListResource struct for ImportListResource
type ImportListResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presets []ImportListResource `json:"presets,omitempty"`
	EnableAutomaticAdd *bool `json:"enableAutomaticAdd,omitempty"`
	SearchForMissingEpisodes *bool `json:"searchForMissingEpisodes,omitempty"`
	ShouldMonitor *MonitorTypes `json:"shouldMonitor,omitempty"`
	MonitorNewItems *NewItemMonitorTypes `json:"monitorNewItems,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	QualityProfileId *int32 `json:"qualityProfileId,omitempty"`
	SeriesType *SeriesTypes `json:"seriesType,omitempty"`
	SeasonFolder *bool `json:"seasonFolder,omitempty"`
	ListType *ImportListType `json:"listType,omitempty"`
	ListOrder *int32 `json:"listOrder,omitempty"`
	MinRefreshInterval *string `json:"minRefreshInterval,omitempty"`
}

// NewImportListResource instantiates a new ImportListResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListResource() *ImportListResource {
	this := ImportListResource{}
	return &this
}

// NewImportListResourceWithDefaults instantiates a new ImportListResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListResourceWithDefaults() *ImportListResource {
	this := ImportListResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportListResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportListResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportListResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImportListResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImportListResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImportListResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImportListResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ImportListResource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *ImportListResource) SetFields(v []Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *ImportListResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *ImportListResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *ImportListResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *ImportListResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *ImportListResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *ImportListResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *ImportListResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *ImportListResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetConfigContract() string {
	if o == nil || IsNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *ImportListResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *ImportListResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *ImportListResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *ImportListResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetInfoLink() string {
	if o == nil || IsNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *ImportListResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *ImportListResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *ImportListResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *ImportListResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImportListResource) GetMessage() ProviderMessage {
	if o == nil || IsNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImportListResource) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *ImportListResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImportListResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ImportListResource) SetTags(v []int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetPresets() []ImportListResource {
	if o == nil {
		var ret []ImportListResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetPresetsOk() ([]ImportListResource, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *ImportListResource) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []ImportListResource and assigns it to the Presets field.
func (o *ImportListResource) SetPresets(v []ImportListResource) {
	o.Presets = v
}

// GetEnableAutomaticAdd returns the EnableAutomaticAdd field value if set, zero value otherwise.
func (o *ImportListResource) GetEnableAutomaticAdd() bool {
	if o == nil || IsNil(o.EnableAutomaticAdd) {
		var ret bool
		return ret
	}
	return *o.EnableAutomaticAdd
}

// GetEnableAutomaticAddOk returns a tuple with the EnableAutomaticAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetEnableAutomaticAddOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutomaticAdd) {
		return nil, false
	}
	return o.EnableAutomaticAdd, true
}

// HasEnableAutomaticAdd returns a boolean if a field has been set.
func (o *ImportListResource) HasEnableAutomaticAdd() bool {
	if o != nil && !IsNil(o.EnableAutomaticAdd) {
		return true
	}

	return false
}

// SetEnableAutomaticAdd gets a reference to the given bool and assigns it to the EnableAutomaticAdd field.
func (o *ImportListResource) SetEnableAutomaticAdd(v bool) {
	o.EnableAutomaticAdd = &v
}

// GetSearchForMissingEpisodes returns the SearchForMissingEpisodes field value if set, zero value otherwise.
func (o *ImportListResource) GetSearchForMissingEpisodes() bool {
	if o == nil || IsNil(o.SearchForMissingEpisodes) {
		var ret bool
		return ret
	}
	return *o.SearchForMissingEpisodes
}

// GetSearchForMissingEpisodesOk returns a tuple with the SearchForMissingEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetSearchForMissingEpisodesOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchForMissingEpisodes) {
		return nil, false
	}
	return o.SearchForMissingEpisodes, true
}

// HasSearchForMissingEpisodes returns a boolean if a field has been set.
func (o *ImportListResource) HasSearchForMissingEpisodes() bool {
	if o != nil && !IsNil(o.SearchForMissingEpisodes) {
		return true
	}

	return false
}

// SetSearchForMissingEpisodes gets a reference to the given bool and assigns it to the SearchForMissingEpisodes field.
func (o *ImportListResource) SetSearchForMissingEpisodes(v bool) {
	o.SearchForMissingEpisodes = &v
}

// GetShouldMonitor returns the ShouldMonitor field value if set, zero value otherwise.
func (o *ImportListResource) GetShouldMonitor() MonitorTypes {
	if o == nil || IsNil(o.ShouldMonitor) {
		var ret MonitorTypes
		return ret
	}
	return *o.ShouldMonitor
}

// GetShouldMonitorOk returns a tuple with the ShouldMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetShouldMonitorOk() (*MonitorTypes, bool) {
	if o == nil || IsNil(o.ShouldMonitor) {
		return nil, false
	}
	return o.ShouldMonitor, true
}

// HasShouldMonitor returns a boolean if a field has been set.
func (o *ImportListResource) HasShouldMonitor() bool {
	if o != nil && !IsNil(o.ShouldMonitor) {
		return true
	}

	return false
}

// SetShouldMonitor gets a reference to the given MonitorTypes and assigns it to the ShouldMonitor field.
func (o *ImportListResource) SetShouldMonitor(v MonitorTypes) {
	o.ShouldMonitor = &v
}

// GetMonitorNewItems returns the MonitorNewItems field value if set, zero value otherwise.
func (o *ImportListResource) GetMonitorNewItems() NewItemMonitorTypes {
	if o == nil || IsNil(o.MonitorNewItems) {
		var ret NewItemMonitorTypes
		return ret
	}
	return *o.MonitorNewItems
}

// GetMonitorNewItemsOk returns a tuple with the MonitorNewItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetMonitorNewItemsOk() (*NewItemMonitorTypes, bool) {
	if o == nil || IsNil(o.MonitorNewItems) {
		return nil, false
	}
	return o.MonitorNewItems, true
}

// HasMonitorNewItems returns a boolean if a field has been set.
func (o *ImportListResource) HasMonitorNewItems() bool {
	if o != nil && !IsNil(o.MonitorNewItems) {
		return true
	}

	return false
}

// SetMonitorNewItems gets a reference to the given NewItemMonitorTypes and assigns it to the MonitorNewItems field.
func (o *ImportListResource) SetMonitorNewItems(v NewItemMonitorTypes) {
	o.MonitorNewItems = &v
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListResource) GetRootFolderPath() string {
	if o == nil || IsNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListResource) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *ImportListResource) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *ImportListResource) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *ImportListResource) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *ImportListResource) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise.
func (o *ImportListResource) GetQualityProfileId() int32 {
	if o == nil || IsNil(o.QualityProfileId) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityProfileId) {
		return nil, false
	}
	return o.QualityProfileId, true
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *ImportListResource) HasQualityProfileId() bool {
	if o != nil && !IsNil(o.QualityProfileId) {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given int32 and assigns it to the QualityProfileId field.
func (o *ImportListResource) SetQualityProfileId(v int32) {
	o.QualityProfileId = &v
}

// GetSeriesType returns the SeriesType field value if set, zero value otherwise.
func (o *ImportListResource) GetSeriesType() SeriesTypes {
	if o == nil || IsNil(o.SeriesType) {
		var ret SeriesTypes
		return ret
	}
	return *o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetSeriesTypeOk() (*SeriesTypes, bool) {
	if o == nil || IsNil(o.SeriesType) {
		return nil, false
	}
	return o.SeriesType, true
}

// HasSeriesType returns a boolean if a field has been set.
func (o *ImportListResource) HasSeriesType() bool {
	if o != nil && !IsNil(o.SeriesType) {
		return true
	}

	return false
}

// SetSeriesType gets a reference to the given SeriesTypes and assigns it to the SeriesType field.
func (o *ImportListResource) SetSeriesType(v SeriesTypes) {
	o.SeriesType = &v
}

// GetSeasonFolder returns the SeasonFolder field value if set, zero value otherwise.
func (o *ImportListResource) GetSeasonFolder() bool {
	if o == nil || IsNil(o.SeasonFolder) {
		var ret bool
		return ret
	}
	return *o.SeasonFolder
}

// GetSeasonFolderOk returns a tuple with the SeasonFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetSeasonFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.SeasonFolder) {
		return nil, false
	}
	return o.SeasonFolder, true
}

// HasSeasonFolder returns a boolean if a field has been set.
func (o *ImportListResource) HasSeasonFolder() bool {
	if o != nil && !IsNil(o.SeasonFolder) {
		return true
	}

	return false
}

// SetSeasonFolder gets a reference to the given bool and assigns it to the SeasonFolder field.
func (o *ImportListResource) SetSeasonFolder(v bool) {
	o.SeasonFolder = &v
}

// GetListType returns the ListType field value if set, zero value otherwise.
func (o *ImportListResource) GetListType() ImportListType {
	if o == nil || IsNil(o.ListType) {
		var ret ImportListType
		return ret
	}
	return *o.ListType
}

// GetListTypeOk returns a tuple with the ListType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetListTypeOk() (*ImportListType, bool) {
	if o == nil || IsNil(o.ListType) {
		return nil, false
	}
	return o.ListType, true
}

// HasListType returns a boolean if a field has been set.
func (o *ImportListResource) HasListType() bool {
	if o != nil && !IsNil(o.ListType) {
		return true
	}

	return false
}

// SetListType gets a reference to the given ImportListType and assigns it to the ListType field.
func (o *ImportListResource) SetListType(v ImportListType) {
	o.ListType = &v
}

// GetListOrder returns the ListOrder field value if set, zero value otherwise.
func (o *ImportListResource) GetListOrder() int32 {
	if o == nil || IsNil(o.ListOrder) {
		var ret int32
		return ret
	}
	return *o.ListOrder
}

// GetListOrderOk returns a tuple with the ListOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetListOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.ListOrder) {
		return nil, false
	}
	return o.ListOrder, true
}

// HasListOrder returns a boolean if a field has been set.
func (o *ImportListResource) HasListOrder() bool {
	if o != nil && !IsNil(o.ListOrder) {
		return true
	}

	return false
}

// SetListOrder gets a reference to the given int32 and assigns it to the ListOrder field.
func (o *ImportListResource) SetListOrder(v int32) {
	o.ListOrder = &v
}

// GetMinRefreshInterval returns the MinRefreshInterval field value if set, zero value otherwise.
func (o *ImportListResource) GetMinRefreshInterval() string {
	if o == nil || IsNil(o.MinRefreshInterval) {
		var ret string
		return ret
	}
	return *o.MinRefreshInterval
}

// GetMinRefreshIntervalOk returns a tuple with the MinRefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListResource) GetMinRefreshIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MinRefreshInterval) {
		return nil, false
	}
	return o.MinRefreshInterval, true
}

// HasMinRefreshInterval returns a boolean if a field has been set.
func (o *ImportListResource) HasMinRefreshInterval() bool {
	if o != nil && !IsNil(o.MinRefreshInterval) {
		return true
	}

	return false
}

// SetMinRefreshInterval gets a reference to the given string and assigns it to the MinRefreshInterval field.
func (o *ImportListResource) SetMinRefreshInterval(v string) {
	o.MinRefreshInterval = &v
}

func (o ImportListResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportListResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ConfigContract.IsSet() {
		toSerialize["configContract"] = o.ConfigContract.Get()
	}
	if o.InfoLink.IsSet() {
		toSerialize["infoLink"] = o.InfoLink.Get()
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if !IsNil(o.EnableAutomaticAdd) {
		toSerialize["enableAutomaticAdd"] = o.EnableAutomaticAdd
	}
	if !IsNil(o.SearchForMissingEpisodes) {
		toSerialize["searchForMissingEpisodes"] = o.SearchForMissingEpisodes
	}
	if !IsNil(o.ShouldMonitor) {
		toSerialize["shouldMonitor"] = o.ShouldMonitor
	}
	if !IsNil(o.MonitorNewItems) {
		toSerialize["monitorNewItems"] = o.MonitorNewItems
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if !IsNil(o.QualityProfileId) {
		toSerialize["qualityProfileId"] = o.QualityProfileId
	}
	if !IsNil(o.SeriesType) {
		toSerialize["seriesType"] = o.SeriesType
	}
	if !IsNil(o.SeasonFolder) {
		toSerialize["seasonFolder"] = o.SeasonFolder
	}
	if !IsNil(o.ListType) {
		toSerialize["listType"] = o.ListType
	}
	if !IsNil(o.ListOrder) {
		toSerialize["listOrder"] = o.ListOrder
	}
	if !IsNil(o.MinRefreshInterval) {
		toSerialize["minRefreshInterval"] = o.MinRefreshInterval
	}
	return toSerialize, nil
}

type NullableImportListResource struct {
	value *ImportListResource
	isSet bool
}

func (v NullableImportListResource) Get() *ImportListResource {
	return v.value
}

func (v *NullableImportListResource) Set(val *ImportListResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListResource(val *ImportListResource) *NullableImportListResource {
	return &NullableImportListResource{value: val, isSet: true}
}

func (v NullableImportListResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


