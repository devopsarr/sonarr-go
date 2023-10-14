# \LogAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLog**](LogAPI.md#GetLog) | **Get** /api/v3/log | 



## GetLog

> LogResourcePagingResource GetLog(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Level(level).Execute()



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
    level := "level_example" // string |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogAPI.GetLog(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLog`: LogResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **level** | **string** |  | 

### Return type

[**LogResourcePagingResource**](LogResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

