# \TagDetailsAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTagDetailById**](TagDetailsAPI.md#GetTagDetailById) | **Get** /api/v3/tag/detail/{id} | 
[**ListTagDetail**](TagDetailsAPI.md#ListTagDetail) | **Get** /api/v3/tag/detail | 



## GetTagDetailById

> TagDetailsResource GetTagDetailById(ctx, id).Execute()



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
    resp, r, err := apiClient.TagDetailsAPI.GetTagDetailById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagDetailsAPI.GetTagDetailById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagDetailById`: TagDetailsResource
    fmt.Fprintf(os.Stdout, "Response from `TagDetailsAPI.GetTagDetailById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagDetailByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagDetailsResource**](TagDetailsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagDetail

> []TagDetailsResource ListTagDetail(ctx).Execute()



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

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagDetailsAPI.ListTagDetail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagDetailsAPI.ListTagDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagDetail`: []TagDetailsResource
    fmt.Fprintf(os.Stdout, "Response from `TagDetailsAPI.ListTagDetail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTagDetailRequest struct via the builder pattern


### Return type

[**[]TagDetailsResource**](TagDetailsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

