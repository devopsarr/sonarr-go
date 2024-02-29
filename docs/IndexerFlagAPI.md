# \IndexerFlagAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListIndexerFlag**](IndexerFlagAPI.md#ListIndexerFlag) | **Get** /api/v3/indexerflag | 



## ListIndexerFlag

> []IndexerFlagResource ListIndexerFlag(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonarrClient "github.com/devopsarr/sonarr-go/sonarr"
)

func main() {

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexerFlagAPI.ListIndexerFlag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexerFlagAPI.ListIndexerFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndexerFlag`: []IndexerFlagResource
	fmt.Fprintf(os.Stdout, "Response from `IndexerFlagAPI.ListIndexerFlag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerFlagRequest struct via the builder pattern


### Return type

[**[]IndexerFlagResource**](IndexerFlagResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

