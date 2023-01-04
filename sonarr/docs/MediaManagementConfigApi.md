# \MediaManagementConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigMediamanagement**](MediaManagementConfigApi.md#GetConfigMediamanagement) | **Get** /api/v3/config/mediamanagement | 
[**GetConfigMediamanagementById**](MediaManagementConfigApi.md#GetConfigMediamanagementById) | **Get** /api/v3/config/mediamanagement/{id} | 
[**UpdateConfigMediamanagement**](MediaManagementConfigApi.md#UpdateConfigMediamanagement) | **Put** /api/v3/config/mediamanagement/{id} | 



## GetConfigMediamanagement

> MediaManagementConfigResource GetConfigMediamanagement(ctx).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.GetConfigMediamanagement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetConfigMediamanagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigMediamanagement`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetConfigMediamanagement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigMediamanagementRequest struct via the builder pattern


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigMediamanagementById

> MediaManagementConfigResource GetConfigMediamanagementById(ctx, id).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.GetConfigMediamanagementById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.GetConfigMediamanagementById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigMediamanagementById`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.GetConfigMediamanagementById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigMediamanagementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigMediamanagement

> MediaManagementConfigResource UpdateConfigMediamanagement(ctx, id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.UpdateConfigMediamanagement(context.Background(), id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.UpdateConfigMediamanagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigMediamanagement`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.UpdateConfigMediamanagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigMediamanagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaManagementConfigResource** | [**MediaManagementConfigResource**](MediaManagementConfigResource.md) |  | 

### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

