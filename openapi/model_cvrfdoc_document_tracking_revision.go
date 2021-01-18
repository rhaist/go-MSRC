/*
 * MSRC Public API
 *
 * This RESTful API can be used to engage the Microsoft Security Response Center (MSRC) in the following ways:    - Get security update summaries and details using the [Common Vulnerability Reporting Framework](https://www.icasi.org/cvrf) (CVRF).    - Report suspected cyberattacks or abuse originating from Microsoft Online Services.    - Notify Microsoft of any planned penetration tests against your Azure assets.    **Sample client code** is available on the Microsoft Security [Updates](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API) and [Engage](https://github.com/Microsoft/MSRC-Microsoft-Engage-API)   Github repositories.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CvrfdocDocumentTrackingRevision struct for CvrfdocDocumentTrackingRevision
type CvrfdocDocumentTrackingRevision struct {
	Number      *string                                     `json:"Number,omitempty"`
	Date        *time.Time                                  `json:"Date,omitempty"`
	Description *CvrfdocDocumentTrackingRevisionDescription `json:"Description,omitempty"`
}

// NewCvrfdocDocumentTrackingRevision instantiates a new CvrfdocDocumentTrackingRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdocDocumentTrackingRevision() *CvrfdocDocumentTrackingRevision {
	this := CvrfdocDocumentTrackingRevision{}
	return &this
}

// NewCvrfdocDocumentTrackingRevisionWithDefaults instantiates a new CvrfdocDocumentTrackingRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocDocumentTrackingRevisionWithDefaults() *CvrfdocDocumentTrackingRevision {
	this := CvrfdocDocumentTrackingRevision{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CvrfdocDocumentTrackingRevision) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentTrackingRevision) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CvrfdocDocumentTrackingRevision) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CvrfdocDocumentTrackingRevision) SetNumber(v string) {
	o.Number = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CvrfdocDocumentTrackingRevision) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentTrackingRevision) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CvrfdocDocumentTrackingRevision) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *CvrfdocDocumentTrackingRevision) SetDate(v time.Time) {
	o.Date = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CvrfdocDocumentTrackingRevision) GetDescription() CvrfdocDocumentTrackingRevisionDescription {
	if o == nil || o.Description == nil {
		var ret CvrfdocDocumentTrackingRevisionDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvrfdocDocumentTrackingRevision) GetDescriptionOk() (*CvrfdocDocumentTrackingRevisionDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CvrfdocDocumentTrackingRevision) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given CvrfdocDocumentTrackingRevisionDescription and assigns it to the Description field.
func (o *CvrfdocDocumentTrackingRevision) SetDescription(v CvrfdocDocumentTrackingRevisionDescription) {
	o.Description = &v
}

func (o CvrfdocDocumentTrackingRevision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["Number"] = o.Number
	}
	if o.Date != nil {
		toSerialize["Date"] = o.Date
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdocDocumentTrackingRevision struct {
	value *CvrfdocDocumentTrackingRevision
	isSet bool
}

func (v NullableCvrfdocDocumentTrackingRevision) Get() *CvrfdocDocumentTrackingRevision {
	return v.value
}

func (v *NullableCvrfdocDocumentTrackingRevision) Set(val *CvrfdocDocumentTrackingRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdocDocumentTrackingRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdocDocumentTrackingRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdocDocumentTrackingRevision(val *CvrfdocDocumentTrackingRevision) *NullableCvrfdocDocumentTrackingRevision {
	return &NullableCvrfdocDocumentTrackingRevision{value: val, isSet: true}
}

func (v NullableCvrfdocDocumentTrackingRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdocDocumentTrackingRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
