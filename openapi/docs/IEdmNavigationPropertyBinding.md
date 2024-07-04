# IEdmNavigationPropertyBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NavigationProperty** | Pointer to [**IEdmNavigationProperty**](IEdmNavigationProperty.md) |  | [optional] 
**Target** | Pointer to [**IEdmNavigationSource**](IEdmNavigationSource.md) |  | [optional] 
**Path** | Pointer to [**IEdmPathExpression**](IEdmPathExpression.md) |  | [optional] 

## Methods

### NewIEdmNavigationPropertyBinding

`func NewIEdmNavigationPropertyBinding() *IEdmNavigationPropertyBinding`

NewIEdmNavigationPropertyBinding instantiates a new IEdmNavigationPropertyBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmNavigationPropertyBindingWithDefaults

`func NewIEdmNavigationPropertyBindingWithDefaults() *IEdmNavigationPropertyBinding`

NewIEdmNavigationPropertyBindingWithDefaults instantiates a new IEdmNavigationPropertyBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNavigationProperty

`func (o *IEdmNavigationPropertyBinding) GetNavigationProperty() IEdmNavigationProperty`

GetNavigationProperty returns the NavigationProperty field if non-nil, zero value otherwise.

### GetNavigationPropertyOk

`func (o *IEdmNavigationPropertyBinding) GetNavigationPropertyOk() (*IEdmNavigationProperty, bool)`

GetNavigationPropertyOk returns a tuple with the NavigationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationProperty

`func (o *IEdmNavigationPropertyBinding) SetNavigationProperty(v IEdmNavigationProperty)`

SetNavigationProperty sets NavigationProperty field to given value.

### HasNavigationProperty

`func (o *IEdmNavigationPropertyBinding) HasNavigationProperty() bool`

HasNavigationProperty returns a boolean if a field has been set.

### GetTarget

`func (o *IEdmNavigationPropertyBinding) GetTarget() IEdmNavigationSource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IEdmNavigationPropertyBinding) GetTargetOk() (*IEdmNavigationSource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IEdmNavigationPropertyBinding) SetTarget(v IEdmNavigationSource)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *IEdmNavigationPropertyBinding) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetPath

`func (o *IEdmNavigationPropertyBinding) GetPath() IEdmPathExpression`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IEdmNavigationPropertyBinding) GetPathOk() (*IEdmPathExpression, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IEdmNavigationPropertyBinding) SetPath(v IEdmPathExpression)`

SetPath sets Path field to given value.

### HasPath

`func (o *IEdmNavigationPropertyBinding) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


