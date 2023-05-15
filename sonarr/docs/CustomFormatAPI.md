# \CustomFormatAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomFormat**](CustomFormatAPI.md#CreateCustomFormat) | **Post** /api/v3/customformat | 
[**DeleteCustomFormat**](CustomFormatAPI.md#DeleteCustomFormat) | **Delete** /api/v3/customformat/{id} | 
[**GetCustomFormatById**](CustomFormatAPI.md#GetCustomFormatById) | **Get** /api/v3/customformat/{id} | 
[**GetCustomFormatSchema**](CustomFormatAPI.md#GetCustomFormatSchema) | **Get** /api/v3/customformat/schema | 
[**ListCustomFormat**](CustomFormatAPI.md#ListCustomFormat) | **Get** /api/v3/customformat | 
[**UpdateCustomFormat**](CustomFormatAPI.md#UpdateCustomFormat) | **Put** /api/v3/customformat/{id} | 



## CreateCustomFormat

> CustomFormatResource CreateCustomFormat(ctx).CustomFormatResource(customFormatResource).Execute()



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
    customFormatResource := *sonarrClient.NewCustomFormatResource() // CustomFormatResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormatAPI.CreateCustomFormat(context.Background()).CustomFormatResource(customFormatResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.CreateCustomFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomFormat`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatAPI.CreateCustomFormat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFormatResource** | [**CustomFormatResource**](CustomFormatResource.md) |  | 

### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomFormat

> DeleteCustomFormat(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFormatAPI.DeleteCustomFormat(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.DeleteCustomFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCustomFormatRequest struct via the builder pattern


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


## GetCustomFormatById

> CustomFormatResource GetCustomFormatById(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFormatAPI.GetCustomFormatById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.GetCustomFormatById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomFormatById`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatAPI.GetCustomFormatById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFormatByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFormatSchema

> GetCustomFormatSchema(ctx).Execute()



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
    resp, r, err := apiClient.CustomFormatAPI.GetCustomFormatSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.GetCustomFormatSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFormatSchemaRequest struct via the builder pattern


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


## ListCustomFormat

> []CustomFormatResource ListCustomFormat(ctx).Execute()



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
    resp, r, err := apiClient.CustomFormatAPI.ListCustomFormat(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.ListCustomFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomFormat`: []CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatAPI.ListCustomFormat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFormatRequest struct via the builder pattern


### Return type

[**[]CustomFormatResource**](CustomFormatResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomFormat

> CustomFormatResource UpdateCustomFormat(ctx, id).CustomFormatResource(customFormatResource).Execute()



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
    customFormatResource := *sonarrClient.NewCustomFormatResource() // CustomFormatResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormatAPI.UpdateCustomFormat(context.Background(), id).CustomFormatResource(customFormatResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatAPI.UpdateCustomFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomFormat`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatAPI.UpdateCustomFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFormatResource** | [**CustomFormatResource**](CustomFormatResource.md) |  | 

### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

