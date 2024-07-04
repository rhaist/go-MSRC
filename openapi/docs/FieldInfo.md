# FieldInfo

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
**MemberType** | Pointer to [**MemberTypes**](MemberTypes.md) |  | [optional] 
**Attributes** | Pointer to [**FieldAttributes**](FieldAttributes.md) |  | [optional] 
**FieldType** | Pointer to [**Type**](Type.md) |  | [optional] 
**IsInitOnly** | Pointer to **bool** |  | [optional] [readonly] 
**IsLiteral** | Pointer to **bool** |  | [optional] [readonly] 
**IsNotSerialized** | Pointer to **bool** |  | [optional] [readonly] 
**IsPinvokeImpl** | Pointer to **bool** |  | [optional] [readonly] 
**IsSpecialName** | Pointer to **bool** |  | [optional] [readonly] 
**IsStatic** | Pointer to **bool** |  | [optional] [readonly] 
**IsAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamily** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamilyAndAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsFamilyOrAssembly** | Pointer to **bool** |  | [optional] [readonly] 
**IsPrivate** | Pointer to **bool** |  | [optional] [readonly] 
**IsPublic** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecuritySafeCritical** | Pointer to **bool** |  | [optional] [readonly] 
**IsSecurityTransparent** | Pointer to **bool** |  | [optional] [readonly] 
**FieldHandle** | Pointer to [**RuntimeFieldHandle**](RuntimeFieldHandle.md) |  | [optional] 

## Methods

### NewFieldInfo

`func NewFieldInfo() *FieldInfo`

NewFieldInfo instantiates a new FieldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldInfoWithDefaults

`func NewFieldInfoWithDefaults() *FieldInfo`

NewFieldInfoWithDefaults instantiates a new FieldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FieldInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FieldInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeclaringType

`func (o *FieldInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *FieldInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *FieldInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *FieldInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetReflectedType

`func (o *FieldInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *FieldInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *FieldInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *FieldInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetModule

`func (o *FieldInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FieldInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FieldInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *FieldInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *FieldInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *FieldInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *FieldInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *FieldInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *FieldInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *FieldInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *FieldInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *FieldInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *FieldInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *FieldInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *FieldInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *FieldInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *FieldInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *FieldInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetMemberType

`func (o *FieldInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *FieldInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *FieldInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *FieldInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetAttributes

`func (o *FieldInfo) GetAttributes() FieldAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FieldInfo) GetAttributesOk() (*FieldAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FieldInfo) SetAttributes(v FieldAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FieldInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetFieldType

`func (o *FieldInfo) GetFieldType() Type`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *FieldInfo) GetFieldTypeOk() (*Type, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *FieldInfo) SetFieldType(v Type)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *FieldInfo) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetIsInitOnly

`func (o *FieldInfo) GetIsInitOnly() bool`

GetIsInitOnly returns the IsInitOnly field if non-nil, zero value otherwise.

### GetIsInitOnlyOk

`func (o *FieldInfo) GetIsInitOnlyOk() (*bool, bool)`

GetIsInitOnlyOk returns a tuple with the IsInitOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitOnly

`func (o *FieldInfo) SetIsInitOnly(v bool)`

SetIsInitOnly sets IsInitOnly field to given value.

### HasIsInitOnly

`func (o *FieldInfo) HasIsInitOnly() bool`

HasIsInitOnly returns a boolean if a field has been set.

### GetIsLiteral

`func (o *FieldInfo) GetIsLiteral() bool`

GetIsLiteral returns the IsLiteral field if non-nil, zero value otherwise.

### GetIsLiteralOk

`func (o *FieldInfo) GetIsLiteralOk() (*bool, bool)`

GetIsLiteralOk returns a tuple with the IsLiteral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLiteral

`func (o *FieldInfo) SetIsLiteral(v bool)`

SetIsLiteral sets IsLiteral field to given value.

### HasIsLiteral

`func (o *FieldInfo) HasIsLiteral() bool`

HasIsLiteral returns a boolean if a field has been set.

### GetIsNotSerialized

`func (o *FieldInfo) GetIsNotSerialized() bool`

GetIsNotSerialized returns the IsNotSerialized field if non-nil, zero value otherwise.

### GetIsNotSerializedOk

`func (o *FieldInfo) GetIsNotSerializedOk() (*bool, bool)`

GetIsNotSerializedOk returns a tuple with the IsNotSerialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotSerialized

`func (o *FieldInfo) SetIsNotSerialized(v bool)`

SetIsNotSerialized sets IsNotSerialized field to given value.

### HasIsNotSerialized

`func (o *FieldInfo) HasIsNotSerialized() bool`

HasIsNotSerialized returns a boolean if a field has been set.

### GetIsPinvokeImpl

`func (o *FieldInfo) GetIsPinvokeImpl() bool`

GetIsPinvokeImpl returns the IsPinvokeImpl field if non-nil, zero value otherwise.

### GetIsPinvokeImplOk

`func (o *FieldInfo) GetIsPinvokeImplOk() (*bool, bool)`

GetIsPinvokeImplOk returns a tuple with the IsPinvokeImpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinvokeImpl

`func (o *FieldInfo) SetIsPinvokeImpl(v bool)`

SetIsPinvokeImpl sets IsPinvokeImpl field to given value.

### HasIsPinvokeImpl

`func (o *FieldInfo) HasIsPinvokeImpl() bool`

HasIsPinvokeImpl returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *FieldInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *FieldInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *FieldInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *FieldInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetIsStatic

`func (o *FieldInfo) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *FieldInfo) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *FieldInfo) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *FieldInfo) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetIsAssembly

`func (o *FieldInfo) GetIsAssembly() bool`

GetIsAssembly returns the IsAssembly field if non-nil, zero value otherwise.

### GetIsAssemblyOk

`func (o *FieldInfo) GetIsAssemblyOk() (*bool, bool)`

GetIsAssemblyOk returns a tuple with the IsAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssembly

`func (o *FieldInfo) SetIsAssembly(v bool)`

SetIsAssembly sets IsAssembly field to given value.

### HasIsAssembly

`func (o *FieldInfo) HasIsAssembly() bool`

HasIsAssembly returns a boolean if a field has been set.

### GetIsFamily

`func (o *FieldInfo) GetIsFamily() bool`

GetIsFamily returns the IsFamily field if non-nil, zero value otherwise.

### GetIsFamilyOk

`func (o *FieldInfo) GetIsFamilyOk() (*bool, bool)`

GetIsFamilyOk returns a tuple with the IsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamily

`func (o *FieldInfo) SetIsFamily(v bool)`

SetIsFamily sets IsFamily field to given value.

### HasIsFamily

`func (o *FieldInfo) HasIsFamily() bool`

HasIsFamily returns a boolean if a field has been set.

### GetIsFamilyAndAssembly

`func (o *FieldInfo) GetIsFamilyAndAssembly() bool`

GetIsFamilyAndAssembly returns the IsFamilyAndAssembly field if non-nil, zero value otherwise.

### GetIsFamilyAndAssemblyOk

`func (o *FieldInfo) GetIsFamilyAndAssemblyOk() (*bool, bool)`

GetIsFamilyAndAssemblyOk returns a tuple with the IsFamilyAndAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyAndAssembly

`func (o *FieldInfo) SetIsFamilyAndAssembly(v bool)`

SetIsFamilyAndAssembly sets IsFamilyAndAssembly field to given value.

### HasIsFamilyAndAssembly

`func (o *FieldInfo) HasIsFamilyAndAssembly() bool`

HasIsFamilyAndAssembly returns a boolean if a field has been set.

### GetIsFamilyOrAssembly

`func (o *FieldInfo) GetIsFamilyOrAssembly() bool`

GetIsFamilyOrAssembly returns the IsFamilyOrAssembly field if non-nil, zero value otherwise.

### GetIsFamilyOrAssemblyOk

`func (o *FieldInfo) GetIsFamilyOrAssemblyOk() (*bool, bool)`

GetIsFamilyOrAssemblyOk returns a tuple with the IsFamilyOrAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFamilyOrAssembly

`func (o *FieldInfo) SetIsFamilyOrAssembly(v bool)`

SetIsFamilyOrAssembly sets IsFamilyOrAssembly field to given value.

### HasIsFamilyOrAssembly

`func (o *FieldInfo) HasIsFamilyOrAssembly() bool`

HasIsFamilyOrAssembly returns a boolean if a field has been set.

### GetIsPrivate

`func (o *FieldInfo) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *FieldInfo) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *FieldInfo) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *FieldInfo) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIsPublic

`func (o *FieldInfo) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *FieldInfo) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *FieldInfo) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *FieldInfo) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetIsSecurityCritical

`func (o *FieldInfo) GetIsSecurityCritical() bool`

GetIsSecurityCritical returns the IsSecurityCritical field if non-nil, zero value otherwise.

### GetIsSecurityCriticalOk

`func (o *FieldInfo) GetIsSecurityCriticalOk() (*bool, bool)`

GetIsSecurityCriticalOk returns a tuple with the IsSecurityCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityCritical

`func (o *FieldInfo) SetIsSecurityCritical(v bool)`

SetIsSecurityCritical sets IsSecurityCritical field to given value.

### HasIsSecurityCritical

`func (o *FieldInfo) HasIsSecurityCritical() bool`

HasIsSecurityCritical returns a boolean if a field has been set.

### GetIsSecuritySafeCritical

`func (o *FieldInfo) GetIsSecuritySafeCritical() bool`

GetIsSecuritySafeCritical returns the IsSecuritySafeCritical field if non-nil, zero value otherwise.

### GetIsSecuritySafeCriticalOk

`func (o *FieldInfo) GetIsSecuritySafeCriticalOk() (*bool, bool)`

GetIsSecuritySafeCriticalOk returns a tuple with the IsSecuritySafeCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySafeCritical

`func (o *FieldInfo) SetIsSecuritySafeCritical(v bool)`

SetIsSecuritySafeCritical sets IsSecuritySafeCritical field to given value.

### HasIsSecuritySafeCritical

`func (o *FieldInfo) HasIsSecuritySafeCritical() bool`

HasIsSecuritySafeCritical returns a boolean if a field has been set.

### GetIsSecurityTransparent

`func (o *FieldInfo) GetIsSecurityTransparent() bool`

GetIsSecurityTransparent returns the IsSecurityTransparent field if non-nil, zero value otherwise.

### GetIsSecurityTransparentOk

`func (o *FieldInfo) GetIsSecurityTransparentOk() (*bool, bool)`

GetIsSecurityTransparentOk returns a tuple with the IsSecurityTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityTransparent

`func (o *FieldInfo) SetIsSecurityTransparent(v bool)`

SetIsSecurityTransparent sets IsSecurityTransparent field to given value.

### HasIsSecurityTransparent

`func (o *FieldInfo) HasIsSecurityTransparent() bool`

HasIsSecurityTransparent returns a boolean if a field has been set.

### GetFieldHandle

`func (o *FieldInfo) GetFieldHandle() RuntimeFieldHandle`

GetFieldHandle returns the FieldHandle field if non-nil, zero value otherwise.

### GetFieldHandleOk

`func (o *FieldInfo) GetFieldHandleOk() (*RuntimeFieldHandle, bool)`

GetFieldHandleOk returns a tuple with the FieldHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldHandle

`func (o *FieldInfo) SetFieldHandle(v RuntimeFieldHandle)`

SetFieldHandle sets FieldHandle field to given value.

### HasFieldHandle

`func (o *FieldInfo) HasFieldHandle() bool`

HasFieldHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


