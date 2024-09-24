# Type

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**IsCollectible** | Pointer to **bool** |  | [optional] [readonly] 
**MetadataToken** | Pointer to **int32** |  | [optional] [readonly] 
**IsInterface** | Pointer to **bool** |  | [optional] [readonly] 
**MemberType** | Pointer to [**MemberTypes**](MemberTypes.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] [readonly] 
**AssemblyQualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**Assembly** | Pointer to [**Assembly**](Assembly.md) |  | [optional] 
**Module** | Pointer to [**Module**](Module.md) |  | [optional] 
**IsNested** | Pointer to **bool** |  | [optional] [readonly] 
**DeclaringType** | Pointer to [**Type**](Type.md) |  | [optional] 
**DeclaringMethod** | Pointer to [**MethodBase**](MethodBase.md) |  | [optional] 
**ReflectedType** | Pointer to [**Type**](Type.md) |  | [optional] 
**UnderlyingSystemType** | Pointer to [**Type**](Type.md) |  | [optional] 
**IsTypeDefinition** | Pointer to **bool** |  | [optional] [readonly] 
**IsArray** | Pointer to **bool** |  | [optional] [readonly] 
**IsByRef** | Pointer to **bool** |  | [optional] [readonly] 
**IsPointer** | Pointer to **bool** |  | [optional] [readonly] 
**IsConstructedGenericType** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericParameter** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericTypeParameter** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericMethodParameter** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericType** | Pointer to **bool** |  | [optional] [readonly] 
**IsGenericTypeDefinition** | Pointer to **bool** |  | [optional] [readonly] 
**IsSZArray** | Pointer to **bool** |  | [optional] [readonly] 
**IsVariableBoundArray** | Pointer to **bool** |  | [optional] [readonly] 
**IsByRefLike** | Pointer to **bool** |  | [optional] [readonly] 
**IsFunctionPointer** | Pointer to **bool** |  | [optional] [readonly] 
**IsUnmanagedFunctionPointer** | Pointer to **bool** |  | [optional] [readonly] 
**HasElementType** | Pointer to **bool** |  | [optional] [readonly] 
**GenericTypeArguments** | Pointer to [**[]Type**](Type.md) |  | [optional] [readonly] 
**GenericParameterPosition** | Pointer to **int32** |  | [optional] [readonly] 
**GenericParameterAttributes** | Pointer to [**GenericParameterAttributes**](GenericParameterAttributes.md) |  | [optional] 
**Attributes** | Pointer to [**TypeAttributes**](TypeAttributes.md) |  | [optional] 
**IsAbstract** | Pointer to **bool** |  | [optional] [readonly] 
**IsImport** | Pointer to **bool** |  | [optional] [readonly] 
**IsSealed** | Pointer to **bool** |  | [optional] [readonly] 
**IsSpecialName** | Pointer to **bool** |  | [optional] [readonly] 
**IsClass** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedFamANDAssem** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedFamily** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedFamORAssem** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedPrivate** | Pointer to **bool** |  | [optional] [readonly] 
**IsNestedPublic** | Pointer to **bool** |  | [optional] [readonly] 
**IsNotPublic** | Pointer to **bool** |  | [optional] [readonly] 
**IsPublic** | Pointer to **bool** |  | [optional] [readonly] 
**IsAutoLayout** | Pointer to **bool** |  | [optional] [readonly] 
**IsExplicitLayout** | Pointer to **bool** |  | [optional] [readonly] 
**IsLayoutSequential** | Pointer to **bool** |  | [optional] [readonly] 
**IsAnsiClass** | Pointer to **bool** |  | [optional] [readonly] 
**IsAutoClass** | Pointer to **bool** |  | [optional] [readonly] 
**IsUnicodeClass** | Pointer to **bool** |  | [optional] [readonly] 
**IsCOMObject** | Pointer to **bool** |  | [optional] [readonly] 
**IsContextful** | Pointer to **bool** |  | [optional] [readonly] 
**IsEnum** | Pointer to **bool** |  | [optional] [readonly] 
**IsMarshalByRef** | Pointer to **bool** |  | [optional] [readonly] 
**IsPrimitive** | Pointer to **bool** |  | [optional] [readonly] 
**IsValueType** | Pointer to **bool** |  | [optional] [readonly] 
**IsSignatureType** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecuritySafeCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityTransparent** | Pointer to **bool** |  | [optional] [readonly] 
**StructLayoutAttribute** | Pointer to [**StructLayoutAttribute**](StructLayoutAttribute.md) |  | [optional] 
**TypeInitializer** | Pointer to [**ConstructorInfo**](ConstructorInfo.md) |  | [optional] 
**TypeHandle** | Pointer to [**RuntimeTypeHandle**](RuntimeTypeHandle.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] [readonly] 
**BaseType** | Pointer to [**Type**](Type.md) |  | [optional] 
**IsSerializable** | Pointer to **bool** |  | [optional] [readonly] 
**ContainsGenericParameters** | Pointer to **bool** |  | [optional] [readonly] 
**IsVisible** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewType

`func NewType() *Type`

NewType instantiates a new Type object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeWithDefaults

`func NewTypeWithDefaults() *Type`

NewTypeWithDefaults instantiates a new Type object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Type) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Type) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Type) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Type) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Type) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Type) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomAttributes

`func (o *Type) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *Type) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *Type) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *Type) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *Type) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *Type) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *Type) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *Type) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *Type) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *Type) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *Type) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *Type) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *Type) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *Type) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetIsInterface

`func (o *Type) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *Type) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *Type) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *Type) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetMemberType

`func (o *Type) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *Type) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *Type) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *Type) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetNamespace

`func (o *Type) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Type) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Type) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Type) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Type) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Type) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetAssemblyQualifiedName

`func (o *Type) GetAssemblyQualifiedName() string`

GetAssemblyQualifiedName returns the AssemblyQualifiedName field if non-nil, zero value otherwise.

### GetAssemblyQualifiedNameOk

`func (o *Type) GetAssemblyQualifiedNameOk() (*string, bool)`

GetAssemblyQualifiedNameOk returns a tuple with the AssemblyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyQualifiedName

`func (o *Type) SetAssemblyQualifiedName(v string)`

SetAssemblyQualifiedName sets AssemblyQualifiedName field to given value.

### HasAssemblyQualifiedName

`func (o *Type) HasAssemblyQualifiedName() bool`

HasAssemblyQualifiedName returns a boolean if a field has been set.

### SetAssemblyQualifiedNameNil

`func (o *Type) SetAssemblyQualifiedNameNil(b bool)`

 SetAssemblyQualifiedNameNil sets the value for AssemblyQualifiedName to be an explicit nil

### UnsetAssemblyQualifiedName
`func (o *Type) UnsetAssemblyQualifiedName()`

UnsetAssemblyQualifiedName ensures that no value is present for AssemblyQualifiedName, not even an explicit nil
### GetFullName

`func (o *Type) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *Type) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *Type) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *Type) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *Type) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *Type) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetAssembly

`func (o *Type) GetAssembly() Assembly`

GetAssembly returns the Assembly field if non-nil, zero value otherwise.

### GetAssemblyOk

`func (o *Type) GetAssemblyOk() (*Assembly, bool)`

GetAssemblyOk returns a tuple with the Assembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembly

`func (o *Type) SetAssembly(v Assembly)`

SetAssembly sets Assembly field to given value.

### HasAssembly

`func (o *Type) HasAssembly() bool`

HasAssembly returns a boolean if a field has been set.

### GetModule

`func (o *Type) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Type) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Type) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *Type) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetIsNested

`func (o *Type) GetIsNested() bool`

GetIsNested returns the IsNested field if non-nil, zero value otherwise.

### GetIsNestedOk

`func (o *Type) GetIsNestedOk() (*bool, bool)`

GetIsNestedOk returns a tuple with the IsNested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNested

`func (o *Type) SetIsNested(v bool)`

SetIsNested sets IsNested field to given value.

### HasIsNested

`func (o *Type) HasIsNested() bool`

HasIsNested returns a boolean if a field has been set.

### GetDeclaringType

`func (o *Type) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *Type) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *Type) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *Type) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetDeclaringMethod

`func (o *Type) GetDeclaringMethod() MethodBase`

GetDeclaringMethod returns the DeclaringMethod field if non-nil, zero value otherwise.

### GetDeclaringMethodOk

`func (o *Type) GetDeclaringMethodOk() (*MethodBase, bool)`

GetDeclaringMethodOk returns a tuple with the DeclaringMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringMethod

`func (o *Type) SetDeclaringMethod(v MethodBase)`

SetDeclaringMethod sets DeclaringMethod field to given value.

### HasDeclaringMethod

`func (o *Type) HasDeclaringMethod() bool`

HasDeclaringMethod returns a boolean if a field has been set.

### GetReflectedType

`func (o *Type) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *Type) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *Type) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *Type) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetUnderlyingSystemType

`func (o *Type) GetUnderlyingSystemType() Type`

GetUnderlyingSystemType returns the UnderlyingSystemType field if non-nil, zero value otherwise.

### GetUnderlyingSystemTypeOk

`func (o *Type) GetUnderlyingSystemTypeOk() (*Type, bool)`

GetUnderlyingSystemTypeOk returns a tuple with the UnderlyingSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSystemType

`func (o *Type) SetUnderlyingSystemType(v Type)`

SetUnderlyingSystemType sets UnderlyingSystemType field to given value.

### HasUnderlyingSystemType

`func (o *Type) HasUnderlyingSystemType() bool`

HasUnderlyingSystemType returns a boolean if a field has been set.

### GetIsTypeDefinition

`func (o *Type) GetIsTypeDefinition() bool`

GetIsTypeDefinition returns the IsTypeDefinition field if non-nil, zero value otherwise.

### GetIsTypeDefinitionOk

`func (o *Type) GetIsTypeDefinitionOk() (*bool, bool)`

GetIsTypeDefinitionOk returns a tuple with the IsTypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTypeDefinition

`func (o *Type) SetIsTypeDefinition(v bool)`

SetIsTypeDefinition sets IsTypeDefinition field to given value.

### HasIsTypeDefinition

`func (o *Type) HasIsTypeDefinition() bool`

HasIsTypeDefinition returns a boolean if a field has been set.

### GetIsArray

`func (o *Type) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *Type) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *Type) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *Type) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetIsByRef

`func (o *Type) GetIsByRef() bool`

GetIsByRef returns the IsByRef field if non-nil, zero value otherwise.

### GetIsByRefOk

`func (o *Type) GetIsByRefOk() (*bool, bool)`

GetIsByRefOk returns a tuple with the IsByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByRef

`func (o *Type) SetIsByRef(v bool)`

SetIsByRef sets IsByRef field to given value.

### HasIsByRef

`func (o *Type) HasIsByRef() bool`

HasIsByRef returns a boolean if a field has been set.

### GetIsPointer

`func (o *Type) GetIsPointer() bool`

GetIsPointer returns the IsPointer field if non-nil, zero value otherwise.

### GetIsPointerOk

`func (o *Type) GetIsPointerOk() (*bool, bool)`

GetIsPointerOk returns a tuple with the IsPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPointer

`func (o *Type) SetIsPointer(v bool)`

SetIsPointer sets IsPointer field to given value.

### HasIsPointer

`func (o *Type) HasIsPointer() bool`

HasIsPointer returns a boolean if a field has been set.

### GetIsConstructedGenericType

`func (o *Type) GetIsConstructedGenericType() bool`

GetIsConstructedGenericType returns the IsConstructedGenericType field if non-nil, zero value otherwise.

### GetIsConstructedGenericTypeOk

`func (o *Type) GetIsConstructedGenericTypeOk() (*bool, bool)`

GetIsConstructedGenericTypeOk returns a tuple with the IsConstructedGenericType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructedGenericType

`func (o *Type) SetIsConstructedGenericType(v bool)`

SetIsConstructedGenericType sets IsConstructedGenericType field to given value.

### HasIsConstructedGenericType

`func (o *Type) HasIsConstructedGenericType() bool`

HasIsConstructedGenericType returns a boolean if a field has been set.

### GetIsGenericParameter

`func (o *Type) GetIsGenericParameter() bool`

GetIsGenericParameter returns the IsGenericParameter field if non-nil, zero value otherwise.

### GetIsGenericParameterOk

`func (o *Type) GetIsGenericParameterOk() (*bool, bool)`

GetIsGenericParameterOk returns a tuple with the IsGenericParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericParameter

`func (o *Type) SetIsGenericParameter(v bool)`

SetIsGenericParameter sets IsGenericParameter field to given value.

### HasIsGenericParameter

`func (o *Type) HasIsGenericParameter() bool`

HasIsGenericParameter returns a boolean if a field has been set.

### GetIsGenericTypeParameter

`func (o *Type) GetIsGenericTypeParameter() bool`

GetIsGenericTypeParameter returns the IsGenericTypeParameter field if non-nil, zero value otherwise.

### GetIsGenericTypeParameterOk

`func (o *Type) GetIsGenericTypeParameterOk() (*bool, bool)`

GetIsGenericTypeParameterOk returns a tuple with the IsGenericTypeParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericTypeParameter

`func (o *Type) SetIsGenericTypeParameter(v bool)`

SetIsGenericTypeParameter sets IsGenericTypeParameter field to given value.

### HasIsGenericTypeParameter

`func (o *Type) HasIsGenericTypeParameter() bool`

HasIsGenericTypeParameter returns a boolean if a field has been set.

### GetIsGenericMethodParameter

`func (o *Type) GetIsGenericMethodParameter() bool`

GetIsGenericMethodParameter returns the IsGenericMethodParameter field if non-nil, zero value otherwise.

### GetIsGenericMethodParameterOk

`func (o *Type) GetIsGenericMethodParameterOk() (*bool, bool)`

GetIsGenericMethodParameterOk returns a tuple with the IsGenericMethodParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethodParameter

`func (o *Type) SetIsGenericMethodParameter(v bool)`

SetIsGenericMethodParameter sets IsGenericMethodParameter field to given value.

### HasIsGenericMethodParameter

`func (o *Type) HasIsGenericMethodParameter() bool`

HasIsGenericMethodParameter returns a boolean if a field has been set.

### GetIsGenericType

`func (o *Type) GetIsGenericType() bool`

GetIsGenericType returns the IsGenericType field if non-nil, zero value otherwise.

### GetIsGenericTypeOk

`func (o *Type) GetIsGenericTypeOk() (*bool, bool)`

GetIsGenericTypeOk returns a tuple with the IsGenericType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericType

`func (o *Type) SetIsGenericType(v bool)`

SetIsGenericType sets IsGenericType field to given value.

### HasIsGenericType

`func (o *Type) HasIsGenericType() bool`

HasIsGenericType returns a boolean if a field has been set.

### GetIsGenericTypeDefinition

`func (o *Type) GetIsGenericTypeDefinition() bool`

GetIsGenericTypeDefinition returns the IsGenericTypeDefinition field if non-nil, zero value otherwise.

### GetIsGenericTypeDefinitionOk

`func (o *Type) GetIsGenericTypeDefinitionOk() (*bool, bool)`

GetIsGenericTypeDefinitionOk returns a tuple with the IsGenericTypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericTypeDefinition

`func (o *Type) SetIsGenericTypeDefinition(v bool)`

SetIsGenericTypeDefinition sets IsGenericTypeDefinition field to given value.

### HasIsGenericTypeDefinition

`func (o *Type) HasIsGenericTypeDefinition() bool`

HasIsGenericTypeDefinition returns a boolean if a field has been set.

### GetIsSZArray

`func (o *Type) GetIsSZArray() bool`

GetIsSZArray returns the IsSZArray field if non-nil, zero value otherwise.

### GetIsSZArrayOk

`func (o *Type) GetIsSZArrayOk() (*bool, bool)`

GetIsSZArrayOk returns a tuple with the IsSZArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSZArray

`func (o *Type) SetIsSZArray(v bool)`

SetIsSZArray sets IsSZArray field to given value.

### HasIsSZArray

`func (o *Type) HasIsSZArray() bool`

HasIsSZArray returns a boolean if a field has been set.

### GetIsVariableBoundArray

`func (o *Type) GetIsVariableBoundArray() bool`

GetIsVariableBoundArray returns the IsVariableBoundArray field if non-nil, zero value otherwise.

### GetIsVariableBoundArrayOk

`func (o *Type) GetIsVariableBoundArrayOk() (*bool, bool)`

GetIsVariableBoundArrayOk returns a tuple with the IsVariableBoundArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariableBoundArray

`func (o *Type) SetIsVariableBoundArray(v bool)`

SetIsVariableBoundArray sets IsVariableBoundArray field to given value.

### HasIsVariableBoundArray

`func (o *Type) HasIsVariableBoundArray() bool`

HasIsVariableBoundArray returns a boolean if a field has been set.

### GetIsByRefLike

`func (o *Type) GetIsByRefLike() bool`

GetIsByRefLike returns the IsByRefLike field if non-nil, zero value otherwise.

### GetIsByRefLikeOk

`func (o *Type) GetIsByRefLikeOk() (*bool, bool)`

GetIsByRefLikeOk returns a tuple with the IsByRefLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByRefLike

`func (o *Type) SetIsByRefLike(v bool)`

SetIsByRefLike sets IsByRefLike field to given value.

### HasIsByRefLike

`func (o *Type) HasIsByRefLike() bool`

HasIsByRefLike returns a boolean if a field has been set.

### GetIsFunctionPointer

`func (o *Type) GetIsFunctionPointer() bool`

GetIsFunctionPointer returns the IsFunctionPointer field if non-nil, zero value otherwise.

### GetIsFunctionPointerOk

`func (o *Type) GetIsFunctionPointerOk() (*bool, bool)`

GetIsFunctionPointerOk returns a tuple with the IsFunctionPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFunctionPointer

`func (o *Type) SetIsFunctionPointer(v bool)`

SetIsFunctionPointer sets IsFunctionPointer field to given value.

### HasIsFunctionPointer

`func (o *Type) HasIsFunctionPointer() bool`

HasIsFunctionPointer returns a boolean if a field has been set.

### GetIsUnmanagedFunctionPointer

`func (o *Type) GetIsUnmanagedFunctionPointer() bool`

GetIsUnmanagedFunctionPointer returns the IsUnmanagedFunctionPointer field if non-nil, zero value otherwise.

### GetIsUnmanagedFunctionPointerOk

`func (o *Type) GetIsUnmanagedFunctionPointerOk() (*bool, bool)`

GetIsUnmanagedFunctionPointerOk returns a tuple with the IsUnmanagedFunctionPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnmanagedFunctionPointer

`func (o *Type) SetIsUnmanagedFunctionPointer(v bool)`

SetIsUnmanagedFunctionPointer sets IsUnmanagedFunctionPointer field to given value.

### HasIsUnmanagedFunctionPointer

`func (o *Type) HasIsUnmanagedFunctionPointer() bool`

HasIsUnmanagedFunctionPointer returns a boolean if a field has been set.

### GetHasElementType

`func (o *Type) GetHasElementType() bool`

GetHasElementType returns the HasElementType field if non-nil, zero value otherwise.

### GetHasElementTypeOk

`func (o *Type) GetHasElementTypeOk() (*bool, bool)`

GetHasElementTypeOk returns a tuple with the HasElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasElementType

`func (o *Type) SetHasElementType(v bool)`

SetHasElementType sets HasElementType field to given value.

### HasHasElementType

`func (o *Type) HasHasElementType() bool`

HasHasElementType returns a boolean if a field has been set.

### GetGenericTypeArguments

`func (o *Type) GetGenericTypeArguments() []Type`

GetGenericTypeArguments returns the GenericTypeArguments field if non-nil, zero value otherwise.

### GetGenericTypeArgumentsOk

`func (o *Type) GetGenericTypeArgumentsOk() (*[]Type, bool)`

GetGenericTypeArgumentsOk returns a tuple with the GenericTypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypeArguments

`func (o *Type) SetGenericTypeArguments(v []Type)`

SetGenericTypeArguments sets GenericTypeArguments field to given value.

### HasGenericTypeArguments

`func (o *Type) HasGenericTypeArguments() bool`

HasGenericTypeArguments returns a boolean if a field has been set.

### SetGenericTypeArgumentsNil

`func (o *Type) SetGenericTypeArgumentsNil(b bool)`

 SetGenericTypeArgumentsNil sets the value for GenericTypeArguments to be an explicit nil

### UnsetGenericTypeArguments
`func (o *Type) UnsetGenericTypeArguments()`

UnsetGenericTypeArguments ensures that no value is present for GenericTypeArguments, not even an explicit nil
### GetGenericParameterPosition

`func (o *Type) GetGenericParameterPosition() int32`

GetGenericParameterPosition returns the GenericParameterPosition field if non-nil, zero value otherwise.

### GetGenericParameterPositionOk

`func (o *Type) GetGenericParameterPositionOk() (*int32, bool)`

GetGenericParameterPositionOk returns a tuple with the GenericParameterPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericParameterPosition

`func (o *Type) SetGenericParameterPosition(v int32)`

SetGenericParameterPosition sets GenericParameterPosition field to given value.

### HasGenericParameterPosition

`func (o *Type) HasGenericParameterPosition() bool`

HasGenericParameterPosition returns a boolean if a field has been set.

### GetGenericParameterAttributes

`func (o *Type) GetGenericParameterAttributes() GenericParameterAttributes`

GetGenericParameterAttributes returns the GenericParameterAttributes field if non-nil, zero value otherwise.

### GetGenericParameterAttributesOk

`func (o *Type) GetGenericParameterAttributesOk() (*GenericParameterAttributes, bool)`

GetGenericParameterAttributesOk returns a tuple with the GenericParameterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericParameterAttributes

`func (o *Type) SetGenericParameterAttributes(v GenericParameterAttributes)`

SetGenericParameterAttributes sets GenericParameterAttributes field to given value.

### HasGenericParameterAttributes

`func (o *Type) HasGenericParameterAttributes() bool`

HasGenericParameterAttributes returns a boolean if a field has been set.

### GetAttributes

`func (o *Type) GetAttributes() TypeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Type) GetAttributesOk() (*TypeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Type) SetAttributes(v TypeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Type) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsAbstract

`func (o *Type) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *Type) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *Type) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *Type) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsImport

`func (o *Type) GetIsImport() bool`

GetIsImport returns the IsImport field if non-nil, zero value otherwise.

### GetIsImportOk

`func (o *Type) GetIsImportOk() (*bool, bool)`

GetIsImportOk returns a tuple with the IsImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImport

`func (o *Type) SetIsImport(v bool)`

SetIsImport sets IsImport field to given value.

### HasIsImport

`func (o *Type) HasIsImport() bool`

HasIsImport returns a boolean if a field has been set.

### GetIsSealed

`func (o *Type) GetIsSealed() bool`

GetIsSealed returns the IsSealed field if non-nil, zero value otherwise.

### GetIsSealedOk

`func (o *Type) GetIsSealedOk() (*bool, bool)`

GetIsSealedOk returns a tuple with the IsSealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSealed

`func (o *Type) SetIsSealed(v bool)`

SetIsSealed sets IsSealed field to given value.

### HasIsSealed

`func (o *Type) HasIsSealed() bool`

HasIsSealed returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *Type) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *Type) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *Type) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *Type) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetIsClass

`func (o *Type) GetIsClass() bool`

GetIsClass returns the IsClass field if non-nil, zero value otherwise.

### GetIsClassOk

`func (o *Type) GetIsClassOk() (*bool, bool)`

GetIsClassOk returns a tuple with the IsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClass

`func (o *Type) SetIsClass(v bool)`

SetIsClass sets IsClass field to given value.

### HasIsClass

`func (o *Type) HasIsClass() bool`

HasIsClass returns a boolean if a field has been set.

### GetIsNestedAssembly

`func (o *Type) GetIsNestedAssembly() bool`

GetIsNestedAssembly returns the IsNestedAssembly field if non-nil, zero value otherwise.

### GetIsNestedAssemblyOk

`func (o *Type) GetIsNestedAssemblyOk() (*bool, bool)`

GetIsNestedAssemblyOk returns a tuple with the IsNestedAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedAssembly

`func (o *Type) SetIsNestedAssembly(v bool)`

SetIsNestedAssembly sets IsNestedAssembly field to given value.

### HasIsNestedAssembly

`func (o *Type) HasIsNestedAssembly() bool`

HasIsNestedAssembly returns a boolean if a field has been set.

### GetIsNestedFamANDAssem

`func (o *Type) GetIsNestedFamANDAssem() bool`

GetIsNestedFamANDAssem returns the IsNestedFamANDAssem field if non-nil, zero value otherwise.

### GetIsNestedFamANDAssemOk

`func (o *Type) GetIsNestedFamANDAssemOk() (*bool, bool)`

GetIsNestedFamANDAssemOk returns a tuple with the IsNestedFamANDAssem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamANDAssem

`func (o *Type) SetIsNestedFamANDAssem(v bool)`

SetIsNestedFamANDAssem sets IsNestedFamANDAssem field to given value.

### HasIsNestedFamANDAssem

`func (o *Type) HasIsNestedFamANDAssem() bool`

HasIsNestedFamANDAssem returns a boolean if a field has been set.

### GetIsNestedFamily

`func (o *Type) GetIsNestedFamily() bool`

GetIsNestedFamily returns the IsNestedFamily field if non-nil, zero value otherwise.

### GetIsNestedFamilyOk

`func (o *Type) GetIsNestedFamilyOk() (*bool, bool)`

GetIsNestedFamilyOk returns a tuple with the IsNestedFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamily

`func (o *Type) SetIsNestedFamily(v bool)`

SetIsNestedFamily sets IsNestedFamily field to given value.

### HasIsNestedFamily

`func (o *Type) HasIsNestedFamily() bool`

HasIsNestedFamily returns a boolean if a field has been set.

### GetIsNestedFamORAssem

`func (o *Type) GetIsNestedFamORAssem() bool`

GetIsNestedFamORAssem returns the IsNestedFamORAssem field if non-nil, zero value otherwise.

### GetIsNestedFamORAssemOk

`func (o *Type) GetIsNestedFamORAssemOk() (*bool, bool)`

GetIsNestedFamORAssemOk returns a tuple with the IsNestedFamORAssem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamORAssem

`func (o *Type) SetIsNestedFamORAssem(v bool)`

SetIsNestedFamORAssem sets IsNestedFamORAssem field to given value.

### HasIsNestedFamORAssem

`func (o *Type) HasIsNestedFamORAssem() bool`

HasIsNestedFamORAssem returns a boolean if a field has been set.

### GetIsNestedPrivate

`func (o *Type) GetIsNestedPrivate() bool`

GetIsNestedPrivate returns the IsNestedPrivate field if non-nil, zero value otherwise.

### GetIsNestedPrivateOk

`func (o *Type) GetIsNestedPrivateOk() (*bool, bool)`

GetIsNestedPrivateOk returns a tuple with the IsNestedPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedPrivate

`func (o *Type) SetIsNestedPrivate(v bool)`

SetIsNestedPrivate sets IsNestedPrivate field to given value.

### HasIsNestedPrivate

`func (o *Type) HasIsNestedPrivate() bool`

HasIsNestedPrivate returns a boolean if a field has been set.

### GetIsNestedPublic

`func (o *Type) GetIsNestedPublic() bool`

GetIsNestedPublic returns the IsNestedPublic field if non-nil, zero value otherwise.

### GetIsNestedPublicOk

`func (o *Type) GetIsNestedPublicOk() (*bool, bool)`

GetIsNestedPublicOk returns a tuple with the IsNestedPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedPublic

`func (o *Type) SetIsNestedPublic(v bool)`

SetIsNestedPublic sets IsNestedPublic field to given value.

### HasIsNestedPublic

`func (o *Type) HasIsNestedPublic() bool`

HasIsNestedPublic returns a boolean if a field has been set.

### GetIsNotPublic

`func (o *Type) GetIsNotPublic() bool`

GetIsNotPublic returns the IsNotPublic field if non-nil, zero value otherwise.

### GetIsNotPublicOk

`func (o *Type) GetIsNotPublicOk() (*bool, bool)`

GetIsNotPublicOk returns a tuple with the IsNotPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotPublic

`func (o *Type) SetIsNotPublic(v bool)`

SetIsNotPublic sets IsNotPublic field to given value.

### HasIsNotPublic

`func (o *Type) HasIsNotPublic() bool`

HasIsNotPublic returns a boolean if a field has been set.

### GetIsPublic

`func (o *Type) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Type) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Type) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Type) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsAutoLayout

`func (o *Type) GetIsAutoLayout() bool`

GetIsAutoLayout returns the IsAutoLayout field if non-nil, zero value otherwise.

### GetIsAutoLayoutOk

`func (o *Type) GetIsAutoLayoutOk() (*bool, bool)`

GetIsAutoLayoutOk returns a tuple with the IsAutoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoLayout

`func (o *Type) SetIsAutoLayout(v bool)`

SetIsAutoLayout sets IsAutoLayout field to given value.

### HasIsAutoLayout

`func (o *Type) HasIsAutoLayout() bool`

HasIsAutoLayout returns a boolean if a field has been set.

### GetIsExplicitLayout

`func (o *Type) GetIsExplicitLayout() bool`

GetIsExplicitLayout returns the IsExplicitLayout field if non-nil, zero value otherwise.

### GetIsExplicitLayoutOk

`func (o *Type) GetIsExplicitLayoutOk() (*bool, bool)`

GetIsExplicitLayoutOk returns a tuple with the IsExplicitLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExplicitLayout

`func (o *Type) SetIsExplicitLayout(v bool)`

SetIsExplicitLayout sets IsExplicitLayout field to given value.

### HasIsExplicitLayout

`func (o *Type) HasIsExplicitLayout() bool`

HasIsExplicitLayout returns a boolean if a field has been set.

### GetIsLayoutSequential

`func (o *Type) GetIsLayoutSequential() bool`

GetIsLayoutSequential returns the IsLayoutSequential field if non-nil, zero value otherwise.

### GetIsLayoutSequentialOk

`func (o *Type) GetIsLayoutSequentialOk() (*bool, bool)`

GetIsLayoutSequentialOk returns a tuple with the IsLayoutSequential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLayoutSequential

`func (o *Type) SetIsLayoutSequential(v bool)`

SetIsLayoutSequential sets IsLayoutSequential field to given value.

### HasIsLayoutSequential

`func (o *Type) HasIsLayoutSequential() bool`

HasIsLayoutSequential returns a boolean if a field has been set.

### GetIsAnsiClass

`func (o *Type) GetIsAnsiClass() bool`

GetIsAnsiClass returns the IsAnsiClass field if non-nil, zero value otherwise.

### GetIsAnsiClassOk

`func (o *Type) GetIsAnsiClassOk() (*bool, bool)`

GetIsAnsiClassOk returns a tuple with the IsAnsiClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnsiClass

`func (o *Type) SetIsAnsiClass(v bool)`

SetIsAnsiClass sets IsAnsiClass field to given value.

### HasIsAnsiClass

`func (o *Type) HasIsAnsiClass() bool`

HasIsAnsiClass returns a boolean if a field has been set.

### GetIsAutoClass

`func (o *Type) GetIsAutoClass() bool`

GetIsAutoClass returns the IsAutoClass field if non-nil, zero value otherwise.

### GetIsAutoClassOk

`func (o *Type) GetIsAutoClassOk() (*bool, bool)`

GetIsAutoClassOk returns a tuple with the IsAutoClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoClass

`func (o *Type) SetIsAutoClass(v bool)`

SetIsAutoClass sets IsAutoClass field to given value.

### HasIsAutoClass

`func (o *Type) HasIsAutoClass() bool`

HasIsAutoClass returns a boolean if a field has been set.

### GetIsUnicodeClass

`func (o *Type) GetIsUnicodeClass() bool`

GetIsUnicodeClass returns the IsUnicodeClass field if non-nil, zero value otherwise.

### GetIsUnicodeClassOk

`func (o *Type) GetIsUnicodeClassOk() (*bool, bool)`

GetIsUnicodeClassOk returns a tuple with the IsUnicodeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnicodeClass

`func (o *Type) SetIsUnicodeClass(v bool)`

SetIsUnicodeClass sets IsUnicodeClass field to given value.

### HasIsUnicodeClass

`func (o *Type) HasIsUnicodeClass() bool`

HasIsUnicodeClass returns a boolean if a field has been set.

### GetIsCOMObject

`func (o *Type) GetIsCOMObject() bool`

GetIsCOMObject returns the IsCOMObject field if non-nil, zero value otherwise.

### GetIsCOMObjectOk

`func (o *Type) GetIsCOMObjectOk() (*bool, bool)`

GetIsCOMObjectOk returns a tuple with the IsCOMObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCOMObject

`func (o *Type) SetIsCOMObject(v bool)`

SetIsCOMObject sets IsCOMObject field to given value.

### HasIsCOMObject

`func (o *Type) HasIsCOMObject() bool`

HasIsCOMObject returns a boolean if a field has been set.

### GetIsContextful

`func (o *Type) GetIsContextful() bool`

GetIsContextful returns the IsContextful field if non-nil, zero value otherwise.

### GetIsContextfulOk

`func (o *Type) GetIsContextfulOk() (*bool, bool)`

GetIsContextfulOk returns a tuple with the IsContextful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextful

`func (o *Type) SetIsContextful(v bool)`

SetIsContextful sets IsContextful field to given value.

### HasIsContextful

`func (o *Type) HasIsContextful() bool`

HasIsContextful returns a boolean if a field has been set.

### GetIsEnum

`func (o *Type) GetIsEnum() bool`

GetIsEnum returns the IsEnum field if non-nil, zero value otherwise.

### GetIsEnumOk

`func (o *Type) GetIsEnumOk() (*bool, bool)`

GetIsEnumOk returns a tuple with the IsEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnum

`func (o *Type) SetIsEnum(v bool)`

SetIsEnum sets IsEnum field to given value.

### HasIsEnum

`func (o *Type) HasIsEnum() bool`

HasIsEnum returns a boolean if a field has been set.

### GetIsMarshalByRef

`func (o *Type) GetIsMarshalByRef() bool`

GetIsMarshalByRef returns the IsMarshalByRef field if non-nil, zero value otherwise.

### GetIsMarshalByRefOk

`func (o *Type) GetIsMarshalByRefOk() (*bool, bool)`

GetIsMarshalByRefOk returns a tuple with the IsMarshalByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarshalByRef

`func (o *Type) SetIsMarshalByRef(v bool)`

SetIsMarshalByRef sets IsMarshalByRef field to given value.

### HasIsMarshalByRef

`func (o *Type) HasIsMarshalByRef() bool`

HasIsMarshalByRef returns a boolean if a field has been set.

### GetIsPrimitive

`func (o *Type) GetIsPrimitive() bool`

GetIsPrimitive returns the IsPrimitive field if non-nil, zero value otherwise.

### GetIsPrimitiveOk

`func (o *Type) GetIsPrimitiveOk() (*bool, bool)`

GetIsPrimitiveOk returns a tuple with the IsPrimitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimitive

`func (o *Type) SetIsPrimitive(v bool)`

SetIsPrimitive sets IsPrimitive field to given value.

### HasIsPrimitive

`func (o *Type) HasIsPrimitive() bool`

HasIsPrimitive returns a boolean if a field has been set.

### GetIsValueType

`func (o *Type) GetIsValueType() bool`

GetIsValueType returns the IsValueType field if non-nil, zero value otherwise.

### GetIsValueTypeOk

`func (o *Type) GetIsValueTypeOk() (*bool, bool)`

GetIsValueTypeOk returns a tuple with the IsValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueType

`func (o *Type) SetIsValueType(v bool)`

SetIsValueType sets IsValueType field to given value.

### HasIsValueType

`func (o *Type) HasIsValueType() bool`

HasIsValueType returns a boolean if a field has been set.

### GetIsSignatureType

`func (o *Type) GetIsSignatureType() bool`

GetIsSignatureType returns the IsSignatureType field if non-nil, zero value otherwise.

### GetIsSignatureTypeOk

`func (o *Type) GetIsSignatureTypeOk() (*bool, bool)`

GetIsSignatureTypeOk returns a tuple with the IsSignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSignatureType

`func (o *Type) SetIsSignatureType(v bool)`

SetIsSignatureType sets IsSignatureType field to given value.

### HasIsSignatureType

`func (o *Type) HasIsSignatureType() bool`

HasIsSignatureType returns a boolean if a field has been set.

### GetIsSecurityCritical

`func (o *Type) GetIsSecurityCritical() bool`

GetIsSecurityCritical returns the IsSecurityCritical field if non-nil, zero value otherwise.

### GetIsSecurityCriticalOk

`func (o *Type) GetIsSecurityCriticalOk() (*bool, bool)`

GetIsSecurityCriticalOk returns a tuple with the IsSecurityCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityCritical

`func (o *Type) SetIsSecurityCritical(v bool)`

SetIsSecurityCritical sets IsSecurityCritical field to given value.

### HasIsSecurityCritical

`func (o *Type) HasIsSecurityCritical() bool`

HasIsSecurityCritical returns a boolean if a field has been set.

### GetIsSecuritySafeCritical

`func (o *Type) GetIsSecuritySafeCritical() bool`

GetIsSecuritySafeCritical returns the IsSecuritySafeCritical field if non-nil, zero value otherwise.

### GetIsSecuritySafeCriticalOk

`func (o *Type) GetIsSecuritySafeCriticalOk() (*bool, bool)`

GetIsSecuritySafeCriticalOk returns a tuple with the IsSecuritySafeCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySafeCritical

`func (o *Type) SetIsSecuritySafeCritical(v bool)`

SetIsSecuritySafeCritical sets IsSecuritySafeCritical field to given value.

### HasIsSecuritySafeCritical

`func (o *Type) HasIsSecuritySafeCritical() bool`

HasIsSecuritySafeCritical returns a boolean if a field has been set.

### GetIsSecurityTransparent

`func (o *Type) GetIsSecurityTransparent() bool`

GetIsSecurityTransparent returns the IsSecurityTransparent field if non-nil, zero value otherwise.

### GetIsSecurityTransparentOk

`func (o *Type) GetIsSecurityTransparentOk() (*bool, bool)`

GetIsSecurityTransparentOk returns a tuple with the IsSecurityTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityTransparent

`func (o *Type) SetIsSecurityTransparent(v bool)`

SetIsSecurityTransparent sets IsSecurityTransparent field to given value.

### HasIsSecurityTransparent

`func (o *Type) HasIsSecurityTransparent() bool`

HasIsSecurityTransparent returns a boolean if a field has been set.

### GetStructLayoutAttribute

`func (o *Type) GetStructLayoutAttribute() StructLayoutAttribute`

GetStructLayoutAttribute returns the StructLayoutAttribute field if non-nil, zero value otherwise.

### GetStructLayoutAttributeOk

`func (o *Type) GetStructLayoutAttributeOk() (*StructLayoutAttribute, bool)`

GetStructLayoutAttributeOk returns a tuple with the StructLayoutAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructLayoutAttribute

`func (o *Type) SetStructLayoutAttribute(v StructLayoutAttribute)`

SetStructLayoutAttribute sets StructLayoutAttribute field to given value.

### HasStructLayoutAttribute

`func (o *Type) HasStructLayoutAttribute() bool`

HasStructLayoutAttribute returns a boolean if a field has been set.

### GetTypeInitializer

`func (o *Type) GetTypeInitializer() ConstructorInfo`

GetTypeInitializer returns the TypeInitializer field if non-nil, zero value otherwise.

### GetTypeInitializerOk

`func (o *Type) GetTypeInitializerOk() (*ConstructorInfo, bool)`

GetTypeInitializerOk returns a tuple with the TypeInitializer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInitializer

`func (o *Type) SetTypeInitializer(v ConstructorInfo)`

SetTypeInitializer sets TypeInitializer field to given value.

### HasTypeInitializer

`func (o *Type) HasTypeInitializer() bool`

HasTypeInitializer returns a boolean if a field has been set.

### GetTypeHandle

`func (o *Type) GetTypeHandle() RuntimeTypeHandle`

GetTypeHandle returns the TypeHandle field if non-nil, zero value otherwise.

### GetTypeHandleOk

`func (o *Type) GetTypeHandleOk() (*RuntimeTypeHandle, bool)`

GetTypeHandleOk returns a tuple with the TypeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeHandle

`func (o *Type) SetTypeHandle(v RuntimeTypeHandle)`

SetTypeHandle sets TypeHandle field to given value.

### HasTypeHandle

`func (o *Type) HasTypeHandle() bool`

HasTypeHandle returns a boolean if a field has been set.

### GetGuid

`func (o *Type) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Type) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Type) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Type) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetBaseType

`func (o *Type) GetBaseType() Type`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *Type) GetBaseTypeOk() (*Type, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *Type) SetBaseType(v Type)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *Type) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetIsSerializable

`func (o *Type) GetIsSerializable() bool`

GetIsSerializable returns the IsSerializable field if non-nil, zero value otherwise.

### GetIsSerializableOk

`func (o *Type) GetIsSerializableOk() (*bool, bool)`

GetIsSerializableOk returns a tuple with the IsSerializable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSerializable

`func (o *Type) SetIsSerializable(v bool)`

SetIsSerializable sets IsSerializable field to given value.

### HasIsSerializable

`func (o *Type) HasIsSerializable() bool`

HasIsSerializable returns a boolean if a field has been set.

### GetContainsGenericParameters

`func (o *Type) GetContainsGenericParameters() bool`

GetContainsGenericParameters returns the ContainsGenericParameters field if non-nil, zero value otherwise.

### GetContainsGenericParametersOk

`func (o *Type) GetContainsGenericParametersOk() (*bool, bool)`

GetContainsGenericParametersOk returns a tuple with the ContainsGenericParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsGenericParameters

`func (o *Type) SetContainsGenericParameters(v bool)`

SetContainsGenericParameters sets ContainsGenericParameters field to given value.

### HasContainsGenericParameters

`func (o *Type) HasContainsGenericParameters() bool`

HasContainsGenericParameters returns a boolean if a field has been set.

### GetIsVisible

`func (o *Type) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *Type) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *Type) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *Type) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


