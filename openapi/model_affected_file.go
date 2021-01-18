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

// AffectedFile struct for AffectedFile
type AffectedFile struct {
	ProductId        *string    `json:"ProductId,omitempty"`
	FileName         *string    `json:"FileName,omitempty"`
	FileVersion      *string    `json:"FileVersion,omitempty"`
	FilePath         *string    `json:"FilePath,omitempty"`
	FileLastModified *time.Time `json:"FileLastModified,omitempty"`
	FileArchitecture *string    `json:"FileArchitecture,omitempty"`
}

// NewAffectedFile instantiates a new AffectedFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffectedFile() *AffectedFile {
	this := AffectedFile{}
	return &this
}

// NewAffectedFileWithDefaults instantiates a new AffectedFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffectedFileWithDefaults() *AffectedFile {
	this := AffectedFile{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AffectedFile) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AffectedFile) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AffectedFile) SetProductId(v string) {
	o.ProductId = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AffectedFile) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AffectedFile) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AffectedFile) SetFileName(v string) {
	o.FileName = &v
}

// GetFileVersion returns the FileVersion field value if set, zero value otherwise.
func (o *AffectedFile) GetFileVersion() string {
	if o == nil || o.FileVersion == nil {
		var ret string
		return ret
	}
	return *o.FileVersion
}

// GetFileVersionOk returns a tuple with the FileVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetFileVersionOk() (*string, bool) {
	if o == nil || o.FileVersion == nil {
		return nil, false
	}
	return o.FileVersion, true
}

// HasFileVersion returns a boolean if a field has been set.
func (o *AffectedFile) HasFileVersion() bool {
	if o != nil && o.FileVersion != nil {
		return true
	}

	return false
}

// SetFileVersion gets a reference to the given string and assigns it to the FileVersion field.
func (o *AffectedFile) SetFileVersion(v string) {
	o.FileVersion = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *AffectedFile) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *AffectedFile) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *AffectedFile) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileLastModified returns the FileLastModified field value if set, zero value otherwise.
func (o *AffectedFile) GetFileLastModified() time.Time {
	if o == nil || o.FileLastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.FileLastModified
}

// GetFileLastModifiedOk returns a tuple with the FileLastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetFileLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.FileLastModified == nil {
		return nil, false
	}
	return o.FileLastModified, true
}

// HasFileLastModified returns a boolean if a field has been set.
func (o *AffectedFile) HasFileLastModified() bool {
	if o != nil && o.FileLastModified != nil {
		return true
	}

	return false
}

// SetFileLastModified gets a reference to the given time.Time and assigns it to the FileLastModified field.
func (o *AffectedFile) SetFileLastModified(v time.Time) {
	o.FileLastModified = &v
}

// GetFileArchitecture returns the FileArchitecture field value if set, zero value otherwise.
func (o *AffectedFile) GetFileArchitecture() string {
	if o == nil || o.FileArchitecture == nil {
		var ret string
		return ret
	}
	return *o.FileArchitecture
}

// GetFileArchitectureOk returns a tuple with the FileArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffectedFile) GetFileArchitectureOk() (*string, bool) {
	if o == nil || o.FileArchitecture == nil {
		return nil, false
	}
	return o.FileArchitecture, true
}

// HasFileArchitecture returns a boolean if a field has been set.
func (o *AffectedFile) HasFileArchitecture() bool {
	if o != nil && o.FileArchitecture != nil {
		return true
	}

	return false
}

// SetFileArchitecture gets a reference to the given string and assigns it to the FileArchitecture field.
func (o *AffectedFile) SetFileArchitecture(v string) {
	o.FileArchitecture = &v
}

func (o AffectedFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductId != nil {
		toSerialize["ProductId"] = o.ProductId
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.FileVersion != nil {
		toSerialize["FileVersion"] = o.FileVersion
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FileLastModified != nil {
		toSerialize["FileLastModified"] = o.FileLastModified
	}
	if o.FileArchitecture != nil {
		toSerialize["FileArchitecture"] = o.FileArchitecture
	}
	return json.Marshal(toSerialize)
}

type NullableAffectedFile struct {
	value *AffectedFile
	isSet bool
}

func (v NullableAffectedFile) Get() *AffectedFile {
	return v.value
}

func (v *NullableAffectedFile) Set(val *AffectedFile) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedFile) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedFile(val *AffectedFile) *NullableAffectedFile {
	return &NullableAffectedFile{value: val, isSet: true}
}

func (v NullableAffectedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}