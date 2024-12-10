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
	"strings"
)


// NamingConfigAPIService NamingConfigAPI service
type NamingConfigAPIService service

type ApiGetNamingConfigRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
}

func (r ApiGetNamingConfigRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetNamingConfigExecute(r)
}

/*
GetNamingConfig Method for GetNamingConfig

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNamingConfigRequest
*/
func (a *NamingConfigAPIService) GetNamingConfig(ctx context.Context) ApiGetNamingConfigRequest {
	return ApiGetNamingConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) GetNamingConfigExecute(r ApiGetNamingConfigRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfig")
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

type ApiGetNamingConfigByIdRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	id int32
}

func (r ApiGetNamingConfigByIdRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.GetNamingConfigByIdExecute(r)
}

/*
GetNamingConfigById Method for GetNamingConfigById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetNamingConfigByIdRequest
*/
func (a *NamingConfigAPIService) GetNamingConfigById(ctx context.Context, id int32) ApiGetNamingConfigByIdRequest {
	return ApiGetNamingConfigByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) GetNamingConfigByIdExecute(r ApiGetNamingConfigByIdRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfigById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
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

type ApiGetNamingConfigExamplesRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	renameEpisodes *bool
	replaceIllegalCharacters *bool
	colonReplacementFormat *int32
	customColonReplacementFormat *string
	multiEpisodeStyle *int32
	standardEpisodeFormat *string
	dailyEpisodeFormat *string
	animeEpisodeFormat *string
	seriesFolderFormat *string
	seasonFolderFormat *string
	specialsFolderFormat *string
	id *int32
	resourceName *string
}

func (r ApiGetNamingConfigExamplesRequest) RenameEpisodes(renameEpisodes bool) ApiGetNamingConfigExamplesRequest {
	r.renameEpisodes = &renameEpisodes
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ReplaceIllegalCharacters(replaceIllegalCharacters bool) ApiGetNamingConfigExamplesRequest {
	r.replaceIllegalCharacters = &replaceIllegalCharacters
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ColonReplacementFormat(colonReplacementFormat int32) ApiGetNamingConfigExamplesRequest {
	r.colonReplacementFormat = &colonReplacementFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) CustomColonReplacementFormat(customColonReplacementFormat string) ApiGetNamingConfigExamplesRequest {
	r.customColonReplacementFormat = &customColonReplacementFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) MultiEpisodeStyle(multiEpisodeStyle int32) ApiGetNamingConfigExamplesRequest {
	r.multiEpisodeStyle = &multiEpisodeStyle
	return r
}

func (r ApiGetNamingConfigExamplesRequest) StandardEpisodeFormat(standardEpisodeFormat string) ApiGetNamingConfigExamplesRequest {
	r.standardEpisodeFormat = &standardEpisodeFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) DailyEpisodeFormat(dailyEpisodeFormat string) ApiGetNamingConfigExamplesRequest {
	r.dailyEpisodeFormat = &dailyEpisodeFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) AnimeEpisodeFormat(animeEpisodeFormat string) ApiGetNamingConfigExamplesRequest {
	r.animeEpisodeFormat = &animeEpisodeFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) SeriesFolderFormat(seriesFolderFormat string) ApiGetNamingConfigExamplesRequest {
	r.seriesFolderFormat = &seriesFolderFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) SeasonFolderFormat(seasonFolderFormat string) ApiGetNamingConfigExamplesRequest {
	r.seasonFolderFormat = &seasonFolderFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) SpecialsFolderFormat(specialsFolderFormat string) ApiGetNamingConfigExamplesRequest {
	r.specialsFolderFormat = &specialsFolderFormat
	return r
}

func (r ApiGetNamingConfigExamplesRequest) Id(id int32) ApiGetNamingConfigExamplesRequest {
	r.id = &id
	return r
}

func (r ApiGetNamingConfigExamplesRequest) ResourceName(resourceName string) ApiGetNamingConfigExamplesRequest {
	r.resourceName = &resourceName
	return r
}

func (r ApiGetNamingConfigExamplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetNamingConfigExamplesExecute(r)
}

/*
GetNamingConfigExamples Method for GetNamingConfigExamples

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNamingConfigExamplesRequest
*/
func (a *NamingConfigAPIService) GetNamingConfigExamples(ctx context.Context) ApiGetNamingConfigExamplesRequest {
	return ApiGetNamingConfigExamplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NamingConfigAPIService) GetNamingConfigExamplesExecute(r ApiGetNamingConfigExamplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.GetNamingConfigExamples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/examples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.renameEpisodes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "renameEpisodes", r.renameEpisodes, "form", "")
	}
	if r.replaceIllegalCharacters != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "replaceIllegalCharacters", r.replaceIllegalCharacters, "form", "")
	}
	if r.colonReplacementFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "colonReplacementFormat", r.colonReplacementFormat, "form", "")
	}
	if r.customColonReplacementFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customColonReplacementFormat", r.customColonReplacementFormat, "form", "")
	}
	if r.multiEpisodeStyle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "multiEpisodeStyle", r.multiEpisodeStyle, "form", "")
	}
	if r.standardEpisodeFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "standardEpisodeFormat", r.standardEpisodeFormat, "form", "")
	}
	if r.dailyEpisodeFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dailyEpisodeFormat", r.dailyEpisodeFormat, "form", "")
	}
	if r.animeEpisodeFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "animeEpisodeFormat", r.animeEpisodeFormat, "form", "")
	}
	if r.seriesFolderFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seriesFolderFormat", r.seriesFolderFormat, "form", "")
	}
	if r.seasonFolderFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seasonFolderFormat", r.seasonFolderFormat, "form", "")
	}
	if r.specialsFolderFormat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "specialsFolderFormat", r.specialsFolderFormat, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.resourceName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resourceName", r.resourceName, "form", "")
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

type ApiUpdateNamingConfigRequest struct {
	ctx context.Context
	ApiService *NamingConfigAPIService
	id string
	namingConfigResource *NamingConfigResource
}

func (r ApiUpdateNamingConfigRequest) NamingConfigResource(namingConfigResource NamingConfigResource) ApiUpdateNamingConfigRequest {
	r.namingConfigResource = &namingConfigResource
	return r
}

func (r ApiUpdateNamingConfigRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.UpdateNamingConfigExecute(r)
}

/*
UpdateNamingConfig Method for UpdateNamingConfig

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateNamingConfigRequest
*/
func (a *NamingConfigAPIService) UpdateNamingConfig(ctx context.Context, id string) ApiUpdateNamingConfigRequest {
	return ApiUpdateNamingConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigAPIService) UpdateNamingConfigExecute(r ApiUpdateNamingConfigRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigAPIService.UpdateNamingConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
