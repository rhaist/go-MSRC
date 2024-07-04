# PipeWriter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanGetUnflushedBytes** | Pointer to **bool** |  | [optional] [readonly] 
**UnflushedBytes** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewPipeWriter

`func NewPipeWriter() *PipeWriter`

NewPipeWriter instantiates a new PipeWriter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipeWriterWithDefaults

`func NewPipeWriterWithDefaults() *PipeWriter`

NewPipeWriterWithDefaults instantiates a new PipeWriter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanGetUnflushedBytes

`func (o *PipeWriter) GetCanGetUnflushedBytes() bool`

GetCanGetUnflushedBytes returns the CanGetUnflushedBytes field if non-nil, zero value otherwise.

### GetCanGetUnflushedBytesOk

`func (o *PipeWriter) GetCanGetUnflushedBytesOk() (*bool, bool)`

GetCanGetUnflushedBytesOk returns a tuple with the CanGetUnflushedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGetUnflushedBytes

`func (o *PipeWriter) SetCanGetUnflushedBytes(v bool)`

SetCanGetUnflushedBytes sets CanGetUnflushedBytes field to given value.

### HasCanGetUnflushedBytes

`func (o *PipeWriter) HasCanGetUnflushedBytes() bool`

HasCanGetUnflushedBytes returns a boolean if a field has been set.

### GetUnflushedBytes

`func (o *PipeWriter) GetUnflushedBytes() int64`

GetUnflushedBytes returns the UnflushedBytes field if non-nil, zero value otherwise.

### GetUnflushedBytesOk

`func (o *PipeWriter) GetUnflushedBytesOk() (*int64, bool)`

GetUnflushedBytesOk returns a tuple with the UnflushedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnflushedBytes

`func (o *PipeWriter) SetUnflushedBytes(v int64)`

SetUnflushedBytes sets UnflushedBytes field to given value.

### HasUnflushedBytes

`func (o *PipeWriter) HasUnflushedBytes() bool`

HasUnflushedBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


