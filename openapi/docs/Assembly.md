# Assembly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefinedTypes** | Pointer to [**[]TypeInfo**](TypeInfo.md) |  | [optional] [readonly] 
**ExportedTypes** | Pointer to [**[]Type**](Type.md) |  | [optional] [readonly] 
**CodeBase** | Pointer to **NullableString** |  | [optional] [readonly] 
**EntryPoint** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**ImageRuntimeVersion** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsDynamic** | Pointer to **bool** |  | [optional] [readonly] 
**Location** | Pointer to **NullableString** |  | [optional] [readonly] 
**ReflectionOnly** | Pointer to **bool** |  | [optional] [readonly] 
**IsCollectible** | Pointer to **bool** |  | [optional] [readonly] 
**IsFullyTrusted** | Pointer to **bool** |  | [optional] [readonly] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**EscapedCodeBase** | Pointer to **NullableString** |  | [optional] [readonly] 
**ManifestModule** | Pointer to [**Module**](Module.md) |  | [optional] 
**Modules** | Pointer to [**[]Module**](Module.md) |  | [optional] [readonly] 
**GlobalAssemblyCache** | Pointer to **bool** |  | [optional] [readonly] 
**HostContext** | Pointer to **int64** |  | [optional] [readonly] 
**SecurityRuleSet** | Pointer to [**SecurityRuleSet**](SecurityRuleSet.md) |  | [optional] 

## Methods

### NewAssembly

`func NewAssembly() *Assembly`

NewAssembly instantiates a new Assembly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssemblyWithDefaults

`func NewAssemblyWithDefaults() *Assembly`

NewAssemblyWithDefaults instantiates a new Assembly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinedTypes

`func (o *Assembly) GetDefinedTypes() []TypeInfo`

GetDefinedTypes returns the DefinedTypes field if non-nil, zero value otherwise.

### GetDefinedTypesOk

`func (o *Assembly) GetDefinedTypesOk() (*[]TypeInfo, bool)`

GetDefinedTypesOk returns a tuple with the DefinedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedTypes

`func (o *Assembly) SetDefinedTypes(v []TypeInfo)`

SetDefinedTypes sets DefinedTypes field to given value.

### HasDefinedTypes

`func (o *Assembly) HasDefinedTypes() bool`

HasDefinedTypes returns a boolean if a field has been set.

### SetDefinedTypesNil

`func (o *Assembly) SetDefinedTypesNil(b bool)`

 SetDefinedTypesNil sets the value for DefinedTypes to be an explicit nil

### UnsetDefinedTypes
`func (o *Assembly) UnsetDefinedTypes()`

UnsetDefinedTypes ensures that no value is present for DefinedTypes, not even an explicit nil
### GetExportedTypes

`func (o *Assembly) GetExportedTypes() []Type`

GetExportedTypes returns the ExportedTypes field if non-nil, zero value otherwise.

### GetExportedTypesOk

`func (o *Assembly) GetExportedTypesOk() (*[]Type, bool)`

GetExportedTypesOk returns a tuple with the ExportedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedTypes

`func (o *Assembly) SetExportedTypes(v []Type)`

SetExportedTypes sets ExportedTypes field to given value.

### HasExportedTypes

`func (o *Assembly) HasExportedTypes() bool`

HasExportedTypes returns a boolean if a field has been set.

### SetExportedTypesNil

`func (o *Assembly) SetExportedTypesNil(b bool)`

 SetExportedTypesNil sets the value for ExportedTypes to be an explicit nil

### UnsetExportedTypes
`func (o *Assembly) UnsetExportedTypes()`

UnsetExportedTypes ensures that no value is present for ExportedTypes, not even an explicit nil
### GetCodeBase

`func (o *Assembly) GetCodeBase() string`

GetCodeBase returns the CodeBase field if non-nil, zero value otherwise.

### GetCodeBaseOk

`func (o *Assembly) GetCodeBaseOk() (*string, bool)`

GetCodeBaseOk returns a tuple with the CodeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeBase

`func (o *Assembly) SetCodeBase(v string)`

SetCodeBase sets CodeBase field to given value.

### HasCodeBase

`func (o *Assembly) HasCodeBase() bool`

HasCodeBase returns a boolean if a field has been set.

### SetCodeBaseNil

`func (o *Assembly) SetCodeBaseNil(b bool)`

 SetCodeBaseNil sets the value for CodeBase to be an explicit nil

### UnsetCodeBase
`func (o *Assembly) UnsetCodeBase()`

UnsetCodeBase ensures that no value is present for CodeBase, not even an explicit nil
### GetEntryPoint

`func (o *Assembly) GetEntryPoint() MethodInfo`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *Assembly) GetEntryPointOk() (*MethodInfo, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *Assembly) SetEntryPoint(v MethodInfo)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *Assembly) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### GetFullName

`func (o *Assembly) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Assembly) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Assembly) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Assembly) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *Assembly) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *Assembly) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetImageRuntimeVersion

`func (o *Assembly) GetImageRuntimeVersion() string`

GetImageRuntimeVersion returns the ImageRuntimeVersion field if non-nil, zero value otherwise.

### GetImageRuntimeVersionOk

`func (o *Assembly) GetImageRuntimeVersionOk() (*string, bool)`

GetImageRuntimeVersionOk returns a tuple with the ImageRuntimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRuntimeVersion

`func (o *Assembly) SetImageRuntimeVersion(v string)`

SetImageRuntimeVersion sets ImageRuntimeVersion field to given value.

### HasImageRuntimeVersion

`func (o *Assembly) HasImageRuntimeVersion() bool`

HasImageRuntimeVersion returns a boolean if a field has been set.

### SetImageRuntimeVersionNil

`func (o *Assembly) SetImageRuntimeVersionNil(b bool)`

 SetImageRuntimeVersionNil sets the value for ImageRuntimeVersion to be an explicit nil

### UnsetImageRuntimeVersion
`func (o *Assembly) UnsetImageRuntimeVersion()`

UnsetImageRuntimeVersion ensures that no value is present for ImageRuntimeVersion, not even an explicit nil
### GetIsDynamic

`func (o *Assembly) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *Assembly) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *Assembly) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.

### HasIsDynamic

`func (o *Assembly) HasIsDynamic() bool`

HasIsDynamic returns a boolean if a field has been set.

### GetLocation

`func (o *Assembly) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Assembly) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Assembly) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Assembly) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Assembly) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Assembly) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetReflectionOnly

`func (o *Assembly) GetReflectionOnly() bool`

GetReflectionOnly returns the ReflectionOnly field if non-nil, zero value otherwise.

### GetReflectionOnlyOk

`func (o *Assembly) GetReflectionOnlyOk() (*bool, bool)`

GetReflectionOnlyOk returns a tuple with the ReflectionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectionOnly

`func (o *Assembly) SetReflectionOnly(v bool)`

SetReflectionOnly sets ReflectionOnly field to given value.

### HasReflectionOnly

`func (o *Assembly) HasReflectionOnly() bool`

HasReflectionOnly returns a boolean if a field has been set.

### GetIsCollectible

`func (o *Assembly) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *Assembly) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *Assembly) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *Assembly) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetIsFullyTrusted

`func (o *Assembly) GetIsFullyTrusted() bool`

GetIsFullyTrusted returns the IsFullyTrusted field if non-nil, zero value otherwise.

### GetIsFullyTrustedOk

`func (o *Assembly) GetIsFullyTrustedOk() (*bool, bool)`

GetIsFullyTrustedOk returns a tuple with the IsFullyTrusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullyTrusted

`func (o *Assembly) SetIsFullyTrusted(v bool)`

SetIsFullyTrusted sets IsFullyTrusted field to given value.

### HasIsFullyTrusted

`func (o *Assembly) HasIsFullyTrusted() bool`

HasIsFullyTrusted returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *Assembly) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *Assembly) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *Assembly) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *Assembly) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *Assembly) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *Assembly) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetEscapedCodeBase

`func (o *Assembly) GetEscapedCodeBase() string`

GetEscapedCodeBase returns the EscapedCodeBase field if non-nil, zero value otherwise.

### GetEscapedCodeBaseOk

`func (o *Assembly) GetEscapedCodeBaseOk() (*string, bool)`

GetEscapedCodeBaseOk returns a tuple with the EscapedCodeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapedCodeBase

`func (o *Assembly) SetEscapedCodeBase(v string)`

SetEscapedCodeBase sets EscapedCodeBase field to given value.

### HasEscapedCodeBase

`func (o *Assembly) HasEscapedCodeBase() bool`

HasEscapedCodeBase returns a boolean if a field has been set.

### SetEscapedCodeBaseNil

`func (o *Assembly) SetEscapedCodeBaseNil(b bool)`

 SetEscapedCodeBaseNil sets the value for EscapedCodeBase to be an explicit nil

### UnsetEscapedCodeBase
`func (o *Assembly) UnsetEscapedCodeBase()`

UnsetEscapedCodeBase ensures that no value is present for EscapedCodeBase, not even an explicit nil
### GetManifestModule

`func (o *Assembly) GetManifestModule() Module`

GetManifestModule returns the ManifestModule field if non-nil, zero value otherwise.

### GetManifestModuleOk

`func (o *Assembly) GetManifestModuleOk() (*Module, bool)`

GetManifestModuleOk returns a tuple with the ManifestModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestModule

`func (o *Assembly) SetManifestModule(v Module)`

SetManifestModule sets ManifestModule field to given value.

### HasManifestModule

`func (o *Assembly) HasManifestModule() bool`

HasManifestModule returns a boolean if a field has been set.

### GetModules

`func (o *Assembly) GetModules() []Module`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *Assembly) GetModulesOk() (*[]Module, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *Assembly) SetModules(v []Module)`

SetModules sets Modules field to given value.

### HasModules

`func (o *Assembly) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *Assembly) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *Assembly) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil
### GetGlobalAssemblyCache

`func (o *Assembly) GetGlobalAssemblyCache() bool`

GetGlobalAssemblyCache returns the GlobalAssemblyCache field if non-nil, zero value otherwise.

### GetGlobalAssemblyCacheOk

`func (o *Assembly) GetGlobalAssemblyCacheOk() (*bool, bool)`

GetGlobalAssemblyCacheOk returns a tuple with the GlobalAssemblyCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAssemblyCache

`func (o *Assembly) SetGlobalAssemblyCache(v bool)`

SetGlobalAssemblyCache sets GlobalAssemblyCache field to given value.

### HasGlobalAssemblyCache

`func (o *Assembly) HasGlobalAssemblyCache() bool`

HasGlobalAssemblyCache returns a boolean if a field has been set.

### GetHostContext

`func (o *Assembly) GetHostContext() int64`

GetHostContext returns the HostContext field if non-nil, zero value otherwise.

### GetHostContextOk

`func (o *Assembly) GetHostContextOk() (*int64, bool)`

GetHostContextOk returns a tuple with the HostContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostContext

`func (o *Assembly) SetHostContext(v int64)`

SetHostContext sets HostContext field to given value.

### HasHostContext

`func (o *Assembly) HasHostContext() bool`

HasHostContext returns a boolean if a field has been set.

### GetSecurityRuleSet

`func (o *Assembly) GetSecurityRuleSet() SecurityRuleSet`

GetSecurityRuleSet returns the SecurityRuleSet field if non-nil, zero value otherwise.

### GetSecurityRuleSetOk

`func (o *Assembly) GetSecurityRuleSetOk() (*SecurityRuleSet, bool)`

GetSecurityRuleSetOk returns a tuple with the SecurityRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRuleSet

`func (o *Assembly) SetSecurityRuleSet(v SecurityRuleSet)`

SetSecurityRuleSet sets SecurityRuleSet field to given value.

### HasSecurityRuleSet

`func (o *Assembly) HasSecurityRuleSet() bool`

HasSecurityRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


