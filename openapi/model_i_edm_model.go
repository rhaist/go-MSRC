/*
MSRC CVRF API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IEdmModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IEdmModel{}

// IEdmModel struct for IEdmModel
type IEdmModel struct {
	SchemaElements                []IEdmSchemaElement        `json:"schemaElements,omitempty"`
	VocabularyAnnotations         []IEdmVocabularyAnnotation `json:"vocabularyAnnotations,omitempty"`
	ReferencedModels              []IEdmModel                `json:"referencedModels,omitempty"`
	DeclaredNamespaces            []string                   `json:"declaredNamespaces,omitempty"`
	DirectValueAnnotationsManager map[string]interface{}     `json:"directValueAnnotationsManager,omitempty"`
	EntityContainer               *IEdmEntityContainer       `json:"entityContainer,omitempty"`
}

// NewIEdmModel instantiates a new IEdmModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIEdmModel() *IEdmModel {
	this := IEdmModel{}
	return &this
}

// NewIEdmModelWithDefaults instantiates a new IEdmModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIEdmModelWithDefaults() *IEdmModel {
	this := IEdmModel{}
	return &this
}

// GetSchemaElements returns the SchemaElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmModel) GetSchemaElements() []IEdmSchemaElement {
	if o == nil {
		var ret []IEdmSchemaElement
		return ret
	}
	return o.SchemaElements
}

// GetSchemaElementsOk returns a tuple with the SchemaElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmModel) GetSchemaElementsOk() ([]IEdmSchemaElement, bool) {
	if o == nil || IsNil(o.SchemaElements) {
		return nil, false
	}
	return o.SchemaElements, true
}

// HasSchemaElements returns a boolean if a field has been set.
func (o *IEdmModel) HasSchemaElements() bool {
	if o != nil && !IsNil(o.SchemaElements) {
		return true
	}

	return false
}

// SetSchemaElements gets a reference to the given []IEdmSchemaElement and assigns it to the SchemaElements field.
func (o *IEdmModel) SetSchemaElements(v []IEdmSchemaElement) {
	o.SchemaElements = v
}

// GetVocabularyAnnotations returns the VocabularyAnnotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmModel) GetVocabularyAnnotations() []IEdmVocabularyAnnotation {
	if o == nil {
		var ret []IEdmVocabularyAnnotation
		return ret
	}
	return o.VocabularyAnnotations
}

// GetVocabularyAnnotationsOk returns a tuple with the VocabularyAnnotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmModel) GetVocabularyAnnotationsOk() ([]IEdmVocabularyAnnotation, bool) {
	if o == nil || IsNil(o.VocabularyAnnotations) {
		return nil, false
	}
	return o.VocabularyAnnotations, true
}

// HasVocabularyAnnotations returns a boolean if a field has been set.
func (o *IEdmModel) HasVocabularyAnnotations() bool {
	if o != nil && !IsNil(o.VocabularyAnnotations) {
		return true
	}

	return false
}

// SetVocabularyAnnotations gets a reference to the given []IEdmVocabularyAnnotation and assigns it to the VocabularyAnnotations field.
func (o *IEdmModel) SetVocabularyAnnotations(v []IEdmVocabularyAnnotation) {
	o.VocabularyAnnotations = v
}

// GetReferencedModels returns the ReferencedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmModel) GetReferencedModels() []IEdmModel {
	if o == nil {
		var ret []IEdmModel
		return ret
	}
	return o.ReferencedModels
}

// GetReferencedModelsOk returns a tuple with the ReferencedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmModel) GetReferencedModelsOk() ([]IEdmModel, bool) {
	if o == nil || IsNil(o.ReferencedModels) {
		return nil, false
	}
	return o.ReferencedModels, true
}

// HasReferencedModels returns a boolean if a field has been set.
func (o *IEdmModel) HasReferencedModels() bool {
	if o != nil && !IsNil(o.ReferencedModels) {
		return true
	}

	return false
}

// SetReferencedModels gets a reference to the given []IEdmModel and assigns it to the ReferencedModels field.
func (o *IEdmModel) SetReferencedModels(v []IEdmModel) {
	o.ReferencedModels = v
}

// GetDeclaredNamespaces returns the DeclaredNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IEdmModel) GetDeclaredNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeclaredNamespaces
}

// GetDeclaredNamespacesOk returns a tuple with the DeclaredNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IEdmModel) GetDeclaredNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeclaredNamespaces) {
		return nil, false
	}
	return o.DeclaredNamespaces, true
}

// HasDeclaredNamespaces returns a boolean if a field has been set.
func (o *IEdmModel) HasDeclaredNamespaces() bool {
	if o != nil && !IsNil(o.DeclaredNamespaces) {
		return true
	}

	return false
}

// SetDeclaredNamespaces gets a reference to the given []string and assigns it to the DeclaredNamespaces field.
func (o *IEdmModel) SetDeclaredNamespaces(v []string) {
	o.DeclaredNamespaces = v
}

// GetDirectValueAnnotationsManager returns the DirectValueAnnotationsManager field value if set, zero value otherwise.
func (o *IEdmModel) GetDirectValueAnnotationsManager() map[string]interface{} {
	if o == nil || IsNil(o.DirectValueAnnotationsManager) {
		var ret map[string]interface{}
		return ret
	}
	return o.DirectValueAnnotationsManager
}

// GetDirectValueAnnotationsManagerOk returns a tuple with the DirectValueAnnotationsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmModel) GetDirectValueAnnotationsManagerOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DirectValueAnnotationsManager) {
		return map[string]interface{}{}, false
	}
	return o.DirectValueAnnotationsManager, true
}

// HasDirectValueAnnotationsManager returns a boolean if a field has been set.
func (o *IEdmModel) HasDirectValueAnnotationsManager() bool {
	if o != nil && !IsNil(o.DirectValueAnnotationsManager) {
		return true
	}

	return false
}

// SetDirectValueAnnotationsManager gets a reference to the given map[string]interface{} and assigns it to the DirectValueAnnotationsManager field.
func (o *IEdmModel) SetDirectValueAnnotationsManager(v map[string]interface{}) {
	o.DirectValueAnnotationsManager = v
}

// GetEntityContainer returns the EntityContainer field value if set, zero value otherwise.
func (o *IEdmModel) GetEntityContainer() IEdmEntityContainer {
	if o == nil || IsNil(o.EntityContainer) {
		var ret IEdmEntityContainer
		return ret
	}
	return *o.EntityContainer
}

// GetEntityContainerOk returns a tuple with the EntityContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IEdmModel) GetEntityContainerOk() (*IEdmEntityContainer, bool) {
	if o == nil || IsNil(o.EntityContainer) {
		return nil, false
	}
	return o.EntityContainer, true
}

// HasEntityContainer returns a boolean if a field has been set.
func (o *IEdmModel) HasEntityContainer() bool {
	if o != nil && !IsNil(o.EntityContainer) {
		return true
	}

	return false
}

// SetEntityContainer gets a reference to the given IEdmEntityContainer and assigns it to the EntityContainer field.
func (o *IEdmModel) SetEntityContainer(v IEdmEntityContainer) {
	o.EntityContainer = &v
}

func (o IEdmModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IEdmModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaElements != nil {
		toSerialize["schemaElements"] = o.SchemaElements
	}
	if o.VocabularyAnnotations != nil {
		toSerialize["vocabularyAnnotations"] = o.VocabularyAnnotations
	}
	if o.ReferencedModels != nil {
		toSerialize["referencedModels"] = o.ReferencedModels
	}
	if o.DeclaredNamespaces != nil {
		toSerialize["declaredNamespaces"] = o.DeclaredNamespaces
	}
	if !IsNil(o.DirectValueAnnotationsManager) {
		toSerialize["directValueAnnotationsManager"] = o.DirectValueAnnotationsManager
	}
	if !IsNil(o.EntityContainer) {
		toSerialize["entityContainer"] = o.EntityContainer
	}
	return toSerialize, nil
}

type NullableIEdmModel struct {
	value *IEdmModel
	isSet bool
}

func (v NullableIEdmModel) Get() *IEdmModel {
	return v.value
}

func (v *NullableIEdmModel) Set(val *IEdmModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIEdmModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIEdmModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIEdmModel(val *IEdmModel) *NullableIEdmModel {
	return &NullableIEdmModel{value: val, isSet: true}
}

func (v NullableIEdmModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIEdmModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
