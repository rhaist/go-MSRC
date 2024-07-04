# IEdmTypeReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNullable** | Pointer to **bool** |  | [optional] [readonly] 
**Definition** | Pointer to [**IEdmType**](IEdmType.md) |  | [optional] 

## Methods

### NewIEdmTypeReference

`func NewIEdmTypeReference() *IEdmTypeReference`

NewIEdmTypeReference instantiates a new IEdmTypeReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmTypeReferenceWithDefaults

`func NewIEdmTypeReferenceWithDefaults() *IEdmTypeReference`

NewIEdmTypeReferenceWithDefaults instantiates a new IEdmTypeReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNullable

`func (o *IEdmTypeReference) GetIsNullable() bool`

GetIsNullable returns the IsNullable field if non-nil, zero value otherwise.

### GetIsNullableOk

`func (o *IEdmTypeReference) GetIsNullableOk() (*bool, bool)`

GetIsNullableOk returns a tuple with the IsNullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNullable

`func (o *IEdmTypeReference) SetIsNullable(v bool)`

SetIsNullable sets IsNullable field to given value.

### HasIsNullable

`func (o *IEdmTypeReference) HasIsNullable() bool`

HasIsNullable returns a boolean if a field has been set.

### GetDefinition

`func (o *IEdmTypeReference) GetDefinition() IEdmType`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *IEdmTypeReference) GetDefinitionOk() (*IEdmType, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *IEdmTypeReference) SetDefinition(v IEdmType)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *IEdmTypeReference) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


