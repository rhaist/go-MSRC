/*
 * MSRC CVRF API
 *
 * Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_
 *
 * API version: v2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CvrfdocDocumentPublisher struct for CvrfdocDocumentPublisher
type CvrfdocDocumentPublisher struct {
	ContactDetails   *CvrfdocDocumentPublisherContactDetails   `json:"ContactDetails,omitempty"`
	IssuingAuthority *CvrfdocDocumentPublisherIssuingAuthority `json:"IssuingAuthority,omitempty"`
	Type             *int32                                    `json:"Type,omitempty"`
	VendorID         *string                                   `json:"VendorID,omitempty"`
}

// NewCvrfdocDocumentPublisher instantiates a new CvrfdocDocumentPublisher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdocDocumentPublisher() *CvrfdocDocumentPublisher {
	this := CvrfdocDocumentPublisher{}
	return &this
}

// NewCvrfdocDocumentPublisherWithDefaults instantiates a new CvrfdocDocumentPublisher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocDocumentPublisherWithDefaults() *CvrfdocDocumentPublisher {
	this := CvrfdocDocumentPublisher{}
	return &this
}

// GetContactDetails returns the ContactDetails field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisher) GetContactDetails() CvrfdocDocumentPublisherContactDetails {
	if o == nil || o.ContactDetails == nil {
		var ret CvrfdocDocumentPublisherContactDetails
		return ret
	}
	return *o.ContactDetails
}

// GetContactDetailsOk returns a tuple with the ContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisher) GetContactDetailsOk() (*CvrfdocDocumentPublisherContactDetails, bool) {
	if o == nil || o.ContactDetails == nil {
		return nil, false
	}
	return o.ContactDetails, true
}

// HasContactDetails returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisher) HasContactDetails() bool {
	if o != nil && o.ContactDetails != nil {
		return true
	}

	return false
}

// SetContactDetails gets a reference to the given CvrfdocDocumentPublisherContactDetails and assigns it to the ContactDetails field.
func (o *CvrfdocDocumentPublisher) SetContactDetails(v CvrfdocDocumentPublisherContactDetails) {
	o.ContactDetails = &v
}

// GetIssuingAuthority returns the IssuingAuthority field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisher) GetIssuingAuthority() CvrfdocDocumentPublisherIssuingAuthority {
	if o == nil || o.IssuingAuthority == nil {
		var ret CvrfdocDocumentPublisherIssuingAuthority
		return ret
	}
	return *o.IssuingAuthority
}

// GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisher) GetIssuingAuthorityOk() (*CvrfdocDocumentPublisherIssuingAuthority, bool) {
	if o == nil || o.IssuingAuthority == nil {
		return nil, false
	}
	return o.IssuingAuthority, true
}

// HasIssuingAuthority returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisher) HasIssuingAuthority() bool {
	if o != nil && o.IssuingAuthority != nil {
		return true
	}

	return false
}

// SetIssuingAuthority gets a reference to the given CvrfdocDocumentPublisherIssuingAuthority and assigns it to the IssuingAuthority field.
func (o *CvrfdocDocumentPublisher) SetIssuingAuthority(v CvrfdocDocumentPublisherIssuingAuthority) {
	o.IssuingAuthority = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisher) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisher) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisher) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *CvrfdocDocumentPublisher) SetType(v int32) {
	o.Type = &v
}

// GetVendorID returns the VendorID field value if set, zero value otherwise.
func (o *CvrfdocDocumentPublisher) GetVendorID() string {
	if o == nil || o.VendorID == nil {
		var ret string
		return ret
	}
	return *o.VendorID
}

// GetVendorIDOk returns a tuple with the VendorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentPublisher) GetVendorIDOk() (*string, bool) {
	if o == nil || o.VendorID == nil {
		return nil, false
	}
	return o.VendorID, true
}

// HasVendorID returns a boolean if a field has been set.
func (o *CvrfdocDocumentPublisher) HasVendorID() bool {
	if o != nil && o.VendorID != nil {
		return true
	}

	return false
}

// SetVendorID gets a reference to the given string and assigns it to the VendorID field.
func (o *CvrfdocDocumentPublisher) SetVendorID(v string) {
	o.VendorID = &v
}

func (o CvrfdocDocumentPublisher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactDetails != nil {
		toSerialize["ContactDetails"] = o.ContactDetails
	}
	if o.IssuingAuthority != nil {
		toSerialize["IssuingAuthority"] = o.IssuingAuthority
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VendorID != nil {
		toSerialize["VendorID"] = o.VendorID
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdocDocumentPublisher struct {
	value *CvrfdocDocumentPublisher
	isSet bool
}

func (v NullableCvrfdocDocumentPublisher) Get() *CvrfdocDocumentPublisher {
	return v.value
}

func (v *NullableCvrfdocDocumentPublisher) Set(val *CvrfdocDocumentPublisher) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdocDocumentPublisher) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdocDocumentPublisher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdocDocumentPublisher(val *CvrfdocDocumentPublisher) *NullableCvrfdocDocumentPublisher {
	return &NullableCvrfdocDocumentPublisher{value: val, isSet: true}
}

func (v NullableCvrfdocDocumentPublisher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdocDocumentPublisher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
