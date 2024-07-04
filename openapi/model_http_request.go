/*
MSRC CVRF API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HttpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpRequest{}

// HttpRequest struct for HttpRequest
type HttpRequest struct {
	HttpContext        *HttpContext                     `json:"httpContext,omitempty"`
	Method             NullableString                   `json:"method,omitempty"`
	Scheme             NullableString                   `json:"scheme,omitempty"`
	IsHttps            *bool                            `json:"isHttps,omitempty"`
	Host               *HostString                      `json:"host,omitempty"`
	PathBase           *PathString                      `json:"pathBase,omitempty"`
	Path               *PathString                      `json:"path,omitempty"`
	QueryString        *QueryString                     `json:"queryString,omitempty"`
	Query              []StringStringValuesKeyValuePair `json:"query,omitempty"`
	Protocol           NullableString                   `json:"protocol,omitempty"`
	Headers            map[string][]string              `json:"headers,omitempty"`
	Cookies            []StringStringKeyValuePair       `json:"cookies,omitempty"`
	ContentLength      NullableInt64                    `json:"contentLength,omitempty"`
	ContentType        NullableString                   `json:"contentType,omitempty"`
	Body               *Stream                          `json:"body,omitempty"`
	BodyReader         map[string]interface{}           `json:"bodyReader,omitempty"`
	HasFormContentType *bool                            `json:"hasFormContentType,omitempty"`
	Form               []StringStringValuesKeyValuePair `json:"form,omitempty"`
	RouteValues        map[string]interface{}           `json:"routeValues,omitempty"`
}

// NewHttpRequest instantiates a new HttpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpRequest() *HttpRequest {
	this := HttpRequest{}
	return &this
}

// NewHttpRequestWithDefaults instantiates a new HttpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpRequestWithDefaults() *HttpRequest {
	this := HttpRequest{}
	return &this
}

// GetHttpContext returns the HttpContext field value if set, zero value otherwise.
func (o *HttpRequest) GetHttpContext() HttpContext {
	if o == nil || IsNil(o.HttpContext) {
		var ret HttpContext
		return ret
	}
	return *o.HttpContext
}

// GetHttpContextOk returns a tuple with the HttpContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetHttpContextOk() (*HttpContext, bool) {
	if o == nil || IsNil(o.HttpContext) {
		return nil, false
	}
	return o.HttpContext, true
}

// HasHttpContext returns a boolean if a field has been set.
func (o *HttpRequest) HasHttpContext() bool {
	if o != nil && !IsNil(o.HttpContext) {
		return true
	}

	return false
}

// SetHttpContext gets a reference to the given HttpContext and assigns it to the HttpContext field.
func (o *HttpRequest) SetHttpContext(v HttpContext) {
	o.HttpContext = &v
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *HttpRequest) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *HttpRequest) SetMethod(v string) {
	o.Method.Set(&v)
}

// SetMethodNil sets the value for Method to be an explicit nil
func (o *HttpRequest) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *HttpRequest) UnsetMethod() {
	o.Method.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetScheme() string {
	if o == nil || IsNil(o.Scheme.Get()) {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *HttpRequest) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *HttpRequest) SetScheme(v string) {
	o.Scheme.Set(&v)
}

// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *HttpRequest) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *HttpRequest) UnsetScheme() {
	o.Scheme.Unset()
}

// GetIsHttps returns the IsHttps field value if set, zero value otherwise.
func (o *HttpRequest) GetIsHttps() bool {
	if o == nil || IsNil(o.IsHttps) {
		var ret bool
		return ret
	}
	return *o.IsHttps
}

// GetIsHttpsOk returns a tuple with the IsHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetIsHttpsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHttps) {
		return nil, false
	}
	return o.IsHttps, true
}

// HasIsHttps returns a boolean if a field has been set.
func (o *HttpRequest) HasIsHttps() bool {
	if o != nil && !IsNil(o.IsHttps) {
		return true
	}

	return false
}

// SetIsHttps gets a reference to the given bool and assigns it to the IsHttps field.
func (o *HttpRequest) SetIsHttps(v bool) {
	o.IsHttps = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HttpRequest) GetHost() HostString {
	if o == nil || IsNil(o.Host) {
		var ret HostString
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetHostOk() (*HostString, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HttpRequest) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given HostString and assigns it to the Host field.
func (o *HttpRequest) SetHost(v HostString) {
	o.Host = &v
}

// GetPathBase returns the PathBase field value if set, zero value otherwise.
func (o *HttpRequest) GetPathBase() PathString {
	if o == nil || IsNil(o.PathBase) {
		var ret PathString
		return ret
	}
	return *o.PathBase
}

// GetPathBaseOk returns a tuple with the PathBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetPathBaseOk() (*PathString, bool) {
	if o == nil || IsNil(o.PathBase) {
		return nil, false
	}
	return o.PathBase, true
}

// HasPathBase returns a boolean if a field has been set.
func (o *HttpRequest) HasPathBase() bool {
	if o != nil && !IsNil(o.PathBase) {
		return true
	}

	return false
}

// SetPathBase gets a reference to the given PathString and assigns it to the PathBase field.
func (o *HttpRequest) SetPathBase(v PathString) {
	o.PathBase = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HttpRequest) GetPath() PathString {
	if o == nil || IsNil(o.Path) {
		var ret PathString
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetPathOk() (*PathString, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HttpRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given PathString and assigns it to the Path field.
func (o *HttpRequest) SetPath(v PathString) {
	o.Path = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *HttpRequest) GetQueryString() QueryString {
	if o == nil || IsNil(o.QueryString) {
		var ret QueryString
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetQueryStringOk() (*QueryString, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *HttpRequest) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given QueryString and assigns it to the QueryString field.
func (o *HttpRequest) SetQueryString(v QueryString) {
	o.QueryString = &v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetQuery() []StringStringValuesKeyValuePair {
	if o == nil {
		var ret []StringStringValuesKeyValuePair
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetQueryOk() ([]StringStringValuesKeyValuePair, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *HttpRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given []StringStringValuesKeyValuePair and assigns it to the Query field.
func (o *HttpRequest) SetQuery(v []StringStringValuesKeyValuePair) {
	o.Query = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *HttpRequest) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *HttpRequest) SetProtocol(v string) {
	o.Protocol.Set(&v)
}

// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *HttpRequest) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *HttpRequest) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetHeaders() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HttpRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *HttpRequest) SetHeaders(v map[string][]string) {
	o.Headers = v
}

// GetCookies returns the Cookies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetCookies() []StringStringKeyValuePair {
	if o == nil {
		var ret []StringStringKeyValuePair
		return ret
	}
	return o.Cookies
}

// GetCookiesOk returns a tuple with the Cookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetCookiesOk() ([]StringStringKeyValuePair, bool) {
	if o == nil || IsNil(o.Cookies) {
		return nil, false
	}
	return o.Cookies, true
}

// HasCookies returns a boolean if a field has been set.
func (o *HttpRequest) HasCookies() bool {
	if o != nil && !IsNil(o.Cookies) {
		return true
	}

	return false
}

// SetCookies gets a reference to the given []StringStringKeyValuePair and assigns it to the Cookies field.
func (o *HttpRequest) SetCookies(v []StringStringKeyValuePair) {
	o.Cookies = v
}

// GetContentLength returns the ContentLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetContentLength() int64 {
	if o == nil || IsNil(o.ContentLength.Get()) {
		var ret int64
		return ret
	}
	return *o.ContentLength.Get()
}

// GetContentLengthOk returns a tuple with the ContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetContentLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentLength.Get(), o.ContentLength.IsSet()
}

// HasContentLength returns a boolean if a field has been set.
func (o *HttpRequest) HasContentLength() bool {
	if o != nil && o.ContentLength.IsSet() {
		return true
	}

	return false
}

// SetContentLength gets a reference to the given NullableInt64 and assigns it to the ContentLength field.
func (o *HttpRequest) SetContentLength(v int64) {
	o.ContentLength.Set(&v)
}

// SetContentLengthNil sets the value for ContentLength to be an explicit nil
func (o *HttpRequest) SetContentLengthNil() {
	o.ContentLength.Set(nil)
}

// UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
func (o *HttpRequest) UnsetContentLength() {
	o.ContentLength.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *HttpRequest) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *HttpRequest) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *HttpRequest) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *HttpRequest) UnsetContentType() {
	o.ContentType.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *HttpRequest) GetBody() Stream {
	if o == nil || IsNil(o.Body) {
		var ret Stream
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetBodyOk() (*Stream, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *HttpRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given Stream and assigns it to the Body field.
func (o *HttpRequest) SetBody(v Stream) {
	o.Body = &v
}

// GetBodyReader returns the BodyReader field value if set, zero value otherwise.
func (o *HttpRequest) GetBodyReader() map[string]interface{} {
	if o == nil || IsNil(o.BodyReader) {
		var ret map[string]interface{}
		return ret
	}
	return o.BodyReader
}

// GetBodyReaderOk returns a tuple with the BodyReader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetBodyReaderOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.BodyReader) {
		return map[string]interface{}{}, false
	}
	return o.BodyReader, true
}

// HasBodyReader returns a boolean if a field has been set.
func (o *HttpRequest) HasBodyReader() bool {
	if o != nil && !IsNil(o.BodyReader) {
		return true
	}

	return false
}

// SetBodyReader gets a reference to the given map[string]interface{} and assigns it to the BodyReader field.
func (o *HttpRequest) SetBodyReader(v map[string]interface{}) {
	o.BodyReader = v
}

// GetHasFormContentType returns the HasFormContentType field value if set, zero value otherwise.
func (o *HttpRequest) GetHasFormContentType() bool {
	if o == nil || IsNil(o.HasFormContentType) {
		var ret bool
		return ret
	}
	return *o.HasFormContentType
}

// GetHasFormContentTypeOk returns a tuple with the HasFormContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequest) GetHasFormContentTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.HasFormContentType) {
		return nil, false
	}
	return o.HasFormContentType, true
}

// HasHasFormContentType returns a boolean if a field has been set.
func (o *HttpRequest) HasHasFormContentType() bool {
	if o != nil && !IsNil(o.HasFormContentType) {
		return true
	}

	return false
}

// SetHasFormContentType gets a reference to the given bool and assigns it to the HasFormContentType field.
func (o *HttpRequest) SetHasFormContentType(v bool) {
	o.HasFormContentType = &v
}

// GetForm returns the Form field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetForm() []StringStringValuesKeyValuePair {
	if o == nil {
		var ret []StringStringValuesKeyValuePair
		return ret
	}
	return o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetFormOk() ([]StringStringValuesKeyValuePair, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *HttpRequest) HasForm() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given []StringStringValuesKeyValuePair and assigns it to the Form field.
func (o *HttpRequest) SetForm(v []StringStringValuesKeyValuePair) {
	o.Form = v
}

// GetRouteValues returns the RouteValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpRequest) GetRouteValues() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RouteValues
}

// GetRouteValuesOk returns a tuple with the RouteValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpRequest) GetRouteValuesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RouteValues) {
		return map[string]interface{}{}, false
	}
	return o.RouteValues, true
}

// HasRouteValues returns a boolean if a field has been set.
func (o *HttpRequest) HasRouteValues() bool {
	if o != nil && !IsNil(o.RouteValues) {
		return true
	}

	return false
}

// SetRouteValues gets a reference to the given map[string]interface{} and assigns it to the RouteValues field.
func (o *HttpRequest) SetRouteValues(v map[string]interface{}) {
	o.RouteValues = v
}

func (o HttpRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpContext) {
		toSerialize["httpContext"] = o.HttpContext
	}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if !IsNil(o.IsHttps) {
		toSerialize["isHttps"] = o.IsHttps
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.PathBase) {
		toSerialize["pathBase"] = o.PathBase
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.QueryString) {
		toSerialize["queryString"] = o.QueryString
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Cookies != nil {
		toSerialize["cookies"] = o.Cookies
	}
	if o.ContentLength.IsSet() {
		toSerialize["contentLength"] = o.ContentLength.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.BodyReader) {
		toSerialize["bodyReader"] = o.BodyReader
	}
	if !IsNil(o.HasFormContentType) {
		toSerialize["hasFormContentType"] = o.HasFormContentType
	}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	if o.RouteValues != nil {
		toSerialize["routeValues"] = o.RouteValues
	}
	return toSerialize, nil
}

type NullableHttpRequest struct {
	value *HttpRequest
	isSet bool
}

func (v NullableHttpRequest) Get() *HttpRequest {
	return v.value
}

func (v *NullableHttpRequest) Set(val *HttpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpRequest(val *HttpRequest) *NullableHttpRequest {
	return &NullableHttpRequest{value: val, isSet: true}
}

func (v NullableHttpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}