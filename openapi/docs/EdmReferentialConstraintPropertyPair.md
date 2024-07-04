# EdmReferentialConstraintPropertyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependentProperty** | Pointer to [**IEdmStructuralProperty**](IEdmStructuralProperty.md) |  | [optional] 
**PrincipalProperty** | Pointer to [**IEdmStructuralProperty**](IEdmStructuralProperty.md) |  | [optional] 

## Methods

### NewEdmReferentialConstraintPropertyPair

`func NewEdmReferentialConstraintPropertyPair() *EdmReferentialConstraintPropertyPair`

NewEdmReferentialConstraintPropertyPair instantiates a new EdmReferentialConstraintPropertyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdmReferentialConstraintPropertyPairWithDefaults

`func NewEdmReferentialConstraintPropertyPairWithDefaults() *EdmReferentialConstraintPropertyPair`

NewEdmReferentialConstraintPropertyPairWithDefaults instantiates a new EdmReferentialConstraintPropertyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependentProperty

`func (o *EdmReferentialConstraintPropertyPair) GetDependentProperty() IEdmStructuralProperty`

GetDependentProperty returns the DependentProperty field if non-nil, zero value otherwise.

### GetDependentPropertyOk

`func (o *EdmReferentialConstraintPropertyPair) GetDependentPropertyOk() (*IEdmStructuralProperty, bool)`

GetDependentPropertyOk returns a tuple with the DependentProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentProperty

`func (o *EdmReferentialConstraintPropertyPair) SetDependentProperty(v IEdmStructuralProperty)`

SetDependentProperty sets DependentProperty field to given value.

### HasDependentProperty

`func (o *EdmReferentialConstraintPropertyPair) HasDependentProperty() bool`

HasDependentProperty returns a boolean if a field has been set.

### GetPrincipalProperty

`func (o *EdmReferentialConstraintPropertyPair) GetPrincipalProperty() IEdmStructuralProperty`

GetPrincipalProperty returns the PrincipalProperty field if non-nil, zero value otherwise.

### GetPrincipalPropertyOk

`func (o *EdmReferentialConstraintPropertyPair) GetPrincipalPropertyOk() (*IEdmStructuralProperty, bool)`

GetPrincipalPropertyOk returns a tuple with the PrincipalProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalProperty

`func (o *EdmReferentialConstraintPropertyPair) SetPrincipalProperty(v IEdmStructuralProperty)`

SetPrincipalProperty sets PrincipalProperty field to given value.

### HasPrincipalProperty

`func (o *EdmReferentialConstraintPropertyPair) HasPrincipalProperty() bool`

HasPrincipalProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


