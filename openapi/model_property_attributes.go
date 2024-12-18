/*
MSRC CVRF API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PropertyAttributes the model 'PropertyAttributes'
type PropertyAttributes string

// List of PropertyAttributes
const (
	NONE            PropertyAttributes = "None"
	SPECIAL_NAME    PropertyAttributes = "SpecialName"
	RT_SPECIAL_NAME PropertyAttributes = "RTSpecialName"
	HAS_DEFAULT     PropertyAttributes = "HasDefault"
	RESERVED2       PropertyAttributes = "Reserved2"
	RESERVED3       PropertyAttributes = "Reserved3"
	RESERVED4       PropertyAttributes = "Reserved4"
	RESERVED_MASK   PropertyAttributes = "ReservedMask"
)

// All allowed values of PropertyAttributes enum
var AllowedPropertyAttributesEnumValues = []PropertyAttributes{
	"None",
	"SpecialName",
	"RTSpecialName",
	"HasDefault",
	"Reserved2",
	"Reserved3",
	"Reserved4",
	"ReservedMask",
}

func (v *PropertyAttributes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PropertyAttributes(value)
	for _, existing := range AllowedPropertyAttributesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PropertyAttributes", value)
}

// NewPropertyAttributesFromValue returns a pointer to a valid PropertyAttributes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPropertyAttributesFromValue(v string) (*PropertyAttributes, error) {
	ev := PropertyAttributes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PropertyAttributes: valid values are %v", v, AllowedPropertyAttributesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PropertyAttributes) IsValid() bool {
	for _, existing := range AllowedPropertyAttributesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PropertyAttributes value
func (v PropertyAttributes) Ptr() *PropertyAttributes {
	return &v
}

type NullablePropertyAttributes struct {
	value *PropertyAttributes
	isSet bool
}

func (v NullablePropertyAttributes) Get() *PropertyAttributes {
	return v.value
}

func (v *NullablePropertyAttributes) Set(val *PropertyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyAttributes(val *PropertyAttributes) *NullablePropertyAttributes {
	return &NullablePropertyAttributes{value: val, isSet: true}
}

func (v NullablePropertyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
