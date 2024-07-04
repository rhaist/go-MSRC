# Claim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **NullableString** |  | [optional] [readonly] 
**OriginalIssuer** | Pointer to **NullableString** |  | [optional] [readonly] 
**Properties** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Subject** | Pointer to [**ClaimsIdentity**](ClaimsIdentity.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] [readonly] 
**Value** | Pointer to **NullableString** |  | [optional] [readonly] 
**ValueType** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewClaim

`func NewClaim() *Claim`

NewClaim instantiates a new Claim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimWithDefaults

`func NewClaimWithDefaults() *Claim`

NewClaimWithDefaults instantiates a new Claim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *Claim) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Claim) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Claim) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Claim) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *Claim) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *Claim) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetOriginalIssuer

`func (o *Claim) GetOriginalIssuer() string`

GetOriginalIssuer returns the OriginalIssuer field if non-nil, zero value otherwise.

### GetOriginalIssuerOk

`func (o *Claim) GetOriginalIssuerOk() (*string, bool)`

GetOriginalIssuerOk returns a tuple with the OriginalIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalIssuer

`func (o *Claim) SetOriginalIssuer(v string)`

SetOriginalIssuer sets OriginalIssuer field to given value.

### HasOriginalIssuer

`func (o *Claim) HasOriginalIssuer() bool`

HasOriginalIssuer returns a boolean if a field has been set.

### SetOriginalIssuerNil

`func (o *Claim) SetOriginalIssuerNil(b bool)`

 SetOriginalIssuerNil sets the value for OriginalIssuer to be an explicit nil

### UnsetOriginalIssuer
`func (o *Claim) UnsetOriginalIssuer()`

UnsetOriginalIssuer ensures that no value is present for OriginalIssuer, not even an explicit nil
### GetProperties

`func (o *Claim) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Claim) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Claim) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Claim) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Claim) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Claim) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSubject

`func (o *Claim) GetSubject() ClaimsIdentity`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Claim) GetSubjectOk() (*ClaimsIdentity, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Claim) SetSubject(v ClaimsIdentity)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Claim) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *Claim) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Claim) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Claim) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Claim) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Claim) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Claim) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *Claim) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Claim) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Claim) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Claim) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Claim) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Claim) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueType

`func (o *Claim) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Claim) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Claim) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Claim) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *Claim) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *Claim) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


