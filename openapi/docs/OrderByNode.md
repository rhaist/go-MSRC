# OrderByNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**OrderByDirection**](OrderByDirection.md) |  | [optional] 

## Methods

### NewOrderByNode

`func NewOrderByNode() *OrderByNode`

NewOrderByNode instantiates a new OrderByNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByNodeWithDefaults

`func NewOrderByNodeWithDefaults() *OrderByNode`

NewOrderByNodeWithDefaults instantiates a new OrderByNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *OrderByNode) GetDirection() OrderByDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OrderByNode) GetDirectionOk() (*OrderByDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OrderByNode) SetDirection(v OrderByDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OrderByNode) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


