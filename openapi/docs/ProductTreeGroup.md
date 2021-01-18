# ProductTreeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**ProductTreeGroupDescription**](ProductTreeGroupDescription.md) |  | [optional] 
**ProductID** | Pointer to **[]string** |  | [optional] 
**GroupID** | Pointer to **string** |  | [optional] 

## Methods

### NewProductTreeGroup

`func NewProductTreeGroup() *ProductTreeGroup`

NewProductTreeGroup instantiates a new ProductTreeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTreeGroupWithDefaults

`func NewProductTreeGroupWithDefaults() *ProductTreeGroup`

NewProductTreeGroupWithDefaults instantiates a new ProductTreeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProductTreeGroup) GetDescription() ProductTreeGroupDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductTreeGroup) GetDescriptionOk() (*ProductTreeGroupDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductTreeGroup) SetDescription(v ProductTreeGroupDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductTreeGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductID

`func (o *ProductTreeGroup) GetProductID() []string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *ProductTreeGroup) GetProductIDOk() (*[]string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *ProductTreeGroup) SetProductID(v []string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *ProductTreeGroup) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetGroupID

`func (o *ProductTreeGroup) GetGroupID() string`

GetGroupID returns the GroupID field if non-nil, zero value otherwise.

### GetGroupIDOk

`func (o *ProductTreeGroup) GetGroupIDOk() (*string, bool)`

GetGroupIDOk returns a tuple with the GroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupID

`func (o *ProductTreeGroup) SetGroupID(v string)`

SetGroupID sets GroupID field to given value.

### HasGroupID

`func (o *ProductTreeGroup) HasGroupID() bool`

HasGroupID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


