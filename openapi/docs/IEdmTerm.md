# IEdmTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**AppliesTo** | Pointer to **NullableString** |  | [optional] [readonly] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] [readonly] 
**SchemaElementKind** | Pointer to [**EdmSchemaElementKind**](EdmSchemaElementKind.md) |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmTerm

`func NewIEdmTerm() *IEdmTerm`

NewIEdmTerm instantiates a new IEdmTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmTermWithDefaults

`func NewIEdmTermWithDefaults() *IEdmTerm`

NewIEdmTermWithDefaults instantiates a new IEdmTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IEdmTerm) GetType() IEdmTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IEdmTerm) GetTypeOk() (*IEdmTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IEdmTerm) SetType(v IEdmTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *IEdmTerm) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAppliesTo

`func (o *IEdmTerm) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *IEdmTerm) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *IEdmTerm) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *IEdmTerm) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *IEdmTerm) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *IEdmTerm) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetDefaultValue

`func (o *IEdmTerm) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *IEdmTerm) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *IEdmTerm) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *IEdmTerm) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *IEdmTerm) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *IEdmTerm) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetSchemaElementKind

`func (o *IEdmTerm) GetSchemaElementKind() EdmSchemaElementKind`

GetSchemaElementKind returns the SchemaElementKind field if non-nil, zero value otherwise.

### GetSchemaElementKindOk

`func (o *IEdmTerm) GetSchemaElementKindOk() (*EdmSchemaElementKind, bool)`

GetSchemaElementKindOk returns a tuple with the SchemaElementKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaElementKind

`func (o *IEdmTerm) SetSchemaElementKind(v EdmSchemaElementKind)`

SetSchemaElementKind sets SchemaElementKind field to given value.

### HasSchemaElementKind

`func (o *IEdmTerm) HasSchemaElementKind() bool`

HasSchemaElementKind returns a boolean if a field has been set.

### GetNamespace

`func (o *IEdmTerm) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IEdmTerm) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IEdmTerm) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IEdmTerm) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *IEdmTerm) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *IEdmTerm) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetName

`func (o *IEdmTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmTerm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmTerm) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmTerm) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmTerm) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


