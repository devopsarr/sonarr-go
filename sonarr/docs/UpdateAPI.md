# \UpdateAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUpdate**](UpdateAPI.md#ListUpdate) | **Get** /api/v3/update | 



## ListUpdate

> []UpdateResource ListUpdate(ctx).Execute()



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
    resp, r, err := apiClient.UpdateAPI.ListUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateAPI.ListUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUpdate`: []UpdateResource
    fmt.Fprintf(os.Stdout, "Response from `UpdateAPI.ListUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUpdateRequest struct via the builder pattern


### Return type

[**[]UpdateResource**](UpdateResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

