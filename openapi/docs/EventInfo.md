# EventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**DeclaringType** | Pointer to [**Type**](Type.md) |  | [optional] 
**ReflectedType** | Pointer to [**Type**](Type.md) |  | [optional] 
**Module** | Pointer to [**Module**](Module.md) |  | [optional] 
**CustomAttributes** | Pointer to [**[]CustomAttributeData**](CustomAttributeData.md) |  | [optional] [readonly] 
**IsCollectible** | Pointer to **bool** |  | [optional] [readonly] 
**MetadataToken** | Pointer to **int32** |  | [optional] [readonly] 
**MemberType** | Pointer to [**MemberTypes**](MemberTypes.md) |  | [optional] 
**Attributes** | Pointer to [**EventAttributes**](EventAttributes.md) |  | [optional] 
**IsSpecialName** | Pointer to **bool** |  | [optional] [readonly] 
**AddMethod** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 
**RemoveMethod** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 
**RaiseMethod** | Pointer to [**MethodInfo**](MethodInfo.md) |  | [optional] 
**IsMulticast** | Pointer to **bool** |  | [optional] [readonly] 
**EventHandlerType** | Pointer to [**Type**](Type.md) |  | [optional] 

## Methods

### NewEventInfo

`func NewEventInfo() *EventInfo`

NewEventInfo instantiates a new EventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInfoWithDefaults

`func NewEventInfoWithDefaults() *EventInfo`

NewEventInfoWithDefaults instantiates a new EventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EventInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EventInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeclaringType

`func (o *EventInfo) GetDeclaringType() Type`

GetDeclaringType returns the DeclaringType field if non-nil, zero value otherwise.

### GetDeclaringTypeOk

`func (o *EventInfo) GetDeclaringTypeOk() (*Type, bool)`

GetDeclaringTypeOk returns a tuple with the DeclaringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaringType

`func (o *EventInfo) SetDeclaringType(v Type)`

SetDeclaringType sets DeclaringType field to given value.

### HasDeclaringType

`func (o *EventInfo) HasDeclaringType() bool`

HasDeclaringType returns a boolean if a field has been set.

### GetReflectedType

`func (o *EventInfo) GetReflectedType() Type`

GetReflectedType returns the ReflectedType field if non-nil, zero value otherwise.

### GetReflectedTypeOk

`func (o *EventInfo) GetReflectedTypeOk() (*Type, bool)`

GetReflectedTypeOk returns a tuple with the ReflectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectedType

`func (o *EventInfo) SetReflectedType(v Type)`

SetReflectedType sets ReflectedType field to given value.

### HasReflectedType

`func (o *EventInfo) HasReflectedType() bool`

HasReflectedType returns a boolean if a field has been set.

### GetModule

`func (o *EventInfo) GetModule() Module`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EventInfo) GetModuleOk() (*Module, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EventInfo) SetModule(v Module)`

SetModule sets Module field to given value.

### HasModule

`func (o *EventInfo) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *EventInfo) GetCustomAttributes() []CustomAttributeData`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *EventInfo) GetCustomAttributesOk() (*[]CustomAttributeData, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *EventInfo) SetCustomAttributes(v []CustomAttributeData)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *EventInfo) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributesNil

`func (o *EventInfo) SetCustomAttributesNil(b bool)`

 SetCustomAttributesNil sets the value for CustomAttributes to be an explicit nil

### UnsetCustomAttributes
`func (o *EventInfo) UnsetCustomAttributes()`

UnsetCustomAttributes ensures that no value is present for CustomAttributes, not even an explicit nil
### GetIsCollectible

`func (o *EventInfo) GetIsCollectible() bool`

GetIsCollectible returns the IsCollectible field if non-nil, zero value otherwise.

### GetIsCollectibleOk

`func (o *EventInfo) GetIsCollectibleOk() (*bool, bool)`

GetIsCollectibleOk returns a tuple with the IsCollectible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollectible

`func (o *EventInfo) SetIsCollectible(v bool)`

SetIsCollectible sets IsCollectible field to given value.

### HasIsCollectible

`func (o *EventInfo) HasIsCollectible() bool`

HasIsCollectible returns a boolean if a field has been set.

### GetMetadataToken

`func (o *EventInfo) GetMetadataToken() int32`

GetMetadataToken returns the MetadataToken field if non-nil, zero value otherwise.

### GetMetadataTokenOk

`func (o *EventInfo) GetMetadataTokenOk() (*int32, bool)`

GetMetadataTokenOk returns a tuple with the MetadataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataToken

`func (o *EventInfo) SetMetadataToken(v int32)`

SetMetadataToken sets MetadataToken field to given value.

### HasMetadataToken

`func (o *EventInfo) HasMetadataToken() bool`

HasMetadataToken returns a boolean if a field has been set.

### GetMemberType

`func (o *EventInfo) GetMemberType() MemberTypes`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *EventInfo) GetMemberTypeOk() (*MemberTypes, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *EventInfo) SetMemberType(v MemberTypes)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *EventInfo) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetAttributes

`func (o *EventInfo) GetAttributes() EventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventInfo) GetAttributesOk() (*EventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventInfo) SetAttributes(v EventAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EventInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsSpecialName

`func (o *EventInfo) GetIsSpecialName() bool`

GetIsSpecialName returns the IsSpecialName field if non-nil, zero value otherwise.

### GetIsSpecialNameOk

`func (o *EventInfo) GetIsSpecialNameOk() (*bool, bool)`

GetIsSpecialNameOk returns a tuple with the IsSpecialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialName

`func (o *EventInfo) SetIsSpecialName(v bool)`

SetIsSpecialName sets IsSpecialName field to given value.

### HasIsSpecialName

`func (o *EventInfo) HasIsSpecialName() bool`

HasIsSpecialName returns a boolean if a field has been set.

### GetAddMethod

`func (o *EventInfo) GetAddMethod() MethodInfo`

GetAddMethod returns the AddMethod field if non-nil, zero value otherwise.

### GetAddMethodOk

`func (o *EventInfo) GetAddMethodOk() (*MethodInfo, bool)`

GetAddMethodOk returns a tuple with the AddMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMethod

`func (o *EventInfo) SetAddMethod(v MethodInfo)`

SetAddMethod sets AddMethod field to given value.

### HasAddMethod

`func (o *EventInfo) HasAddMethod() bool`

HasAddMethod returns a boolean if a field has been set.

### GetRemoveMethod

`func (o *EventInfo) GetRemoveMethod() MethodInfo`

GetRemoveMethod returns the RemoveMethod field if non-nil, zero value otherwise.

### GetRemoveMethodOk

`func (o *EventInfo) GetRemoveMethodOk() (*MethodInfo, bool)`

GetRemoveMethodOk returns a tuple with the RemoveMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveMethod

`func (o *EventInfo) SetRemoveMethod(v MethodInfo)`

SetRemoveMethod sets RemoveMethod field to given value.

### HasRemoveMethod

`func (o *EventInfo) HasRemoveMethod() bool`

HasRemoveMethod returns a boolean if a field has been set.

### GetRaiseMethod

`func (o *EventInfo) GetRaiseMethod() MethodInfo`

GetRaiseMethod returns the RaiseMethod field if non-nil, zero value otherwise.

### GetRaiseMethodOk

`func (o *EventInfo) GetRaiseMethodOk() (*MethodInfo, bool)`

GetRaiseMethodOk returns a tuple with the RaiseMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseMethod

`func (o *EventInfo) SetRaiseMethod(v MethodInfo)`

SetRaiseMethod sets RaiseMethod field to given value.

### HasRaiseMethod

`func (o *EventInfo) HasRaiseMethod() bool`

HasRaiseMethod returns a boolean if a field has been set.

### GetIsMulticast

`func (o *EventInfo) GetIsMulticast() bool`

GetIsMulticast returns the IsMulticast field if non-nil, zero value otherwise.

### GetIsMulticastOk

`func (o *EventInfo) GetIsMulticastOk() (*bool, bool)`

GetIsMulticastOk returns a tuple with the IsMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulticast

`func (o *EventInfo) SetIsMulticast(v bool)`

SetIsMulticast sets IsMulticast field to given value.

### HasIsMulticast

`func (o *EventInfo) HasIsMulticast() bool`

HasIsMulticast returns a boolean if a field has been set.

### GetEventHandlerType

`func (o *EventInfo) GetEventHandlerType() Type`

GetEventHandlerType returns the EventHandlerType field if non-nil, zero value otherwise.

### GetEventHandlerTypeOk

`func (o *EventInfo) GetEventHandlerTypeOk() (*Type, bool)`

GetEventHandlerTypeOk returns a tuple with the EventHandlerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerType

`func (o *EventInfo) SetEventHandlerType(v Type)`

SetEventHandlerType sets EventHandlerType field to given value.

### HasEventHandlerType

`func (o *EventInfo) HasEventHandlerType() bool`

HasEventHandlerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


