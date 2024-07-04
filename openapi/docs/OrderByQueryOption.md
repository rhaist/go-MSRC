# OrderByQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**OrderByNodes** | Pointer to [**[]OrderByNode**](OrderByNode.md) |  | [optional] [readonly] 
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**Compute** | Pointer to [**ComputeQueryOption**](ComputeQueryOption.md) |  | [optional] 
**OrderByClause** | Pointer to [**OrderByClause**](OrderByClause.md) |  | [optional] 

## Methods

### NewOrderByQueryOption

`func NewOrderByQueryOption() *OrderByQueryOption`

NewOrderByQueryOption instantiates a new OrderByQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByQueryOptionWithDefaults

`func NewOrderByQueryOptionWithDefaults() *OrderByQueryOption`

NewOrderByQueryOptionWithDefaults instantiates a new OrderByQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OrderByQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OrderByQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OrderByQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *OrderByQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetOrderByNodes

`func (o *OrderByQueryOption) GetOrderByNodes() []OrderByNode`

GetOrderByNodes returns the OrderByNodes field if non-nil, zero value otherwise.

### GetOrderByNodesOk

`func (o *OrderByQueryOption) GetOrderByNodesOk() (*[]OrderByNode, bool)`

GetOrderByNodesOk returns a tuple with the OrderByNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByNodes

`func (o *OrderByQueryOption) SetOrderByNodes(v []OrderByNode)`

SetOrderByNodes sets OrderByNodes field to given value.

### HasOrderByNodes

`func (o *OrderByQueryOption) HasOrderByNodes() bool`

HasOrderByNodes returns a boolean if a field has been set.

### SetOrderByNodesNil

`func (o *OrderByQueryOption) SetOrderByNodesNil(b bool)`

 SetOrderByNodesNil sets the value for OrderByNodes to be an explicit nil

### UnsetOrderByNodes
`func (o *OrderByQueryOption) UnsetOrderByNodes()`

UnsetOrderByNodes ensures that no value is present for OrderByNodes, not even an explicit nil
### GetRawValue

`func (o *OrderByQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *OrderByQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *OrderByQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *OrderByQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *OrderByQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *OrderByQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetValidator

`func (o *OrderByQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *OrderByQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *OrderByQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *OrderByQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetCompute

`func (o *OrderByQueryOption) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *OrderByQueryOption) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *OrderByQueryOption) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *OrderByQueryOption) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetOrderByClause

`func (o *OrderByQueryOption) GetOrderByClause() OrderByClause`

GetOrderByClause returns the OrderByClause field if non-nil, zero value otherwise.

### GetOrderByClauseOk

`func (o *OrderByQueryOption) GetOrderByClauseOk() (*OrderByClause, bool)`

GetOrderByClauseOk returns a tuple with the OrderByClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByClause

`func (o *OrderByQueryOption) SetOrderByClause(v OrderByClause)`

SetOrderByClause sets OrderByClause field to given value.

### HasOrderByClause

`func (o *OrderByQueryOption) HasOrderByClause() bool`

HasOrderByClause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


