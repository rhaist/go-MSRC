# RangeVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**TypeReference** | Pointer to [**IEdmTypeReference**](IEdmTypeReference.md) |  | [optional] 
**Kind** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewRangeVariable

`func NewRangeVariable() *RangeVariable`

NewRangeVariable instantiates a new RangeVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeVariableWithDefaults

`func NewRangeVariableWithDefaults() *RangeVariable`

NewRangeVariableWithDefaults instantiates a new RangeVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RangeVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RangeVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RangeVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RangeVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RangeVariable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RangeVariable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTypeReference

`func (o *RangeVariable) GetTypeReference() IEdmTypeReference`

GetTypeReference returns the TypeReference field if non-nil, zero value otherwise.

### GetTypeReferenceOk

`func (o *RangeVariable) GetTypeReferenceOk() (*IEdmTypeReference, bool)`

GetTypeReferenceOk returns a tuple with the TypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeReference

`func (o *RangeVariable) SetTypeReference(v IEdmTypeReference)`

SetTypeReference sets TypeReference field to given value.

### HasTypeReference

`func (o *RangeVariable) HasTypeReference() bool`

HasTypeReference returns a boolean if a field has been set.

### GetKind

`func (o *RangeVariable) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RangeVariable) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RangeVariable) SetKind(v int32)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RangeVariable) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


