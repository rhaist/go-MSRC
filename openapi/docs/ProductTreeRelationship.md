# ProductTreeRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullProductName** | Pointer to [**[]FullProductName**](FullProductName.md) |  | [optional] 
**ProductReference** | Pointer to **string** |  | [optional] 
**RelationType** | Pointer to **int32** |  | [optional] 
**RelatesToProductReference** | Pointer to **string** |  | [optional] 

## Methods

### NewProductTreeRelationship

`func NewProductTreeRelationship() *ProductTreeRelationship`

NewProductTreeRelationship instantiates a new ProductTreeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTreeRelationshipWithDefaults

`func NewProductTreeRelationshipWithDefaults() *ProductTreeRelationship`

NewProductTreeRelationshipWithDefaults instantiates a new ProductTreeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullProductName

`func (o *ProductTreeRelationship) GetFullProductName() []FullProductName`

GetFullProductName returns the FullProductName field if non-nil, zero value otherwise.

### GetFullProductNameOk

`func (o *ProductTreeRelationship) GetFullProductNameOk() (*[]FullProductName, bool)`

GetFullProductNameOk returns a tuple with the FullProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProductName

`func (o *ProductTreeRelationship) SetFullProductName(v []FullProductName)`

SetFullProductName sets FullProductName field to given value.

### HasFullProductName

`func (o *ProductTreeRelationship) HasFullProductName() bool`

HasFullProductName returns a boolean if a field has been set.

### GetProductReference

`func (o *ProductTreeRelationship) GetProductReference() string`

GetProductReference returns the ProductReference field if non-nil, zero value otherwise.

### GetProductReferenceOk

`func (o *ProductTreeRelationship) GetProductReferenceOk() (*string, bool)`

GetProductReferenceOk returns a tuple with the ProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductReference

`func (o *ProductTreeRelationship) SetProductReference(v string)`

SetProductReference sets ProductReference field to given value.

### HasProductReference

`func (o *ProductTreeRelationship) HasProductReference() bool`

HasProductReference returns a boolean if a field has been set.

### GetRelationType

`func (o *ProductTreeRelationship) GetRelationType() int32`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *ProductTreeRelationship) GetRelationTypeOk() (*int32, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *ProductTreeRelationship) SetRelationType(v int32)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *ProductTreeRelationship) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### GetRelatesToProductReference

`func (o *ProductTreeRelationship) GetRelatesToProductReference() string`

GetRelatesToProductReference returns the RelatesToProductReference field if non-nil, zero value otherwise.

### GetRelatesToProductReferenceOk

`func (o *ProductTreeRelationship) GetRelatesToProductReferenceOk() (*string, bool)`

GetRelatesToProductReferenceOk returns a tuple with the RelatesToProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatesToProductReference

`func (o *ProductTreeRelationship) SetRelatesToProductReference(v string)`

SetRelatesToProductReference sets RelatesToProductReference field to given value.

### HasRelatesToProductReference

`func (o *ProductTreeRelationship) HasRelatesToProductReference() bool`

HasRelatesToProductReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


