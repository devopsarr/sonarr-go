# \QueueAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteQueue**](QueueAPI.md#DeleteQueue) | **Delete** /api/v3/queue/{id} | 
[**DeleteQueueBulk**](QueueAPI.md#DeleteQueueBulk) | **Delete** /api/v3/queue/bulk | 
[**GetQueue**](QueueAPI.md#GetQueue) | **Get** /api/v3/queue | 



## DeleteQueue

> DeleteQueue(ctx, id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipRedownload(skipRedownload).ChangeCategory(changeCategory).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipRedownload := true // bool |  (optional) (default to false)
    changeCategory := true // bool |  (optional) (default to false)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueAPI.DeleteQueue(context.Background(), id).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipRedownload(skipRedownload).ChangeCategory(changeCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.DeleteQueue``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipRedownload** | **bool** |  | [default to false]
 **changeCategory** | **bool** |  | [default to false]

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


## DeleteQueueBulk

> DeleteQueueBulk(ctx).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipRedownload(skipRedownload).ChangeCategory(changeCategory).QueueBulkResource(queueBulkResource).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    skipRedownload := true // bool |  (optional) (default to false)
    changeCategory := true // bool |  (optional) (default to false)
    queueBulkResource := *sonarrClient.NewQueueBulkResource() // QueueBulkResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueAPI.DeleteQueueBulk(context.Background()).RemoveFromClient(removeFromClient).Blocklist(blocklist).SkipRedownload(skipRedownload).ChangeCategory(changeCategory).QueueBulkResource(queueBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.DeleteQueueBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **skipRedownload** | **bool** |  | [default to false]
 **changeCategory** | **bool** |  | [default to false]
 **queueBulkResource** | [**QueueBulkResource**](QueueBulkResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueue

> QueueResourcePagingResource GetQueue(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeUnknownSeriesItems(includeUnknownSeriesItems).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).SeriesIds(seriesIds).Protocol(protocol).Languages(languages).Quality(quality).Execute()



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
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 10)
    sortKey := "sortKey_example" // string |  (optional)
    sortDirection := sonarrClient.SortDirection("default") // SortDirection |  (optional)
    includeUnknownSeriesItems := true // bool |  (optional) (default to false)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisode := true // bool |  (optional) (default to false)
    seriesIds := []int32{int32(123)} // []int32 |  (optional)
    protocol := sonarrClient.DownloadProtocol("unknown") // DownloadProtocol |  (optional)
    languages := []int32{int32(123)} // []int32 |  (optional)
    quality := int32(56) // int32 |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueAPI.GetQueue(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeUnknownSeriesItems(includeUnknownSeriesItems).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).SeriesIds(seriesIds).Protocol(protocol).Languages(languages).Quality(quality).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueAPI.GetQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueue`: QueueResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `QueueAPI.GetQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeUnknownSeriesItems** | **bool** |  | [default to false]
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]
 **seriesIds** | **[]int32** |  | 
 **protocol** | [**DownloadProtocol**](DownloadProtocol.md) |  | 
 **languages** | **[]int32** |  | 
 **quality** | **int32** |  | 

### Return type

[**QueueResourcePagingResource**](QueueResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

