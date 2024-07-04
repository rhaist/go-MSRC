# CancellationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCancellationRequested** | Pointer to **bool** |  | [optional] [readonly] 
**CanBeCanceled** | Pointer to **bool** |  | [optional] [readonly] 
**WaitHandle** | Pointer to [**WaitHandle**](WaitHandle.md) |  | [optional] 

## Methods

### NewCancellationToken

`func NewCancellationToken() *CancellationToken`

NewCancellationToken instantiates a new CancellationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationTokenWithDefaults

`func NewCancellationTokenWithDefaults() *CancellationToken`

NewCancellationTokenWithDefaults instantiates a new CancellationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCancellationRequested

`func (o *CancellationToken) GetIsCancellationRequested() bool`

GetIsCancellationRequested returns the IsCancellationRequested field if non-nil, zero value otherwise.

### GetIsCancellationRequestedOk

`func (o *CancellationToken) GetIsCancellationRequestedOk() (*bool, bool)`

GetIsCancellationRequestedOk returns a tuple with the IsCancellationRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancellationRequested

`func (o *CancellationToken) SetIsCancellationRequested(v bool)`

SetIsCancellationRequested sets IsCancellationRequested field to given value.

### HasIsCancellationRequested

`func (o *CancellationToken) HasIsCancellationRequested() bool`

HasIsCancellationRequested returns a boolean if a field has been set.

### GetCanBeCanceled

`func (o *CancellationToken) GetCanBeCanceled() bool`

GetCanBeCanceled returns the CanBeCanceled field if non-nil, zero value otherwise.

### GetCanBeCanceledOk

`func (o *CancellationToken) GetCanBeCanceledOk() (*bool, bool)`

GetCanBeCanceledOk returns a tuple with the CanBeCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeCanceled

`func (o *CancellationToken) SetCanBeCanceled(v bool)`

SetCanBeCanceled sets CanBeCanceled field to given value.

### HasCanBeCanceled

`func (o *CancellationToken) HasCanBeCanceled() bool`

HasCanBeCanceled returns a boolean if a field has been set.

### GetWaitHandle

`func (o *CancellationToken) GetWaitHandle() WaitHandle`

GetWaitHandle returns the WaitHandle field if non-nil, zero value otherwise.

### GetWaitHandleOk

`func (o *CancellationToken) GetWaitHandleOk() (*WaitHandle, bool)`

GetWaitHandleOk returns a tuple with the WaitHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitHandle

`func (o *CancellationToken) SetWaitHandle(v WaitHandle)`

SetWaitHandle sets WaitHandle field to given value.

### HasWaitHandle

`func (o *CancellationToken) HasWaitHandle() bool`

HasWaitHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


