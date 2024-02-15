# \LocalizationAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocalization**](LocalizationAPI.md#GetLocalization) | **Get** /api/v3/localization | 
[**GetLocalizationById**](LocalizationAPI.md#GetLocalizationById) | **Get** /api/v3/localization/{id} | 
[**GetLocalizationLanguage**](LocalizationAPI.md#GetLocalizationLanguage) | **Get** /api/v3/localization/language | 



## GetLocalization

> LocalizationResource GetLocalization(ctx).Execute()



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
	resp, r, err := apiClient.LocalizationAPI.GetLocalization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetLocalization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalization`: LocalizationResource
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetLocalization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationRequest struct via the builder pattern


### Return type

[**LocalizationResource**](LocalizationResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationById

> LocalizationResource GetLocalizationById(ctx, id).Execute()



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
	resp, r, err := apiClient.LocalizationAPI.GetLocalizationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetLocalizationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationById`: LocalizationResource
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetLocalizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocalizationResource**](LocalizationResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationLanguage

> LocalizationLanguageResource GetLocalizationLanguage(ctx).Execute()



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
	resp, r, err := apiClient.LocalizationAPI.GetLocalizationLanguage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetLocalizationLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationLanguage`: LocalizationLanguageResource
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetLocalizationLanguage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationLanguageRequest struct via the builder pattern


### Return type

[**LocalizationLanguageResource**](LocalizationLanguageResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

