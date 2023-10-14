# \CustomFilterAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomFilter**](CustomFilterAPI.md#CreateCustomFilter) | **Post** /api/v3/customfilter | 
[**DeleteCustomFilter**](CustomFilterAPI.md#DeleteCustomFilter) | **Delete** /api/v3/customfilter/{id} | 
[**GetCustomFilterById**](CustomFilterAPI.md#GetCustomFilterById) | **Get** /api/v3/customfilter/{id} | 
[**ListCustomFilter**](CustomFilterAPI.md#ListCustomFilter) | **Get** /api/v3/customfilter | 
[**UpdateCustomFilter**](CustomFilterAPI.md#UpdateCustomFilter) | **Put** /api/v3/customfilter/{id} | 



## CreateCustomFilter

> CustomFilterResource CreateCustomFilter(ctx).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *sonarrClient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterAPI.CreateCustomFilter(context.Background()).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterAPI.CreateCustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomFilter`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterAPI.CreateCustomFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomFilter

> DeleteCustomFilter(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterAPI.DeleteCustomFilter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterAPI.DeleteCustomFilter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCustomFilterRequest struct via the builder pattern


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


## GetCustomFilterById

> CustomFilterResource GetCustomFilterById(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterAPI.GetCustomFilterById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterAPI.GetCustomFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomFilterById`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterAPI.GetCustomFilterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFilter

> []CustomFilterResource ListCustomFilter(ctx).Execute()



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
    resp, r, err := apiClient.CustomFilterAPI.ListCustomFilter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterAPI.ListCustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomFilter`: []CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterAPI.ListCustomFilter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFilterRequest struct via the builder pattern


### Return type

[**[]CustomFilterResource**](CustomFilterResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomFilter

> CustomFilterResource UpdateCustomFilter(ctx, id).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *sonarrClient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterAPI.UpdateCustomFilter(context.Background(), id).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterAPI.UpdateCustomFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomFilter`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterAPI.UpdateCustomFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

