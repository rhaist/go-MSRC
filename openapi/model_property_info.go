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

// checks if the PropertyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyInfo{}

// PropertyInfo struct for PropertyInfo
type PropertyInfo struct {
	Name             NullableString        `json:"name,omitempty"`
	DeclaringType    *Type                 `json:"declaringType,omitempty"`
	ReflectedType    *Type                 `json:"reflectedType,omitempty"`
	Module           *Module               `json:"module,omitempty"`
	CustomAttributes []CustomAttributeData `json:"customAttributes,omitempty"`
	IsCollectible    *bool                 `json:"isCollectible,omitempty"`
	MetadataToken    *int32                `json:"metadataToken,omitempty"`
	MemberType       *MemberTypes          `json:"memberType,omitempty"`
	PropertyType     *Type                 `json:"propertyType,omitempty"`
	Attributes       *PropertyAttributes   `json:"attributes,omitempty"`
	IsSpecialName    *bool                 `json:"isSpecialName,omitempty"`
	CanRead          *bool                 `json:"canRead,omitempty"`
	CanWrite         *bool                 `json:"canWrite,omitempty"`
	GetMethod        *MethodInfo           `json:"getMethod,omitempty"`
	SetMethod        *MethodInfo           `json:"setMethod,omitempty"`
}

// NewPropertyInfo instantiates a new PropertyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyInfo() *PropertyInfo {
	this := PropertyInfo{}
	return &this
}

// NewPropertyInfoWithDefaults instantiates a new PropertyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyInfoWithDefaults() *PropertyInfo {
	this := PropertyInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PropertyInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PropertyInfo) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PropertyInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PropertyInfo) UnsetName() {
	o.Name.Unset()
}

// GetDeclaringType returns the DeclaringType field value if set, zero value otherwise.
func (o *PropertyInfo) GetDeclaringType() Type {
	if o == nil || IsNil(o.DeclaringType) {
		var ret Type
		return ret
	}
	return *o.DeclaringType
}

// GetDeclaringTypeOk returns a tuple with the DeclaringType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetDeclaringTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.DeclaringType) {
		return nil, false
	}
	return o.DeclaringType, true
}

// HasDeclaringType returns a boolean if a field has been set.
func (o *PropertyInfo) HasDeclaringType() bool {
	if o != nil && !IsNil(o.DeclaringType) {
		return true
	}

	return false
}

// SetDeclaringType gets a reference to the given Type and assigns it to the DeclaringType field.
func (o *PropertyInfo) SetDeclaringType(v Type) {
	o.DeclaringType = &v
}

// GetReflectedType returns the ReflectedType field value if set, zero value otherwise.
func (o *PropertyInfo) GetReflectedType() Type {
	if o == nil || IsNil(o.ReflectedType) {
		var ret Type
		return ret
	}
	return *o.ReflectedType
}

// GetReflectedTypeOk returns a tuple with the ReflectedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetReflectedTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.ReflectedType) {
		return nil, false
	}
	return o.ReflectedType, true
}

// HasReflectedType returns a boolean if a field has been set.
func (o *PropertyInfo) HasReflectedType() bool {
	if o != nil && !IsNil(o.ReflectedType) {
		return true
	}

	return false
}

// SetReflectedType gets a reference to the given Type and assigns it to the ReflectedType field.
func (o *PropertyInfo) SetReflectedType(v Type) {
	o.ReflectedType = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *PropertyInfo) GetModule() Module {
	if o == nil || IsNil(o.Module) {
		var ret Module
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetModuleOk() (*Module, bool) {
	if o == nil || IsNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *PropertyInfo) HasModule() bool {
	if o != nil && !IsNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given Module and assigns it to the Module field.
func (o *PropertyInfo) SetModule(v Module) {
	o.Module = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyInfo) GetCustomAttributes() []CustomAttributeData {
	if o == nil {
		var ret []CustomAttributeData
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyInfo) GetCustomAttributesOk() ([]CustomAttributeData, bool) {
	if o == nil || IsNil(o.CustomAttributes) {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *PropertyInfo) HasCustomAttributes() bool {
	if o != nil && !IsNil(o.CustomAttributes) {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []CustomAttributeData and assigns it to the CustomAttributes field.
func (o *PropertyInfo) SetCustomAttributes(v []CustomAttributeData) {
	o.CustomAttributes = v
}

// GetIsCollectible returns the IsCollectible field value if set, zero value otherwise.
func (o *PropertyInfo) GetIsCollectible() bool {
	if o == nil || IsNil(o.IsCollectible) {
		var ret bool
		return ret
	}
	return *o.IsCollectible
}

// GetIsCollectibleOk returns a tuple with the IsCollectible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetIsCollectibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollectible) {
		return nil, false
	}
	return o.IsCollectible, true
}

// HasIsCollectible returns a boolean if a field has been set.
func (o *PropertyInfo) HasIsCollectible() bool {
	if o != nil && !IsNil(o.IsCollectible) {
		return true
	}

	return false
}

// SetIsCollectible gets a reference to the given bool and assigns it to the IsCollectible field.
func (o *PropertyInfo) SetIsCollectible(v bool) {
	o.IsCollectible = &v
}

// GetMetadataToken returns the MetadataToken field value if set, zero value otherwise.
func (o *PropertyInfo) GetMetadataToken() int32 {
	if o == nil || IsNil(o.MetadataToken) {
		var ret int32
		return ret
	}
	return *o.MetadataToken
}

// GetMetadataTokenOk returns a tuple with the MetadataToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetMetadataTokenOk() (*int32, bool) {
	if o == nil || IsNil(o.MetadataToken) {
		return nil, false
	}
	return o.MetadataToken, true
}

// HasMetadataToken returns a boolean if a field has been set.
func (o *PropertyInfo) HasMetadataToken() bool {
	if o != nil && !IsNil(o.MetadataToken) {
		return true
	}

	return false
}

// SetMetadataToken gets a reference to the given int32 and assigns it to the MetadataToken field.
func (o *PropertyInfo) SetMetadataToken(v int32) {
	o.MetadataToken = &v
}

// GetMemberType returns the MemberType field value if set, zero value otherwise.
func (o *PropertyInfo) GetMemberType() MemberTypes {
	if o == nil || IsNil(o.MemberType) {
		var ret MemberTypes
		return ret
	}
	return *o.MemberType
}

// GetMemberTypeOk returns a tuple with the MemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetMemberTypeOk() (*MemberTypes, bool) {
	if o == nil || IsNil(o.MemberType) {
		return nil, false
	}
	return o.MemberType, true
}

// HasMemberType returns a boolean if a field has been set.
func (o *PropertyInfo) HasMemberType() bool {
	if o != nil && !IsNil(o.MemberType) {
		return true
	}

	return false
}

// SetMemberType gets a reference to the given MemberTypes and assigns it to the MemberType field.
func (o *PropertyInfo) SetMemberType(v MemberTypes) {
	o.MemberType = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *PropertyInfo) GetPropertyType() Type {
	if o == nil || IsNil(o.PropertyType) {
		var ret Type
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetPropertyTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *PropertyInfo) HasPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given Type and assigns it to the PropertyType field.
func (o *PropertyInfo) SetPropertyType(v Type) {
	o.PropertyType = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PropertyInfo) GetAttributes() PropertyAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PropertyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetAttributesOk() (*PropertyAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PropertyInfo) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PropertyAttributes and assigns it to the Attributes field.
func (o *PropertyInfo) SetAttributes(v PropertyAttributes) {
	o.Attributes = &v
}

// GetIsSpecialName returns the IsSpecialName field value if set, zero value otherwise.
func (o *PropertyInfo) GetIsSpecialName() bool {
	if o == nil || IsNil(o.IsSpecialName) {
		var ret bool
		return ret
	}
	return *o.IsSpecialName
}

// GetIsSpecialNameOk returns a tuple with the IsSpecialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetIsSpecialNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSpecialName) {
		return nil, false
	}
	return o.IsSpecialName, true
}

// HasIsSpecialName returns a boolean if a field has been set.
func (o *PropertyInfo) HasIsSpecialName() bool {
	if o != nil && !IsNil(o.IsSpecialName) {
		return true
	}

	return false
}

// SetIsSpecialName gets a reference to the given bool and assigns it to the IsSpecialName field.
func (o *PropertyInfo) SetIsSpecialName(v bool) {
	o.IsSpecialName = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *PropertyInfo) GetCanRead() bool {
	if o == nil || IsNil(o.CanRead) {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetCanReadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRead) {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *PropertyInfo) HasCanRead() bool {
	if o != nil && !IsNil(o.CanRead) {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *PropertyInfo) SetCanRead(v bool) {
	o.CanRead = &v
}

// GetCanWrite returns the CanWrite field value if set, zero value otherwise.
func (o *PropertyInfo) GetCanWrite() bool {
	if o == nil || IsNil(o.CanWrite) {
		var ret bool
		return ret
	}
	return *o.CanWrite
}

// GetCanWriteOk returns a tuple with the CanWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetCanWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanWrite) {
		return nil, false
	}
	return o.CanWrite, true
}

// HasCanWrite returns a boolean if a field has been set.
func (o *PropertyInfo) HasCanWrite() bool {
	if o != nil && !IsNil(o.CanWrite) {
		return true
	}

	return false
}

// SetCanWrite gets a reference to the given bool and assigns it to the CanWrite field.
func (o *PropertyInfo) SetCanWrite(v bool) {
	o.CanWrite = &v
}

// GetGetMethod returns the GetMethod field value if set, zero value otherwise.
func (o *PropertyInfo) GetGetMethod() MethodInfo {
	if o == nil || IsNil(o.GetMethod) {
		var ret MethodInfo
		return ret
	}
	return *o.GetMethod
}

// GetGetMethodOk returns a tuple with the GetMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetGetMethodOk() (*MethodInfo, bool) {
	if o == nil || IsNil(o.GetMethod) {
		return nil, false
	}
	return o.GetMethod, true
}

// HasGetMethod returns a boolean if a field has been set.
func (o *PropertyInfo) HasGetMethod() bool {
	if o != nil && !IsNil(o.GetMethod) {
		return true
	}

	return false
}

// SetGetMethod gets a reference to the given MethodInfo and assigns it to the GetMethod field.
func (o *PropertyInfo) SetGetMethod(v MethodInfo) {
	o.GetMethod = &v
}

// GetSetMethod returns the SetMethod field value if set, zero value otherwise.
func (o *PropertyInfo) GetSetMethod() MethodInfo {
	if o == nil || IsNil(o.SetMethod) {
		var ret MethodInfo
		return ret
	}
	return *o.SetMethod
}

// GetSetMethodOk returns a tuple with the SetMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyInfo) GetSetMethodOk() (*MethodInfo, bool) {
	if o == nil || IsNil(o.SetMethod) {
		return nil, false
	}
	return o.SetMethod, true
}

// HasSetMethod returns a boolean if a field has been set.
func (o *PropertyInfo) HasSetMethod() bool {
	if o != nil && !IsNil(o.SetMethod) {
		return true
	}

	return false
}

// SetSetMethod gets a reference to the given MethodInfo and assigns it to the SetMethod field.
func (o *PropertyInfo) SetSetMethod(v MethodInfo) {
	o.SetMethod = &v
}

func (o PropertyInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.MemberType) {
		toSerialize["memberType"] = o.MemberType
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.IsSpecialName) {
		toSerialize["isSpecialName"] = o.IsSpecialName
	}
	if !IsNil(o.CanRead) {
		toSerialize["canRead"] = o.CanRead
	}
	if !IsNil(o.CanWrite) {
		toSerialize["canWrite"] = o.CanWrite
	}
	if !IsNil(o.GetMethod) {
		toSerialize["getMethod"] = o.GetMethod
	}
	if !IsNil(o.SetMethod) {
		toSerialize["setMethod"] = o.SetMethod
	}
	return toSerialize, nil
}

type NullablePropertyInfo struct {
	value *PropertyInfo
	isSet bool
}

func (v NullablePropertyInfo) Get() *PropertyInfo {
	return v.value
}

func (v *NullablePropertyInfo) Set(val *PropertyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyInfo(val *PropertyInfo) *NullablePropertyInfo {
	return &NullablePropertyInfo{value: val, isSet: true}
}

func (v NullablePropertyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}