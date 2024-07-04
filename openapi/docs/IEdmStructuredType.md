# IEdmStructuredType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAbstract** | Pointer to **bool** |  | [optional] [readonly] 
**IsOpen** | Pointer to **bool** |  | [optional] [readonly] 
**BaseType** | Pointer to [**IEdmStructuredType**](IEdmStructuredType.md) |  | [optional] 
**DeclaredProperties** | Pointer to [**[]IEdmProperty**](IEdmProperty.md) |  | [optional] [readonly] 
**TypeKind** | Pointer to [**EdmTypeKind**](EdmTypeKind.md) |  | [optional] 

## Methods

### NewIEdmStructuredType

`func NewIEdmStructuredType() *IEdmStructuredType`

NewIEdmStructuredType instantiates a new IEdmStructuredType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmStructuredTypeWithDefaults

`func NewIEdmStructuredTypeWithDefaults() *IEdmStructuredType`

NewIEdmStructuredTypeWithDefaults instantiates a new IEdmStructuredType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAbstract

`func (o *IEdmStructuredType) GetIsAbstract() bool`

GetIsAbstract returns the IsAbstract field if non-nil, zero value otherwise.

### GetIsAbstractOk

`func (o *IEdmStructuredType) GetIsAbstractOk() (*bool, bool)`

GetIsAbstractOk returns a tuple with the IsAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbstract

`func (o *IEdmStructuredType) SetIsAbstract(v bool)`

SetIsAbstract sets IsAbstract field to given value.

### HasIsAbstract

`func (o *IEdmStructuredType) HasIsAbstract() bool`

HasIsAbstract returns a boolean if a field has been set.

### GetIsOpen

`func (o *IEdmStructuredType) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *IEdmStructuredType) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *IEdmStructuredType) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *IEdmStructuredType) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetBaseType

`func (o *IEdmStructuredType) GetBaseType() IEdmStructuredType`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *IEdmStructuredType) GetBaseTypeOk() (*IEdmStructuredType, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *IEdmStructuredType) SetBaseType(v IEdmStructuredType)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *IEdmStructuredType) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetDeclaredProperties

`func (o *IEdmStructuredType) GetDeclaredProperties() []IEdmProperty`

GetDeclaredProperties returns the DeclaredProperties field if non-nil, zero value otherwise.

### GetDeclaredPropertiesOk

`func (o *IEdmStructuredType) GetDeclaredPropertiesOk() (*[]IEdmProperty, bool)`

GetDeclaredPropertiesOk returns a tuple with the DeclaredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredProperties

`func (o *IEdmStructuredType) SetDeclaredProperties(v []IEdmProperty)`

SetDeclaredProperties sets DeclaredProperties field to given value.

### HasDeclaredProperties

`func (o *IEdmStructuredType) HasDeclaredProperties() bool`

HasDeclaredProperties returns a boolean if a field has been set.

### SetDeclaredPropertiesNil

`func (o *IEdmStructuredType) SetDeclaredPropertiesNil(b bool)`

 SetDeclaredPropertiesNil sets the value for DeclaredProperties to be an explicit nil

### UnsetDeclaredProperties
`func (o *IEdmStructuredType) UnsetDeclaredProperties()`

UnsetDeclaredProperties ensures that no value is present for DeclaredProperties, not even an explicit nil
### GetTypeKind

`func (o *IEdmStructuredType) GetTypeKind() EdmTypeKind`

GetTypeKind returns the TypeKind field if non-nil, zero value otherwise.

### GetTypeKindOk

`func (o *IEdmStructuredType) GetTypeKindOk() (*EdmTypeKind, bool)`

GetTypeKindOk returns a tuple with the TypeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeKind

`func (o *IEdmStructuredType) SetTypeKind(v EdmTypeKind)`

SetTypeKind sets TypeKind field to given value.

### HasTypeKind

`func (o *IEdmStructuredType) HasTypeKind() bool`

HasTypeKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


