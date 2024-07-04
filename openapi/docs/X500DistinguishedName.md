# X500DistinguishedName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | Pointer to [**Oid**](Oid.md) |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewX500DistinguishedName

`func NewX500DistinguishedName() *X500DistinguishedName`

NewX500DistinguishedName instantiates a new X500DistinguishedName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX500DistinguishedNameWithDefaults

`func NewX500DistinguishedNameWithDefaults() *X500DistinguishedName`

NewX500DistinguishedNameWithDefaults instantiates a new X500DistinguishedName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *X500DistinguishedName) GetOid() Oid`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *X500DistinguishedName) GetOidOk() (*Oid, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *X500DistinguishedName) SetOid(v Oid)`

SetOid sets Oid field to given value.

### HasOid

`func (o *X500DistinguishedName) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetRawData

`func (o *X500DistinguishedName) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *X500DistinguishedName) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *X500DistinguishedName) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *X500DistinguishedName) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *X500DistinguishedName) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *X500DistinguishedName) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetName

`func (o *X500DistinguishedName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *X500DistinguishedName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *X500DistinguishedName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *X500DistinguishedName) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *X500DistinguishedName) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *X500DistinguishedName) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


