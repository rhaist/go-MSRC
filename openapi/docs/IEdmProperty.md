# IEdmProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyKind** | Pointer to [**EdmPropertyKind**](EdmPropertyKind.md) |  | [optional] 
**Type** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**DeclaringType** | Pointer to [**IEdmStructuredType**](IEdmStructuredType.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmProperty

`func NewIEdmProperty() *IEdmProperty`

NewIEdmProperty instantiates a new IEdmProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmPropertyWithDefaults

`func NewIEdmPropertyWithDefaults() *IEdmProperty`

NewIEdmPropertyWithDefaults instantiates a new IEdmProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyKind

`func (o *IEdmProperty) GetPropertyKind() EdmPropertyKind`

GetPropertyKind returns the PropertyKind field if non-nil, zero value otherwise.

### GetPropertyKindOk

`func (o *IEdmProperty) GetPropertyKindOk() (*EdmPropertyKind, bool)`

GetPropertyKindOk returns a tuple with the PropertyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKind

`func (o *IEdmProperty) SetPropertyKind(v EdmPropertyKind)`

SetPropertyKind sets PropertyKind field to given value.

### HasPropertyKind

`func (o *IEdmProperty) HasPropertyKind() bool`

HasPropertyKind returns a boolean if a field has been set.

### GetType

`func (o *IEdmProperty) GetType() IEdmTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IEdmProperty) GetTypeOk() (*IEdmTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IEdmProperty) SetType(v IEdmTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *IEdmProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeclaringType

`func (o *IEdmProperty) GetDeclaringType() IEdmStructuredType`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *IEdmProperty) GetDeclaringTypeOk() (*IEdmStructuredType, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *IEdmProperty) SetDeclaringType(v IEdmStructuredType)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *IEdmProperty) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetName

`func (o *IEdmProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


