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

// checks if the AsnEncodedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsnEncodedData{}

// AsnEncodedData struct for AsnEncodedData
type AsnEncodedData struct {
	Oid     *Oid           `json:"oid,omitempty"`
	RawData NullableString `json:"rawData,omitempty"`
}

// NewAsnEncodedData instantiates a new AsnEncodedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsnEncodedData() *AsnEncodedData {
	this := AsnEncodedData{}
	return &this
}

// NewAsnEncodedDataWithDefaults instantiates a new AsnEncodedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsnEncodedDataWithDefaults() *AsnEncodedData {
	this := AsnEncodedData{}
	return &this
}

// GetOid returns the Oid field value if set, zero value otherwise.
func (o *AsnEncodedData) GetOid() Oid {
	if o == nil || IsNil(o.Oid) {
		var ret Oid
		return ret
	}
	return *o.Oid
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsnEncodedData) GetOidOk() (*Oid, bool) {
	if o == nil || IsNil(o.Oid) {
		return nil, false
	}
	return o.Oid, true
}

// HasOid returns a boolean if a field has been set.
func (o *AsnEncodedData) HasOid() bool {
	if o != nil && !IsNil(o.Oid) {
		return true
	}

	return false
}

// SetOid gets a reference to the given Oid and assigns it to the Oid field.
func (o *AsnEncodedData) SetOid(v Oid) {
	o.Oid = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AsnEncodedData) GetRawData() string {
	if o == nil || IsNil(o.RawData.Get()) {
		var ret string
		return ret
	}
	return *o.RawData.Get()
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AsnEncodedData) GetRawDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawData.Get(), o.RawData.IsSet()
}

// HasRawData returns a boolean if a field has been set.
func (o *AsnEncodedData) HasRawData() bool {
	if o != nil && o.RawData.IsSet() {
		return true
	}

	return false
}

// SetRawData gets a reference to the given NullableString and assigns it to the RawData field.
func (o *AsnEncodedData) SetRawData(v string) {
	o.RawData.Set(&v)
}

// SetRawDataNil sets the value for RawData to be an explicit nil
func (o *AsnEncodedData) SetRawDataNil() {
	o.RawData.Set(nil)
}

// UnsetRawData ensures that no value is present for RawData, not even an explicit nil
func (o *AsnEncodedData) UnsetRawData() {
	o.RawData.Unset()
}

func (o AsnEncodedData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsnEncodedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oid) {
		toSerialize["oid"] = o.Oid
	}
	if o.RawData.IsSet() {
		toSerialize["rawData"] = o.RawData.Get()
	}
	return toSerialize, nil
}

type NullableAsnEncodedData struct {
	value *AsnEncodedData
	isSet bool
}

func (v NullableAsnEncodedData) Get() *AsnEncodedData {
	return v.value
}

func (v *NullableAsnEncodedData) Set(val *AsnEncodedData) {
	v.value = val
	v.isSet = true
}

func (v NullableAsnEncodedData) IsSet() bool {
	return v.isSet
}

func (v *NullableAsnEncodedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsnEncodedData(val *AsnEncodedData) *NullableAsnEncodedData {
	return &NullableAsnEncodedData{value: val, isSet: true}
}

func (v NullableAsnEncodedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsnEncodedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
