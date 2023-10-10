# \CutoffApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWantedCutoff**](CutoffApi.md#GetWantedCutoff) | **Get** /api/v3/wanted/cutoff | 
[**GetWantedCutoffById**](CutoffApi.md#GetWantedCutoffById) | **Get** /api/v3/wanted/cutoff/{id} | 



## GetWantedCutoff

> EpisodeResourcePagingResource GetWantedCutoff(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Monitored(monitored).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sonarrClient "./openapi"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 10)
    sortKey := "sortKey_example" // string |  (optional)
    sortDirection := sonarrClient.SortDirection("default") // SortDirection |  (optional)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisodeFile := true // bool |  (optional) (default to false)
    includeImages := true // bool |  (optional) (default to false)
    monitored := true // bool |  (optional) (default to true)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CutoffApi.GetWantedCutoff(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Monitored(monitored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.GetWantedCutoff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWantedCutoff`: EpisodeResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.GetWantedCutoff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedCutoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisodeFile** | **bool** |  | [default to false]
 **includeImages** | **bool** |  | [default to false]
 **monitored** | **bool** |  | [default to true]

### Return type

[**EpisodeResourcePagingResource**](EpisodeResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWantedCutoffById

> EpisodeResource GetWantedCutoffById(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sonarrClient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CutoffApi.GetWantedCutoffById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.GetWantedCutoffById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWantedCutoffById`: EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.GetWantedCutoffById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedCutoffByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeResource**](EpisodeResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

