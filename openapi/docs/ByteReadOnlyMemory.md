# ByteReadOnlyMemory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int32** |  | [optional] [readonly] 
**IsEmpty** | Pointer to **bool** |  | [optional] [readonly] 
**Span** | Pointer to [**ByteReadOnlySpan**](ByteReadOnlySpan.md) |  | [optional] 

## Methods

### NewByteReadOnlyMemory

`func NewByteReadOnlyMemory() *ByteReadOnlyMemory`

NewByteReadOnlyMemory instantiates a new ByteReadOnlyMemory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByteReadOnlyMemoryWithDefaults

`func NewByteReadOnlyMemoryWithDefaults() *ByteReadOnlyMemory`

NewByteReadOnlyMemoryWithDefaults instantiates a new ByteReadOnlyMemory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *ByteReadOnlyMemory) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ByteReadOnlyMemory) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ByteReadOnlyMemory) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ByteReadOnlyMemory) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetIsEmpty

`func (o *ByteReadOnlyMemory) GetIsEmpty() bool`

GetIsEmpty returns the IsEmpty field if non-nil, zero value otherwise.

### GetIsEmptyOk

`func (o *ByteReadOnlyMemory) GetIsEmptyOk() (*bool, bool)`

GetIsEmptyOk returns a tuple with the IsEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmpty

`func (o *ByteReadOnlyMemory) SetIsEmpty(v bool)`

SetIsEmpty sets IsEmpty field to given value.

### HasIsEmpty

`func (o *ByteReadOnlyMemory) HasIsEmpty() bool`

HasIsEmpty returns a boolean if a field has been set.

### GetSpan

`func (o *ByteReadOnlyMemory) GetSpan() ByteReadOnlySpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ByteReadOnlyMemory) GetSpanOk() (*ByteReadOnlySpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ByteReadOnlyMemory) SetSpan(v ByteReadOnlySpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ByteReadOnlyMemory) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


