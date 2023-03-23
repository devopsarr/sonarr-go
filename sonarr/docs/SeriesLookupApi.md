# \SeriesLookupApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSeriesLookup**](SeriesLookupApi.md#GetSeriesLookup) | **Get** /api/v3/series/lookup | 



## GetSeriesLookup

> []SeriesResource GetSeriesLookup(ctx).Term(term).Execute()



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
    term := "term_example" // string |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesLookupApi.GetSeriesLookup(context.Background()).Term(term).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesLookupApi.GetSeriesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeriesLookup`: []SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesLookupApi.GetSeriesLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

### Return type

[**[]SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

