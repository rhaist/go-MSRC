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

// Report1 abuse report
type Report1 struct {
	// tag and group reports
	BatchID *string `json:"batchID,omitempty"`
	// IDs of known related cases
	RelatedCases *[]string `json:"relatedCases,omitempty"`
	// a brief summary of what occurred
	ReportNotes string `json:"reportNotes"`
	// [Traffic Light Protocol](https://www.us-cert.gov/tlp) color (RED, AMBER, GREEN, or WHITE) to regulate the level of  disclosure
	TLP *string `json:"TLP,omitempty"`
	// any additional limits on disclosure
	DisclosureNotes *string `json:"disclosureNotes,omitempty"`
	// instances of abuse
	Threats []Threat `json:"threats"`
}

// NewReport1 instantiates a new Report1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport1(reportNotes string, threats []Threat) *Report1 {
	this := Report1{}
	this.ReportNotes = reportNotes
	this.Threats = threats
	return &this
}

// NewReport1WithDefaults instantiates a new Report1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReport1WithDefaults() *Report1 {
	this := Report1{}
	return &this
}

// GetBatchID returns the BatchID field value if set, zero value otherwise.
func (o *Report1) GetBatchID() string {
	if o == nil || o.BatchID == nil {
		var ret string
		return ret
	}
	return *o.BatchID
}

// GetBatchIDOk returns a tuple with the BatchID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report1) GetBatchIDOk() (*string, bool) {
	if o == nil || o.BatchID == nil {
		return nil, false
	}
	return o.BatchID, true
}

// HasBatchID returns a boolean if a field has been set.
func (o *Report1) HasBatchID() bool {
	if o != nil && o.BatchID != nil {
		return true
	}

	return false
}

// SetBatchID gets a reference to the given string and assigns it to the BatchID field.
func (o *Report1) SetBatchID(v string) {
	o.BatchID = &v
}

// GetRelatedCases returns the RelatedCases field value if set, zero value otherwise.
func (o *Report1) GetRelatedCases() []string {
	if o == nil || o.RelatedCases == nil {
		var ret []string
		return ret
	}
	return *o.RelatedCases
}

// GetRelatedCasesOk returns a tuple with the RelatedCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report1) GetRelatedCasesOk() (*[]string, bool) {
	if o == nil || o.RelatedCases == nil {
		return nil, false
	}
	return o.RelatedCases, true
}

// HasRelatedCases returns a boolean if a field has been set.
func (o *Report1) HasRelatedCases() bool {
	if o != nil && o.RelatedCases != nil {
		return true
	}

	return false
}

// SetRelatedCases gets a reference to the given []string and assigns it to the RelatedCases field.
func (o *Report1) SetRelatedCases(v []string) {
	o.RelatedCases = &v
}

// GetReportNotes returns the ReportNotes field value
func (o *Report1) GetReportNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportNotes
}

// GetReportNotesOk returns a tuple with the ReportNotes field value
// and a boolean to check if the value has been set.
func (o *Report1) GetReportNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportNotes, true
}

// SetReportNotes sets field value
func (o *Report1) SetReportNotes(v string) {
	o.ReportNotes = v
}

// GetTLP returns the TLP field value if set, zero value otherwise.
func (o *Report1) GetTLP() string {
	if o == nil || o.TLP == nil {
		var ret string
		return ret
	}
	return *o.TLP
}

// GetTLPOk returns a tuple with the TLP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report1) GetTLPOk() (*string, bool) {
	if o == nil || o.TLP == nil {
		return nil, false
	}
	return o.TLP, true
}

// HasTLP returns a boolean if a field has been set.
func (o *Report1) HasTLP() bool {
	if o != nil && o.TLP != nil {
		return true
	}

	return false
}

// SetTLP gets a reference to the given string and assigns it to the TLP field.
func (o *Report1) SetTLP(v string) {
	o.TLP = &v
}

// GetDisclosureNotes returns the DisclosureNotes field value if set, zero value otherwise.
func (o *Report1) GetDisclosureNotes() string {
	if o == nil || o.DisclosureNotes == nil {
		var ret string
		return ret
	}
	return *o.DisclosureNotes
}

// GetDisclosureNotesOk returns a tuple with the DisclosureNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report1) GetDisclosureNotesOk() (*string, bool) {
	if o == nil || o.DisclosureNotes == nil {
		return nil, false
	}
	return o.DisclosureNotes, true
}

// HasDisclosureNotes returns a boolean if a field has been set.
func (o *Report1) HasDisclosureNotes() bool {
	if o != nil && o.DisclosureNotes != nil {
		return true
	}

	return false
}

// SetDisclosureNotes gets a reference to the given string and assigns it to the DisclosureNotes field.
func (o *Report1) SetDisclosureNotes(v string) {
	o.DisclosureNotes = &v
}

// GetThreats returns the Threats field value
func (o *Report1) GetThreats() []Threat {
	if o == nil {
		var ret []Threat
		return ret
	}

	return o.Threats
}

// GetThreatsOk returns a tuple with the Threats field value
// and a boolean to check if the value has been set.
func (o *Report1) GetThreatsOk() (*[]Threat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threats, true
}

// SetThreats sets field value
func (o *Report1) SetThreats(v []Threat) {
	o.Threats = v
}

func (o Report1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BatchID != nil {
		toSerialize["batchID"] = o.BatchID
	}
	if o.RelatedCases != nil {
		toSerialize["relatedCases"] = o.RelatedCases
	}
	if true {
		toSerialize["reportNotes"] = o.ReportNotes
	}
	if o.TLP != nil {
		toSerialize["TLP"] = o.TLP
	}
	if o.DisclosureNotes != nil {
		toSerialize["disclosureNotes"] = o.DisclosureNotes
	}
	if true {
		toSerialize["threats"] = o.Threats
	}
	return json.Marshal(toSerialize)
}

type NullableReport1 struct {
	value *Report1
	isSet bool
}

func (v NullableReport1) Get() *Report1 {
	return v.value
}

func (v *NullableReport1) Set(val *Report1) {
	v.value = val
	v.isSet = true
}

func (v NullableReport1) IsSet() bool {
	return v.isSet
}

func (v *NullableReport1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport1(val *Report1) *NullableReport1 {
	return &NullableReport1{value: val, isSet: true}
}

func (v NullableReport1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
