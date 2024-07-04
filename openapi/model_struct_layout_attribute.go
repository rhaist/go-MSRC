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

// checks if the StructLayoutAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructLayoutAttribute{}

// StructLayoutAttribute struct for StructLayoutAttribute
type StructLayoutAttribute struct {
	TypeId interface{} `json:"typeId,omitempty"`
	Value  *LayoutKind `json:"value,omitempty"`
}

// NewStructLayoutAttribute instantiates a new StructLayoutAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructLayoutAttribute() *StructLayoutAttribute {
	this := StructLayoutAttribute{}
	return &this
}

// NewStructLayoutAttributeWithDefaults instantiates a new StructLayoutAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructLayoutAttributeWithDefaults() *StructLayoutAttribute {
	this := StructLayoutAttribute{}
	return &this
}

// GetTypeId returns the TypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StructLayoutAttribute) GetTypeId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StructLayoutAttribute) GetTypeIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return &o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *StructLayoutAttribute) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given interface{} and assigns it to the TypeId field.
func (o *StructLayoutAttribute) SetTypeId(v interface{}) {
	o.TypeId = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StructLayoutAttribute) GetValue() LayoutKind {
	if o == nil || IsNil(o.Value) {
		var ret LayoutKind
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructLayoutAttribute) GetValueOk() (*LayoutKind, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StructLayoutAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given LayoutKind and assigns it to the Value field.
func (o *StructLayoutAttribute) SetValue(v LayoutKind) {
	o.Value = &v
}

func (o StructLayoutAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructLayoutAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeId != nil {
		toSerialize["typeId"] = o.TypeId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableStructLayoutAttribute struct {
	value *StructLayoutAttribute
	isSet bool
}

func (v NullableStructLayoutAttribute) Get() *StructLayoutAttribute {
	return v.value
}

func (v *NullableStructLayoutAttribute) Set(val *StructLayoutAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableStructLayoutAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableStructLayoutAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructLayoutAttribute(val *StructLayoutAttribute) *NullableStructLayoutAttribute {
	return &NullableStructLayoutAttribute{value: val, isSet: true}
}

func (v NullableStructLayoutAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructLayoutAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}