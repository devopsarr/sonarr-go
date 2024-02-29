# \NamingConfigAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNamingConfig**](NamingConfigAPI.md#GetNamingConfig) | **Get** /api/v3/config/naming | 
[**GetNamingConfigById**](NamingConfigAPI.md#GetNamingConfigById) | **Get** /api/v3/config/naming/{id} | 
[**GetNamingConfigExamples**](NamingConfigAPI.md#GetNamingConfigExamples) | **Get** /api/v3/config/naming/examples | 
[**UpdateNamingConfig**](NamingConfigAPI.md#UpdateNamingConfig) | **Put** /api/v3/config/naming/{id} | 



## GetNamingConfig

> NamingConfigResource GetNamingConfig(ctx).Execute()



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
	resp, r, err := apiClient.NamingConfigAPI.GetNamingConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingConfig`: NamingConfigResource
	fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.GetNamingConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigRequest struct via the builder pattern


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingConfigById

> NamingConfigResource GetNamingConfigById(ctx, id).Execute()



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
	resp, r, err := apiClient.NamingConfigAPI.GetNamingConfigById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfigById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingConfigById`: NamingConfigResource
	fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.GetNamingConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingConfigExamples

> GetNamingConfigExamples(ctx).RenameEpisodes(renameEpisodes).ReplaceIllegalCharacters(replaceIllegalCharacters).ColonReplacementFormat(colonReplacementFormat).MultiEpisodeStyle(multiEpisodeStyle).StandardEpisodeFormat(standardEpisodeFormat).DailyEpisodeFormat(dailyEpisodeFormat).AnimeEpisodeFormat(animeEpisodeFormat).SeriesFolderFormat(seriesFolderFormat).SeasonFolderFormat(seasonFolderFormat).SpecialsFolderFormat(specialsFolderFormat).Id(id).ResourceName(resourceName).Execute()



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
	renameEpisodes := true // bool |  (optional)
	replaceIllegalCharacters := true // bool |  (optional)
	colonReplacementFormat := int32(56) // int32 |  (optional)
	multiEpisodeStyle := int32(56) // int32 |  (optional)
	standardEpisodeFormat := "standardEpisodeFormat_example" // string |  (optional)
	dailyEpisodeFormat := "dailyEpisodeFormat_example" // string |  (optional)
	animeEpisodeFormat := "animeEpisodeFormat_example" // string |  (optional)
	seriesFolderFormat := "seriesFolderFormat_example" // string |  (optional)
	seasonFolderFormat := "seasonFolderFormat_example" // string |  (optional)
	specialsFolderFormat := "specialsFolderFormat_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	resourceName := "resourceName_example" // string |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.NamingConfigAPI.GetNamingConfigExamples(context.Background()).RenameEpisodes(renameEpisodes).ReplaceIllegalCharacters(replaceIllegalCharacters).ColonReplacementFormat(colonReplacementFormat).MultiEpisodeStyle(multiEpisodeStyle).StandardEpisodeFormat(standardEpisodeFormat).DailyEpisodeFormat(dailyEpisodeFormat).AnimeEpisodeFormat(animeEpisodeFormat).SeriesFolderFormat(seriesFolderFormat).SeasonFolderFormat(seasonFolderFormat).SpecialsFolderFormat(specialsFolderFormat).Id(id).ResourceName(resourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.GetNamingConfigExamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingConfigExamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameEpisodes** | **bool** |  | 
 **replaceIllegalCharacters** | **bool** |  | 
 **colonReplacementFormat** | **int32** |  | 
 **multiEpisodeStyle** | **int32** |  | 
 **standardEpisodeFormat** | **string** |  | 
 **dailyEpisodeFormat** | **string** |  | 
 **animeEpisodeFormat** | **string** |  | 
 **seriesFolderFormat** | **string** |  | 
 **seasonFolderFormat** | **string** |  | 
 **specialsFolderFormat** | **string** |  | 
 **id** | **int32** |  | 
 **resourceName** | **string** |  | 

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


## UpdateNamingConfig

> NamingConfigResource UpdateNamingConfig(ctx, id).NamingConfigResource(namingConfigResource).Execute()



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
	namingConfigResource := *sonarrClient.NewNamingConfigResource() // NamingConfigResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingConfigAPI.UpdateNamingConfig(context.Background(), id).NamingConfigResource(namingConfigResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigAPI.UpdateNamingConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNamingConfig`: NamingConfigResource
	fmt.Fprintf(os.Stdout, "Response from `NamingConfigAPI.UpdateNamingConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamingConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namingConfigResource** | [**NamingConfigResource**](NamingConfigResource.md) |  | 

### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

