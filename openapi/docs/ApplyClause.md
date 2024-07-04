# ApplyClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transformations** | Pointer to [**[]TransformationNode**](TransformationNode.md) |  | [optional] 

## Methods

### NewApplyClause

`func NewApplyClause() *ApplyClause`

NewApplyClause instantiates a new ApplyClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyClauseWithDefaults

`func NewApplyClauseWithDefaults() *ApplyClause`

NewApplyClauseWithDefaults instantiates a new ApplyClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransformations

`func (o *ApplyClause) GetTransformations() []TransformationNode`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *ApplyClause) GetTransformationsOk() (*[]TransformationNode, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *ApplyClause) SetTransformations(v []TransformationNode)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *ApplyClause) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### SetTransformationsNil

`func (o *ApplyClause) SetTransformationsNil(b bool)`

 SetTransformationsNil sets the value for Transformations to be an explicit nil

### UnsetTransformations
`func (o *ApplyClause) UnsetTransformations()`

UnsetTransformations ensures that no value is present for Transformations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


