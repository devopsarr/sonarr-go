/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// CalendarFeedApiService CalendarFeedApi service
type CalendarFeedApiService service
type ApiGetFeedV3CalendarSonarrIcsRequest struct {
	ctx context.Context
	ApiService *CalendarFeedApiService
	pastDays *int32
	futureDays *int32
	tagList *string
	unmonitored *bool
	premieresOnly *bool
	asAllDay *bool
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) PastDays(pastDays int32) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.pastDays = &pastDays
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) FutureDays(futureDays int32) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.futureDays = &futureDays
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) TagList(tagList string) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.tagList = &tagList
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) Unmonitored(unmonitored bool) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.unmonitored = &unmonitored
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) PremieresOnly(premieresOnly bool) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.premieresOnly = &premieresOnly
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) AsAllDay(asAllDay bool) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.asAllDay = &asAllDay
	return r
}

func (r ApiGetFeedV3CalendarSonarrIcsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetFeedV3CalendarSonarrIcsExecute(r)
}

/*
GetFeedV3CalendarSonarrIcs Method for GetFeedV3CalendarSonarrIcs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFeedV3CalendarSonarrIcsRequest
*/
func (a *CalendarFeedApiService) GetFeedV3CalendarSonarrIcs(ctx context.Context) ApiGetFeedV3CalendarSonarrIcsRequest {
	return ApiGetFeedV3CalendarSonarrIcsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CalendarFeedApiService) GetFeedV3CalendarSonarrIcsExecute(r ApiGetFeedV3CalendarSonarrIcsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarFeedApiService.GetFeedV3CalendarSonarrIcs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feed/v3/calendar/sonarr.ics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pastDays != nil {
		localVarQueryParams.Add("pastDays", parameterToString(*r.pastDays, ""))
	}
	if r.futureDays != nil {
		localVarQueryParams.Add("futureDays", parameterToString(*r.futureDays, ""))
	}
	if r.tagList != nil {
		localVarQueryParams.Add("tagList", parameterToString(*r.tagList, ""))
	}
	if r.unmonitored != nil {
		localVarQueryParams.Add("unmonitored", parameterToString(*r.unmonitored, ""))
	}
	if r.premieresOnly != nil {
		localVarQueryParams.Add("premieresOnly", parameterToString(*r.premieresOnly, ""))
	}
	if r.asAllDay != nil {
		localVarQueryParams.Add("asAllDay", parameterToString(*r.asAllDay, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
