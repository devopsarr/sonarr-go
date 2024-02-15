# \LanguageProfileAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanguageProfile**](LanguageProfileAPI.md#CreateLanguageProfile) | **Post** /api/v3/languageprofile | 
[**DeleteLanguageProfile**](LanguageProfileAPI.md#DeleteLanguageProfile) | **Delete** /api/v3/languageprofile/{id} | 
[**GetLanguageProfileById**](LanguageProfileAPI.md#GetLanguageProfileById) | **Get** /api/v3/languageprofile/{id} | 
[**ListLanguageProfile**](LanguageProfileAPI.md#ListLanguageProfile) | **Get** /api/v3/languageprofile | 
[**UpdateLanguageProfile**](LanguageProfileAPI.md#UpdateLanguageProfile) | **Put** /api/v3/languageprofile/{id} | 



## CreateLanguageProfile

> LanguageProfileResource CreateLanguageProfile(ctx).LanguageProfileResource(languageProfileResource).Execute()



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
	languageProfileResource := *sonarrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageProfileAPI.CreateLanguageProfile(context.Background()).LanguageProfileResource(languageProfileResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileAPI.CreateLanguageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanguageProfile`: LanguageProfileResource
	fmt.Fprintf(os.Stdout, "Response from `LanguageProfileAPI.CreateLanguageProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanguageProfile

> DeleteLanguageProfile(ctx, id).Execute()



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
	r, err := apiClient.LanguageProfileAPI.DeleteLanguageProfile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileAPI.DeleteLanguageProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLanguageProfileRequest struct via the builder pattern


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


## GetLanguageProfileById

> LanguageProfileResource GetLanguageProfileById(ctx, id).Execute()



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
	resp, r, err := apiClient.LanguageProfileAPI.GetLanguageProfileById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileAPI.GetLanguageProfileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanguageProfileById`: LanguageProfileResource
	fmt.Fprintf(os.Stdout, "Response from `LanguageProfileAPI.GetLanguageProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguageProfile

> []LanguageProfileResource ListLanguageProfile(ctx).Execute()



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
	resp, r, err := apiClient.LanguageProfileAPI.ListLanguageProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileAPI.ListLanguageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLanguageProfile`: []LanguageProfileResource
	fmt.Fprintf(os.Stdout, "Response from `LanguageProfileAPI.ListLanguageProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLanguageProfileRequest struct via the builder pattern


### Return type

[**[]LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguageProfile

> LanguageProfileResource UpdateLanguageProfile(ctx, id).LanguageProfileResource(languageProfileResource).Execute()



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
	languageProfileResource := *sonarrClient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageProfileAPI.UpdateLanguageProfile(context.Background(), id).LanguageProfileResource(languageProfileResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileAPI.UpdateLanguageProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLanguageProfile`: LanguageProfileResource
	fmt.Fprintf(os.Stdout, "Response from `LanguageProfileAPI.UpdateLanguageProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguageProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

