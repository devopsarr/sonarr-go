# \ImportListAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImportList**](ImportListAPI.md#CreateImportList) | **Post** /api/v3/importlist | 
[**CreateImportListActionByName**](ImportListAPI.md#CreateImportListActionByName) | **Post** /api/v3/importlist/action/{name} | 
[**DeleteImportList**](ImportListAPI.md#DeleteImportList) | **Delete** /api/v3/importlist/{id} | 
[**DeleteImportListBulk**](ImportListAPI.md#DeleteImportListBulk) | **Delete** /api/v3/importlist/bulk | 
[**GetImportListById**](ImportListAPI.md#GetImportListById) | **Get** /api/v3/importlist/{id} | 
[**ListImportList**](ImportListAPI.md#ListImportList) | **Get** /api/v3/importlist | 
[**ListImportListSchema**](ImportListAPI.md#ListImportListSchema) | **Get** /api/v3/importlist/schema | 
[**PutImportListBulk**](ImportListAPI.md#PutImportListBulk) | **Put** /api/v3/importlist/bulk | 
[**TestImportList**](ImportListAPI.md#TestImportList) | **Post** /api/v3/importlist/test | 
[**TestallImportList**](ImportListAPI.md#TestallImportList) | **Post** /api/v3/importlist/testall | 
[**UpdateImportList**](ImportListAPI.md#UpdateImportList) | **Put** /api/v3/importlist/{id} | 



## CreateImportList

> ImportListResource CreateImportList(ctx).ForceSave(forceSave).ImportListResource(importListResource).Execute()



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
	forceSave := true // bool |  (optional) (default to false)
	importListResource := *sonarrClient.NewImportListResource() // ImportListResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListAPI.CreateImportList(context.Background()).ForceSave(forceSave).ImportListResource(importListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.CreateImportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImportList`: ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.CreateImportList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSave** | **bool** |  | [default to false]
 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateImportListActionByName

> CreateImportListActionByName(ctx, name).ImportListResource(importListResource).Execute()



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
	name := "name_example" // string | 
	importListResource := *sonarrClient.NewImportListResource() // ImportListResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportListAPI.CreateImportListActionByName(context.Background(), name).ImportListResource(importListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.CreateImportListActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateImportListActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

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


## DeleteImportList

> DeleteImportList(ctx, id).Execute()



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
	r, err := apiClient.ImportListAPI.DeleteImportList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.DeleteImportList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteImportListRequest struct via the builder pattern


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


## DeleteImportListBulk

> DeleteImportListBulk(ctx).ImportListBulkResource(importListBulkResource).Execute()



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
	importListBulkResource := *sonarrClient.NewImportListBulkResource() // ImportListBulkResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportListAPI.DeleteImportListBulk(context.Background()).ImportListBulkResource(importListBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.DeleteImportListBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImportListBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListBulkResource** | [**ImportListBulkResource**](ImportListBulkResource.md) |  | 

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


## GetImportListById

> ImportListResource GetImportListById(ctx, id).Execute()



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
	resp, r, err := apiClient.ImportListAPI.GetImportListById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.GetImportListById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportListById`: ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.GetImportListById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportListByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportList

> []ImportListResource ListImportList(ctx).Execute()



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
	resp, r, err := apiClient.ImportListAPI.ListImportList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.ListImportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImportList`: []ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.ListImportList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListImportListRequest struct via the builder pattern


### Return type

[**[]ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImportListSchema

> []ImportListResource ListImportListSchema(ctx).Execute()



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
	resp, r, err := apiClient.ImportListAPI.ListImportListSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.ListImportListSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImportListSchema`: []ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.ListImportListSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListImportListSchemaRequest struct via the builder pattern


### Return type

[**[]ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutImportListBulk

> ImportListResource PutImportListBulk(ctx).ImportListBulkResource(importListBulkResource).Execute()



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
	importListBulkResource := *sonarrClient.NewImportListBulkResource() // ImportListBulkResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListAPI.PutImportListBulk(context.Background()).ImportListBulkResource(importListBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.PutImportListBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutImportListBulk`: ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.PutImportListBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutImportListBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListBulkResource** | [**ImportListBulkResource**](ImportListBulkResource.md) |  | 

### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestImportList

> TestImportList(ctx).ForceTest(forceTest).ImportListResource(importListResource).Execute()



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
	forceTest := true // bool |  (optional) (default to false)
	importListResource := *sonarrClient.NewImportListResource() // ImportListResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportListAPI.TestImportList(context.Background()).ForceTest(forceTest).ImportListResource(importListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.TestImportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceTest** | **bool** |  | [default to false]
 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

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


## TestallImportList

> TestallImportList(ctx).Execute()



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
	r, err := apiClient.ImportListAPI.TestallImportList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.TestallImportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallImportListRequest struct via the builder pattern


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


## UpdateImportList

> ImportListResource UpdateImportList(ctx, id).ForceSave(forceSave).ImportListResource(importListResource).Execute()



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
	forceSave := true // bool |  (optional) (default to false)
	importListResource := *sonarrClient.NewImportListResource() // ImportListResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListAPI.UpdateImportList(context.Background(), id).ForceSave(forceSave).ImportListResource(importListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListAPI.UpdateImportList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImportList`: ImportListResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListAPI.UpdateImportList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImportListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceSave** | **bool** |  | [default to false]
 **importListResource** | [**ImportListResource**](ImportListResource.md) |  | 

### Return type

[**ImportListResource**](ImportListResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

