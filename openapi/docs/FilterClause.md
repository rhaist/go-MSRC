# FilterClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to [**SingleValueNode**](SingleValueNode.md) |  | [optional] 
**RangeVariable** | Pointer to [**RangeVariable**](RangeVariable.md) |  | [optional] 
**ItemType** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 

## Methods

### NewFilterClause

`func NewFilterClause() *FilterClause`

NewFilterClause instantiates a new FilterClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterClauseWithDefaults

`func NewFilterClauseWithDefaults() *FilterClause`

NewFilterClauseWithDefaults instantiates a new FilterClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *FilterClause) GetExpression() SingleValueNode`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *FilterClause) GetExpressionOk() (*SingleValueNode, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *FilterClause) SetExpression(v SingleValueNode)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *FilterClause) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetRangeVariable

`func (o *FilterClause) GetRangeVariable() RangeVariable`

GetRangeVariable returns the RangeVariable field if non-nil, zero value otherwise.

### GetRangeVariableOk

`func (o *FilterClause) GetRangeVariableOk() (*RangeVariable, bool)`

GetRangeVariableOk returns a tuple with the RangeVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeVariable

`func (o *FilterClause) SetRangeVariable(v RangeVariable)`

SetRangeVariable sets RangeVariable field to given value.

### HasRangeVariable

`func (o *FilterClause) HasRangeVariable() bool`

HasRangeVariable returns a boolean if a field has been set.

### GetItemType

`func (o *FilterClause) GetItemType() IEdmTypeReference`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *FilterClause) GetItemTypeOk() (*IEdmTypeReference, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *FilterClause) SetItemType(v IEdmTypeReference)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *FilterClause) HasItemType() bool`

HasItemType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


