# PropertyInfo

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
**PropertyType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Attributes** | Pointer to [**PropertyAttributes**](PropertyAttributes.md) |  | [optional] 
**IsSpecialName** | Pointer to **bool** |  | [optional] [readonly] 
**CanRead** | Pointer to **bool** |  | [optional] [readonly] 
**CanWrite** | Pointer to **bool** |  | [optional] [readonly] 
**GetMethod** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 
**SetMethod** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 

## Methods

### NewPropertyInfo

`func NewPropertyInfo() *PropertyInfo`

NewPropertyInfo instantiates a new PropertyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyInfoWithDefaults

`func NewPropertyInfoWithDefaults() *PropertyInfo`

NewPropertyInfoWithDefaults instantiates a new PropertyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PropertyInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PropertyInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeclaringType

`func (o *PropertyInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *PropertyInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *PropertyInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *PropertyInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetReflectedType

`func (o *PropertyInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *PropertyInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *PropertyInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *PropertyInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetModule

`func (o *PropertyInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PropertyInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PropertyInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *PropertyInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *PropertyInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *PropertyInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *PropertyInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *PropertyInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *PropertyInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *PropertyInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *PropertyInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *PropertyInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *PropertyInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *PropertyInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *PropertyInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *PropertyInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *PropertyInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *PropertyInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetMemberType

`func (o *PropertyInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *PropertyInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *PropertyInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *PropertyInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetPropertyType

`func (o *PropertyInfo) GetPropertyType() Type`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PropertyInfo) GetPropertyTypeOk() (*Type, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PropertyInfo) SetPropertyType(v Type)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *PropertyInfo) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### GetAttributes

`func (o *PropertyInfo) GetAttributes() PropertyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PropertyInfo) GetAttributesOk() (*PropertyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PropertyInfo) SetAttributes(v PropertyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PropertyInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *PropertyInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *PropertyInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *PropertyInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *PropertyInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetCanRead

`func (o *PropertyInfo) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *PropertyInfo) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *PropertyInfo) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *PropertyInfo) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetCanWrite

`func (o *PropertyInfo) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *PropertyInfo) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *PropertyInfo) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *PropertyInfo) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.

### GetGetMethod

`func (o *PropertyInfo) GetGetMethod() MethodInfo`

GetGetMethod returns the GetMethod field if non-nil, zero value otherwise.

### GetGetMethodOk

`func (o *PropertyInfo) GetGetMethodOk() (*MethodInfo, bool)`

GetGetMethodOk returns a tuple with the GetMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetMethod

`func (o *PropertyInfo) SetGetMethod(v MethodInfo)`

SetGetMethod sets GetMethod field to given value.

### HasGetMethod

`func (o *PropertyInfo) HasGetMethod() bool`

HasGetMethod returns a boolean if a field has been set.

### GetSetMethod

`func (o *PropertyInfo) GetSetMethod() MethodInfo`

GetSetMethod returns the SetMethod field if non-nil, zero value otherwise.

### GetSetMethodOk

`func (o *PropertyInfo) GetSetMethodOk() (*MethodInfo, bool)`

GetSetMethodOk returns a tuple with the SetMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMethod

`func (o *PropertyInfo) SetSetMethod(v MethodInfo)`

SetSetMethod sets SetMethod field to given value.

### HasSetMethod

`func (o *PropertyInfo) HasSetMethod() bool`

HasSetMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


