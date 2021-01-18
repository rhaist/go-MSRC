# CarsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSubmission** | **bool** | is this a test submission? If so, the API will validate the input but won&#39;t submit it to MSRC for action | 
**CarsAdvancedCheck** | Pointer to **bool** | do the reports include detailed fields? This can be false only for form submissions (API users should ignore it) | [optional] 
**ReporterInfo** | [**ReporterInfo**](ReporterInfo.md) |  | 
**Reports** | [**[]Report**](Report.md) | reports of abuse | 

## Methods

### NewCarsDto

`func NewCarsDto(testSubmission bool, reporterInfo ReporterInfo, reports []Report, ) *CarsDto`

NewCarsDto instantiates a new CarsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarsDtoWithDefaults

`func NewCarsDtoWithDefaults() *CarsDto`

NewCarsDtoWithDefaults instantiates a new CarsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSubmission

`func (o *CarsDto) GetTestSubmission() bool`

GetTestSubmission returns the TestSubmission field if non-nil, zero value otherwise.

### GetTestSubmissionOk

`func (o *CarsDto) GetTestSubmissionOk() (*bool, bool)`

GetTestSubmissionOk returns a tuple with the TestSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSubmission

`func (o *CarsDto) SetTestSubmission(v bool)`

SetTestSubmission sets TestSubmission field to given value.


### GetCarsAdvancedCheck

`func (o *CarsDto) GetCarsAdvancedCheck() bool`

GetCarsAdvancedCheck returns the CarsAdvancedCheck field if non-nil, zero value otherwise.

### GetCarsAdvancedCheckOk

`func (o *CarsDto) GetCarsAdvancedCheckOk() (*bool, bool)`

GetCarsAdvancedCheckOk returns a tuple with the CarsAdvancedCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarsAdvancedCheck

`func (o *CarsDto) SetCarsAdvancedCheck(v bool)`

SetCarsAdvancedCheck sets CarsAdvancedCheck field to given value.

### HasCarsAdvancedCheck

`func (o *CarsDto) HasCarsAdvancedCheck() bool`

HasCarsAdvancedCheck returns a boolean if a field has been set.

### GetReporterInfo

`func (o *CarsDto) GetReporterInfo() ReporterInfo`

GetReporterInfo returns the ReporterInfo field if non-nil, zero value otherwise.

### GetReporterInfoOk

`func (o *CarsDto) GetReporterInfoOk() (*ReporterInfo, bool)`

GetReporterInfoOk returns a tuple with the ReporterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterInfo

`func (o *CarsDto) SetReporterInfo(v ReporterInfo)`

SetReporterInfo sets ReporterInfo field to given value.


### GetReports

`func (o *CarsDto) GetReports() []Report`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *CarsDto) GetReportsOk() (*[]Report, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *CarsDto) SetReports(v []Report)`

SetReports sets Reports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


