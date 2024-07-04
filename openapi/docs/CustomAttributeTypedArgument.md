# CustomAttributeTypedArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgumentType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCustomAttributeTypedArgument

`func NewCustomAttributeTypedArgument() *CustomAttributeTypedArgument`

NewCustomAttributeTypedArgument instantiates a new CustomAttributeTypedArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeTypedArgumentWithDefaults

`func NewCustomAttributeTypedArgumentWithDefaults() *CustomAttributeTypedArgument`

NewCustomAttributeTypedArgumentWithDefaults instantiates a new CustomAttributeTypedArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgumentType

`func (o *CustomAttributeTypedArgument) GetArgumentType() Type`

GetArgumentType returns the ArgumentType field if non-nil, zero value otherwise.

### GetArgumentTypeOk

`func (o *CustomAttributeTypedArgument) GetArgumentTypeOk() (*Type, bool)`

GetArgumentTypeOk returns a tuple with the ArgumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgumentType

`func (o *CustomAttributeTypedArgument) SetArgumentType(v Type)`

SetArgumentType sets ArgumentType field to given value.

### HasArgumentType

`func (o *CustomAttributeTypedArgument) HasArgumentType() bool`

HasArgumentType returns a boolean if a field has been set.

### GetValue

`func (o *CustomAttributeTypedArgument) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomAttributeTypedArgument) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomAttributeTypedArgument) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomAttributeTypedArgument) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomAttributeTypedArgument) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomAttributeTypedArgument) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


