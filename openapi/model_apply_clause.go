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

// checks if the ApplyClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyClause{}

// ApplyClause struct for ApplyClause
type ApplyClause struct {
	Transformations []TransformationNode `json:"transformations,omitempty"`
}

// NewApplyClause instantiates a new ApplyClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyClause() *ApplyClause {
	this := ApplyClause{}
	return &this
}

// NewApplyClauseWithDefaults instantiates a new ApplyClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyClauseWithDefaults() *ApplyClause {
	this := ApplyClause{}
	return &this
}

// GetTransformations returns the Transformations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplyClause) GetTransformations() []TransformationNode {
	if o == nil {
		var ret []TransformationNode
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplyClause) GetTransformationsOk() ([]TransformationNode, bool) {
	if o == nil || IsNil(o.Transformations) {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *ApplyClause) HasTransformations() bool {
	if o != nil && !IsNil(o.Transformations) {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []TransformationNode and assigns it to the Transformations field.
func (o *ApplyClause) SetTransformations(v []TransformationNode) {
	o.Transformations = v
}

func (o ApplyClause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Transformations != nil {
		toSerialize["transformations"] = o.Transformations
	}
	return toSerialize, nil
}

type NullableApplyClause struct {
	value *ApplyClause
	isSet bool
}

func (v NullableApplyClause) Get() *ApplyClause {
	return v.value
}

func (v *NullableApplyClause) Set(val *ApplyClause) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyClause) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyClause(val *ApplyClause) *NullableApplyClause {
	return &NullableApplyClause{value: val, isSet: true}
}

func (v NullableApplyClause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}