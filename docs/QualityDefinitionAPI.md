# \QualityDefinitionAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQualityDefinitionById**](QualityDefinitionAPI.md#GetQualityDefinitionById) | **Get** /api/v3/qualitydefinition/{id} | 
[**GetQualityDefinitionLimits**](QualityDefinitionAPI.md#GetQualityDefinitionLimits) | **Get** /api/v3/qualitydefinition/limits | 
[**ListQualityDefinition**](QualityDefinitionAPI.md#ListQualityDefinition) | **Get** /api/v3/qualitydefinition | 
[**PutQualityDefinitionUpdate**](QualityDefinitionAPI.md#PutQualityDefinitionUpdate) | **Put** /api/v3/qualitydefinition/update | 
[**UpdateQualityDefinition**](QualityDefinitionAPI.md#UpdateQualityDefinition) | **Put** /api/v3/qualitydefinition/{id} | 



## GetQualityDefinitionById

> QualityDefinitionResource GetQualityDefinitionById(ctx, id).Execute()



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
	resp, r, err := apiClient.QualityDefinitionAPI.GetQualityDefinitionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionAPI.GetQualityDefinitionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQualityDefinitionById`: QualityDefinitionResource
	fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionAPI.GetQualityDefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityDefinitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQualityDefinitionLimits

> QualityDefinitionLimitsResource GetQualityDefinitionLimits(ctx).Execute()



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
	resp, r, err := apiClient.QualityDefinitionAPI.GetQualityDefinitionLimits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionAPI.GetQualityDefinitionLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQualityDefinitionLimits`: QualityDefinitionLimitsResource
	fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionAPI.GetQualityDefinitionLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityDefinitionLimitsRequest struct via the builder pattern


### Return type

[**QualityDefinitionLimitsResource**](QualityDefinitionLimitsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQualityDefinition

> []QualityDefinitionResource ListQualityDefinition(ctx).Execute()



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
	resp, r, err := apiClient.QualityDefinitionAPI.ListQualityDefinition(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionAPI.ListQualityDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQualityDefinition`: []QualityDefinitionResource
	fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionAPI.ListQualityDefinition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListQualityDefinitionRequest struct via the builder pattern


### Return type

[**[]QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutQualityDefinitionUpdate

> PutQualityDefinitionUpdate(ctx).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
	qualityDefinitionResource := []sonarrClient.QualityDefinitionResource{*sonarrClient.NewQualityDefinitionResource()} // []QualityDefinitionResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.QualityDefinitionAPI.PutQualityDefinitionUpdate(context.Background()).QualityDefinitionResource(qualityDefinitionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionAPI.PutQualityDefinitionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutQualityDefinitionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityDefinitionResource** | [**[]QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQualityDefinition

> QualityDefinitionResource UpdateQualityDefinition(ctx, id).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
	qualityDefinitionResource := *sonarrClient.NewQualityDefinitionResource() // QualityDefinitionResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityDefinitionAPI.UpdateQualityDefinition(context.Background(), id).QualityDefinitionResource(qualityDefinitionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionAPI.UpdateQualityDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQualityDefinition`: QualityDefinitionResource
	fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionAPI.UpdateQualityDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQualityDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityDefinitionResource** | [**QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

