# \SeriesAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSeries**](SeriesAPI.md#CreateSeries) | **Post** /api/v3/series | 
[**DeleteSeries**](SeriesAPI.md#DeleteSeries) | **Delete** /api/v3/series/{id} | 
[**GetSeriesById**](SeriesAPI.md#GetSeriesById) | **Get** /api/v3/series/{id} | 
[**ListSeries**](SeriesAPI.md#ListSeries) | **Get** /api/v3/series | 
[**UpdateSeries**](SeriesAPI.md#UpdateSeries) | **Put** /api/v3/series/{id} | 



## CreateSeries

> SeriesResource CreateSeries(ctx).SeriesResource(seriesResource).Execute()



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
	seriesResource := *sonarrClient.NewSeriesResource() // SeriesResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeriesAPI.CreateSeries(context.Background()).SeriesResource(seriesResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.CreateSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSeries`: SeriesResource
	fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.CreateSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesResource** | [**SeriesResource**](SeriesResource.md) |  | 

### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSeries

> DeleteSeries(ctx, id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()



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
	deleteFiles := true // bool |  (optional) (default to false)
	addImportListExclusion := true // bool |  (optional) (default to false)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.SeriesAPI.DeleteSeries(context.Background(), id).DeleteFiles(deleteFiles).AddImportListExclusion(addImportListExclusion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.DeleteSeries``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFiles** | **bool** |  | [default to false]
 **addImportListExclusion** | **bool** |  | [default to false]

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


## GetSeriesById

> SeriesResource GetSeriesById(ctx, id).IncludeSeasonImages(includeSeasonImages).Execute()



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
	includeSeasonImages := true // bool |  (optional) (default to false)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeriesAPI.GetSeriesById(context.Background(), id).IncludeSeasonImages(includeSeasonImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.GetSeriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeriesById`: SeriesResource
	fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.GetSeriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSeasonImages** | **bool** |  | [default to false]

### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSeries

> []SeriesResource ListSeries(ctx).TvdbId(tvdbId).IncludeSeasonImages(includeSeasonImages).Execute()



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
	tvdbId := int32(56) // int32 |  (optional)
	includeSeasonImages := true // bool |  (optional) (default to false)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeriesAPI.ListSeries(context.Background()).TvdbId(tvdbId).IncludeSeasonImages(includeSeasonImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.ListSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSeries`: []SeriesResource
	fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.ListSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **int32** |  | 
 **includeSeasonImages** | **bool** |  | [default to false]

### Return type

[**[]SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSeries

> SeriesResource UpdateSeries(ctx, id).MoveFiles(moveFiles).SeriesResource(seriesResource).Execute()



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
	moveFiles := true // bool |  (optional) (default to false)
	seriesResource := *sonarrClient.NewSeriesResource() // SeriesResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeriesAPI.UpdateSeries(context.Background(), id).MoveFiles(moveFiles).SeriesResource(seriesResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeriesAPI.UpdateSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSeries`: SeriesResource
	fmt.Fprintf(os.Stdout, "Response from `SeriesAPI.UpdateSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveFiles** | **bool** |  | [default to false]
 **seriesResource** | [**SeriesResource**](SeriesResource.md) |  | 

### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

