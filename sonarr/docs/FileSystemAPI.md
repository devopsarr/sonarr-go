# \FileSystemAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileSystem**](FileSystemAPI.md#GetFileSystem) | **Get** /api/v3/filesystem | 
[**GetFileSystemMediafiles**](FileSystemAPI.md#GetFileSystemMediafiles) | **Get** /api/v3/filesystem/mediafiles | 
[**GetFileSystemType**](FileSystemAPI.md#GetFileSystemType) | **Get** /api/v3/filesystem/type | 



## GetFileSystem

> GetFileSystem(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sonarrClient "./openapi"
)

func main() {
    path := "path_example" // string |  (optional)
    includeFiles := true // bool |  (optional) (default to false)
    allowFoldersWithoutTrailingSlashes := true // bool |  (optional) (default to false)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemAPI.GetFileSystem(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemAPI.GetFileSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** |  | [default to false]
 **allowFoldersWithoutTrailingSlashes** | **bool** |  | [default to false]

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


## GetFileSystemMediafiles

> GetFileSystemMediafiles(ctx).Path(path).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sonarrClient "./openapi"
)

func main() {
    path := "path_example" // string |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemAPI.GetFileSystemMediafiles(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemAPI.GetFileSystemMediafiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSystemMediafilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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


## GetFileSystemType

> GetFileSystemType(ctx).Path(path).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sonarrClient "./openapi"
)

func main() {
    path := "path_example" // string |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemAPI.GetFileSystemType(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemAPI.GetFileSystemType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSystemTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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

