# \DownloadClientConfigAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadClientConfig**](DownloadClientConfigAPI.md#GetDownloadClientConfig) | **Get** /api/v3/config/downloadclient | 
[**GetDownloadClientConfigById**](DownloadClientConfigAPI.md#GetDownloadClientConfigById) | **Get** /api/v3/config/downloadclient/{id} | 
[**UpdateDownloadClientConfig**](DownloadClientConfigAPI.md#UpdateDownloadClientConfig) | **Put** /api/v3/config/downloadclient/{id} | 



## GetDownloadClientConfig

> DownloadClientConfigResource GetDownloadClientConfig(ctx).Execute()



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
	resp, r, err := apiClient.DownloadClientConfigAPI.GetDownloadClientConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigAPI.GetDownloadClientConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownloadClientConfig`: DownloadClientConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigAPI.GetDownloadClientConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientConfigRequest struct via the builder pattern


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadClientConfigById

> DownloadClientConfigResource GetDownloadClientConfigById(ctx, id).Execute()



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
	resp, r, err := apiClient.DownloadClientConfigAPI.GetDownloadClientConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigAPI.GetDownloadClientConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownloadClientConfigById`: DownloadClientConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigAPI.GetDownloadClientConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDownloadClientConfig

> DownloadClientConfigResource UpdateDownloadClientConfig(ctx, id).DownloadClientConfigResource(downloadClientConfigResource).Execute()



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
	downloadClientConfigResource := *sonarrClient.NewDownloadClientConfigResource() // DownloadClientConfigResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadClientConfigAPI.UpdateDownloadClientConfig(context.Background(), id).DownloadClientConfigResource(downloadClientConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigAPI.UpdateDownloadClientConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDownloadClientConfig`: DownloadClientConfigResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigAPI.UpdateDownloadClientConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDownloadClientConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientConfigResource** | [**DownloadClientConfigResource**](DownloadClientConfigResource.md) |  | 

### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

