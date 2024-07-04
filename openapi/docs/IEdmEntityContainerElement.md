# IEdmEntityContainerElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerElementKind** | Pointer to [**EdmContainerElementKind**](EdmContainerElementKind.md) |  | [optional] 
**Container** | Pointer to [**IEdmEntityContainer**](IEdmEntityContainer.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewIEdmEntityContainerElement

`func NewIEdmEntityContainerElement() *IEdmEntityContainerElement`

NewIEdmEntityContainerElement instantiates a new IEdmEntityContainerElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmEntityContainerElementWithDefaults

`func NewIEdmEntityContainerElementWithDefaults() *IEdmEntityContainerElement`

NewIEdmEntityContainerElementWithDefaults instantiates a new IEdmEntityContainerElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerElementKind

`func (o *IEdmEntityContainerElement) GetContainerElementKind() EdmContainerElementKind`

GetContainerElementKind returns the ContainerElementKind field if non-nil, zero value otherwise.

### GetContainerElementKindOk

`func (o *IEdmEntityContainerElement) GetContainerElementKindOk() (*EdmContainerElementKind, bool)`

GetContainerElementKindOk returns a tuple with the ContainerElementKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerElementKind

`func (o *IEdmEntityContainerElement) SetContainerElementKind(v EdmContainerElementKind)`

SetContainerElementKind sets ContainerElementKind field to given value.

### HasContainerElementKind

`func (o *IEdmEntityContainerElement) HasContainerElementKind() bool`

HasContainerElementKind returns a boolean if a field has been set.

### GetContainer

`func (o *IEdmEntityContainerElement) GetContainer() IEdmEntityContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *IEdmEntityContainerElement) GetContainerOk() (*IEdmEntityContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *IEdmEntityContainerElement) SetContainer(v IEdmEntityContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *IEdmEntityContainerElement) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetName

`func (o *IEdmEntityContainerElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IEdmEntityContainerElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IEdmEntityContainerElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IEdmEntityContainerElement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IEdmEntityContainerElement) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IEdmEntityContainerElement) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


