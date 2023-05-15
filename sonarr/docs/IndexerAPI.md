# \IndexerAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndexer**](IndexerAPI.md#CreateIndexer) | **Post** /api/v3/indexer | 
[**CreateIndexerActionByName**](IndexerAPI.md#CreateIndexerActionByName) | **Post** /api/v3/indexer/action/{name} | 
[**DeleteIndexer**](IndexerAPI.md#DeleteIndexer) | **Delete** /api/v3/indexer/{id} | 
[**GetIndexerById**](IndexerAPI.md#GetIndexerById) | **Get** /api/v3/indexer/{id} | 
[**ListIndexer**](IndexerAPI.md#ListIndexer) | **Get** /api/v3/indexer | 
[**ListIndexerSchema**](IndexerAPI.md#ListIndexerSchema) | **Get** /api/v3/indexer/schema | 
[**TestIndexer**](IndexerAPI.md#TestIndexer) | **Post** /api/v3/indexer/test | 
[**TestallIndexer**](IndexerAPI.md#TestallIndexer) | **Post** /api/v3/indexer/testall | 
[**UpdateIndexer**](IndexerAPI.md#UpdateIndexer) | **Put** /api/v3/indexer/{id} | 



## CreateIndexer

> IndexerResource CreateIndexer(ctx).IndexerResource(indexerResource).Execute()



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
    indexerResource := *sonarrClient.NewIndexerResource() // IndexerResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.CreateIndexer(context.Background()).IndexerResource(indexerResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.CreateIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndexer`: IndexerResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerAPI.CreateIndexer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerResource** | [**IndexerResource**](IndexerResource.md) |  | 

### Return type

[**IndexerResource**](IndexerResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndexerActionByName

> CreateIndexerActionByName(ctx, name).IndexerResource(indexerResource).Execute()



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
    name := "name_example" // string | 
    indexerResource := *sonarrClient.NewIndexerResource() // IndexerResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.CreateIndexerActionByName(context.Background(), name).IndexerResource(indexerResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.CreateIndexerActionByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexerActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerResource** | [**IndexerResource**](IndexerResource.md) |  | 

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


## DeleteIndexer

> DeleteIndexer(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.DeleteIndexer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.DeleteIndexer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndexerRequest struct via the builder pattern


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


## GetIndexerById

> IndexerResource GetIndexerById(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.GetIndexerById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.GetIndexerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexerById`: IndexerResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerAPI.GetIndexerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerResource**](IndexerResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexer

> []IndexerResource ListIndexer(ctx).Execute()



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

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.ListIndexer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.ListIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexer`: []IndexerResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerAPI.ListIndexer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerRequest struct via the builder pattern


### Return type

[**[]IndexerResource**](IndexerResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIndexerSchema

> []IndexerResource ListIndexerSchema(ctx).Execute()



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

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.ListIndexerSchema(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.ListIndexerSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIndexerSchema`: []IndexerResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerAPI.ListIndexerSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIndexerSchemaRequest struct via the builder pattern


### Return type

[**[]IndexerResource**](IndexerResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIndexer

> TestIndexer(ctx).IndexerResource(indexerResource).Execute()



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
    indexerResource := *sonarrClient.NewIndexerResource() // IndexerResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.TestIndexer(context.Background()).IndexerResource(indexerResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.TestIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestIndexerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexerResource** | [**IndexerResource**](IndexerResource.md) |  | 

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


## TestallIndexer

> TestallIndexer(ctx).Execute()



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

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.TestallIndexer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.TestallIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallIndexerRequest struct via the builder pattern


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


## UpdateIndexer

> IndexerResource UpdateIndexer(ctx, id).IndexerResource(indexerResource).Execute()



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
    id := "id_example" // string | 
    indexerResource := *sonarrClient.NewIndexerResource() // IndexerResource |  (optional)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerAPI.UpdateIndexer(context.Background(), id).IndexerResource(indexerResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerAPI.UpdateIndexer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndexer`: IndexerResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerAPI.UpdateIndexer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndexerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerResource** | [**IndexerResource**](IndexerResource.md) |  | 

### Return type

[**IndexerResource**](IndexerResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

