# WebSocketManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsWebSocketRequest** | Pointer to **bool** |  | [optional] [readonly] 
**WebSocketRequestedProtocols** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewWebSocketManager

`func NewWebSocketManager() *WebSocketManager`

NewWebSocketManager instantiates a new WebSocketManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSocketManagerWithDefaults

`func NewWebSocketManagerWithDefaults() *WebSocketManager`

NewWebSocketManagerWithDefaults instantiates a new WebSocketManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsWebSocketRequest

`func (o *WebSocketManager) GetIsWebSocketRequest() bool`

GetIsWebSocketRequest returns the IsWebSocketRequest field if non-nil, zero value otherwise.

### GetIsWebSocketRequestOk

`func (o *WebSocketManager) GetIsWebSocketRequestOk() (*bool, bool)`

GetIsWebSocketRequestOk returns a tuple with the IsWebSocketRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebSocketRequest

`func (o *WebSocketManager) SetIsWebSocketRequest(v bool)`

SetIsWebSocketRequest sets IsWebSocketRequest field to given value.

### HasIsWebSocketRequest

`func (o *WebSocketManager) HasIsWebSocketRequest() bool`

HasIsWebSocketRequest returns a boolean if a field has been set.

### GetWebSocketRequestedProtocols

`func (o *WebSocketManager) GetWebSocketRequestedProtocols() []string`

GetWebSocketRequestedProtocols returns the WebSocketRequestedProtocols field if non-nil, zero value otherwise.

### GetWebSocketRequestedProtocolsOk

`func (o *WebSocketManager) GetWebSocketRequestedProtocolsOk() (*[]string, bool)`

GetWebSocketRequestedProtocolsOk returns a tuple with the WebSocketRequestedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSocketRequestedProtocols

`func (o *WebSocketManager) SetWebSocketRequestedProtocols(v []string)`

SetWebSocketRequestedProtocols sets WebSocketRequestedProtocols field to given value.

### HasWebSocketRequestedProtocols

`func (o *WebSocketManager) HasWebSocketRequestedProtocols() bool`

HasWebSocketRequestedProtocols returns a boolean if a field has been set.

### SetWebSocketRequestedProtocolsNil

`func (o *WebSocketManager) SetWebSocketRequestedProtocolsNil(b bool)`

 SetWebSocketRequestedProtocolsNil sets the value for WebSocketRequestedProtocols to be an explicit nil

### UnsetWebSocketRequestedProtocols
`func (o *WebSocketManager) UnsetWebSocketRequestedProtocols()`

UnsetWebSocketRequestedProtocols ensures that no value is present for WebSocketRequestedProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


