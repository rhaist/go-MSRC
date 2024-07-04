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

// checks if the OrderByQueryOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderByQueryOption{}

// OrderByQueryOption struct for OrderByQueryOption
type OrderByQueryOption struct {
	Context       *ODataQueryContext     `json:"context,omitempty"`
	OrderByNodes  []OrderByNode          `json:"orderByNodes,omitempty"`
	RawValue      NullableString         `json:"rawValue,omitempty"`
	Validator     map[string]interface{} `json:"validator,omitempty"`
	Compute       *ComputeQueryOption    `json:"compute,omitempty"`
	OrderByClause *OrderByClause         `json:"orderByClause,omitempty"`
}

// NewOrderByQueryOption instantiates a new OrderByQueryOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderByQueryOption() *OrderByQueryOption {
	this := OrderByQueryOption{}
	return &this
}

// NewOrderByQueryOptionWithDefaults instantiates a new OrderByQueryOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderByQueryOptionWithDefaults() *OrderByQueryOption {
	this := OrderByQueryOption{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *OrderByQueryOption) GetContext() ODataQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret ODataQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByQueryOption) GetContextOk() (*ODataQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given ODataQueryContext and assigns it to the Context field.
func (o *OrderByQueryOption) SetContext(v ODataQueryContext) {
	o.Context = &v
}

// GetOrderByNodes returns the OrderByNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderByQueryOption) GetOrderByNodes() []OrderByNode {
	if o == nil {
		var ret []OrderByNode
		return ret
	}
	return o.OrderByNodes
}

// GetOrderByNodesOk returns a tuple with the OrderByNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderByQueryOption) GetOrderByNodesOk() ([]OrderByNode, bool) {
	if o == nil || IsNil(o.OrderByNodes) {
		return nil, false
	}
	return o.OrderByNodes, true
}

// HasOrderByNodes returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasOrderByNodes() bool {
	if o != nil && !IsNil(o.OrderByNodes) {
		return true
	}

	return false
}

// SetOrderByNodes gets a reference to the given []OrderByNode and assigns it to the OrderByNodes field.
func (o *OrderByQueryOption) SetOrderByNodes(v []OrderByNode) {
	o.OrderByNodes = v
}

// GetRawValue returns the RawValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderByQueryOption) GetRawValue() string {
	if o == nil || IsNil(o.RawValue.Get()) {
		var ret string
		return ret
	}
	return *o.RawValue.Get()
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderByQueryOption) GetRawValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawValue.Get(), o.RawValue.IsSet()
}

// HasRawValue returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasRawValue() bool {
	if o != nil && o.RawValue.IsSet() {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given NullableString and assigns it to the RawValue field.
func (o *OrderByQueryOption) SetRawValue(v string) {
	o.RawValue.Set(&v)
}

// SetRawValueNil sets the value for RawValue to be an explicit nil
func (o *OrderByQueryOption) SetRawValueNil() {
	o.RawValue.Set(nil)
}

// UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
func (o *OrderByQueryOption) UnsetRawValue() {
	o.RawValue.Unset()
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *OrderByQueryOption) GetValidator() map[string]interface{} {
	if o == nil || IsNil(o.Validator) {
		var ret map[string]interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByQueryOption) GetValidatorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Validator) {
		return map[string]interface{}{}, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasValidator() bool {
	if o != nil && !IsNil(o.Validator) {
		return true
	}

	return false
}

// SetValidator gets a reference to the given map[string]interface{} and assigns it to the Validator field.
func (o *OrderByQueryOption) SetValidator(v map[string]interface{}) {
	o.Validator = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *OrderByQueryOption) GetCompute() ComputeQueryOption {
	if o == nil || IsNil(o.Compute) {
		var ret ComputeQueryOption
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByQueryOption) GetComputeOk() (*ComputeQueryOption, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ComputeQueryOption and assigns it to the Compute field.
func (o *OrderByQueryOption) SetCompute(v ComputeQueryOption) {
	o.Compute = &v
}

// GetOrderByClause returns the OrderByClause field value if set, zero value otherwise.
func (o *OrderByQueryOption) GetOrderByClause() OrderByClause {
	if o == nil || IsNil(o.OrderByClause) {
		var ret OrderByClause
		return ret
	}
	return *o.OrderByClause
}

// GetOrderByClauseOk returns a tuple with the OrderByClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByQueryOption) GetOrderByClauseOk() (*OrderByClause, bool) {
	if o == nil || IsNil(o.OrderByClause) {
		return nil, false
	}
	return o.OrderByClause, true
}

// HasOrderByClause returns a boolean if a field has been set.
func (o *OrderByQueryOption) HasOrderByClause() bool {
	if o != nil && !IsNil(o.OrderByClause) {
		return true
	}

	return false
}

// SetOrderByClause gets a reference to the given OrderByClause and assigns it to the OrderByClause field.
func (o *OrderByQueryOption) SetOrderByClause(v OrderByClause) {
	o.OrderByClause = &v
}

func (o OrderByQueryOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderByQueryOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if o.OrderByNodes != nil {
		toSerialize["orderByNodes"] = o.OrderByNodes
	}
	if o.RawValue.IsSet() {
		toSerialize["rawValue"] = o.RawValue.Get()
	}
	if !IsNil(o.Validator) {
		toSerialize["validator"] = o.Validator
	}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.OrderByClause) {
		toSerialize["orderByClause"] = o.OrderByClause
	}
	return toSerialize, nil
}

type NullableOrderByQueryOption struct {
	value *OrderByQueryOption
	isSet bool
}

func (v NullableOrderByQueryOption) Get() *OrderByQueryOption {
	return v.value
}

func (v *NullableOrderByQueryOption) Set(val *OrderByQueryOption) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByQueryOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByQueryOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByQueryOption(val *OrderByQueryOption) *NullableOrderByQueryOption {
	return &NullableOrderByQueryOption{value: val, isSet: true}
}

func (v NullableOrderByQueryOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByQueryOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
