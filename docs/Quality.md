# Quality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to [**QualitySource**](QualitySource.md) |  | [optional] 
**Resolution** | Pointer to **int32** |  | [optional] 

## Methods

### NewQuality

`func NewQuality() *Quality`

NewQuality instantiates a new Quality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityWithDefaults

`func NewQualityWithDefaults() *Quality`

NewQualityWithDefaults instantiates a new Quality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Quality) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quality) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quality) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Quality) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Quality) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Quality) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Quality) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Quality) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Quality) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Quality) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSource

`func (o *Quality) GetSource() QualitySource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Quality) GetSourceOk() (*QualitySource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Quality) SetSource(v QualitySource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Quality) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetResolution

`func (o *Quality) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Quality) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Quality) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Quality) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


