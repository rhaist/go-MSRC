# OrderByClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThenBy** | Pointer to [**OrderByClause**](OrderByClause.md) |  | [optional] 
**Expression** | Pointer to [**SingleValueNode**](SingleValueNode.md) |  | [optional] 
**Direction** | Pointer to [**OrderByDirection**](OrderByDirection.md) |  | [optional] 
**RangeVariable** | Pointer to [**RangeVariable**](RangeVariable.md) |  | [optional] 
**ItemType** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 

## Methods

### NewOrderByClause

`func NewOrderByClause() *OrderByClause`

NewOrderByClause instantiates a new OrderByClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByClauseWithDefaults

`func NewOrderByClauseWithDefaults() *OrderByClause`

NewOrderByClauseWithDefaults instantiates a new OrderByClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThenBy

`func (o *OrderByClause) GetThenBy() OrderByClause`

GetThenBy returns the ThenBy field if non-nil, zero value otherwise.

### GetThenByOk

`func (o *OrderByClause) GetThenByOk() (*OrderByClause, bool)`

GetThenByOk returns a tuple with the ThenBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThenBy

`func (o *OrderByClause) SetThenBy(v OrderByClause)`

SetThenBy sets ThenBy field to given value.

### HasThenBy

`func (o *OrderByClause) HasThenBy() bool`

HasThenBy returns a boolean if a field has been set.

### GetExpression

`func (o *OrderByClause) GetExpression() SingleValueNode`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *OrderByClause) GetExpressionOk() (*SingleValueNode, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *OrderByClause) SetExpression(v SingleValueNode)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *OrderByClause) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetDirection

`func (o *OrderByClause) GetDirection() OrderByDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OrderByClause) GetDirectionOk() (*OrderByDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OrderByClause) SetDirection(v OrderByDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OrderByClause) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRangeVariable

`func (o *OrderByClause) GetRangeVariable() RangeVariable`

GetRangeVariable returns the RangeVariable field if non-nil, zero value otherwise.

### GetRangeVariableOk

`func (o *OrderByClause) GetRangeVariableOk() (*RangeVariable, bool)`

GetRangeVariableOk returns a tuple with the RangeVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeVariable

`func (o *OrderByClause) SetRangeVariable(v RangeVariable)`

SetRangeVariable sets RangeVariable field to given value.

### HasRangeVariable

`func (o *OrderByClause) HasRangeVariable() bool`

HasRangeVariable returns a boolean if a field has been set.

### GetItemType

`func (o *OrderByClause) GetItemType() IEdmTypeReference`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *OrderByClause) GetItemTypeOk() (*IEdmTypeReference, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *OrderByClause) SetItemType(v IEdmTypeReference)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *OrderByClause) HasItemType() bool`

HasItemType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


