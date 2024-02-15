# \RemotePathMappingAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemotePathMapping**](RemotePathMappingAPI.md#CreateRemotePathMapping) | **Post** /api/v3/remotepathmapping | 
[**DeleteRemotePathMapping**](RemotePathMappingAPI.md#DeleteRemotePathMapping) | **Delete** /api/v3/remotepathmapping/{id} | 
[**GetRemotePathMappingById**](RemotePathMappingAPI.md#GetRemotePathMappingById) | **Get** /api/v3/remotepathmapping/{id} | 
[**ListRemotePathMapping**](RemotePathMappingAPI.md#ListRemotePathMapping) | **Get** /api/v3/remotepathmapping | 
[**UpdateRemotePathMapping**](RemotePathMappingAPI.md#UpdateRemotePathMapping) | **Put** /api/v3/remotepathmapping/{id} | 



## CreateRemotePathMapping

> RemotePathMappingResource CreateRemotePathMapping(ctx).RemotePathMappingResource(remotePathMappingResource).Execute()



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
	remotePathMappingResource := *sonarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotePathMappingAPI.CreateRemotePathMapping(context.Background()).RemotePathMappingResource(remotePathMappingResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingAPI.CreateRemotePathMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRemotePathMapping`: RemotePathMappingResource
	fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingAPI.CreateRemotePathMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemotePathMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemotePathMapping

> DeleteRemotePathMapping(ctx, id).Execute()



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
	r, err := apiClient.RemotePathMappingAPI.DeleteRemotePathMapping(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingAPI.DeleteRemotePathMapping``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRemotePathMappingRequest struct via the builder pattern


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


## GetRemotePathMappingById

> RemotePathMappingResource GetRemotePathMappingById(ctx, id).Execute()



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
	resp, r, err := apiClient.RemotePathMappingAPI.GetRemotePathMappingById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingAPI.GetRemotePathMappingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemotePathMappingById`: RemotePathMappingResource
	fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingAPI.GetRemotePathMappingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemotePathMappingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemotePathMapping

> []RemotePathMappingResource ListRemotePathMapping(ctx).Execute()



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
	resp, r, err := apiClient.RemotePathMappingAPI.ListRemotePathMapping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingAPI.ListRemotePathMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRemotePathMapping`: []RemotePathMappingResource
	fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingAPI.ListRemotePathMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRemotePathMappingRequest struct via the builder pattern


### Return type

[**[]RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemotePathMapping

> RemotePathMappingResource UpdateRemotePathMapping(ctx, id).RemotePathMappingResource(remotePathMappingResource).Execute()



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
	remotePathMappingResource := *sonarrClient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemotePathMappingAPI.UpdateRemotePathMapping(context.Background(), id).RemotePathMappingResource(remotePathMappingResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingAPI.UpdateRemotePathMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemotePathMapping`: RemotePathMappingResource
	fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingAPI.UpdateRemotePathMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemotePathMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

