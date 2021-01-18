# ReporterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReporterEmail** | **string** | reporter email address, may be used for further correspondence | 
**ReporterName** | **string** | reporter name | 
**ReporterPhone** | Pointer to **string** | reporter phone number | [optional] 
**ReporterOrg** | Pointer to **string** | reporter organization | [optional] 
**DiscloseEmail** | **string** | allow response organization to share reporter email with an external incident resolver as appropriate? | 
**ReporterNotes** | **string** | any other relevant information about the reporter or reporter organization | 

## Methods

### NewReporterInfo

`func NewReporterInfo(reporterEmail string, reporterName string, discloseEmail string, reporterNotes string, ) *ReporterInfo`

NewReporterInfo instantiates a new ReporterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReporterInfoWithDefaults

`func NewReporterInfoWithDefaults() *ReporterInfo`

NewReporterInfoWithDefaults instantiates a new ReporterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReporterEmail

`func (o *ReporterInfo) GetReporterEmail() string`

GetReporterEmail returns the ReporterEmail field if non-nil, zero value otherwise.

### GetReporterEmailOk

`func (o *ReporterInfo) GetReporterEmailOk() (*string, bool)`

GetReporterEmailOk returns a tuple with the ReporterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterEmail

`func (o *ReporterInfo) SetReporterEmail(v string)`

SetReporterEmail sets ReporterEmail field to given value.


### GetReporterName

`func (o *ReporterInfo) GetReporterName() string`

GetReporterName returns the ReporterName field if non-nil, zero value otherwise.

### GetReporterNameOk

`func (o *ReporterInfo) GetReporterNameOk() (*string, bool)`

GetReporterNameOk returns a tuple with the ReporterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterName

`func (o *ReporterInfo) SetReporterName(v string)`

SetReporterName sets ReporterName field to given value.


### GetReporterPhone

`func (o *ReporterInfo) GetReporterPhone() string`

GetReporterPhone returns the ReporterPhone field if non-nil, zero value otherwise.

### GetReporterPhoneOk

`func (o *ReporterInfo) GetReporterPhoneOk() (*string, bool)`

GetReporterPhoneOk returns a tuple with the ReporterPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterPhone

`func (o *ReporterInfo) SetReporterPhone(v string)`

SetReporterPhone sets ReporterPhone field to given value.

### HasReporterPhone

`func (o *ReporterInfo) HasReporterPhone() bool`

HasReporterPhone returns a boolean if a field has been set.

### GetReporterOrg

`func (o *ReporterInfo) GetReporterOrg() string`

GetReporterOrg returns the ReporterOrg field if non-nil, zero value otherwise.

### GetReporterOrgOk

`func (o *ReporterInfo) GetReporterOrgOk() (*string, bool)`

GetReporterOrgOk returns a tuple with the ReporterOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterOrg

`func (o *ReporterInfo) SetReporterOrg(v string)`

SetReporterOrg sets ReporterOrg field to given value.

### HasReporterOrg

`func (o *ReporterInfo) HasReporterOrg() bool`

HasReporterOrg returns a boolean if a field has been set.

### GetDiscloseEmail

`func (o *ReporterInfo) GetDiscloseEmail() string`

GetDiscloseEmail returns the DiscloseEmail field if non-nil, zero value otherwise.

### GetDiscloseEmailOk

`func (o *ReporterInfo) GetDiscloseEmailOk() (*string, bool)`

GetDiscloseEmailOk returns a tuple with the DiscloseEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscloseEmail

`func (o *ReporterInfo) SetDiscloseEmail(v string)`

SetDiscloseEmail sets DiscloseEmail field to given value.


### GetReporterNotes

`func (o *ReporterInfo) GetReporterNotes() string`

GetReporterNotes returns the ReporterNotes field if non-nil, zero value otherwise.

### GetReporterNotesOk

`func (o *ReporterInfo) GetReporterNotesOk() (*string, bool)`

GetReporterNotesOk returns a tuple with the ReporterNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterNotes

`func (o *ReporterInfo) SetReporterNotes(v string)`

SetReporterNotes sets ReporterNotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


