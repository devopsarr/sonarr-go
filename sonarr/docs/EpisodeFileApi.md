# \EpisodeFileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEpisodefile**](EpisodeFileApi.md#DeleteEpisodefile) | **Delete** /api/v3/episodefile/{id} | 
[**DeleteEpisodefileBulk**](EpisodeFileApi.md#DeleteEpisodefileBulk) | **Delete** /api/v3/episodefile/bulk | 
[**GetEpisodefileById**](EpisodeFileApi.md#GetEpisodefileById) | **Get** /api/v3/episodefile/{id} | 
[**ListEpisodefile**](EpisodeFileApi.md#ListEpisodefile) | **Get** /api/v3/episodefile | 
[**PutEpisodefileBulk**](EpisodeFileApi.md#PutEpisodefileBulk) | **Put** /api/v3/episodefile/bulk | 
[**PutEpisodefileEditor**](EpisodeFileApi.md#PutEpisodefileEditor) | **Put** /api/v3/episodefile/editor | 
[**UpdateEpisodefile**](EpisodeFileApi.md#UpdateEpisodefile) | **Put** /api/v3/episodefile/{id} | 



## DeleteEpisodefile

> DeleteEpisodefile(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.DeleteEpisodefile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.DeleteEpisodefile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEpisodefileRequest struct via the builder pattern


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


## DeleteEpisodefileBulk

> DeleteEpisodefileBulk(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    episodeFileListResource := *sonarrClient.NewEpisodeFileListResource() // EpisodeFileListResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.DeleteEpisodefileBulk(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.DeleteEpisodefileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEpisodefileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileListResource** | [**EpisodeFileListResource**](EpisodeFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpisodefileById

> EpisodeFileResource GetEpisodefileById(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.GetEpisodefileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.GetEpisodefileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEpisodefileById`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.GetEpisodefileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpisodefileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpisodefile

> []EpisodeFileResource ListEpisodefile(ctx).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()



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
    seriesId := int32(56) // int32 |  (optional)
    episodeFileIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ListEpisodefile(context.Background()).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ListEpisodefile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEpisodefile`: []EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.ListEpisodefile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEpisodefileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **episodeFileIds** | **[]int32** |  | 

### Return type

[**[]EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEpisodefileBulk

> PutEpisodefileBulk(ctx).EpisodeFileResource(episodeFileResource).Execute()



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
    episodeFileResource := []sonarrClient.EpisodeFileResource{*sonarrClient.NewEpisodeFileResource()} // []EpisodeFileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.PutEpisodefileBulk(context.Background()).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.PutEpisodefileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEpisodefileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileResource** | [**[]EpisodeFileResource**](EpisodeFileResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEpisodefileEditor

> PutEpisodefileEditor(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    episodeFileListResource := *sonarrClient.NewEpisodeFileListResource() // EpisodeFileListResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.PutEpisodefileEditor(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.PutEpisodefileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEpisodefileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileListResource** | [**EpisodeFileListResource**](EpisodeFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEpisodefile

> EpisodeFileResource UpdateEpisodefile(ctx, id).EpisodeFileResource(episodeFileResource).Execute()



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
    episodeFileResource := *sonarrClient.NewEpisodeFileResource() // EpisodeFileResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.UpdateEpisodefile(context.Background(), id).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.UpdateEpisodefile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEpisodefile`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.UpdateEpisodefile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpisodefileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **episodeFileResource** | [**EpisodeFileResource**](EpisodeFileResource.md) |  | 

### Return type

[**EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

