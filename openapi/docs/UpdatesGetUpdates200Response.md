# UpdatesGetUpdates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | OData metadata for the request | 
**Value** | [**[]Update**](Update.md) | array of security update summaries | 

## Methods

### NewUpdatesGetUpdates200Response

`func NewUpdatesGetUpdates200Response(odataContext string, value []Update, ) *UpdatesGetUpdates200Response`

NewUpdatesGetUpdates200Response instantiates a new UpdatesGetUpdates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatesGetUpdates200ResponseWithDefaults

`func NewUpdatesGetUpdates200ResponseWithDefaults() *UpdatesGetUpdates200Response`

NewUpdatesGetUpdates200ResponseWithDefaults instantiates a new UpdatesGetUpdates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOdataContext

`func (o *UpdatesGetUpdates200Response) GetOdataContext() string`

GetOdataContext returns the OdataContext field if non-nil, zero value otherwise.

### GetOdataContextOk

`func (o *UpdatesGetUpdates200Response) GetOdataContextOk() (*string, bool)`

GetOdataContextOk returns a tuple with the OdataContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataContext

`func (o *UpdatesGetUpdates200Response) SetOdataContext(v string)`

SetOdataContext sets OdataContext field to given value.


### GetValue

`func (o *UpdatesGetUpdates200Response) GetValue() []Update`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatesGetUpdates200Response) GetValueOk() (*[]Update, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatesGetUpdates200Response) SetValue(v []Update)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


