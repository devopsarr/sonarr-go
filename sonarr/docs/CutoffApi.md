# \CutoffApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWantedCutoff**](CutoffApi.md#GetWantedCutoff) | **Get** /api/v3/wanted/cutoff | 
[**GetWantedCutoffById**](CutoffApi.md#GetWantedCutoffById) | **Get** /api/v3/wanted/cutoff/{id} | 



## GetWantedCutoff

> EpisodeResourcePagingResource GetWantedCutoff(ctx).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Execute()



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
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisodeFile := true // bool |  (optional) (default to false)
    includeImages := true // bool |  (optional) (default to false)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CutoffApi.GetWantedCutoff(context.Background()).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Execute()
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
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisodeFile** | **bool** |  | [default to false]
 **includeImages** | **bool** |  | [default to false]

### Return type

[**EpisodeResourcePagingResource**](EpisodeResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

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

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

