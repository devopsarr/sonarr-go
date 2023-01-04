# \LanguageProfileSchemaApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLanguageprofileSchema**](LanguageProfileSchemaApi.md#GetLanguageprofileSchema) | **Get** /api/v3/languageprofile/schema | 
[**GetLanguageprofileSchemaById**](LanguageProfileSchemaApi.md#GetLanguageprofileSchemaById) | **Get** /api/v3/languageprofile/schema/{id} | 



## GetLanguageprofileSchema

> LanguageProfileResource GetLanguageprofileSchema(ctx).Execute()



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
    resp, r, err := apiClient.LanguageProfileSchemaApi.GetLanguageprofileSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileSchemaApi.GetLanguageprofileSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguageprofileSchema`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileSchemaApi.GetLanguageprofileSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageprofileSchemaRequest struct via the builder pattern


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


## GetLanguageprofileSchemaById

> LanguageProfileResource GetLanguageprofileSchemaById(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileSchemaApi.GetLanguageprofileSchemaById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileSchemaApi.GetLanguageprofileSchemaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguageprofileSchemaById`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileSchemaApi.GetLanguageprofileSchemaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageprofileSchemaByIdRequest struct via the builder pattern


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

