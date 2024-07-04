# ISession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAvailable** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **NullableString** |  | [optional] [readonly] 
**Keys** | Pointer to **[]string** |  | [optional] [readonly] 

## Methods

### NewISession

`func NewISession() *ISession`

NewISession instantiates a new ISession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISessionWithDefaults

`func NewISessionWithDefaults() *ISession`

NewISessionWithDefaults instantiates a new ISession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAvailable

`func (o *ISession) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *ISession) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *ISession) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *ISession) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetId

`func (o *ISession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ISession) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ISession) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ISession) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKeys

`func (o *ISession) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ISession) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ISession) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ISession) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *ISession) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *ISession) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


