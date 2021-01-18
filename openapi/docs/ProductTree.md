# ProductTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**[]BranchType**](BranchType.md) |  | [optional] 
**FullProductName** | Pointer to [**[]FullProductName**](FullProductName.md) |  | [optional] 
**Relationship** | Pointer to [**[]ProductTreeRelationship**](ProductTreeRelationship.md) |  | [optional] 
**ProductGroups** | Pointer to [**[]ProductTreeGroup**](ProductTreeGroup.md) |  | [optional] 

## Methods

### NewProductTree

`func NewProductTree() *ProductTree`

NewProductTree instantiates a new ProductTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTreeWithDefaults

`func NewProductTreeWithDefaults() *ProductTree`

NewProductTreeWithDefaults instantiates a new ProductTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ProductTree) GetBranch() []BranchType`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ProductTree) GetBranchOk() (*[]BranchType, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ProductTree) SetBranch(v []BranchType)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ProductTree) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFullProductName

`func (o *ProductTree) GetFullProductName() []FullProductName`

GetFullProductName returns the FullProductName field if non-nil, zero value otherwise.

### GetFullProductNameOk

`func (o *ProductTree) GetFullProductNameOk() (*[]FullProductName, bool)`

GetFullProductNameOk returns a tuple with the FullProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProductName

`func (o *ProductTree) SetFullProductName(v []FullProductName)`

SetFullProductName sets FullProductName field to given value.

### HasFullProductName

`func (o *ProductTree) HasFullProductName() bool`

HasFullProductName returns a boolean if a field has been set.

### GetRelationship

`func (o *ProductTree) GetRelationship() []ProductTreeRelationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *ProductTree) GetRelationshipOk() (*[]ProductTreeRelationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *ProductTree) SetRelationship(v []ProductTreeRelationship)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *ProductTree) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetProductGroups

`func (o *ProductTree) GetProductGroups() []ProductTreeGroup`

GetProductGroups returns the ProductGroups field if non-nil, zero value otherwise.

### GetProductGroupsOk

`func (o *ProductTree) GetProductGroupsOk() (*[]ProductTreeGroup, bool)`

GetProductGroupsOk returns a tuple with the ProductGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroups

`func (o *ProductTree) SetProductGroups(v []ProductTreeGroup)`

SetProductGroups sets ProductGroups field to given value.

### HasProductGroups

`func (o *ProductTree) HasProductGroups() bool`

HasProductGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


