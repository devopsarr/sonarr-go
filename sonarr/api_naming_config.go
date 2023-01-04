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


// NamingConfigApiService NamingConfigApi service
type NamingConfigApiService service
type ApiGetConfigNamingRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
}

func (r ApiGetConfigNamingRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetConfigNamingExecute(r)
}

/*
GetConfigNaming Method for GetConfigNaming

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfigNamingRequest
*/
func (a *NamingConfigApiService) GetConfigNaming(ctx context.Context) ApiGetConfigNamingRequest {
	return ApiGetConfigNamingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) GetConfigNamingExecute(r ApiGetConfigNamingRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.GetConfigNaming")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming"

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
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
type ApiGetConfigNamingByIdRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	id int32
}

func (r ApiGetConfigNamingByIdRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetConfigNamingByIdExecute(r)
}

/*
GetConfigNamingById Method for GetConfigNamingById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetConfigNamingByIdRequest
*/
func (a *NamingConfigApiService) GetConfigNamingById(ctx context.Context, id int32) ApiGetConfigNamingByIdRequest {
	return ApiGetConfigNamingByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) GetConfigNamingByIdExecute(r ApiGetConfigNamingByIdRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.GetConfigNamingById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
type ApiGetConfigNamingExamplesRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	renameEpisodes *bool
	replaceIllegalCharacters *bool
	multiEpisodeStyle *int32
	standardEpisodeFormat *string
	dailyEpisodeFormat *string
	animeEpisodeFormat *string
	seriesFolderFormat *string
	seasonFolderFormat *string
	specialsFolderFormat *string
	includeSeriesTitle *bool
	includeEpisodeTitle *bool
	includeQuality *bool
	replaceSpaces *bool
	separator *string
	numberStyle *string
	id *int32
	resourceName *string
}

func (r ApiGetConfigNamingExamplesRequest) RenameEpisodes(renameEpisodes bool) ApiGetConfigNamingExamplesRequest {
	r.renameEpisodes = &renameEpisodes
	return r
}

func (r ApiGetConfigNamingExamplesRequest) ReplaceIllegalCharacters(replaceIllegalCharacters bool) ApiGetConfigNamingExamplesRequest {
	r.replaceIllegalCharacters = &replaceIllegalCharacters
	return r
}

func (r ApiGetConfigNamingExamplesRequest) MultiEpisodeStyle(multiEpisodeStyle int32) ApiGetConfigNamingExamplesRequest {
	r.multiEpisodeStyle = &multiEpisodeStyle
	return r
}

func (r ApiGetConfigNamingExamplesRequest) StandardEpisodeFormat(standardEpisodeFormat string) ApiGetConfigNamingExamplesRequest {
	r.standardEpisodeFormat = &standardEpisodeFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) DailyEpisodeFormat(dailyEpisodeFormat string) ApiGetConfigNamingExamplesRequest {
	r.dailyEpisodeFormat = &dailyEpisodeFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) AnimeEpisodeFormat(animeEpisodeFormat string) ApiGetConfigNamingExamplesRequest {
	r.animeEpisodeFormat = &animeEpisodeFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) SeriesFolderFormat(seriesFolderFormat string) ApiGetConfigNamingExamplesRequest {
	r.seriesFolderFormat = &seriesFolderFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) SeasonFolderFormat(seasonFolderFormat string) ApiGetConfigNamingExamplesRequest {
	r.seasonFolderFormat = &seasonFolderFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) SpecialsFolderFormat(specialsFolderFormat string) ApiGetConfigNamingExamplesRequest {
	r.specialsFolderFormat = &specialsFolderFormat
	return r
}

func (r ApiGetConfigNamingExamplesRequest) IncludeSeriesTitle(includeSeriesTitle bool) ApiGetConfigNamingExamplesRequest {
	r.includeSeriesTitle = &includeSeriesTitle
	return r
}

func (r ApiGetConfigNamingExamplesRequest) IncludeEpisodeTitle(includeEpisodeTitle bool) ApiGetConfigNamingExamplesRequest {
	r.includeEpisodeTitle = &includeEpisodeTitle
	return r
}

func (r ApiGetConfigNamingExamplesRequest) IncludeQuality(includeQuality bool) ApiGetConfigNamingExamplesRequest {
	r.includeQuality = &includeQuality
	return r
}

func (r ApiGetConfigNamingExamplesRequest) ReplaceSpaces(replaceSpaces bool) ApiGetConfigNamingExamplesRequest {
	r.replaceSpaces = &replaceSpaces
	return r
}

func (r ApiGetConfigNamingExamplesRequest) Separator(separator string) ApiGetConfigNamingExamplesRequest {
	r.separator = &separator
	return r
}

func (r ApiGetConfigNamingExamplesRequest) NumberStyle(numberStyle string) ApiGetConfigNamingExamplesRequest {
	r.numberStyle = &numberStyle
	return r
}

func (r ApiGetConfigNamingExamplesRequest) Id(id int32) ApiGetConfigNamingExamplesRequest {
	r.id = &id
	return r
}

func (r ApiGetConfigNamingExamplesRequest) ResourceName(resourceName string) ApiGetConfigNamingExamplesRequest {
	r.resourceName = &resourceName
	return r
}

func (r ApiGetConfigNamingExamplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetConfigNamingExamplesExecute(r)
}

/*
GetConfigNamingExamples Method for GetConfigNamingExamples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfigNamingExamplesRequest
*/
func (a *NamingConfigApiService) GetConfigNamingExamples(ctx context.Context) ApiGetConfigNamingExamplesRequest {
	return ApiGetConfigNamingExamplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NamingConfigApiService) GetConfigNamingExamplesExecute(r ApiGetConfigNamingExamplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.GetConfigNamingExamples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/examples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.renameEpisodes != nil {
		localVarQueryParams.Add("RenameEpisodes", parameterToString(*r.renameEpisodes, ""))
	}
	if r.replaceIllegalCharacters != nil {
		localVarQueryParams.Add("ReplaceIllegalCharacters", parameterToString(*r.replaceIllegalCharacters, ""))
	}
	if r.multiEpisodeStyle != nil {
		localVarQueryParams.Add("MultiEpisodeStyle", parameterToString(*r.multiEpisodeStyle, ""))
	}
	if r.standardEpisodeFormat != nil {
		localVarQueryParams.Add("StandardEpisodeFormat", parameterToString(*r.standardEpisodeFormat, ""))
	}
	if r.dailyEpisodeFormat != nil {
		localVarQueryParams.Add("DailyEpisodeFormat", parameterToString(*r.dailyEpisodeFormat, ""))
	}
	if r.animeEpisodeFormat != nil {
		localVarQueryParams.Add("AnimeEpisodeFormat", parameterToString(*r.animeEpisodeFormat, ""))
	}
	if r.seriesFolderFormat != nil {
		localVarQueryParams.Add("SeriesFolderFormat", parameterToString(*r.seriesFolderFormat, ""))
	}
	if r.seasonFolderFormat != nil {
		localVarQueryParams.Add("SeasonFolderFormat", parameterToString(*r.seasonFolderFormat, ""))
	}
	if r.specialsFolderFormat != nil {
		localVarQueryParams.Add("SpecialsFolderFormat", parameterToString(*r.specialsFolderFormat, ""))
	}
	if r.includeSeriesTitle != nil {
		localVarQueryParams.Add("IncludeSeriesTitle", parameterToString(*r.includeSeriesTitle, ""))
	}
	if r.includeEpisodeTitle != nil {
		localVarQueryParams.Add("IncludeEpisodeTitle", parameterToString(*r.includeEpisodeTitle, ""))
	}
	if r.includeQuality != nil {
		localVarQueryParams.Add("IncludeQuality", parameterToString(*r.includeQuality, ""))
	}
	if r.replaceSpaces != nil {
		localVarQueryParams.Add("ReplaceSpaces", parameterToString(*r.replaceSpaces, ""))
	}
	if r.separator != nil {
		localVarQueryParams.Add("Separator", parameterToString(*r.separator, ""))
	}
	if r.numberStyle != nil {
		localVarQueryParams.Add("NumberStyle", parameterToString(*r.numberStyle, ""))
	}
	if r.id != nil {
		localVarQueryParams.Add("Id", parameterToString(*r.id, ""))
	}
	if r.resourceName != nil {
		localVarQueryParams.Add("ResourceName", parameterToString(*r.resourceName, ""))
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
type ApiUpdateConfigNamingRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	id string
	namingConfigResource *NamingConfigResource
}

func (r ApiUpdateConfigNamingRequest) NamingConfigResource(namingConfigResource NamingConfigResource) ApiUpdateConfigNamingRequest {
	r.namingConfigResource = &namingConfigResource
	return r
}

func (r ApiUpdateConfigNamingRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.UpdateConfigNamingExecute(r)
}

/*
UpdateConfigNaming Method for UpdateConfigNaming

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateConfigNamingRequest
*/
func (a *NamingConfigApiService) UpdateConfigNaming(ctx context.Context, id string) ApiUpdateConfigNamingRequest {
	return ApiUpdateConfigNamingRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) UpdateConfigNamingExecute(r ApiUpdateConfigNamingRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.UpdateConfigNaming")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

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
	localVarPostBody = r.namingConfigResource
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
