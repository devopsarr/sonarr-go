# \DelayProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelayprofile**](DelayProfileApi.md#CreateDelayprofile) | **Post** /api/v3/delayprofile | 
[**DeleteDelayprofile**](DelayProfileApi.md#DeleteDelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
[**GetDelayprofileById**](DelayProfileApi.md#GetDelayprofileById) | **Get** /api/v3/delayprofile/{id} | 
[**ListDelayprofile**](DelayProfileApi.md#ListDelayprofile) | **Get** /api/v3/delayprofile | 
[**UpdateDelayprofile**](DelayProfileApi.md#UpdateDelayprofile) | **Put** /api/v3/delayprofile/{id} | 
[**UpdateDelayprofileReorder**](DelayProfileApi.md#UpdateDelayprofileReorder) | **Put** /api/v3/delayprofile/reorder/{id} | 



## CreateDelayprofile

> DelayProfileResource CreateDelayprofile(ctx).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *sonarrClient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.CreateDelayprofile(context.Background()).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.CreateDelayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDelayprofile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.CreateDelayprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDelayprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelayprofile

> DeleteDelayprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.DeleteDelayprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.DeleteDelayprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDelayprofileRequest struct via the builder pattern


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


## GetDelayprofileById

> DelayProfileResource GetDelayprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.GetDelayprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.GetDelayprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelayprofileById`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.GetDelayprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelayprofileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelayprofile

> []DelayProfileResource ListDelayprofile(ctx).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ListDelayprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ListDelayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelayprofile`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ListDelayprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDelayprofileRequest struct via the builder pattern


### Return type

[**[]DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelayprofile

> DelayProfileResource UpdateDelayprofile(ctx, id).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *sonarrClient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.UpdateDelayprofile(context.Background(), id).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateDelayprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelayprofile`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.UpdateDelayprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelayprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelayprofileReorder

> []DelayProfileResource UpdateDelayprofileReorder(ctx, id).After(after).Execute()



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
    after := int32(56) // int32 |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.UpdateDelayprofileReorder(context.Background(), id).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.UpdateDelayprofileReorder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelayprofileReorder`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.UpdateDelayprofileReorder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelayprofileReorderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** |  | 

### Return type

[**[]DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

