# ApplyQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**ResultClrType** | Pointer to [**Type**](Type.md) |  | [optional] 
**ApplyClause** | Pointer to [**ApplyClause**](ApplyClause.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApplyQueryOption

`func NewApplyQueryOption() *ApplyQueryOption`

NewApplyQueryOption instantiates a new ApplyQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyQueryOptionWithDefaults

`func NewApplyQueryOptionWithDefaults() *ApplyQueryOption`

NewApplyQueryOptionWithDefaults instantiates a new ApplyQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ApplyQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ApplyQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ApplyQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ApplyQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResultClrType

`func (o *ApplyQueryOption) GetResultClrType() Type`

GetResultClrType returns the ResultClrType field if non-nil, zero value otherwise.

### GetResultClrTypeOk

`func (o *ApplyQueryOption) GetResultClrTypeOk() (*Type, bool)`

GetResultClrTypeOk returns a tuple with the ResultClrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultClrType

`func (o *ApplyQueryOption) SetResultClrType(v Type)`

SetResultClrType sets ResultClrType field to given value.

### HasResultClrType

`func (o *ApplyQueryOption) HasResultClrType() bool`

HasResultClrType returns a boolean if a field has been set.

### GetApplyClause

`func (o *ApplyQueryOption) GetApplyClause() ApplyClause`

GetApplyClause returns the ApplyClause field if non-nil, zero value otherwise.

### GetApplyClauseOk

`func (o *ApplyQueryOption) GetApplyClauseOk() (*ApplyClause, bool)`

GetApplyClauseOk returns a tuple with the ApplyClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyClause

`func (o *ApplyQueryOption) SetApplyClause(v ApplyClause)`

SetApplyClause sets ApplyClause field to given value.

### HasApplyClause

`func (o *ApplyQueryOption) HasApplyClause() bool`

HasApplyClause returns a boolean if a field has been set.

### GetRawValue

`func (o *ApplyQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *ApplyQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *ApplyQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *ApplyQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *ApplyQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *ApplyQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


