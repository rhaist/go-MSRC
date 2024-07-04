# SingleValueNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeReference** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**Kind** | Pointer to [**QueryNodeKind**](QueryNodeKind.md) |  | [optional] 

## Methods

### NewSingleValueNode

`func NewSingleValueNode() *SingleValueNode`

NewSingleValueNode instantiates a new SingleValueNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleValueNodeWithDefaults

`func NewSingleValueNodeWithDefaults() *SingleValueNode`

NewSingleValueNodeWithDefaults instantiates a new SingleValueNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeReference

`func (o *SingleValueNode) GetTypeReference() IEdmTypeReference`

GetTypeReference returns the TypeReference field if non-nil, zero value otherwise.

### GetTypeReferenceOk

`func (o *SingleValueNode) GetTypeReferenceOk() (*IEdmTypeReference, bool)`

GetTypeReferenceOk returns a tuple with the TypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeReference

`func (o *SingleValueNode) SetTypeReference(v IEdmTypeReference)`

SetTypeReference sets TypeReference field to given value.

### HasTypeReference

`func (o *SingleValueNode) HasTypeReference() bool`

HasTypeReference returns a boolean if a field has been set.

### GetKind

`func (o *SingleValueNode) GetKind() QueryNodeKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SingleValueNode) GetKindOk() (*QueryNodeKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SingleValueNode) SetKind(v QueryNodeKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SingleValueNode) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


