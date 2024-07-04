# ODataQueryContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultQueryConfigurations** | Pointer to [**DefaultQueryConfigurations**](DefaultQueryConfigurations.md) |  | [optional] 
**Model** | Pointer to [**IEdmModel**](IEdmModel.md) |  | [optional] 
**ElementType** | Pointer to [**IEdmType**](IEdmType.md) |  | [optional] 
**NavigationSource** | Pointer to [**IEdmNavigationSource**](IEdmNavigationSource.md) |  | [optional] 
**ElementClrType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Path** | Pointer to [**[]ODataPathSegment**](ODataPathSegment.md) |  | [optional] 
**RequestContainer** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewODataQueryContext

`func NewODataQueryContext() *ODataQueryContext`

NewODataQueryContext instantiates a new ODataQueryContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewODataQueryContextWithDefaults

`func NewODataQueryContextWithDefaults() *ODataQueryContext`

NewODataQueryContextWithDefaults instantiates a new ODataQueryContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultQueryConfigurations

`func (o *ODataQueryContext) GetDefaultQueryConfigurations() DefaultQueryConfigurations`

GetDefaultQueryConfigurations returns the DefaultQueryConfigurations field if non-nil, zero value otherwise.

### GetDefaultQueryConfigurationsOk

`func (o *ODataQueryContext) GetDefaultQueryConfigurationsOk() (*DefaultQueryConfigurations, bool)`

GetDefaultQueryConfigurationsOk returns a tuple with the DefaultQueryConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQueryConfigurations

`func (o *ODataQueryContext) SetDefaultQueryConfigurations(v DefaultQueryConfigurations)`

SetDefaultQueryConfigurations sets DefaultQueryConfigurations field to given value.

### HasDefaultQueryConfigurations

`func (o *ODataQueryContext) HasDefaultQueryConfigurations() bool`

HasDefaultQueryConfigurations returns a boolean if a field has been set.

### GetModel

`func (o *ODataQueryContext) GetModel() IEdmModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ODataQueryContext) GetModelOk() (*IEdmModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ODataQueryContext) SetModel(v IEdmModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *ODataQueryContext) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetElementType

`func (o *ODataQueryContext) GetElementType() IEdmType`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ODataQueryContext) GetElementTypeOk() (*IEdmType, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ODataQueryContext) SetElementType(v IEdmType)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *ODataQueryContext) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetNavigationSource

`func (o *ODataQueryContext) GetNavigationSource() IEdmNavigationSource`

GetNavigationSource returns the NavigationSource field if non-nil, zero value otherwise.

### GetNavigationSourceOk

`func (o *ODataQueryContext) GetNavigationSourceOk() (*IEdmNavigationSource, bool)`

GetNavigationSourceOk returns a tuple with the NavigationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationSource

`func (o *ODataQueryContext) SetNavigationSource(v IEdmNavigationSource)`

SetNavigationSource sets NavigationSource field to given value.

### HasNavigationSource

`func (o *ODataQueryContext) HasNavigationSource() bool`

HasNavigationSource returns a boolean if a field has been set.

### GetElementClrType

`func (o *ODataQueryContext) GetElementClrType() Type`

GetElementClrType returns the ElementClrType field if non-nil, zero value otherwise.

### GetElementClrTypeOk

`func (o *ODataQueryContext) GetElementClrTypeOk() (*Type, bool)`

GetElementClrTypeOk returns a tuple with the ElementClrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementClrType

`func (o *ODataQueryContext) SetElementClrType(v Type)`

SetElementClrType sets ElementClrType field to given value.

### HasElementClrType

`func (o *ODataQueryContext) HasElementClrType() bool`

HasElementClrType returns a boolean if a field has been set.

### GetPath

`func (o *ODataQueryContext) GetPath() []ODataPathSegment`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ODataQueryContext) GetPathOk() (*[]ODataPathSegment, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ODataQueryContext) SetPath(v []ODataPathSegment)`

SetPath sets Path field to given value.

### HasPath

`func (o *ODataQueryContext) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ODataQueryContext) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ODataQueryContext) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRequestContainer

`func (o *ODataQueryContext) GetRequestContainer() map[string]interface{}`

GetRequestContainer returns the RequestContainer field if non-nil, zero value otherwise.

### GetRequestContainerOk

`func (o *ODataQueryContext) GetRequestContainerOk() (*map[string]interface{}, bool)`

GetRequestContainerOk returns a tuple with the RequestContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContainer

`func (o *ODataQueryContext) SetRequestContainer(v map[string]interface{})`

SetRequestContainer sets RequestContainer field to given value.

### HasRequestContainer

`func (o *ODataQueryContext) HasRequestContainer() bool`

HasRequestContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


