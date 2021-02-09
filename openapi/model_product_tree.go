/*
 * MSRC CVRF API
 *
 * Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_
 *
 * API version: v2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductTree struct for ProductTree
type ProductTree struct {
	Branch          *[]BranchType              `json:"Branch,omitempty"`
	FullProductName *[]FullProductName         `json:"FullProductName,omitempty"`
	Relationship    *[]ProductTreeRelationship `json:"Relationship,omitempty"`
	ProductGroups   *[]ProductTreeGroup        `json:"ProductGroups,omitempty"`
}

// NewProductTree instantiates a new ProductTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTree() *ProductTree {
	this := ProductTree{}
	return &this
}

// NewProductTreeWithDefaults instantiates a new ProductTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTreeWithDefaults() *ProductTree {
	this := ProductTree{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ProductTree) GetBranch() []BranchType {
	if o == nil || o.Branch == nil {
		var ret []BranchType
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTree) GetBranchOk() (*[]BranchType, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ProductTree) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given []BranchType and assigns it to the Branch field.
func (o *ProductTree) SetBranch(v []BranchType) {
	o.Branch = &v
}

// GetFullProductName returns the FullProductName field value if set, zero value otherwise.
func (o *ProductTree) GetFullProductName() []FullProductName {
	if o == nil || o.FullProductName == nil {
		var ret []FullProductName
		return ret
	}
	return *o.FullProductName
}

// GetFullProductNameOk returns a tuple with the FullProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTree) GetFullProductNameOk() (*[]FullProductName, bool) {
	if o == nil || o.FullProductName == nil {
		return nil, false
	}
	return o.FullProductName, true
}

// HasFullProductName returns a boolean if a field has been set.
func (o *ProductTree) HasFullProductName() bool {
	if o != nil && o.FullProductName != nil {
		return true
	}

	return false
}

// SetFullProductName gets a reference to the given []FullProductName and assigns it to the FullProductName field.
func (o *ProductTree) SetFullProductName(v []FullProductName) {
	o.FullProductName = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *ProductTree) GetRelationship() []ProductTreeRelationship {
	if o == nil || o.Relationship == nil {
		var ret []ProductTreeRelationship
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTree) GetRelationshipOk() (*[]ProductTreeRelationship, bool) {
	if o == nil || o.Relationship == nil {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *ProductTree) HasRelationship() bool {
	if o != nil && o.Relationship != nil {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given []ProductTreeRelationship and assigns it to the Relationship field.
func (o *ProductTree) SetRelationship(v []ProductTreeRelationship) {
	o.Relationship = &v
}

// GetProductGroups returns the ProductGroups field value if set, zero value otherwise.
func (o *ProductTree) GetProductGroups() []ProductTreeGroup {
	if o == nil || o.ProductGroups == nil {
		var ret []ProductTreeGroup
		return ret
	}
	return *o.ProductGroups
}

// GetProductGroupsOk returns a tuple with the ProductGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTree) GetProductGroupsOk() (*[]ProductTreeGroup, bool) {
	if o == nil || o.ProductGroups == nil {
		return nil, false
	}
	return o.ProductGroups, true
}

// HasProductGroups returns a boolean if a field has been set.
func (o *ProductTree) HasProductGroups() bool {
	if o != nil && o.ProductGroups != nil {
		return true
	}

	return false
}

// SetProductGroups gets a reference to the given []ProductTreeGroup and assigns it to the ProductGroups field.
func (o *ProductTree) SetProductGroups(v []ProductTreeGroup) {
	o.ProductGroups = &v
}

func (o ProductTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Branch != nil {
		toSerialize["Branch"] = o.Branch
	}
	if o.FullProductName != nil {
		toSerialize["FullProductName"] = o.FullProductName
	}
	if o.Relationship != nil {
		toSerialize["Relationship"] = o.Relationship
	}
	if o.ProductGroups != nil {
		toSerialize["ProductGroups"] = o.ProductGroups
	}
	return json.Marshal(toSerialize)
}

type NullableProductTree struct {
	value *ProductTree
	isSet bool
}

func (v NullableProductTree) Get() *ProductTree {
	return v.value
}

func (v *NullableProductTree) Set(val *ProductTree) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTree) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTree(val *ProductTree) *NullableProductTree {
	return &NullableProductTree{value: val, isSet: true}
}

func (v NullableProductTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
