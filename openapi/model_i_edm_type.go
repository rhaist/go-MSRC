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

// checks if the IEdmType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmType{}

// IEdmType struct for IEdmType
type IEdmType struct {
	TypeKind *EdmTypeKind `json:"typeKind,omitempty"`
}

// NewIEdmType instantiates a new IEdmType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmType() *IEdmType {
	this := IEdmType{}
	return &this
}

// NewIEdmTypeWithDefaults instantiates a new IEdmType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmTypeWithDefaults() *IEdmType {
	this := IEdmType{}
	return &this
}

// GetTypeKind returns the TypeKind field value if set, zero value otherwise.
func (o *IEdmType) GetTypeKind() EdmTypeKind {
	if o == nil || IsNil(o.TypeKind) {
		var ret EdmTypeKind
		return ret
	}
	return *o.TypeKind
}

// GetTypeKindOk returns a tuple with the TypeKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmType) GetTypeKindOk() (*EdmTypeKind, bool) {
	if o == nil || IsNil(o.TypeKind) {
		return nil, false
	}
	return o.TypeKind, true
}

// HasTypeKind returns a boolean if a field has been set.
func (o *IEdmType) HasTypeKind() bool {
	if o != nil && !IsNil(o.TypeKind) {
		return true
	}

	return false
}

// SetTypeKind gets a reference to the given EdmTypeKind and assigns it to the TypeKind field.
func (o *IEdmType) SetTypeKind(v EdmTypeKind) {
	o.TypeKind = &v
}

func (o IEdmType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TypeKind) {
		toSerialize["typeKind"] = o.TypeKind
	}
	return toSerialize, nil
}

type NullableIEdmType struct {
	value *IEdmType
	isSet bool
}

func (v NullableIEdmType) Get() *IEdmType {
	return v.value
}

func (v *NullableIEdmType) Set(val *IEdmType) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmType) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmType(val *IEdmType) *NullableIEdmType {
	return &NullableIEdmType{value: val, isSet: true}
}

func (v NullableIEdmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}