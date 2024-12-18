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

// checks if the MemberInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberInfo{}

// MemberInfo struct for MemberInfo
type MemberInfo struct {
	MemberType       *MemberTypes          `json:"memberType,omitempty"`
	Name             NullableString        `json:"name,omitempty"`
	DeclaringType    *Type                 `json:"declaringType,omitempty"`
	ReflectedType    *Type                 `json:"reflectedType,omitempty"`
	Module           *Module               `json:"module,omitempty"`
	CustomAttributes []CustomAttributeData `json:"customAttributes,omitempty"`
	IsCollectible    *bool                 `json:"isCollectible,omitempty"`
	MetadataToken    *int32                `json:"metadataToken,omitempty"`
}

// NewMemberInfo instantiates a new MemberInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberInfo() *MemberInfo {
	this := MemberInfo{}
	return &this
}

// NewMemberInfoWithDefaults instantiates a new MemberInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberInfoWithDefaults() *MemberInfo {
	this := MemberInfo{}
	return &this
}

// GetMemberType returns the MemberType field value if set, zero value otherwise.
func (o *MemberInfo) GetMemberType() MemberTypes {
	if o == nil || IsNil(o.MemberType) {
		var ret MemberTypes
		return ret
	}
	return *o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetMemberTypeOk() (*MemberTypes, bool) {
	if o == nil || IsNil(o.MemberType) {
		return nil, false
	}
	return o.MemberType, true
}

// HasMemberType returns a boolean if a field has been set.
func (o *MemberInfo) HasMemberType() bool {
	if o != nil && !IsNil(o.MemberType) {
		return true
	}

	return false
}

// SetMemberType gets a reference to the given MemberTypes and assigns it to the MemberType field.
func (o *MemberInfo) SetMemberType(v MemberTypes) {
	o.MemberType = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MemberInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MemberInfo) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *MemberInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MemberInfo) UnsetName() {
	o.Name.Unset()
}

// GetDeclaringType returns the DeclaringType field value if set, zero value otherwise.
func (o *MemberInfo) GetDeclaringType() Type {
	if o == nil || IsNil(o.DeclaringType) {
		var ret Type
		return ret
	}
	return *o.DeclaringType
}

// GetDeclaringTypeOk returns a tuple with the DeclaringType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetDeclaringTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.DeclaringType) {
		return nil, false
	}
	return o.DeclaringType, true
}

// HasDeclaringType returns a boolean if a field has been set.
func (o *MemberInfo) HasDeclaringType() bool {
	if o != nil && !IsNil(o.DeclaringType) {
		return true
	}

	return false
}

// SetDeclaringType gets a reference to the given Type and assigns it to the DeclaringType field.
func (o *MemberInfo) SetDeclaringType(v Type) {
	o.DeclaringType = &v
}

// GetReflectedType returns the ReflectedType field value if set, zero value otherwise.
func (o *MemberInfo) GetReflectedType() Type {
	if o == nil || IsNil(o.ReflectedType) {
		var ret Type
		return ret
	}
	return *o.ReflectedType
}

// GetReflectedTypeOk returns a tuple with the ReflectedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetReflectedTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.ReflectedType) {
		return nil, false
	}
	return o.ReflectedType, true
}

// HasReflectedType returns a boolean if a field has been set.
func (o *MemberInfo) HasReflectedType() bool {
	if o != nil && !IsNil(o.ReflectedType) {
		return true
	}

	return false
}

// SetReflectedType gets a reference to the given Type and assigns it to the ReflectedType field.
func (o *MemberInfo) SetReflectedType(v Type) {
	o.ReflectedType = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *MemberInfo) GetModule() Module {
	if o == nil || IsNil(o.Module) {
		var ret Module
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetModuleOk() (*Module, bool) {
	if o == nil || IsNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *MemberInfo) HasModule() bool {
	if o != nil && !IsNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given Module and assigns it to the Module field.
func (o *MemberInfo) SetModule(v Module) {
	o.Module = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberInfo) GetCustomAttributes() []CustomAttributeData {
	if o == nil {
		var ret []CustomAttributeData
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberInfo) GetCustomAttributesOk() ([]CustomAttributeData, bool) {
	if o == nil || IsNil(o.CustomAttributes) {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *MemberInfo) HasCustomAttributes() bool {
	if o != nil && !IsNil(o.CustomAttributes) {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []CustomAttributeData and assigns it to the CustomAttributes field.
func (o *MemberInfo) SetCustomAttributes(v []CustomAttributeData) {
	o.CustomAttributes = v
}

// GetIsCollectible returns the IsCollectible field value if set, zero value otherwise.
func (o *MemberInfo) GetIsCollectible() bool {
	if o == nil || IsNil(o.IsCollectible) {
		var ret bool
		return ret
	}
	return *o.IsCollectible
}

// GetIsCollectibleOk returns a tuple with the IsCollectible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetIsCollectibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollectible) {
		return nil, false
	}
	return o.IsCollectible, true
}

// HasIsCollectible returns a boolean if a field has been set.
func (o *MemberInfo) HasIsCollectible() bool {
	if o != nil && !IsNil(o.IsCollectible) {
		return true
	}

	return false
}

// SetIsCollectible gets a reference to the given bool and assigns it to the IsCollectible field.
func (o *MemberInfo) SetIsCollectible(v bool) {
	o.IsCollectible = &v
}

// GetMetadataToken returns the MetadataToken field value if set, zero value otherwise.
func (o *MemberInfo) GetMetadataToken() int32 {
	if o == nil || IsNil(o.MetadataToken) {
		var ret int32
		return ret
	}
	return *o.MetadataToken
}

// GetMetadataTokenOk returns a tuple with the MetadataToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberInfo) GetMetadataTokenOk() (*int32, bool) {
	if o == nil || IsNil(o.MetadataToken) {
		return nil, false
	}
	return o.MetadataToken, true
}

// HasMetadataToken returns a boolean if a field has been set.
func (o *MemberInfo) HasMetadataToken() bool {
	if o != nil && !IsNil(o.MetadataToken) {
		return true
	}

	return false
}

// SetMetadataToken gets a reference to the given int32 and assigns it to the MetadataToken field.
func (o *MemberInfo) SetMetadataToken(v int32) {
	o.MetadataToken = &v
}

func (o MemberInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberType) {
		toSerialize["memberType"] = o.MemberType
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.DeclaringType) {
		toSerialize["declaringType"] = o.DeclaringType
	}
	if !IsNil(o.ReflectedType) {
		toSerialize["reflectedType"] = o.ReflectedType
	}
	if !IsNil(o.Module) {
		toSerialize["module"] = o.Module
	}
	if o.CustomAttributes != nil {
		toSerialize["customAttributes"] = o.CustomAttributes
	}
	if !IsNil(o.IsCollectible) {
		toSerialize["isCollectible"] = o.IsCollectible
	}
	if !IsNil(o.MetadataToken) {
		toSerialize["metadataToken"] = o.MetadataToken
	}
	return toSerialize, nil
}

type NullableMemberInfo struct {
	value *MemberInfo
	isSet bool
}

func (v NullableMemberInfo) Get() *MemberInfo {
	return v.value
}

func (v *NullableMemberInfo) Set(val *MemberInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberInfo(val *MemberInfo) *NullableMemberInfo {
	return &NullableMemberInfo{value: val, isSet: true}
}

func (v NullableMemberInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
