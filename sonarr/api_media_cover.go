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
	"strings"
)


// MediaCoverApiService MediaCoverApi service
type MediaCoverApiService service
type ApiGetMediaCoverseriesIdByFilenameRequest struct {
	ctx context.Context
	ApiService *MediaCoverApiService
	seriesId int32
	filename string
}

func (r ApiGetMediaCoverseriesIdByFilenameRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetMediaCoverseriesIdByFilenameExecute(r)
}

/*
GetMediaCoverseriesIdByFilename Method for GetMediaCoverseriesIdByFilename

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seriesId
 @param filename
 @return ApiGetMediaCoverseriesIdByFilenameRequest
*/
func (a *MediaCoverApiService) GetMediaCoverseriesIdByFilename(ctx context.Context, seriesId int32, filename string) ApiGetMediaCoverseriesIdByFilenameRequest {
	return ApiGetMediaCoverseriesIdByFilenameRequest{
		ApiService: a,
		ctx: ctx,
		seriesId: seriesId,
		filename: filename,
	}
}

// Execute executes the request
func (a *MediaCoverApiService) GetMediaCoverseriesIdByFilenameExecute(r ApiGetMediaCoverseriesIdByFilenameRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaCoverApiService.GetMediaCoverseriesIdByFilename")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/mediacover/{seriesId}/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"seriesId"+"}", url.PathEscape(parameterToString(r.seriesId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterToString(r.filename, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
