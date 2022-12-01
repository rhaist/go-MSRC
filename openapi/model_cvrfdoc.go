/*
MSRC CVRF API

Get security updates programmatically using this RESTful API. Sample client code is available on the [Microsoft Security Updates GitHub](https://github.com/microsoft/MSRC-Microsoft-Security-Updates-API). For more details, please visit [msrc.microsoft.com/update-guide](https://msrc.microsoft.com/update-guide).    _**NOTE: If you're looking for the Engage API (CARS), please refer to the new [Abuse Reporting developer portal](https://msrc.microsoft.com/report/developer).**_

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Cvrfdoc struct for Cvrfdoc
type Cvrfdoc struct {
	DocumentTitle        *CvrfdocDocumentTitle        `json:"DocumentTitle,omitempty"`
	DocumentType         *CvrfdocDocumentType         `json:"DocumentType,omitempty"`
	DocumentPublisher    *CvrfdocDocumentPublisher    `json:"DocumentPublisher,omitempty"`
	DocumentTracking     *CvrfdocDocumentTracking     `json:"DocumentTracking,omitempty"`
	DocumentNotes        []CvrfdocNote                `json:"DocumentNotes,omitempty"`
	DocumentDistribution *CvrfdocDocumentDistribution `json:"DocumentDistribution,omitempty"`
	AggregateSeverity    *CvrfdocAggregateSeverity    `json:"AggregateSeverity,omitempty"`
	DocumentReferences   []CvrfdocReference           `json:"DocumentReferences,omitempty"`
	Acknowledgments      []CvrfdocAcknowledgment      `json:"Acknowledgments,omitempty"`
	ProductTree          *ProductTree                 `json:"ProductTree,omitempty"`
	Vulnerability        []Vulnerability              `json:"Vulnerability,omitempty"`
}

// NewCvrfdoc instantiates a new Cvrfdoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvrfdoc() *Cvrfdoc {
	this := Cvrfdoc{}
	return &this
}

// NewCvrfdocWithDefaults instantiates a new Cvrfdoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvrfdocWithDefaults() *Cvrfdoc {
	this := Cvrfdoc{}
	return &this
}

// GetDocumentTitle returns the DocumentTitle field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentTitle() CvrfdocDocumentTitle {
	if o == nil || isNil(o.DocumentTitle) {
		var ret CvrfdocDocumentTitle
		return ret
	}
	return *o.DocumentTitle
}

// GetDocumentTitleOk returns a tuple with the DocumentTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentTitleOk() (*CvrfdocDocumentTitle, bool) {
	if o == nil || isNil(o.DocumentTitle) {
		return nil, false
	}
	return o.DocumentTitle, true
}

// HasDocumentTitle returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentTitle() bool {
	if o != nil && !isNil(o.DocumentTitle) {
		return true
	}

	return false
}

// SetDocumentTitle gets a reference to the given CvrfdocDocumentTitle and assigns it to the DocumentTitle field.
func (o *Cvrfdoc) SetDocumentTitle(v CvrfdocDocumentTitle) {
	o.DocumentTitle = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentType() CvrfdocDocumentType {
	if o == nil || isNil(o.DocumentType) {
		var ret CvrfdocDocumentType
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentTypeOk() (*CvrfdocDocumentType, bool) {
	if o == nil || isNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentType() bool {
	if o != nil && !isNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given CvrfdocDocumentType and assigns it to the DocumentType field.
func (o *Cvrfdoc) SetDocumentType(v CvrfdocDocumentType) {
	o.DocumentType = &v
}

// GetDocumentPublisher returns the DocumentPublisher field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentPublisher() CvrfdocDocumentPublisher {
	if o == nil || isNil(o.DocumentPublisher) {
		var ret CvrfdocDocumentPublisher
		return ret
	}
	return *o.DocumentPublisher
}

// GetDocumentPublisherOk returns a tuple with the DocumentPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentPublisherOk() (*CvrfdocDocumentPublisher, bool) {
	if o == nil || isNil(o.DocumentPublisher) {
		return nil, false
	}
	return o.DocumentPublisher, true
}

// HasDocumentPublisher returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentPublisher() bool {
	if o != nil && !isNil(o.DocumentPublisher) {
		return true
	}

	return false
}

// SetDocumentPublisher gets a reference to the given CvrfdocDocumentPublisher and assigns it to the DocumentPublisher field.
func (o *Cvrfdoc) SetDocumentPublisher(v CvrfdocDocumentPublisher) {
	o.DocumentPublisher = &v
}

// GetDocumentTracking returns the DocumentTracking field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentTracking() CvrfdocDocumentTracking {
	if o == nil || isNil(o.DocumentTracking) {
		var ret CvrfdocDocumentTracking
		return ret
	}
	return *o.DocumentTracking
}

// GetDocumentTrackingOk returns a tuple with the DocumentTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentTrackingOk() (*CvrfdocDocumentTracking, bool) {
	if o == nil || isNil(o.DocumentTracking) {
		return nil, false
	}
	return o.DocumentTracking, true
}

// HasDocumentTracking returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentTracking() bool {
	if o != nil && !isNil(o.DocumentTracking) {
		return true
	}

	return false
}

// SetDocumentTracking gets a reference to the given CvrfdocDocumentTracking and assigns it to the DocumentTracking field.
func (o *Cvrfdoc) SetDocumentTracking(v CvrfdocDocumentTracking) {
	o.DocumentTracking = &v
}

// GetDocumentNotes returns the DocumentNotes field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentNotes() []CvrfdocNote {
	if o == nil || isNil(o.DocumentNotes) {
		var ret []CvrfdocNote
		return ret
	}
	return o.DocumentNotes
}

// GetDocumentNotesOk returns a tuple with the DocumentNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentNotesOk() ([]CvrfdocNote, bool) {
	if o == nil || isNil(o.DocumentNotes) {
		return nil, false
	}
	return o.DocumentNotes, true
}

// HasDocumentNotes returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentNotes() bool {
	if o != nil && !isNil(o.DocumentNotes) {
		return true
	}

	return false
}

// SetDocumentNotes gets a reference to the given []CvrfdocNote and assigns it to the DocumentNotes field.
func (o *Cvrfdoc) SetDocumentNotes(v []CvrfdocNote) {
	o.DocumentNotes = v
}

// GetDocumentDistribution returns the DocumentDistribution field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentDistribution() CvrfdocDocumentDistribution {
	if o == nil || isNil(o.DocumentDistribution) {
		var ret CvrfdocDocumentDistribution
		return ret
	}
	return *o.DocumentDistribution
}

// GetDocumentDistributionOk returns a tuple with the DocumentDistribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentDistributionOk() (*CvrfdocDocumentDistribution, bool) {
	if o == nil || isNil(o.DocumentDistribution) {
		return nil, false
	}
	return o.DocumentDistribution, true
}

// HasDocumentDistribution returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentDistribution() bool {
	if o != nil && !isNil(o.DocumentDistribution) {
		return true
	}

	return false
}

// SetDocumentDistribution gets a reference to the given CvrfdocDocumentDistribution and assigns it to the DocumentDistribution field.
func (o *Cvrfdoc) SetDocumentDistribution(v CvrfdocDocumentDistribution) {
	o.DocumentDistribution = &v
}

// GetAggregateSeverity returns the AggregateSeverity field value if set, zero value otherwise.
func (o *Cvrfdoc) GetAggregateSeverity() CvrfdocAggregateSeverity {
	if o == nil || isNil(o.AggregateSeverity) {
		var ret CvrfdocAggregateSeverity
		return ret
	}
	return *o.AggregateSeverity
}

// GetAggregateSeverityOk returns a tuple with the AggregateSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetAggregateSeverityOk() (*CvrfdocAggregateSeverity, bool) {
	if o == nil || isNil(o.AggregateSeverity) {
		return nil, false
	}
	return o.AggregateSeverity, true
}

// HasAggregateSeverity returns a boolean if a field has been set.
func (o *Cvrfdoc) HasAggregateSeverity() bool {
	if o != nil && !isNil(o.AggregateSeverity) {
		return true
	}

	return false
}

// SetAggregateSeverity gets a reference to the given CvrfdocAggregateSeverity and assigns it to the AggregateSeverity field.
func (o *Cvrfdoc) SetAggregateSeverity(v CvrfdocAggregateSeverity) {
	o.AggregateSeverity = &v
}

// GetDocumentReferences returns the DocumentReferences field value if set, zero value otherwise.
func (o *Cvrfdoc) GetDocumentReferences() []CvrfdocReference {
	if o == nil || isNil(o.DocumentReferences) {
		var ret []CvrfdocReference
		return ret
	}
	return o.DocumentReferences
}

// GetDocumentReferencesOk returns a tuple with the DocumentReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetDocumentReferencesOk() ([]CvrfdocReference, bool) {
	if o == nil || isNil(o.DocumentReferences) {
		return nil, false
	}
	return o.DocumentReferences, true
}

// HasDocumentReferences returns a boolean if a field has been set.
func (o *Cvrfdoc) HasDocumentReferences() bool {
	if o != nil && !isNil(o.DocumentReferences) {
		return true
	}

	return false
}

// SetDocumentReferences gets a reference to the given []CvrfdocReference and assigns it to the DocumentReferences field.
func (o *Cvrfdoc) SetDocumentReferences(v []CvrfdocReference) {
	o.DocumentReferences = v
}

// GetAcknowledgments returns the Acknowledgments field value if set, zero value otherwise.
func (o *Cvrfdoc) GetAcknowledgments() []CvrfdocAcknowledgment {
	if o == nil || isNil(o.Acknowledgments) {
		var ret []CvrfdocAcknowledgment
		return ret
	}
	return o.Acknowledgments
}

// GetAcknowledgmentsOk returns a tuple with the Acknowledgments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetAcknowledgmentsOk() ([]CvrfdocAcknowledgment, bool) {
	if o == nil || isNil(o.Acknowledgments) {
		return nil, false
	}
	return o.Acknowledgments, true
}

// HasAcknowledgments returns a boolean if a field has been set.
func (o *Cvrfdoc) HasAcknowledgments() bool {
	if o != nil && !isNil(o.Acknowledgments) {
		return true
	}

	return false
}

// SetAcknowledgments gets a reference to the given []CvrfdocAcknowledgment and assigns it to the Acknowledgments field.
func (o *Cvrfdoc) SetAcknowledgments(v []CvrfdocAcknowledgment) {
	o.Acknowledgments = v
}

// GetProductTree returns the ProductTree field value if set, zero value otherwise.
func (o *Cvrfdoc) GetProductTree() ProductTree {
	if o == nil || isNil(o.ProductTree) {
		var ret ProductTree
		return ret
	}
	return *o.ProductTree
}

// GetProductTreeOk returns a tuple with the ProductTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetProductTreeOk() (*ProductTree, bool) {
	if o == nil || isNil(o.ProductTree) {
		return nil, false
	}
	return o.ProductTree, true
}

// HasProductTree returns a boolean if a field has been set.
func (o *Cvrfdoc) HasProductTree() bool {
	if o != nil && !isNil(o.ProductTree) {
		return true
	}

	return false
}

// SetProductTree gets a reference to the given ProductTree and assigns it to the ProductTree field.
func (o *Cvrfdoc) SetProductTree(v ProductTree) {
	o.ProductTree = &v
}

// GetVulnerability returns the Vulnerability field value if set, zero value otherwise.
func (o *Cvrfdoc) GetVulnerability() []Vulnerability {
	if o == nil || isNil(o.Vulnerability) {
		var ret []Vulnerability
		return ret
	}
	return o.Vulnerability
}

// GetVulnerabilityOk returns a tuple with the Vulnerability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cvrfdoc) GetVulnerabilityOk() ([]Vulnerability, bool) {
	if o == nil || isNil(o.Vulnerability) {
		return nil, false
	}
	return o.Vulnerability, true
}

// HasVulnerability returns a boolean if a field has been set.
func (o *Cvrfdoc) HasVulnerability() bool {
	if o != nil && !isNil(o.Vulnerability) {
		return true
	}

	return false
}

// SetVulnerability gets a reference to the given []Vulnerability and assigns it to the Vulnerability field.
func (o *Cvrfdoc) SetVulnerability(v []Vulnerability) {
	o.Vulnerability = v
}

func (o Cvrfdoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DocumentTitle) {
		toSerialize["DocumentTitle"] = o.DocumentTitle
	}
	if !isNil(o.DocumentType) {
		toSerialize["DocumentType"] = o.DocumentType
	}
	if !isNil(o.DocumentPublisher) {
		toSerialize["DocumentPublisher"] = o.DocumentPublisher
	}
	if !isNil(o.DocumentTracking) {
		toSerialize["DocumentTracking"] = o.DocumentTracking
	}
	if !isNil(o.DocumentNotes) {
		toSerialize["DocumentNotes"] = o.DocumentNotes
	}
	if !isNil(o.DocumentDistribution) {
		toSerialize["DocumentDistribution"] = o.DocumentDistribution
	}
	if !isNil(o.AggregateSeverity) {
		toSerialize["AggregateSeverity"] = o.AggregateSeverity
	}
	if !isNil(o.DocumentReferences) {
		toSerialize["DocumentReferences"] = o.DocumentReferences
	}
	if !isNil(o.Acknowledgments) {
		toSerialize["Acknowledgments"] = o.Acknowledgments
	}
	if !isNil(o.ProductTree) {
		toSerialize["ProductTree"] = o.ProductTree
	}
	if !isNil(o.Vulnerability) {
		toSerialize["Vulnerability"] = o.Vulnerability
	}
	return json.Marshal(toSerialize)
}

type NullableCvrfdoc struct {
	value *Cvrfdoc
	isSet bool
}

func (v NullableCvrfdoc) Get() *Cvrfdoc {
	return v.value
}

func (v *NullableCvrfdoc) Set(val *Cvrfdoc) {
	v.value = val
	v.isSet = true
}

func (v NullableCvrfdoc) IsSet() bool {
	return v.isSet
}

func (v *NullableCvrfdoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvrfdoc(val *Cvrfdoc) *NullableCvrfdoc {
	return &NullableCvrfdoc{value: val, isSet: true}
}

func (v NullableCvrfdoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvrfdoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
