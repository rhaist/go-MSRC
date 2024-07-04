# HttpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpContext** | Pointer to [**HttpContext**](HttpContext.md) |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] [readonly] 
**Body** | Pointer to [**Stream**](Stream.md) |  | [optional] 
**BodyWriter** | Pointer to [**PipeWriter**](PipeWriter.md) |  | [optional] 
**ContentLength** | Pointer to **NullableInt64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**Cookies** | Pointer to **map[string]interface{}** |  | [optional] 
**HasStarted** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewHttpResponse

`func NewHttpResponse() *HttpResponse`

NewHttpResponse instantiates a new HttpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResponseWithDefaults

`func NewHttpResponseWithDefaults() *HttpResponse`

NewHttpResponseWithDefaults instantiates a new HttpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpContext

`func (o *HttpResponse) GetHttpContext() HttpContext`

GetHttpContext returns the HttpContext field if non-nil, zero value otherwise.

### GetHttpContextOk

`func (o *HttpResponse) GetHttpContextOk() (*HttpContext, bool)`

GetHttpContextOk returns a tuple with the HttpContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContext

`func (o *HttpResponse) SetHttpContext(v HttpContext)`

SetHttpContext sets HttpContext field to given value.

### HasHttpContext

`func (o *HttpResponse) HasHttpContext() bool`

HasHttpContext returns a boolean if a field has been set.

### GetStatusCode

`func (o *HttpResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HttpResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HttpResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HttpResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpResponse) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpResponse) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpResponse) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *HttpResponse) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *HttpResponse) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetBody

`func (o *HttpResponse) GetBody() Stream`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HttpResponse) GetBodyOk() (*Stream, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HttpResponse) SetBody(v Stream)`

SetBody sets Body field to given value.

### HasBody

`func (o *HttpResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyWriter

`func (o *HttpResponse) GetBodyWriter() PipeWriter`

GetBodyWriter returns the BodyWriter field if non-nil, zero value otherwise.

### GetBodyWriterOk

`func (o *HttpResponse) GetBodyWriterOk() (*PipeWriter, bool)`

GetBodyWriterOk returns a tuple with the BodyWriter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyWriter

`func (o *HttpResponse) SetBodyWriter(v PipeWriter)`

SetBodyWriter sets BodyWriter field to given value.

### HasBodyWriter

`func (o *HttpResponse) HasBodyWriter() bool`

HasBodyWriter returns a boolean if a field has been set.

### GetContentLength

`func (o *HttpResponse) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *HttpResponse) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *HttpResponse) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *HttpResponse) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *HttpResponse) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *HttpResponse) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *HttpResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HttpResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HttpResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HttpResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *HttpResponse) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *HttpResponse) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetCookies

`func (o *HttpResponse) GetCookies() map[string]interface{}`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *HttpResponse) GetCookiesOk() (*map[string]interface{}, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *HttpResponse) SetCookies(v map[string]interface{})`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *HttpResponse) HasCookies() bool`

HasCookies returns a boolean if a field has been set.

### GetHasStarted

`func (o *HttpResponse) GetHasStarted() bool`

GetHasStarted returns the HasStarted field if non-nil, zero value otherwise.

### GetHasStartedOk

`func (o *HttpResponse) GetHasStartedOk() (*bool, bool)`

GetHasStartedOk returns a tuple with the HasStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStarted

`func (o *HttpResponse) SetHasStarted(v bool)`

SetHasStarted sets HasStarted field to given value.

### HasHasStarted

`func (o *HttpResponse) HasHasStarted() bool`

HasHasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


