# MethodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**DeclaringType** | Pointer to [**Type**](Type.md) |  | [optional] 
**ReflectedType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Module** | Pointer to [**Module**](Module.md) |  | [optional] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**IsCollectible** | Pointer to **bool** |  | [optional] [readonly] 
**MetadataToken** | Pointer to **int32** |  | [optional] [readonly] 
**Attributes** | Pointer to [**MethodAttributes**](MethodAttributes.md) |  | [optional] 
**MethodImplementationFlags** | Pointer to [**MethodImplAttributes**](MethodImplAttributes.md) |  | [optional] 
**CallingConvention** | Pointer to [**CallingConventions**](CallingConventions.md) |  | [optional] 
**IsAbstract** | Pointer to **bool** |  | [optional] [readonly] 
**IsConstructor** | Pointer to **bool** |  | [optional] [readonly] 
**IsFinal** | Pointer to **bool** |  | [optional] [readonly] 
**IsHideBySig** | Pointer to **bool** |  | [optional] [readonly] 
**IsSpecialName** | Pointer to **bool** |  | [optional] [readonly] 
**IsStatic** | Pointer to **bool** |  | [optional] [readonly] 
**IsVirtual** | Pointer to **bool** |  | [optional] [readonly] 
**IsAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamily** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamilyAndAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamilyOrAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsPrivate** | Pointer to **bool** |  | [optional] [readonly] 
**IsPublic** | Pointer to **bool** |  | [optional] [readonly] 
**IsConstructedGenericMethod** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericMethod** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericMethodDefinition** | Pointer to **bool** |  | [optional] [readonly] 
**ContainsGenericParameters** | Pointer to **bool** |  | [optional] [readonly] 
**MethodHandle** | Pointer to [**RuntimeMethodHandle**](RuntimeMethodHandle.md) |  | [optional] 
**IsSecurityCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecuritySafeCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityTransparent** | Pointer to **bool** |  | [optional] [readonly] 
**MemberType** | Pointer to [**MemberTypes**](MemberTypes.md) |  | [optional] 
**ReturnParameter** | Pointer to [**ParameterInfo**](ParameterInfo.md) |  | [optional] 
**ReturnType** | Pointer to [**Type**](Type.md) |  | [optional] 
**ReturnTypeCustomAttributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMethodInfo

`func NewMethodInfo() *MethodInfo`

NewMethodInfo instantiates a new MethodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodInfoWithDefaults

`func NewMethodInfoWithDefaults() *MethodInfo`

NewMethodInfoWithDefaults instantiates a new MethodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MethodInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MethodInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MethodInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MethodInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MethodInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MethodInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeclaringType

`func (o *MethodInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *MethodInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *MethodInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *MethodInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetReflectedType

`func (o *MethodInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *MethodInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *MethodInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *MethodInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetModule

`func (o *MethodInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *MethodInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *MethodInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *MethodInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *MethodInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *MethodInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *MethodInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *MethodInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *MethodInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *MethodInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *MethodInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *MethodInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *MethodInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *MethodInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *MethodInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *MethodInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *MethodInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *MethodInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetAttributes

`func (o *MethodInfo) GetAttributes() MethodAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MethodInfo) GetAttributesOk() (*MethodAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MethodInfo) SetAttributes(v MethodAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MethodInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMethodImplementationFlags

`func (o *MethodInfo) GetMethodImplementationFlags() MethodImplAttributes`

GetMethodImplementationFlags returns the MethodImplementationFlags field if non-nil, zero value otherwise.

### GetMethodImplementationFlagsOk

`func (o *MethodInfo) GetMethodImplementationFlagsOk() (*MethodImplAttributes, bool)`

GetMethodImplementationFlagsOk returns a tuple with the MethodImplementationFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodImplementationFlags

`func (o *MethodInfo) SetMethodImplementationFlags(v MethodImplAttributes)`

SetMethodImplementationFlags sets MethodImplementationFlags field to given value.

### HasMethodImplementationFlags

`func (o *MethodInfo) HasMethodImplementationFlags() bool`

HasMethodImplementationFlags returns a boolean if a field has been set.

### GetCallingConvention

`func (o *MethodInfo) GetCallingConvention() CallingConventions`

GetCallingConvention returns the CallingConvention field if non-nil, zero value otherwise.

### GetCallingConventionOk

`func (o *MethodInfo) GetCallingConventionOk() (*CallingConventions, bool)`

GetCallingConventionOk returns a tuple with the CallingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingConvention

`func (o *MethodInfo) SetCallingConvention(v CallingConventions)`

SetCallingConvention sets CallingConvention field to given value.

### HasCallingConvention

`func (o *MethodInfo) HasCallingConvention() bool`

HasCallingConvention returns a boolean if a field has been set.

### GetIsAbstract

`func (o *MethodInfo) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *MethodInfo) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *MethodInfo) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *MethodInfo) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsConstructor

`func (o *MethodInfo) GetIsConstructor() bool`

GetIsConstructor returns the IsConstructor field if non-nil, zero value otherwise.

### GetIsConstructorOk

`func (o *MethodInfo) GetIsConstructorOk() (*bool, bool)`

GetIsConstructorOk returns a tuple with the IsConstructor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructor

`func (o *MethodInfo) SetIsConstructor(v bool)`

SetIsConstructor sets IsConstructor field to given value.

### HasIsConstructor

`func (o *MethodInfo) HasIsConstructor() bool`

HasIsConstructor returns a boolean if a field has been set.

### GetIsFinal

`func (o *MethodInfo) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *MethodInfo) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *MethodInfo) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *MethodInfo) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetIsHideBySig

`func (o *MethodInfo) GetIsHideBySig() bool`

GetIsHideBySig returns the IsHideBySig field if non-nil, zero value otherwise.

### GetIsHideBySigOk

`func (o *MethodInfo) GetIsHideBySigOk() (*bool, bool)`

GetIsHideBySigOk returns a tuple with the IsHideBySig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHideBySig

`func (o *MethodInfo) SetIsHideBySig(v bool)`

SetIsHideBySig sets IsHideBySig field to given value.

### HasIsHideBySig

`func (o *MethodInfo) HasIsHideBySig() bool`

HasIsHideBySig returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *MethodInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *MethodInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *MethodInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *MethodInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetIsStatic

`func (o *MethodInfo) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *MethodInfo) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *MethodInfo) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *MethodInfo) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetIsVirtual

`func (o *MethodInfo) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *MethodInfo) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *MethodInfo) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *MethodInfo) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsAssembly

`func (o *MethodInfo) GetIsAssembly() bool`

GetIsAssembly returns the IsAssembly field if non-nil, zero value otherwise.

### GetIsAssemblyOk

`func (o *MethodInfo) GetIsAssemblyOk() (*bool, bool)`

GetIsAssemblyOk returns a tuple with the IsAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssembly

`func (o *MethodInfo) SetIsAssembly(v bool)`

SetIsAssembly sets IsAssembly field to given value.

### HasIsAssembly

`func (o *MethodInfo) HasIsAssembly() bool`

HasIsAssembly returns a boolean if a field has been set.

### GetIsFamily

`func (o *MethodInfo) GetIsFamily() bool`

GetIsFamily returns the IsFamily field if non-nil, zero value otherwise.

### GetIsFamilyOk

`func (o *MethodInfo) GetIsFamilyOk() (*bool, bool)`

GetIsFamilyOk returns a tuple with the IsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamily

`func (o *MethodInfo) SetIsFamily(v bool)`

SetIsFamily sets IsFamily field to given value.

### HasIsFamily

`func (o *MethodInfo) HasIsFamily() bool`

HasIsFamily returns a boolean if a field has been set.

### GetIsFamilyAndAssembly

`func (o *MethodInfo) GetIsFamilyAndAssembly() bool`

GetIsFamilyAndAssembly returns the IsFamilyAndAssembly field if non-nil, zero value otherwise.

### GetIsFamilyAndAssemblyOk

`func (o *MethodInfo) GetIsFamilyAndAssemblyOk() (*bool, bool)`

GetIsFamilyAndAssemblyOk returns a tuple with the IsFamilyAndAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyAndAssembly

`func (o *MethodInfo) SetIsFamilyAndAssembly(v bool)`

SetIsFamilyAndAssembly sets IsFamilyAndAssembly field to given value.

### HasIsFamilyAndAssembly

`func (o *MethodInfo) HasIsFamilyAndAssembly() bool`

HasIsFamilyAndAssembly returns a boolean if a field has been set.

### GetIsFamilyOrAssembly

`func (o *MethodInfo) GetIsFamilyOrAssembly() bool`

GetIsFamilyOrAssembly returns the IsFamilyOrAssembly field if non-nil, zero value otherwise.

### GetIsFamilyOrAssemblyOk

`func (o *MethodInfo) GetIsFamilyOrAssemblyOk() (*bool, bool)`

GetIsFamilyOrAssemblyOk returns a tuple with the IsFamilyOrAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyOrAssembly

`func (o *MethodInfo) SetIsFamilyOrAssembly(v bool)`

SetIsFamilyOrAssembly sets IsFamilyOrAssembly field to given value.

### HasIsFamilyOrAssembly

`func (o *MethodInfo) HasIsFamilyOrAssembly() bool`

HasIsFamilyOrAssembly returns a boolean if a field has been set.

### GetIsPrivate

`func (o *MethodInfo) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *MethodInfo) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *MethodInfo) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *MethodInfo) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsPublic

`func (o *MethodInfo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *MethodInfo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *MethodInfo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *MethodInfo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsConstructedGenericMethod

`func (o *MethodInfo) GetIsConstructedGenericMethod() bool`

GetIsConstructedGenericMethod returns the IsConstructedGenericMethod field if non-nil, zero value otherwise.

### GetIsConstructedGenericMethodOk

`func (o *MethodInfo) GetIsConstructedGenericMethodOk() (*bool, bool)`

GetIsConstructedGenericMethodOk returns a tuple with the IsConstructedGenericMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructedGenericMethod

`func (o *MethodInfo) SetIsConstructedGenericMethod(v bool)`

SetIsConstructedGenericMethod sets IsConstructedGenericMethod field to given value.

### HasIsConstructedGenericMethod

`func (o *MethodInfo) HasIsConstructedGenericMethod() bool`

HasIsConstructedGenericMethod returns a boolean if a field has been set.

### GetIsGenericMethod

`func (o *MethodInfo) GetIsGenericMethod() bool`

GetIsGenericMethod returns the IsGenericMethod field if non-nil, zero value otherwise.

### GetIsGenericMethodOk

`func (o *MethodInfo) GetIsGenericMethodOk() (*bool, bool)`

GetIsGenericMethodOk returns a tuple with the IsGenericMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethod

`func (o *MethodInfo) SetIsGenericMethod(v bool)`

SetIsGenericMethod sets IsGenericMethod field to given value.

### HasIsGenericMethod

`func (o *MethodInfo) HasIsGenericMethod() bool`

HasIsGenericMethod returns a boolean if a field has been set.

### GetIsGenericMethodDefinition

`func (o *MethodInfo) GetIsGenericMethodDefinition() bool`

GetIsGenericMethodDefinition returns the IsGenericMethodDefinition field if non-nil, zero value otherwise.

### GetIsGenericMethodDefinitionOk

`func (o *MethodInfo) GetIsGenericMethodDefinitionOk() (*bool, bool)`

GetIsGenericMethodDefinitionOk returns a tuple with the IsGenericMethodDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethodDefinition

`func (o *MethodInfo) SetIsGenericMethodDefinition(v bool)`

SetIsGenericMethodDefinition sets IsGenericMethodDefinition field to given value.

### HasIsGenericMethodDefinition

`func (o *MethodInfo) HasIsGenericMethodDefinition() bool`

HasIsGenericMethodDefinition returns a boolean if a field has been set.

### GetContainsGenericParameters

`func (o *MethodInfo) GetContainsGenericParameters() bool`

GetContainsGenericParameters returns the ContainsGenericParameters field if non-nil, zero value otherwise.

### GetContainsGenericParametersOk

`func (o *MethodInfo) GetContainsGenericParametersOk() (*bool, bool)`

GetContainsGenericParametersOk returns a tuple with the ContainsGenericParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsGenericParameters

`func (o *MethodInfo) SetContainsGenericParameters(v bool)`

SetContainsGenericParameters sets ContainsGenericParameters field to given value.

### HasContainsGenericParameters

`func (o *MethodInfo) HasContainsGenericParameters() bool`

HasContainsGenericParameters returns a boolean if a field has been set.

### GetMethodHandle

`func (o *MethodInfo) GetMethodHandle() RuntimeMethodHandle`

GetMethodHandle returns the MethodHandle field if non-nil, zero value otherwise.

### GetMethodHandleOk

`func (o *MethodInfo) GetMethodHandleOk() (*RuntimeMethodHandle, bool)`

GetMethodHandleOk returns a tuple with the MethodHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodHandle

`func (o *MethodInfo) SetMethodHandle(v RuntimeMethodHandle)`

SetMethodHandle sets MethodHandle field to given value.

### HasMethodHandle

`func (o *MethodInfo) HasMethodHandle() bool`

HasMethodHandle returns a boolean if a field has been set.

### GetIsSecurityCritical

`func (o *MethodInfo) GetIsSecurityCritical() bool`

GetIsSecurityCritical returns the IsSecurityCritical field if non-nil, zero value otherwise.

### GetIsSecurityCriticalOk

`func (o *MethodInfo) GetIsSecurityCriticalOk() (*bool, bool)`

GetIsSecurityCriticalOk returns a tuple with the IsSecurityCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityCritical

`func (o *MethodInfo) SetIsSecurityCritical(v bool)`

SetIsSecurityCritical sets IsSecurityCritical field to given value.

### HasIsSecurityCritical

`func (o *MethodInfo) HasIsSecurityCritical() bool`

HasIsSecurityCritical returns a boolean if a field has been set.

### GetIsSecuritySafeCritical

`func (o *MethodInfo) GetIsSecuritySafeCritical() bool`

GetIsSecuritySafeCritical returns the IsSecuritySafeCritical field if non-nil, zero value otherwise.

### GetIsSecuritySafeCriticalOk

`func (o *MethodInfo) GetIsSecuritySafeCriticalOk() (*bool, bool)`

GetIsSecuritySafeCriticalOk returns a tuple with the IsSecuritySafeCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySafeCritical

`func (o *MethodInfo) SetIsSecuritySafeCritical(v bool)`

SetIsSecuritySafeCritical sets IsSecuritySafeCritical field to given value.

### HasIsSecuritySafeCritical

`func (o *MethodInfo) HasIsSecuritySafeCritical() bool`

HasIsSecuritySafeCritical returns a boolean if a field has been set.

### GetIsSecurityTransparent

`func (o *MethodInfo) GetIsSecurityTransparent() bool`

GetIsSecurityTransparent returns the IsSecurityTransparent field if non-nil, zero value otherwise.

### GetIsSecurityTransparentOk

`func (o *MethodInfo) GetIsSecurityTransparentOk() (*bool, bool)`

GetIsSecurityTransparentOk returns a tuple with the IsSecurityTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityTransparent

`func (o *MethodInfo) SetIsSecurityTransparent(v bool)`

SetIsSecurityTransparent sets IsSecurityTransparent field to given value.

### HasIsSecurityTransparent

`func (o *MethodInfo) HasIsSecurityTransparent() bool`

HasIsSecurityTransparent returns a boolean if a field has been set.

### GetMemberType

`func (o *MethodInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *MethodInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *MethodInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *MethodInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetReturnParameter

`func (o *MethodInfo) GetReturnParameter() ParameterInfo`

GetReturnParameter returns the ReturnParameter field if non-nil, zero value otherwise.

### GetReturnParameterOk

`func (o *MethodInfo) GetReturnParameterOk() (*ParameterInfo, bool)`

GetReturnParameterOk returns a tuple with the ReturnParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnParameter

`func (o *MethodInfo) SetReturnParameter(v ParameterInfo)`

SetReturnParameter sets ReturnParameter field to given value.

### HasReturnParameter

`func (o *MethodInfo) HasReturnParameter() bool`

HasReturnParameter returns a boolean if a field has been set.

### GetReturnType

`func (o *MethodInfo) GetReturnType() Type`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *MethodInfo) GetReturnTypeOk() (*Type, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *MethodInfo) SetReturnType(v Type)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *MethodInfo) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetReturnTypeCustomAttributes

`func (o *MethodInfo) GetReturnTypeCustomAttributes() map[string]interface{}`

GetReturnTypeCustomAttributes returns the ReturnTypeCustomAttributes field if non-nil, zero value otherwise.

### GetReturnTypeCustomAttributesOk

`func (o *MethodInfo) GetReturnTypeCustomAttributesOk() (*map[string]interface{}, bool)`

GetReturnTypeCustomAttributesOk returns a tuple with the ReturnTypeCustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTypeCustomAttributes

`func (o *MethodInfo) SetReturnTypeCustomAttributes(v map[string]interface{})`

SetReturnTypeCustomAttributes sets ReturnTypeCustomAttributes field to given value.

### HasReturnTypeCustomAttributes

`func (o *MethodInfo) HasReturnTypeCustomAttributes() bool`

HasReturnTypeCustomAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


