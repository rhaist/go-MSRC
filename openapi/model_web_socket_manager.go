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

// checks if the WebSocketManager type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebSocketManager{}

// WebSocketManager struct for WebSocketManager
type WebSocketManager struct {
	IsWebSocketRequest          *bool    `json:"isWebSocketRequest,omitempty"`
	WebSocketRequestedProtocols []string `json:"webSocketRequestedProtocols,omitempty"`
}

// NewWebSocketManager instantiates a new WebSocketManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebSocketManager() *WebSocketManager {
	this := WebSocketManager{}
	return &this
}

// NewWebSocketManagerWithDefaults instantiates a new WebSocketManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebSocketManagerWithDefaults() *WebSocketManager {
	this := WebSocketManager{}
	return &this
}

// GetIsWebSocketRequest returns the IsWebSocketRequest field value if set, zero value otherwise.
func (o *WebSocketManager) GetIsWebSocketRequest() bool {
	if o == nil || IsNil(o.IsWebSocketRequest) {
		var ret bool
		return ret
	}
	return *o.IsWebSocketRequest
}

// GetIsWebSocketRequestOk returns a tuple with the IsWebSocketRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebSocketManager) GetIsWebSocketRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWebSocketRequest) {
		return nil, false
	}
	return o.IsWebSocketRequest, true
}

// HasIsWebSocketRequest returns a boolean if a field has been set.
func (o *WebSocketManager) HasIsWebSocketRequest() bool {
	if o != nil && !IsNil(o.IsWebSocketRequest) {
		return true
	}

	return false
}

// SetIsWebSocketRequest gets a reference to the given bool and assigns it to the IsWebSocketRequest field.
func (o *WebSocketManager) SetIsWebSocketRequest(v bool) {
	o.IsWebSocketRequest = &v
}

// GetWebSocketRequestedProtocols returns the WebSocketRequestedProtocols field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketManager) GetWebSocketRequestedProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WebSocketRequestedProtocols
}

// GetWebSocketRequestedProtocolsOk returns a tuple with the WebSocketRequestedProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketManager) GetWebSocketRequestedProtocolsOk() ([]string, bool) {
	if o == nil || IsNil(o.WebSocketRequestedProtocols) {
		return nil, false
	}
	return o.WebSocketRequestedProtocols, true
}

// HasWebSocketRequestedProtocols returns a boolean if a field has been set.
func (o *WebSocketManager) HasWebSocketRequestedProtocols() bool {
	if o != nil && !IsNil(o.WebSocketRequestedProtocols) {
		return true
	}

	return false
}

// SetWebSocketRequestedProtocols gets a reference to the given []string and assigns it to the WebSocketRequestedProtocols field.
func (o *WebSocketManager) SetWebSocketRequestedProtocols(v []string) {
	o.WebSocketRequestedProtocols = v
}

func (o WebSocketManager) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebSocketManager) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsWebSocketRequest) {
		toSerialize["isWebSocketRequest"] = o.IsWebSocketRequest
	}
	if o.WebSocketRequestedProtocols != nil {
		toSerialize["webSocketRequestedProtocols"] = o.WebSocketRequestedProtocols
	}
	return toSerialize, nil
}

type NullableWebSocketManager struct {
	value *WebSocketManager
	isSet bool
}

func (v NullableWebSocketManager) Get() *WebSocketManager {
	return v.value
}

func (v *NullableWebSocketManager) Set(val *WebSocketManager) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSocketManager) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSocketManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSocketManager(val *WebSocketManager) *NullableWebSocketManager {
	return &NullableWebSocketManager{value: val, isSet: true}
}

func (v NullableWebSocketManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSocketManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
