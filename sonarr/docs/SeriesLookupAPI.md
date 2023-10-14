# \SeriesLookupAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSeriesLookup**](SeriesLookupAPI.md#ListSeriesLookup) | **Get** /api/v3/series/lookup | 



## ListSeriesLookup

> []SeriesResource ListSeriesLookup(ctx).Term(term).Execute()



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
    resp, r, err := apiClient.SeriesLookupAPI.ListSeriesLookup(context.Background()).Term(term).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesLookupAPI.ListSeriesLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSeriesLookup`: []SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesLookupAPI.ListSeriesLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSeriesLookupRequest struct via the builder pattern


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

