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

// Update summary information about a security release (CVRF document)
type Update struct {
	// unique identifier for CVRF document, often yyyy-mmm
	ID string `json:"ID"`
	// alternative ID for CVRF document
	Alias string `json:"Alias"`
	// title of CVRF document
	DocumentTitle string `json:"DocumentTitle"`
	// aggregate severity
	Severity *string `json:"Severity,omitempty"`
	// date that the CVRF document was initially released
	InitialReleaseDate time.Time `json:"InitialReleaseDate"`
	// date that the CVRF document was last modified
	CurrentReleaseDate time.Time `json:"CurrentReleaseDate"`
	// URL of CVRF document (for CVRF endpoint)
	CvrfUrl string `json:"CvrfUrl"`
}

// NewUpdate instantiates a new Update object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdate(iD string, alias string, documentTitle string, initialReleaseDate time.Time, currentReleaseDate time.Time, cvrfUrl string) *Update {
	this := Update{}
	this.ID = iD
	this.Alias = alias
	this.DocumentTitle = documentTitle
	this.InitialReleaseDate = initialReleaseDate
	this.CurrentReleaseDate = currentReleaseDate
	this.CvrfUrl = cvrfUrl
	return &this
}

// NewUpdateWithDefaults instantiates a new Update object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWithDefaults() *Update {
	this := Update{}
	return &this
}

// GetID returns the ID field value
func (o *Update) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *Update) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *Update) SetID(v string) {
	o.ID = v
}

// GetAlias returns the Alias field value
func (o *Update) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *Update) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *Update) SetAlias(v string) {
	o.Alias = v
}

// GetDocumentTitle returns the DocumentTitle field value
func (o *Update) GetDocumentTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentTitle
}

// GetDocumentTitleOk returns a tuple with the DocumentTitle field value
// and a boolean to check if the value has been set.
func (o *Update) GetDocumentTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentTitle, true
}

// SetDocumentTitle sets field value
func (o *Update) SetDocumentTitle(v string) {
	o.DocumentTitle = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Update) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Update) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *Update) SetSeverity(v string) {
	o.Severity = &v
}

// GetInitialReleaseDate returns the InitialReleaseDate field value
func (o *Update) GetInitialReleaseDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.InitialReleaseDate
}

// GetInitialReleaseDateOk returns a tuple with the InitialReleaseDate field value
// and a boolean to check if the value has been set.
func (o *Update) GetInitialReleaseDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialReleaseDate, true
}

// SetInitialReleaseDate sets field value
func (o *Update) SetInitialReleaseDate(v time.Time) {
	o.InitialReleaseDate = v
}

// GetCurrentReleaseDate returns the CurrentReleaseDate field value
func (o *Update) GetCurrentReleaseDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentReleaseDate
}

// GetCurrentReleaseDateOk returns a tuple with the CurrentReleaseDate field value
// and a boolean to check if the value has been set.
func (o *Update) GetCurrentReleaseDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentReleaseDate, true
}

// SetCurrentReleaseDate sets field value
func (o *Update) SetCurrentReleaseDate(v time.Time) {
	o.CurrentReleaseDate = v
}

// GetCvrfUrl returns the CvrfUrl field value
func (o *Update) GetCvrfUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CvrfUrl
}

// GetCvrfUrlOk returns a tuple with the CvrfUrl field value
// and a boolean to check if the value has been set.
func (o *Update) GetCvrfUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CvrfUrl, true
}

// SetCvrfUrl sets field value
func (o *Update) SetCvrfUrl(v string) {
	o.CvrfUrl = v
}

func (o Update) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ID"] = o.ID
	}
	if true {
		toSerialize["Alias"] = o.Alias
	}
	if true {
		toSerialize["DocumentTitle"] = o.DocumentTitle
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if true {
		toSerialize["InitialReleaseDate"] = o.InitialReleaseDate
	}
	if true {
		toSerialize["CurrentReleaseDate"] = o.CurrentReleaseDate
	}
	if true {
		toSerialize["CvrfUrl"] = o.CvrfUrl
	}
	return json.Marshal(toSerialize)
}

type NullableUpdate struct {
	value *Update
	isSet bool
}

func (v NullableUpdate) Get() *Update {
	return v.value
}

func (v *NullableUpdate) Set(val *Update) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdate(val *Update) *NullableUpdate {
	return &NullableUpdate{value: val, isSet: true}
}

func (v NullableUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
