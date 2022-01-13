/*
MSRC CVRF API

Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// GetSecurityUpdatesApiService GetSecurityUpdatesApi service
type GetSecurityUpdatesApiService service

type ApiCvrfGetCvrfRequest struct {
	ctx        _context.Context
	ApiService *GetSecurityUpdatesApiService
	id         string
}

func (r ApiCvrfGetCvrfRequest) Execute() (Cvrfdoc, *_nethttp.Response, error) {
	return r.ApiService.CvrfGetCvrfExecute(r)
}

/*
CvrfGetCvrf Get security update details in CVRF format

Get detailed Microsoft security update, formatted according to the [Common Vulnerability Reporting
Framework](http://www.icasi.org/cvrf/). MSRC investigates all reports of
security vulnerabilities affecting Microsoft products and services, and provides these updates as part of the
ongoing effort to help you manage security risks and help keep your systems protected.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id CVRF document ID (yyyy-mmm)
 @return ApiCvrfGetCvrfRequest
*/
func (a *GetSecurityUpdatesApiService) CvrfGetCvrf(ctx _context.Context, id string) ApiCvrfGetCvrfRequest {
	return ApiCvrfGetCvrfRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return Cvrfdoc
func (a *GetSecurityUpdatesApiService) CvrfGetCvrfExecute(r ApiCvrfGetCvrfRequest) (Cvrfdoc, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue Cvrfdoc
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetSecurityUpdatesApiService.CvrfGetCvrf")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cvrf/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatesGetUpdatesRequest struct {
	ctx        _context.Context
	ApiService *GetSecurityUpdatesApiService
}

func (r ApiUpdatesGetUpdatesRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.UpdatesGetUpdatesExecute(r)
}

/*
UpdatesGetUpdates Get all security update summaries

Get a list of all Microsoft security updates. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf). This list can be manipulated using [OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356) on current and initial release dates.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatesGetUpdatesRequest
*/
func (a *GetSecurityUpdatesApiService) UpdatesGetUpdates(ctx _context.Context) ApiUpdatesGetUpdatesRequest {
	return ApiUpdatesGetUpdatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return InlineResponse200
func (a *GetSecurityUpdatesApiService) UpdatesGetUpdatesExecute(r ApiUpdatesGetUpdatesRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetSecurityUpdatesApiService.UpdatesGetUpdates")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/updates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatesGetUpdatesByKeyRequest struct {
	ctx        _context.Context
	ApiService *GetSecurityUpdatesApiService
	key        string
}

func (r ApiUpdatesGetUpdatesByKeyRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.UpdatesGetUpdatesByKeyExecute(r)
}

/*
UpdatesGetUpdatesByKey Get security update summaries by key

Get a list of Microsoft security updates by update ID, vulnerability ID ([CVE](https://cve.mitre.org/about)), or year. Each includes a link to the update details, formatted according to the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf). This list can be manipulated using [OData URL filtering](https://docs.oasis-open.org/odata/odata/v4.0/errata03/os/complete/part2-url-conventions/odata-v4.0-errata03-os-part2-url-conventions-complete.html#_Toc453752356) on current and initial release dates.

NOTE: this endpoint is case-sensitive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param key update ID (yyyy-mmm), vulnerability ID (CVE number), or year (yyyy)
 @return ApiUpdatesGetUpdatesByKeyRequest
*/
func (a *GetSecurityUpdatesApiService) UpdatesGetUpdatesByKey(ctx _context.Context, key string) ApiUpdatesGetUpdatesByKeyRequest {
	return ApiUpdatesGetUpdatesByKeyRequest{
		ApiService: a,
		ctx:        ctx,
		key:        key,
	}
}

// Execute executes the request
//  @return InlineResponse200
func (a *GetSecurityUpdatesApiService) UpdatesGetUpdatesByKeyExecute(r ApiUpdatesGetUpdatesByKeyRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GetSecurityUpdatesApiService.UpdatesGetUpdatesByKey")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Updates('{key}')"
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", _neturl.PathEscape(parameterToString(r.key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
