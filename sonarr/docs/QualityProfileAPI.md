# \QualityProfileAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQualityProfile**](QualityProfileAPI.md#CreateQualityProfile) | **Post** /api/v3/qualityprofile | 
[**DeleteQualityProfile**](QualityProfileAPI.md#DeleteQualityProfile) | **Delete** /api/v3/qualityprofile/{id} | 
[**GetQualityProfileById**](QualityProfileAPI.md#GetQualityProfileById) | **Get** /api/v3/qualityprofile/{id} | 
[**ListQualityProfile**](QualityProfileAPI.md#ListQualityProfile) | **Get** /api/v3/qualityprofile | 
[**UpdateQualityProfile**](QualityProfileAPI.md#UpdateQualityProfile) | **Put** /api/v3/qualityprofile/{id} | 



## CreateQualityProfile

> QualityProfileResource CreateQualityProfile(ctx).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *sonarrClient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileAPI.CreateQualityProfile(context.Background()).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileAPI.CreateQualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQualityProfile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileAPI.CreateQualityProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQualityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQualityProfile

> DeleteQualityProfile(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileAPI.DeleteQualityProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileAPI.DeleteQualityProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteQualityProfileRequest struct via the builder pattern


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


## GetQualityProfileById

> QualityProfileResource GetQualityProfileById(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileAPI.GetQualityProfileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileAPI.GetQualityProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQualityProfileById`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileAPI.GetQualityProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQualityProfile

> []QualityProfileResource ListQualityProfile(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileAPI.ListQualityProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileAPI.ListQualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQualityProfile`: []QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileAPI.ListQualityProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListQualityProfileRequest struct via the builder pattern


### Return type

[**[]QualityProfileResource**](QualityProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQualityProfile

> QualityProfileResource UpdateQualityProfile(ctx, id).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *sonarrClient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileAPI.UpdateQualityProfile(context.Background(), id).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileAPI.UpdateQualityProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQualityProfile`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileAPI.UpdateQualityProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQualityProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

