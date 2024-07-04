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

// checks if the TransformationNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationNode{}

// TransformationNode struct for TransformationNode
type TransformationNode struct {
	Kind *TransformationNodeKind `json:"kind,omitempty"`
}

// NewTransformationNode instantiates a new TransformationNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationNode() *TransformationNode {
	this := TransformationNode{}
	return &this
}

// NewTransformationNodeWithDefaults instantiates a new TransformationNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationNodeWithDefaults() *TransformationNode {
	this := TransformationNode{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TransformationNode) GetKind() TransformationNodeKind {
	if o == nil || IsNil(o.Kind) {
		var ret TransformationNodeKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationNode) GetKindOk() (*TransformationNodeKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TransformationNode) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given TransformationNodeKind and assigns it to the Kind field.
func (o *TransformationNode) SetKind(v TransformationNodeKind) {
	o.Kind = &v
}

func (o TransformationNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableTransformationNode struct {
	value *TransformationNode
	isSet bool
}

func (v NullableTransformationNode) Get() *TransformationNode {
	return v.value
}

func (v *NullableTransformationNode) Set(val *TransformationNode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationNode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationNode(val *TransformationNode) *NullableTransformationNode {
	return &NullableTransformationNode{value: val, isSet: true}
}

func (v NullableTransformationNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
