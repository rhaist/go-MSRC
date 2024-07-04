# HostString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** |  | [optional] 
**HasValue** | Pointer to **bool** |  | [optional] [readonly] 
**Host** | Pointer to **NullableString** |  | [optional] [readonly] 
**Port** | Pointer to **NullableInt32** |  | [optional] [readonly] 

## Methods

### NewHostString

`func NewHostString() *HostString`

NewHostString instantiates a new HostString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostStringWithDefaults

`func NewHostStringWithDefaults() *HostString`

NewHostStringWithDefaults instantiates a new HostString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *HostString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HostString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HostString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HostString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HostString) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HostString) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetHasValue

`func (o *HostString) GetHasValue() bool`

GetHasValue returns the HasValue field if non-nil, zero value otherwise.

### GetHasValueOk

`func (o *HostString) GetHasValueOk() (*bool, bool)`

GetHasValueOk returns a tuple with the HasValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValue

`func (o *HostString) SetHasValue(v bool)`

SetHasValue sets HasValue field to given value.

### HasHasValue

`func (o *HostString) HasHasValue() bool`

HasHasValue returns a boolean if a field has been set.

### GetHost

`func (o *HostString) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostString) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostString) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostString) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HostString) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HostString) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HostString) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HostString) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HostString) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HostString) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HostString) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HostString) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


