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

// checks if the ClaimsPrincipal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimsPrincipal{}

// ClaimsPrincipal struct for ClaimsPrincipal
type ClaimsPrincipal struct {
	Claims     []Claim          `json:"claims,omitempty"`
	Identities []ClaimsIdentity `json:"identities,omitempty"`
	Identity   *IIdentity       `json:"identity,omitempty"`
}

// NewClaimsPrincipal instantiates a new ClaimsPrincipal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimsPrincipal() *ClaimsPrincipal {
	this := ClaimsPrincipal{}
	return &this
}

// NewClaimsPrincipalWithDefaults instantiates a new ClaimsPrincipal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimsPrincipalWithDefaults() *ClaimsPrincipal {
	this := ClaimsPrincipal{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClaimsPrincipal) GetClaims() []Claim {
	if o == nil {
		var ret []Claim
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClaimsPrincipal) GetClaimsOk() ([]Claim, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *ClaimsPrincipal) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []Claim and assigns it to the Claims field.
func (o *ClaimsPrincipal) SetClaims(v []Claim) {
	o.Claims = v
}

// GetIdentities returns the Identities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClaimsPrincipal) GetIdentities() []ClaimsIdentity {
	if o == nil {
		var ret []ClaimsIdentity
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClaimsPrincipal) GetIdentitiesOk() ([]ClaimsIdentity, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *ClaimsPrincipal) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []ClaimsIdentity and assigns it to the Identities field.
func (o *ClaimsPrincipal) SetIdentities(v []ClaimsIdentity) {
	o.Identities = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ClaimsPrincipal) GetIdentity() IIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret IIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimsPrincipal) GetIdentityOk() (*IIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ClaimsPrincipal) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given IIdentity and assigns it to the Identity field.
func (o *ClaimsPrincipal) SetIdentity(v IIdentity) {
	o.Identity = &v
}

func (o ClaimsPrincipal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimsPrincipal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return toSerialize, nil
}

type NullableClaimsPrincipal struct {
	value *ClaimsPrincipal
	isSet bool
}

func (v NullableClaimsPrincipal) Get() *ClaimsPrincipal {
	return v.value
}

func (v *NullableClaimsPrincipal) Set(val *ClaimsPrincipal) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimsPrincipal) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimsPrincipal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimsPrincipal(val *ClaimsPrincipal) *NullableClaimsPrincipal {
	return &NullableClaimsPrincipal{value: val, isSet: true}
}

func (v NullableClaimsPrincipal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimsPrincipal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}