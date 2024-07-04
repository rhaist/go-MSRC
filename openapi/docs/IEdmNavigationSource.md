# IEdmNavigationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NavigationPropertyBindings** | Pointer to [**[]IEdmNavigationPropertyBinding**](IEdmNavigationPropertyBinding.md) |  | [optional] [readonly] 
**Path** | Pointer to [**IEdmPathExpression**](IEdmPathExpression.md) |  | [optional] 
**Type** | Pointer to [**IEdmType**](IEdmType.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmNavigationSource

`func NewIEdmNavigationSource() *IEdmNavigationSource`

NewIEdmNavigationSource instantiates a new IEdmNavigationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmNavigationSourceWithDefaults

`func NewIEdmNavigationSourceWithDefaults() *IEdmNavigationSource`

NewIEdmNavigationSourceWithDefaults instantiates a new IEdmNavigationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNavigationPropertyBindings

`func (o *IEdmNavigationSource) GetNavigationPropertyBindings() []IEdmNavigationPropertyBinding`

GetNavigationPropertyBindings returns the NavigationPropertyBindings field if non-nil, zero value otherwise.

### GetNavigationPropertyBindingsOk

`func (o *IEdmNavigationSource) GetNavigationPropertyBindingsOk() (*[]IEdmNavigationPropertyBinding, bool)`

GetNavigationPropertyBindingsOk returns a tuple with the NavigationPropertyBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationPropertyBindings

`func (o *IEdmNavigationSource) SetNavigationPropertyBindings(v []IEdmNavigationPropertyBinding)`

SetNavigationPropertyBindings sets NavigationPropertyBindings field to given value.

### HasNavigationPropertyBindings

`func (o *IEdmNavigationSource) HasNavigationPropertyBindings() bool`

HasNavigationPropertyBindings returns a boolean if a field has been set.

### SetNavigationPropertyBindingsNil

`func (o *IEdmNavigationSource) SetNavigationPropertyBindingsNil(b bool)`

 SetNavigationPropertyBindingsNil sets the value for NavigationPropertyBindings to be an explicit nil

### UnsetNavigationPropertyBindings
`func (o *IEdmNavigationSource) UnsetNavigationPropertyBindings()`

UnsetNavigationPropertyBindings ensures that no value is present for NavigationPropertyBindings, not even an explicit nil
### GetPath

`func (o *IEdmNavigationSource) GetPath() IEdmPathExpression`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IEdmNavigationSource) GetPathOk() (*IEdmPathExpression, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IEdmNavigationSource) SetPath(v IEdmPathExpression)`

SetPath sets Path field to given value.

### HasPath

`func (o *IEdmNavigationSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *IEdmNavigationSource) GetType() IEdmType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IEdmNavigationSource) GetTypeOk() (*IEdmType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IEdmNavigationSource) SetType(v IEdmType)`

SetType sets Type field to given value.

### HasType

`func (o *IEdmNavigationSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IEdmNavigationSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmNavigationSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmNavigationSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmNavigationSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmNavigationSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmNavigationSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


