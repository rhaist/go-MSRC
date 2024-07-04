# ODataRawQueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **NullableString** |  | [optional] [readonly] 
**Apply** | Pointer to **NullableString** |  | [optional] [readonly] 
**Compute** | Pointer to **NullableString** |  | [optional] [readonly] 
**Search** | Pointer to **NullableString** |  | [optional] [readonly] 
**OrderBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**Top** | Pointer to **NullableString** |  | [optional] [readonly] 
**Skip** | Pointer to **NullableString** |  | [optional] [readonly] 
**Select** | Pointer to **NullableString** |  | [optional] [readonly] 
**Expand** | Pointer to **NullableString** |  | [optional] [readonly] 
**Count** | Pointer to **NullableString** |  | [optional] [readonly] 
**Format** | Pointer to **NullableString** |  | [optional] [readonly] 
**SkipToken** | Pointer to **NullableString** |  | [optional] [readonly] 
**DeltaToken** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewODataRawQueryOptions

`func NewODataRawQueryOptions() *ODataRawQueryOptions`

NewODataRawQueryOptions instantiates a new ODataRawQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewODataRawQueryOptionsWithDefaults

`func NewODataRawQueryOptionsWithDefaults() *ODataRawQueryOptions`

NewODataRawQueryOptionsWithDefaults instantiates a new ODataRawQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ODataRawQueryOptions) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ODataRawQueryOptions) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ODataRawQueryOptions) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ODataRawQueryOptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ODataRawQueryOptions) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ODataRawQueryOptions) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetApply

`func (o *ODataRawQueryOptions) GetApply() string`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *ODataRawQueryOptions) GetApplyOk() (*string, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *ODataRawQueryOptions) SetApply(v string)`

SetApply sets Apply field to given value.

### HasApply

`func (o *ODataRawQueryOptions) HasApply() bool`

HasApply returns a boolean if a field has been set.

### SetApplyNil

`func (o *ODataRawQueryOptions) SetApplyNil(b bool)`

 SetApplyNil sets the value for Apply to be an explicit nil

### UnsetApply
`func (o *ODataRawQueryOptions) UnsetApply()`

UnsetApply ensures that no value is present for Apply, not even an explicit nil
### GetCompute

`func (o *ODataRawQueryOptions) GetCompute() string`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *ODataRawQueryOptions) GetComputeOk() (*string, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *ODataRawQueryOptions) SetCompute(v string)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *ODataRawQueryOptions) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *ODataRawQueryOptions) SetComputeNil(b bool)`

 SetComputeNil sets the value for Compute to be an explicit nil

### UnsetCompute
`func (o *ODataRawQueryOptions) UnsetCompute()`

UnsetCompute ensures that no value is present for Compute, not even an explicit nil
### GetSearch

`func (o *ODataRawQueryOptions) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ODataRawQueryOptions) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ODataRawQueryOptions) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ODataRawQueryOptions) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ODataRawQueryOptions) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ODataRawQueryOptions) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetOrderBy

`func (o *ODataRawQueryOptions) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ODataRawQueryOptions) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ODataRawQueryOptions) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ODataRawQueryOptions) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *ODataRawQueryOptions) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *ODataRawQueryOptions) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetTop

`func (o *ODataRawQueryOptions) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *ODataRawQueryOptions) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *ODataRawQueryOptions) SetTop(v string)`

SetTop sets Top field to given value.

### HasTop

`func (o *ODataRawQueryOptions) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTopNil

`func (o *ODataRawQueryOptions) SetTopNil(b bool)`

 SetTopNil sets the value for Top to be an explicit nil

### UnsetTop
`func (o *ODataRawQueryOptions) UnsetTop()`

UnsetTop ensures that no value is present for Top, not even an explicit nil
### GetSkip

`func (o *ODataRawQueryOptions) GetSkip() string`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ODataRawQueryOptions) GetSkipOk() (*string, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ODataRawQueryOptions) SetSkip(v string)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ODataRawQueryOptions) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### SetSkipNil

`func (o *ODataRawQueryOptions) SetSkipNil(b bool)`

 SetSkipNil sets the value for Skip to be an explicit nil

### UnsetSkip
`func (o *ODataRawQueryOptions) UnsetSkip()`

UnsetSkip ensures that no value is present for Skip, not even an explicit nil
### GetSelect

`func (o *ODataRawQueryOptions) GetSelect() string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *ODataRawQueryOptions) GetSelectOk() (*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *ODataRawQueryOptions) SetSelect(v string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *ODataRawQueryOptions) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *ODataRawQueryOptions) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *ODataRawQueryOptions) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetExpand

`func (o *ODataRawQueryOptions) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *ODataRawQueryOptions) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *ODataRawQueryOptions) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *ODataRawQueryOptions) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### SetExpandNil

`func (o *ODataRawQueryOptions) SetExpandNil(b bool)`

 SetExpandNil sets the value for Expand to be an explicit nil

### UnsetExpand
`func (o *ODataRawQueryOptions) UnsetExpand()`

UnsetExpand ensures that no value is present for Expand, not even an explicit nil
### GetCount

`func (o *ODataRawQueryOptions) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ODataRawQueryOptions) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ODataRawQueryOptions) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *ODataRawQueryOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ODataRawQueryOptions) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ODataRawQueryOptions) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetFormat

`func (o *ODataRawQueryOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ODataRawQueryOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ODataRawQueryOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ODataRawQueryOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *ODataRawQueryOptions) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *ODataRawQueryOptions) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetSkipToken

`func (o *ODataRawQueryOptions) GetSkipToken() string`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *ODataRawQueryOptions) GetSkipTokenOk() (*string, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *ODataRawQueryOptions) SetSkipToken(v string)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *ODataRawQueryOptions) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### SetSkipTokenNil

`func (o *ODataRawQueryOptions) SetSkipTokenNil(b bool)`

 SetSkipTokenNil sets the value for SkipToken to be an explicit nil

### UnsetSkipToken
`func (o *ODataRawQueryOptions) UnsetSkipToken()`

UnsetSkipToken ensures that no value is present for SkipToken, not even an explicit nil
### GetDeltaToken

`func (o *ODataRawQueryOptions) GetDeltaToken() string`

GetDeltaToken returns the DeltaToken field if non-nil, zero value otherwise.

### GetDeltaTokenOk

`func (o *ODataRawQueryOptions) GetDeltaTokenOk() (*string, bool)`

GetDeltaTokenOk returns a tuple with the DeltaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaToken

`func (o *ODataRawQueryOptions) SetDeltaToken(v string)`

SetDeltaToken sets DeltaToken field to given value.

### HasDeltaToken

`func (o *ODataRawQueryOptions) HasDeltaToken() bool`

HasDeltaToken returns a boolean if a field has been set.

### SetDeltaTokenNil

`func (o *ODataRawQueryOptions) SetDeltaTokenNil(b bool)`

 SetDeltaTokenNil sets the value for DeltaToken to be an explicit nil

### UnsetDeltaToken
`func (o *ODataRawQueryOptions) UnsetDeltaToken()`

UnsetDeltaToken ensures that no value is present for DeltaToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


