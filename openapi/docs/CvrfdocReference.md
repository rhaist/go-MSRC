# CvrfdocReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Description** | Pointer to [**CvrfdocReferenceDescription**](CvrfdocReferenceDescription.md) |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewCvrfdocReference

`func NewCvrfdocReference() *CvrfdocReference`

NewCvrfdocReference instantiates a new CvrfdocReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvrfdocReferenceWithDefaults

`func NewCvrfdocReferenceWithDefaults() *CvrfdocReference`

NewCvrfdocReferenceWithDefaults instantiates a new CvrfdocReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *CvrfdocReference) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *CvrfdocReference) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *CvrfdocReference) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *CvrfdocReference) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetDescription

`func (o *CvrfdocReference) GetDescription() CvrfdocReferenceDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CvrfdocReference) GetDescriptionOk() (*CvrfdocReferenceDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CvrfdocReference) SetDescription(v CvrfdocReferenceDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CvrfdocReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CvrfdocReference) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CvrfdocReference) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CvrfdocReference) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CvrfdocReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


