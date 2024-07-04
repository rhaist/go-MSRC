# IEdmModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaElements** | Pointer to [**[]IEdmSchemaElement**](IEdmSchemaElement.md) |  | [optional] [readonly] 
**VocabularyAnnotations** | Pointer to [**[]IEdmVocabularyAnnotation**](IEdmVocabularyAnnotation.md) |  | [optional] [readonly] 
**ReferencedModels** | Pointer to [**[]IEdmModel**](IEdmModel.md) |  | [optional] [readonly] 
**DeclaredNamespaces** | Pointer to **[]string** |  | [optional] [readonly] 
**DirectValueAnnotationsManager** | Pointer to **map[string]interface{}** |  | [optional] 
**EntityContainer** | Pointer to [**IEdmEntityContainer**](IEdmEntityContainer.md) |  | [optional] 

## Methods

### NewIEdmModel

`func NewIEdmModel() *IEdmModel`

NewIEdmModel instantiates a new IEdmModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmModelWithDefaults

`func NewIEdmModelWithDefaults() *IEdmModel`

NewIEdmModelWithDefaults instantiates a new IEdmModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaElements

`func (o *IEdmModel) GetSchemaElements() []IEdmSchemaElement`

GetSchemaElements returns the SchemaElements field if non-nil, zero value otherwise.

### GetSchemaElementsOk

`func (o *IEdmModel) GetSchemaElementsOk() (*[]IEdmSchemaElement, bool)`

GetSchemaElementsOk returns a tuple with the SchemaElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaElements

`func (o *IEdmModel) SetSchemaElements(v []IEdmSchemaElement)`

SetSchemaElements sets SchemaElements field to given value.

### HasSchemaElements

`func (o *IEdmModel) HasSchemaElements() bool`

HasSchemaElements returns a boolean if a field has been set.

### SetSchemaElementsNil

`func (o *IEdmModel) SetSchemaElementsNil(b bool)`

 SetSchemaElementsNil sets the value for SchemaElements to be an explicit nil

### UnsetSchemaElements
`func (o *IEdmModel) UnsetSchemaElements()`

UnsetSchemaElements ensures that no value is present for SchemaElements, not even an explicit nil
### GetVocabularyAnnotations

`func (o *IEdmModel) GetVocabularyAnnotations() []IEdmVocabularyAnnotation`

GetVocabularyAnnotations returns the VocabularyAnnotations field if non-nil, zero value otherwise.

### GetVocabularyAnnotationsOk

`func (o *IEdmModel) GetVocabularyAnnotationsOk() (*[]IEdmVocabularyAnnotation, bool)`

GetVocabularyAnnotationsOk returns a tuple with the VocabularyAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVocabularyAnnotations

`func (o *IEdmModel) SetVocabularyAnnotations(v []IEdmVocabularyAnnotation)`

SetVocabularyAnnotations sets VocabularyAnnotations field to given value.

### HasVocabularyAnnotations

`func (o *IEdmModel) HasVocabularyAnnotations() bool`

HasVocabularyAnnotations returns a boolean if a field has been set.

### SetVocabularyAnnotationsNil

`func (o *IEdmModel) SetVocabularyAnnotationsNil(b bool)`

 SetVocabularyAnnotationsNil sets the value for VocabularyAnnotations to be an explicit nil

### UnsetVocabularyAnnotations
`func (o *IEdmModel) UnsetVocabularyAnnotations()`

UnsetVocabularyAnnotations ensures that no value is present for VocabularyAnnotations, not even an explicit nil
### GetReferencedModels

`func (o *IEdmModel) GetReferencedModels() []IEdmModel`

GetReferencedModels returns the ReferencedModels field if non-nil, zero value otherwise.

### GetReferencedModelsOk

`func (o *IEdmModel) GetReferencedModelsOk() (*[]IEdmModel, bool)`

GetReferencedModelsOk returns a tuple with the ReferencedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedModels

`func (o *IEdmModel) SetReferencedModels(v []IEdmModel)`

SetReferencedModels sets ReferencedModels field to given value.

### HasReferencedModels

`func (o *IEdmModel) HasReferencedModels() bool`

HasReferencedModels returns a boolean if a field has been set.

### SetReferencedModelsNil

`func (o *IEdmModel) SetReferencedModelsNil(b bool)`

 SetReferencedModelsNil sets the value for ReferencedModels to be an explicit nil

### UnsetReferencedModels
`func (o *IEdmModel) UnsetReferencedModels()`

UnsetReferencedModels ensures that no value is present for ReferencedModels, not even an explicit nil
### GetDeclaredNamespaces

`func (o *IEdmModel) GetDeclaredNamespaces() []string`

GetDeclaredNamespaces returns the DeclaredNamespaces field if non-nil, zero value otherwise.

### GetDeclaredNamespacesOk

`func (o *IEdmModel) GetDeclaredNamespacesOk() (*[]string, bool)`

GetDeclaredNamespacesOk returns a tuple with the DeclaredNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredNamespaces

`func (o *IEdmModel) SetDeclaredNamespaces(v []string)`

SetDeclaredNamespaces sets DeclaredNamespaces field to given value.

### HasDeclaredNamespaces

`func (o *IEdmModel) HasDeclaredNamespaces() bool`

HasDeclaredNamespaces returns a boolean if a field has been set.

### SetDeclaredNamespacesNil

`func (o *IEdmModel) SetDeclaredNamespacesNil(b bool)`

 SetDeclaredNamespacesNil sets the value for DeclaredNamespaces to be an explicit nil

### UnsetDeclaredNamespaces
`func (o *IEdmModel) UnsetDeclaredNamespaces()`

UnsetDeclaredNamespaces ensures that no value is present for DeclaredNamespaces, not even an explicit nil
### GetDirectValueAnnotationsManager

`func (o *IEdmModel) GetDirectValueAnnotationsManager() map[string]interface{}`

GetDirectValueAnnotationsManager returns the DirectValueAnnotationsManager field if non-nil, zero value otherwise.

### GetDirectValueAnnotationsManagerOk

`func (o *IEdmModel) GetDirectValueAnnotationsManagerOk() (*map[string]interface{}, bool)`

GetDirectValueAnnotationsManagerOk returns a tuple with the DirectValueAnnotationsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectValueAnnotationsManager

`func (o *IEdmModel) SetDirectValueAnnotationsManager(v map[string]interface{})`

SetDirectValueAnnotationsManager sets DirectValueAnnotationsManager field to given value.

### HasDirectValueAnnotationsManager

`func (o *IEdmModel) HasDirectValueAnnotationsManager() bool`

HasDirectValueAnnotationsManager returns a boolean if a field has been set.

### GetEntityContainer

`func (o *IEdmModel) GetEntityContainer() IEdmEntityContainer`

GetEntityContainer returns the EntityContainer field if non-nil, zero value otherwise.

### GetEntityContainerOk

`func (o *IEdmModel) GetEntityContainerOk() (*IEdmEntityContainer, bool)`

GetEntityContainerOk returns a tuple with the EntityContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityContainer

`func (o *IEdmModel) SetEntityContainer(v IEdmEntityContainer)`

SetEntityContainer sets EntityContainer field to given value.

### HasEntityContainer

`func (o *IEdmModel) HasEntityContainer() bool`

HasEntityContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


