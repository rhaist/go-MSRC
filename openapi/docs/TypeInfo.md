# TypeInfo

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
**GenericTypeParameters** | Pointer to [**[]Type**](Type.md) |  | [optional] [readonly] 
**DeclaredConstructors** | Pointer to [**[]ConstructorInfo**](ConstructorInfo.md) |  | [optional] [readonly] 
**DeclaredEvents** | Pointer to [**[]EventInfo**](EventInfo.md) |  | [optional] [readonly] 
**DeclaredFields** | Pointer to [**[]FieldInfo**](FieldInfo.md) |  | [optional] [readonly] 
**DeclaredMembers** | Pointer to [**[]MemberInfo**](MemberInfo.md) |  | [optional] [readonly] 
**DeclaredMethods** | Pointer to [**[]MethodInfo**](MethodInfo.md) |  | [optional] [readonly] 
**DeclaredNestedTypes** | Pointer to [**[]TypeInfo**](TypeInfo.md) |  | [optional] [readonly] 
**DeclaredProperties** | Pointer to [**[]PropertyInfo**](PropertyInfo.md) |  | [optional] [readonly] 
**ImplementedInterfaces** | Pointer to [**[]Type**](Type.md) |  | [optional] [readonly] 

## Methods

### NewTypeInfo

`func NewTypeInfo() *TypeInfo`

NewTypeInfo instantiates a new TypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeInfoWithDefaults

`func NewTypeInfoWithDefaults() *TypeInfo`

NewTypeInfoWithDefaults instantiates a new TypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TypeInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TypeInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomAttributes

`func (o *TypeInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *TypeInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *TypeInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *TypeInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *TypeInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *TypeInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *TypeInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *TypeInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *TypeInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *TypeInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *TypeInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *TypeInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *TypeInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *TypeInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetIsInterface

`func (o *TypeInfo) GetIsInterface() bool`

GetIsInterface returns the IsInterface field if non-nil, zero value otherwise.

### GetIsInterfaceOk

`func (o *TypeInfo) GetIsInterfaceOk() (*bool, bool)`

GetIsInterfaceOk returns a tuple with the IsInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterface

`func (o *TypeInfo) SetIsInterface(v bool)`

SetIsInterface sets IsInterface field to given value.

### HasIsInterface

`func (o *TypeInfo) HasIsInterface() bool`

HasIsInterface returns a boolean if a field has been set.

### GetMemberType

`func (o *TypeInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *TypeInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *TypeInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *TypeInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetNamespace

`func (o *TypeInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TypeInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TypeInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TypeInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *TypeInfo) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *TypeInfo) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetAssemblyQualifiedName

`func (o *TypeInfo) GetAssemblyQualifiedName() string`

GetAssemblyQualifiedName returns the AssemblyQualifiedName field if non-nil, zero value otherwise.

### GetAssemblyQualifiedNameOk

`func (o *TypeInfo) GetAssemblyQualifiedNameOk() (*string, bool)`

GetAssemblyQualifiedNameOk returns a tuple with the AssemblyQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyQualifiedName

`func (o *TypeInfo) SetAssemblyQualifiedName(v string)`

SetAssemblyQualifiedName sets AssemblyQualifiedName field to given value.

### HasAssemblyQualifiedName

`func (o *TypeInfo) HasAssemblyQualifiedName() bool`

HasAssemblyQualifiedName returns a boolean if a field has been set.

### SetAssemblyQualifiedNameNil

`func (o *TypeInfo) SetAssemblyQualifiedNameNil(b bool)`

 SetAssemblyQualifiedNameNil sets the value for AssemblyQualifiedName to be an explicit nil

### UnsetAssemblyQualifiedName
`func (o *TypeInfo) UnsetAssemblyQualifiedName()`

UnsetAssemblyQualifiedName ensures that no value is present for AssemblyQualifiedName, not even an explicit nil
### GetFullName

`func (o *TypeInfo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *TypeInfo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *TypeInfo) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *TypeInfo) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *TypeInfo) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *TypeInfo) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetAssembly

`func (o *TypeInfo) GetAssembly() Assembly`

GetAssembly returns the Assembly field if non-nil, zero value otherwise.

### GetAssemblyOk

`func (o *TypeInfo) GetAssemblyOk() (*Assembly, bool)`

GetAssemblyOk returns a tuple with the Assembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembly

`func (o *TypeInfo) SetAssembly(v Assembly)`

SetAssembly sets Assembly field to given value.

### HasAssembly

`func (o *TypeInfo) HasAssembly() bool`

HasAssembly returns a boolean if a field has been set.

### GetModule

`func (o *TypeInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *TypeInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *TypeInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *TypeInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetIsNested

`func (o *TypeInfo) GetIsNested() bool`

GetIsNested returns the IsNested field if non-nil, zero value otherwise.

### GetIsNestedOk

`func (o *TypeInfo) GetIsNestedOk() (*bool, bool)`

GetIsNestedOk returns a tuple with the IsNested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNested

`func (o *TypeInfo) SetIsNested(v bool)`

SetIsNested sets IsNested field to given value.

### HasIsNested

`func (o *TypeInfo) HasIsNested() bool`

HasIsNested returns a boolean if a field has been set.

### GetDeclaringType

`func (o *TypeInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *TypeInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *TypeInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *TypeInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetDeclaringMethod

`func (o *TypeInfo) GetDeclaringMethod() MethodBase`

GetDeclaringMethod returns the DeclaringMethod field if non-nil, zero value otherwise.

### GetDeclaringMethodOk

`func (o *TypeInfo) GetDeclaringMethodOk() (*MethodBase, bool)`

GetDeclaringMethodOk returns a tuple with the DeclaringMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringMethod

`func (o *TypeInfo) SetDeclaringMethod(v MethodBase)`

SetDeclaringMethod sets DeclaringMethod field to given value.

### HasDeclaringMethod

`func (o *TypeInfo) HasDeclaringMethod() bool`

HasDeclaringMethod returns a boolean if a field has been set.

### GetReflectedType

`func (o *TypeInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *TypeInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *TypeInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *TypeInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetUnderlyingSystemType

`func (o *TypeInfo) GetUnderlyingSystemType() Type`

GetUnderlyingSystemType returns the UnderlyingSystemType field if non-nil, zero value otherwise.

### GetUnderlyingSystemTypeOk

`func (o *TypeInfo) GetUnderlyingSystemTypeOk() (*Type, bool)`

GetUnderlyingSystemTypeOk returns a tuple with the UnderlyingSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSystemType

`func (o *TypeInfo) SetUnderlyingSystemType(v Type)`

SetUnderlyingSystemType sets UnderlyingSystemType field to given value.

### HasUnderlyingSystemType

`func (o *TypeInfo) HasUnderlyingSystemType() bool`

HasUnderlyingSystemType returns a boolean if a field has been set.

### GetIsTypeDefinition

`func (o *TypeInfo) GetIsTypeDefinition() bool`

GetIsTypeDefinition returns the IsTypeDefinition field if non-nil, zero value otherwise.

### GetIsTypeDefinitionOk

`func (o *TypeInfo) GetIsTypeDefinitionOk() (*bool, bool)`

GetIsTypeDefinitionOk returns a tuple with the IsTypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTypeDefinition

`func (o *TypeInfo) SetIsTypeDefinition(v bool)`

SetIsTypeDefinition sets IsTypeDefinition field to given value.

### HasIsTypeDefinition

`func (o *TypeInfo) HasIsTypeDefinition() bool`

HasIsTypeDefinition returns a boolean if a field has been set.

### GetIsArray

`func (o *TypeInfo) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *TypeInfo) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *TypeInfo) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *TypeInfo) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetIsByRef

`func (o *TypeInfo) GetIsByRef() bool`

GetIsByRef returns the IsByRef field if non-nil, zero value otherwise.

### GetIsByRefOk

`func (o *TypeInfo) GetIsByRefOk() (*bool, bool)`

GetIsByRefOk returns a tuple with the IsByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByRef

`func (o *TypeInfo) SetIsByRef(v bool)`

SetIsByRef sets IsByRef field to given value.

### HasIsByRef

`func (o *TypeInfo) HasIsByRef() bool`

HasIsByRef returns a boolean if a field has been set.

### GetIsPointer

`func (o *TypeInfo) GetIsPointer() bool`

GetIsPointer returns the IsPointer field if non-nil, zero value otherwise.

### GetIsPointerOk

`func (o *TypeInfo) GetIsPointerOk() (*bool, bool)`

GetIsPointerOk returns a tuple with the IsPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPointer

`func (o *TypeInfo) SetIsPointer(v bool)`

SetIsPointer sets IsPointer field to given value.

### HasIsPointer

`func (o *TypeInfo) HasIsPointer() bool`

HasIsPointer returns a boolean if a field has been set.

### GetIsConstructedGenericType

`func (o *TypeInfo) GetIsConstructedGenericType() bool`

GetIsConstructedGenericType returns the IsConstructedGenericType field if non-nil, zero value otherwise.

### GetIsConstructedGenericTypeOk

`func (o *TypeInfo) GetIsConstructedGenericTypeOk() (*bool, bool)`

GetIsConstructedGenericTypeOk returns a tuple with the IsConstructedGenericType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstructedGenericType

`func (o *TypeInfo) SetIsConstructedGenericType(v bool)`

SetIsConstructedGenericType sets IsConstructedGenericType field to given value.

### HasIsConstructedGenericType

`func (o *TypeInfo) HasIsConstructedGenericType() bool`

HasIsConstructedGenericType returns a boolean if a field has been set.

### GetIsGenericParameter

`func (o *TypeInfo) GetIsGenericParameter() bool`

GetIsGenericParameter returns the IsGenericParameter field if non-nil, zero value otherwise.

### GetIsGenericParameterOk

`func (o *TypeInfo) GetIsGenericParameterOk() (*bool, bool)`

GetIsGenericParameterOk returns a tuple with the IsGenericParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericParameter

`func (o *TypeInfo) SetIsGenericParameter(v bool)`

SetIsGenericParameter sets IsGenericParameter field to given value.

### HasIsGenericParameter

`func (o *TypeInfo) HasIsGenericParameter() bool`

HasIsGenericParameter returns a boolean if a field has been set.

### GetIsGenericTypeParameter

`func (o *TypeInfo) GetIsGenericTypeParameter() bool`

GetIsGenericTypeParameter returns the IsGenericTypeParameter field if non-nil, zero value otherwise.

### GetIsGenericTypeParameterOk

`func (o *TypeInfo) GetIsGenericTypeParameterOk() (*bool, bool)`

GetIsGenericTypeParameterOk returns a tuple with the IsGenericTypeParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericTypeParameter

`func (o *TypeInfo) SetIsGenericTypeParameter(v bool)`

SetIsGenericTypeParameter sets IsGenericTypeParameter field to given value.

### HasIsGenericTypeParameter

`func (o *TypeInfo) HasIsGenericTypeParameter() bool`

HasIsGenericTypeParameter returns a boolean if a field has been set.

### GetIsGenericMethodParameter

`func (o *TypeInfo) GetIsGenericMethodParameter() bool`

GetIsGenericMethodParameter returns the IsGenericMethodParameter field if non-nil, zero value otherwise.

### GetIsGenericMethodParameterOk

`func (o *TypeInfo) GetIsGenericMethodParameterOk() (*bool, bool)`

GetIsGenericMethodParameterOk returns a tuple with the IsGenericMethodParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericMethodParameter

`func (o *TypeInfo) SetIsGenericMethodParameter(v bool)`

SetIsGenericMethodParameter sets IsGenericMethodParameter field to given value.

### HasIsGenericMethodParameter

`func (o *TypeInfo) HasIsGenericMethodParameter() bool`

HasIsGenericMethodParameter returns a boolean if a field has been set.

### GetIsGenericType

`func (o *TypeInfo) GetIsGenericType() bool`

GetIsGenericType returns the IsGenericType field if non-nil, zero value otherwise.

### GetIsGenericTypeOk

`func (o *TypeInfo) GetIsGenericTypeOk() (*bool, bool)`

GetIsGenericTypeOk returns a tuple with the IsGenericType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericType

`func (o *TypeInfo) SetIsGenericType(v bool)`

SetIsGenericType sets IsGenericType field to given value.

### HasIsGenericType

`func (o *TypeInfo) HasIsGenericType() bool`

HasIsGenericType returns a boolean if a field has been set.

### GetIsGenericTypeDefinition

`func (o *TypeInfo) GetIsGenericTypeDefinition() bool`

GetIsGenericTypeDefinition returns the IsGenericTypeDefinition field if non-nil, zero value otherwise.

### GetIsGenericTypeDefinitionOk

`func (o *TypeInfo) GetIsGenericTypeDefinitionOk() (*bool, bool)`

GetIsGenericTypeDefinitionOk returns a tuple with the IsGenericTypeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGenericTypeDefinition

`func (o *TypeInfo) SetIsGenericTypeDefinition(v bool)`

SetIsGenericTypeDefinition sets IsGenericTypeDefinition field to given value.

### HasIsGenericTypeDefinition

`func (o *TypeInfo) HasIsGenericTypeDefinition() bool`

HasIsGenericTypeDefinition returns a boolean if a field has been set.

### GetIsSZArray

`func (o *TypeInfo) GetIsSZArray() bool`

GetIsSZArray returns the IsSZArray field if non-nil, zero value otherwise.

### GetIsSZArrayOk

`func (o *TypeInfo) GetIsSZArrayOk() (*bool, bool)`

GetIsSZArrayOk returns a tuple with the IsSZArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSZArray

`func (o *TypeInfo) SetIsSZArray(v bool)`

SetIsSZArray sets IsSZArray field to given value.

### HasIsSZArray

`func (o *TypeInfo) HasIsSZArray() bool`

HasIsSZArray returns a boolean if a field has been set.

### GetIsVariableBoundArray

`func (o *TypeInfo) GetIsVariableBoundArray() bool`

GetIsVariableBoundArray returns the IsVariableBoundArray field if non-nil, zero value otherwise.

### GetIsVariableBoundArrayOk

`func (o *TypeInfo) GetIsVariableBoundArrayOk() (*bool, bool)`

GetIsVariableBoundArrayOk returns a tuple with the IsVariableBoundArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariableBoundArray

`func (o *TypeInfo) SetIsVariableBoundArray(v bool)`

SetIsVariableBoundArray sets IsVariableBoundArray field to given value.

### HasIsVariableBoundArray

`func (o *TypeInfo) HasIsVariableBoundArray() bool`

HasIsVariableBoundArray returns a boolean if a field has been set.

### GetIsByRefLike

`func (o *TypeInfo) GetIsByRefLike() bool`

GetIsByRefLike returns the IsByRefLike field if non-nil, zero value otherwise.

### GetIsByRefLikeOk

`func (o *TypeInfo) GetIsByRefLikeOk() (*bool, bool)`

GetIsByRefLikeOk returns a tuple with the IsByRefLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsByRefLike

`func (o *TypeInfo) SetIsByRefLike(v bool)`

SetIsByRefLike sets IsByRefLike field to given value.

### HasIsByRefLike

`func (o *TypeInfo) HasIsByRefLike() bool`

HasIsByRefLike returns a boolean if a field has been set.

### GetIsFunctionPointer

`func (o *TypeInfo) GetIsFunctionPointer() bool`

GetIsFunctionPointer returns the IsFunctionPointer field if non-nil, zero value otherwise.

### GetIsFunctionPointerOk

`func (o *TypeInfo) GetIsFunctionPointerOk() (*bool, bool)`

GetIsFunctionPointerOk returns a tuple with the IsFunctionPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFunctionPointer

`func (o *TypeInfo) SetIsFunctionPointer(v bool)`

SetIsFunctionPointer sets IsFunctionPointer field to given value.

### HasIsFunctionPointer

`func (o *TypeInfo) HasIsFunctionPointer() bool`

HasIsFunctionPointer returns a boolean if a field has been set.

### GetIsUnmanagedFunctionPointer

`func (o *TypeInfo) GetIsUnmanagedFunctionPointer() bool`

GetIsUnmanagedFunctionPointer returns the IsUnmanagedFunctionPointer field if non-nil, zero value otherwise.

### GetIsUnmanagedFunctionPointerOk

`func (o *TypeInfo) GetIsUnmanagedFunctionPointerOk() (*bool, bool)`

GetIsUnmanagedFunctionPointerOk returns a tuple with the IsUnmanagedFunctionPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnmanagedFunctionPointer

`func (o *TypeInfo) SetIsUnmanagedFunctionPointer(v bool)`

SetIsUnmanagedFunctionPointer sets IsUnmanagedFunctionPointer field to given value.

### HasIsUnmanagedFunctionPointer

`func (o *TypeInfo) HasIsUnmanagedFunctionPointer() bool`

HasIsUnmanagedFunctionPointer returns a boolean if a field has been set.

### GetHasElementType

`func (o *TypeInfo) GetHasElementType() bool`

GetHasElementType returns the HasElementType field if non-nil, zero value otherwise.

### GetHasElementTypeOk

`func (o *TypeInfo) GetHasElementTypeOk() (*bool, bool)`

GetHasElementTypeOk returns a tuple with the HasElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasElementType

`func (o *TypeInfo) SetHasElementType(v bool)`

SetHasElementType sets HasElementType field to given value.

### HasHasElementType

`func (o *TypeInfo) HasHasElementType() bool`

HasHasElementType returns a boolean if a field has been set.

### GetGenericTypeArguments

`func (o *TypeInfo) GetGenericTypeArguments() []Type`

GetGenericTypeArguments returns the GenericTypeArguments field if non-nil, zero value otherwise.

### GetGenericTypeArgumentsOk

`func (o *TypeInfo) GetGenericTypeArgumentsOk() (*[]Type, bool)`

GetGenericTypeArgumentsOk returns a tuple with the GenericTypeArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypeArguments

`func (o *TypeInfo) SetGenericTypeArguments(v []Type)`

SetGenericTypeArguments sets GenericTypeArguments field to given value.

### HasGenericTypeArguments

`func (o *TypeInfo) HasGenericTypeArguments() bool`

HasGenericTypeArguments returns a boolean if a field has been set.

### SetGenericTypeArgumentsNil

`func (o *TypeInfo) SetGenericTypeArgumentsNil(b bool)`

 SetGenericTypeArgumentsNil sets the value for GenericTypeArguments to be an explicit nil

### UnsetGenericTypeArguments
`func (o *TypeInfo) UnsetGenericTypeArguments()`

UnsetGenericTypeArguments ensures that no value is present for GenericTypeArguments, not even an explicit nil
### GetGenericParameterPosition

`func (o *TypeInfo) GetGenericParameterPosition() int32`

GetGenericParameterPosition returns the GenericParameterPosition field if non-nil, zero value otherwise.

### GetGenericParameterPositionOk

`func (o *TypeInfo) GetGenericParameterPositionOk() (*int32, bool)`

GetGenericParameterPositionOk returns a tuple with the GenericParameterPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericParameterPosition

`func (o *TypeInfo) SetGenericParameterPosition(v int32)`

SetGenericParameterPosition sets GenericParameterPosition field to given value.

### HasGenericParameterPosition

`func (o *TypeInfo) HasGenericParameterPosition() bool`

HasGenericParameterPosition returns a boolean if a field has been set.

### GetGenericParameterAttributes

`func (o *TypeInfo) GetGenericParameterAttributes() GenericParameterAttributes`

GetGenericParameterAttributes returns the GenericParameterAttributes field if non-nil, zero value otherwise.

### GetGenericParameterAttributesOk

`func (o *TypeInfo) GetGenericParameterAttributesOk() (*GenericParameterAttributes, bool)`

GetGenericParameterAttributesOk returns a tuple with the GenericParameterAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericParameterAttributes

`func (o *TypeInfo) SetGenericParameterAttributes(v GenericParameterAttributes)`

SetGenericParameterAttributes sets GenericParameterAttributes field to given value.

### HasGenericParameterAttributes

`func (o *TypeInfo) HasGenericParameterAttributes() bool`

HasGenericParameterAttributes returns a boolean if a field has been set.

### GetAttributes

`func (o *TypeInfo) GetAttributes() TypeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TypeInfo) GetAttributesOk() (*TypeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TypeInfo) SetAttributes(v TypeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TypeInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsAbstract

`func (o *TypeInfo) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *TypeInfo) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *TypeInfo) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *TypeInfo) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsImport

`func (o *TypeInfo) GetIsImport() bool`

GetIsImport returns the IsImport field if non-nil, zero value otherwise.

### GetIsImportOk

`func (o *TypeInfo) GetIsImportOk() (*bool, bool)`

GetIsImportOk returns a tuple with the IsImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImport

`func (o *TypeInfo) SetIsImport(v bool)`

SetIsImport sets IsImport field to given value.

### HasIsImport

`func (o *TypeInfo) HasIsImport() bool`

HasIsImport returns a boolean if a field has been set.

### GetIsSealed

`func (o *TypeInfo) GetIsSealed() bool`

GetIsSealed returns the IsSealed field if non-nil, zero value otherwise.

### GetIsSealedOk

`func (o *TypeInfo) GetIsSealedOk() (*bool, bool)`

GetIsSealedOk returns a tuple with the IsSealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSealed

`func (o *TypeInfo) SetIsSealed(v bool)`

SetIsSealed sets IsSealed field to given value.

### HasIsSealed

`func (o *TypeInfo) HasIsSealed() bool`

HasIsSealed returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *TypeInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *TypeInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *TypeInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *TypeInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetIsClass

`func (o *TypeInfo) GetIsClass() bool`

GetIsClass returns the IsClass field if non-nil, zero value otherwise.

### GetIsClassOk

`func (o *TypeInfo) GetIsClassOk() (*bool, bool)`

GetIsClassOk returns a tuple with the IsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClass

`func (o *TypeInfo) SetIsClass(v bool)`

SetIsClass sets IsClass field to given value.

### HasIsClass

`func (o *TypeInfo) HasIsClass() bool`

HasIsClass returns a boolean if a field has been set.

### GetIsNestedAssembly

`func (o *TypeInfo) GetIsNestedAssembly() bool`

GetIsNestedAssembly returns the IsNestedAssembly field if non-nil, zero value otherwise.

### GetIsNestedAssemblyOk

`func (o *TypeInfo) GetIsNestedAssemblyOk() (*bool, bool)`

GetIsNestedAssemblyOk returns a tuple with the IsNestedAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedAssembly

`func (o *TypeInfo) SetIsNestedAssembly(v bool)`

SetIsNestedAssembly sets IsNestedAssembly field to given value.

### HasIsNestedAssembly

`func (o *TypeInfo) HasIsNestedAssembly() bool`

HasIsNestedAssembly returns a boolean if a field has been set.

### GetIsNestedFamANDAssem

`func (o *TypeInfo) GetIsNestedFamANDAssem() bool`

GetIsNestedFamANDAssem returns the IsNestedFamANDAssem field if non-nil, zero value otherwise.

### GetIsNestedFamANDAssemOk

`func (o *TypeInfo) GetIsNestedFamANDAssemOk() (*bool, bool)`

GetIsNestedFamANDAssemOk returns a tuple with the IsNestedFamANDAssem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamANDAssem

`func (o *TypeInfo) SetIsNestedFamANDAssem(v bool)`

SetIsNestedFamANDAssem sets IsNestedFamANDAssem field to given value.

### HasIsNestedFamANDAssem

`func (o *TypeInfo) HasIsNestedFamANDAssem() bool`

HasIsNestedFamANDAssem returns a boolean if a field has been set.

### GetIsNestedFamily

`func (o *TypeInfo) GetIsNestedFamily() bool`

GetIsNestedFamily returns the IsNestedFamily field if non-nil, zero value otherwise.

### GetIsNestedFamilyOk

`func (o *TypeInfo) GetIsNestedFamilyOk() (*bool, bool)`

GetIsNestedFamilyOk returns a tuple with the IsNestedFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamily

`func (o *TypeInfo) SetIsNestedFamily(v bool)`

SetIsNestedFamily sets IsNestedFamily field to given value.

### HasIsNestedFamily

`func (o *TypeInfo) HasIsNestedFamily() bool`

HasIsNestedFamily returns a boolean if a field has been set.

### GetIsNestedFamORAssem

`func (o *TypeInfo) GetIsNestedFamORAssem() bool`

GetIsNestedFamORAssem returns the IsNestedFamORAssem field if non-nil, zero value otherwise.

### GetIsNestedFamORAssemOk

`func (o *TypeInfo) GetIsNestedFamORAssemOk() (*bool, bool)`

GetIsNestedFamORAssemOk returns a tuple with the IsNestedFamORAssem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedFamORAssem

`func (o *TypeInfo) SetIsNestedFamORAssem(v bool)`

SetIsNestedFamORAssem sets IsNestedFamORAssem field to given value.

### HasIsNestedFamORAssem

`func (o *TypeInfo) HasIsNestedFamORAssem() bool`

HasIsNestedFamORAssem returns a boolean if a field has been set.

### GetIsNestedPrivate

`func (o *TypeInfo) GetIsNestedPrivate() bool`

GetIsNestedPrivate returns the IsNestedPrivate field if non-nil, zero value otherwise.

### GetIsNestedPrivateOk

`func (o *TypeInfo) GetIsNestedPrivateOk() (*bool, bool)`

GetIsNestedPrivateOk returns a tuple with the IsNestedPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedPrivate

`func (o *TypeInfo) SetIsNestedPrivate(v bool)`

SetIsNestedPrivate sets IsNestedPrivate field to given value.

### HasIsNestedPrivate

`func (o *TypeInfo) HasIsNestedPrivate() bool`

HasIsNestedPrivate returns a boolean if a field has been set.

### GetIsNestedPublic

`func (o *TypeInfo) GetIsNestedPublic() bool`

GetIsNestedPublic returns the IsNestedPublic field if non-nil, zero value otherwise.

### GetIsNestedPublicOk

`func (o *TypeInfo) GetIsNestedPublicOk() (*bool, bool)`

GetIsNestedPublicOk returns a tuple with the IsNestedPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNestedPublic

`func (o *TypeInfo) SetIsNestedPublic(v bool)`

SetIsNestedPublic sets IsNestedPublic field to given value.

### HasIsNestedPublic

`func (o *TypeInfo) HasIsNestedPublic() bool`

HasIsNestedPublic returns a boolean if a field has been set.

### GetIsNotPublic

`func (o *TypeInfo) GetIsNotPublic() bool`

GetIsNotPublic returns the IsNotPublic field if non-nil, zero value otherwise.

### GetIsNotPublicOk

`func (o *TypeInfo) GetIsNotPublicOk() (*bool, bool)`

GetIsNotPublicOk returns a tuple with the IsNotPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotPublic

`func (o *TypeInfo) SetIsNotPublic(v bool)`

SetIsNotPublic sets IsNotPublic field to given value.

### HasIsNotPublic

`func (o *TypeInfo) HasIsNotPublic() bool`

HasIsNotPublic returns a boolean if a field has been set.

### GetIsPublic

`func (o *TypeInfo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *TypeInfo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *TypeInfo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *TypeInfo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsAutoLayout

`func (o *TypeInfo) GetIsAutoLayout() bool`

GetIsAutoLayout returns the IsAutoLayout field if non-nil, zero value otherwise.

### GetIsAutoLayoutOk

`func (o *TypeInfo) GetIsAutoLayoutOk() (*bool, bool)`

GetIsAutoLayoutOk returns a tuple with the IsAutoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoLayout

`func (o *TypeInfo) SetIsAutoLayout(v bool)`

SetIsAutoLayout sets IsAutoLayout field to given value.

### HasIsAutoLayout

`func (o *TypeInfo) HasIsAutoLayout() bool`

HasIsAutoLayout returns a boolean if a field has been set.

### GetIsExplicitLayout

`func (o *TypeInfo) GetIsExplicitLayout() bool`

GetIsExplicitLayout returns the IsExplicitLayout field if non-nil, zero value otherwise.

### GetIsExplicitLayoutOk

`func (o *TypeInfo) GetIsExplicitLayoutOk() (*bool, bool)`

GetIsExplicitLayoutOk returns a tuple with the IsExplicitLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExplicitLayout

`func (o *TypeInfo) SetIsExplicitLayout(v bool)`

SetIsExplicitLayout sets IsExplicitLayout field to given value.

### HasIsExplicitLayout

`func (o *TypeInfo) HasIsExplicitLayout() bool`

HasIsExplicitLayout returns a boolean if a field has been set.

### GetIsLayoutSequential

`func (o *TypeInfo) GetIsLayoutSequential() bool`

GetIsLayoutSequential returns the IsLayoutSequential field if non-nil, zero value otherwise.

### GetIsLayoutSequentialOk

`func (o *TypeInfo) GetIsLayoutSequentialOk() (*bool, bool)`

GetIsLayoutSequentialOk returns a tuple with the IsLayoutSequential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLayoutSequential

`func (o *TypeInfo) SetIsLayoutSequential(v bool)`

SetIsLayoutSequential sets IsLayoutSequential field to given value.

### HasIsLayoutSequential

`func (o *TypeInfo) HasIsLayoutSequential() bool`

HasIsLayoutSequential returns a boolean if a field has been set.

### GetIsAnsiClass

`func (o *TypeInfo) GetIsAnsiClass() bool`

GetIsAnsiClass returns the IsAnsiClass field if non-nil, zero value otherwise.

### GetIsAnsiClassOk

`func (o *TypeInfo) GetIsAnsiClassOk() (*bool, bool)`

GetIsAnsiClassOk returns a tuple with the IsAnsiClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnsiClass

`func (o *TypeInfo) SetIsAnsiClass(v bool)`

SetIsAnsiClass sets IsAnsiClass field to given value.

### HasIsAnsiClass

`func (o *TypeInfo) HasIsAnsiClass() bool`

HasIsAnsiClass returns a boolean if a field has been set.

### GetIsAutoClass

`func (o *TypeInfo) GetIsAutoClass() bool`

GetIsAutoClass returns the IsAutoClass field if non-nil, zero value otherwise.

### GetIsAutoClassOk

`func (o *TypeInfo) GetIsAutoClassOk() (*bool, bool)`

GetIsAutoClassOk returns a tuple with the IsAutoClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoClass

`func (o *TypeInfo) SetIsAutoClass(v bool)`

SetIsAutoClass sets IsAutoClass field to given value.

### HasIsAutoClass

`func (o *TypeInfo) HasIsAutoClass() bool`

HasIsAutoClass returns a boolean if a field has been set.

### GetIsUnicodeClass

`func (o *TypeInfo) GetIsUnicodeClass() bool`

GetIsUnicodeClass returns the IsUnicodeClass field if non-nil, zero value otherwise.

### GetIsUnicodeClassOk

`func (o *TypeInfo) GetIsUnicodeClassOk() (*bool, bool)`

GetIsUnicodeClassOk returns a tuple with the IsUnicodeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnicodeClass

`func (o *TypeInfo) SetIsUnicodeClass(v bool)`

SetIsUnicodeClass sets IsUnicodeClass field to given value.

### HasIsUnicodeClass

`func (o *TypeInfo) HasIsUnicodeClass() bool`

HasIsUnicodeClass returns a boolean if a field has been set.

### GetIsCOMObject

`func (o *TypeInfo) GetIsCOMObject() bool`

GetIsCOMObject returns the IsCOMObject field if non-nil, zero value otherwise.

### GetIsCOMObjectOk

`func (o *TypeInfo) GetIsCOMObjectOk() (*bool, bool)`

GetIsCOMObjectOk returns a tuple with the IsCOMObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCOMObject

`func (o *TypeInfo) SetIsCOMObject(v bool)`

SetIsCOMObject sets IsCOMObject field to given value.

### HasIsCOMObject

`func (o *TypeInfo) HasIsCOMObject() bool`

HasIsCOMObject returns a boolean if a field has been set.

### GetIsContextful

`func (o *TypeInfo) GetIsContextful() bool`

GetIsContextful returns the IsContextful field if non-nil, zero value otherwise.

### GetIsContextfulOk

`func (o *TypeInfo) GetIsContextfulOk() (*bool, bool)`

GetIsContextfulOk returns a tuple with the IsContextful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContextful

`func (o *TypeInfo) SetIsContextful(v bool)`

SetIsContextful sets IsContextful field to given value.

### HasIsContextful

`func (o *TypeInfo) HasIsContextful() bool`

HasIsContextful returns a boolean if a field has been set.

### GetIsEnum

`func (o *TypeInfo) GetIsEnum() bool`

GetIsEnum returns the IsEnum field if non-nil, zero value otherwise.

### GetIsEnumOk

`func (o *TypeInfo) GetIsEnumOk() (*bool, bool)`

GetIsEnumOk returns a tuple with the IsEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnum

`func (o *TypeInfo) SetIsEnum(v bool)`

SetIsEnum sets IsEnum field to given value.

### HasIsEnum

`func (o *TypeInfo) HasIsEnum() bool`

HasIsEnum returns a boolean if a field has been set.

### GetIsMarshalByRef

`func (o *TypeInfo) GetIsMarshalByRef() bool`

GetIsMarshalByRef returns the IsMarshalByRef field if non-nil, zero value otherwise.

### GetIsMarshalByRefOk

`func (o *TypeInfo) GetIsMarshalByRefOk() (*bool, bool)`

GetIsMarshalByRefOk returns a tuple with the IsMarshalByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarshalByRef

`func (o *TypeInfo) SetIsMarshalByRef(v bool)`

SetIsMarshalByRef sets IsMarshalByRef field to given value.

### HasIsMarshalByRef

`func (o *TypeInfo) HasIsMarshalByRef() bool`

HasIsMarshalByRef returns a boolean if a field has been set.

### GetIsPrimitive

`func (o *TypeInfo) GetIsPrimitive() bool`

GetIsPrimitive returns the IsPrimitive field if non-nil, zero value otherwise.

### GetIsPrimitiveOk

`func (o *TypeInfo) GetIsPrimitiveOk() (*bool, bool)`

GetIsPrimitiveOk returns a tuple with the IsPrimitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimitive

`func (o *TypeInfo) SetIsPrimitive(v bool)`

SetIsPrimitive sets IsPrimitive field to given value.

### HasIsPrimitive

`func (o *TypeInfo) HasIsPrimitive() bool`

HasIsPrimitive returns a boolean if a field has been set.

### GetIsValueType

`func (o *TypeInfo) GetIsValueType() bool`

GetIsValueType returns the IsValueType field if non-nil, zero value otherwise.

### GetIsValueTypeOk

`func (o *TypeInfo) GetIsValueTypeOk() (*bool, bool)`

GetIsValueTypeOk returns a tuple with the IsValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueType

`func (o *TypeInfo) SetIsValueType(v bool)`

SetIsValueType sets IsValueType field to given value.

### HasIsValueType

`func (o *TypeInfo) HasIsValueType() bool`

HasIsValueType returns a boolean if a field has been set.

### GetIsSignatureType

`func (o *TypeInfo) GetIsSignatureType() bool`

GetIsSignatureType returns the IsSignatureType field if non-nil, zero value otherwise.

### GetIsSignatureTypeOk

`func (o *TypeInfo) GetIsSignatureTypeOk() (*bool, bool)`

GetIsSignatureTypeOk returns a tuple with the IsSignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSignatureType

`func (o *TypeInfo) SetIsSignatureType(v bool)`

SetIsSignatureType sets IsSignatureType field to given value.

### HasIsSignatureType

`func (o *TypeInfo) HasIsSignatureType() bool`

HasIsSignatureType returns a boolean if a field has been set.

### GetIsSecurityCritical

`func (o *TypeInfo) GetIsSecurityCritical() bool`

GetIsSecurityCritical returns the IsSecurityCritical field if non-nil, zero value otherwise.

### GetIsSecurityCriticalOk

`func (o *TypeInfo) GetIsSecurityCriticalOk() (*bool, bool)`

GetIsSecurityCriticalOk returns a tuple with the IsSecurityCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityCritical

`func (o *TypeInfo) SetIsSecurityCritical(v bool)`

SetIsSecurityCritical sets IsSecurityCritical field to given value.

### HasIsSecurityCritical

`func (o *TypeInfo) HasIsSecurityCritical() bool`

HasIsSecurityCritical returns a boolean if a field has been set.

### GetIsSecuritySafeCritical

`func (o *TypeInfo) GetIsSecuritySafeCritical() bool`

GetIsSecuritySafeCritical returns the IsSecuritySafeCritical field if non-nil, zero value otherwise.

### GetIsSecuritySafeCriticalOk

`func (o *TypeInfo) GetIsSecuritySafeCriticalOk() (*bool, bool)`

GetIsSecuritySafeCriticalOk returns a tuple with the IsSecuritySafeCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySafeCritical

`func (o *TypeInfo) SetIsSecuritySafeCritical(v bool)`

SetIsSecuritySafeCritical sets IsSecuritySafeCritical field to given value.

### HasIsSecuritySafeCritical

`func (o *TypeInfo) HasIsSecuritySafeCritical() bool`

HasIsSecuritySafeCritical returns a boolean if a field has been set.

### GetIsSecurityTransparent

`func (o *TypeInfo) GetIsSecurityTransparent() bool`

GetIsSecurityTransparent returns the IsSecurityTransparent field if non-nil, zero value otherwise.

### GetIsSecurityTransparentOk

`func (o *TypeInfo) GetIsSecurityTransparentOk() (*bool, bool)`

GetIsSecurityTransparentOk returns a tuple with the IsSecurityTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityTransparent

`func (o *TypeInfo) SetIsSecurityTransparent(v bool)`

SetIsSecurityTransparent sets IsSecurityTransparent field to given value.

### HasIsSecurityTransparent

`func (o *TypeInfo) HasIsSecurityTransparent() bool`

HasIsSecurityTransparent returns a boolean if a field has been set.

### GetStructLayoutAttribute

`func (o *TypeInfo) GetStructLayoutAttribute() StructLayoutAttribute`

GetStructLayoutAttribute returns the StructLayoutAttribute field if non-nil, zero value otherwise.

### GetStructLayoutAttributeOk

`func (o *TypeInfo) GetStructLayoutAttributeOk() (*StructLayoutAttribute, bool)`

GetStructLayoutAttributeOk returns a tuple with the StructLayoutAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructLayoutAttribute

`func (o *TypeInfo) SetStructLayoutAttribute(v StructLayoutAttribute)`

SetStructLayoutAttribute sets StructLayoutAttribute field to given value.

### HasStructLayoutAttribute

`func (o *TypeInfo) HasStructLayoutAttribute() bool`

HasStructLayoutAttribute returns a boolean if a field has been set.

### GetTypeInitializer

`func (o *TypeInfo) GetTypeInitializer() ConstructorInfo`

GetTypeInitializer returns the TypeInitializer field if non-nil, zero value otherwise.

### GetTypeInitializerOk

`func (o *TypeInfo) GetTypeInitializerOk() (*ConstructorInfo, bool)`

GetTypeInitializerOk returns a tuple with the TypeInitializer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInitializer

`func (o *TypeInfo) SetTypeInitializer(v ConstructorInfo)`

SetTypeInitializer sets TypeInitializer field to given value.

### HasTypeInitializer

`func (o *TypeInfo) HasTypeInitializer() bool`

HasTypeInitializer returns a boolean if a field has been set.

### GetTypeHandle

`func (o *TypeInfo) GetTypeHandle() RuntimeTypeHandle`

GetTypeHandle returns the TypeHandle field if non-nil, zero value otherwise.

### GetTypeHandleOk

`func (o *TypeInfo) GetTypeHandleOk() (*RuntimeTypeHandle, bool)`

GetTypeHandleOk returns a tuple with the TypeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeHandle

`func (o *TypeInfo) SetTypeHandle(v RuntimeTypeHandle)`

SetTypeHandle sets TypeHandle field to given value.

### HasTypeHandle

`func (o *TypeInfo) HasTypeHandle() bool`

HasTypeHandle returns a boolean if a field has been set.

### GetGuid

`func (o *TypeInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *TypeInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *TypeInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *TypeInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetBaseType

`func (o *TypeInfo) GetBaseType() Type`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TypeInfo) GetBaseTypeOk() (*Type, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TypeInfo) SetBaseType(v Type)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TypeInfo) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetIsSerializable

`func (o *TypeInfo) GetIsSerializable() bool`

GetIsSerializable returns the IsSerializable field if non-nil, zero value otherwise.

### GetIsSerializableOk

`func (o *TypeInfo) GetIsSerializableOk() (*bool, bool)`

GetIsSerializableOk returns a tuple with the IsSerializable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSerializable

`func (o *TypeInfo) SetIsSerializable(v bool)`

SetIsSerializable sets IsSerializable field to given value.

### HasIsSerializable

`func (o *TypeInfo) HasIsSerializable() bool`

HasIsSerializable returns a boolean if a field has been set.

### GetContainsGenericParameters

`func (o *TypeInfo) GetContainsGenericParameters() bool`

GetContainsGenericParameters returns the ContainsGenericParameters field if non-nil, zero value otherwise.

### GetContainsGenericParametersOk

`func (o *TypeInfo) GetContainsGenericParametersOk() (*bool, bool)`

GetContainsGenericParametersOk returns a tuple with the ContainsGenericParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsGenericParameters

`func (o *TypeInfo) SetContainsGenericParameters(v bool)`

SetContainsGenericParameters sets ContainsGenericParameters field to given value.

### HasContainsGenericParameters

`func (o *TypeInfo) HasContainsGenericParameters() bool`

HasContainsGenericParameters returns a boolean if a field has been set.

### GetIsVisible

`func (o *TypeInfo) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *TypeInfo) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *TypeInfo) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *TypeInfo) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetGenericTypeParameters

`func (o *TypeInfo) GetGenericTypeParameters() []Type`

GetGenericTypeParameters returns the GenericTypeParameters field if non-nil, zero value otherwise.

### GetGenericTypeParametersOk

`func (o *TypeInfo) GetGenericTypeParametersOk() (*[]Type, bool)`

GetGenericTypeParametersOk returns a tuple with the GenericTypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypeParameters

`func (o *TypeInfo) SetGenericTypeParameters(v []Type)`

SetGenericTypeParameters sets GenericTypeParameters field to given value.

### HasGenericTypeParameters

`func (o *TypeInfo) HasGenericTypeParameters() bool`

HasGenericTypeParameters returns a boolean if a field has been set.

### SetGenericTypeParametersNil

`func (o *TypeInfo) SetGenericTypeParametersNil(b bool)`

 SetGenericTypeParametersNil sets the value for GenericTypeParameters to be an explicit nil

### UnsetGenericTypeParameters
`func (o *TypeInfo) UnsetGenericTypeParameters()`

UnsetGenericTypeParameters ensures that no value is present for GenericTypeParameters, not even an explicit nil
### GetDeclaredConstructors

`func (o *TypeInfo) GetDeclaredConstructors() []ConstructorInfo`

GetDeclaredConstructors returns the DeclaredConstructors field if non-nil, zero value otherwise.

### GetDeclaredConstructorsOk

`func (o *TypeInfo) GetDeclaredConstructorsOk() (*[]ConstructorInfo, bool)`

GetDeclaredConstructorsOk returns a tuple with the DeclaredConstructors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredConstructors

`func (o *TypeInfo) SetDeclaredConstructors(v []ConstructorInfo)`

SetDeclaredConstructors sets DeclaredConstructors field to given value.

### HasDeclaredConstructors

`func (o *TypeInfo) HasDeclaredConstructors() bool`

HasDeclaredConstructors returns a boolean if a field has been set.

### SetDeclaredConstructorsNil

`func (o *TypeInfo) SetDeclaredConstructorsNil(b bool)`

 SetDeclaredConstructorsNil sets the value for DeclaredConstructors to be an explicit nil

### UnsetDeclaredConstructors
`func (o *TypeInfo) UnsetDeclaredConstructors()`

UnsetDeclaredConstructors ensures that no value is present for DeclaredConstructors, not even an explicit nil
### GetDeclaredEvents

`func (o *TypeInfo) GetDeclaredEvents() []EventInfo`

GetDeclaredEvents returns the DeclaredEvents field if non-nil, zero value otherwise.

### GetDeclaredEventsOk

`func (o *TypeInfo) GetDeclaredEventsOk() (*[]EventInfo, bool)`

GetDeclaredEventsOk returns a tuple with the DeclaredEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredEvents

`func (o *TypeInfo) SetDeclaredEvents(v []EventInfo)`

SetDeclaredEvents sets DeclaredEvents field to given value.

### HasDeclaredEvents

`func (o *TypeInfo) HasDeclaredEvents() bool`

HasDeclaredEvents returns a boolean if a field has been set.

### SetDeclaredEventsNil

`func (o *TypeInfo) SetDeclaredEventsNil(b bool)`

 SetDeclaredEventsNil sets the value for DeclaredEvents to be an explicit nil

### UnsetDeclaredEvents
`func (o *TypeInfo) UnsetDeclaredEvents()`

UnsetDeclaredEvents ensures that no value is present for DeclaredEvents, not even an explicit nil
### GetDeclaredFields

`func (o *TypeInfo) GetDeclaredFields() []FieldInfo`

GetDeclaredFields returns the DeclaredFields field if non-nil, zero value otherwise.

### GetDeclaredFieldsOk

`func (o *TypeInfo) GetDeclaredFieldsOk() (*[]FieldInfo, bool)`

GetDeclaredFieldsOk returns a tuple with the DeclaredFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredFields

`func (o *TypeInfo) SetDeclaredFields(v []FieldInfo)`

SetDeclaredFields sets DeclaredFields field to given value.

### HasDeclaredFields

`func (o *TypeInfo) HasDeclaredFields() bool`

HasDeclaredFields returns a boolean if a field has been set.

### SetDeclaredFieldsNil

`func (o *TypeInfo) SetDeclaredFieldsNil(b bool)`

 SetDeclaredFieldsNil sets the value for DeclaredFields to be an explicit nil

### UnsetDeclaredFields
`func (o *TypeInfo) UnsetDeclaredFields()`

UnsetDeclaredFields ensures that no value is present for DeclaredFields, not even an explicit nil
### GetDeclaredMembers

`func (o *TypeInfo) GetDeclaredMembers() []MemberInfo`

GetDeclaredMembers returns the DeclaredMembers field if non-nil, zero value otherwise.

### GetDeclaredMembersOk

`func (o *TypeInfo) GetDeclaredMembersOk() (*[]MemberInfo, bool)`

GetDeclaredMembersOk returns a tuple with the DeclaredMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMembers

`func (o *TypeInfo) SetDeclaredMembers(v []MemberInfo)`

SetDeclaredMembers sets DeclaredMembers field to given value.

### HasDeclaredMembers

`func (o *TypeInfo) HasDeclaredMembers() bool`

HasDeclaredMembers returns a boolean if a field has been set.

### SetDeclaredMembersNil

`func (o *TypeInfo) SetDeclaredMembersNil(b bool)`

 SetDeclaredMembersNil sets the value for DeclaredMembers to be an explicit nil

### UnsetDeclaredMembers
`func (o *TypeInfo) UnsetDeclaredMembers()`

UnsetDeclaredMembers ensures that no value is present for DeclaredMembers, not even an explicit nil
### GetDeclaredMethods

`func (o *TypeInfo) GetDeclaredMethods() []MethodInfo`

GetDeclaredMethods returns the DeclaredMethods field if non-nil, zero value otherwise.

### GetDeclaredMethodsOk

`func (o *TypeInfo) GetDeclaredMethodsOk() (*[]MethodInfo, bool)`

GetDeclaredMethodsOk returns a tuple with the DeclaredMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredMethods

`func (o *TypeInfo) SetDeclaredMethods(v []MethodInfo)`

SetDeclaredMethods sets DeclaredMethods field to given value.

### HasDeclaredMethods

`func (o *TypeInfo) HasDeclaredMethods() bool`

HasDeclaredMethods returns a boolean if a field has been set.

### SetDeclaredMethodsNil

`func (o *TypeInfo) SetDeclaredMethodsNil(b bool)`

 SetDeclaredMethodsNil sets the value for DeclaredMethods to be an explicit nil

### UnsetDeclaredMethods
`func (o *TypeInfo) UnsetDeclaredMethods()`

UnsetDeclaredMethods ensures that no value is present for DeclaredMethods, not even an explicit nil
### GetDeclaredNestedTypes

`func (o *TypeInfo) GetDeclaredNestedTypes() []TypeInfo`

GetDeclaredNestedTypes returns the DeclaredNestedTypes field if non-nil, zero value otherwise.

### GetDeclaredNestedTypesOk

`func (o *TypeInfo) GetDeclaredNestedTypesOk() (*[]TypeInfo, bool)`

GetDeclaredNestedTypesOk returns a tuple with the DeclaredNestedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredNestedTypes

`func (o *TypeInfo) SetDeclaredNestedTypes(v []TypeInfo)`

SetDeclaredNestedTypes sets DeclaredNestedTypes field to given value.

### HasDeclaredNestedTypes

`func (o *TypeInfo) HasDeclaredNestedTypes() bool`

HasDeclaredNestedTypes returns a boolean if a field has been set.

### SetDeclaredNestedTypesNil

`func (o *TypeInfo) SetDeclaredNestedTypesNil(b bool)`

 SetDeclaredNestedTypesNil sets the value for DeclaredNestedTypes to be an explicit nil

### UnsetDeclaredNestedTypes
`func (o *TypeInfo) UnsetDeclaredNestedTypes()`

UnsetDeclaredNestedTypes ensures that no value is present for DeclaredNestedTypes, not even an explicit nil
### GetDeclaredProperties

`func (o *TypeInfo) GetDeclaredProperties() []PropertyInfo`

GetDeclaredProperties returns the DeclaredProperties field if non-nil, zero value otherwise.

### GetDeclaredPropertiesOk

`func (o *TypeInfo) GetDeclaredPropertiesOk() (*[]PropertyInfo, bool)`

GetDeclaredPropertiesOk returns a tuple with the DeclaredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredProperties

`func (o *TypeInfo) SetDeclaredProperties(v []PropertyInfo)`

SetDeclaredProperties sets DeclaredProperties field to given value.

### HasDeclaredProperties

`func (o *TypeInfo) HasDeclaredProperties() bool`

HasDeclaredProperties returns a boolean if a field has been set.

### SetDeclaredPropertiesNil

`func (o *TypeInfo) SetDeclaredPropertiesNil(b bool)`

 SetDeclaredPropertiesNil sets the value for DeclaredProperties to be an explicit nil

### UnsetDeclaredProperties
`func (o *TypeInfo) UnsetDeclaredProperties()`

UnsetDeclaredProperties ensures that no value is present for DeclaredProperties, not even an explicit nil
### GetImplementedInterfaces

`func (o *TypeInfo) GetImplementedInterfaces() []Type`

GetImplementedInterfaces returns the ImplementedInterfaces field if non-nil, zero value otherwise.

### GetImplementedInterfacesOk

`func (o *TypeInfo) GetImplementedInterfacesOk() (*[]Type, bool)`

GetImplementedInterfacesOk returns a tuple with the ImplementedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedInterfaces

`func (o *TypeInfo) SetImplementedInterfaces(v []Type)`

SetImplementedInterfaces sets ImplementedInterfaces field to given value.

### HasImplementedInterfaces

`func (o *TypeInfo) HasImplementedInterfaces() bool`

HasImplementedInterfaces returns a boolean if a field has been set.

### SetImplementedInterfacesNil

`func (o *TypeInfo) SetImplementedInterfacesNil(b bool)`

 SetImplementedInterfacesNil sets the value for ImplementedInterfaces to be an explicit nil

### UnsetImplementedInterfaces
`func (o *TypeInfo) UnsetImplementedInterfaces()`

UnsetImplementedInterfaces ensures that no value is present for ImplementedInterfaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


