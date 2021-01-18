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

// CvrfdocDocumentPublisherIssuingAuthority struct for CvrfdocDocumentPublisherIssuingAuthority
type CvrfdocDocumentPublisherIssuingAuthority struct {
	Lang  *string `json:"lang,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewCvrfdocDocumentPublisherIssuingAuthority instantiates a new CvrfdocDocumentPublisherIssuingAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdocDocumentPublisherIssuingAuthority() *CvrfdocDocumentPublisherIssuingAuthority {
	this := CvrfdocDocumentPublisherIssuingAuthority{}
	return &this
}

// NewCvrfdocDocumentPublisherIssuingAuthorityWithDefaults instantiates a new CvrfdocDocumentPublisherIssuingAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocDocumentPublisherIssuingAuthorityWithDefaults() *CvrfdocDocumentPublisherIssuingAuthority {
	this := CvrfdocDocumentPublisherIssuingAuthority{}
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisherIssuingAuthority) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisherIssuingAuthority) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisherIssuingAuthority) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *CvrfdocDocumentPublisherIssuingAuthority) SetLang(v string) {
	o.Lang = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisherIssuingAuthority) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisherIssuingAuthority) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisherIssuingAuthority) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CvrfdocDocumentPublisherIssuingAuthority) SetValue(v string) {
	o.Value = &v
}

func (o CvrfdocDocumentPublisherIssuingAuthority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdocDocumentPublisherIssuingAuthority struct {
	value *CvrfdocDocumentPublisherIssuingAuthority
	isSet bool
}

func (v NullableCvrfdocDocumentPublisherIssuingAuthority) Get() *CvrfdocDocumentPublisherIssuingAuthority {
	return v.value
}

func (v *NullableCvrfdocDocumentPublisherIssuingAuthority) Set(val *CvrfdocDocumentPublisherIssuingAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdocDocumentPublisherIssuingAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdocDocumentPublisherIssuingAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdocDocumentPublisherIssuingAuthority(val *CvrfdocDocumentPublisherIssuingAuthority) *NullableCvrfdocDocumentPublisherIssuingAuthority {
	return &NullableCvrfdocDocumentPublisherIssuingAuthority{value: val, isSet: true}
}

func (v NullableCvrfdocDocumentPublisherIssuingAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdocDocumentPublisherIssuingAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
