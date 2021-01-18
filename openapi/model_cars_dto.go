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
)

// CarsDto a Common Abuse Reporting System (CARS) submission
type CarsDto struct {
	// is this a test submission? If so, the API will validate the input but won't submit it to MSRC for action
	TestSubmission bool `json:"testSubmission"`
	// do the reports include detailed fields? This can be false only for form submissions (API users should ignore it)
	CarsAdvancedCheck *bool        `json:"carsAdvancedCheck,omitempty"`
	ReporterInfo      ReporterInfo `json:"reporterInfo"`
	// reports of abuse
	Reports []Report `json:"reports"`
}

// NewCarsDto instantiates a new CarsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarsDto(testSubmission bool, reporterInfo ReporterInfo, reports []Report) *CarsDto {
	this := CarsDto{}
	this.TestSubmission = testSubmission
	this.ReporterInfo = reporterInfo
	this.Reports = reports
	return &this
}

// NewCarsDtoWithDefaults instantiates a new CarsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarsDtoWithDefaults() *CarsDto {
	this := CarsDto{}
	return &this
}

// GetTestSubmission returns the TestSubmission field value
func (o *CarsDto) GetTestSubmission() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TestSubmission
}

// GetTestSubmissionOk returns a tuple with the TestSubmission field value
// and a boolean to check if the value has been set.
func (o *CarsDto) GetTestSubmissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSubmission, true
}

// SetTestSubmission sets field value
func (o *CarsDto) SetTestSubmission(v bool) {
	o.TestSubmission = v
}

// GetCarsAdvancedCheck returns the CarsAdvancedCheck field value if set, zero value otherwise.
func (o *CarsDto) GetCarsAdvancedCheck() bool {
	if o == nil || o.CarsAdvancedCheck == nil {
		var ret bool
		return ret
	}
	return *o.CarsAdvancedCheck
}

// GetCarsAdvancedCheckOk returns a tuple with the CarsAdvancedCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarsDto) GetCarsAdvancedCheckOk() (*bool, bool) {
	if o == nil || o.CarsAdvancedCheck == nil {
		return nil, false
	}
	return o.CarsAdvancedCheck, true
}

// HasCarsAdvancedCheck returns a boolean if a field has been set.
func (o *CarsDto) HasCarsAdvancedCheck() bool {
	if o != nil && o.CarsAdvancedCheck != nil {
		return true
	}

	return false
}

// SetCarsAdvancedCheck gets a reference to the given bool and assigns it to the CarsAdvancedCheck field.
func (o *CarsDto) SetCarsAdvancedCheck(v bool) {
	o.CarsAdvancedCheck = &v
}

// GetReporterInfo returns the ReporterInfo field value
func (o *CarsDto) GetReporterInfo() ReporterInfo {
	if o == nil {
		var ret ReporterInfo
		return ret
	}

	return o.ReporterInfo
}

// GetReporterInfoOk returns a tuple with the ReporterInfo field value
// and a boolean to check if the value has been set.
func (o *CarsDto) GetReporterInfoOk() (*ReporterInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReporterInfo, true
}

// SetReporterInfo sets field value
func (o *CarsDto) SetReporterInfo(v ReporterInfo) {
	o.ReporterInfo = v
}

// GetReports returns the Reports field value
func (o *CarsDto) GetReports() []Report {
	if o == nil {
		var ret []Report
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *CarsDto) GetReportsOk() (*[]Report, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reports, true
}

// SetReports sets field value
func (o *CarsDto) SetReports(v []Report) {
	o.Reports = v
}

func (o CarsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["testSubmission"] = o.TestSubmission
	}
	if o.CarsAdvancedCheck != nil {
		toSerialize["carsAdvancedCheck"] = o.CarsAdvancedCheck
	}
	if true {
		toSerialize["reporterInfo"] = o.ReporterInfo
	}
	if true {
		toSerialize["reports"] = o.Reports
	}
	return json.Marshal(toSerialize)
}

type NullableCarsDto struct {
	value *CarsDto
	isSet bool
}

func (v NullableCarsDto) Get() *CarsDto {
	return v.value
}

func (v *NullableCarsDto) Set(val *CarsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCarsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCarsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarsDto(val *CarsDto) *NullableCarsDto {
	return &NullableCarsDto{value: val, isSet: true}
}

func (v NullableCarsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}