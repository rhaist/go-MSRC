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

// checks if the SelectExpandQueryOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectExpandQueryOption{}

// SelectExpandQueryOption struct for SelectExpandQueryOption
type SelectExpandQueryOption struct {
	Context                        *ODataQueryContext     `json:"context,omitempty"`
	RawSelect                      NullableString         `json:"rawSelect,omitempty"`
	RawExpand                      NullableString         `json:"rawExpand,omitempty"`
	Compute                        *ComputeQueryOption    `json:"compute,omitempty"`
	Validator                      map[string]interface{} `json:"validator,omitempty"`
	SelectExpandClause             *SelectExpandClause    `json:"selectExpandClause,omitempty"`
	LevelsMaxLiteralExpansionDepth *int32                 `json:"levelsMaxLiteralExpansionDepth,omitempty"`
}

// NewSelectExpandQueryOption instantiates a new SelectExpandQueryOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectExpandQueryOption() *SelectExpandQueryOption {
	this := SelectExpandQueryOption{}
	return &this
}

// NewSelectExpandQueryOptionWithDefaults instantiates a new SelectExpandQueryOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectExpandQueryOptionWithDefaults() *SelectExpandQueryOption {
	this := SelectExpandQueryOption{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *SelectExpandQueryOption) GetContext() ODataQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret ODataQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectExpandQueryOption) GetContextOk() (*ODataQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given ODataQueryContext and assigns it to the Context field.
func (o *SelectExpandQueryOption) SetContext(v ODataQueryContext) {
	o.Context = &v
}

// GetRawSelect returns the RawSelect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectExpandQueryOption) GetRawSelect() string {
	if o == nil || IsNil(o.RawSelect.Get()) {
		var ret string
		return ret
	}
	return *o.RawSelect.Get()
}

// GetRawSelectOk returns a tuple with the RawSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectExpandQueryOption) GetRawSelectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawSelect.Get(), o.RawSelect.IsSet()
}

// HasRawSelect returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasRawSelect() bool {
	if o != nil && o.RawSelect.IsSet() {
		return true
	}

	return false
}

// SetRawSelect gets a reference to the given NullableString and assigns it to the RawSelect field.
func (o *SelectExpandQueryOption) SetRawSelect(v string) {
	o.RawSelect.Set(&v)
}

// SetRawSelectNil sets the value for RawSelect to be an explicit nil
func (o *SelectExpandQueryOption) SetRawSelectNil() {
	o.RawSelect.Set(nil)
}

// UnsetRawSelect ensures that no value is present for RawSelect, not even an explicit nil
func (o *SelectExpandQueryOption) UnsetRawSelect() {
	o.RawSelect.Unset()
}

// GetRawExpand returns the RawExpand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectExpandQueryOption) GetRawExpand() string {
	if o == nil || IsNil(o.RawExpand.Get()) {
		var ret string
		return ret
	}
	return *o.RawExpand.Get()
}

// GetRawExpandOk returns a tuple with the RawExpand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectExpandQueryOption) GetRawExpandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawExpand.Get(), o.RawExpand.IsSet()
}

// HasRawExpand returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasRawExpand() bool {
	if o != nil && o.RawExpand.IsSet() {
		return true
	}

	return false
}

// SetRawExpand gets a reference to the given NullableString and assigns it to the RawExpand field.
func (o *SelectExpandQueryOption) SetRawExpand(v string) {
	o.RawExpand.Set(&v)
}

// SetRawExpandNil sets the value for RawExpand to be an explicit nil
func (o *SelectExpandQueryOption) SetRawExpandNil() {
	o.RawExpand.Set(nil)
}

// UnsetRawExpand ensures that no value is present for RawExpand, not even an explicit nil
func (o *SelectExpandQueryOption) UnsetRawExpand() {
	o.RawExpand.Unset()
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SelectExpandQueryOption) GetCompute() ComputeQueryOption {
	if o == nil || IsNil(o.Compute) {
		var ret ComputeQueryOption
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectExpandQueryOption) GetComputeOk() (*ComputeQueryOption, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ComputeQueryOption and assigns it to the Compute field.
func (o *SelectExpandQueryOption) SetCompute(v ComputeQueryOption) {
	o.Compute = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *SelectExpandQueryOption) GetValidator() map[string]interface{} {
	if o == nil || IsNil(o.Validator) {
		var ret map[string]interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectExpandQueryOption) GetValidatorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Validator) {
		return map[string]interface{}{}, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasValidator() bool {
	if o != nil && !IsNil(o.Validator) {
		return true
	}

	return false
}

// SetValidator gets a reference to the given map[string]interface{} and assigns it to the Validator field.
func (o *SelectExpandQueryOption) SetValidator(v map[string]interface{}) {
	o.Validator = v
}

// GetSelectExpandClause returns the SelectExpandClause field value if set, zero value otherwise.
func (o *SelectExpandQueryOption) GetSelectExpandClause() SelectExpandClause {
	if o == nil || IsNil(o.SelectExpandClause) {
		var ret SelectExpandClause
		return ret
	}
	return *o.SelectExpandClause
}

// GetSelectExpandClauseOk returns a tuple with the SelectExpandClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectExpandQueryOption) GetSelectExpandClauseOk() (*SelectExpandClause, bool) {
	if o == nil || IsNil(o.SelectExpandClause) {
		return nil, false
	}
	return o.SelectExpandClause, true
}

// HasSelectExpandClause returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasSelectExpandClause() bool {
	if o != nil && !IsNil(o.SelectExpandClause) {
		return true
	}

	return false
}

// SetSelectExpandClause gets a reference to the given SelectExpandClause and assigns it to the SelectExpandClause field.
func (o *SelectExpandQueryOption) SetSelectExpandClause(v SelectExpandClause) {
	o.SelectExpandClause = &v
}

// GetLevelsMaxLiteralExpansionDepth returns the LevelsMaxLiteralExpansionDepth field value if set, zero value otherwise.
func (o *SelectExpandQueryOption) GetLevelsMaxLiteralExpansionDepth() int32 {
	if o == nil || IsNil(o.LevelsMaxLiteralExpansionDepth) {
		var ret int32
		return ret
	}
	return *o.LevelsMaxLiteralExpansionDepth
}

// GetLevelsMaxLiteralExpansionDepthOk returns a tuple with the LevelsMaxLiteralExpansionDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectExpandQueryOption) GetLevelsMaxLiteralExpansionDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.LevelsMaxLiteralExpansionDepth) {
		return nil, false
	}
	return o.LevelsMaxLiteralExpansionDepth, true
}

// HasLevelsMaxLiteralExpansionDepth returns a boolean if a field has been set.
func (o *SelectExpandQueryOption) HasLevelsMaxLiteralExpansionDepth() bool {
	if o != nil && !IsNil(o.LevelsMaxLiteralExpansionDepth) {
		return true
	}

	return false
}

// SetLevelsMaxLiteralExpansionDepth gets a reference to the given int32 and assigns it to the LevelsMaxLiteralExpansionDepth field.
func (o *SelectExpandQueryOption) SetLevelsMaxLiteralExpansionDepth(v int32) {
	o.LevelsMaxLiteralExpansionDepth = &v
}

func (o SelectExpandQueryOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectExpandQueryOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if o.RawSelect.IsSet() {
		toSerialize["rawSelect"] = o.RawSelect.Get()
	}
	if o.RawExpand.IsSet() {
		toSerialize["rawExpand"] = o.RawExpand.Get()
	}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.Validator) {
		toSerialize["validator"] = o.Validator
	}
	if !IsNil(o.SelectExpandClause) {
		toSerialize["selectExpandClause"] = o.SelectExpandClause
	}
	if !IsNil(o.LevelsMaxLiteralExpansionDepth) {
		toSerialize["levelsMaxLiteralExpansionDepth"] = o.LevelsMaxLiteralExpansionDepth
	}
	return toSerialize, nil
}

type NullableSelectExpandQueryOption struct {
	value *SelectExpandQueryOption
	isSet bool
}

func (v NullableSelectExpandQueryOption) Get() *SelectExpandQueryOption {
	return v.value
}

func (v *NullableSelectExpandQueryOption) Set(val *SelectExpandQueryOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectExpandQueryOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectExpandQueryOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectExpandQueryOption(val *SelectExpandQueryOption) *NullableSelectExpandQueryOption {
	return &NullableSelectExpandQueryOption{value: val, isSet: true}
}

func (v NullableSelectExpandQueryOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectExpandQueryOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
