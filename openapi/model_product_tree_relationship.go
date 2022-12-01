/*
MSRC CVRF API

Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductTreeRelationship struct for ProductTreeRelationship
type ProductTreeRelationship struct {
	FullProductName           []FullProductName `json:"FullProductName,omitempty"`
	ProductReference          *string           `json:"ProductReference,omitempty"`
	RelationType              *int32            `json:"RelationType,omitempty"`
	RelatesToProductReference *string           `json:"RelatesToProductReference,omitempty"`
}

// NewProductTreeRelationship instantiates a new ProductTreeRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTreeRelationship() *ProductTreeRelationship {
	this := ProductTreeRelationship{}
	return &this
}

// NewProductTreeRelationshipWithDefaults instantiates a new ProductTreeRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTreeRelationshipWithDefaults() *ProductTreeRelationship {
	this := ProductTreeRelationship{}
	return &this
}

// GetFullProductName returns the FullProductName field value if set, zero value otherwise.
func (o *ProductTreeRelationship) GetFullProductName() []FullProductName {
	if o == nil || isNil(o.FullProductName) {
		var ret []FullProductName
		return ret
	}
	return o.FullProductName
}

// GetFullProductNameOk returns a tuple with the FullProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeRelationship) GetFullProductNameOk() ([]FullProductName, bool) {
	if o == nil || isNil(o.FullProductName) {
		return nil, false
	}
	return o.FullProductName, true
}

// HasFullProductName returns a boolean if a field has been set.
func (o *ProductTreeRelationship) HasFullProductName() bool {
	if o != nil && !isNil(o.FullProductName) {
		return true
	}

	return false
}

// SetFullProductName gets a reference to the given []FullProductName and assigns it to the FullProductName field.
func (o *ProductTreeRelationship) SetFullProductName(v []FullProductName) {
	o.FullProductName = v
}

// GetProductReference returns the ProductReference field value if set, zero value otherwise.
func (o *ProductTreeRelationship) GetProductReference() string {
	if o == nil || isNil(o.ProductReference) {
		var ret string
		return ret
	}
	return *o.ProductReference
}

// GetProductReferenceOk returns a tuple with the ProductReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeRelationship) GetProductReferenceOk() (*string, bool) {
	if o == nil || isNil(o.ProductReference) {
		return nil, false
	}
	return o.ProductReference, true
}

// HasProductReference returns a boolean if a field has been set.
func (o *ProductTreeRelationship) HasProductReference() bool {
	if o != nil && !isNil(o.ProductReference) {
		return true
	}

	return false
}

// SetProductReference gets a reference to the given string and assigns it to the ProductReference field.
func (o *ProductTreeRelationship) SetProductReference(v string) {
	o.ProductReference = &v
}

// GetRelationType returns the RelationType field value if set, zero value otherwise.
func (o *ProductTreeRelationship) GetRelationType() int32 {
	if o == nil || isNil(o.RelationType) {
		var ret int32
		return ret
	}
	return *o.RelationType
}

// GetRelationTypeOk returns a tuple with the RelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeRelationship) GetRelationTypeOk() (*int32, bool) {
	if o == nil || isNil(o.RelationType) {
		return nil, false
	}
	return o.RelationType, true
}

// HasRelationType returns a boolean if a field has been set.
func (o *ProductTreeRelationship) HasRelationType() bool {
	if o != nil && !isNil(o.RelationType) {
		return true
	}

	return false
}

// SetRelationType gets a reference to the given int32 and assigns it to the RelationType field.
func (o *ProductTreeRelationship) SetRelationType(v int32) {
	o.RelationType = &v
}

// GetRelatesToProductReference returns the RelatesToProductReference field value if set, zero value otherwise.
func (o *ProductTreeRelationship) GetRelatesToProductReference() string {
	if o == nil || isNil(o.RelatesToProductReference) {
		var ret string
		return ret
	}
	return *o.RelatesToProductReference
}

// GetRelatesToProductReferenceOk returns a tuple with the RelatesToProductReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTreeRelationship) GetRelatesToProductReferenceOk() (*string, bool) {
	if o == nil || isNil(o.RelatesToProductReference) {
		return nil, false
	}
	return o.RelatesToProductReference, true
}

// HasRelatesToProductReference returns a boolean if a field has been set.
func (o *ProductTreeRelationship) HasRelatesToProductReference() bool {
	if o != nil && !isNil(o.RelatesToProductReference) {
		return true
	}

	return false
}

// SetRelatesToProductReference gets a reference to the given string and assigns it to the RelatesToProductReference field.
func (o *ProductTreeRelationship) SetRelatesToProductReference(v string) {
	o.RelatesToProductReference = &v
}

func (o ProductTreeRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FullProductName) {
		toSerialize["FullProductName"] = o.FullProductName
	}
	if !isNil(o.ProductReference) {
		toSerialize["ProductReference"] = o.ProductReference
	}
	if !isNil(o.RelationType) {
		toSerialize["RelationType"] = o.RelationType
	}
	if !isNil(o.RelatesToProductReference) {
		toSerialize["RelatesToProductReference"] = o.RelatesToProductReference
	}
	return json.Marshal(toSerialize)
}

type NullableProductTreeRelationship struct {
	value *ProductTreeRelationship
	isSet bool
}

func (v NullableProductTreeRelationship) Get() *ProductTreeRelationship {
	return v.value
}

func (v *NullableProductTreeRelationship) Set(val *ProductTreeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTreeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTreeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTreeRelationship(val *ProductTreeRelationship) *NullableProductTreeRelationship {
	return &NullableProductTreeRelationship{value: val, isSet: true}
}

func (v NullableProductTreeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTreeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
