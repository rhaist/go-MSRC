# CustomAttributeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Constructor** | Pointer to [**ConstructorInfo**](ConstructorInfo.md) |  | [optional] 
**ConstructorArguments** | Pointer to [**[]CustomAttributeTypedArgument**](CustomAttributeTypedArgument.md) |  | [optional] [readonly] 
**NamedArguments** | Pointer to [**[]CustomAttributeNamedArgument**](CustomAttributeNamedArgument.md) |  | [optional] [readonly] 

## Methods

### NewCustomAttributeData

`func NewCustomAttributeData() *CustomAttributeData`

NewCustomAttributeData instantiates a new CustomAttributeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeDataWithDefaults

`func NewCustomAttributeDataWithDefaults() *CustomAttributeData`

NewCustomAttributeDataWithDefaults instantiates a new CustomAttributeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *CustomAttributeData) GetAttributeType() Type`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *CustomAttributeData) GetAttributeTypeOk() (*Type, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *CustomAttributeData) SetAttributeType(v Type)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *CustomAttributeData) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetConstructor

`func (o *CustomAttributeData) GetConstructor() ConstructorInfo`

GetConstructor returns the Constructor field if non-nil, zero value otherwise.

### GetConstructorOk

`func (o *CustomAttributeData) GetConstructorOk() (*ConstructorInfo, bool)`

GetConstructorOk returns a tuple with the Constructor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructor

`func (o *CustomAttributeData) SetConstructor(v ConstructorInfo)`

SetConstructor sets Constructor field to given value.

### HasConstructor

`func (o *CustomAttributeData) HasConstructor() bool`

HasConstructor returns a boolean if a field has been set.

### GetConstructorArguments

`func (o *CustomAttributeData) GetConstructorArguments() []CustomAttributeTypedArgument`

GetConstructorArguments returns the ConstructorArguments field if non-nil, zero value otherwise.

### GetConstructorArgumentsOk

`func (o *CustomAttributeData) GetConstructorArgumentsOk() (*[]CustomAttributeTypedArgument, bool)`

GetConstructorArgumentsOk returns a tuple with the ConstructorArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructorArguments

`func (o *CustomAttributeData) SetConstructorArguments(v []CustomAttributeTypedArgument)`

SetConstructorArguments sets ConstructorArguments field to given value.

### HasConstructorArguments

`func (o *CustomAttributeData) HasConstructorArguments() bool`

HasConstructorArguments returns a boolean if a field has been set.

### SetConstructorArgumentsNil

`func (o *CustomAttributeData) SetConstructorArgumentsNil(b bool)`

 SetConstructorArgumentsNil sets the value for ConstructorArguments to be an explicit nil

### UnsetConstructorArguments
`func (o *CustomAttributeData) UnsetConstructorArguments()`

UnsetConstructorArguments ensures that no value is present for ConstructorArguments, not even an explicit nil
### GetNamedArguments

`func (o *CustomAttributeData) GetNamedArguments() []CustomAttributeNamedArgument`

GetNamedArguments returns the NamedArguments field if non-nil, zero value otherwise.

### GetNamedArgumentsOk

`func (o *CustomAttributeData) GetNamedArgumentsOk() (*[]CustomAttributeNamedArgument, bool)`

GetNamedArgumentsOk returns a tuple with the NamedArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedArguments

`func (o *CustomAttributeData) SetNamedArguments(v []CustomAttributeNamedArgument)`

SetNamedArguments sets NamedArguments field to given value.

### HasNamedArguments

`func (o *CustomAttributeData) HasNamedArguments() bool`

HasNamedArguments returns a boolean if a field has been set.

### SetNamedArgumentsNil

`func (o *CustomAttributeData) SetNamedArgumentsNil(b bool)`

 SetNamedArgumentsNil sets the value for NamedArguments to be an explicit nil

### UnsetNamedArguments
`func (o *CustomAttributeData) UnsetNamedArguments()`

UnsetNamedArguments ensures that no value is present for NamedArguments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


