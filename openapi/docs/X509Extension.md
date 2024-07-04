# X509Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | Pointer to [**Oid**](Oid.md) |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 
**Critical** | Pointer to **bool** |  | [optional] 

## Methods

### NewX509Extension

`func NewX509Extension() *X509Extension`

NewX509Extension instantiates a new X509Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509ExtensionWithDefaults

`func NewX509ExtensionWithDefaults() *X509Extension`

NewX509ExtensionWithDefaults instantiates a new X509Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *X509Extension) GetOid() Oid`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *X509Extension) GetOidOk() (*Oid, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *X509Extension) SetOid(v Oid)`

SetOid sets Oid field to given value.

### HasOid

`func (o *X509Extension) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetRawData

`func (o *X509Extension) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *X509Extension) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *X509Extension) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *X509Extension) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *X509Extension) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *X509Extension) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetCritical

`func (o *X509Extension) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *X509Extension) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *X509Extension) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *X509Extension) HasCritical() bool`

HasCritical returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


