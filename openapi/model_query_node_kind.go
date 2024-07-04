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

// QueryNodeKind the model 'QueryNodeKind'
type QueryNodeKind string

// List of QueryNodeKind
const (
	NONE                                  QueryNodeKind = "None"
	CONSTANT                              QueryNodeKind = "Constant"
	CONVERT                               QueryNodeKind = "Convert"
	NON_RESOURCE_RANGE_VARIABLE_REFERENCE QueryNodeKind = "NonResourceRangeVariableReference"
	BINARY_OPERATOR                       QueryNodeKind = "BinaryOperator"
	UNARY_OPERATOR                        QueryNodeKind = "UnaryOperator"
	SINGLE_VALUE_PROPERTY_ACCESS          QueryNodeKind = "SingleValuePropertyAccess"
	COLLECTION_PROPERTY_ACCESS            QueryNodeKind = "CollectionPropertyAccess"
	SINGLE_VALUE_FUNCTION_CALL            QueryNodeKind = "SingleValueFunctionCall"
	ANY                                   QueryNodeKind = "Any"
	COLLECTION_NAVIGATION_NODE            QueryNodeKind = "CollectionNavigationNode"
	SINGLE_NAVIGATION_NODE                QueryNodeKind = "SingleNavigationNode"
	SINGLE_VALUE_OPEN_PROPERTY_ACCESS     QueryNodeKind = "SingleValueOpenPropertyAccess"
	SINGLE_RESOURCE_CAST                  QueryNodeKind = "SingleResourceCast"
	ALL                                   QueryNodeKind = "All"
	COLLECTION_RESOURCE_CAST              QueryNodeKind = "CollectionResourceCast"
	RESOURCE_RANGE_VARIABLE_REFERENCE     QueryNodeKind = "ResourceRangeVariableReference"
	SINGLE_RESOURCE_FUNCTION_CALL         QueryNodeKind = "SingleResourceFunctionCall"
	COLLECTION_FUNCTION_CALL              QueryNodeKind = "CollectionFunctionCall"
	COLLECTION_RESOURCE_FUNCTION_CALL     QueryNodeKind = "CollectionResourceFunctionCall"
	NAMED_FUNCTION_PARAMETER              QueryNodeKind = "NamedFunctionParameter"
	PARAMETER_ALIAS                       QueryNodeKind = "ParameterAlias"
	ENTITY_SET                            QueryNodeKind = "EntitySet"
	KEY_LOOKUP                            QueryNodeKind = "KeyLookup"
	SEARCH_TERM                           QueryNodeKind = "SearchTerm"
	COLLECTION_OPEN_PROPERTY_ACCESS       QueryNodeKind = "CollectionOpenPropertyAccess"
	COLLECTION_COMPLEX_NODE               QueryNodeKind = "CollectionComplexNode"
	SINGLE_COMPLEX_NODE                   QueryNodeKind = "SingleComplexNode"
	COUNT                                 QueryNodeKind = "Count"
	SINGLE_VALUE_CAST                     QueryNodeKind = "SingleValueCast"
	COLLECTION_PROPERTY_NODE              QueryNodeKind = "CollectionPropertyNode"
	AGGREGATED_COLLECTION_PROPERTY_NODE   QueryNodeKind = "AggregatedCollectionPropertyNode"
	IN                                    QueryNodeKind = "In"
	COLLECTION_CONSTANT                   QueryNodeKind = "CollectionConstant"
)

// All allowed values of QueryNodeKind enum
var AllowedQueryNodeKindEnumValues = []QueryNodeKind{
	"None",
	"Constant",
	"Convert",
	"NonResourceRangeVariableReference",
	"BinaryOperator",
	"UnaryOperator",
	"SingleValuePropertyAccess",
	"CollectionPropertyAccess",
	"SingleValueFunctionCall",
	"Any",
	"CollectionNavigationNode",
	"SingleNavigationNode",
	"SingleValueOpenPropertyAccess",
	"SingleResourceCast",
	"All",
	"CollectionResourceCast",
	"ResourceRangeVariableReference",
	"SingleResourceFunctionCall",
	"CollectionFunctionCall",
	"CollectionResourceFunctionCall",
	"NamedFunctionParameter",
	"ParameterAlias",
	"EntitySet",
	"KeyLookup",
	"SearchTerm",
	"CollectionOpenPropertyAccess",
	"CollectionComplexNode",
	"SingleComplexNode",
	"Count",
	"SingleValueCast",
	"CollectionPropertyNode",
	"AggregatedCollectionPropertyNode",
	"In",
	"CollectionConstant",
}

func (v *QueryNodeKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryNodeKind(value)
	for _, existing := range AllowedQueryNodeKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryNodeKind", value)
}

// NewQueryNodeKindFromValue returns a pointer to a valid QueryNodeKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryNodeKindFromValue(v string) (*QueryNodeKind, error) {
	ev := QueryNodeKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryNodeKind: valid values are %v", v, AllowedQueryNodeKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryNodeKind) IsValid() bool {
	for _, existing := range AllowedQueryNodeKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryNodeKind value
func (v QueryNodeKind) Ptr() *QueryNodeKind {
	return &v
}

type NullableQueryNodeKind struct {
	value *QueryNodeKind
	isSet bool
}

func (v NullableQueryNodeKind) Get() *QueryNodeKind {
	return v.value
}

func (v *NullableQueryNodeKind) Set(val *QueryNodeKind) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryNodeKind) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryNodeKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryNodeKind(val *QueryNodeKind) *NullableQueryNodeKind {
	return &NullableQueryNodeKind{value: val, isSet: true}
}

func (v NullableQueryNodeKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryNodeKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
