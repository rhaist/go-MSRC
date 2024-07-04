# SafeWaitHandle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsClosed** | Pointer to **bool** |  | [optional] [readonly] 
**IsInvalid** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewSafeWaitHandle

`func NewSafeWaitHandle() *SafeWaitHandle`

NewSafeWaitHandle instantiates a new SafeWaitHandle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeWaitHandleWithDefaults

`func NewSafeWaitHandleWithDefaults() *SafeWaitHandle`

NewSafeWaitHandleWithDefaults instantiates a new SafeWaitHandle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsClosed

`func (o *SafeWaitHandle) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *SafeWaitHandle) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *SafeWaitHandle) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *SafeWaitHandle) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsInvalid

`func (o *SafeWaitHandle) GetIsInvalid() bool`

GetIsInvalid returns the IsInvalid field if non-nil, zero value otherwise.

### GetIsInvalidOk

`func (o *SafeWaitHandle) GetIsInvalidOk() (*bool, bool)`

GetIsInvalidOk returns a tuple with the IsInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalid

`func (o *SafeWaitHandle) SetIsInvalid(v bool)`

SetIsInvalid sets IsInvalid field to given value.

### HasIsInvalid

`func (o *SafeWaitHandle) HasIsInvalid() bool`

HasIsInvalid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


