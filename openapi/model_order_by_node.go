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

// checks if the OrderByNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderByNode{}

// OrderByNode struct for OrderByNode
type OrderByNode struct {
	Direction *OrderByDirection `json:"direction,omitempty"`
}

// NewOrderByNode instantiates a new OrderByNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderByNode() *OrderByNode {
	this := OrderByNode{}
	return &this
}

// NewOrderByNodeWithDefaults instantiates a new OrderByNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderByNodeWithDefaults() *OrderByNode {
	this := OrderByNode{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *OrderByNode) GetDirection() OrderByDirection {
	if o == nil || IsNil(o.Direction) {
		var ret OrderByDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByNode) GetDirectionOk() (*OrderByDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *OrderByNode) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given OrderByDirection and assigns it to the Direction field.
func (o *OrderByNode) SetDirection(v OrderByDirection) {
	o.Direction = &v
}

func (o OrderByNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderByNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	return toSerialize, nil
}

type NullableOrderByNode struct {
	value *OrderByNode
	isSet bool
}

func (v NullableOrderByNode) Get() *OrderByNode {
	return v.value
}

func (v *NullableOrderByNode) Set(val *OrderByNode) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByNode) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByNode(val *OrderByNode) *NullableOrderByNode {
	return &NullableOrderByNode{value: val, isSet: true}
}

func (v NullableOrderByNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
