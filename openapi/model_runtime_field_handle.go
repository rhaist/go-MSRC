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

// checks if the RuntimeFieldHandle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuntimeFieldHandle{}

// RuntimeFieldHandle struct for RuntimeFieldHandle
type RuntimeFieldHandle struct {
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewRuntimeFieldHandle instantiates a new RuntimeFieldHandle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeFieldHandle() *RuntimeFieldHandle {
	this := RuntimeFieldHandle{}
	return &this
}

// NewRuntimeFieldHandleWithDefaults instantiates a new RuntimeFieldHandle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeFieldHandleWithDefaults() *RuntimeFieldHandle {
	this := RuntimeFieldHandle{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RuntimeFieldHandle) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeFieldHandle) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RuntimeFieldHandle) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *RuntimeFieldHandle) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o RuntimeFieldHandle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuntimeFieldHandle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRuntimeFieldHandle struct {
	value *RuntimeFieldHandle
	isSet bool
}

func (v NullableRuntimeFieldHandle) Get() *RuntimeFieldHandle {
	return v.value
}

func (v *NullableRuntimeFieldHandle) Set(val *RuntimeFieldHandle) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeFieldHandle) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeFieldHandle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeFieldHandle(val *RuntimeFieldHandle) *NullableRuntimeFieldHandle {
	return &NullableRuntimeFieldHandle{value: val, isSet: true}
}

func (v NullableRuntimeFieldHandle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeFieldHandle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}