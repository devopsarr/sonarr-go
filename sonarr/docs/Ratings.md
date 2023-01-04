# Ratings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Votes** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewRatings

`func NewRatings() *Ratings`

NewRatings instantiates a new Ratings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingsWithDefaults

`func NewRatingsWithDefaults() *Ratings`

NewRatingsWithDefaults instantiates a new Ratings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVotes

`func (o *Ratings) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *Ratings) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *Ratings) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *Ratings) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetValue

`func (o *Ratings) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Ratings) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Ratings) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Ratings) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


