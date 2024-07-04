# ConstructorInfo

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

## Methods

### NewConstructorInfo

`func NewConstructorInfo() *ConstructorInfo`

NewConstructorInfo instantiates a new ConstructorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructorInfoWithDefaults

`func NewConstructorInfoWithDefaults() *ConstructorInfo`

NewConstructorInfoWithDefaults instantiates a new ConstructorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConstructorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConstructorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConstructorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConstructorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConstructorInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConstructorInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeclaringType

`func (o *ConstructorInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *ConstructorInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *ConstructorInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *ConstructorInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetReflectedType

`func (o *ConstructorInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *ConstructorInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *ConstructorInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *ConstructorInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetModule

`func (o *ConstructorInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ConstructorInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ConstructorInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *ConstructorInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ConstructorInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ConstructorInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ConstructorInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ConstructorInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *ConstructorInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *ConstructorInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *ConstructorInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *ConstructorInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *ConstructorInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *ConstructorInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *ConstructorInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *ConstructorInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *ConstructorInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *ConstructorInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetAttributes

`func (o *ConstructorInfo) GetAttributes() MethodAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConstructorInfo) GetAttributesOk() (*MethodAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConstructorInfo) SetAttributes(v MethodAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConstructorInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMethodImplementationFlags

`func (o *ConstructorInfo) GetMethodImplementationFlags() MethodImplAttributes`

GetMethodImplementationFlags returns the MethodImplementationFlags field if non-nil, zero value otherwise.

### GetMethodImplementationFlagsOk

`func (o *ConstructorInfo) GetMethodImplementationFlagsOk() (*MethodImplAttributes, bool)`

GetMethodImplementationFlagsOk returns a tuple with the MethodImplementationFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodImplementationFlags

`func (o *ConstructorInfo) SetMethodImplementationFlags(v MethodImplAttributes)`

SetMethodImplementationFlags sets MethodImplementationFlags field to given value.

### HasMethodImplementationFlags

`func (o *ConstructorInfo) HasMethodImplementationFlags() bool`

HasMethodImplementationFlags returns a boolean if a field has been set.

### GetCallingConvention

`func (o *ConstructorInfo) GetCallingConvention() CallingConventions`

GetCallingConvention returns the CallingConvention field if non-nil, zero value otherwise.

### GetCallingConventionOk

`func (o *ConstructorInfo) GetCallingConventionOk() (*CallingConventions, bool)`

GetCallingConventionOk returns a tuple with the CallingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingConvention

`func (o *ConstructorInfo) SetCallingConvention(v CallingConventions)`

SetCallingConvention sets CallingConvention field to given value.

### HasCallingConvention

`func (o *ConstructorInfo) HasCallingConvention() bool`

HasCallingConvention returns a boolean if a field has been set.

### GetIsAbstract

`func (o *ConstructorInfo) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *ConstructorInfo) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *ConstructorInfo) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *ConstructorInfo) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsConstructor

`func (o *ConstructorInfo) GetIsConstructor() bool`

GetIsConstructor returns the IsConstructor field if non-nil, zero value otherwise.

### GetIsConstructorOk

`func (o *ConstructorInfo) GetIsConstructorOk() (*bool, bool)`

GetIsConstructorOk returns a tuple with the IsConstructor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructor

`func (o *ConstructorInfo) SetIsConstructor(v bool)`

SetIsConstructor sets IsConstructor field to given value.

### HasIsConstructor

`func (o *ConstructorInfo) HasIsConstructor() bool`

HasIsConstructor returns a boolean if a field has been set.

### GetIsFinal

`func (o *ConstructorInfo) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *ConstructorInfo) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *ConstructorInfo) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *ConstructorInfo) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetIsHideBySig

`func (o *ConstructorInfo) GetIsHideBySig() bool`

GetIsHideBySig returns the IsHideBySig field if non-nil, zero value otherwise.

### GetIsHideBySigOk

`func (o *ConstructorInfo) GetIsHideBySigOk() (*bool, bool)`

GetIsHideBySigOk returns a tuple with the IsHideBySig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHideBySig

`func (o *ConstructorInfo) SetIsHideBySig(v bool)`

SetIsHideBySig sets IsHideBySig field to given value.

### HasIsHideBySig

`func (o *ConstructorInfo) HasIsHideBySig() bool`

HasIsHideBySig returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *ConstructorInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *ConstructorInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *ConstructorInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *ConstructorInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetIsStatic

`func (o *ConstructorInfo) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *ConstructorInfo) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *ConstructorInfo) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *ConstructorInfo) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ConstructorInfo) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ConstructorInfo) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ConstructorInfo) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ConstructorInfo) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsAssembly

`func (o *ConstructorInfo) GetIsAssembly() bool`

GetIsAssembly returns the IsAssembly field if non-nil, zero value otherwise.

### GetIsAssemblyOk

`func (o *ConstructorInfo) GetIsAssemblyOk() (*bool, bool)`

GetIsAssemblyOk returns a tuple with the IsAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssembly

`func (o *ConstructorInfo) SetIsAssembly(v bool)`

SetIsAssembly sets IsAssembly field to given value.

### HasIsAssembly

`func (o *ConstructorInfo) HasIsAssembly() bool`

HasIsAssembly returns a boolean if a field has been set.

### GetIsFamily

`func (o *ConstructorInfo) GetIsFamily() bool`

GetIsFamily returns the IsFamily field if non-nil, zero value otherwise.

### GetIsFamilyOk

`func (o *ConstructorInfo) GetIsFamilyOk() (*bool, bool)`

GetIsFamilyOk returns a tuple with the IsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamily

`func (o *ConstructorInfo) SetIsFamily(v bool)`

SetIsFamily sets IsFamily field to given value.

### HasIsFamily

`func (o *ConstructorInfo) HasIsFamily() bool`

HasIsFamily returns a boolean if a field has been set.

### GetIsFamilyAndAssembly

`func (o *ConstructorInfo) GetIsFamilyAndAssembly() bool`

GetIsFamilyAndAssembly returns the IsFamilyAndAssembly field if non-nil, zero value otherwise.

### GetIsFamilyAndAssemblyOk

`func (o *ConstructorInfo) GetIsFamilyAndAssemblyOk() (*bool, bool)`

GetIsFamilyAndAssemblyOk returns a tuple with the IsFamilyAndAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyAndAssembly

`func (o *ConstructorInfo) SetIsFamilyAndAssembly(v bool)`

SetIsFamilyAndAssembly sets IsFamilyAndAssembly field to given value.

### HasIsFamilyAndAssembly

`func (o *ConstructorInfo) HasIsFamilyAndAssembly() bool`

HasIsFamilyAndAssembly returns a boolean if a field has been set.

### GetIsFamilyOrAssembly

`func (o *ConstructorInfo) GetIsFamilyOrAssembly() bool`

GetIsFamilyOrAssembly returns the IsFamilyOrAssembly field if non-nil, zero value otherwise.

### GetIsFamilyOrAssemblyOk

`func (o *ConstructorInfo) GetIsFamilyOrAssemblyOk() (*bool, bool)`

GetIsFamilyOrAssemblyOk returns a tuple with the IsFamilyOrAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyOrAssembly

`func (o *ConstructorInfo) SetIsFamilyOrAssembly(v bool)`

SetIsFamilyOrAssembly sets IsFamilyOrAssembly field to given value.

### HasIsFamilyOrAssembly

`func (o *ConstructorInfo) HasIsFamilyOrAssembly() bool`

HasIsFamilyOrAssembly returns a boolean if a field has been set.

### GetIsPrivate

`func (o *ConstructorInfo) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *ConstructorInfo) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *ConstructorInfo) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *ConstructorInfo) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsPublic

`func (o *ConstructorInfo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ConstructorInfo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ConstructorInfo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ConstructorInfo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsConstructedGenericMethod

`func (o *ConstructorInfo) GetIsConstructedGenericMethod() bool`

GetIsConstructedGenericMethod returns the IsConstructedGenericMethod field if non-nil, zero value otherwise.

### GetIsConstructedGenericMethodOk

`func (o *ConstructorInfo) GetIsConstructedGenericMethodOk() (*bool, bool)`

GetIsConstructedGenericMethodOk returns a tuple with the IsConstructedGenericMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructedGenericMethod

`func (o *ConstructorInfo) SetIsConstructedGenericMethod(v bool)`

SetIsConstructedGenericMethod sets IsConstructedGenericMethod field to given value.

### HasIsConstructedGenericMethod

`func (o *ConstructorInfo) HasIsConstructedGenericMethod() bool`

HasIsConstructedGenericMethod returns a boolean if a field has been set.

### GetIsGenericMethod

`func (o *ConstructorInfo) GetIsGenericMethod() bool`

GetIsGenericMethod returns the IsGenericMethod field if non-nil, zero value otherwise.

### GetIsGenericMethodOk

`func (o *ConstructorInfo) GetIsGenericMethodOk() (*bool, bool)`

GetIsGenericMethodOk returns a tuple with the IsGenericMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethod

`func (o *ConstructorInfo) SetIsGenericMethod(v bool)`

SetIsGenericMethod sets IsGenericMethod field to given value.

### HasIsGenericMethod

`func (o *ConstructorInfo) HasIsGenericMethod() bool`

HasIsGenericMethod returns a boolean if a field has been set.

### GetIsGenericMethodDefinition

`func (o *ConstructorInfo) GetIsGenericMethodDefinition() bool`

GetIsGenericMethodDefinition returns the IsGenericMethodDefinition field if non-nil, zero value otherwise.

### GetIsGenericMethodDefinitionOk

`func (o *ConstructorInfo) GetIsGenericMethodDefinitionOk() (*bool, bool)`

GetIsGenericMethodDefinitionOk returns a tuple with the IsGenericMethodDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethodDefinition

`func (o *ConstructorInfo) SetIsGenericMethodDefinition(v bool)`

SetIsGenericMethodDefinition sets IsGenericMethodDefinition field to given value.

### HasIsGenericMethodDefinition

`func (o *ConstructorInfo) HasIsGenericMethodDefinition() bool`

HasIsGenericMethodDefinition returns a boolean if a field has been set.

### GetContainsGenericParameters

`func (o *ConstructorInfo) GetContainsGenericParameters() bool`

GetContainsGenericParameters returns the ContainsGenericParameters field if non-nil, zero value otherwise.

### GetContainsGenericParametersOk

`func (o *ConstructorInfo) GetContainsGenericParametersOk() (*bool, bool)`

GetContainsGenericParametersOk returns a tuple with the ContainsGenericParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsGenericParameters

`func (o *ConstructorInfo) SetContainsGenericParameters(v bool)`

SetContainsGenericParameters sets ContainsGenericParameters field to given value.

### HasContainsGenericParameters

`func (o *ConstructorInfo) HasContainsGenericParameters() bool`

HasContainsGenericParameters returns a boolean if a field has been set.

### GetMethodHandle

`func (o *ConstructorInfo) GetMethodHandle() RuntimeMethodHandle`

GetMethodHandle returns the MethodHandle field if non-nil, zero value otherwise.

### GetMethodHandleOk

`func (o *ConstructorInfo) GetMethodHandleOk() (*RuntimeMethodHandle, bool)`

GetMethodHandleOk returns a tuple with the MethodHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodHandle

`func (o *ConstructorInfo) SetMethodHandle(v RuntimeMethodHandle)`

SetMethodHandle sets MethodHandle field to given value.

### HasMethodHandle

`func (o *ConstructorInfo) HasMethodHandle() bool`

HasMethodHandle returns a boolean if a field has been set.

### GetIsSecurityCritical

`func (o *ConstructorInfo) GetIsSecurityCritical() bool`

GetIsSecurityCritical returns the IsSecurityCritical field if non-nil, zero value otherwise.

### GetIsSecurityCriticalOk

`func (o *ConstructorInfo) GetIsSecurityCriticalOk() (*bool, bool)`

GetIsSecurityCriticalOk returns a tuple with the IsSecurityCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityCritical

`func (o *ConstructorInfo) SetIsSecurityCritical(v bool)`

SetIsSecurityCritical sets IsSecurityCritical field to given value.

### HasIsSecurityCritical

`func (o *ConstructorInfo) HasIsSecurityCritical() bool`

HasIsSecurityCritical returns a boolean if a field has been set.

### GetIsSecuritySafeCritical

`func (o *ConstructorInfo) GetIsSecuritySafeCritical() bool`

GetIsSecuritySafeCritical returns the IsSecuritySafeCritical field if non-nil, zero value otherwise.

### GetIsSecuritySafeCriticalOk

`func (o *ConstructorInfo) GetIsSecuritySafeCriticalOk() (*bool, bool)`

GetIsSecuritySafeCriticalOk returns a tuple with the IsSecuritySafeCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySafeCritical

`func (o *ConstructorInfo) SetIsSecuritySafeCritical(v bool)`

SetIsSecuritySafeCritical sets IsSecuritySafeCritical field to given value.

### HasIsSecuritySafeCritical

`func (o *ConstructorInfo) HasIsSecuritySafeCritical() bool`

HasIsSecuritySafeCritical returns a boolean if a field has been set.

### GetIsSecurityTransparent

`func (o *ConstructorInfo) GetIsSecurityTransparent() bool`

GetIsSecurityTransparent returns the IsSecurityTransparent field if non-nil, zero value otherwise.

### GetIsSecurityTransparentOk

`func (o *ConstructorInfo) GetIsSecurityTransparentOk() (*bool, bool)`

GetIsSecurityTransparentOk returns a tuple with the IsSecurityTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityTransparent

`func (o *ConstructorInfo) SetIsSecurityTransparent(v bool)`

SetIsSecurityTransparent sets IsSecurityTransparent field to given value.

### HasIsSecurityTransparent

`func (o *ConstructorInfo) HasIsSecurityTransparent() bool`

HasIsSecurityTransparent returns a boolean if a field has been set.

### GetMemberType

`func (o *ConstructorInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *ConstructorInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *ConstructorInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *ConstructorInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


