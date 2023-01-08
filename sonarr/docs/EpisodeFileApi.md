# \EpisodeFileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEpisodeFile**](EpisodeFileApi.md#DeleteEpisodeFile) | **Delete** /api/v3/episodefile/{id} | 
[**DeleteEpisodeFileBulk**](EpisodeFileApi.md#DeleteEpisodeFileBulk) | **Delete** /api/v3/episodefile/bulk | 
[**GetEpisodeFileById**](EpisodeFileApi.md#GetEpisodeFileById) | **Get** /api/v3/episodefile/{id} | 
[**ListEpisodeFile**](EpisodeFileApi.md#ListEpisodeFile) | **Get** /api/v3/episodefile | 
[**PutEpisodeFileBulk**](EpisodeFileApi.md#PutEpisodeFileBulk) | **Put** /api/v3/episodefile/bulk | 
[**PutEpisodeFileEditor**](EpisodeFileApi.md#PutEpisodeFileEditor) | **Put** /api/v3/episodefile/editor | 
[**UpdateEpisodeFile**](EpisodeFileApi.md#UpdateEpisodeFile) | **Put** /api/v3/episodefile/{id} | 



## DeleteEpisodeFile

> DeleteEpisodeFile(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.DeleteEpisodeFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.DeleteEpisodeFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEpisodeFileRequest struct via the builder pattern


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


## DeleteEpisodeFileBulk

> DeleteEpisodeFileBulk(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.DeleteEpisodeFileBulk(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.DeleteEpisodeFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEpisodeFileBulkRequest struct via the builder pattern


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


## GetEpisodeFileById

> EpisodeFileResource GetEpisodeFileById(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.GetEpisodeFileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.GetEpisodeFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEpisodeFileById`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.GetEpisodeFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpisodeFileByIdRequest struct via the builder pattern


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


## ListEpisodeFile

> []EpisodeFileResource ListEpisodeFile(ctx).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.ListEpisodeFile(context.Background()).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ListEpisodeFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEpisodeFile`: []EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.ListEpisodeFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEpisodeFileRequest struct via the builder pattern


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


## PutEpisodeFileBulk

> PutEpisodeFileBulk(ctx).EpisodeFileResource(episodeFileResource).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.PutEpisodeFileBulk(context.Background()).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.PutEpisodeFileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEpisodeFileBulkRequest struct via the builder pattern


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


## PutEpisodeFileEditor

> PutEpisodeFileEditor(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.PutEpisodeFileEditor(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.PutEpisodeFileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEpisodeFileEditorRequest struct via the builder pattern


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


## UpdateEpisodeFile

> EpisodeFileResource UpdateEpisodeFile(ctx, id).EpisodeFileResource(episodeFileResource).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.UpdateEpisodeFile(context.Background(), id).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.UpdateEpisodeFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEpisodeFile`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.UpdateEpisodeFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpisodeFileRequest struct via the builder pattern


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

