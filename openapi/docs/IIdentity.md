# IIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**AuthenticationType** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsAuthenticated** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewIIdentity

`func NewIIdentity() *IIdentity`

NewIIdentity instantiates a new IIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIIdentityWithDefaults

`func NewIIdentityWithDefaults() *IIdentity`

NewIIdentityWithDefaults instantiates a new IIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IIdentity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IIdentity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthenticationType

`func (o *IIdentity) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *IIdentity) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *IIdentity) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *IIdentity) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *IIdentity) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *IIdentity) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetIsAuthenticated

`func (o *IIdentity) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *IIdentity) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *IIdentity) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *IIdentity) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


