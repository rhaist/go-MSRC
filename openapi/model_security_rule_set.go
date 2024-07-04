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

// SecurityRuleSet the model 'SecurityRuleSet'
type SecurityRuleSet string

// List of SecurityRuleSet
const (
	NONE   SecurityRuleSet = "None"
	LEVEL1 SecurityRuleSet = "Level1"
	LEVEL2 SecurityRuleSet = "Level2"
)

// All allowed values of SecurityRuleSet enum
var AllowedSecurityRuleSetEnumValues = []SecurityRuleSet{
	"None",
	"Level1",
	"Level2",
}

func (v *SecurityRuleSet) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityRuleSet(value)
	for _, existing := range AllowedSecurityRuleSetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityRuleSet", value)
}

// NewSecurityRuleSetFromValue returns a pointer to a valid SecurityRuleSet
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityRuleSetFromValue(v string) (*SecurityRuleSet, error) {
	ev := SecurityRuleSet(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityRuleSet: valid values are %v", v, AllowedSecurityRuleSetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityRuleSet) IsValid() bool {
	for _, existing := range AllowedSecurityRuleSetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuleSet value
func (v SecurityRuleSet) Ptr() *SecurityRuleSet {
	return &v
}

type NullableSecurityRuleSet struct {
	value *SecurityRuleSet
	isSet bool
}

func (v NullableSecurityRuleSet) Get() *SecurityRuleSet {
	return v.value
}

func (v *NullableSecurityRuleSet) Set(val *SecurityRuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityRuleSet(val *SecurityRuleSet) *NullableSecurityRuleSet {
	return &NullableSecurityRuleSet{value: val, isSet: true}
}

func (v NullableSecurityRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
