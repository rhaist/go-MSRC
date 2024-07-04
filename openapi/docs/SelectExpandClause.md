# SelectExpandClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedItems** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AllSelected** | Pointer to **bool** |  | [optional] 

## Methods

### NewSelectExpandClause

`func NewSelectExpandClause() *SelectExpandClause`

NewSelectExpandClause instantiates a new SelectExpandClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectExpandClauseWithDefaults

`func NewSelectExpandClauseWithDefaults() *SelectExpandClause`

NewSelectExpandClauseWithDefaults instantiates a new SelectExpandClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedItems

`func (o *SelectExpandClause) GetSelectedItems() []map[string]interface{}`

GetSelectedItems returns the SelectedItems field if non-nil, zero value otherwise.

### GetSelectedItemsOk

`func (o *SelectExpandClause) GetSelectedItemsOk() (*[]map[string]interface{}, bool)`

GetSelectedItemsOk returns a tuple with the SelectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedItems

`func (o *SelectExpandClause) SetSelectedItems(v []map[string]interface{})`

SetSelectedItems sets SelectedItems field to given value.

### HasSelectedItems

`func (o *SelectExpandClause) HasSelectedItems() bool`

HasSelectedItems returns a boolean if a field has been set.

### SetSelectedItemsNil

`func (o *SelectExpandClause) SetSelectedItemsNil(b bool)`

 SetSelectedItemsNil sets the value for SelectedItems to be an explicit nil

### UnsetSelectedItems
`func (o *SelectExpandClause) UnsetSelectedItems()`

UnsetSelectedItems ensures that no value is present for SelectedItems, not even an explicit nil
### GetAllSelected

`func (o *SelectExpandClause) GetAllSelected() bool`

GetAllSelected returns the AllSelected field if non-nil, zero value otherwise.

### GetAllSelectedOk

`func (o *SelectExpandClause) GetAllSelectedOk() (*bool, bool)`

GetAllSelectedOk returns a tuple with the AllSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSelected

`func (o *SelectExpandClause) SetAllSelected(v bool)`

SetAllSelected sets AllSelected field to given value.

### HasAllSelected

`func (o *SelectExpandClause) HasAllSelected() bool`

HasAllSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


