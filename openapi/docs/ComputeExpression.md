# ComputeExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to [**SingleValueNode**](SingleValueNode.md) |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**TypeReference** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 

## Methods

### NewComputeExpression

`func NewComputeExpression() *ComputeExpression`

NewComputeExpression instantiates a new ComputeExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeExpressionWithDefaults

`func NewComputeExpressionWithDefaults() *ComputeExpression`

NewComputeExpressionWithDefaults instantiates a new ComputeExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ComputeExpression) GetExpression() SingleValueNode`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ComputeExpression) GetExpressionOk() (*SingleValueNode, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ComputeExpression) SetExpression(v SingleValueNode)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ComputeExpression) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetAlias

`func (o *ComputeExpression) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ComputeExpression) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ComputeExpression) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ComputeExpression) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ComputeExpression) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ComputeExpression) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetTypeReference

`func (o *ComputeExpression) GetTypeReference() IEdmTypeReference`

GetTypeReference returns the TypeReference field if non-nil, zero value otherwise.

### GetTypeReferenceOk

`func (o *ComputeExpression) GetTypeReferenceOk() (*IEdmTypeReference, bool)`

GetTypeReferenceOk returns a tuple with the TypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeReference

`func (o *ComputeExpression) SetTypeReference(v IEdmTypeReference)`

SetTypeReference sets TypeReference field to given value.

### HasTypeReference

`func (o *ComputeExpression) HasTypeReference() bool`

HasTypeReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


