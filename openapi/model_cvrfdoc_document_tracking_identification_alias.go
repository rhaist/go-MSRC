/*
 * MSRC Public API
 *
 * This RESTful API can be used to engage the Microsoft Security Response Center (MSRC) in the following ways:    - Get security update summaries and details using the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf) (CVRF).    - Report suspected cyberattacks or abuse originating from Microsoft Online Services.    - Notify Microsoft of any planned penetration tests against your Azure assets.    **Sample client code** is available on the Microsoft Security [Updates](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API) and [Engage](https://github.com/Microsoft/MSRC-Microsoft-Engage-API)   Github repositories.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CvrfdocDocumentTrackingIdentificationAlias struct for CvrfdocDocumentTrackingIdentificationAlias
type CvrfdocDocumentTrackingIdentificationAlias struct {
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewCvrfdocDocumentTrackingIdentificationAlias instantiates a new CvrfdocDocumentTrackingIdentificationAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdocDocumentTrackingIdentificationAlias() *CvrfdocDocumentTrackingIdentificationAlias {
	this := CvrfdocDocumentTrackingIdentificationAlias{}
	return &this
}

// NewCvrfdocDocumentTrackingIdentificationAliasWithDefaults instantiates a new CvrfdocDocumentTrackingIdentificationAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocDocumentTrackingIdentificationAliasWithDefaults() *CvrfdocDocumentTrackingIdentificationAlias {
	this := CvrfdocDocumentTrackingIdentificationAlias{}
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *CvrfdocDocumentTrackingIdentificationAlias) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentTrackingIdentificationAlias) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *CvrfdocDocumentTrackingIdentificationAlias) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *CvrfdocDocumentTrackingIdentificationAlias) SetLang(v string) {
	o.Lang = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CvrfdocDocumentTrackingIdentificationAlias) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentTrackingIdentificationAlias) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CvrfdocDocumentTrackingIdentificationAlias) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CvrfdocDocumentTrackingIdentificationAlias) SetValue(v string) {
	o.Value = &v
}

func (o CvrfdocDocumentTrackingIdentificationAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdocDocumentTrackingIdentificationAlias struct {
	value *CvrfdocDocumentTrackingIdentificationAlias
	isSet bool
}

func (v NullableCvrfdocDocumentTrackingIdentificationAlias) Get() *CvrfdocDocumentTrackingIdentificationAlias {
	return v.value
}

func (v *NullableCvrfdocDocumentTrackingIdentificationAlias) Set(val *CvrfdocDocumentTrackingIdentificationAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdocDocumentTrackingIdentificationAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdocDocumentTrackingIdentificationAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdocDocumentTrackingIdentificationAlias(val *CvrfdocDocumentTrackingIdentificationAlias) *NullableCvrfdocDocumentTrackingIdentificationAlias {
	return &NullableCvrfdocDocumentTrackingIdentificationAlias{value: val, isSet: true}
}

func (v NullableCvrfdocDocumentTrackingIdentificationAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdocDocumentTrackingIdentificationAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
