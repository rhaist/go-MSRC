# IEdmEntityContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]IEdmEntityContainerElement**](IEdmEntityContainerElement.md) |  | [optional] [readonly] 
**SchemaElementKind** | Pointer to [**EdmSchemaElementKind**](EdmSchemaElementKind.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmEntityContainer

`func NewIEdmEntityContainer() *IEdmEntityContainer`

NewIEdmEntityContainer instantiates a new IEdmEntityContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmEntityContainerWithDefaults

`func NewIEdmEntityContainerWithDefaults() *IEdmEntityContainer`

NewIEdmEntityContainerWithDefaults instantiates a new IEdmEntityContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *IEdmEntityContainer) GetElements() []IEdmEntityContainerElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *IEdmEntityContainer) GetElementsOk() (*[]IEdmEntityContainerElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *IEdmEntityContainer) SetElements(v []IEdmEntityContainerElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *IEdmEntityContainer) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElementsNil

`func (o *IEdmEntityContainer) SetElementsNil(b bool)`

 SetElementsNil sets the value for Elements to be an explicit nil

### UnsetElements
`func (o *IEdmEntityContainer) UnsetElements()`

UnsetElements ensures that no value is present for Elements, not even an explicit nil
### GetSchemaElementKind

`func (o *IEdmEntityContainer) GetSchemaElementKind() EdmSchemaElementKind`

GetSchemaElementKind returns the SchemaElementKind field if non-nil, zero value otherwise.

### GetSchemaElementKindOk

`func (o *IEdmEntityContainer) GetSchemaElementKindOk() (*EdmSchemaElementKind, bool)`

GetSchemaElementKindOk returns a tuple with the SchemaElementKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaElementKind

`func (o *IEdmEntityContainer) SetSchemaElementKind(v EdmSchemaElementKind)`

SetSchemaElementKind sets SchemaElementKind field to given value.

### HasSchemaElementKind

`func (o *IEdmEntityContainer) HasSchemaElementKind() bool`

HasSchemaElementKind returns a boolean if a field has been set.

### GetNamespace

`func (o *IEdmEntityContainer) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IEdmEntityContainer) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IEdmEntityContainer) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IEdmEntityContainer) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *IEdmEntityContainer) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *IEdmEntityContainer) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetName

`func (o *IEdmEntityContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmEntityContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmEntityContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmEntityContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmEntityContainer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmEntityContainer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


