# \AutoTaggingAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoTagging**](AutoTaggingAPI.md#CreateAutoTagging) | **Post** /api/v3/autotagging | 
[**DeleteAutoTagging**](AutoTaggingAPI.md#DeleteAutoTagging) | **Delete** /api/v3/autotagging/{id} | 
[**GetAutoTaggingById**](AutoTaggingAPI.md#GetAutoTaggingById) | **Get** /api/v3/autotagging/{id} | 
[**ListAutoTagging**](AutoTaggingAPI.md#ListAutoTagging) | **Get** /api/v3/autotagging | 
[**ListAutoTaggingSchema**](AutoTaggingAPI.md#ListAutoTaggingSchema) | **Get** /api/v3/autotagging/schema | 
[**UpdateAutoTagging**](AutoTaggingAPI.md#UpdateAutoTagging) | **Put** /api/v3/autotagging/{id} | 



## CreateAutoTagging

> AutoTaggingResource CreateAutoTagging(ctx).AutoTaggingResource(autoTaggingResource).Execute()



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
	autoTaggingResource := *sonarrClient.NewAutoTaggingResource() // AutoTaggingResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTaggingAPI.CreateAutoTagging(context.Background()).AutoTaggingResource(autoTaggingResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.CreateAutoTagging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoTagging`: AutoTaggingResource
	fmt.Fprintf(os.Stdout, "Response from `AutoTaggingAPI.CreateAutoTagging`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTaggingResource** | [**AutoTaggingResource**](AutoTaggingResource.md) |  | 

### Return type

[**AutoTaggingResource**](AutoTaggingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoTagging

> DeleteAutoTagging(ctx, id).Execute()



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
	r, err := apiClient.AutoTaggingAPI.DeleteAutoTagging(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.DeleteAutoTagging``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAutoTaggingRequest struct via the builder pattern


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


## GetAutoTaggingById

> AutoTaggingResource GetAutoTaggingById(ctx, id).Execute()



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
	resp, r, err := apiClient.AutoTaggingAPI.GetAutoTaggingById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.GetAutoTaggingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoTaggingById`: AutoTaggingResource
	fmt.Fprintf(os.Stdout, "Response from `AutoTaggingAPI.GetAutoTaggingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTaggingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTaggingResource**](AutoTaggingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoTagging

> []AutoTaggingResource ListAutoTagging(ctx).Execute()



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
	resp, r, err := apiClient.AutoTaggingAPI.ListAutoTagging(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.ListAutoTagging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoTagging`: []AutoTaggingResource
	fmt.Fprintf(os.Stdout, "Response from `AutoTaggingAPI.ListAutoTagging`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoTaggingRequest struct via the builder pattern


### Return type

[**[]AutoTaggingResource**](AutoTaggingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoTaggingSchema

> []AutoTaggingSpecificationSchema ListAutoTaggingSchema(ctx).Execute()



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
	resp, r, err := apiClient.AutoTaggingAPI.ListAutoTaggingSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.ListAutoTaggingSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoTaggingSchema`: []AutoTaggingSpecificationSchema
	fmt.Fprintf(os.Stdout, "Response from `AutoTaggingAPI.ListAutoTaggingSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoTaggingSchemaRequest struct via the builder pattern


### Return type

[**[]AutoTaggingSpecificationSchema**](AutoTaggingSpecificationSchema.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoTagging

> AutoTaggingResource UpdateAutoTagging(ctx, id).AutoTaggingResource(autoTaggingResource).Execute()



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
	autoTaggingResource := *sonarrClient.NewAutoTaggingResource() // AutoTaggingResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTaggingAPI.UpdateAutoTagging(context.Background(), id).AutoTaggingResource(autoTaggingResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTaggingAPI.UpdateAutoTagging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoTagging`: AutoTaggingResource
	fmt.Fprintf(os.Stdout, "Response from `AutoTaggingAPI.UpdateAutoTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTaggingResource** | [**AutoTaggingResource**](AutoTaggingResource.md) |  | 

### Return type

[**AutoTaggingResource**](AutoTaggingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

