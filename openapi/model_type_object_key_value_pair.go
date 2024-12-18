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

// checks if the TypeObjectKeyValuePair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypeObjectKeyValuePair{}

// TypeObjectKeyValuePair struct for TypeObjectKeyValuePair
type TypeObjectKeyValuePair struct {
	Key   *Type       `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewTypeObjectKeyValuePair instantiates a new TypeObjectKeyValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeObjectKeyValuePair() *TypeObjectKeyValuePair {
	this := TypeObjectKeyValuePair{}
	return &this
}

// NewTypeObjectKeyValuePairWithDefaults instantiates a new TypeObjectKeyValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeObjectKeyValuePairWithDefaults() *TypeObjectKeyValuePair {
	this := TypeObjectKeyValuePair{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TypeObjectKeyValuePair) GetKey() Type {
	if o == nil || IsNil(o.Key) {
		var ret Type
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeObjectKeyValuePair) GetKeyOk() (*Type, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TypeObjectKeyValuePair) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given Type and assigns it to the Key field.
func (o *TypeObjectKeyValuePair) SetKey(v Type) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TypeObjectKeyValuePair) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TypeObjectKeyValuePair) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TypeObjectKeyValuePair) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *TypeObjectKeyValuePair) SetValue(v interface{}) {
	o.Value = v
}

func (o TypeObjectKeyValuePair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypeObjectKeyValuePair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTypeObjectKeyValuePair struct {
	value *TypeObjectKeyValuePair
	isSet bool
}

func (v NullableTypeObjectKeyValuePair) Get() *TypeObjectKeyValuePair {
	return v.value
}

func (v *NullableTypeObjectKeyValuePair) Set(val *TypeObjectKeyValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeObjectKeyValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeObjectKeyValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeObjectKeyValuePair(val *TypeObjectKeyValuePair) *NullableTypeObjectKeyValuePair {
	return &NullableTypeObjectKeyValuePair{value: val, isSet: true}
}

func (v NullableTypeObjectKeyValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeObjectKeyValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
