# HttpContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]TypeObjectKeyValuePair**](TypeObjectKeyValuePair.md) |  | [optional] [readonly] 
**Request** | Pointer to [**HttpRequest**](HttpRequest.md) |  | [optional] 
**Response** | Pointer to [**HttpResponse**](HttpResponse.md) |  | [optional] 
**Connection** | Pointer to [**ConnectionInfo**](ConnectionInfo.md) |  | [optional] 
**WebSockets** | Pointer to [**WebSocketManager**](WebSocketManager.md) |  | [optional] 
**User** | Pointer to [**ClaimsPrincipal**](ClaimsPrincipal.md) |  | [optional] 
**Items** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestServices** | Pointer to **map[string]interface{}** |  | [optional] 
**RequestAborted** | Pointer to [**CancellationToken**](CancellationToken.md) |  | [optional] 
**TraceIdentifier** | Pointer to **NullableString** |  | [optional] 
**Session** | Pointer to [**ISession**](ISession.md) |  | [optional] 

## Methods

### NewHttpContext

`func NewHttpContext() *HttpContext`

NewHttpContext instantiates a new HttpContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpContextWithDefaults

`func NewHttpContextWithDefaults() *HttpContext`

NewHttpContextWithDefaults instantiates a new HttpContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *HttpContext) GetFeatures() []TypeObjectKeyValuePair`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *HttpContext) GetFeaturesOk() (*[]TypeObjectKeyValuePair, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *HttpContext) SetFeatures(v []TypeObjectKeyValuePair)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *HttpContext) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *HttpContext) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *HttpContext) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetRequest

`func (o *HttpContext) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *HttpContext) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *HttpContext) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *HttpContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *HttpContext) GetResponse() HttpResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *HttpContext) GetResponseOk() (*HttpResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *HttpContext) SetResponse(v HttpResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *HttpContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetConnection

`func (o *HttpContext) GetConnection() ConnectionInfo`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *HttpContext) GetConnectionOk() (*ConnectionInfo, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *HttpContext) SetConnection(v ConnectionInfo)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *HttpContext) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetWebSockets

`func (o *HttpContext) GetWebSockets() WebSocketManager`

GetWebSockets returns the WebSockets field if non-nil, zero value otherwise.

### GetWebSocketsOk

`func (o *HttpContext) GetWebSocketsOk() (*WebSocketManager, bool)`

GetWebSocketsOk returns a tuple with the WebSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSockets

`func (o *HttpContext) SetWebSockets(v WebSocketManager)`

SetWebSockets sets WebSockets field to given value.

### HasWebSockets

`func (o *HttpContext) HasWebSockets() bool`

HasWebSockets returns a boolean if a field has been set.

### GetUser

`func (o *HttpContext) GetUser() ClaimsPrincipal`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HttpContext) GetUserOk() (*ClaimsPrincipal, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HttpContext) SetUser(v ClaimsPrincipal)`

SetUser sets User field to given value.

### HasUser

`func (o *HttpContext) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetItems

`func (o *HttpContext) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HttpContext) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HttpContext) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *HttpContext) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *HttpContext) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *HttpContext) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetRequestServices

`func (o *HttpContext) GetRequestServices() map[string]interface{}`

GetRequestServices returns the RequestServices field if non-nil, zero value otherwise.

### GetRequestServicesOk

`func (o *HttpContext) GetRequestServicesOk() (*map[string]interface{}, bool)`

GetRequestServicesOk returns a tuple with the RequestServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestServices

`func (o *HttpContext) SetRequestServices(v map[string]interface{})`

SetRequestServices sets RequestServices field to given value.

### HasRequestServices

`func (o *HttpContext) HasRequestServices() bool`

HasRequestServices returns a boolean if a field has been set.

### GetRequestAborted

`func (o *HttpContext) GetRequestAborted() CancellationToken`

GetRequestAborted returns the RequestAborted field if non-nil, zero value otherwise.

### GetRequestAbortedOk

`func (o *HttpContext) GetRequestAbortedOk() (*CancellationToken, bool)`

GetRequestAbortedOk returns a tuple with the RequestAborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAborted

`func (o *HttpContext) SetRequestAborted(v CancellationToken)`

SetRequestAborted sets RequestAborted field to given value.

### HasRequestAborted

`func (o *HttpContext) HasRequestAborted() bool`

HasRequestAborted returns a boolean if a field has been set.

### GetTraceIdentifier

`func (o *HttpContext) GetTraceIdentifier() string`

GetTraceIdentifier returns the TraceIdentifier field if non-nil, zero value otherwise.

### GetTraceIdentifierOk

`func (o *HttpContext) GetTraceIdentifierOk() (*string, bool)`

GetTraceIdentifierOk returns a tuple with the TraceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceIdentifier

`func (o *HttpContext) SetTraceIdentifier(v string)`

SetTraceIdentifier sets TraceIdentifier field to given value.

### HasTraceIdentifier

`func (o *HttpContext) HasTraceIdentifier() bool`

HasTraceIdentifier returns a boolean if a field has been set.

### SetTraceIdentifierNil

`func (o *HttpContext) SetTraceIdentifierNil(b bool)`

 SetTraceIdentifierNil sets the value for TraceIdentifier to be an explicit nil

### UnsetTraceIdentifier
`func (o *HttpContext) UnsetTraceIdentifier()`

UnsetTraceIdentifier ensures that no value is present for TraceIdentifier, not even an explicit nil
### GetSession

`func (o *HttpContext) GetSession() ISession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *HttpContext) GetSessionOk() (*ISession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *HttpContext) SetSession(v ISession)`

SetSession sets Session field to given value.

### HasSession

`func (o *HttpContext) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


