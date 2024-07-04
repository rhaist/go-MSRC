# ModuleHandle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MdStreamVersion** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewModuleHandle

`func NewModuleHandle() *ModuleHandle`

NewModuleHandle instantiates a new ModuleHandle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleHandleWithDefaults

`func NewModuleHandleWithDefaults() *ModuleHandle`

NewModuleHandleWithDefaults instantiates a new ModuleHandle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdStreamVersion

`func (o *ModuleHandle) GetMdStreamVersion() int32`

GetMdStreamVersion returns the MdStreamVersion field if non-nil, zero value otherwise.

### GetMdStreamVersionOk

`func (o *ModuleHandle) GetMdStreamVersionOk() (*int32, bool)`

GetMdStreamVersionOk returns a tuple with the MdStreamVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdStreamVersion

`func (o *ModuleHandle) SetMdStreamVersion(v int32)`

SetMdStreamVersion sets MdStreamVersion field to given value.

### HasMdStreamVersion

`func (o *ModuleHandle) HasMdStreamVersion() bool`

HasMdStreamVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


