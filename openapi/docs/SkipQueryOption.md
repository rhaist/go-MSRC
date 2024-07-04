# SkipQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] [readonly] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSkipQueryOption

`func NewSkipQueryOption() *SkipQueryOption`

NewSkipQueryOption instantiates a new SkipQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkipQueryOptionWithDefaults

`func NewSkipQueryOptionWithDefaults() *SkipQueryOption`

NewSkipQueryOptionWithDefaults instantiates a new SkipQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SkipQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SkipQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SkipQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SkipQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValue

`func (o *SkipQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *SkipQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *SkipQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *SkipQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *SkipQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *SkipQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetValue

`func (o *SkipQueryOption) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SkipQueryOption) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SkipQueryOption) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SkipQueryOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValidator

`func (o *SkipQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *SkipQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *SkipQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *SkipQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


