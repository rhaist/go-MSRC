# HttpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpContext** | Pointer to [**HttpContext**](HttpContext.md) |  | [optional] 
**Method** | Pointer to **NullableString** |  | [optional] 
**Scheme** | Pointer to **NullableString** |  | [optional] 
**IsHttps** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to [**HostString**](HostString.md) |  | [optional] 
**PathBase** | Pointer to [**PathString**](PathString.md) |  | [optional] 
**Path** | Pointer to [**PathString**](PathString.md) |  | [optional] 
**QueryString** | Pointer to [**QueryString**](QueryString.md) |  | [optional] 
**Query** | Pointer to [**[]StringStringValuesKeyValuePair**](StringStringValuesKeyValuePair.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] [readonly] 
**Cookies** | Pointer to [**[]StringStringKeyValuePair**](StringStringKeyValuePair.md) |  | [optional] 
**ContentLength** | Pointer to **NullableInt64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to [**Stream**](Stream.md) |  | [optional] 
**BodyReader** | Pointer to **map[string]interface{}** |  | [optional] 
**HasFormContentType** | Pointer to **bool** |  | [optional] [readonly] 
**Form** | Pointer to [**[]StringStringValuesKeyValuePair**](StringStringValuesKeyValuePair.md) |  | [optional] 
**RouteValues** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewHttpRequest

`func NewHttpRequest() *HttpRequest`

NewHttpRequest instantiates a new HttpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpRequestWithDefaults

`func NewHttpRequestWithDefaults() *HttpRequest`

NewHttpRequestWithDefaults instantiates a new HttpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpContext

`func (o *HttpRequest) GetHttpContext() HttpContext`

GetHttpContext returns the HttpContext field if non-nil, zero value otherwise.

### GetHttpContextOk

`func (o *HttpRequest) GetHttpContextOk() (*HttpContext, bool)`

GetHttpContextOk returns a tuple with the HttpContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContext

`func (o *HttpRequest) SetHttpContext(v HttpContext)`

SetHttpContext sets HttpContext field to given value.

### HasHttpContext

`func (o *HttpRequest) HasHttpContext() bool`

HasHttpContext returns a boolean if a field has been set.

### GetMethod

`func (o *HttpRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *HttpRequest) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *HttpRequest) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetScheme

`func (o *HttpRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *HttpRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *HttpRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *HttpRequest) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *HttpRequest) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *HttpRequest) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetIsHttps

`func (o *HttpRequest) GetIsHttps() bool`

GetIsHttps returns the IsHttps field if non-nil, zero value otherwise.

### GetIsHttpsOk

`func (o *HttpRequest) GetIsHttpsOk() (*bool, bool)`

GetIsHttpsOk returns a tuple with the IsHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHttps

`func (o *HttpRequest) SetIsHttps(v bool)`

SetIsHttps sets IsHttps field to given value.

### HasIsHttps

`func (o *HttpRequest) HasIsHttps() bool`

HasIsHttps returns a boolean if a field has been set.

### GetHost

`func (o *HttpRequest) GetHost() HostString`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpRequest) GetHostOk() (*HostString, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpRequest) SetHost(v HostString)`

SetHost sets Host field to given value.

### HasHost

`func (o *HttpRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPathBase

`func (o *HttpRequest) GetPathBase() PathString`

GetPathBase returns the PathBase field if non-nil, zero value otherwise.

### GetPathBaseOk

`func (o *HttpRequest) GetPathBaseOk() (*PathString, bool)`

GetPathBaseOk returns a tuple with the PathBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathBase

`func (o *HttpRequest) SetPathBase(v PathString)`

SetPathBase sets PathBase field to given value.

### HasPathBase

`func (o *HttpRequest) HasPathBase() bool`

HasPathBase returns a boolean if a field has been set.

### GetPath

`func (o *HttpRequest) GetPath() PathString`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HttpRequest) GetPathOk() (*PathString, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HttpRequest) SetPath(v PathString)`

SetPath sets Path field to given value.

### HasPath

`func (o *HttpRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQueryString

`func (o *HttpRequest) GetQueryString() QueryString`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *HttpRequest) GetQueryStringOk() (*QueryString, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *HttpRequest) SetQueryString(v QueryString)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *HttpRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetQuery

`func (o *HttpRequest) GetQuery() []StringStringValuesKeyValuePair`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *HttpRequest) GetQueryOk() (*[]StringStringValuesKeyValuePair, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *HttpRequest) SetQuery(v []StringStringValuesKeyValuePair)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *HttpRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *HttpRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *HttpRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetProtocol

`func (o *HttpRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HttpRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HttpRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *HttpRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *HttpRequest) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *HttpRequest) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetHeaders

`func (o *HttpRequest) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpRequest) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpRequest) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *HttpRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *HttpRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetCookies

`func (o *HttpRequest) GetCookies() []StringStringKeyValuePair`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *HttpRequest) GetCookiesOk() (*[]StringStringKeyValuePair, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *HttpRequest) SetCookies(v []StringStringKeyValuePair)`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *HttpRequest) HasCookies() bool`

HasCookies returns a boolean if a field has been set.

### SetCookiesNil

`func (o *HttpRequest) SetCookiesNil(b bool)`

 SetCookiesNil sets the value for Cookies to be an explicit nil

### UnsetCookies
`func (o *HttpRequest) UnsetCookies()`

UnsetCookies ensures that no value is present for Cookies, not even an explicit nil
### GetContentLength

`func (o *HttpRequest) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *HttpRequest) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *HttpRequest) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *HttpRequest) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *HttpRequest) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *HttpRequest) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *HttpRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HttpRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HttpRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HttpRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *HttpRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *HttpRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetBody

`func (o *HttpRequest) GetBody() Stream`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpRequest) GetBodyOk() (*Stream, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpRequest) SetBody(v Stream)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyReader

`func (o *HttpRequest) GetBodyReader() map[string]interface{}`

GetBodyReader returns the BodyReader field if non-nil, zero value otherwise.

### GetBodyReaderOk

`func (o *HttpRequest) GetBodyReaderOk() (*map[string]interface{}, bool)`

GetBodyReaderOk returns a tuple with the BodyReader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyReader

`func (o *HttpRequest) SetBodyReader(v map[string]interface{})`

SetBodyReader sets BodyReader field to given value.

### HasBodyReader

`func (o *HttpRequest) HasBodyReader() bool`

HasBodyReader returns a boolean if a field has been set.

### GetHasFormContentType

`func (o *HttpRequest) GetHasFormContentType() bool`

GetHasFormContentType returns the HasFormContentType field if non-nil, zero value otherwise.

### GetHasFormContentTypeOk

`func (o *HttpRequest) GetHasFormContentTypeOk() (*bool, bool)`

GetHasFormContentTypeOk returns a tuple with the HasFormContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFormContentType

`func (o *HttpRequest) SetHasFormContentType(v bool)`

SetHasFormContentType sets HasFormContentType field to given value.

### HasHasFormContentType

`func (o *HttpRequest) HasHasFormContentType() bool`

HasHasFormContentType returns a boolean if a field has been set.

### GetForm

`func (o *HttpRequest) GetForm() []StringStringValuesKeyValuePair`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *HttpRequest) GetFormOk() (*[]StringStringValuesKeyValuePair, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *HttpRequest) SetForm(v []StringStringValuesKeyValuePair)`

SetForm sets Form field to given value.

### HasForm

`func (o *HttpRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *HttpRequest) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *HttpRequest) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil
### GetRouteValues

`func (o *HttpRequest) GetRouteValues() map[string]interface{}`

GetRouteValues returns the RouteValues field if non-nil, zero value otherwise.

### GetRouteValuesOk

`func (o *HttpRequest) GetRouteValuesOk() (*map[string]interface{}, bool)`

GetRouteValuesOk returns a tuple with the RouteValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteValues

`func (o *HttpRequest) SetRouteValues(v map[string]interface{})`

SetRouteValues sets RouteValues field to given value.

### HasRouteValues

`func (o *HttpRequest) HasRouteValues() bool`

HasRouteValues returns a boolean if a field has been set.

### SetRouteValuesNil

`func (o *HttpRequest) SetRouteValuesNil(b bool)`

 SetRouteValuesNil sets the value for RouteValues to be an explicit nil

### UnsetRouteValues
`func (o *HttpRequest) UnsetRouteValues()`

UnsetRouteValues ensures that no value is present for RouteValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


