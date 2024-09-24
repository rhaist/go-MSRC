# X509Certificate2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **map[string]interface{}** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] [readonly] 
**Subject** | Pointer to **NullableString** |  | [optional] [readonly] 
**SerialNumberBytes** | Pointer to [**ByteReadOnlyMemory**](ByteReadOnlyMemory.md) |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Extensions** | Pointer to [**[]X509Extension**](X509Extension.md) |  | [optional] [readonly] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] [readonly] 
**PrivateKey** | Pointer to [**AsymmetricAlgorithm**](AsymmetricAlgorithm.md) |  | [optional] 
**IssuerName** | Pointer to [**X500DistinguishedName**](X500DistinguishedName.md) |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] [readonly] 
**NotBefore** | Pointer to **time.Time** |  | [optional] [readonly] 
**PublicKey** | Pointer to [**PublicKey**](PublicKey.md) |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] [readonly] 
**RawDataMemory** | Pointer to [**ByteReadOnlyMemory**](ByteReadOnlyMemory.md) |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] [readonly] 
**SignatureAlgorithm** | Pointer to [**Oid**](Oid.md) |  | [optional] 
**SubjectName** | Pointer to [**X500DistinguishedName**](X500DistinguishedName.md) |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewX509Certificate2

`func NewX509Certificate2() *X509Certificate2`

NewX509Certificate2 instantiates a new X509Certificate2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509Certificate2WithDefaults

`func NewX509Certificate2WithDefaults() *X509Certificate2`

NewX509Certificate2WithDefaults instantiates a new X509Certificate2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *X509Certificate2) GetHandle() map[string]interface{}`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *X509Certificate2) GetHandleOk() (*map[string]interface{}, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *X509Certificate2) SetHandle(v map[string]interface{})`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *X509Certificate2) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIssuer

`func (o *X509Certificate2) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *X509Certificate2) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *X509Certificate2) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *X509Certificate2) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *X509Certificate2) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *X509Certificate2) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetSubject

`func (o *X509Certificate2) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *X509Certificate2) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *X509Certificate2) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *X509Certificate2) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *X509Certificate2) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *X509Certificate2) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSerialNumberBytes

`func (o *X509Certificate2) GetSerialNumberBytes() ByteReadOnlyMemory`

GetSerialNumberBytes returns the SerialNumberBytes field if non-nil, zero value otherwise.

### GetSerialNumberBytesOk

`func (o *X509Certificate2) GetSerialNumberBytesOk() (*ByteReadOnlyMemory, bool)`

GetSerialNumberBytesOk returns a tuple with the SerialNumberBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberBytes

`func (o *X509Certificate2) SetSerialNumberBytes(v ByteReadOnlyMemory)`

SetSerialNumberBytes sets SerialNumberBytes field to given value.

### HasSerialNumberBytes

`func (o *X509Certificate2) HasSerialNumberBytes() bool`

HasSerialNumberBytes returns a boolean if a field has been set.

### GetArchived

`func (o *X509Certificate2) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *X509Certificate2) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *X509Certificate2) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *X509Certificate2) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetExtensions

`func (o *X509Certificate2) GetExtensions() []X509Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *X509Certificate2) GetExtensionsOk() (*[]X509Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *X509Certificate2) SetExtensions(v []X509Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *X509Certificate2) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *X509Certificate2) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *X509Certificate2) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetFriendlyName

`func (o *X509Certificate2) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *X509Certificate2) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *X509Certificate2) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *X509Certificate2) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *X509Certificate2) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *X509Certificate2) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetHasPrivateKey

`func (o *X509Certificate2) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *X509Certificate2) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *X509Certificate2) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *X509Certificate2) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *X509Certificate2) GetPrivateKey() AsymmetricAlgorithm`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *X509Certificate2) GetPrivateKeyOk() (*AsymmetricAlgorithm, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *X509Certificate2) SetPrivateKey(v AsymmetricAlgorithm)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *X509Certificate2) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetIssuerName

`func (o *X509Certificate2) GetIssuerName() X500DistinguishedName`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *X509Certificate2) GetIssuerNameOk() (*X500DistinguishedName, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *X509Certificate2) SetIssuerName(v X500DistinguishedName)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *X509Certificate2) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### GetNotAfter

`func (o *X509Certificate2) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *X509Certificate2) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *X509Certificate2) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *X509Certificate2) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *X509Certificate2) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *X509Certificate2) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *X509Certificate2) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *X509Certificate2) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetPublicKey

`func (o *X509Certificate2) GetPublicKey() PublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *X509Certificate2) GetPublicKeyOk() (*PublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *X509Certificate2) SetPublicKey(v PublicKey)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *X509Certificate2) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRawData

`func (o *X509Certificate2) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *X509Certificate2) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *X509Certificate2) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *X509Certificate2) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *X509Certificate2) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *X509Certificate2) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil
### GetRawDataMemory

`func (o *X509Certificate2) GetRawDataMemory() ByteReadOnlyMemory`

GetRawDataMemory returns the RawDataMemory field if non-nil, zero value otherwise.

### GetRawDataMemoryOk

`func (o *X509Certificate2) GetRawDataMemoryOk() (*ByteReadOnlyMemory, bool)`

GetRawDataMemoryOk returns a tuple with the RawDataMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataMemory

`func (o *X509Certificate2) SetRawDataMemory(v ByteReadOnlyMemory)`

SetRawDataMemory sets RawDataMemory field to given value.

### HasRawDataMemory

`func (o *X509Certificate2) HasRawDataMemory() bool`

HasRawDataMemory returns a boolean if a field has been set.

### GetSerialNumber

`func (o *X509Certificate2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *X509Certificate2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *X509Certificate2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *X509Certificate2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *X509Certificate2) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *X509Certificate2) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSignatureAlgorithm

`func (o *X509Certificate2) GetSignatureAlgorithm() Oid`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *X509Certificate2) GetSignatureAlgorithmOk() (*Oid, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *X509Certificate2) SetSignatureAlgorithm(v Oid)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *X509Certificate2) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSubjectName

`func (o *X509Certificate2) GetSubjectName() X500DistinguishedName`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *X509Certificate2) GetSubjectNameOk() (*X500DistinguishedName, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *X509Certificate2) SetSubjectName(v X500DistinguishedName)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *X509Certificate2) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetThumbprint

`func (o *X509Certificate2) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *X509Certificate2) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *X509Certificate2) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *X509Certificate2) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *X509Certificate2) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *X509Certificate2) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetVersion

`func (o *X509Certificate2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *X509Certificate2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *X509Certificate2) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *X509Certificate2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


