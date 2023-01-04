# EpisodeResourcePagingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**SortKey** | Pointer to **NullableString** |  | [optional] 
**SortDirection** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] 
**Filters** | Pointer to [**[]PagingResourceFilter**](PagingResourceFilter.md) |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]EpisodeResource**](EpisodeResource.md) |  | [optional] 

## Methods

### NewEpisodeResourcePagingResource

`func NewEpisodeResourcePagingResource() *EpisodeResourcePagingResource`

NewEpisodeResourcePagingResource instantiates a new EpisodeResourcePagingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeResourcePagingResourceWithDefaults

`func NewEpisodeResourcePagingResourceWithDefaults() *EpisodeResourcePagingResource`

NewEpisodeResourcePagingResourceWithDefaults instantiates a new EpisodeResourcePagingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *EpisodeResourcePagingResource) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EpisodeResourcePagingResource) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EpisodeResourcePagingResource) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *EpisodeResourcePagingResource) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *EpisodeResourcePagingResource) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EpisodeResourcePagingResource) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EpisodeResourcePagingResource) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EpisodeResourcePagingResource) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *EpisodeResourcePagingResource) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *EpisodeResourcePagingResource) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *EpisodeResourcePagingResource) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *EpisodeResourcePagingResource) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### SetSortKeyNil

`func (o *EpisodeResourcePagingResource) SetSortKeyNil(b bool)`

 SetSortKeyNil sets the value for SortKey to be an explicit nil

### UnsetSortKey
`func (o *EpisodeResourcePagingResource) UnsetSortKey()`

UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
### GetSortDirection

`func (o *EpisodeResourcePagingResource) GetSortDirection() SortDirection`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *EpisodeResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *EpisodeResourcePagingResource) SetSortDirection(v SortDirection)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *EpisodeResourcePagingResource) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetFilters

`func (o *EpisodeResourcePagingResource) GetFilters() []PagingResourceFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EpisodeResourcePagingResource) GetFiltersOk() (*[]PagingResourceFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EpisodeResourcePagingResource) SetFilters(v []PagingResourceFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EpisodeResourcePagingResource) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *EpisodeResourcePagingResource) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *EpisodeResourcePagingResource) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetTotalRecords

`func (o *EpisodeResourcePagingResource) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *EpisodeResourcePagingResource) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *EpisodeResourcePagingResource) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *EpisodeResourcePagingResource) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetRecords

`func (o *EpisodeResourcePagingResource) GetRecords() []EpisodeResource`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *EpisodeResourcePagingResource) GetRecordsOk() (*[]EpisodeResource, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *EpisodeResourcePagingResource) SetRecords(v []EpisodeResource)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *EpisodeResourcePagingResource) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *EpisodeResourcePagingResource) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *EpisodeResourcePagingResource) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


