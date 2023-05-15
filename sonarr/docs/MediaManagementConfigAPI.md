# \MediaManagementConfigAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMediaManagementConfig**](MediaManagementConfigAPI.md#GetMediaManagementConfig) | **Get** /api/v3/config/mediamanagement | 
[**GetMediaManagementConfigById**](MediaManagementConfigAPI.md#GetMediaManagementConfigById) | **Get** /api/v3/config/mediamanagement/{id} | 
[**UpdateMediaManagementConfig**](MediaManagementConfigAPI.md#UpdateMediaManagementConfig) | **Put** /api/v3/config/mediamanagement/{id} | 



## GetMediaManagementConfig

> MediaManagementConfigResource GetMediaManagementConfig(ctx).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigAPI.GetMediaManagementConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigAPI.GetMediaManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaManagementConfig`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigAPI.GetMediaManagementConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaManagementConfigRequest struct via the builder pattern


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaManagementConfigById

> MediaManagementConfigResource GetMediaManagementConfigById(ctx, id).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigAPI.GetMediaManagementConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigAPI.GetMediaManagementConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaManagementConfigById`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigAPI.GetMediaManagementConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaManagementConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMediaManagementConfig

> MediaManagementConfigResource UpdateMediaManagementConfig(ctx, id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()



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
    id := "id_example" // string | 
    mediaManagementConfigResource := *sonarrClient.NewMediaManagementConfigResource() // MediaManagementConfigResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigAPI.UpdateMediaManagementConfig(context.Background(), id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigAPI.UpdateMediaManagementConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMediaManagementConfig`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigAPI.UpdateMediaManagementConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaManagementConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaManagementConfigResource** | [**MediaManagementConfigResource**](MediaManagementConfigResource.md) |  | 

### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

