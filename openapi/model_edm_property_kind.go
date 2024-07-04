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

// EdmPropertyKind the model 'EdmPropertyKind'
type EdmPropertyKind string

// List of EdmPropertyKind
const (
	NONE       EdmPropertyKind = "None"
	STRUCTURAL EdmPropertyKind = "Structural"
	NAVIGATION EdmPropertyKind = "Navigation"
)

// All allowed values of EdmPropertyKind enum
var AllowedEdmPropertyKindEnumValues = []EdmPropertyKind{
	"None",
	"Structural",
	"Navigation",
}

func (v *EdmPropertyKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EdmPropertyKind(value)
	for _, existing := range AllowedEdmPropertyKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EdmPropertyKind", value)
}

// NewEdmPropertyKindFromValue returns a pointer to a valid EdmPropertyKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEdmPropertyKindFromValue(v string) (*EdmPropertyKind, error) {
	ev := EdmPropertyKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EdmPropertyKind: valid values are %v", v, AllowedEdmPropertyKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EdmPropertyKind) IsValid() bool {
	for _, existing := range AllowedEdmPropertyKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EdmPropertyKind value
func (v EdmPropertyKind) Ptr() *EdmPropertyKind {
	return &v
}

type NullableEdmPropertyKind struct {
	value *EdmPropertyKind
	isSet bool
}

func (v NullableEdmPropertyKind) Get() *EdmPropertyKind {
	return v.value
}

func (v *NullableEdmPropertyKind) Set(val *EdmPropertyKind) {
	v.value = val
	v.isSet = true
}

func (v NullableEdmPropertyKind) IsSet() bool {
	return v.isSet
}

func (v *NullableEdmPropertyKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdmPropertyKind(val *EdmPropertyKind) *NullableEdmPropertyKind {
	return &NullableEdmPropertyKind{value: val, isSet: true}
}

func (v NullableEdmPropertyKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdmPropertyKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}