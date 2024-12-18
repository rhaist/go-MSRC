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

// checks if the UpdateODataQueryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateODataQueryOptions{}

// UpdateODataQueryOptions struct for UpdateODataQueryOptions
type UpdateODataQueryOptions struct {
	Request      *HttpRequest             `json:"request,omitempty"`
	Context      *ODataQueryContext       `json:"context,omitempty"`
	RawValues    *ODataRawQueryOptions    `json:"rawValues,omitempty"`
	SelectExpand *SelectExpandQueryOption `json:"selectExpand,omitempty"`
	Apply        *ApplyQueryOption        `json:"apply,omitempty"`
	Compute      *ComputeQueryOption      `json:"compute,omitempty"`
	Filter       *FilterQueryOption       `json:"filter,omitempty"`
	Search       *SearchQueryOption       `json:"search,omitempty"`
	OrderBy      *OrderByQueryOption      `json:"orderBy,omitempty"`
	Skip         *SkipQueryOption         `json:"skip,omitempty"`
	SkipToken    *SkipTokenQueryOption    `json:"skipToken,omitempty"`
	Top          *TopQueryOption          `json:"top,omitempty"`
	Count        *CountQueryOption        `json:"count,omitempty"`
	Validator    map[string]interface{}   `json:"validator,omitempty"`
	IfMatch      *UpdateETag              `json:"ifMatch,omitempty"`
	IfNoneMatch  *UpdateETag              `json:"ifNoneMatch,omitempty"`
}

// NewUpdateODataQueryOptions instantiates a new UpdateODataQueryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateODataQueryOptions() *UpdateODataQueryOptions {
	this := UpdateODataQueryOptions{}
	return &this
}

// NewUpdateODataQueryOptionsWithDefaults instantiates a new UpdateODataQueryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateODataQueryOptionsWithDefaults() *UpdateODataQueryOptions {
	this := UpdateODataQueryOptions{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetRequest() HttpRequest {
	if o == nil || IsNil(o.Request) {
		var ret HttpRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetRequestOk() (*HttpRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given HttpRequest and assigns it to the Request field.
func (o *UpdateODataQueryOptions) SetRequest(v HttpRequest) {
	o.Request = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetContext() ODataQueryContext {
	if o == nil || IsNil(o.Context) {
		var ret ODataQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetContextOk() (*ODataQueryContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given ODataQueryContext and assigns it to the Context field.
func (o *UpdateODataQueryOptions) SetContext(v ODataQueryContext) {
	o.Context = &v
}

// GetRawValues returns the RawValues field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetRawValues() ODataRawQueryOptions {
	if o == nil || IsNil(o.RawValues) {
		var ret ODataRawQueryOptions
		return ret
	}
	return *o.RawValues
}

// GetRawValuesOk returns a tuple with the RawValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetRawValuesOk() (*ODataRawQueryOptions, bool) {
	if o == nil || IsNil(o.RawValues) {
		return nil, false
	}
	return o.RawValues, true
}

// HasRawValues returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasRawValues() bool {
	if o != nil && !IsNil(o.RawValues) {
		return true
	}

	return false
}

// SetRawValues gets a reference to the given ODataRawQueryOptions and assigns it to the RawValues field.
func (o *UpdateODataQueryOptions) SetRawValues(v ODataRawQueryOptions) {
	o.RawValues = &v
}

// GetSelectExpand returns the SelectExpand field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetSelectExpand() SelectExpandQueryOption {
	if o == nil || IsNil(o.SelectExpand) {
		var ret SelectExpandQueryOption
		return ret
	}
	return *o.SelectExpand
}

// GetSelectExpandOk returns a tuple with the SelectExpand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetSelectExpandOk() (*SelectExpandQueryOption, bool) {
	if o == nil || IsNil(o.SelectExpand) {
		return nil, false
	}
	return o.SelectExpand, true
}

// HasSelectExpand returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasSelectExpand() bool {
	if o != nil && !IsNil(o.SelectExpand) {
		return true
	}

	return false
}

// SetSelectExpand gets a reference to the given SelectExpandQueryOption and assigns it to the SelectExpand field.
func (o *UpdateODataQueryOptions) SetSelectExpand(v SelectExpandQueryOption) {
	o.SelectExpand = &v
}

// GetApply returns the Apply field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetApply() ApplyQueryOption {
	if o == nil || IsNil(o.Apply) {
		var ret ApplyQueryOption
		return ret
	}
	return *o.Apply
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetApplyOk() (*ApplyQueryOption, bool) {
	if o == nil || IsNil(o.Apply) {
		return nil, false
	}
	return o.Apply, true
}

// HasApply returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasApply() bool {
	if o != nil && !IsNil(o.Apply) {
		return true
	}

	return false
}

// SetApply gets a reference to the given ApplyQueryOption and assigns it to the Apply field.
func (o *UpdateODataQueryOptions) SetApply(v ApplyQueryOption) {
	o.Apply = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetCompute() ComputeQueryOption {
	if o == nil || IsNil(o.Compute) {
		var ret ComputeQueryOption
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetComputeOk() (*ComputeQueryOption, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ComputeQueryOption and assigns it to the Compute field.
func (o *UpdateODataQueryOptions) SetCompute(v ComputeQueryOption) {
	o.Compute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetFilter() FilterQueryOption {
	if o == nil || IsNil(o.Filter) {
		var ret FilterQueryOption
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetFilterOk() (*FilterQueryOption, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given FilterQueryOption and assigns it to the Filter field.
func (o *UpdateODataQueryOptions) SetFilter(v FilterQueryOption) {
	o.Filter = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetSearch() SearchQueryOption {
	if o == nil || IsNil(o.Search) {
		var ret SearchQueryOption
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetSearchOk() (*SearchQueryOption, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given SearchQueryOption and assigns it to the Search field.
func (o *UpdateODataQueryOptions) SetSearch(v SearchQueryOption) {
	o.Search = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetOrderBy() OrderByQueryOption {
	if o == nil || IsNil(o.OrderBy) {
		var ret OrderByQueryOption
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetOrderByOk() (*OrderByQueryOption, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given OrderByQueryOption and assigns it to the OrderBy field.
func (o *UpdateODataQueryOptions) SetOrderBy(v OrderByQueryOption) {
	o.OrderBy = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetSkip() SkipQueryOption {
	if o == nil || IsNil(o.Skip) {
		var ret SkipQueryOption
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetSkipOk() (*SkipQueryOption, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given SkipQueryOption and assigns it to the Skip field.
func (o *UpdateODataQueryOptions) SetSkip(v SkipQueryOption) {
	o.Skip = &v
}

// GetSkipToken returns the SkipToken field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetSkipToken() SkipTokenQueryOption {
	if o == nil || IsNil(o.SkipToken) {
		var ret SkipTokenQueryOption
		return ret
	}
	return *o.SkipToken
}

// GetSkipTokenOk returns a tuple with the SkipToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetSkipTokenOk() (*SkipTokenQueryOption, bool) {
	if o == nil || IsNil(o.SkipToken) {
		return nil, false
	}
	return o.SkipToken, true
}

// HasSkipToken returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasSkipToken() bool {
	if o != nil && !IsNil(o.SkipToken) {
		return true
	}

	return false
}

// SetSkipToken gets a reference to the given SkipTokenQueryOption and assigns it to the SkipToken field.
func (o *UpdateODataQueryOptions) SetSkipToken(v SkipTokenQueryOption) {
	o.SkipToken = &v
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetTop() TopQueryOption {
	if o == nil || IsNil(o.Top) {
		var ret TopQueryOption
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetTopOk() (*TopQueryOption, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given TopQueryOption and assigns it to the Top field.
func (o *UpdateODataQueryOptions) SetTop(v TopQueryOption) {
	o.Top = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetCount() CountQueryOption {
	if o == nil || IsNil(o.Count) {
		var ret CountQueryOption
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetCountOk() (*CountQueryOption, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given CountQueryOption and assigns it to the Count field.
func (o *UpdateODataQueryOptions) SetCount(v CountQueryOption) {
	o.Count = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetValidator() map[string]interface{} {
	if o == nil || IsNil(o.Validator) {
		var ret map[string]interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetValidatorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Validator) {
		return map[string]interface{}{}, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasValidator() bool {
	if o != nil && !IsNil(o.Validator) {
		return true
	}

	return false
}

// SetValidator gets a reference to the given map[string]interface{} and assigns it to the Validator field.
func (o *UpdateODataQueryOptions) SetValidator(v map[string]interface{}) {
	o.Validator = v
}

// GetIfMatch returns the IfMatch field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetIfMatch() UpdateETag {
	if o == nil || IsNil(o.IfMatch) {
		var ret UpdateETag
		return ret
	}
	return *o.IfMatch
}

// GetIfMatchOk returns a tuple with the IfMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetIfMatchOk() (*UpdateETag, bool) {
	if o == nil || IsNil(o.IfMatch) {
		return nil, false
	}
	return o.IfMatch, true
}

// HasIfMatch returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasIfMatch() bool {
	if o != nil && !IsNil(o.IfMatch) {
		return true
	}

	return false
}

// SetIfMatch gets a reference to the given UpdateETag and assigns it to the IfMatch field.
func (o *UpdateODataQueryOptions) SetIfMatch(v UpdateETag) {
	o.IfMatch = &v
}

// GetIfNoneMatch returns the IfNoneMatch field value if set, zero value otherwise.
func (o *UpdateODataQueryOptions) GetIfNoneMatch() UpdateETag {
	if o == nil || IsNil(o.IfNoneMatch) {
		var ret UpdateETag
		return ret
	}
	return *o.IfNoneMatch
}

// GetIfNoneMatchOk returns a tuple with the IfNoneMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateODataQueryOptions) GetIfNoneMatchOk() (*UpdateETag, bool) {
	if o == nil || IsNil(o.IfNoneMatch) {
		return nil, false
	}
	return o.IfNoneMatch, true
}

// HasIfNoneMatch returns a boolean if a field has been set.
func (o *UpdateODataQueryOptions) HasIfNoneMatch() bool {
	if o != nil && !IsNil(o.IfNoneMatch) {
		return true
	}

	return false
}

// SetIfNoneMatch gets a reference to the given UpdateETag and assigns it to the IfNoneMatch field.
func (o *UpdateODataQueryOptions) SetIfNoneMatch(v UpdateETag) {
	o.IfNoneMatch = &v
}

func (o UpdateODataQueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateODataQueryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.RawValues) {
		toSerialize["rawValues"] = o.RawValues
	}
	if !IsNil(o.SelectExpand) {
		toSerialize["selectExpand"] = o.SelectExpand
	}
	if !IsNil(o.Apply) {
		toSerialize["apply"] = o.Apply
	}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.SkipToken) {
		toSerialize["skipToken"] = o.SkipToken
	}
	if !IsNil(o.Top) {
		toSerialize["top"] = o.Top
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Validator) {
		toSerialize["validator"] = o.Validator
	}
	if !IsNil(o.IfMatch) {
		toSerialize["ifMatch"] = o.IfMatch
	}
	if !IsNil(o.IfNoneMatch) {
		toSerialize["ifNoneMatch"] = o.IfNoneMatch
	}
	return toSerialize, nil
}

type NullableUpdateODataQueryOptions struct {
	value *UpdateODataQueryOptions
	isSet bool
}

func (v NullableUpdateODataQueryOptions) Get() *UpdateODataQueryOptions {
	return v.value
}

func (v *NullableUpdateODataQueryOptions) Set(val *UpdateODataQueryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateODataQueryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateODataQueryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateODataQueryOptions(val *UpdateODataQueryOptions) *NullableUpdateODataQueryOptions {
	return &NullableUpdateODataQueryOptions{value: val, isSet: true}
}

func (v NullableUpdateODataQueryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateODataQueryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
