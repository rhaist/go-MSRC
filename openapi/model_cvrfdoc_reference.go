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

// CvrfdocReference struct for CvrfdocReference
type CvrfdocReference struct {
	URL         *string                      `json:"URL,omitempty"`
	Description *CvrfdocReferenceDescription `json:"Description,omitempty"`
	Type        *int32                       `json:"Type,omitempty"`
}

// NewCvrfdocReference instantiates a new CvrfdocReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdocReference() *CvrfdocReference {
	this := CvrfdocReference{}
	return &this
}

// NewCvrfdocReferenceWithDefaults instantiates a new CvrfdocReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocReferenceWithDefaults() *CvrfdocReference {
	this := CvrfdocReference{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *CvrfdocReference) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocReference) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *CvrfdocReference) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *CvrfdocReference) SetURL(v string) {
	o.URL = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CvrfdocReference) GetDescription() CvrfdocReferenceDescription {
	if o == nil || o.Description == nil {
		var ret CvrfdocReferenceDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocReference) GetDescriptionOk() (*CvrfdocReferenceDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CvrfdocReference) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given CvrfdocReferenceDescription and assigns it to the Description field.
func (o *CvrfdocReference) SetDescription(v CvrfdocReferenceDescription) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CvrfdocReference) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocReference) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CvrfdocReference) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *CvrfdocReference) SetType(v int32) {
	o.Type = &v
}

func (o CvrfdocReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.URL != nil {
		toSerialize["URL"] = o.URL
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdocReference struct {
	value *CvrfdocReference
	isSet bool
}

func (v NullableCvrfdocReference) Get() *CvrfdocReference {
	return v.value
}

func (v *NullableCvrfdocReference) Set(val *CvrfdocReference) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdocReference) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdocReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdocReference(val *CvrfdocReference) *NullableCvrfdocReference {
	return &NullableCvrfdocReference{value: val, isSet: true}
}

func (v NullableCvrfdocReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdocReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
