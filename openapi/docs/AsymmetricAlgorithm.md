# AsymmetricAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySize** | Pointer to **int32** |  | [optional] 
**LegalKeySizes** | Pointer to [**[]KeySizes**](KeySizes.md) |  | [optional] [readonly] 
**SignatureAlgorithm** | Pointer to **NullableString** |  | [optional] [readonly] 
**KeyExchangeAlgorithm** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewAsymmetricAlgorithm

`func NewAsymmetricAlgorithm() *AsymmetricAlgorithm`

NewAsymmetricAlgorithm instantiates a new AsymmetricAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsymmetricAlgorithmWithDefaults

`func NewAsymmetricAlgorithmWithDefaults() *AsymmetricAlgorithm`

NewAsymmetricAlgorithmWithDefaults instantiates a new AsymmetricAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySize

`func (o *AsymmetricAlgorithm) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *AsymmetricAlgorithm) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *AsymmetricAlgorithm) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *AsymmetricAlgorithm) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetLegalKeySizes

`func (o *AsymmetricAlgorithm) GetLegalKeySizes() []KeySizes`

GetLegalKeySizes returns the LegalKeySizes field if non-nil, zero value otherwise.

### GetLegalKeySizesOk

`func (o *AsymmetricAlgorithm) GetLegalKeySizesOk() (*[]KeySizes, bool)`

GetLegalKeySizesOk returns a tuple with the LegalKeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalKeySizes

`func (o *AsymmetricAlgorithm) SetLegalKeySizes(v []KeySizes)`

SetLegalKeySizes sets LegalKeySizes field to given value.

### HasLegalKeySizes

`func (o *AsymmetricAlgorithm) HasLegalKeySizes() bool`

HasLegalKeySizes returns a boolean if a field has been set.

### SetLegalKeySizesNil

`func (o *AsymmetricAlgorithm) SetLegalKeySizesNil(b bool)`

 SetLegalKeySizesNil sets the value for LegalKeySizes to be an explicit nil

### UnsetLegalKeySizes
`func (o *AsymmetricAlgorithm) UnsetLegalKeySizes()`

UnsetLegalKeySizes ensures that no value is present for LegalKeySizes, not even an explicit nil
### GetSignatureAlgorithm

`func (o *AsymmetricAlgorithm) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *AsymmetricAlgorithm) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *AsymmetricAlgorithm) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *AsymmetricAlgorithm) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### SetSignatureAlgorithmNil

`func (o *AsymmetricAlgorithm) SetSignatureAlgorithmNil(b bool)`

 SetSignatureAlgorithmNil sets the value for SignatureAlgorithm to be an explicit nil

### UnsetSignatureAlgorithm
`func (o *AsymmetricAlgorithm) UnsetSignatureAlgorithm()`

UnsetSignatureAlgorithm ensures that no value is present for SignatureAlgorithm, not even an explicit nil
### GetKeyExchangeAlgorithm

`func (o *AsymmetricAlgorithm) GetKeyExchangeAlgorithm() string`

GetKeyExchangeAlgorithm returns the KeyExchangeAlgorithm field if non-nil, zero value otherwise.

### GetKeyExchangeAlgorithmOk

`func (o *AsymmetricAlgorithm) GetKeyExchangeAlgorithmOk() (*string, bool)`

GetKeyExchangeAlgorithmOk returns a tuple with the KeyExchangeAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchangeAlgorithm

`func (o *AsymmetricAlgorithm) SetKeyExchangeAlgorithm(v string)`

SetKeyExchangeAlgorithm sets KeyExchangeAlgorithm field to given value.

### HasKeyExchangeAlgorithm

`func (o *AsymmetricAlgorithm) HasKeyExchangeAlgorithm() bool`

HasKeyExchangeAlgorithm returns a boolean if a field has been set.

### SetKeyExchangeAlgorithmNil

`func (o *AsymmetricAlgorithm) SetKeyExchangeAlgorithmNil(b bool)`

 SetKeyExchangeAlgorithmNil sets the value for KeyExchangeAlgorithm to be an explicit nil

### UnsetKeyExchangeAlgorithm
`func (o *AsymmetricAlgorithm) UnsetKeyExchangeAlgorithm()`

UnsetKeyExchangeAlgorithm ensures that no value is present for KeyExchangeAlgorithm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


