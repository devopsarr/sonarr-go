# \DownloadClientConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigDownloadclient**](DownloadClientConfigApi.md#GetConfigDownloadclient) | **Get** /api/v3/config/downloadclient | 
[**GetConfigDownloadclientById**](DownloadClientConfigApi.md#GetConfigDownloadclientById) | **Get** /api/v3/config/downloadclient/{id} | 
[**UpdateConfigDownloadclient**](DownloadClientConfigApi.md#UpdateConfigDownloadclient) | **Put** /api/v3/config/downloadclient/{id} | 



## GetConfigDownloadclient

> DownloadClientConfigResource GetConfigDownloadclient(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.GetConfigDownloadclient(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.GetConfigDownloadclient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigDownloadclient`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.GetConfigDownloadclient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigDownloadclientRequest struct via the builder pattern


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigDownloadclientById

> DownloadClientConfigResource GetConfigDownloadclientById(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.GetConfigDownloadclientById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.GetConfigDownloadclientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigDownloadclientById`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.GetConfigDownloadclientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigDownloadclientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigDownloadclient

> DownloadClientConfigResource UpdateConfigDownloadclient(ctx, id).DownloadClientConfigResource(downloadClientConfigResource).Execute()



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
    downloadClientConfigResource := *sonarrClient.NewDownloadClientConfigResource() // DownloadClientConfigResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientConfigApi.UpdateConfigDownloadclient(context.Background(), id).DownloadClientConfigResource(downloadClientConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.UpdateConfigDownloadclient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigDownloadclient`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.UpdateConfigDownloadclient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigDownloadclientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientConfigResource** | [**DownloadClientConfigResource**](DownloadClientConfigResource.md) |  | 

### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

