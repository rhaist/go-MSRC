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

// TransformationNodeKind the model 'TransformationNodeKind'
type TransformationNodeKind string

// List of TransformationNodeKind
const (
	AGGREGATE TransformationNodeKind = "Aggregate"
	GROUP_BY  TransformationNodeKind = "GroupBy"
	FILTER    TransformationNodeKind = "Filter"
	COMPUTE   TransformationNodeKind = "Compute"
	EXPAND    TransformationNodeKind = "Expand"
)

// All allowed values of TransformationNodeKind enum
var AllowedTransformationNodeKindEnumValues = []TransformationNodeKind{
	"Aggregate",
	"GroupBy",
	"Filter",
	"Compute",
	"Expand",
}

func (v *TransformationNodeKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransformationNodeKind(value)
	for _, existing := range AllowedTransformationNodeKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransformationNodeKind", value)
}

// NewTransformationNodeKindFromValue returns a pointer to a valid TransformationNodeKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransformationNodeKindFromValue(v string) (*TransformationNodeKind, error) {
	ev := TransformationNodeKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransformationNodeKind: valid values are %v", v, AllowedTransformationNodeKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransformationNodeKind) IsValid() bool {
	for _, existing := range AllowedTransformationNodeKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransformationNodeKind value
func (v TransformationNodeKind) Ptr() *TransformationNodeKind {
	return &v
}

type NullableTransformationNodeKind struct {
	value *TransformationNodeKind
	isSet bool
}

func (v NullableTransformationNodeKind) Get() *TransformationNodeKind {
	return v.value
}

func (v *NullableTransformationNodeKind) Set(val *TransformationNodeKind) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationNodeKind) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationNodeKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationNodeKind(val *TransformationNodeKind) *NullableTransformationNodeKind {
	return &NullableTransformationNodeKind{value: val, isSet: true}
}

func (v NullableTransformationNodeKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationNodeKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
