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

// checks if the IIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IIdentity{}

// IIdentity struct for IIdentity
type IIdentity struct {
	Name               NullableString `json:"name,omitempty"`
	AuthenticationType NullableString `json:"authenticationType,omitempty"`
	IsAuthenticated    *bool          `json:"isAuthenticated,omitempty"`
}

// NewIIdentity instantiates a new IIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIIdentity() *IIdentity {
	this := IIdentity{}
	return &this
}

// NewIIdentityWithDefaults instantiates a new IIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIIdentityWithDefaults() *IIdentity {
	this := IIdentity{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IIdentity) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IIdentity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IIdentity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IIdentity) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *IIdentity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IIdentity) UnsetName() {
	o.Name.Unset()
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IIdentity) GetAuthenticationType() string {
	if o == nil || IsNil(o.AuthenticationType.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationType.Get()
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IIdentity) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationType.Get(), o.AuthenticationType.IsSet()
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *IIdentity) HasAuthenticationType() bool {
	if o != nil && o.AuthenticationType.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given NullableString and assigns it to the AuthenticationType field.
func (o *IIdentity) SetAuthenticationType(v string) {
	o.AuthenticationType.Set(&v)
}

// SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil
func (o *IIdentity) SetAuthenticationTypeNil() {
	o.AuthenticationType.Set(nil)
}

// UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
func (o *IIdentity) UnsetAuthenticationType() {
	o.AuthenticationType.Unset()
}

// GetIsAuthenticated returns the IsAuthenticated field value if set, zero value otherwise.
func (o *IIdentity) GetIsAuthenticated() bool {
	if o == nil || IsNil(o.IsAuthenticated) {
		var ret bool
		return ret
	}
	return *o.IsAuthenticated
}

// GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IIdentity) GetIsAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthenticated) {
		return nil, false
	}
	return o.IsAuthenticated, true
}

// HasIsAuthenticated returns a boolean if a field has been set.
func (o *IIdentity) HasIsAuthenticated() bool {
	if o != nil && !IsNil(o.IsAuthenticated) {
		return true
	}

	return false
}

// SetIsAuthenticated gets a reference to the given bool and assigns it to the IsAuthenticated field.
func (o *IIdentity) SetIsAuthenticated(v bool) {
	o.IsAuthenticated = &v
}

func (o IIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AuthenticationType.IsSet() {
		toSerialize["authenticationType"] = o.AuthenticationType.Get()
	}
	if !IsNil(o.IsAuthenticated) {
		toSerialize["isAuthenticated"] = o.IsAuthenticated
	}
	return toSerialize, nil
}

type NullableIIdentity struct {
	value *IIdentity
	isSet bool
}

func (v NullableIIdentity) Get() *IIdentity {
	return v.value
}

func (v *NullableIIdentity) Set(val *IIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableIIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIIdentity(val *IIdentity) *NullableIIdentity {
	return &NullableIIdentity{value: val, isSet: true}
}

func (v NullableIIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
