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

// BranchType struct for BranchType
type BranchType struct {
	Items []map[string]interface{} `json:"Items,omitempty"`
	Type  *int32                   `json:"Type,omitempty"`
	Name  *string                  `json:"Name,omitempty"`
}

// NewBranchType instantiates a new BranchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchType() *BranchType {
	this := BranchType{}
	return &this
}

// NewBranchTypeWithDefaults instantiates a new BranchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchTypeWithDefaults() *BranchType {
	this := BranchType{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BranchType) GetItems() []map[string]interface{} {
	if o == nil || o.Items == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchType) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BranchType) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *BranchType) SetItems(v []map[string]interface{}) {
	o.Items = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BranchType) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchType) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BranchType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *BranchType) SetType(v int32) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BranchType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BranchType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BranchType) SetName(v string) {
	o.Name = &v
}

func (o BranchType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBranchType struct {
	value *BranchType
	isSet bool
}

func (v NullableBranchType) Get() *BranchType {
	return v.value
}

func (v *NullableBranchType) Set(val *BranchType) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchType) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchType(val *BranchType) *NullableBranchType {
	return &NullableBranchType{value: val, isSet: true}
}

func (v NullableBranchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
