# AsnEncodedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | Pointer to [**Oid**](Oid.md) |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAsnEncodedData

`func NewAsnEncodedData() *AsnEncodedData`

NewAsnEncodedData instantiates a new AsnEncodedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnEncodedDataWithDefaults

`func NewAsnEncodedDataWithDefaults() *AsnEncodedData`

NewAsnEncodedDataWithDefaults instantiates a new AsnEncodedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *AsnEncodedData) GetOid() Oid`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *AsnEncodedData) GetOidOk() (*Oid, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *AsnEncodedData) SetOid(v Oid)`

SetOid sets Oid field to given value.

### HasOid

`func (o *AsnEncodedData) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetRawData

`func (o *AsnEncodedData) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *AsnEncodedData) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *AsnEncodedData) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *AsnEncodedData) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *AsnEncodedData) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *AsnEncodedData) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


