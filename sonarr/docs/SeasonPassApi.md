# \SeasonPassApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSeasonpass**](SeasonPassApi.md#CreateSeasonpass) | **Post** /api/v3/seasonpass | 



## CreateSeasonpass

> CreateSeasonpass(ctx).SeasonPassResource(seasonPassResource).Execute()



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
    seasonPassResource := *sonarrClient.NewSeasonPassResource() // SeasonPassResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeasonPassApi.CreateSeasonpass(context.Background()).SeasonPassResource(seasonPassResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeasonPassApi.CreateSeasonpass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeasonpassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seasonPassResource** | [**SeasonPassResource**](SeasonPassResource.md) |  | 

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

