# PathString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**HasValue** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewPathString

`func NewPathString() *PathString`

NewPathString instantiates a new PathString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathStringWithDefaults

`func NewPathStringWithDefaults() *PathString`

NewPathStringWithDefaults instantiates a new PathString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PathString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PathString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PathString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PathString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PathString) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PathString) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetHasValue

`func (o *PathString) GetHasValue() bool`

GetHasValue returns the HasValue field if non-nil, zero value otherwise.

### GetHasValueOk

`func (o *PathString) GetHasValueOk() (*bool, bool)`

GetHasValueOk returns a tuple with the HasValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValue

`func (o *PathString) SetHasValue(v bool)`

SetHasValue sets HasValue field to given value.

### HasHasValue

`func (o *PathString) HasHasValue() bool`

HasHasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


