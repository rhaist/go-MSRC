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

// GenericParameterAttributes the model 'GenericParameterAttributes'
type GenericParameterAttributes string

// List of GenericParameterAttributes
const (
	NONE                               GenericParameterAttributes = "None"
	COVARIANT                          GenericParameterAttributes = "Covariant"
	CONTRAVARIANT                      GenericParameterAttributes = "Contravariant"
	VARIANCE_MASK                      GenericParameterAttributes = "VarianceMask"
	REFERENCE_TYPE_CONSTRAINT          GenericParameterAttributes = "ReferenceTypeConstraint"
	NOT_NULLABLE_VALUE_TYPE_CONSTRAINT GenericParameterAttributes = "NotNullableValueTypeConstraint"
	DEFAULT_CONSTRUCTOR_CONSTRAINT     GenericParameterAttributes = "DefaultConstructorConstraint"
	SPECIAL_CONSTRAINT_MASK            GenericParameterAttributes = "SpecialConstraintMask"
)

// All allowed values of GenericParameterAttributes enum
var AllowedGenericParameterAttributesEnumValues = []GenericParameterAttributes{
	"None",
	"Covariant",
	"Contravariant",
	"VarianceMask",
	"ReferenceTypeConstraint",
	"NotNullableValueTypeConstraint",
	"DefaultConstructorConstraint",
	"SpecialConstraintMask",
}

func (v *GenericParameterAttributes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GenericParameterAttributes(value)
	for _, existing := range AllowedGenericParameterAttributesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GenericParameterAttributes", value)
}

// NewGenericParameterAttributesFromValue returns a pointer to a valid GenericParameterAttributes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGenericParameterAttributesFromValue(v string) (*GenericParameterAttributes, error) {
	ev := GenericParameterAttributes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GenericParameterAttributes: valid values are %v", v, AllowedGenericParameterAttributesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GenericParameterAttributes) IsValid() bool {
	for _, existing := range AllowedGenericParameterAttributesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GenericParameterAttributes value
func (v GenericParameterAttributes) Ptr() *GenericParameterAttributes {
	return &v
}

type NullableGenericParameterAttributes struct {
	value *GenericParameterAttributes
	isSet bool
}

func (v NullableGenericParameterAttributes) Get() *GenericParameterAttributes {
	return v.value
}

func (v *NullableGenericParameterAttributes) Set(val *GenericParameterAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericParameterAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericParameterAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericParameterAttributes(val *GenericParameterAttributes) *NullableGenericParameterAttributes {
	return &NullableGenericParameterAttributes{value: val, isSet: true}
}

func (v NullableGenericParameterAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericParameterAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
