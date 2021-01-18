# TestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | date testing will commence | 
**EndDate** | **string** | date testing will end | 
**TestDescription** | **string** | detailed summary of the planned testing, including types of tests and a list of targeted assets (IP or FQDN) | 

## Methods

### NewTestInfo

`func NewTestInfo(startDate string, endDate string, testDescription string, ) *TestInfo`

NewTestInfo instantiates a new TestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoWithDefaults

`func NewTestInfoWithDefaults() *TestInfo`

NewTestInfoWithDefaults instantiates a new TestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *TestInfo) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestInfo) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestInfo) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *TestInfo) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestInfo) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestInfo) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetTestDescription

`func (o *TestInfo) GetTestDescription() string`

GetTestDescription returns the TestDescription field if non-nil, zero value otherwise.

### GetTestDescriptionOk

`func (o *TestInfo) GetTestDescriptionOk() (*string, bool)`

GetTestDescriptionOk returns a tuple with the TestDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDescription

`func (o *TestInfo) SetTestDescription(v string)`

SetTestDescription sets TestDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


