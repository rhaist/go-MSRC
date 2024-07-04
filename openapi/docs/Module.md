# Module

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assembly** | Pointer to [**Assembly**](Assembly.md) |  | [optional] 
**FullyQualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**MdStreamVersion** | Pointer to **int32** |  | [optional] [readonly] 
**ModuleVersionId** | Pointer to **string** |  | [optional] [readonly] 
**ScopeName** | Pointer to **NullableString** |  | [optional] [readonly] 
**ModuleHandle** | Pointer to [**ModuleHandle**](ModuleHandle.md) |  | [optional] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**MetadataToken** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewModule

`func NewModule() *Module`

NewModule instantiates a new Module object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleWithDefaults

`func NewModuleWithDefaults() *Module`

NewModuleWithDefaults instantiates a new Module object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssembly

`func (o *Module) GetAssembly() Assembly`

GetAssembly returns the Assembly field if non-nil, zero value otherwise.

### GetAssemblyOk

`func (o *Module) GetAssemblyOk() (*Assembly, bool)`

GetAssemblyOk returns a tuple with the Assembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembly

`func (o *Module) SetAssembly(v Assembly)`

SetAssembly sets Assembly field to given value.

### HasAssembly

`func (o *Module) HasAssembly() bool`

HasAssembly returns a boolean if a field has been set.

### GetFullyQualifiedName

`func (o *Module) GetFullyQualifiedName() string`

GetFullyQualifiedName returns the FullyQualifiedName field if non-nil, zero value otherwise.

### GetFullyQualifiedNameOk

`func (o *Module) GetFullyQualifiedNameOk() (*string, bool)`

GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedName

`func (o *Module) SetFullyQualifiedName(v string)`

SetFullyQualifiedName sets FullyQualifiedName field to given value.

### HasFullyQualifiedName

`func (o *Module) HasFullyQualifiedName() bool`

HasFullyQualifiedName returns a boolean if a field has been set.

### SetFullyQualifiedNameNil

`func (o *Module) SetFullyQualifiedNameNil(b bool)`

 SetFullyQualifiedNameNil sets the value for FullyQualifiedName to be an explicit nil

### UnsetFullyQualifiedName
`func (o *Module) UnsetFullyQualifiedName()`

UnsetFullyQualifiedName ensures that no value is present for FullyQualifiedName, not even an explicit nil
### GetName

`func (o *Module) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Module) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Module) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Module) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Module) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Module) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMdStreamVersion

`func (o *Module) GetMdStreamVersion() int32`

GetMdStreamVersion returns the MdStreamVersion field if non-nil, zero value otherwise.

### GetMdStreamVersionOk

`func (o *Module) GetMdStreamVersionOk() (*int32, bool)`

GetMdStreamVersionOk returns a tuple with the MdStreamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdStreamVersion

`func (o *Module) SetMdStreamVersion(v int32)`

SetMdStreamVersion sets MdStreamVersion field to given value.

### HasMdStreamVersion

`func (o *Module) HasMdStreamVersion() bool`

HasMdStreamVersion returns a boolean if a field has been set.

### GetModuleVersionId

`func (o *Module) GetModuleVersionId() string`

GetModuleVersionId returns the ModuleVersionId field if non-nil, zero value otherwise.

### GetModuleVersionIdOk

`func (o *Module) GetModuleVersionIdOk() (*string, bool)`

GetModuleVersionIdOk returns a tuple with the ModuleVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersionId

`func (o *Module) SetModuleVersionId(v string)`

SetModuleVersionId sets ModuleVersionId field to given value.

### HasModuleVersionId

`func (o *Module) HasModuleVersionId() bool`

HasModuleVersionId returns a boolean if a field has been set.

### GetScopeName

`func (o *Module) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *Module) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *Module) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *Module) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### SetScopeNameNil

`func (o *Module) SetScopeNameNil(b bool)`

 SetScopeNameNil sets the value for ScopeName to be an explicit nil

### UnsetScopeName
`func (o *Module) UnsetScopeName()`

UnsetScopeName ensures that no value is present for ScopeName, not even an explicit nil
### GetModuleHandle

`func (o *Module) GetModuleHandle() ModuleHandle`

GetModuleHandle returns the ModuleHandle field if non-nil, zero value otherwise.

### GetModuleHandleOk

`func (o *Module) GetModuleHandleOk() (*ModuleHandle, bool)`

GetModuleHandleOk returns a tuple with the ModuleHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleHandle

`func (o *Module) SetModuleHandle(v ModuleHandle)`

SetModuleHandle sets ModuleHandle field to given value.

### HasModuleHandle

`func (o *Module) HasModuleHandle() bool`

HasModuleHandle returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *Module) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *Module) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *Module) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *Module) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *Module) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *Module) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetMetadataToken

`func (o *Module) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *Module) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *Module) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *Module) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


