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

// checks if the IEdmReferentialConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmReferentialConstraint{}

// IEdmReferentialConstraint struct for IEdmReferentialConstraint
type IEdmReferentialConstraint struct {
	PropertyPairs []EdmReferentialConstraintPropertyPair `json:"propertyPairs,omitempty"`
}

// NewIEdmReferentialConstraint instantiates a new IEdmReferentialConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmReferentialConstraint() *IEdmReferentialConstraint {
	this := IEdmReferentialConstraint{}
	return &this
}

// NewIEdmReferentialConstraintWithDefaults instantiates a new IEdmReferentialConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmReferentialConstraintWithDefaults() *IEdmReferentialConstraint {
	this := IEdmReferentialConstraint{}
	return &this
}

// GetPropertyPairs returns the PropertyPairs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmReferentialConstraint) GetPropertyPairs() []EdmReferentialConstraintPropertyPair {
	if o == nil {
		var ret []EdmReferentialConstraintPropertyPair
		return ret
	}
	return o.PropertyPairs
}

// GetPropertyPairsOk returns a tuple with the PropertyPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmReferentialConstraint) GetPropertyPairsOk() ([]EdmReferentialConstraintPropertyPair, bool) {
	if o == nil || IsNil(o.PropertyPairs) {
		return nil, false
	}
	return o.PropertyPairs, true
}

// HasPropertyPairs returns a boolean if a field has been set.
func (o *IEdmReferentialConstraint) HasPropertyPairs() bool {
	if o != nil && !IsNil(o.PropertyPairs) {
		return true
	}

	return false
}

// SetPropertyPairs gets a reference to the given []EdmReferentialConstraintPropertyPair and assigns it to the PropertyPairs field.
func (o *IEdmReferentialConstraint) SetPropertyPairs(v []EdmReferentialConstraintPropertyPair) {
	o.PropertyPairs = v
}

func (o IEdmReferentialConstraint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmReferentialConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PropertyPairs != nil {
		toSerialize["propertyPairs"] = o.PropertyPairs
	}
	return toSerialize, nil
}

type NullableIEdmReferentialConstraint struct {
	value *IEdmReferentialConstraint
	isSet bool
}

func (v NullableIEdmReferentialConstraint) Get() *IEdmReferentialConstraint {
	return v.value
}

func (v *NullableIEdmReferentialConstraint) Set(val *IEdmReferentialConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmReferentialConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmReferentialConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmReferentialConstraint(val *IEdmReferentialConstraint) *NullableIEdmReferentialConstraint {
	return &NullableIEdmReferentialConstraint{value: val, isSet: true}
}

func (v NullableIEdmReferentialConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmReferentialConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}