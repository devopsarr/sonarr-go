# \EpisodeAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEpisodeById**](EpisodeAPI.md#GetEpisodeById) | **Get** /api/v3/episode/{id} | 
[**ListEpisode**](EpisodeAPI.md#ListEpisode) | **Get** /api/v3/episode | 
[**PutEpisodeMonitor**](EpisodeAPI.md#PutEpisodeMonitor) | **Put** /api/v3/episode/monitor | 
[**UpdateEpisode**](EpisodeAPI.md#UpdateEpisode) | **Put** /api/v3/episode/{id} | 



## GetEpisodeById

> EpisodeResource GetEpisodeById(ctx, id).Execute()



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
	resp, r, err := apiClient.EpisodeAPI.GetEpisodeById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpisodeAPI.GetEpisodeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEpisodeById`: EpisodeResource
	fmt.Fprintf(os.Stdout, "Response from `EpisodeAPI.GetEpisodeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpisodeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeResource**](EpisodeResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpisode

> []EpisodeResource ListEpisode(ctx).SeriesId(seriesId).SeasonNumber(seasonNumber).EpisodeIds(episodeIds).EpisodeFileId(episodeFileId).IncludeImages(includeImages).Execute()



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
	seriesId := int32(56) // int32 |  (optional)
	seasonNumber := int32(56) // int32 |  (optional)
	episodeIds := []int32{int32(123)} // []int32 |  (optional)
	episodeFileId := int32(56) // int32 |  (optional)
	includeImages := true // bool |  (optional) (default to false)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpisodeAPI.ListEpisode(context.Background()).SeriesId(seriesId).SeasonNumber(seasonNumber).EpisodeIds(episodeIds).EpisodeFileId(episodeFileId).IncludeImages(includeImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpisodeAPI.ListEpisode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEpisode`: []EpisodeResource
	fmt.Fprintf(os.Stdout, "Response from `EpisodeAPI.ListEpisode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEpisodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 
 **episodeIds** | **[]int32** |  | 
 **episodeFileId** | **int32** |  | 
 **includeImages** | **bool** |  | [default to false]

### Return type

[**[]EpisodeResource**](EpisodeResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEpisodeMonitor

> PutEpisodeMonitor(ctx).IncludeImages(includeImages).EpisodesMonitoredResource(episodesMonitoredResource).Execute()



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
	includeImages := true // bool |  (optional) (default to false)
	episodesMonitoredResource := *sonarrClient.NewEpisodesMonitoredResource() // EpisodesMonitoredResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.EpisodeAPI.PutEpisodeMonitor(context.Background()).IncludeImages(includeImages).EpisodesMonitoredResource(episodesMonitoredResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpisodeAPI.PutEpisodeMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEpisodeMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeImages** | **bool** |  | [default to false]
 **episodesMonitoredResource** | [**EpisodesMonitoredResource**](EpisodesMonitoredResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEpisode

> EpisodeResource UpdateEpisode(ctx, id).EpisodeResource(episodeResource).Execute()



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
	episodeResource := *sonarrClient.NewEpisodeResource() // EpisodeResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpisodeAPI.UpdateEpisode(context.Background(), id).EpisodeResource(episodeResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpisodeAPI.UpdateEpisode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEpisode`: EpisodeResource
	fmt.Fprintf(os.Stdout, "Response from `EpisodeAPI.UpdateEpisode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpisodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **episodeResource** | [**EpisodeResource**](EpisodeResource.md) |  | 

### Return type

[**EpisodeResource**](EpisodeResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

