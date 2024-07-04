# ComputeQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**ResultClrType** | Pointer to [**Type**](Type.md) |  | [optional] 
**ComputeClause** | Pointer to [**ComputeClause**](ComputeClause.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewComputeQueryOption

`func NewComputeQueryOption() *ComputeQueryOption`

NewComputeQueryOption instantiates a new ComputeQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeQueryOptionWithDefaults

`func NewComputeQueryOptionWithDefaults() *ComputeQueryOption`

NewComputeQueryOptionWithDefaults instantiates a new ComputeQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ComputeQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ComputeQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ComputeQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ComputeQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResultClrType

`func (o *ComputeQueryOption) GetResultClrType() Type`

GetResultClrType returns the ResultClrType field if non-nil, zero value otherwise.

### GetResultClrTypeOk

`func (o *ComputeQueryOption) GetResultClrTypeOk() (*Type, bool)`

GetResultClrTypeOk returns a tuple with the ResultClrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultClrType

`func (o *ComputeQueryOption) SetResultClrType(v Type)`

SetResultClrType sets ResultClrType field to given value.

### HasResultClrType

`func (o *ComputeQueryOption) HasResultClrType() bool`

HasResultClrType returns a boolean if a field has been set.

### GetComputeClause

`func (o *ComputeQueryOption) GetComputeClause() ComputeClause`

GetComputeClause returns the ComputeClause field if non-nil, zero value otherwise.

### GetComputeClauseOk

`func (o *ComputeQueryOption) GetComputeClauseOk() (*ComputeClause, bool)`

GetComputeClauseOk returns a tuple with the ComputeClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeClause

`func (o *ComputeQueryOption) SetComputeClause(v ComputeClause)`

SetComputeClause sets ComputeClause field to given value.

### HasComputeClause

`func (o *ComputeQueryOption) HasComputeClause() bool`

HasComputeClause returns a boolean if a field has been set.

### GetRawValue

`func (o *ComputeQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *ComputeQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *ComputeQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *ComputeQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *ComputeQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *ComputeQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetValidator

`func (o *ComputeQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *ComputeQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *ComputeQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *ComputeQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


