# SkipTokenQueryOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawValue** | Pointer to **NullableString** |  | [optional] 
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**Handler** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSkipTokenQueryOption

`func NewSkipTokenQueryOption() *SkipTokenQueryOption`

NewSkipTokenQueryOption instantiates a new SkipTokenQueryOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkipTokenQueryOptionWithDefaults

`func NewSkipTokenQueryOptionWithDefaults() *SkipTokenQueryOption`

NewSkipTokenQueryOptionWithDefaults instantiates a new SkipTokenQueryOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawValue

`func (o *SkipTokenQueryOption) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *SkipTokenQueryOption) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *SkipTokenQueryOption) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *SkipTokenQueryOption) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *SkipTokenQueryOption) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *SkipTokenQueryOption) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
### GetContext

`func (o *SkipTokenQueryOption) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SkipTokenQueryOption) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SkipTokenQueryOption) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SkipTokenQueryOption) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetValidator

`func (o *SkipTokenQueryOption) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *SkipTokenQueryOption) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *SkipTokenQueryOption) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *SkipTokenQueryOption) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetHandler

`func (o *SkipTokenQueryOption) GetHandler() map[string]interface{}`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *SkipTokenQueryOption) GetHandlerOk() (*map[string]interface{}, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *SkipTokenQueryOption) SetHandler(v map[string]interface{})`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *SkipTokenQueryOption) HasHandler() bool`

HasHandler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


