# \DownloadClientAPI

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDownloadClient**](DownloadClientAPI.md#CreateDownloadClient) | **Post** /api/v3/downloadclient | 
[**CreateDownloadClientActionByName**](DownloadClientAPI.md#CreateDownloadClientActionByName) | **Post** /api/v3/downloadclient/action/{name} | 
[**DeleteDownloadClient**](DownloadClientAPI.md#DeleteDownloadClient) | **Delete** /api/v3/downloadclient/{id} | 
[**DeleteDownloadClientBulk**](DownloadClientAPI.md#DeleteDownloadClientBulk) | **Delete** /api/v3/downloadclient/bulk | 
[**GetDownloadClientById**](DownloadClientAPI.md#GetDownloadClientById) | **Get** /api/v3/downloadclient/{id} | 
[**ListDownloadClient**](DownloadClientAPI.md#ListDownloadClient) | **Get** /api/v3/downloadclient | 
[**ListDownloadClientSchema**](DownloadClientAPI.md#ListDownloadClientSchema) | **Get** /api/v3/downloadclient/schema | 
[**PutDownloadClientBulk**](DownloadClientAPI.md#PutDownloadClientBulk) | **Put** /api/v3/downloadclient/bulk | 
[**TestDownloadClient**](DownloadClientAPI.md#TestDownloadClient) | **Post** /api/v3/downloadclient/test | 
[**TestallDownloadClient**](DownloadClientAPI.md#TestallDownloadClient) | **Post** /api/v3/downloadclient/testall | 
[**UpdateDownloadClient**](DownloadClientAPI.md#UpdateDownloadClient) | **Put** /api/v3/downloadclient/{id} | 



## CreateDownloadClient

> DownloadClientResource CreateDownloadClient(ctx).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()



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
	downloadClientResource := *sonarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadClientAPI.CreateDownloadClient(context.Background()).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.CreateDownloadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDownloadClient`: DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.CreateDownloadClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceSave** | **bool** |  | [default to false]
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDownloadClientActionByName

> CreateDownloadClientActionByName(ctx, name).DownloadClientResource(downloadClientResource).Execute()



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
	downloadClientResource := *sonarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.DownloadClientAPI.CreateDownloadClientActionByName(context.Background(), name).DownloadClientResource(downloadClientResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.CreateDownloadClientActionByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateDownloadClientActionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

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


## DeleteDownloadClient

> DeleteDownloadClient(ctx, id).Execute()



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
	r, err := apiClient.DownloadClientAPI.DeleteDownloadClient(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.DeleteDownloadClient``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDownloadClientRequest struct via the builder pattern


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


## DeleteDownloadClientBulk

> DeleteDownloadClientBulk(ctx).DownloadClientBulkResource(downloadClientBulkResource).Execute()



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
	downloadClientBulkResource := *sonarrClient.NewDownloadClientBulkResource() // DownloadClientBulkResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.DownloadClientAPI.DeleteDownloadClientBulk(context.Background()).DownloadClientBulkResource(downloadClientBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.DeleteDownloadClientBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDownloadClientBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientBulkResource** | [**DownloadClientBulkResource**](DownloadClientBulkResource.md) |  | 

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


## GetDownloadClientById

> DownloadClientResource GetDownloadClientById(ctx, id).Execute()



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
	resp, r, err := apiClient.DownloadClientAPI.GetDownloadClientById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.GetDownloadClientById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownloadClientById`: DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.GetDownloadClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClient

> []DownloadClientResource ListDownloadClient(ctx).Execute()



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
	resp, r, err := apiClient.DownloadClientAPI.ListDownloadClient(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.ListDownloadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDownloadClient`: []DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.ListDownloadClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDownloadClientSchema

> []DownloadClientResource ListDownloadClientSchema(ctx).Execute()



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
	resp, r, err := apiClient.DownloadClientAPI.ListDownloadClientSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.ListDownloadClientSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDownloadClientSchema`: []DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.ListDownloadClientSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDownloadClientSchemaRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDownloadClientBulk

> DownloadClientResource PutDownloadClientBulk(ctx).DownloadClientBulkResource(downloadClientBulkResource).Execute()



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
	downloadClientBulkResource := *sonarrClient.NewDownloadClientBulkResource() // DownloadClientBulkResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadClientAPI.PutDownloadClientBulk(context.Background()).DownloadClientBulkResource(downloadClientBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.PutDownloadClientBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDownloadClientBulk`: DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.PutDownloadClientBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDownloadClientBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientBulkResource** | [**DownloadClientBulkResource**](DownloadClientBulkResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDownloadClient

> TestDownloadClient(ctx).ForceTest(forceTest).DownloadClientResource(downloadClientResource).Execute()



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
	downloadClientResource := *sonarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	r, err := apiClient.DownloadClientAPI.TestDownloadClient(context.Background()).ForceTest(forceTest).DownloadClientResource(downloadClientResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.TestDownloadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceTest** | **bool** |  | [default to false]
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

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


## TestallDownloadClient

> TestallDownloadClient(ctx).Execute()



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
	r, err := apiClient.DownloadClientAPI.TestallDownloadClient(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.TestallDownloadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestallDownloadClientRequest struct via the builder pattern


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


## UpdateDownloadClient

> DownloadClientResource UpdateDownloadClient(ctx, id).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()



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
	forceSave := true // bool |  (optional) (default to false)
	downloadClientResource := *sonarrClient.NewDownloadClientResource() // DownloadClientResource |  (optional)

	configuration := sonarrClient.NewConfiguration()
	apiClient := sonarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadClientAPI.UpdateDownloadClient(context.Background(), id).ForceSave(forceSave).DownloadClientResource(downloadClientResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientAPI.UpdateDownloadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDownloadClient`: DownloadClientResource
	fmt.Fprintf(os.Stdout, "Response from `DownloadClientAPI.UpdateDownloadClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDownloadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceSave** | **bool** |  | [default to false]
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

