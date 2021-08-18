/*
MSRC CVRF API

Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductTreeGroupDescription struct for ProductTreeGroupDescription
type ProductTreeGroupDescription struct {
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewProductTreeGroupDescription instantiates a new ProductTreeGroupDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTreeGroupDescription() *ProductTreeGroupDescription {
	this := ProductTreeGroupDescription{}
	return &this
}

// NewProductTreeGroupDescriptionWithDefaults instantiates a new ProductTreeGroupDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTreeGroupDescriptionWithDefaults() *ProductTreeGroupDescription {
	this := ProductTreeGroupDescription{}
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ProductTreeGroupDescription) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeGroupDescription) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ProductTreeGroupDescription) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ProductTreeGroupDescription) SetLang(v string) {
	o.Lang = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductTreeGroupDescription) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeGroupDescription) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductTreeGroupDescription) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProductTreeGroupDescription) SetValue(v string) {
	o.Value = &v
}

func (o ProductTreeGroupDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableProductTreeGroupDescription struct {
	value *ProductTreeGroupDescription
	isSet bool
}

func (v NullableProductTreeGroupDescription) Get() *ProductTreeGroupDescription {
	return v.value
}

func (v *NullableProductTreeGroupDescription) Set(val *ProductTreeGroupDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTreeGroupDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTreeGroupDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTreeGroupDescription(val *ProductTreeGroupDescription) *NullableProductTreeGroupDescription {
	return &NullableProductTreeGroupDescription{value: val, isSet: true}
}

func (v NullableProductTreeGroupDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTreeGroupDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
