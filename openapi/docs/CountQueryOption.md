# CountQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] [readonly] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCountQueryOption

`func NewCountQueryOption() *CountQueryOption`

NewCountQueryOption instantiates a new CountQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountQueryOptionWithDefaults

`func NewCountQueryOptionWithDefaults() *CountQueryOption`

NewCountQueryOptionWithDefaults instantiates a new CountQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CountQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CountQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CountQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *CountQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValue

`func (o *CountQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *CountQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *CountQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *CountQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *CountQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *CountQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetValue

`func (o *CountQueryOption) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CountQueryOption) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CountQueryOption) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *CountQueryOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValidator

`func (o *CountQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *CountQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *CountQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *CountQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


