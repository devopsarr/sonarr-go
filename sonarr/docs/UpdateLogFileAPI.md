# \UpdateLogFileAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogFileUpdateByFilename**](UpdateLogFileAPI.md#GetLogFileUpdateByFilename) | **Get** /api/v3/log/file/update/{filename} | 
[**ListLogFileUpdate**](UpdateLogFileAPI.md#ListLogFileUpdate) | **Get** /api/v3/log/file/update | 



## GetLogFileUpdateByFilename

> GetLogFileUpdateByFilename(ctx, filename).Execute()



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
    filename := "filename_example" // string | 

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateLogFileAPI.GetLogFileUpdateByFilename(context.Background(), filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateLogFileAPI.GetLogFileUpdateByFilename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFileUpdateByFilenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogFileUpdate

> []LogFileResource ListLogFileUpdate(ctx).Execute()



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
    resp, r, err := apiClient.UpdateLogFileAPI.ListLogFileUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateLogFileAPI.ListLogFileUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogFileUpdate`: []LogFileResource
    fmt.Fprintf(os.Stdout, "Response from `UpdateLogFileAPI.ListLogFileUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogFileUpdateRequest struct via the builder pattern


### Return type

[**[]LogFileResource**](LogFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

