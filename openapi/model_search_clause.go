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

// checks if the SearchClause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchClause{}

// SearchClause struct for SearchClause
type SearchClause struct {
	Expression *SingleValueNode `json:"expression,omitempty"`
}

// NewSearchClause instantiates a new SearchClause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchClause() *SearchClause {
	this := SearchClause{}
	return &this
}

// NewSearchClauseWithDefaults instantiates a new SearchClause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchClauseWithDefaults() *SearchClause {
	this := SearchClause{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SearchClause) GetExpression() SingleValueNode {
	if o == nil || IsNil(o.Expression) {
		var ret SingleValueNode
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchClause) GetExpressionOk() (*SingleValueNode, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SearchClause) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given SingleValueNode and assigns it to the Expression field.
func (o *SearchClause) SetExpression(v SingleValueNode) {
	o.Expression = &v
}

func (o SearchClause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchClause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

type NullableSearchClause struct {
	value *SearchClause
	isSet bool
}

func (v NullableSearchClause) Get() *SearchClause {
	return v.value
}

func (v *NullableSearchClause) Set(val *SearchClause) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchClause) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchClause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchClause(val *SearchClause) *NullableSearchClause {
	return &NullableSearchClause{value: val, isSet: true}
}

func (v NullableSearchClause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchClause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}