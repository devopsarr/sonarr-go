# \LanguageProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanguageprofile**](LanguageProfileApi.md#CreateLanguageprofile) | **Post** /api/v3/languageprofile | 
[**DeleteLanguageprofile**](LanguageProfileApi.md#DeleteLanguageprofile) | **Delete** /api/v3/languageprofile/{id} | 
[**GetLanguageprofileById**](LanguageProfileApi.md#GetLanguageprofileById) | **Get** /api/v3/languageprofile/{id} | 
[**ListLanguageprofile**](LanguageProfileApi.md#ListLanguageprofile) | **Get** /api/v3/languageprofile | 
[**UpdateLanguageprofile**](LanguageProfileApi.md#UpdateLanguageprofile) | **Put** /api/v3/languageprofile/{id} | 



## CreateLanguageprofile

> LanguageProfileResource CreateLanguageprofile(ctx).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *sonarrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.CreateLanguageprofile(context.Background()).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.CreateLanguageprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanguageprofile`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.CreateLanguageprofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanguageprofile

> DeleteLanguageprofile(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.DeleteLanguageprofile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.DeleteLanguageprofile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLanguageprofileRequest struct via the builder pattern


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


## GetLanguageprofileById

> LanguageProfileResource GetLanguageprofileById(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.GetLanguageprofileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.GetLanguageprofileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguageprofileById`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.GetLanguageprofileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageprofileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguageprofile

> []LanguageProfileResource ListLanguageprofile(ctx).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.ListLanguageprofile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ListLanguageprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLanguageprofile`: []LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ListLanguageprofile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLanguageprofileRequest struct via the builder pattern


### Return type

[**[]LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguageprofile

> LanguageProfileResource UpdateLanguageprofile(ctx, id).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *sonarrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.UpdateLanguageprofile(context.Background(), id).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.UpdateLanguageprofile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLanguageprofile`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.UpdateLanguageprofile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguageprofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

