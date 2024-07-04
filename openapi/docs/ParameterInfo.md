# ParameterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ParameterAttributes**](ParameterAttributes.md) |  | [optional] 
**Member** | Pointer to [**MemberInfo**](MemberInfo.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**ParameterType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] [readonly] 
**IsIn** | Pointer to **bool** |  | [optional] [readonly] 
**IsLcid** | Pointer to **bool** |  | [optional] [readonly] 
**IsOptional** | Pointer to **bool** |  | [optional] [readonly] 
**IsOut** | Pointer to **bool** |  | [optional] [readonly] 
**IsRetval** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] [readonly] 
**RawDefaultValue** | Pointer to **interface{}** |  | [optional] [readonly] 
**HasDefaultValue** | Pointer to **bool** |  | [optional] [readonly] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**MetadataToken** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewParameterInfo

`func NewParameterInfo() *ParameterInfo`

NewParameterInfo instantiates a new ParameterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterInfoWithDefaults

`func NewParameterInfoWithDefaults() *ParameterInfo`

NewParameterInfoWithDefaults instantiates a new ParameterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ParameterInfo) GetAttributes() ParameterAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParameterInfo) GetAttributesOk() (*ParameterAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParameterInfo) SetAttributes(v ParameterAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ParameterInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMember

`func (o *ParameterInfo) GetMember() MemberInfo`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ParameterInfo) GetMemberOk() (*MemberInfo, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ParameterInfo) SetMember(v MemberInfo)`

SetMember sets Member field to given value.

### HasMember

`func (o *ParameterInfo) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *ParameterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParameterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ParameterInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ParameterInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParameterType

`func (o *ParameterInfo) GetParameterType() Type`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ParameterInfo) GetParameterTypeOk() (*Type, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ParameterInfo) SetParameterType(v Type)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *ParameterInfo) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### GetPosition

`func (o *ParameterInfo) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ParameterInfo) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ParameterInfo) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ParameterInfo) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsIn

`func (o *ParameterInfo) GetIsIn() bool`

GetIsIn returns the IsIn field if non-nil, zero value otherwise.

### GetIsInOk

`func (o *ParameterInfo) GetIsInOk() (*bool, bool)`

GetIsInOk returns a tuple with the IsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIn

`func (o *ParameterInfo) SetIsIn(v bool)`

SetIsIn sets IsIn field to given value.

### HasIsIn

`func (o *ParameterInfo) HasIsIn() bool`

HasIsIn returns a boolean if a field has been set.

### GetIsLcid

`func (o *ParameterInfo) GetIsLcid() bool`

GetIsLcid returns the IsLcid field if non-nil, zero value otherwise.

### GetIsLcidOk

`func (o *ParameterInfo) GetIsLcidOk() (*bool, bool)`

GetIsLcidOk returns a tuple with the IsLcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLcid

`func (o *ParameterInfo) SetIsLcid(v bool)`

SetIsLcid sets IsLcid field to given value.

### HasIsLcid

`func (o *ParameterInfo) HasIsLcid() bool`

HasIsLcid returns a boolean if a field has been set.

### GetIsOptional

`func (o *ParameterInfo) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *ParameterInfo) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *ParameterInfo) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *ParameterInfo) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetIsOut

`func (o *ParameterInfo) GetIsOut() bool`

GetIsOut returns the IsOut field if non-nil, zero value otherwise.

### GetIsOutOk

`func (o *ParameterInfo) GetIsOutOk() (*bool, bool)`

GetIsOutOk returns a tuple with the IsOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOut

`func (o *ParameterInfo) SetIsOut(v bool)`

SetIsOut sets IsOut field to given value.

### HasIsOut

`func (o *ParameterInfo) HasIsOut() bool`

HasIsOut returns a boolean if a field has been set.

### GetIsRetval

`func (o *ParameterInfo) GetIsRetval() bool`

GetIsRetval returns the IsRetval field if non-nil, zero value otherwise.

### GetIsRetvalOk

`func (o *ParameterInfo) GetIsRetvalOk() (*bool, bool)`

GetIsRetvalOk returns a tuple with the IsRetval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetval

`func (o *ParameterInfo) SetIsRetval(v bool)`

SetIsRetval sets IsRetval field to given value.

### HasIsRetval

`func (o *ParameterInfo) HasIsRetval() bool`

HasIsRetval returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ParameterInfo) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ParameterInfo) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ParameterInfo) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ParameterInfo) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ParameterInfo) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ParameterInfo) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRawDefaultValue

`func (o *ParameterInfo) GetRawDefaultValue() interface{}`

GetRawDefaultValue returns the RawDefaultValue field if non-nil, zero value otherwise.

### GetRawDefaultValueOk

`func (o *ParameterInfo) GetRawDefaultValueOk() (*interface{}, bool)`

GetRawDefaultValueOk returns a tuple with the RawDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDefaultValue

`func (o *ParameterInfo) SetRawDefaultValue(v interface{})`

SetRawDefaultValue sets RawDefaultValue field to given value.

### HasRawDefaultValue

`func (o *ParameterInfo) HasRawDefaultValue() bool`

HasRawDefaultValue returns a boolean if a field has been set.

### SetRawDefaultValueNil

`func (o *ParameterInfo) SetRawDefaultValueNil(b bool)`

 SetRawDefaultValueNil sets the value for RawDefaultValue to be an explicit nil

### UnsetRawDefaultValue
`func (o *ParameterInfo) UnsetRawDefaultValue()`

UnsetRawDefaultValue ensures that no value is present for RawDefaultValue, not even an explicit nil
### GetHasDefaultValue

`func (o *ParameterInfo) GetHasDefaultValue() bool`

GetHasDefaultValue returns the HasDefaultValue field if non-nil, zero value otherwise.

### GetHasDefaultValueOk

`func (o *ParameterInfo) GetHasDefaultValueOk() (*bool, bool)`

GetHasDefaultValueOk returns a tuple with the HasDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultValue

`func (o *ParameterInfo) SetHasDefaultValue(v bool)`

SetHasDefaultValue sets HasDefaultValue field to given value.

### HasHasDefaultValue

`func (o *ParameterInfo) HasHasDefaultValue() bool`

HasHasDefaultValue returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ParameterInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ParameterInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ParameterInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ParameterInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *ParameterInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *ParameterInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetMetadataToken

`func (o *ParameterInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *ParameterInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *ParameterInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *ParameterInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


