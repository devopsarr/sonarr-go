# \SeriesEditorAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSeriesEditor**](SeriesEditorAPI.md#DeleteSeriesEditor) | **Delete** /api/v3/series/editor | 
[**PutSeriesEditor**](SeriesEditorAPI.md#PutSeriesEditor) | **Put** /api/v3/series/editor | 



## DeleteSeriesEditor

> DeleteSeriesEditor(ctx).SeriesEditorResource(seriesEditorResource).Execute()



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
    seriesEditorResource := *sonarrClient.NewSeriesEditorResource() // SeriesEditorResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesEditorAPI.DeleteSeriesEditor(context.Background()).SeriesEditorResource(seriesEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesEditorAPI.DeleteSeriesEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeriesEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesEditorResource** | [**SeriesEditorResource**](SeriesEditorResource.md) |  | 

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


## PutSeriesEditor

> PutSeriesEditor(ctx).SeriesEditorResource(seriesEditorResource).Execute()



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
    seriesEditorResource := *sonarrClient.NewSeriesEditorResource() // SeriesEditorResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesEditorAPI.PutSeriesEditor(context.Background()).SeriesEditorResource(seriesEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesEditorAPI.PutSeriesEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSeriesEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesEditorResource** | [**SeriesEditorResource**](SeriesEditorResource.md) |  | 

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

