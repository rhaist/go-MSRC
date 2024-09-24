/*
MSRC CVRF API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the X509Certificate2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X509Certificate2{}

// X509Certificate2 struct for X509Certificate2
type X509Certificate2 struct {
	Handle             map[string]interface{} `json:"handle,omitempty"`
	Issuer             NullableString         `json:"issuer,omitempty"`
	Subject            NullableString         `json:"subject,omitempty"`
	SerialNumberBytes  *ByteReadOnlyMemory    `json:"serialNumberBytes,omitempty"`
	Archived           *bool                  `json:"archived,omitempty"`
	Extensions         []X509Extension        `json:"extensions,omitempty"`
	FriendlyName       NullableString         `json:"friendlyName,omitempty"`
	HasPrivateKey      *bool                  `json:"hasPrivateKey,omitempty"`
	PrivateKey         *AsymmetricAlgorithm   `json:"privateKey,omitempty"`
	IssuerName         *X500DistinguishedName `json:"issuerName,omitempty"`
	NotAfter           *time.Time             `json:"notAfter,omitempty"`
	NotBefore          *time.Time             `json:"notBefore,omitempty"`
	PublicKey          *PublicKey             `json:"publicKey,omitempty"`
	RawData            NullableString         `json:"rawData,omitempty"`
	RawDataMemory      *ByteReadOnlyMemory    `json:"rawDataMemory,omitempty"`
	SerialNumber       NullableString         `json:"serialNumber,omitempty"`
	SignatureAlgorithm *Oid                   `json:"signatureAlgorithm,omitempty"`
	SubjectName        *X500DistinguishedName `json:"subjectName,omitempty"`
	Thumbprint         NullableString         `json:"thumbprint,omitempty"`
	Version            *int32                 `json:"version,omitempty"`
}

// NewX509Certificate2 instantiates a new X509Certificate2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate2() *X509Certificate2 {
	this := X509Certificate2{}
	return &this
}

// NewX509Certificate2WithDefaults instantiates a new X509Certificate2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509Certificate2WithDefaults() *X509Certificate2 {
	this := X509Certificate2{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *X509Certificate2) GetHandle() map[string]interface{} {
	if o == nil || IsNil(o.Handle) {
		var ret map[string]interface{}
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetHandleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Handle) {
		return map[string]interface{}{}, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *X509Certificate2) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given map[string]interface{} and assigns it to the Handle field.
func (o *X509Certificate2) SetHandle(v map[string]interface{}) {
	o.Handle = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *X509Certificate2) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *X509Certificate2) SetIssuer(v string) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *X509Certificate2) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *X509Certificate2) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *X509Certificate2) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *X509Certificate2) SetSubject(v string) {
	o.Subject.Set(&v)
}

// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *X509Certificate2) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *X509Certificate2) UnsetSubject() {
	o.Subject.Unset()
}

// GetSerialNumberBytes returns the SerialNumberBytes field value if set, zero value otherwise.
func (o *X509Certificate2) GetSerialNumberBytes() ByteReadOnlyMemory {
	if o == nil || IsNil(o.SerialNumberBytes) {
		var ret ByteReadOnlyMemory
		return ret
	}
	return *o.SerialNumberBytes
}

// GetSerialNumberBytesOk returns a tuple with the SerialNumberBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetSerialNumberBytesOk() (*ByteReadOnlyMemory, bool) {
	if o == nil || IsNil(o.SerialNumberBytes) {
		return nil, false
	}
	return o.SerialNumberBytes, true
}

// HasSerialNumberBytes returns a boolean if a field has been set.
func (o *X509Certificate2) HasSerialNumberBytes() bool {
	if o != nil && !IsNil(o.SerialNumberBytes) {
		return true
	}

	return false
}

// SetSerialNumberBytes gets a reference to the given ByteReadOnlyMemory and assigns it to the SerialNumberBytes field.
func (o *X509Certificate2) SetSerialNumberBytes(v ByteReadOnlyMemory) {
	o.SerialNumberBytes = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *X509Certificate2) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *X509Certificate2) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *X509Certificate2) SetArchived(v bool) {
	o.Archived = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetExtensions() []X509Extension {
	if o == nil {
		var ret []X509Extension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetExtensionsOk() ([]X509Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *X509Certificate2) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []X509Extension and assigns it to the Extensions field.
func (o *X509Certificate2) SetExtensions(v []X509Extension) {
	o.Extensions = v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName.Get()) {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *X509Certificate2) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *X509Certificate2) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *X509Certificate2) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *X509Certificate2) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetHasPrivateKey returns the HasPrivateKey field value if set, zero value otherwise.
func (o *X509Certificate2) GetHasPrivateKey() bool {
	if o == nil || IsNil(o.HasPrivateKey) {
		var ret bool
		return ret
	}
	return *o.HasPrivateKey
}

// GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetHasPrivateKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPrivateKey) {
		return nil, false
	}
	return o.HasPrivateKey, true
}

// HasHasPrivateKey returns a boolean if a field has been set.
func (o *X509Certificate2) HasHasPrivateKey() bool {
	if o != nil && !IsNil(o.HasPrivateKey) {
		return true
	}

	return false
}

// SetHasPrivateKey gets a reference to the given bool and assigns it to the HasPrivateKey field.
func (o *X509Certificate2) SetHasPrivateKey(v bool) {
	o.HasPrivateKey = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *X509Certificate2) GetPrivateKey() AsymmetricAlgorithm {
	if o == nil || IsNil(o.PrivateKey) {
		var ret AsymmetricAlgorithm
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetPrivateKeyOk() (*AsymmetricAlgorithm, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *X509Certificate2) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given AsymmetricAlgorithm and assigns it to the PrivateKey field.
func (o *X509Certificate2) SetPrivateKey(v AsymmetricAlgorithm) {
	o.PrivateKey = &v
}

// GetIssuerName returns the IssuerName field value if set, zero value otherwise.
func (o *X509Certificate2) GetIssuerName() X500DistinguishedName {
	if o == nil || IsNil(o.IssuerName) {
		var ret X500DistinguishedName
		return ret
	}
	return *o.IssuerName
}

// GetIssuerNameOk returns a tuple with the IssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetIssuerNameOk() (*X500DistinguishedName, bool) {
	if o == nil || IsNil(o.IssuerName) {
		return nil, false
	}
	return o.IssuerName, true
}

// HasIssuerName returns a boolean if a field has been set.
func (o *X509Certificate2) HasIssuerName() bool {
	if o != nil && !IsNil(o.IssuerName) {
		return true
	}

	return false
}

// SetIssuerName gets a reference to the given X500DistinguishedName and assigns it to the IssuerName field.
func (o *X509Certificate2) SetIssuerName(v X500DistinguishedName) {
	o.IssuerName = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509Certificate2) GetNotAfter() time.Time {
	if o == nil || IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509Certificate2) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509Certificate2) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509Certificate2) GetNotBefore() time.Time {
	if o == nil || IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509Certificate2) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509Certificate2) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *X509Certificate2) GetPublicKey() PublicKey {
	if o == nil || IsNil(o.PublicKey) {
		var ret PublicKey
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetPublicKeyOk() (*PublicKey, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *X509Certificate2) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given PublicKey and assigns it to the PublicKey field.
func (o *X509Certificate2) SetPublicKey(v PublicKey) {
	o.PublicKey = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetRawData() string {
	if o == nil || IsNil(o.RawData.Get()) {
		var ret string
		return ret
	}
	return *o.RawData.Get()
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetRawDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawData.Get(), o.RawData.IsSet()
}

// HasRawData returns a boolean if a field has been set.
func (o *X509Certificate2) HasRawData() bool {
	if o != nil && o.RawData.IsSet() {
		return true
	}

	return false
}

// SetRawData gets a reference to the given NullableString and assigns it to the RawData field.
func (o *X509Certificate2) SetRawData(v string) {
	o.RawData.Set(&v)
}

// SetRawDataNil sets the value for RawData to be an explicit nil
func (o *X509Certificate2) SetRawDataNil() {
	o.RawData.Set(nil)
}

// UnsetRawData ensures that no value is present for RawData, not even an explicit nil
func (o *X509Certificate2) UnsetRawData() {
	o.RawData.Unset()
}

// GetRawDataMemory returns the RawDataMemory field value if set, zero value otherwise.
func (o *X509Certificate2) GetRawDataMemory() ByteReadOnlyMemory {
	if o == nil || IsNil(o.RawDataMemory) {
		var ret ByteReadOnlyMemory
		return ret
	}
	return *o.RawDataMemory
}

// GetRawDataMemoryOk returns a tuple with the RawDataMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetRawDataMemoryOk() (*ByteReadOnlyMemory, bool) {
	if o == nil || IsNil(o.RawDataMemory) {
		return nil, false
	}
	return o.RawDataMemory, true
}

// HasRawDataMemory returns a boolean if a field has been set.
func (o *X509Certificate2) HasRawDataMemory() bool {
	if o != nil && !IsNil(o.RawDataMemory) {
		return true
	}

	return false
}

// SetRawDataMemory gets a reference to the given ByteReadOnlyMemory and assigns it to the RawDataMemory field.
func (o *X509Certificate2) SetRawDataMemory(v ByteReadOnlyMemory) {
	o.RawDataMemory = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber.Get()) {
		var ret string
		return ret
	}
	return *o.SerialNumber.Get()
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SerialNumber.Get(), o.SerialNumber.IsSet()
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *X509Certificate2) HasSerialNumber() bool {
	if o != nil && o.SerialNumber.IsSet() {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given NullableString and assigns it to the SerialNumber field.
func (o *X509Certificate2) SetSerialNumber(v string) {
	o.SerialNumber.Set(&v)
}

// SetSerialNumberNil sets the value for SerialNumber to be an explicit nil
func (o *X509Certificate2) SetSerialNumberNil() {
	o.SerialNumber.Set(nil)
}

// UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
func (o *X509Certificate2) UnsetSerialNumber() {
	o.SerialNumber.Unset()
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *X509Certificate2) GetSignatureAlgorithm() Oid {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret Oid
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetSignatureAlgorithmOk() (*Oid, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *X509Certificate2) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given Oid and assigns it to the SignatureAlgorithm field.
func (o *X509Certificate2) SetSignatureAlgorithm(v Oid) {
	o.SignatureAlgorithm = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *X509Certificate2) GetSubjectName() X500DistinguishedName {
	if o == nil || IsNil(o.SubjectName) {
		var ret X500DistinguishedName
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetSubjectNameOk() (*X500DistinguishedName, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *X509Certificate2) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given X500DistinguishedName and assigns it to the SubjectName field.
func (o *X509Certificate2) SetSubjectName(v X500DistinguishedName) {
	o.SubjectName = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate2) GetThumbprint() string {
	if o == nil || IsNil(o.Thumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbprint.Get()
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate2) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbprint.Get(), o.Thumbprint.IsSet()
}

// HasThumbprint returns a boolean if a field has been set.
func (o *X509Certificate2) HasThumbprint() bool {
	if o != nil && o.Thumbprint.IsSet() {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given NullableString and assigns it to the Thumbprint field.
func (o *X509Certificate2) SetThumbprint(v string) {
	o.Thumbprint.Set(&v)
}

// SetThumbprintNil sets the value for Thumbprint to be an explicit nil
func (o *X509Certificate2) SetThumbprintNil() {
	o.Thumbprint.Set(nil)
}

// UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
func (o *X509Certificate2) UnsetThumbprint() {
	o.Thumbprint.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *X509Certificate2) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate2) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *X509Certificate2) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *X509Certificate2) SetVersion(v int32) {
	o.Version = &v
}

func (o X509Certificate2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o X509Certificate2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if !IsNil(o.SerialNumberBytes) {
		toSerialize["serialNumberBytes"] = o.SerialNumberBytes
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendlyName"] = o.FriendlyName.Get()
	}
	if !IsNil(o.HasPrivateKey) {
		toSerialize["hasPrivateKey"] = o.HasPrivateKey
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.IssuerName) {
		toSerialize["issuerName"] = o.IssuerName
	}
	if !IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.RawData.IsSet() {
		toSerialize["rawData"] = o.RawData.Get()
	}
	if !IsNil(o.RawDataMemory) {
		toSerialize["rawDataMemory"] = o.RawDataMemory
	}
	if o.SerialNumber.IsSet() {
		toSerialize["serialNumber"] = o.SerialNumber.Get()
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if o.Thumbprint.IsSet() {
		toSerialize["thumbprint"] = o.Thumbprint.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableX509Certificate2 struct {
	value *X509Certificate2
	isSet bool
}

func (v NullableX509Certificate2) Get() *X509Certificate2 {
	return v.value
}

func (v *NullableX509Certificate2) Set(val *X509Certificate2) {
	v.value = val
	v.isSet = true
}

func (v NullableX509Certificate2) IsSet() bool {
	return v.isSet
}

func (v *NullableX509Certificate2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509Certificate2(val *X509Certificate2) *NullableX509Certificate2 {
	return &NullableX509Certificate2{value: val, isSet: true}
}

func (v NullableX509Certificate2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509Certificate2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
