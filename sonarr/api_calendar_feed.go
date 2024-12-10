/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.11.2680
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


// CalendarFeedAPIService CalendarFeedAPI service
type CalendarFeedAPIService service

type ApiGetFeedV3CalendarSonarrIcsRequest struct {
	ctx context.Context
	ApiService *CalendarFeedAPIService
	pastDays *int32
	futureDays *int32
	tags *string
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

func (r ApiGetFeedV3CalendarSonarrIcsRequest) Tags(tags string) ApiGetFeedV3CalendarSonarrIcsRequest {
	r.tags = &tags
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
func (a *CalendarFeedAPIService) GetFeedV3CalendarSonarrIcs(ctx context.Context) ApiGetFeedV3CalendarSonarrIcsRequest {
	return ApiGetFeedV3CalendarSonarrIcsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CalendarFeedAPIService) GetFeedV3CalendarSonarrIcsExecute(r ApiGetFeedV3CalendarSonarrIcsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarFeedAPIService.GetFeedV3CalendarSonarrIcs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feed/v3/calendar/sonarr.ics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pastDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pastDays", r.pastDays, "form", "")
	} else {
		var defaultValue int32 = 7
		r.pastDays = &defaultValue
	}
	if r.futureDays != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "futureDays", r.futureDays, "form", "")
	} else {
		var defaultValue int32 = 28
		r.futureDays = &defaultValue
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
	} else {
		var defaultValue string = ""
		r.tags = &defaultValue
	}
	if r.unmonitored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unmonitored", r.unmonitored, "form", "")
	} else {
		var defaultValue bool = false
		r.unmonitored = &defaultValue
	}
	if r.premieresOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "premieresOnly", r.premieresOnly, "form", "")
	} else {
		var defaultValue bool = false
		r.premieresOnly = &defaultValue
	}
	if r.asAllDay != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asAllDay", r.asAllDay, "form", "")
	} else {
		var defaultValue bool = false
		r.asAllDay = &defaultValue
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
