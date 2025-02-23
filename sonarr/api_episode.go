/*
Sonarr

Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.

API version: v4.0.13.2932
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
	"reflect"
)


// EpisodeAPIService EpisodeAPI service
type EpisodeAPIService service

type ApiGetEpisodeByIdRequest struct {
	ctx context.Context
	ApiService *EpisodeAPIService
	id int32
}

func (r ApiGetEpisodeByIdRequest) Execute() (*EpisodeResource, *http.Response, error) {
	return r.ApiService.GetEpisodeByIdExecute(r)
}

/*
GetEpisodeById Method for GetEpisodeById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetEpisodeByIdRequest
*/
func (a *EpisodeAPIService) GetEpisodeById(ctx context.Context, id int32) ApiGetEpisodeByIdRequest {
	return ApiGetEpisodeByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EpisodeResource
func (a *EpisodeAPIService) GetEpisodeByIdExecute(r ApiGetEpisodeByIdRequest) (*EpisodeResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EpisodeResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EpisodeAPIService.GetEpisodeById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/episode/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListEpisodeRequest struct {
	ctx context.Context
	ApiService *EpisodeAPIService
	seriesId *int32
	seasonNumber *int32
	episodeIds *[]int32
	episodeFileId *int32
	includeSeries *bool
	includeEpisodeFile *bool
	includeImages *bool
}

func (r ApiListEpisodeRequest) SeriesId(seriesId int32) ApiListEpisodeRequest {
	r.seriesId = &seriesId
	return r
}

func (r ApiListEpisodeRequest) SeasonNumber(seasonNumber int32) ApiListEpisodeRequest {
	r.seasonNumber = &seasonNumber
	return r
}

func (r ApiListEpisodeRequest) EpisodeIds(episodeIds []int32) ApiListEpisodeRequest {
	r.episodeIds = &episodeIds
	return r
}

func (r ApiListEpisodeRequest) EpisodeFileId(episodeFileId int32) ApiListEpisodeRequest {
	r.episodeFileId = &episodeFileId
	return r
}

func (r ApiListEpisodeRequest) IncludeSeries(includeSeries bool) ApiListEpisodeRequest {
	r.includeSeries = &includeSeries
	return r
}

func (r ApiListEpisodeRequest) IncludeEpisodeFile(includeEpisodeFile bool) ApiListEpisodeRequest {
	r.includeEpisodeFile = &includeEpisodeFile
	return r
}

func (r ApiListEpisodeRequest) IncludeImages(includeImages bool) ApiListEpisodeRequest {
	r.includeImages = &includeImages
	return r
}

func (r ApiListEpisodeRequest) Execute() ([]EpisodeResource, *http.Response, error) {
	return r.ApiService.ListEpisodeExecute(r)
}

/*
ListEpisode Method for ListEpisode

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListEpisodeRequest
*/
func (a *EpisodeAPIService) ListEpisode(ctx context.Context) ApiListEpisodeRequest {
	return ApiListEpisodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []EpisodeResource
func (a *EpisodeAPIService) ListEpisodeExecute(r ApiListEpisodeRequest) ([]EpisodeResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EpisodeResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EpisodeAPIService.ListEpisode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/episode"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.seriesId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seriesId", r.seriesId, "form", "")
	}
	if r.seasonNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seasonNumber", r.seasonNumber, "form", "")
	}
	if r.episodeIds != nil {
		t := *r.episodeIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "episodeIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "episodeIds", t, "form", "multi")
		}
	}
	if r.episodeFileId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "episodeFileId", r.episodeFileId, "form", "")
	}
	if r.includeSeries != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeSeries", r.includeSeries, "form", "")
	} else {
		var defaultValue bool = false
		r.includeSeries = &defaultValue
	}
	if r.includeEpisodeFile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeEpisodeFile", r.includeEpisodeFile, "form", "")
	} else {
		var defaultValue bool = false
		r.includeEpisodeFile = &defaultValue
	}
	if r.includeImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeImages", r.includeImages, "form", "")
	} else {
		var defaultValue bool = false
		r.includeImages = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutEpisodeMonitorRequest struct {
	ctx context.Context
	ApiService *EpisodeAPIService
	includeImages *bool
	episodesMonitoredResource *EpisodesMonitoredResource
}

func (r ApiPutEpisodeMonitorRequest) IncludeImages(includeImages bool) ApiPutEpisodeMonitorRequest {
	r.includeImages = &includeImages
	return r
}

func (r ApiPutEpisodeMonitorRequest) EpisodesMonitoredResource(episodesMonitoredResource EpisodesMonitoredResource) ApiPutEpisodeMonitorRequest {
	r.episodesMonitoredResource = &episodesMonitoredResource
	return r
}

func (r ApiPutEpisodeMonitorRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutEpisodeMonitorExecute(r)
}

/*
PutEpisodeMonitor Method for PutEpisodeMonitor

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPutEpisodeMonitorRequest
*/
func (a *EpisodeAPIService) PutEpisodeMonitor(ctx context.Context) ApiPutEpisodeMonitorRequest {
	return ApiPutEpisodeMonitorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EpisodeAPIService) PutEpisodeMonitorExecute(r ApiPutEpisodeMonitorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EpisodeAPIService.PutEpisodeMonitor")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/episode/monitor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeImages", r.includeImages, "form", "")
	} else {
		var defaultValue bool = false
		r.includeImages = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.episodesMonitoredResource
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

type ApiUpdateEpisodeRequest struct {
	ctx context.Context
	ApiService *EpisodeAPIService
	id int32
	episodeResource *EpisodeResource
}

func (r ApiUpdateEpisodeRequest) EpisodeResource(episodeResource EpisodeResource) ApiUpdateEpisodeRequest {
	r.episodeResource = &episodeResource
	return r
}

func (r ApiUpdateEpisodeRequest) Execute() (*EpisodeResource, *http.Response, error) {
	return r.ApiService.UpdateEpisodeExecute(r)
}

/*
UpdateEpisode Method for UpdateEpisode

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateEpisodeRequest
*/
func (a *EpisodeAPIService) UpdateEpisode(ctx context.Context, id int32) ApiUpdateEpisodeRequest {
	return ApiUpdateEpisodeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EpisodeResource
func (a *EpisodeAPIService) UpdateEpisodeExecute(r ApiUpdateEpisodeRequest) (*EpisodeResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EpisodeResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EpisodeAPIService.UpdateEpisode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/episode/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.episodeResource
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
