# ClaimsIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsAuthenticated** | Pointer to **bool** |  | [optional] [readonly] 
**Actor** | Pointer to [**ClaimsIdentity**](ClaimsIdentity.md) |  | [optional] 
**BootstrapContext** | Pointer to **interface{}** |  | [optional] 
**Claims** | Pointer to [**[]Claim**](Claim.md) |  | [optional] [readonly] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**NameClaimType** | Pointer to **NullableString** |  | [optional] [readonly] 
**RoleClaimType** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewClaimsIdentity

`func NewClaimsIdentity() *ClaimsIdentity`

NewClaimsIdentity instantiates a new ClaimsIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimsIdentityWithDefaults

`func NewClaimsIdentityWithDefaults() *ClaimsIdentity`

NewClaimsIdentityWithDefaults instantiates a new ClaimsIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *ClaimsIdentity) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ClaimsIdentity) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ClaimsIdentity) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ClaimsIdentity) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *ClaimsIdentity) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *ClaimsIdentity) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetIsAuthenticated

`func (o *ClaimsIdentity) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *ClaimsIdentity) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *ClaimsIdentity) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *ClaimsIdentity) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### GetActor

`func (o *ClaimsIdentity) GetActor() ClaimsIdentity`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ClaimsIdentity) GetActorOk() (*ClaimsIdentity, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ClaimsIdentity) SetActor(v ClaimsIdentity)`

SetActor sets Actor field to given value.

### HasActor

`func (o *ClaimsIdentity) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetBootstrapContext

`func (o *ClaimsIdentity) GetBootstrapContext() interface{}`

GetBootstrapContext returns the BootstrapContext field if non-nil, zero value otherwise.

### GetBootstrapContextOk

`func (o *ClaimsIdentity) GetBootstrapContextOk() (*interface{}, bool)`

GetBootstrapContextOk returns a tuple with the BootstrapContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapContext

`func (o *ClaimsIdentity) SetBootstrapContext(v interface{})`

SetBootstrapContext sets BootstrapContext field to given value.

### HasBootstrapContext

`func (o *ClaimsIdentity) HasBootstrapContext() bool`

HasBootstrapContext returns a boolean if a field has been set.

### SetBootstrapContextNil

`func (o *ClaimsIdentity) SetBootstrapContextNil(b bool)`

 SetBootstrapContextNil sets the value for BootstrapContext to be an explicit nil

### UnsetBootstrapContext
`func (o *ClaimsIdentity) UnsetBootstrapContext()`

UnsetBootstrapContext ensures that no value is present for BootstrapContext, not even an explicit nil
### GetClaims

`func (o *ClaimsIdentity) GetClaims() []Claim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ClaimsIdentity) GetClaimsOk() (*[]Claim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ClaimsIdentity) SetClaims(v []Claim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ClaimsIdentity) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *ClaimsIdentity) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *ClaimsIdentity) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetLabel

`func (o *ClaimsIdentity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ClaimsIdentity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ClaimsIdentity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ClaimsIdentity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ClaimsIdentity) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ClaimsIdentity) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetName

`func (o *ClaimsIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClaimsIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClaimsIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClaimsIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClaimsIdentity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClaimsIdentity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameClaimType

`func (o *ClaimsIdentity) GetNameClaimType() string`

GetNameClaimType returns the NameClaimType field if non-nil, zero value otherwise.

### GetNameClaimTypeOk

`func (o *ClaimsIdentity) GetNameClaimTypeOk() (*string, bool)`

GetNameClaimTypeOk returns a tuple with the NameClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameClaimType

`func (o *ClaimsIdentity) SetNameClaimType(v string)`

SetNameClaimType sets NameClaimType field to given value.

### HasNameClaimType

`func (o *ClaimsIdentity) HasNameClaimType() bool`

HasNameClaimType returns a boolean if a field has been set.

### SetNameClaimTypeNil

`func (o *ClaimsIdentity) SetNameClaimTypeNil(b bool)`

 SetNameClaimTypeNil sets the value for NameClaimType to be an explicit nil

### UnsetNameClaimType
`func (o *ClaimsIdentity) UnsetNameClaimType()`

UnsetNameClaimType ensures that no value is present for NameClaimType, not even an explicit nil
### GetRoleClaimType

`func (o *ClaimsIdentity) GetRoleClaimType() string`

GetRoleClaimType returns the RoleClaimType field if non-nil, zero value otherwise.

### GetRoleClaimTypeOk

`func (o *ClaimsIdentity) GetRoleClaimTypeOk() (*string, bool)`

GetRoleClaimTypeOk returns a tuple with the RoleClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleClaimType

`func (o *ClaimsIdentity) SetRoleClaimType(v string)`

SetRoleClaimType sets RoleClaimType field to given value.

### HasRoleClaimType

`func (o *ClaimsIdentity) HasRoleClaimType() bool`

HasRoleClaimType returns a boolean if a field has been set.

### SetRoleClaimTypeNil

`func (o *ClaimsIdentity) SetRoleClaimTypeNil(b bool)`

 SetRoleClaimTypeNil sets the value for RoleClaimType to be an explicit nil

### UnsetRoleClaimType
`func (o *ClaimsIdentity) UnsetRoleClaimType()`

UnsetRoleClaimType ensures that no value is present for RoleClaimType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


