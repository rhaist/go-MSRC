/*
MSRC CVRF API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// UpdatesAPIService UpdatesAPI service
type UpdatesAPIService service

type ApiUpdatesRequest struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	options    *UpdateODataQueryOptions
}

// OData query options
func (r ApiUpdatesRequest) Options(options UpdateODataQueryOptions) ApiUpdatesRequest {
	r.options = &options
	return r
}

func (r ApiUpdatesRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesExecute(r)
}

/*
Updates Get all security update summaries

Get a list of all Microsoft security updates. Each includes a link to the update details, formatted according to
the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesRequest
*/
func (a *UpdatesAPIService) Updates(ctx context.Context) ApiUpdatesRequest {
	return ApiUpdatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesExecute(r ApiUpdatesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.Updates")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesCountRequest struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	options    *UpdateODataQueryOptions
}

// OData query options
func (r ApiUpdatesCountRequest) Options(options UpdateODataQueryOptions) ApiUpdatesCountRequest {
	r.options = &options
	return r
}

func (r ApiUpdatesCountRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesCountExecute(r)
}

/*
UpdatesCount Get all security update summaries

Get a list of all Microsoft security updates. Each includes a link to the update details, formatted according to
the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesCountRequest
*/
func (a *UpdatesAPIService) UpdatesCount(ctx context.Context) ApiUpdatesCountRequest {
	return ApiUpdatesCountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesCountExecute(r ApiUpdatesCountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesCount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Updates/$count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesGetRequest struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	options    *UpdateODataQueryOptions
}

// OData query options
func (r ApiUpdatesGetRequest) Options(options UpdateODataQueryOptions) ApiUpdatesGetRequest {
	r.options = &options
	return r
}

func (r ApiUpdatesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesGetExecute(r)
}

/*
UpdatesGet Get all security update summaries

Get a list of all Microsoft security updates. Each includes a link to the update details, formatted according to
the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesGetRequest
*/
func (a *UpdatesAPIService) UpdatesGet(ctx context.Context) ApiUpdatesGetRequest {
	return ApiUpdatesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesGetExecute(r ApiUpdatesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesKeyRequest struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	key        *string
	options    *UpdateODataQueryOptions
}

// update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)
func (r ApiUpdatesKeyRequest) Key(key string) ApiUpdatesKeyRequest {
	r.key = &key
	return r
}

// OData query options
func (r ApiUpdatesKeyRequest) Options(options UpdateODataQueryOptions) ApiUpdatesKeyRequest {
	r.options = &options
	return r
}

func (r ApiUpdatesKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesKeyExecute(r)
}

/*
UpdatesKey Get security update summaries by key

Get a list of Microsoft security updates by update ID, vulnerability ID ([CVE](https://cve.mitre.org/about)), or
year. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting
Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesKeyRequest
*/
func (a *UpdatesAPIService) UpdatesKey(ctx context.Context) ApiUpdatesKeyRequest {
	return ApiUpdatesKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesKeyExecute(r ApiUpdatesKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Updates({key})"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "form", "")
	}
	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesKeyGetRequest struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	key        *string
	options    *UpdateODataQueryOptions
}

// update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)
func (r ApiUpdatesKeyGetRequest) Key(key string) ApiUpdatesKeyGetRequest {
	r.key = &key
	return r
}

// OData query options
func (r ApiUpdatesKeyGetRequest) Options(options UpdateODataQueryOptions) ApiUpdatesKeyGetRequest {
	r.options = &options
	return r
}

func (r ApiUpdatesKeyGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesKeyGetExecute(r)
}

/*
UpdatesKeyGet Get security update summaries by key

Get a list of Microsoft security updates by update ID, vulnerability ID ([CVE](https://cve.mitre.org/about)), or
year. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting
Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesKeyGetRequest
*/
func (a *UpdatesAPIService) UpdatesKeyGet(ctx context.Context) ApiUpdatesKeyGetRequest {
	return ApiUpdatesKeyGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesKeyGetExecute(r ApiUpdatesKeyGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesKeyGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/updates({key})"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "form", "")
	}
	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesKeyGet_0Request struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	key        *string
	options    *UpdateODataQueryOptions
}

// update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)
func (r ApiUpdatesKeyGet_0Request) Key(key string) ApiUpdatesKeyGet_0Request {
	r.key = &key
	return r
}

// OData query options
func (r ApiUpdatesKeyGet_0Request) Options(options UpdateODataQueryOptions) ApiUpdatesKeyGet_0Request {
	r.options = &options
	return r
}

func (r ApiUpdatesKeyGet_0Request) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesKeyGet_1Execute(r)
}

/*
UpdatesKeyGet_0 Get security update summaries by key

Get a list of Microsoft security updates by update ID, vulnerability ID ([CVE](https://cve.mitre.org/about)), or
year. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting
Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesKeyGet_0Request
*/
func (a *UpdatesAPIService) UpdatesKeyGet_1(ctx context.Context) ApiUpdatesKeyGet_0Request {
	return ApiUpdatesKeyGet_0Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesKeyGet_1Execute(r ApiUpdatesKeyGet_0Request) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesKeyGet_1")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/updates/{key}"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "form", "")
	}
	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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

type ApiUpdatesKey_0Request struct {
	ctx        context.Context
	ApiService *UpdatesAPIService
	key        *string
	options    *UpdateODataQueryOptions
}

// update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)
func (r ApiUpdatesKey_0Request) Key(key string) ApiUpdatesKey_0Request {
	r.key = &key
	return r
}

// OData query options
func (r ApiUpdatesKey_0Request) Options(options UpdateODataQueryOptions) ApiUpdatesKey_0Request {
	r.options = &options
	return r
}

func (r ApiUpdatesKey_0Request) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesKey_2Execute(r)
}

/*
UpdatesKey_0 Get security update summaries by key

Get a list of Microsoft security updates by update ID, vulnerability ID ([CVE](https://cve.mitre.org/about)), or
year. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting
Framework](https://www.icasi.org/cvrf). This list can be manipulated using
[OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356)
on current and initial release dates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesKey_0Request
*/
func (a *UpdatesAPIService) UpdatesKey_2(ctx context.Context) ApiUpdatesKey_0Request {
	return ApiUpdatesKey_0Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UpdatesAPIService) UpdatesKey_2Execute(r ApiUpdatesKey_0Request) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdatesAPIService.UpdatesKey_2")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Updates/{key}"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.key != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "form", "")
	}
	if r.options != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "options", r.options, "form", "")
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
