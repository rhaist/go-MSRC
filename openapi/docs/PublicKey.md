# PublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedKeyValue** | Pointer to [**AsnEncodedData**](AsnEncodedData.md) |  | [optional] 
**EncodedParameters** | Pointer to [**AsnEncodedData**](AsnEncodedData.md) |  | [optional] 
**Key** | Pointer to [**AsymmetricAlgorithm**](AsymmetricAlgorithm.md) |  | [optional] 
**Oid** | Pointer to [**Oid**](Oid.md) |  | [optional] 

## Methods

### NewPublicKey

`func NewPublicKey() *PublicKey`

NewPublicKey instantiates a new PublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyWithDefaults

`func NewPublicKeyWithDefaults() *PublicKey`

NewPublicKeyWithDefaults instantiates a new PublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedKeyValue

`func (o *PublicKey) GetEncodedKeyValue() AsnEncodedData`

GetEncodedKeyValue returns the EncodedKeyValue field if non-nil, zero value otherwise.

### GetEncodedKeyValueOk

`func (o *PublicKey) GetEncodedKeyValueOk() (*AsnEncodedData, bool)`

GetEncodedKeyValueOk returns a tuple with the EncodedKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedKeyValue

`func (o *PublicKey) SetEncodedKeyValue(v AsnEncodedData)`

SetEncodedKeyValue sets EncodedKeyValue field to given value.

### HasEncodedKeyValue

`func (o *PublicKey) HasEncodedKeyValue() bool`

HasEncodedKeyValue returns a boolean if a field has been set.

### GetEncodedParameters

`func (o *PublicKey) GetEncodedParameters() AsnEncodedData`

GetEncodedParameters returns the EncodedParameters field if non-nil, zero value otherwise.

### GetEncodedParametersOk

`func (o *PublicKey) GetEncodedParametersOk() (*AsnEncodedData, bool)`

GetEncodedParametersOk returns a tuple with the EncodedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedParameters

`func (o *PublicKey) SetEncodedParameters(v AsnEncodedData)`

SetEncodedParameters sets EncodedParameters field to given value.

### HasEncodedParameters

`func (o *PublicKey) HasEncodedParameters() bool`

HasEncodedParameters returns a boolean if a field has been set.

### GetKey

`func (o *PublicKey) GetKey() AsymmetricAlgorithm`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublicKey) GetKeyOk() (*AsymmetricAlgorithm, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublicKey) SetKey(v AsymmetricAlgorithm)`

SetKey sets Key field to given value.

### HasKey

`func (o *PublicKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOid

`func (o *PublicKey) GetOid() Oid`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *PublicKey) GetOidOk() (*Oid, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *PublicKey) SetOid(v Oid)`

SetOid sets Oid field to given value.

### HasOid

`func (o *PublicKey) HasOid() bool`

HasOid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


