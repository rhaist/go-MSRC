# SearchClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to [**SingleValueNode**](SingleValueNode.md) |  | [optional] 

## Methods

### NewSearchClause

`func NewSearchClause() *SearchClause`

NewSearchClause instantiates a new SearchClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchClauseWithDefaults

`func NewSearchClauseWithDefaults() *SearchClause`

NewSearchClauseWithDefaults instantiates a new SearchClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *SearchClause) GetExpression() SingleValueNode`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SearchClause) GetExpressionOk() (*SingleValueNode, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SearchClause) SetExpression(v SingleValueNode)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SearchClause) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


