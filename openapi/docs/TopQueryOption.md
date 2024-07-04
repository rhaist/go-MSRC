# TopQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] [readonly] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTopQueryOption

`func NewTopQueryOption() *TopQueryOption`

NewTopQueryOption instantiates a new TopQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopQueryOptionWithDefaults

`func NewTopQueryOptionWithDefaults() *TopQueryOption`

NewTopQueryOptionWithDefaults instantiates a new TopQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TopQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TopQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TopQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TopQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValue

`func (o *TopQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *TopQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *TopQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *TopQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *TopQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *TopQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetValue

`func (o *TopQueryOption) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TopQueryOption) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TopQueryOption) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *TopQueryOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValidator

`func (o *TopQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *TopQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *TopQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *TopQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


