# Report1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchID** | Pointer to **string** | tag and group reports | [optional] 
**RelatedCases** | Pointer to **[]string** | IDs of known related cases | [optional] 
**ReportNotes** | **string** | a brief summary of what occurred | 
**TLP** | Pointer to **string** | [Traffic Light Protocol](https://www.us-cert.gov/tlp) color (RED, AMBER, GREEN, or WHITE) to regulate the level of  disclosure | [optional] 
**DisclosureNotes** | Pointer to **string** | any additional limits on disclosure | [optional] 
**Threats** | [**[]Threat**](Threat.md) | instances of abuse | 

## Methods

### NewReport1

`func NewReport1(reportNotes string, threats []Threat, ) *Report1`

NewReport1 instantiates a new Report1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReport1WithDefaults

`func NewReport1WithDefaults() *Report1`

NewReport1WithDefaults instantiates a new Report1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchID

`func (o *Report1) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *Report1) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *Report1) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.

### HasBatchID

`func (o *Report1) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.

### GetRelatedCases

`func (o *Report1) GetRelatedCases() []string`

GetRelatedCases returns the RelatedCases field if non-nil, zero value otherwise.

### GetRelatedCasesOk

`func (o *Report1) GetRelatedCasesOk() (*[]string, bool)`

GetRelatedCasesOk returns a tuple with the RelatedCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCases

`func (o *Report1) SetRelatedCases(v []string)`

SetRelatedCases sets RelatedCases field to given value.

### HasRelatedCases

`func (o *Report1) HasRelatedCases() bool`

HasRelatedCases returns a boolean if a field has been set.

### GetReportNotes

`func (o *Report1) GetReportNotes() string`

GetReportNotes returns the ReportNotes field if non-nil, zero value otherwise.

### GetReportNotesOk

`func (o *Report1) GetReportNotesOk() (*string, bool)`

GetReportNotesOk returns a tuple with the ReportNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportNotes

`func (o *Report1) SetReportNotes(v string)`

SetReportNotes sets ReportNotes field to given value.


### GetTLP

`func (o *Report1) GetTLP() string`

GetTLP returns the TLP field if non-nil, zero value otherwise.

### GetTLPOk

`func (o *Report1) GetTLPOk() (*string, bool)`

GetTLPOk returns a tuple with the TLP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLP

`func (o *Report1) SetTLP(v string)`

SetTLP sets TLP field to given value.

### HasTLP

`func (o *Report1) HasTLP() bool`

HasTLP returns a boolean if a field has been set.

### GetDisclosureNotes

`func (o *Report1) GetDisclosureNotes() string`

GetDisclosureNotes returns the DisclosureNotes field if non-nil, zero value otherwise.

### GetDisclosureNotesOk

`func (o *Report1) GetDisclosureNotesOk() (*string, bool)`

GetDisclosureNotesOk returns a tuple with the DisclosureNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureNotes

`func (o *Report1) SetDisclosureNotes(v string)`

SetDisclosureNotes sets DisclosureNotes field to given value.

### HasDisclosureNotes

`func (o *Report1) HasDisclosureNotes() bool`

HasDisclosureNotes returns a boolean if a field has been set.

### GetThreats

`func (o *Report1) GetThreats() []Threat`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *Report1) GetThreatsOk() (*[]Threat, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *Report1) SetThreats(v []Threat)`

SetThreats sets Threats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


