# \CalendarFeedApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedV3CalendarSonarrIcs**](CalendarFeedApi.md#GetFeedV3CalendarSonarrIcs) | **Get** /feed/v3/calendar/sonarr.ics | 



## GetFeedV3CalendarSonarrIcs

> GetFeedV3CalendarSonarrIcs(ctx).PastDays(pastDays).FutureDays(futureDays).Tags(tags).Unmonitored(unmonitored).PremieresOnly(premieresOnly).AsAllDay(asAllDay).Execute()



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
    pastDays := int32(56) // int32 |  (optional) (default to 7)
    futureDays := int32(56) // int32 |  (optional) (default to 28)
    tags := "tags_example" // string |  (optional) (default to "")
    unmonitored := true // bool |  (optional) (default to false)
    premieresOnly := true // bool |  (optional) (default to false)
    asAllDay := true // bool |  (optional) (default to false)

    configuration := sonarrClient.NewConfiguration()
    apiClient := sonarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarFeedApi.GetFeedV3CalendarSonarrIcs(context.Background()).PastDays(pastDays).FutureDays(futureDays).Tags(tags).Unmonitored(unmonitored).PremieresOnly(premieresOnly).AsAllDay(asAllDay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarFeedApi.GetFeedV3CalendarSonarrIcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedV3CalendarSonarrIcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pastDays** | **int32** |  | [default to 7]
 **futureDays** | **int32** |  | [default to 28]
 **tags** | **string** |  | [default to &quot;&quot;]
 **unmonitored** | **bool** |  | [default to false]
 **premieresOnly** | **bool** |  | [default to false]
 **asAllDay** | **bool** |  | [default to false]

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

