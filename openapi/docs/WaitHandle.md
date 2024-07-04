# WaitHandle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **map[string]interface{}** |  | [optional] 
**SafeWaitHandle** | Pointer to [**SafeWaitHandle**](SafeWaitHandle.md) |  | [optional] 

## Methods

### NewWaitHandle

`func NewWaitHandle() *WaitHandle`

NewWaitHandle instantiates a new WaitHandle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitHandleWithDefaults

`func NewWaitHandleWithDefaults() *WaitHandle`

NewWaitHandleWithDefaults instantiates a new WaitHandle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *WaitHandle) GetHandle() map[string]interface{}`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *WaitHandle) GetHandleOk() (*map[string]interface{}, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *WaitHandle) SetHandle(v map[string]interface{})`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *WaitHandle) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetSafeWaitHandle

`func (o *WaitHandle) GetSafeWaitHandle() SafeWaitHandle`

GetSafeWaitHandle returns the SafeWaitHandle field if non-nil, zero value otherwise.

### GetSafeWaitHandleOk

`func (o *WaitHandle) GetSafeWaitHandleOk() (*SafeWaitHandle, bool)`

GetSafeWaitHandleOk returns a tuple with the SafeWaitHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeWaitHandle

`func (o *WaitHandle) SetSafeWaitHandle(v SafeWaitHandle)`

SetSafeWaitHandle sets SafeWaitHandle field to given value.

### HasSafeWaitHandle

`func (o *WaitHandle) HasSafeWaitHandle() bool`

HasSafeWaitHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


