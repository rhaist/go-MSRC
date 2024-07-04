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

// checks if the X500DistinguishedName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X500DistinguishedName{}

// X500DistinguishedName struct for X500DistinguishedName
type X500DistinguishedName struct {
	Oid     *Oid           `json:"oid,omitempty"`
	RawData NullableString `json:"rawData,omitempty"`
	Name    NullableString `json:"name,omitempty"`
}

// NewX500DistinguishedName instantiates a new X500DistinguishedName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX500DistinguishedName() *X500DistinguishedName {
	this := X500DistinguishedName{}
	return &this
}

// NewX500DistinguishedNameWithDefaults instantiates a new X500DistinguishedName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX500DistinguishedNameWithDefaults() *X500DistinguishedName {
	this := X500DistinguishedName{}
	return &this
}

// GetOid returns the Oid field value if set, zero value otherwise.
func (o *X500DistinguishedName) GetOid() Oid {
	if o == nil || IsNil(o.Oid) {
		var ret Oid
		return ret
	}
	return *o.Oid
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X500DistinguishedName) GetOidOk() (*Oid, bool) {
	if o == nil || IsNil(o.Oid) {
		return nil, false
	}
	return o.Oid, true
}

// HasOid returns a boolean if a field has been set.
func (o *X500DistinguishedName) HasOid() bool {
	if o != nil && !IsNil(o.Oid) {
		return true
	}

	return false
}

// SetOid gets a reference to the given Oid and assigns it to the Oid field.
func (o *X500DistinguishedName) SetOid(v Oid) {
	o.Oid = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X500DistinguishedName) GetRawData() string {
	if o == nil || IsNil(o.RawData.Get()) {
		var ret string
		return ret
	}
	return *o.RawData.Get()
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X500DistinguishedName) GetRawDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawData.Get(), o.RawData.IsSet()
}

// HasRawData returns a boolean if a field has been set.
func (o *X500DistinguishedName) HasRawData() bool {
	if o != nil && o.RawData.IsSet() {
		return true
	}

	return false
}

// SetRawData gets a reference to the given NullableString and assigns it to the RawData field.
func (o *X500DistinguishedName) SetRawData(v string) {
	o.RawData.Set(&v)
}

// SetRawDataNil sets the value for RawData to be an explicit nil
func (o *X500DistinguishedName) SetRawDataNil() {
	o.RawData.Set(nil)
}

// UnsetRawData ensures that no value is present for RawData, not even an explicit nil
func (o *X500DistinguishedName) UnsetRawData() {
	o.RawData.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X500DistinguishedName) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X500DistinguishedName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *X500DistinguishedName) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *X500DistinguishedName) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *X500DistinguishedName) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *X500DistinguishedName) UnsetName() {
	o.Name.Unset()
}

func (o X500DistinguishedName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o X500DistinguishedName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oid) {
		toSerialize["oid"] = o.Oid
	}
	if o.RawData.IsSet() {
		toSerialize["rawData"] = o.RawData.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableX500DistinguishedName struct {
	value *X500DistinguishedName
	isSet bool
}

func (v NullableX500DistinguishedName) Get() *X500DistinguishedName {
	return v.value
}

func (v *NullableX500DistinguishedName) Set(val *X500DistinguishedName) {
	v.value = val
	v.isSet = true
}

func (v NullableX500DistinguishedName) IsSet() bool {
	return v.isSet
}

func (v *NullableX500DistinguishedName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX500DistinguishedName(val *X500DistinguishedName) *NullableX500DistinguishedName {
	return &NullableX500DistinguishedName{value: val, isSet: true}
}

func (v NullableX500DistinguishedName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX500DistinguishedName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
