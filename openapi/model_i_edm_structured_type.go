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

// checks if the IEdmStructuredType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmStructuredType{}

// IEdmStructuredType struct for IEdmStructuredType
type IEdmStructuredType struct {
	IsAbstract         *bool               `json:"isAbstract,omitempty"`
	IsOpen             *bool               `json:"isOpen,omitempty"`
	BaseType           *IEdmStructuredType `json:"baseType,omitempty"`
	DeclaredProperties []IEdmProperty      `json:"declaredProperties,omitempty"`
	TypeKind           *EdmTypeKind        `json:"typeKind,omitempty"`
}

// NewIEdmStructuredType instantiates a new IEdmStructuredType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmStructuredType() *IEdmStructuredType {
	this := IEdmStructuredType{}
	return &this
}

// NewIEdmStructuredTypeWithDefaults instantiates a new IEdmStructuredType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmStructuredTypeWithDefaults() *IEdmStructuredType {
	this := IEdmStructuredType{}
	return &this
}

// GetIsAbstract returns the IsAbstract field value if set, zero value otherwise.
func (o *IEdmStructuredType) GetIsAbstract() bool {
	if o == nil || IsNil(o.IsAbstract) {
		var ret bool
		return ret
	}
	return *o.IsAbstract
}

// GetIsAbstractOk returns a tuple with the IsAbstract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmStructuredType) GetIsAbstractOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAbstract) {
		return nil, false
	}
	return o.IsAbstract, true
}

// HasIsAbstract returns a boolean if a field has been set.
func (o *IEdmStructuredType) HasIsAbstract() bool {
	if o != nil && !IsNil(o.IsAbstract) {
		return true
	}

	return false
}

// SetIsAbstract gets a reference to the given bool and assigns it to the IsAbstract field.
func (o *IEdmStructuredType) SetIsAbstract(v bool) {
	o.IsAbstract = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *IEdmStructuredType) GetIsOpen() bool {
	if o == nil || IsNil(o.IsOpen) {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmStructuredType) GetIsOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpen) {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *IEdmStructuredType) HasIsOpen() bool {
	if o != nil && !IsNil(o.IsOpen) {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *IEdmStructuredType) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *IEdmStructuredType) GetBaseType() IEdmStructuredType {
	if o == nil || IsNil(o.BaseType) {
		var ret IEdmStructuredType
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmStructuredType) GetBaseTypeOk() (*IEdmStructuredType, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *IEdmStructuredType) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given IEdmStructuredType and assigns it to the BaseType field.
func (o *IEdmStructuredType) SetBaseType(v IEdmStructuredType) {
	o.BaseType = &v
}

// GetDeclaredProperties returns the DeclaredProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmStructuredType) GetDeclaredProperties() []IEdmProperty {
	if o == nil {
		var ret []IEdmProperty
		return ret
	}
	return o.DeclaredProperties
}

// GetDeclaredPropertiesOk returns a tuple with the DeclaredProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmStructuredType) GetDeclaredPropertiesOk() ([]IEdmProperty, bool) {
	if o == nil || IsNil(o.DeclaredProperties) {
		return nil, false
	}
	return o.DeclaredProperties, true
}

// HasDeclaredProperties returns a boolean if a field has been set.
func (o *IEdmStructuredType) HasDeclaredProperties() bool {
	if o != nil && !IsNil(o.DeclaredProperties) {
		return true
	}

	return false
}

// SetDeclaredProperties gets a reference to the given []IEdmProperty and assigns it to the DeclaredProperties field.
func (o *IEdmStructuredType) SetDeclaredProperties(v []IEdmProperty) {
	o.DeclaredProperties = v
}

// GetTypeKind returns the TypeKind field value if set, zero value otherwise.
func (o *IEdmStructuredType) GetTypeKind() EdmTypeKind {
	if o == nil || IsNil(o.TypeKind) {
		var ret EdmTypeKind
		return ret
	}
	return *o.TypeKind
}

// GetTypeKindOk returns a tuple with the TypeKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmStructuredType) GetTypeKindOk() (*EdmTypeKind, bool) {
	if o == nil || IsNil(o.TypeKind) {
		return nil, false
	}
	return o.TypeKind, true
}

// HasTypeKind returns a boolean if a field has been set.
func (o *IEdmStructuredType) HasTypeKind() bool {
	if o != nil && !IsNil(o.TypeKind) {
		return true
	}

	return false
}

// SetTypeKind gets a reference to the given EdmTypeKind and assigns it to the TypeKind field.
func (o *IEdmStructuredType) SetTypeKind(v EdmTypeKind) {
	o.TypeKind = &v
}

func (o IEdmStructuredType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmStructuredType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAbstract) {
		toSerialize["isAbstract"] = o.IsAbstract
	}
	if !IsNil(o.IsOpen) {
		toSerialize["isOpen"] = o.IsOpen
	}
	if !IsNil(o.BaseType) {
		toSerialize["baseType"] = o.BaseType
	}
	if o.DeclaredProperties != nil {
		toSerialize["declaredProperties"] = o.DeclaredProperties
	}
	if !IsNil(o.TypeKind) {
		toSerialize["typeKind"] = o.TypeKind
	}
	return toSerialize, nil
}

type NullableIEdmStructuredType struct {
	value *IEdmStructuredType
	isSet bool
}

func (v NullableIEdmStructuredType) Get() *IEdmStructuredType {
	return v.value
}

func (v *NullableIEdmStructuredType) Set(val *IEdmStructuredType) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmStructuredType) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmStructuredType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmStructuredType(val *IEdmStructuredType) *NullableIEdmStructuredType {
	return &NullableIEdmStructuredType{value: val, isSet: true}
}

func (v NullableIEdmStructuredType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmStructuredType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
