# UpdateODataQueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**HttpRequest**](HttpRequest.md) |  | [optional] 
**Context** | Pointer to [**ODataQueryContext**](ODataQueryContext.md) |  | [optional] 
**RawValues** | Pointer to [**ODataRawQueryOptions**](ODataRawQueryOptions.md) |  | [optional] 
**SelectExpand** | Pointer to [**SelectExpandQueryOption**](SelectExpandQueryOption.md) |  | [optional] 
**Apply** | Pointer to [**ApplyQueryOption**](ApplyQueryOption.md) |  | [optional] 
**Compute** | Pointer to [**ComputeQueryOption**](ComputeQueryOption.md) |  | [optional] 
**Filter** | Pointer to [**FilterQueryOption**](FilterQueryOption.md) |  | [optional] 
**Search** | Pointer to [**SearchQueryOption**](SearchQueryOption.md) |  | [optional] 
**OrderBy** | Pointer to [**OrderByQueryOption**](OrderByQueryOption.md) |  | [optional] 
**Skip** | Pointer to [**SkipQueryOption**](SkipQueryOption.md) |  | [optional] 
**SkipToken** | Pointer to [**SkipTokenQueryOption**](SkipTokenQueryOption.md) |  | [optional] 
**Top** | Pointer to [**TopQueryOption**](TopQueryOption.md) |  | [optional] 
**Count** | Pointer to [**CountQueryOption**](CountQueryOption.md) |  | [optional] 
**Validator** | Pointer to **map[string]interface{}** |  | [optional] 
**IfMatch** | Pointer to [**UpdateETag**](UpdateETag.md) |  | [optional] 
**IfNoneMatch** | Pointer to [**UpdateETag**](UpdateETag.md) |  | [optional] 

## Methods

### NewUpdateODataQueryOptions

`func NewUpdateODataQueryOptions() *UpdateODataQueryOptions`

NewUpdateODataQueryOptions instantiates a new UpdateODataQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateODataQueryOptionsWithDefaults

`func NewUpdateODataQueryOptionsWithDefaults() *UpdateODataQueryOptions`

NewUpdateODataQueryOptionsWithDefaults instantiates a new UpdateODataQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *UpdateODataQueryOptions) GetRequest() HttpRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *UpdateODataQueryOptions) GetRequestOk() (*HttpRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *UpdateODataQueryOptions) SetRequest(v HttpRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *UpdateODataQueryOptions) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetContext

`func (o *UpdateODataQueryOptions) GetContext() ODataQueryContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UpdateODataQueryOptions) GetContextOk() (*ODataQueryContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UpdateODataQueryOptions) SetContext(v ODataQueryContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *UpdateODataQueryOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRawValues

`func (o *UpdateODataQueryOptions) GetRawValues() ODataRawQueryOptions`

GetRawValues returns the RawValues field if non-nil, zero value otherwise.

### GetRawValuesOk

`func (o *UpdateODataQueryOptions) GetRawValuesOk() (*ODataRawQueryOptions, bool)`

GetRawValuesOk returns a tuple with the RawValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValues

`func (o *UpdateODataQueryOptions) SetRawValues(v ODataRawQueryOptions)`

SetRawValues sets RawValues field to given value.

### HasRawValues

`func (o *UpdateODataQueryOptions) HasRawValues() bool`

HasRawValues returns a boolean if a field has been set.

### GetSelectExpand

`func (o *UpdateODataQueryOptions) GetSelectExpand() SelectExpandQueryOption`

GetSelectExpand returns the SelectExpand field if non-nil, zero value otherwise.

### GetSelectExpandOk

`func (o *UpdateODataQueryOptions) GetSelectExpandOk() (*SelectExpandQueryOption, bool)`

GetSelectExpandOk returns a tuple with the SelectExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectExpand

`func (o *UpdateODataQueryOptions) SetSelectExpand(v SelectExpandQueryOption)`

SetSelectExpand sets SelectExpand field to given value.

### HasSelectExpand

`func (o *UpdateODataQueryOptions) HasSelectExpand() bool`

HasSelectExpand returns a boolean if a field has been set.

### GetApply

`func (o *UpdateODataQueryOptions) GetApply() ApplyQueryOption`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *UpdateODataQueryOptions) GetApplyOk() (*ApplyQueryOption, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *UpdateODataQueryOptions) SetApply(v ApplyQueryOption)`

SetApply sets Apply field to given value.

### HasApply

`func (o *UpdateODataQueryOptions) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetCompute

`func (o *UpdateODataQueryOptions) GetCompute() ComputeQueryOption`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *UpdateODataQueryOptions) GetComputeOk() (*ComputeQueryOption, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *UpdateODataQueryOptions) SetCompute(v ComputeQueryOption)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *UpdateODataQueryOptions) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateODataQueryOptions) GetFilter() FilterQueryOption`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateODataQueryOptions) GetFilterOk() (*FilterQueryOption, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateODataQueryOptions) SetFilter(v FilterQueryOption)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateODataQueryOptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearch

`func (o *UpdateODataQueryOptions) GetSearch() SearchQueryOption`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *UpdateODataQueryOptions) GetSearchOk() (*SearchQueryOption, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *UpdateODataQueryOptions) SetSearch(v SearchQueryOption)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *UpdateODataQueryOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetOrderBy

`func (o *UpdateODataQueryOptions) GetOrderBy() OrderByQueryOption`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *UpdateODataQueryOptions) GetOrderByOk() (*OrderByQueryOption, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *UpdateODataQueryOptions) SetOrderBy(v OrderByQueryOption)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *UpdateODataQueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSkip

`func (o *UpdateODataQueryOptions) GetSkip() SkipQueryOption`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *UpdateODataQueryOptions) GetSkipOk() (*SkipQueryOption, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *UpdateODataQueryOptions) SetSkip(v SkipQueryOption)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *UpdateODataQueryOptions) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSkipToken

`func (o *UpdateODataQueryOptions) GetSkipToken() SkipTokenQueryOption`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *UpdateODataQueryOptions) GetSkipTokenOk() (*SkipTokenQueryOption, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *UpdateODataQueryOptions) SetSkipToken(v SkipTokenQueryOption)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *UpdateODataQueryOptions) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### GetTop

`func (o *UpdateODataQueryOptions) GetTop() TopQueryOption`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *UpdateODataQueryOptions) GetTopOk() (*TopQueryOption, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *UpdateODataQueryOptions) SetTop(v TopQueryOption)`

SetTop sets Top field to given value.

### HasTop

`func (o *UpdateODataQueryOptions) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetCount

`func (o *UpdateODataQueryOptions) GetCount() CountQueryOption`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UpdateODataQueryOptions) GetCountOk() (*CountQueryOption, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UpdateODataQueryOptions) SetCount(v CountQueryOption)`

SetCount sets Count field to given value.

### HasCount

`func (o *UpdateODataQueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetValidator

`func (o *UpdateODataQueryOptions) GetValidator() map[string]interface{}`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *UpdateODataQueryOptions) GetValidatorOk() (*map[string]interface{}, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *UpdateODataQueryOptions) SetValidator(v map[string]interface{})`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *UpdateODataQueryOptions) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetIfMatch

`func (o *UpdateODataQueryOptions) GetIfMatch() UpdateETag`

GetIfMatch returns the IfMatch field if non-nil, zero value otherwise.

### GetIfMatchOk

`func (o *UpdateODataQueryOptions) GetIfMatchOk() (*UpdateETag, bool)`

GetIfMatchOk returns a tuple with the IfMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfMatch

`func (o *UpdateODataQueryOptions) SetIfMatch(v UpdateETag)`

SetIfMatch sets IfMatch field to given value.

### HasIfMatch

`func (o *UpdateODataQueryOptions) HasIfMatch() bool`

HasIfMatch returns a boolean if a field has been set.

### GetIfNoneMatch

`func (o *UpdateODataQueryOptions) GetIfNoneMatch() UpdateETag`

GetIfNoneMatch returns the IfNoneMatch field if non-nil, zero value otherwise.

### GetIfNoneMatchOk

`func (o *UpdateODataQueryOptions) GetIfNoneMatchOk() (*UpdateETag, bool)`

GetIfNoneMatchOk returns a tuple with the IfNoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfNoneMatch

`func (o *UpdateODataQueryOptions) SetIfNoneMatch(v UpdateETag)`

SetIfNoneMatch sets IfNoneMatch field to given value.

### HasIfNoneMatch

`func (o *UpdateODataQueryOptions) HasIfNoneMatch() bool`

HasIfNoneMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


