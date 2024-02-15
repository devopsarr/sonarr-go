# \ManualImportAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualImport**](ManualImportAPI.md#CreateManualImport) | **Post** /api/v3/manualimport | 
[**ListManualImport**](ManualImportAPI.md#ListManualImport) | **Get** /api/v3/manualimport | 



## CreateManualImport

> CreateManualImport(ctx).ManualImportReprocessResource(manualImportReprocessResource).Execute()



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
	manualImportReprocessResource := []sonarrClient.ManualImportReprocessResource{*sonarrClient.NewManualImportReprocessResource()} // []ManualImportReprocessResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.ManualImportAPI.CreateManualImport(context.Background()).ManualImportReprocessResource(manualImportReprocessResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualImportAPI.CreateManualImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualImportReprocessResource** | [**[]ManualImportReprocessResource**](ManualImportReprocessResource.md) |  | 

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


## ListManualImport

> []ManualImportResource ListManualImport(ctx).Folder(folder).DownloadId(downloadId).SeriesId(seriesId).SeasonNumber(seasonNumber).FilterExistingFiles(filterExistingFiles).Execute()



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
	folder := "folder_example" // string |  (optional)
	downloadId := "downloadId_example" // string |  (optional)
	seriesId := int32(56) // int32 |  (optional)
	seasonNumber := int32(56) // int32 |  (optional)
	filterExistingFiles := true // bool |  (optional) (default to true)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManualImportAPI.ListManualImport(context.Background()).Folder(folder).DownloadId(downloadId).SeriesId(seriesId).SeasonNumber(seasonNumber).FilterExistingFiles(filterExistingFiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManualImportAPI.ListManualImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListManualImport`: []ManualImportResource
	fmt.Fprintf(os.Stdout, "Response from `ManualImportAPI.ListManualImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folder** | **string** |  | 
 **downloadId** | **string** |  | 
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 
 **filterExistingFiles** | **bool** |  | [default to true]

### Return type

[**[]ManualImportResource**](ManualImportResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

