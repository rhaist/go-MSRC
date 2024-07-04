# CustomAttributeNamedArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberInfo** | Pointer to [**MemberInfo**](MemberInfo.md) |  | [optional] 
**TypedValue** | Pointer to [**CustomAttributeTypedArgument**](CustomAttributeTypedArgument.md) |  | [optional] 
**MemberName** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsField** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewCustomAttributeNamedArgument

`func NewCustomAttributeNamedArgument() *CustomAttributeNamedArgument`

NewCustomAttributeNamedArgument instantiates a new CustomAttributeNamedArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeNamedArgumentWithDefaults

`func NewCustomAttributeNamedArgumentWithDefaults() *CustomAttributeNamedArgument`

NewCustomAttributeNamedArgumentWithDefaults instantiates a new CustomAttributeNamedArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberInfo

`func (o *CustomAttributeNamedArgument) GetMemberInfo() MemberInfo`

GetMemberInfo returns the MemberInfo field if non-nil, zero value otherwise.

### GetMemberInfoOk

`func (o *CustomAttributeNamedArgument) GetMemberInfoOk() (*MemberInfo, bool)`

GetMemberInfoOk returns a tuple with the MemberInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberInfo

`func (o *CustomAttributeNamedArgument) SetMemberInfo(v MemberInfo)`

SetMemberInfo sets MemberInfo field to given value.

### HasMemberInfo

`func (o *CustomAttributeNamedArgument) HasMemberInfo() bool`

HasMemberInfo returns a boolean if a field has been set.

### GetTypedValue

`func (o *CustomAttributeNamedArgument) GetTypedValue() CustomAttributeTypedArgument`

GetTypedValue returns the TypedValue field if non-nil, zero value otherwise.

### GetTypedValueOk

`func (o *CustomAttributeNamedArgument) GetTypedValueOk() (*CustomAttributeTypedArgument, bool)`

GetTypedValueOk returns a tuple with the TypedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypedValue

`func (o *CustomAttributeNamedArgument) SetTypedValue(v CustomAttributeTypedArgument)`

SetTypedValue sets TypedValue field to given value.

### HasTypedValue

`func (o *CustomAttributeNamedArgument) HasTypedValue() bool`

HasTypedValue returns a boolean if a field has been set.

### GetMemberName

`func (o *CustomAttributeNamedArgument) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *CustomAttributeNamedArgument) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *CustomAttributeNamedArgument) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *CustomAttributeNamedArgument) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### SetMemberNameNil

`func (o *CustomAttributeNamedArgument) SetMemberNameNil(b bool)`

 SetMemberNameNil sets the value for MemberName to be an explicit nil

### UnsetMemberName
`func (o *CustomAttributeNamedArgument) UnsetMemberName()`

UnsetMemberName ensures that no value is present for MemberName, not even an explicit nil
### GetIsField

`func (o *CustomAttributeNamedArgument) GetIsField() bool`

GetIsField returns the IsField field if non-nil, zero value otherwise.

### GetIsFieldOk

`func (o *CustomAttributeNamedArgument) GetIsFieldOk() (*bool, bool)`

GetIsFieldOk returns a tuple with the IsField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsField

`func (o *CustomAttributeNamedArgument) SetIsField(v bool)`

SetIsField sets IsField field to given value.

### HasIsField

`func (o *CustomAttributeNamedArgument) HasIsField() bool`

HasIsField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


