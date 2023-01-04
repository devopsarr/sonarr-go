# \RenameEpisodeApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRename**](RenameEpisodeApi.md#ListRename) | **Get** /api/v3/rename | 



## ListRename

> []RenameEpisodeResource ListRename(ctx).SeriesId(seriesId).SeasonNumber(seasonNumber).Execute()



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
    seasonNumber := int32(56) // int32 |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RenameEpisodeApi.ListRename(context.Background()).SeriesId(seriesId).SeasonNumber(seasonNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RenameEpisodeApi.ListRename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRename`: []RenameEpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `RenameEpisodeApi.ListRename`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 

### Return type

[**[]RenameEpisodeResource**](RenameEpisodeResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

