# TypeObjectKeyValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**Type**](Type.md) |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTypeObjectKeyValuePair

`func NewTypeObjectKeyValuePair() *TypeObjectKeyValuePair`

NewTypeObjectKeyValuePair instantiates a new TypeObjectKeyValuePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeObjectKeyValuePairWithDefaults

`func NewTypeObjectKeyValuePairWithDefaults() *TypeObjectKeyValuePair`

NewTypeObjectKeyValuePairWithDefaults instantiates a new TypeObjectKeyValuePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TypeObjectKeyValuePair) GetKey() Type`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TypeObjectKeyValuePair) GetKeyOk() (*Type, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TypeObjectKeyValuePair) SetKey(v Type)`

SetKey sets Key field to given value.

### HasKey

`func (o *TypeObjectKeyValuePair) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *TypeObjectKeyValuePair) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TypeObjectKeyValuePair) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TypeObjectKeyValuePair) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TypeObjectKeyValuePair) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TypeObjectKeyValuePair) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TypeObjectKeyValuePair) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


