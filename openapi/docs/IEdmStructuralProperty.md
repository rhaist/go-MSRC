# IEdmStructuralProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValueString** | Pointer to **NullableString** |  | [optional] [readonly] 
**PropertyKind** | Pointer to [**EdmPropertyKind**](EdmPropertyKind.md) |  | [optional] 
**Type** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**DeclaringType** | Pointer to [**IEdmStructuredType**](IEdmStructuredType.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmStructuralProperty

`func NewIEdmStructuralProperty() *IEdmStructuralProperty`

NewIEdmStructuralProperty instantiates a new IEdmStructuralProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmStructuralPropertyWithDefaults

`func NewIEdmStructuralPropertyWithDefaults() *IEdmStructuralProperty`

NewIEdmStructuralPropertyWithDefaults instantiates a new IEdmStructuralProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValueString

`func (o *IEdmStructuralProperty) GetDefaultValueString() string`

GetDefaultValueString returns the DefaultValueString field if non-nil, zero value otherwise.

### GetDefaultValueStringOk

`func (o *IEdmStructuralProperty) GetDefaultValueStringOk() (*string, bool)`

GetDefaultValueStringOk returns a tuple with the DefaultValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValueString

`func (o *IEdmStructuralProperty) SetDefaultValueString(v string)`

SetDefaultValueString sets DefaultValueString field to given value.

### HasDefaultValueString

`func (o *IEdmStructuralProperty) HasDefaultValueString() bool`

HasDefaultValueString returns a boolean if a field has been set.

### SetDefaultValueStringNil

`func (o *IEdmStructuralProperty) SetDefaultValueStringNil(b bool)`

 SetDefaultValueStringNil sets the value for DefaultValueString to be an explicit nil

### UnsetDefaultValueString
`func (o *IEdmStructuralProperty) UnsetDefaultValueString()`

UnsetDefaultValueString ensures that no value is present for DefaultValueString, not even an explicit nil
### GetPropertyKind

`func (o *IEdmStructuralProperty) GetPropertyKind() EdmPropertyKind`

GetPropertyKind returns the PropertyKind field if non-nil, zero value otherwise.

### GetPropertyKindOk

`func (o *IEdmStructuralProperty) GetPropertyKindOk() (*EdmPropertyKind, bool)`

GetPropertyKindOk returns a tuple with the PropertyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKind

`func (o *IEdmStructuralProperty) SetPropertyKind(v EdmPropertyKind)`

SetPropertyKind sets PropertyKind field to given value.

### HasPropertyKind

`func (o *IEdmStructuralProperty) HasPropertyKind() bool`

HasPropertyKind returns a boolean if a field has been set.

### GetType

`func (o *IEdmStructuralProperty) GetType() IEdmTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IEdmStructuralProperty) GetTypeOk() (*IEdmTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IEdmStructuralProperty) SetType(v IEdmTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *IEdmStructuralProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeclaringType

`func (o *IEdmStructuralProperty) GetDeclaringType() IEdmStructuredType`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *IEdmStructuralProperty) GetDeclaringTypeOk() (*IEdmStructuredType, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *IEdmStructuralProperty) SetDeclaringType(v IEdmStructuredType)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *IEdmStructuralProperty) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetName

`func (o *IEdmStructuralProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmStructuralProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmStructuralProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmStructuralProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmStructuralProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmStructuralProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


