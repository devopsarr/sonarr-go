# \QueueDetailsAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQueueDetails**](QueueDetailsAPI.md#ListQueueDetails) | **Get** /api/v3/queue/details | 



## ListQueueDetails

> []QueueResource ListQueueDetails(ctx).SeriesId(seriesId).EpisodeIds(episodeIds).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



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
	episodeIds := []int32{int32(123)} // []int32 |  (optional)
	includeSeries := true // bool |  (optional) (default to false)
	includeEpisode := true // bool |  (optional) (default to false)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueDetailsAPI.ListQueueDetails(context.Background()).SeriesId(seriesId).EpisodeIds(episodeIds).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsAPI.ListQueueDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueueDetails`: []QueueResource
	fmt.Fprintf(os.Stdout, "Response from `QueueDetailsAPI.ListQueueDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **episodeIds** | **[]int32** |  | 
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

