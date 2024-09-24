# ByteReadOnlySpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int32** |  | [optional] [readonly] 
**IsEmpty** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewByteReadOnlySpan

`func NewByteReadOnlySpan() *ByteReadOnlySpan`

NewByteReadOnlySpan instantiates a new ByteReadOnlySpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByteReadOnlySpanWithDefaults

`func NewByteReadOnlySpanWithDefaults() *ByteReadOnlySpan`

NewByteReadOnlySpanWithDefaults instantiates a new ByteReadOnlySpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *ByteReadOnlySpan) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ByteReadOnlySpan) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ByteReadOnlySpan) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ByteReadOnlySpan) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetIsEmpty

`func (o *ByteReadOnlySpan) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *ByteReadOnlySpan) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *ByteReadOnlySpan) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.

### HasIsEmpty

`func (o *ByteReadOnlySpan) HasIsEmpty() bool`

HasIsEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


