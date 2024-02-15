# \ImportListExclusionAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImportListExclusion**](ImportListExclusionAPI.md#CreateImportListExclusion) | **Post** /api/v3/importlistexclusion | 
[**DeleteImportListExclusion**](ImportListExclusionAPI.md#DeleteImportListExclusion) | **Delete** /api/v3/importlistexclusion/{id} | 
[**GetImportListExclusionById**](ImportListExclusionAPI.md#GetImportListExclusionById) | **Get** /api/v3/importlistexclusion/{id} | 
[**ListImportListExclusion**](ImportListExclusionAPI.md#ListImportListExclusion) | **Get** /api/v3/importlistexclusion | 
[**UpdateImportListExclusion**](ImportListExclusionAPI.md#UpdateImportListExclusion) | **Put** /api/v3/importlistexclusion/{id} | 



## CreateImportListExclusion

> ImportListExclusionResource CreateImportListExclusion(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
	importListExclusionResource := *sonarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListExclusionAPI.CreateImportListExclusion(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.CreateImportListExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImportListExclusion`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.CreateImportListExclusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImportListExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImportListExclusion

> DeleteImportListExclusion(ctx, id).Execute()



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
	r, err := apiClient.ImportListExclusionAPI.DeleteImportListExclusion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.DeleteImportListExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImportListExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportListExclusionById

> ImportListExclusionResource GetImportListExclusionById(ctx, id).Execute()



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
	resp, r, err := apiClient.ImportListExclusionAPI.GetImportListExclusionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.GetImportListExclusionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportListExclusionById`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.GetImportListExclusionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportListExclusionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportListExclusion

> []ImportListExclusionResource ListImportListExclusion(ctx).Execute()



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
	resp, r, err := apiClient.ImportListExclusionAPI.ListImportListExclusion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.ListImportListExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImportListExclusion`: []ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.ListImportListExclusion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListImportListExclusionRequest struct via the builder pattern


### Return type

[**[]ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImportListExclusion

> ImportListExclusionResource UpdateImportListExclusion(ctx, id).ImportListExclusionResource(importListExclusionResource).Execute()



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
	importListExclusionResource := *sonarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListExclusionAPI.UpdateImportListExclusion(context.Background(), id).ImportListExclusionResource(importListExclusionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.UpdateImportListExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImportListExclusion`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.UpdateImportListExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImportListExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

