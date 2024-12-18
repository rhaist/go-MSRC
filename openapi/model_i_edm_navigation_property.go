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

// checks if the IEdmNavigationProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmNavigationProperty{}

// IEdmNavigationProperty struct for IEdmNavigationProperty
type IEdmNavigationProperty struct {
	Partner               *IEdmNavigationProperty    `json:"partner,omitempty"`
	OnDelete              *EdmOnDeleteAction         `json:"onDelete,omitempty"`
	ContainsTarget        *bool                      `json:"containsTarget,omitempty"`
	ReferentialConstraint *IEdmReferentialConstraint `json:"referentialConstraint,omitempty"`
	PropertyKind          *EdmPropertyKind           `json:"propertyKind,omitempty"`
	Type                  *IEdmTypeReference         `json:"type,omitempty"`
	DeclaringType         *IEdmStructuredType        `json:"declaringType,omitempty"`
	Name                  NullableString             `json:"name,omitempty"`
}

// NewIEdmNavigationProperty instantiates a new IEdmNavigationProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmNavigationProperty() *IEdmNavigationProperty {
	this := IEdmNavigationProperty{}
	return &this
}

// NewIEdmNavigationPropertyWithDefaults instantiates a new IEdmNavigationProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmNavigationPropertyWithDefaults() *IEdmNavigationProperty {
	this := IEdmNavigationProperty{}
	return &this
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetPartner() IEdmNavigationProperty {
	if o == nil || IsNil(o.Partner) {
		var ret IEdmNavigationProperty
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetPartnerOk() (*IEdmNavigationProperty, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given IEdmNavigationProperty and assigns it to the Partner field.
func (o *IEdmNavigationProperty) SetPartner(v IEdmNavigationProperty) {
	o.Partner = &v
}

// GetOnDelete returns the OnDelete field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetOnDelete() EdmOnDeleteAction {
	if o == nil || IsNil(o.OnDelete) {
		var ret EdmOnDeleteAction
		return ret
	}
	return *o.OnDelete
}

// GetOnDeleteOk returns a tuple with the OnDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetOnDeleteOk() (*EdmOnDeleteAction, bool) {
	if o == nil || IsNil(o.OnDelete) {
		return nil, false
	}
	return o.OnDelete, true
}

// HasOnDelete returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasOnDelete() bool {
	if o != nil && !IsNil(o.OnDelete) {
		return true
	}

	return false
}

// SetOnDelete gets a reference to the given EdmOnDeleteAction and assigns it to the OnDelete field.
func (o *IEdmNavigationProperty) SetOnDelete(v EdmOnDeleteAction) {
	o.OnDelete = &v
}

// GetContainsTarget returns the ContainsTarget field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetContainsTarget() bool {
	if o == nil || IsNil(o.ContainsTarget) {
		var ret bool
		return ret
	}
	return *o.ContainsTarget
}

// GetContainsTargetOk returns a tuple with the ContainsTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetContainsTargetOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainsTarget) {
		return nil, false
	}
	return o.ContainsTarget, true
}

// HasContainsTarget returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasContainsTarget() bool {
	if o != nil && !IsNil(o.ContainsTarget) {
		return true
	}

	return false
}

// SetContainsTarget gets a reference to the given bool and assigns it to the ContainsTarget field.
func (o *IEdmNavigationProperty) SetContainsTarget(v bool) {
	o.ContainsTarget = &v
}

// GetReferentialConstraint returns the ReferentialConstraint field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetReferentialConstraint() IEdmReferentialConstraint {
	if o == nil || IsNil(o.ReferentialConstraint) {
		var ret IEdmReferentialConstraint
		return ret
	}
	return *o.ReferentialConstraint
}

// GetReferentialConstraintOk returns a tuple with the ReferentialConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetReferentialConstraintOk() (*IEdmReferentialConstraint, bool) {
	if o == nil || IsNil(o.ReferentialConstraint) {
		return nil, false
	}
	return o.ReferentialConstraint, true
}

// HasReferentialConstraint returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasReferentialConstraint() bool {
	if o != nil && !IsNil(o.ReferentialConstraint) {
		return true
	}

	return false
}

// SetReferentialConstraint gets a reference to the given IEdmReferentialConstraint and assigns it to the ReferentialConstraint field.
func (o *IEdmNavigationProperty) SetReferentialConstraint(v IEdmReferentialConstraint) {
	o.ReferentialConstraint = &v
}

// GetPropertyKind returns the PropertyKind field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetPropertyKind() EdmPropertyKind {
	if o == nil || IsNil(o.PropertyKind) {
		var ret EdmPropertyKind
		return ret
	}
	return *o.PropertyKind
}

// GetPropertyKindOk returns a tuple with the PropertyKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetPropertyKindOk() (*EdmPropertyKind, bool) {
	if o == nil || IsNil(o.PropertyKind) {
		return nil, false
	}
	return o.PropertyKind, true
}

// HasPropertyKind returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasPropertyKind() bool {
	if o != nil && !IsNil(o.PropertyKind) {
		return true
	}

	return false
}

// SetPropertyKind gets a reference to the given EdmPropertyKind and assigns it to the PropertyKind field.
func (o *IEdmNavigationProperty) SetPropertyKind(v EdmPropertyKind) {
	o.PropertyKind = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetType() IEdmTypeReference {
	if o == nil || IsNil(o.Type) {
		var ret IEdmTypeReference
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetTypeOk() (*IEdmTypeReference, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IEdmTypeReference and assigns it to the Type field.
func (o *IEdmNavigationProperty) SetType(v IEdmTypeReference) {
	o.Type = &v
}

// GetDeclaringType returns the DeclaringType field value if set, zero value otherwise.
func (o *IEdmNavigationProperty) GetDeclaringType() IEdmStructuredType {
	if o == nil || IsNil(o.DeclaringType) {
		var ret IEdmStructuredType
		return ret
	}
	return *o.DeclaringType
}

// GetDeclaringTypeOk returns a tuple with the DeclaringType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmNavigationProperty) GetDeclaringTypeOk() (*IEdmStructuredType, bool) {
	if o == nil || IsNil(o.DeclaringType) {
		return nil, false
	}
	return o.DeclaringType, true
}

// HasDeclaringType returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasDeclaringType() bool {
	if o != nil && !IsNil(o.DeclaringType) {
		return true
	}

	return false
}

// SetDeclaringType gets a reference to the given IEdmStructuredType and assigns it to the DeclaringType field.
func (o *IEdmNavigationProperty) SetDeclaringType(v IEdmStructuredType) {
	o.DeclaringType = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmNavigationProperty) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmNavigationProperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IEdmNavigationProperty) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IEdmNavigationProperty) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *IEdmNavigationProperty) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IEdmNavigationProperty) UnsetName() {
	o.Name.Unset()
}

func (o IEdmNavigationProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmNavigationProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.OnDelete) {
		toSerialize["onDelete"] = o.OnDelete
	}
	if !IsNil(o.ContainsTarget) {
		toSerialize["containsTarget"] = o.ContainsTarget
	}
	if !IsNil(o.ReferentialConstraint) {
		toSerialize["referentialConstraint"] = o.ReferentialConstraint
	}
	if !IsNil(o.PropertyKind) {
		toSerialize["propertyKind"] = o.PropertyKind
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DeclaringType) {
		toSerialize["declaringType"] = o.DeclaringType
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableIEdmNavigationProperty struct {
	value *IEdmNavigationProperty
	isSet bool
}

func (v NullableIEdmNavigationProperty) Get() *IEdmNavigationProperty {
	return v.value
}

func (v *NullableIEdmNavigationProperty) Set(val *IEdmNavigationProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmNavigationProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmNavigationProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmNavigationProperty(val *IEdmNavigationProperty) *NullableIEdmNavigationProperty {
	return &NullableIEdmNavigationProperty{value: val, isSet: true}
}

func (v NullableIEdmNavigationProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmNavigationProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
