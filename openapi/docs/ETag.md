# ETag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsWellFormed** | Pointer to **bool** |  | [optional] 
**EntityType** | Pointer to [**Type**](Type.md) |  | [optional] 
**IsAny** | Pointer to **bool** |  | [optional] 
**IsIfNoneMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewETag

`func NewETag() *ETag`

NewETag instantiates a new ETag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETagWithDefaults

`func NewETagWithDefaults() *ETag`

NewETagWithDefaults instantiates a new ETag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsWellFormed

`func (o *ETag) GetIsWellFormed() bool`

GetIsWellFormed returns the IsWellFormed field if non-nil, zero value otherwise.

### GetIsWellFormedOk

`func (o *ETag) GetIsWellFormedOk() (*bool, bool)`

GetIsWellFormedOk returns a tuple with the IsWellFormed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWellFormed

`func (o *ETag) SetIsWellFormed(v bool)`

SetIsWellFormed sets IsWellFormed field to given value.

### HasIsWellFormed

`func (o *ETag) HasIsWellFormed() bool`

HasIsWellFormed returns a boolean if a field has been set.

### GetEntityType

`func (o *ETag) GetEntityType() Type`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ETag) GetEntityTypeOk() (*Type, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ETag) SetEntityType(v Type)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ETag) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIsAny

`func (o *ETag) GetIsAny() bool`

GetIsAny returns the IsAny field if non-nil, zero value otherwise.

### GetIsAnyOk

`func (o *ETag) GetIsAnyOk() (*bool, bool)`

GetIsAnyOk returns a tuple with the IsAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAny

`func (o *ETag) SetIsAny(v bool)`

SetIsAny sets IsAny field to given value.

### HasIsAny

`func (o *ETag) HasIsAny() bool`

HasIsAny returns a boolean if a field has been set.

### GetIsIfNoneMatch

`func (o *ETag) GetIsIfNoneMatch() bool`

GetIsIfNoneMatch returns the IsIfNoneMatch field if non-nil, zero value otherwise.

### GetIsIfNoneMatchOk

`func (o *ETag) GetIsIfNoneMatchOk() (*bool, bool)`

GetIsIfNoneMatchOk returns a tuple with the IsIfNoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIfNoneMatch

`func (o *ETag) SetIsIfNoneMatch(v bool)`

SetIsIfNoneMatch sets IsIfNoneMatch field to given value.

### HasIsIfNoneMatch

`func (o *ETag) HasIsIfNoneMatch() bool`

HasIsIfNoneMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


