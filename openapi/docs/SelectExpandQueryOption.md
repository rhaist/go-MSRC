# SelectExpandQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawSelect** | Pointer to **NullableString** |  | [optional] [readonly] 
**RawExpand** | Pointer to **NullableString** |  | [optional] [readonly] 
**Compute** | Pointer to [**ComputeQueryOption**](ComputeQueryOption.md) |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**SelectExpandClause** | Pointer to [**SelectExpandClause**](SelectExpandClause.md) |  | [optional] 
**LevelsMaxLiteralExpansionDepth** | Pointer to **int32** |  | [optional] 

## Methods

### NewSelectExpandQueryOption

`func NewSelectExpandQueryOption() *SelectExpandQueryOption`

NewSelectExpandQueryOption instantiates a new SelectExpandQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectExpandQueryOptionWithDefaults

`func NewSelectExpandQueryOptionWithDefaults() *SelectExpandQueryOption`

NewSelectExpandQueryOptionWithDefaults instantiates a new SelectExpandQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SelectExpandQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SelectExpandQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SelectExpandQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SelectExpandQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawSelect

`func (o *SelectExpandQueryOption) GetRawSelect() string`

GetRawSelect returns the RawSelect field if non-nil, zero value otherwise.

### GetRawSelectOk

`func (o *SelectExpandQueryOption) GetRawSelectOk() (*string, bool)`

GetRawSelectOk returns a tuple with the RawSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSelect

`func (o *SelectExpandQueryOption) SetRawSelect(v string)`

SetRawSelect sets RawSelect field to given value.

### HasRawSelect

`func (o *SelectExpandQueryOption) HasRawSelect() bool`

HasRawSelect returns a boolean if a field has been set.

### SetRawSelectNil

`func (o *SelectExpandQueryOption) SetRawSelectNil(b bool)`

 SetRawSelectNil sets the value for RawSelect to be an explicit nil

### UnsetRawSelect
`func (o *SelectExpandQueryOption) UnsetRawSelect()`

UnsetRawSelect ensures that no value is present for RawSelect, not even an explicit nil
### GetRawExpand

`func (o *SelectExpandQueryOption) GetRawExpand() string`

GetRawExpand returns the RawExpand field if non-nil, zero value otherwise.

### GetRawExpandOk

`func (o *SelectExpandQueryOption) GetRawExpandOk() (*string, bool)`

GetRawExpandOk returns a tuple with the RawExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawExpand

`func (o *SelectExpandQueryOption) SetRawExpand(v string)`

SetRawExpand sets RawExpand field to given value.

### HasRawExpand

`func (o *SelectExpandQueryOption) HasRawExpand() bool`

HasRawExpand returns a boolean if a field has been set.

### SetRawExpandNil

`func (o *SelectExpandQueryOption) SetRawExpandNil(b bool)`

 SetRawExpandNil sets the value for RawExpand to be an explicit nil

### UnsetRawExpand
`func (o *SelectExpandQueryOption) UnsetRawExpand()`

UnsetRawExpand ensures that no value is present for RawExpand, not even an explicit nil
### GetCompute

`func (o *SelectExpandQueryOption) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *SelectExpandQueryOption) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *SelectExpandQueryOption) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *SelectExpandQueryOption) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetValidator

`func (o *SelectExpandQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *SelectExpandQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *SelectExpandQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *SelectExpandQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetSelectExpandClause

`func (o *SelectExpandQueryOption) GetSelectExpandClause() SelectExpandClause`

GetSelectExpandClause returns the SelectExpandClause field if non-nil, zero value otherwise.

### GetSelectExpandClauseOk

`func (o *SelectExpandQueryOption) GetSelectExpandClauseOk() (*SelectExpandClause, bool)`

GetSelectExpandClauseOk returns a tuple with the SelectExpandClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpandClause

`func (o *SelectExpandQueryOption) SetSelectExpandClause(v SelectExpandClause)`

SetSelectExpandClause sets SelectExpandClause field to given value.

### HasSelectExpandClause

`func (o *SelectExpandQueryOption) HasSelectExpandClause() bool`

HasSelectExpandClause returns a boolean if a field has been set.

### GetLevelsMaxLiteralExpansionDepth

`func (o *SelectExpandQueryOption) GetLevelsMaxLiteralExpansionDepth() int32`

GetLevelsMaxLiteralExpansionDepth returns the LevelsMaxLiteralExpansionDepth field if non-nil, zero value otherwise.

### GetLevelsMaxLiteralExpansionDepthOk

`func (o *SelectExpandQueryOption) GetLevelsMaxLiteralExpansionDepthOk() (*int32, bool)`

GetLevelsMaxLiteralExpansionDepthOk returns a tuple with the LevelsMaxLiteralExpansionDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelsMaxLiteralExpansionDepth

`func (o *SelectExpandQueryOption) SetLevelsMaxLiteralExpansionDepth(v int32)`

SetLevelsMaxLiteralExpansionDepth sets LevelsMaxLiteralExpansionDepth field to given value.

### HasLevelsMaxLiteralExpansionDepth

`func (o *SelectExpandQueryOption) HasLevelsMaxLiteralExpansionDepth() bool`

HasLevelsMaxLiteralExpansionDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


