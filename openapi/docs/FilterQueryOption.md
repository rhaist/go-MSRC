# FilterQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**Compute** | Pointer to [**ComputeQueryOption**](ComputeQueryOption.md) |  | [optional] 
**FilterClause** | Pointer to [**FilterClause**](FilterClause.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFilterQueryOption

`func NewFilterQueryOption() *FilterQueryOption`

NewFilterQueryOption instantiates a new FilterQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterQueryOptionWithDefaults

`func NewFilterQueryOptionWithDefaults() *FilterQueryOption`

NewFilterQueryOptionWithDefaults instantiates a new FilterQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *FilterQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FilterQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FilterQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *FilterQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetValidator

`func (o *FilterQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *FilterQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *FilterQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *FilterQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetCompute

`func (o *FilterQueryOption) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *FilterQueryOption) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *FilterQueryOption) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *FilterQueryOption) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilterClause

`func (o *FilterQueryOption) GetFilterClause() FilterClause`

GetFilterClause returns the FilterClause field if non-nil, zero value otherwise.

### GetFilterClauseOk

`func (o *FilterQueryOption) GetFilterClauseOk() (*FilterClause, bool)`

GetFilterClauseOk returns a tuple with the FilterClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterClause

`func (o *FilterQueryOption) SetFilterClause(v FilterClause)`

SetFilterClause sets FilterClause field to given value.

### HasFilterClause

`func (o *FilterQueryOption) HasFilterClause() bool`

HasFilterClause returns a boolean if a field has been set.

### GetRawValue

`func (o *FilterQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *FilterQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *FilterQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *FilterQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *FilterQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *FilterQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


