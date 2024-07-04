# IEdmNavigationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partner** | Pointer to [**IEdmNavigationProperty**](IEdmNavigationProperty.md) |  | [optional] 
**OnDelete** | Pointer to [**EdmOnDeleteAction**](EdmOnDeleteAction.md) |  | [optional] 
**ContainsTarget** | Pointer to **bool** |  | [optional] [readonly] 
**ReferentialConstraint** | Pointer to [**IEdmReferentialConstraint**](IEdmReferentialConstraint.md) |  | [optional] 
**PropertyKind** | Pointer to [**EdmPropertyKind**](EdmPropertyKind.md) |  | [optional] 
**Type** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**DeclaringType** | Pointer to [**IEdmStructuredType**](IEdmStructuredType.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmNavigationProperty

`func NewIEdmNavigationProperty() *IEdmNavigationProperty`

NewIEdmNavigationProperty instantiates a new IEdmNavigationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmNavigationPropertyWithDefaults

`func NewIEdmNavigationPropertyWithDefaults() *IEdmNavigationProperty`

NewIEdmNavigationPropertyWithDefaults instantiates a new IEdmNavigationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartner

`func (o *IEdmNavigationProperty) GetPartner() IEdmNavigationProperty`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *IEdmNavigationProperty) GetPartnerOk() (*IEdmNavigationProperty, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *IEdmNavigationProperty) SetPartner(v IEdmNavigationProperty)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *IEdmNavigationProperty) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetOnDelete

`func (o *IEdmNavigationProperty) GetOnDelete() EdmOnDeleteAction`

GetOnDelete returns the OnDelete field if non-nil, zero value otherwise.

### GetOnDeleteOk

`func (o *IEdmNavigationProperty) GetOnDeleteOk() (*EdmOnDeleteAction, bool)`

GetOnDeleteOk returns a tuple with the OnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDelete

`func (o *IEdmNavigationProperty) SetOnDelete(v EdmOnDeleteAction)`

SetOnDelete sets OnDelete field to given value.

### HasOnDelete

`func (o *IEdmNavigationProperty) HasOnDelete() bool`

HasOnDelete returns a boolean if a field has been set.

### GetContainsTarget

`func (o *IEdmNavigationProperty) GetContainsTarget() bool`

GetContainsTarget returns the ContainsTarget field if non-nil, zero value otherwise.

### GetContainsTargetOk

`func (o *IEdmNavigationProperty) GetContainsTargetOk() (*bool, bool)`

GetContainsTargetOk returns a tuple with the ContainsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsTarget

`func (o *IEdmNavigationProperty) SetContainsTarget(v bool)`

SetContainsTarget sets ContainsTarget field to given value.

### HasContainsTarget

`func (o *IEdmNavigationProperty) HasContainsTarget() bool`

HasContainsTarget returns a boolean if a field has been set.

### GetReferentialConstraint

`func (o *IEdmNavigationProperty) GetReferentialConstraint() IEdmReferentialConstraint`

GetReferentialConstraint returns the ReferentialConstraint field if non-nil, zero value otherwise.

### GetReferentialConstraintOk

`func (o *IEdmNavigationProperty) GetReferentialConstraintOk() (*IEdmReferentialConstraint, bool)`

GetReferentialConstraintOk returns a tuple with the ReferentialConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferentialConstraint

`func (o *IEdmNavigationProperty) SetReferentialConstraint(v IEdmReferentialConstraint)`

SetReferentialConstraint sets ReferentialConstraint field to given value.

### HasReferentialConstraint

`func (o *IEdmNavigationProperty) HasReferentialConstraint() bool`

HasReferentialConstraint returns a boolean if a field has been set.

### GetPropertyKind

`func (o *IEdmNavigationProperty) GetPropertyKind() EdmPropertyKind`

GetPropertyKind returns the PropertyKind field if non-nil, zero value otherwise.

### GetPropertyKindOk

`func (o *IEdmNavigationProperty) GetPropertyKindOk() (*EdmPropertyKind, bool)`

GetPropertyKindOk returns a tuple with the PropertyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKind

`func (o *IEdmNavigationProperty) SetPropertyKind(v EdmPropertyKind)`

SetPropertyKind sets PropertyKind field to given value.

### HasPropertyKind

`func (o *IEdmNavigationProperty) HasPropertyKind() bool`

HasPropertyKind returns a boolean if a field has been set.

### GetType

`func (o *IEdmNavigationProperty) GetType() IEdmTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IEdmNavigationProperty) GetTypeOk() (*IEdmTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IEdmNavigationProperty) SetType(v IEdmTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *IEdmNavigationProperty) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeclaringType

`func (o *IEdmNavigationProperty) GetDeclaringType() IEdmStructuredType`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *IEdmNavigationProperty) GetDeclaringTypeOk() (*IEdmStructuredType, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *IEdmNavigationProperty) SetDeclaringType(v IEdmStructuredType)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *IEdmNavigationProperty) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetName

`func (o *IEdmNavigationProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmNavigationProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmNavigationProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmNavigationProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmNavigationProperty) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmNavigationProperty) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


