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

// checks if the Claim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Claim{}

// Claim struct for Claim
type Claim struct {
	Issuer         NullableString    `json:"issuer,omitempty"`
	OriginalIssuer NullableString    `json:"originalIssuer,omitempty"`
	Properties     map[string]string `json:"properties,omitempty"`
	Subject        *ClaimsIdentity   `json:"subject,omitempty"`
	Type           NullableString    `json:"type,omitempty"`
	Value          NullableString    `json:"value,omitempty"`
	ValueType      NullableString    `json:"valueType,omitempty"`
}

// NewClaim instantiates a new Claim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaim() *Claim {
	this := Claim{}
	return &this
}

// NewClaimWithDefaults instantiates a new Claim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimWithDefaults() *Claim {
	this := Claim{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *Claim) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *Claim) SetIssuer(v string) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *Claim) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *Claim) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetOriginalIssuer returns the OriginalIssuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetOriginalIssuer() string {
	if o == nil || IsNil(o.OriginalIssuer.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalIssuer.Get()
}

// GetOriginalIssuerOk returns a tuple with the OriginalIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetOriginalIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalIssuer.Get(), o.OriginalIssuer.IsSet()
}

// HasOriginalIssuer returns a boolean if a field has been set.
func (o *Claim) HasOriginalIssuer() bool {
	if o != nil && o.OriginalIssuer.IsSet() {
		return true
	}

	return false
}

// SetOriginalIssuer gets a reference to the given NullableString and assigns it to the OriginalIssuer field.
func (o *Claim) SetOriginalIssuer(v string) {
	o.OriginalIssuer.Set(&v)
}

// SetOriginalIssuerNil sets the value for OriginalIssuer to be an explicit nil
func (o *Claim) SetOriginalIssuerNil() {
	o.OriginalIssuer.Set(nil)
}

// UnsetOriginalIssuer ensures that no value is present for OriginalIssuer, not even an explicit nil
func (o *Claim) UnsetOriginalIssuer() {
	o.OriginalIssuer.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Claim) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *Claim) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Claim) GetSubject() ClaimsIdentity {
	if o == nil || IsNil(o.Subject) {
		var ret ClaimsIdentity
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Claim) GetSubjectOk() (*ClaimsIdentity, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Claim) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given ClaimsIdentity and assigns it to the Subject field.
func (o *Claim) SetSubject(v ClaimsIdentity) {
	o.Subject = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Claim) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Claim) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *Claim) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Claim) UnsetType() {
	o.Type.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Claim) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *Claim) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *Claim) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Claim) UnsetValue() {
	o.Value.Unset()
}

// GetValueType returns the ValueType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Claim) GetValueType() string {
	if o == nil || IsNil(o.ValueType.Get()) {
		var ret string
		return ret
	}
	return *o.ValueType.Get()
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Claim) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueType.Get(), o.ValueType.IsSet()
}

// HasValueType returns a boolean if a field has been set.
func (o *Claim) HasValueType() bool {
	if o != nil && o.ValueType.IsSet() {
		return true
	}

	return false
}

// SetValueType gets a reference to the given NullableString and assigns it to the ValueType field.
func (o *Claim) SetValueType(v string) {
	o.ValueType.Set(&v)
}

// SetValueTypeNil sets the value for ValueType to be an explicit nil
func (o *Claim) SetValueTypeNil() {
	o.ValueType.Set(nil)
}

// UnsetValueType ensures that no value is present for ValueType, not even an explicit nil
func (o *Claim) UnsetValueType() {
	o.ValueType.Unset()
}

func (o Claim) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Claim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.OriginalIssuer.IsSet() {
		toSerialize["originalIssuer"] = o.OriginalIssuer.Get()
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.ValueType.IsSet() {
		toSerialize["valueType"] = o.ValueType.Get()
	}
	return toSerialize, nil
}

type NullableClaim struct {
	value *Claim
	isSet bool
}

func (v NullableClaim) Get() *Claim {
	return v.value
}

func (v *NullableClaim) Set(val *Claim) {
	v.value = val
	v.isSet = true
}

func (v NullableClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaim(val *Claim) *NullableClaim {
	return &NullableClaim{value: val, isSet: true}
}

func (v NullableClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
