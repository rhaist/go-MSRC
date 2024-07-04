# ClaimsPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]Claim**](Claim.md) |  | [optional] [readonly] 
**Identities** | Pointer to [**[]ClaimsIdentity**](ClaimsIdentity.md) |  | [optional] [readonly] 
**Identity** | Pointer to [**IIdentity**](IIdentity.md) |  | [optional] 

## Methods

### NewClaimsPrincipal

`func NewClaimsPrincipal() *ClaimsPrincipal`

NewClaimsPrincipal instantiates a new ClaimsPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimsPrincipalWithDefaults

`func NewClaimsPrincipalWithDefaults() *ClaimsPrincipal`

NewClaimsPrincipalWithDefaults instantiates a new ClaimsPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *ClaimsPrincipal) GetClaims() []Claim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ClaimsPrincipal) GetClaimsOk() (*[]Claim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ClaimsPrincipal) SetClaims(v []Claim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ClaimsPrincipal) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *ClaimsPrincipal) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *ClaimsPrincipal) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil
### GetIdentities

`func (o *ClaimsPrincipal) GetIdentities() []ClaimsIdentity`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ClaimsPrincipal) GetIdentitiesOk() (*[]ClaimsIdentity, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ClaimsPrincipal) SetIdentities(v []ClaimsIdentity)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *ClaimsPrincipal) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *ClaimsPrincipal) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *ClaimsPrincipal) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
### GetIdentity

`func (o *ClaimsPrincipal) GetIdentity() IIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ClaimsPrincipal) GetIdentityOk() (*IIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ClaimsPrincipal) SetIdentity(v IIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ClaimsPrincipal) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


