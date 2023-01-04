# \RemotePathMappingApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemotepathmapping**](RemotePathMappingApi.md#CreateRemotepathmapping) | **Post** /api/v3/remotepathmapping | 
[**DeleteRemotepathmapping**](RemotePathMappingApi.md#DeleteRemotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
[**GetRemotepathmappingById**](RemotePathMappingApi.md#GetRemotepathmappingById) | **Get** /api/v3/remotepathmapping/{id} | 
[**ListRemotepathmapping**](RemotePathMappingApi.md#ListRemotepathmapping) | **Get** /api/v3/remotepathmapping | 
[**UpdateRemotepathmapping**](RemotePathMappingApi.md#UpdateRemotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 



## CreateRemotepathmapping

> RemotePathMappingResource CreateRemotepathmapping(ctx).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    remotePathMappingResource := *sonarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.CreateRemotepathmapping(context.Background()).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.CreateRemotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemotepathmapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.CreateRemotepathmapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemotepathmappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemotepathmapping

> DeleteRemotepathmapping(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.DeleteRemotepathmapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.DeleteRemotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemotepathmappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemotepathmappingById

> RemotePathMappingResource GetRemotepathmappingById(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.GetRemotepathmappingById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.GetRemotepathmappingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemotepathmappingById`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.GetRemotepathmappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemotepathmappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemotepathmapping

> []RemotePathMappingResource ListRemotepathmapping(ctx).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.ListRemotepathmapping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ListRemotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemotepathmapping`: []RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ListRemotepathmapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRemotepathmappingRequest struct via the builder pattern


### Return type

[**[]RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemotepathmapping

> RemotePathMappingResource UpdateRemotepathmapping(ctx, id).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    remotePathMappingResource := *sonarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.UpdateRemotepathmapping(context.Background(), id).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.UpdateRemotepathmapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemotepathmapping`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.UpdateRemotepathmapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemotepathmappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

