# Stream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanRead** | Pointer to **bool** |  | [optional] [readonly] 
**CanWrite** | Pointer to **bool** |  | [optional] [readonly] 
**CanSeek** | Pointer to **bool** |  | [optional] [readonly] 
**CanTimeout** | Pointer to **bool** |  | [optional] [readonly] 
**Length** | Pointer to **int64** |  | [optional] [readonly] 
**Position** | Pointer to **int64** |  | [optional] 
**ReadTimeout** | Pointer to **int32** |  | [optional] 
**WriteTimeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewStream

`func NewStream() *Stream`

NewStream instantiates a new Stream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWithDefaults

`func NewStreamWithDefaults() *Stream`

NewStreamWithDefaults instantiates a new Stream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanRead

`func (o *Stream) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *Stream) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *Stream) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *Stream) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetCanWrite

`func (o *Stream) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *Stream) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *Stream) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *Stream) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.

### GetCanSeek

`func (o *Stream) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *Stream) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *Stream) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *Stream) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetCanTimeout

`func (o *Stream) GetCanTimeout() bool`

GetCanTimeout returns the CanTimeout field if non-nil, zero value otherwise.

### GetCanTimeoutOk

`func (o *Stream) GetCanTimeoutOk() (*bool, bool)`

GetCanTimeoutOk returns a tuple with the CanTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTimeout

`func (o *Stream) SetCanTimeout(v bool)`

SetCanTimeout sets CanTimeout field to given value.

### HasCanTimeout

`func (o *Stream) HasCanTimeout() bool`

HasCanTimeout returns a boolean if a field has been set.

### GetLength

`func (o *Stream) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Stream) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Stream) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *Stream) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPosition

`func (o *Stream) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Stream) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Stream) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Stream) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetReadTimeout

`func (o *Stream) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *Stream) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *Stream) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *Stream) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetWriteTimeout

`func (o *Stream) GetWriteTimeout() int32`

GetWriteTimeout returns the WriteTimeout field if non-nil, zero value otherwise.

### GetWriteTimeoutOk

`func (o *Stream) GetWriteTimeoutOk() (*int32, bool)`

GetWriteTimeoutOk returns a tuple with the WriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteTimeout

`func (o *Stream) SetWriteTimeout(v int32)`

SetWriteTimeout sets WriteTimeout field to given value.

### HasWriteTimeout

`func (o *Stream) HasWriteTimeout() bool`

HasWriteTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


