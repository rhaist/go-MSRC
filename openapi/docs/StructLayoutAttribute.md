# StructLayoutAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | Pointer to **interface{}** |  | [optional] [readonly] 
**Value** | Pointer to [**LayoutKind**](LayoutKind.md) |  | [optional] 

## Methods

### NewStructLayoutAttribute

`func NewStructLayoutAttribute() *StructLayoutAttribute`

NewStructLayoutAttribute instantiates a new StructLayoutAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructLayoutAttributeWithDefaults

`func NewStructLayoutAttributeWithDefaults() *StructLayoutAttribute`

NewStructLayoutAttributeWithDefaults instantiates a new StructLayoutAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *StructLayoutAttribute) GetTypeId() interface{}`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *StructLayoutAttribute) GetTypeIdOk() (*interface{}, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *StructLayoutAttribute) SetTypeId(v interface{})`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *StructLayoutAttribute) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *StructLayoutAttribute) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *StructLayoutAttribute) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetValue

`func (o *StructLayoutAttribute) GetValue() LayoutKind`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StructLayoutAttribute) GetValueOk() (*LayoutKind, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StructLayoutAttribute) SetValue(v LayoutKind)`

SetValue sets Value field to given value.

### HasValue

`func (o *StructLayoutAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


