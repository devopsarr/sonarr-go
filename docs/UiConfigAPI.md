# \UiConfigAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUiConfig**](UiConfigAPI.md#GetUiConfig) | **Get** /api/v3/config/ui | 
[**GetUiConfigById**](UiConfigAPI.md#GetUiConfigById) | **Get** /api/v3/config/ui/{id} | 
[**UpdateUiConfig**](UiConfigAPI.md#UpdateUiConfig) | **Put** /api/v3/config/ui/{id} | 



## GetUiConfig

> UiConfigResource GetUiConfig(ctx).Execute()



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
	resp, r, err := apiClient.UiConfigAPI.GetUiConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UiConfigAPI.GetUiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiConfig`: UiConfigResource
	fmt.Fprintf(os.Stdout, "Response from `UiConfigAPI.GetUiConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiConfigRequest struct via the builder pattern


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUiConfigById

> UiConfigResource GetUiConfigById(ctx, id).Execute()



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
	resp, r, err := apiClient.UiConfigAPI.GetUiConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UiConfigAPI.GetUiConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiConfigById`: UiConfigResource
	fmt.Fprintf(os.Stdout, "Response from `UiConfigAPI.GetUiConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUiConfig

> UiConfigResource UpdateUiConfig(ctx, id).UiConfigResource(uiConfigResource).Execute()



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
	uiConfigResource := *sonarrClient.NewUiConfigResource() // UiConfigResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.UiConfigAPI.UpdateUiConfig(context.Background(), id).UiConfigResource(uiConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UiConfigAPI.UpdateUiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUiConfig`: UiConfigResource
	fmt.Fprintf(os.Stdout, "Response from `UiConfigAPI.UpdateUiConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUiConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiConfigResource** | [**UiConfigResource**](UiConfigResource.md) |  | 

### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

