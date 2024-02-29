# \ImportListConfigAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportListConfig**](ImportListConfigAPI.md#GetImportListConfig) | **Get** /api/v3/config/importlist | 
[**GetImportListConfigById**](ImportListConfigAPI.md#GetImportListConfigById) | **Get** /api/v3/config/importlist/{id} | 
[**UpdateImportListConfig**](ImportListConfigAPI.md#UpdateImportListConfig) | **Put** /api/v3/config/importlist/{id} | 



## GetImportListConfig

> ImportListConfigResource GetImportListConfig(ctx).Execute()



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
	resp, r, err := apiClient.ImportListConfigAPI.GetImportListConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigAPI.GetImportListConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportListConfig`: ImportListConfigResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListConfigAPI.GetImportListConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportListConfigRequest struct via the builder pattern


### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportListConfigById

> ImportListConfigResource GetImportListConfigById(ctx, id).Execute()



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
	id := int32(56) // int32 | 

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListConfigAPI.GetImportListConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigAPI.GetImportListConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportListConfigById`: ImportListConfigResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListConfigAPI.GetImportListConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportListConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImportListConfig

> ImportListConfigResource UpdateImportListConfig(ctx, id).ImportListConfigResource(importListConfigResource).Execute()



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
	id := "id_example" // string | 
	importListConfigResource := *sonarrClient.NewImportListConfigResource() // ImportListConfigResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListConfigAPI.UpdateImportListConfig(context.Background(), id).ImportListConfigResource(importListConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListConfigAPI.UpdateImportListConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImportListConfig`: ImportListConfigResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListConfigAPI.UpdateImportListConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImportListConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListConfigResource** | [**ImportListConfigResource**](ImportListConfigResource.md) |  | 

### Return type

[**ImportListConfigResource**](ImportListConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

