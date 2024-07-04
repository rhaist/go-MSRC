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

// checks if the Assembly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Assembly{}

// Assembly struct for Assembly
type Assembly struct {
	DefinedTypes        []TypeInfo            `json:"definedTypes,omitempty"`
	ExportedTypes       []Type                `json:"exportedTypes,omitempty"`
	CodeBase            NullableString        `json:"codeBase,omitempty"`
	EntryPoint          *MethodInfo           `json:"entryPoint,omitempty"`
	FullName            NullableString        `json:"fullName,omitempty"`
	ImageRuntimeVersion NullableString        `json:"imageRuntimeVersion,omitempty"`
	IsDynamic           *bool                 `json:"isDynamic,omitempty"`
	Location            NullableString        `json:"location,omitempty"`
	ReflectionOnly      *bool                 `json:"reflectionOnly,omitempty"`
	IsCollectible       *bool                 `json:"isCollectible,omitempty"`
	IsFullyTrusted      *bool                 `json:"isFullyTrusted,omitempty"`
	CustomAttributes    []CustomAttributeData `json:"customAttributes,omitempty"`
	EscapedCodeBase     NullableString        `json:"escapedCodeBase,omitempty"`
	ManifestModule      *Module               `json:"manifestModule,omitempty"`
	Modules             []Module              `json:"modules,omitempty"`
	// Deprecated
	GlobalAssemblyCache *bool            `json:"globalAssemblyCache,omitempty"`
	HostContext         *int64           `json:"hostContext,omitempty"`
	SecurityRuleSet     *SecurityRuleSet `json:"securityRuleSet,omitempty"`
}

// NewAssembly instantiates a new Assembly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssembly() *Assembly {
	this := Assembly{}
	return &this
}

// NewAssemblyWithDefaults instantiates a new Assembly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssemblyWithDefaults() *Assembly {
	this := Assembly{}
	return &this
}

// GetDefinedTypes returns the DefinedTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetDefinedTypes() []TypeInfo {
	if o == nil {
		var ret []TypeInfo
		return ret
	}
	return o.DefinedTypes
}

// GetDefinedTypesOk returns a tuple with the DefinedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetDefinedTypesOk() ([]TypeInfo, bool) {
	if o == nil || IsNil(o.DefinedTypes) {
		return nil, false
	}
	return o.DefinedTypes, true
}

// HasDefinedTypes returns a boolean if a field has been set.
func (o *Assembly) HasDefinedTypes() bool {
	if o != nil && !IsNil(o.DefinedTypes) {
		return true
	}

	return false
}

// SetDefinedTypes gets a reference to the given []TypeInfo and assigns it to the DefinedTypes field.
func (o *Assembly) SetDefinedTypes(v []TypeInfo) {
	o.DefinedTypes = v
}

// GetExportedTypes returns the ExportedTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetExportedTypes() []Type {
	if o == nil {
		var ret []Type
		return ret
	}
	return o.ExportedTypes
}

// GetExportedTypesOk returns a tuple with the ExportedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetExportedTypesOk() ([]Type, bool) {
	if o == nil || IsNil(o.ExportedTypes) {
		return nil, false
	}
	return o.ExportedTypes, true
}

// HasExportedTypes returns a boolean if a field has been set.
func (o *Assembly) HasExportedTypes() bool {
	if o != nil && !IsNil(o.ExportedTypes) {
		return true
	}

	return false
}

// SetExportedTypes gets a reference to the given []Type and assigns it to the ExportedTypes field.
func (o *Assembly) SetExportedTypes(v []Type) {
	o.ExportedTypes = v
}

// GetCodeBase returns the CodeBase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetCodeBase() string {
	if o == nil || IsNil(o.CodeBase.Get()) {
		var ret string
		return ret
	}
	return *o.CodeBase.Get()
}

// GetCodeBaseOk returns a tuple with the CodeBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetCodeBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeBase.Get(), o.CodeBase.IsSet()
}

// HasCodeBase returns a boolean if a field has been set.
func (o *Assembly) HasCodeBase() bool {
	if o != nil && o.CodeBase.IsSet() {
		return true
	}

	return false
}

// SetCodeBase gets a reference to the given NullableString and assigns it to the CodeBase field.
func (o *Assembly) SetCodeBase(v string) {
	o.CodeBase.Set(&v)
}

// SetCodeBaseNil sets the value for CodeBase to be an explicit nil
func (o *Assembly) SetCodeBaseNil() {
	o.CodeBase.Set(nil)
}

// UnsetCodeBase ensures that no value is present for CodeBase, not even an explicit nil
func (o *Assembly) UnsetCodeBase() {
	o.CodeBase.Unset()
}

// GetEntryPoint returns the EntryPoint field value if set, zero value otherwise.
func (o *Assembly) GetEntryPoint() MethodInfo {
	if o == nil || IsNil(o.EntryPoint) {
		var ret MethodInfo
		return ret
	}
	return *o.EntryPoint
}

// GetEntryPointOk returns a tuple with the EntryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetEntryPointOk() (*MethodInfo, bool) {
	if o == nil || IsNil(o.EntryPoint) {
		return nil, false
	}
	return o.EntryPoint, true
}

// HasEntryPoint returns a boolean if a field has been set.
func (o *Assembly) HasEntryPoint() bool {
	if o != nil && !IsNil(o.EntryPoint) {
		return true
	}

	return false
}

// SetEntryPoint gets a reference to the given MethodInfo and assigns it to the EntryPoint field.
func (o *Assembly) SetEntryPoint(v MethodInfo) {
	o.EntryPoint = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetFullName() string {
	if o == nil || IsNil(o.FullName.Get()) {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *Assembly) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *Assembly) SetFullName(v string) {
	o.FullName.Set(&v)
}

// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *Assembly) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *Assembly) UnsetFullName() {
	o.FullName.Unset()
}

// GetImageRuntimeVersion returns the ImageRuntimeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetImageRuntimeVersion() string {
	if o == nil || IsNil(o.ImageRuntimeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ImageRuntimeVersion.Get()
}

// GetImageRuntimeVersionOk returns a tuple with the ImageRuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetImageRuntimeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageRuntimeVersion.Get(), o.ImageRuntimeVersion.IsSet()
}

// HasImageRuntimeVersion returns a boolean if a field has been set.
func (o *Assembly) HasImageRuntimeVersion() bool {
	if o != nil && o.ImageRuntimeVersion.IsSet() {
		return true
	}

	return false
}

// SetImageRuntimeVersion gets a reference to the given NullableString and assigns it to the ImageRuntimeVersion field.
func (o *Assembly) SetImageRuntimeVersion(v string) {
	o.ImageRuntimeVersion.Set(&v)
}

// SetImageRuntimeVersionNil sets the value for ImageRuntimeVersion to be an explicit nil
func (o *Assembly) SetImageRuntimeVersionNil() {
	o.ImageRuntimeVersion.Set(nil)
}

// UnsetImageRuntimeVersion ensures that no value is present for ImageRuntimeVersion, not even an explicit nil
func (o *Assembly) UnsetImageRuntimeVersion() {
	o.ImageRuntimeVersion.Unset()
}

// GetIsDynamic returns the IsDynamic field value if set, zero value otherwise.
func (o *Assembly) GetIsDynamic() bool {
	if o == nil || IsNil(o.IsDynamic) {
		var ret bool
		return ret
	}
	return *o.IsDynamic
}

// GetIsDynamicOk returns a tuple with the IsDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetIsDynamicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDynamic) {
		return nil, false
	}
	return o.IsDynamic, true
}

// HasIsDynamic returns a boolean if a field has been set.
func (o *Assembly) HasIsDynamic() bool {
	if o != nil && !IsNil(o.IsDynamic) {
		return true
	}

	return false
}

// SetIsDynamic gets a reference to the given bool and assigns it to the IsDynamic field.
func (o *Assembly) SetIsDynamic(v bool) {
	o.IsDynamic = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Assembly) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *Assembly) SetLocation(v string) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *Assembly) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Assembly) UnsetLocation() {
	o.Location.Unset()
}

// GetReflectionOnly returns the ReflectionOnly field value if set, zero value otherwise.
func (o *Assembly) GetReflectionOnly() bool {
	if o == nil || IsNil(o.ReflectionOnly) {
		var ret bool
		return ret
	}
	return *o.ReflectionOnly
}

// GetReflectionOnlyOk returns a tuple with the ReflectionOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetReflectionOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReflectionOnly) {
		return nil, false
	}
	return o.ReflectionOnly, true
}

// HasReflectionOnly returns a boolean if a field has been set.
func (o *Assembly) HasReflectionOnly() bool {
	if o != nil && !IsNil(o.ReflectionOnly) {
		return true
	}

	return false
}

// SetReflectionOnly gets a reference to the given bool and assigns it to the ReflectionOnly field.
func (o *Assembly) SetReflectionOnly(v bool) {
	o.ReflectionOnly = &v
}

// GetIsCollectible returns the IsCollectible field value if set, zero value otherwise.
func (o *Assembly) GetIsCollectible() bool {
	if o == nil || IsNil(o.IsCollectible) {
		var ret bool
		return ret
	}
	return *o.IsCollectible
}

// GetIsCollectibleOk returns a tuple with the IsCollectible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetIsCollectibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollectible) {
		return nil, false
	}
	return o.IsCollectible, true
}

// HasIsCollectible returns a boolean if a field has been set.
func (o *Assembly) HasIsCollectible() bool {
	if o != nil && !IsNil(o.IsCollectible) {
		return true
	}

	return false
}

// SetIsCollectible gets a reference to the given bool and assigns it to the IsCollectible field.
func (o *Assembly) SetIsCollectible(v bool) {
	o.IsCollectible = &v
}

// GetIsFullyTrusted returns the IsFullyTrusted field value if set, zero value otherwise.
func (o *Assembly) GetIsFullyTrusted() bool {
	if o == nil || IsNil(o.IsFullyTrusted) {
		var ret bool
		return ret
	}
	return *o.IsFullyTrusted
}

// GetIsFullyTrustedOk returns a tuple with the IsFullyTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetIsFullyTrustedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullyTrusted) {
		return nil, false
	}
	return o.IsFullyTrusted, true
}

// HasIsFullyTrusted returns a boolean if a field has been set.
func (o *Assembly) HasIsFullyTrusted() bool {
	if o != nil && !IsNil(o.IsFullyTrusted) {
		return true
	}

	return false
}

// SetIsFullyTrusted gets a reference to the given bool and assigns it to the IsFullyTrusted field.
func (o *Assembly) SetIsFullyTrusted(v bool) {
	o.IsFullyTrusted = &v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetCustomAttributes() []CustomAttributeData {
	if o == nil {
		var ret []CustomAttributeData
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetCustomAttributesOk() ([]CustomAttributeData, bool) {
	if o == nil || IsNil(o.CustomAttributes) {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *Assembly) HasCustomAttributes() bool {
	if o != nil && !IsNil(o.CustomAttributes) {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []CustomAttributeData and assigns it to the CustomAttributes field.
func (o *Assembly) SetCustomAttributes(v []CustomAttributeData) {
	o.CustomAttributes = v
}

// GetEscapedCodeBase returns the EscapedCodeBase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetEscapedCodeBase() string {
	if o == nil || IsNil(o.EscapedCodeBase.Get()) {
		var ret string
		return ret
	}
	return *o.EscapedCodeBase.Get()
}

// GetEscapedCodeBaseOk returns a tuple with the EscapedCodeBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetEscapedCodeBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EscapedCodeBase.Get(), o.EscapedCodeBase.IsSet()
}

// HasEscapedCodeBase returns a boolean if a field has been set.
func (o *Assembly) HasEscapedCodeBase() bool {
	if o != nil && o.EscapedCodeBase.IsSet() {
		return true
	}

	return false
}

// SetEscapedCodeBase gets a reference to the given NullableString and assigns it to the EscapedCodeBase field.
func (o *Assembly) SetEscapedCodeBase(v string) {
	o.EscapedCodeBase.Set(&v)
}

// SetEscapedCodeBaseNil sets the value for EscapedCodeBase to be an explicit nil
func (o *Assembly) SetEscapedCodeBaseNil() {
	o.EscapedCodeBase.Set(nil)
}

// UnsetEscapedCodeBase ensures that no value is present for EscapedCodeBase, not even an explicit nil
func (o *Assembly) UnsetEscapedCodeBase() {
	o.EscapedCodeBase.Unset()
}

// GetManifestModule returns the ManifestModule field value if set, zero value otherwise.
func (o *Assembly) GetManifestModule() Module {
	if o == nil || IsNil(o.ManifestModule) {
		var ret Module
		return ret
	}
	return *o.ManifestModule
}

// GetManifestModuleOk returns a tuple with the ManifestModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetManifestModuleOk() (*Module, bool) {
	if o == nil || IsNil(o.ManifestModule) {
		return nil, false
	}
	return o.ManifestModule, true
}

// HasManifestModule returns a boolean if a field has been set.
func (o *Assembly) HasManifestModule() bool {
	if o != nil && !IsNil(o.ManifestModule) {
		return true
	}

	return false
}

// SetManifestModule gets a reference to the given Module and assigns it to the ManifestModule field.
func (o *Assembly) SetManifestModule(v Module) {
	o.ManifestModule = &v
}

// GetModules returns the Modules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Assembly) GetModules() []Module {
	if o == nil {
		var ret []Module
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Assembly) GetModulesOk() ([]Module, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *Assembly) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given []Module and assigns it to the Modules field.
func (o *Assembly) SetModules(v []Module) {
	o.Modules = v
}

// GetGlobalAssemblyCache returns the GlobalAssemblyCache field value if set, zero value otherwise.
// Deprecated
func (o *Assembly) GetGlobalAssemblyCache() bool {
	if o == nil || IsNil(o.GlobalAssemblyCache) {
		var ret bool
		return ret
	}
	return *o.GlobalAssemblyCache
}

// GetGlobalAssemblyCacheOk returns a tuple with the GlobalAssemblyCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Assembly) GetGlobalAssemblyCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalAssemblyCache) {
		return nil, false
	}
	return o.GlobalAssemblyCache, true
}

// HasGlobalAssemblyCache returns a boolean if a field has been set.
func (o *Assembly) HasGlobalAssemblyCache() bool {
	if o != nil && !IsNil(o.GlobalAssemblyCache) {
		return true
	}

	return false
}

// SetGlobalAssemblyCache gets a reference to the given bool and assigns it to the GlobalAssemblyCache field.
// Deprecated
func (o *Assembly) SetGlobalAssemblyCache(v bool) {
	o.GlobalAssemblyCache = &v
}

// GetHostContext returns the HostContext field value if set, zero value otherwise.
func (o *Assembly) GetHostContext() int64 {
	if o == nil || IsNil(o.HostContext) {
		var ret int64
		return ret
	}
	return *o.HostContext
}

// GetHostContextOk returns a tuple with the HostContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetHostContextOk() (*int64, bool) {
	if o == nil || IsNil(o.HostContext) {
		return nil, false
	}
	return o.HostContext, true
}

// HasHostContext returns a boolean if a field has been set.
func (o *Assembly) HasHostContext() bool {
	if o != nil && !IsNil(o.HostContext) {
		return true
	}

	return false
}

// SetHostContext gets a reference to the given int64 and assigns it to the HostContext field.
func (o *Assembly) SetHostContext(v int64) {
	o.HostContext = &v
}

// GetSecurityRuleSet returns the SecurityRuleSet field value if set, zero value otherwise.
func (o *Assembly) GetSecurityRuleSet() SecurityRuleSet {
	if o == nil || IsNil(o.SecurityRuleSet) {
		var ret SecurityRuleSet
		return ret
	}
	return *o.SecurityRuleSet
}

// GetSecurityRuleSetOk returns a tuple with the SecurityRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Assembly) GetSecurityRuleSetOk() (*SecurityRuleSet, bool) {
	if o == nil || IsNil(o.SecurityRuleSet) {
		return nil, false
	}
	return o.SecurityRuleSet, true
}

// HasSecurityRuleSet returns a boolean if a field has been set.
func (o *Assembly) HasSecurityRuleSet() bool {
	if o != nil && !IsNil(o.SecurityRuleSet) {
		return true
	}

	return false
}

// SetSecurityRuleSet gets a reference to the given SecurityRuleSet and assigns it to the SecurityRuleSet field.
func (o *Assembly) SetSecurityRuleSet(v SecurityRuleSet) {
	o.SecurityRuleSet = &v
}

func (o Assembly) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Assembly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefinedTypes != nil {
		toSerialize["definedTypes"] = o.DefinedTypes
	}
	if o.ExportedTypes != nil {
		toSerialize["exportedTypes"] = o.ExportedTypes
	}
	if o.CodeBase.IsSet() {
		toSerialize["codeBase"] = o.CodeBase.Get()
	}
	if !IsNil(o.EntryPoint) {
		toSerialize["entryPoint"] = o.EntryPoint
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.ImageRuntimeVersion.IsSet() {
		toSerialize["imageRuntimeVersion"] = o.ImageRuntimeVersion.Get()
	}
	if !IsNil(o.IsDynamic) {
		toSerialize["isDynamic"] = o.IsDynamic
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if !IsNil(o.ReflectionOnly) {
		toSerialize["reflectionOnly"] = o.ReflectionOnly
	}
	if !IsNil(o.IsCollectible) {
		toSerialize["isCollectible"] = o.IsCollectible
	}
	if !IsNil(o.IsFullyTrusted) {
		toSerialize["isFullyTrusted"] = o.IsFullyTrusted
	}
	if o.CustomAttributes != nil {
		toSerialize["customAttributes"] = o.CustomAttributes
	}
	if o.EscapedCodeBase.IsSet() {
		toSerialize["escapedCodeBase"] = o.EscapedCodeBase.Get()
	}
	if !IsNil(o.ManifestModule) {
		toSerialize["manifestModule"] = o.ManifestModule
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.GlobalAssemblyCache) {
		toSerialize["globalAssemblyCache"] = o.GlobalAssemblyCache
	}
	if !IsNil(o.HostContext) {
		toSerialize["hostContext"] = o.HostContext
	}
	if !IsNil(o.SecurityRuleSet) {
		toSerialize["securityRuleSet"] = o.SecurityRuleSet
	}
	return toSerialize, nil
}

type NullableAssembly struct {
	value *Assembly
	isSet bool
}

func (v NullableAssembly) Get() *Assembly {
	return v.value
}

func (v *NullableAssembly) Set(val *Assembly) {
	v.value = val
	v.isSet = true
}

func (v NullableAssembly) IsSet() bool {
	return v.isSet
}

func (v *NullableAssembly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssembly(val *Assembly) *NullableAssembly {
	return &NullableAssembly{value: val, isSet: true}
}

func (v NullableAssembly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssembly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}