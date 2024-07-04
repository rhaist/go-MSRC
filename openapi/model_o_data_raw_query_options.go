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

// checks if the ODataRawQueryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ODataRawQueryOptions{}

// ODataRawQueryOptions struct for ODataRawQueryOptions
type ODataRawQueryOptions struct {
	Filter     NullableString `json:"filter,omitempty"`
	Apply      NullableString `json:"apply,omitempty"`
	Compute    NullableString `json:"compute,omitempty"`
	Search     NullableString `json:"search,omitempty"`
	OrderBy    NullableString `json:"orderBy,omitempty"`
	Top        NullableString `json:"top,omitempty"`
	Skip       NullableString `json:"skip,omitempty"`
	Select     NullableString `json:"select,omitempty"`
	Expand     NullableString `json:"expand,omitempty"`
	Count      NullableString `json:"count,omitempty"`
	Format     NullableString `json:"format,omitempty"`
	SkipToken  NullableString `json:"skipToken,omitempty"`
	DeltaToken NullableString `json:"deltaToken,omitempty"`
}

// NewODataRawQueryOptions instantiates a new ODataRawQueryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewODataRawQueryOptions() *ODataRawQueryOptions {
	this := ODataRawQueryOptions{}
	return &this
}

// NewODataRawQueryOptionsWithDefaults instantiates a new ODataRawQueryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewODataRawQueryOptionsWithDefaults() *ODataRawQueryOptions {
	this := ODataRawQueryOptions{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetFilter() string {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *ODataRawQueryOptions) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil
func (o *ODataRawQueryOptions) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetFilter() {
	o.Filter.Unset()
}

// GetApply returns the Apply field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetApply() string {
	if o == nil || IsNil(o.Apply.Get()) {
		var ret string
		return ret
	}
	return *o.Apply.Get()
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetApplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Apply.Get(), o.Apply.IsSet()
}

// HasApply returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasApply() bool {
	if o != nil && o.Apply.IsSet() {
		return true
	}

	return false
}

// SetApply gets a reference to the given NullableString and assigns it to the Apply field.
func (o *ODataRawQueryOptions) SetApply(v string) {
	o.Apply.Set(&v)
}

// SetApplyNil sets the value for Apply to be an explicit nil
func (o *ODataRawQueryOptions) SetApplyNil() {
	o.Apply.Set(nil)
}

// UnsetApply ensures that no value is present for Apply, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetApply() {
	o.Apply.Unset()
}

// GetCompute returns the Compute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetCompute() string {
	if o == nil || IsNil(o.Compute.Get()) {
		var ret string
		return ret
	}
	return *o.Compute.Get()
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetComputeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Compute.Get(), o.Compute.IsSet()
}

// HasCompute returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasCompute() bool {
	if o != nil && o.Compute.IsSet() {
		return true
	}

	return false
}

// SetCompute gets a reference to the given NullableString and assigns it to the Compute field.
func (o *ODataRawQueryOptions) SetCompute(v string) {
	o.Compute.Set(&v)
}

// SetComputeNil sets the value for Compute to be an explicit nil
func (o *ODataRawQueryOptions) SetComputeNil() {
	o.Compute.Set(nil)
}

// UnsetCompute ensures that no value is present for Compute, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetCompute() {
	o.Compute.Unset()
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetSearch() string {
	if o == nil || IsNil(o.Search.Get()) {
		var ret string
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableString and assigns it to the Search field.
func (o *ODataRawQueryOptions) SetSearch(v string) {
	o.Search.Set(&v)
}

// SetSearchNil sets the value for Search to be an explicit nil
func (o *ODataRawQueryOptions) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetSearch() {
	o.Search.Unset()
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy.Get()) {
		var ret string
		return ret
	}
	return *o.OrderBy.Get()
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderBy.Get(), o.OrderBy.IsSet()
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasOrderBy() bool {
	if o != nil && o.OrderBy.IsSet() {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given NullableString and assigns it to the OrderBy field.
func (o *ODataRawQueryOptions) SetOrderBy(v string) {
	o.OrderBy.Set(&v)
}

// SetOrderByNil sets the value for OrderBy to be an explicit nil
func (o *ODataRawQueryOptions) SetOrderByNil() {
	o.OrderBy.Set(nil)
}

// UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetOrderBy() {
	o.OrderBy.Unset()
}

// GetTop returns the Top field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetTop() string {
	if o == nil || IsNil(o.Top.Get()) {
		var ret string
		return ret
	}
	return *o.Top.Get()
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetTopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Top.Get(), o.Top.IsSet()
}

// HasTop returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasTop() bool {
	if o != nil && o.Top.IsSet() {
		return true
	}

	return false
}

// SetTop gets a reference to the given NullableString and assigns it to the Top field.
func (o *ODataRawQueryOptions) SetTop(v string) {
	o.Top.Set(&v)
}

// SetTopNil sets the value for Top to be an explicit nil
func (o *ODataRawQueryOptions) SetTopNil() {
	o.Top.Set(nil)
}

// UnsetTop ensures that no value is present for Top, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetTop() {
	o.Top.Unset()
}

// GetSkip returns the Skip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetSkip() string {
	if o == nil || IsNil(o.Skip.Get()) {
		var ret string
		return ret
	}
	return *o.Skip.Get()
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetSkipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skip.Get(), o.Skip.IsSet()
}

// HasSkip returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasSkip() bool {
	if o != nil && o.Skip.IsSet() {
		return true
	}

	return false
}

// SetSkip gets a reference to the given NullableString and assigns it to the Skip field.
func (o *ODataRawQueryOptions) SetSkip(v string) {
	o.Skip.Set(&v)
}

// SetSkipNil sets the value for Skip to be an explicit nil
func (o *ODataRawQueryOptions) SetSkipNil() {
	o.Skip.Set(nil)
}

// UnsetSkip ensures that no value is present for Skip, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetSkip() {
	o.Skip.Unset()
}

// GetSelect returns the Select field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetSelect() string {
	if o == nil || IsNil(o.Select.Get()) {
		var ret string
		return ret
	}
	return *o.Select.Get()
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetSelectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Select.Get(), o.Select.IsSet()
}

// HasSelect returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasSelect() bool {
	if o != nil && o.Select.IsSet() {
		return true
	}

	return false
}

// SetSelect gets a reference to the given NullableString and assigns it to the Select field.
func (o *ODataRawQueryOptions) SetSelect(v string) {
	o.Select.Set(&v)
}

// SetSelectNil sets the value for Select to be an explicit nil
func (o *ODataRawQueryOptions) SetSelectNil() {
	o.Select.Set(nil)
}

// UnsetSelect ensures that no value is present for Select, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetSelect() {
	o.Select.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetExpand() string {
	if o == nil || IsNil(o.Expand.Get()) {
		var ret string
		return ret
	}
	return *o.Expand.Get()
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetExpandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expand.Get(), o.Expand.IsSet()
}

// HasExpand returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasExpand() bool {
	if o != nil && o.Expand.IsSet() {
		return true
	}

	return false
}

// SetExpand gets a reference to the given NullableString and assigns it to the Expand field.
func (o *ODataRawQueryOptions) SetExpand(v string) {
	o.Expand.Set(&v)
}

// SetExpandNil sets the value for Expand to be an explicit nil
func (o *ODataRawQueryOptions) SetExpandNil() {
	o.Expand.Set(nil)
}

// UnsetExpand ensures that no value is present for Expand, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetExpand() {
	o.Expand.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetCount() string {
	if o == nil || IsNil(o.Count.Get()) {
		var ret string
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableString and assigns it to the Count field.
func (o *ODataRawQueryOptions) SetCount(v string) {
	o.Count.Set(&v)
}

// SetCountNil sets the value for Count to be an explicit nil
func (o *ODataRawQueryOptions) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetCount() {
	o.Count.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetFormat() string {
	if o == nil || IsNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *ODataRawQueryOptions) SetFormat(v string) {
	o.Format.Set(&v)
}

// SetFormatNil sets the value for Format to be an explicit nil
func (o *ODataRawQueryOptions) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetFormat() {
	o.Format.Unset()
}

// GetSkipToken returns the SkipToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetSkipToken() string {
	if o == nil || IsNil(o.SkipToken.Get()) {
		var ret string
		return ret
	}
	return *o.SkipToken.Get()
}

// GetSkipTokenOk returns a tuple with the SkipToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetSkipTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipToken.Get(), o.SkipToken.IsSet()
}

// HasSkipToken returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasSkipToken() bool {
	if o != nil && o.SkipToken.IsSet() {
		return true
	}

	return false
}

// SetSkipToken gets a reference to the given NullableString and assigns it to the SkipToken field.
func (o *ODataRawQueryOptions) SetSkipToken(v string) {
	o.SkipToken.Set(&v)
}

// SetSkipTokenNil sets the value for SkipToken to be an explicit nil
func (o *ODataRawQueryOptions) SetSkipTokenNil() {
	o.SkipToken.Set(nil)
}

// UnsetSkipToken ensures that no value is present for SkipToken, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetSkipToken() {
	o.SkipToken.Unset()
}

// GetDeltaToken returns the DeltaToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataRawQueryOptions) GetDeltaToken() string {
	if o == nil || IsNil(o.DeltaToken.Get()) {
		var ret string
		return ret
	}
	return *o.DeltaToken.Get()
}

// GetDeltaTokenOk returns a tuple with the DeltaToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataRawQueryOptions) GetDeltaTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeltaToken.Get(), o.DeltaToken.IsSet()
}

// HasDeltaToken returns a boolean if a field has been set.
func (o *ODataRawQueryOptions) HasDeltaToken() bool {
	if o != nil && o.DeltaToken.IsSet() {
		return true
	}

	return false
}

// SetDeltaToken gets a reference to the given NullableString and assigns it to the DeltaToken field.
func (o *ODataRawQueryOptions) SetDeltaToken(v string) {
	o.DeltaToken.Set(&v)
}

// SetDeltaTokenNil sets the value for DeltaToken to be an explicit nil
func (o *ODataRawQueryOptions) SetDeltaTokenNil() {
	o.DeltaToken.Set(nil)
}

// UnsetDeltaToken ensures that no value is present for DeltaToken, not even an explicit nil
func (o *ODataRawQueryOptions) UnsetDeltaToken() {
	o.DeltaToken.Unset()
}

func (o ODataRawQueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ODataRawQueryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.Apply.IsSet() {
		toSerialize["apply"] = o.Apply.Get()
	}
	if o.Compute.IsSet() {
		toSerialize["compute"] = o.Compute.Get()
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	if o.OrderBy.IsSet() {
		toSerialize["orderBy"] = o.OrderBy.Get()
	}
	if o.Top.IsSet() {
		toSerialize["top"] = o.Top.Get()
	}
	if o.Skip.IsSet() {
		toSerialize["skip"] = o.Skip.Get()
	}
	if o.Select.IsSet() {
		toSerialize["select"] = o.Select.Get()
	}
	if o.Expand.IsSet() {
		toSerialize["expand"] = o.Expand.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	if o.SkipToken.IsSet() {
		toSerialize["skipToken"] = o.SkipToken.Get()
	}
	if o.DeltaToken.IsSet() {
		toSerialize["deltaToken"] = o.DeltaToken.Get()
	}
	return toSerialize, nil
}

type NullableODataRawQueryOptions struct {
	value *ODataRawQueryOptions
	isSet bool
}

func (v NullableODataRawQueryOptions) Get() *ODataRawQueryOptions {
	return v.value
}

func (v *NullableODataRawQueryOptions) Set(val *ODataRawQueryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableODataRawQueryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableODataRawQueryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableODataRawQueryOptions(val *ODataRawQueryOptions) *NullableODataRawQueryOptions {
	return &NullableODataRawQueryOptions{value: val, isSet: true}
}

func (v NullableODataRawQueryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableODataRawQueryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}