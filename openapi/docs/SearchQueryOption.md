# SearchQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**ResultClrType** | Pointer to [**Type**](Type.md) |  | [optional] 
**SearchClause** | Pointer to [**SearchClause**](SearchClause.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSearchQueryOption

`func NewSearchQueryOption() *SearchQueryOption`

NewSearchQueryOption instantiates a new SearchQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryOptionWithDefaults

`func NewSearchQueryOptionWithDefaults() *SearchQueryOption`

NewSearchQueryOptionWithDefaults instantiates a new SearchQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SearchQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResultClrType

`func (o *SearchQueryOption) GetResultClrType() Type`

GetResultClrType returns the ResultClrType field if non-nil, zero value otherwise.

### GetResultClrTypeOk

`func (o *SearchQueryOption) GetResultClrTypeOk() (*Type, bool)`

GetResultClrTypeOk returns a tuple with the ResultClrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultClrType

`func (o *SearchQueryOption) SetResultClrType(v Type)`

SetResultClrType sets ResultClrType field to given value.

### HasResultClrType

`func (o *SearchQueryOption) HasResultClrType() bool`

HasResultClrType returns a boolean if a field has been set.

### GetSearchClause

`func (o *SearchQueryOption) GetSearchClause() SearchClause`

GetSearchClause returns the SearchClause field if non-nil, zero value otherwise.

### GetSearchClauseOk

`func (o *SearchQueryOption) GetSearchClauseOk() (*SearchClause, bool)`

GetSearchClauseOk returns a tuple with the SearchClause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchClause

`func (o *SearchQueryOption) SetSearchClause(v SearchClause)`

SetSearchClause sets SearchClause field to given value.

### HasSearchClause

`func (o *SearchQueryOption) HasSearchClause() bool`

HasSearchClause returns a boolean if a field has been set.

### GetRawValue

`func (o *SearchQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *SearchQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *SearchQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *SearchQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *SearchQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *SearchQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


