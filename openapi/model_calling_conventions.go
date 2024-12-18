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

// CallingConventions the model 'CallingConventions'
type CallingConventions string

// List of CallingConventions
const (
	STANDARD      CallingConventions = "Standard"
	VAR_ARGS      CallingConventions = "VarArgs"
	ANY           CallingConventions = "Any"
	HAS_THIS      CallingConventions = "HasThis"
	EXPLICIT_THIS CallingConventions = "ExplicitThis"
)

// All allowed values of CallingConventions enum
var AllowedCallingConventionsEnumValues = []CallingConventions{
	"Standard",
	"VarArgs",
	"Any",
	"HasThis",
	"ExplicitThis",
}

func (v *CallingConventions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CallingConventions(value)
	for _, existing := range AllowedCallingConventionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CallingConventions", value)
}

// NewCallingConventionsFromValue returns a pointer to a valid CallingConventions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCallingConventionsFromValue(v string) (*CallingConventions, error) {
	ev := CallingConventions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CallingConventions: valid values are %v", v, AllowedCallingConventionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CallingConventions) IsValid() bool {
	for _, existing := range AllowedCallingConventionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CallingConventions value
func (v CallingConventions) Ptr() *CallingConventions {
	return &v
}

type NullableCallingConventions struct {
	value *CallingConventions
	isSet bool
}

func (v NullableCallingConventions) Get() *CallingConventions {
	return v.value
}

func (v *NullableCallingConventions) Set(val *CallingConventions) {
	v.value = val
	v.isSet = true
}

func (v NullableCallingConventions) IsSet() bool {
	return v.isSet
}

func (v *NullableCallingConventions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallingConventions(val *CallingConventions) *NullableCallingConventions {
	return &NullableCallingConventions{value: val, isSet: true}
}

func (v NullableCallingConventions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallingConventions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
