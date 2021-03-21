# CvrfdocDocumentTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identification** | Pointer to [**CvrfdocDocumentTrackingIdentification**](CvrfdocDocumentTrackingIdentification.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**RevisionHistory** | Pointer to [**[]CvrfdocDocumentTrackingRevision**](CvrfdocDocumentTrackingRevision.md) |  | [optional] 
**InitialReleaseDate** | Pointer to **time.Time** |  | [optional] 
**CurrentReleaseDate** | Pointer to **time.Time** |  | [optional] 
**Generator** | Pointer to [**CvrfdocDocumentTrackingGenerator**](CvrfdocDocumentTrackingGenerator.md) |  | [optional] 

## Methods

### NewCvrfdocDocumentTracking

`func NewCvrfdocDocumentTracking() *CvrfdocDocumentTracking`

NewCvrfdocDocumentTracking instantiates a new CvrfdocDocumentTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvrfdocDocumentTrackingWithDefaults

`func NewCvrfdocDocumentTrackingWithDefaults() *CvrfdocDocumentTracking`

NewCvrfdocDocumentTrackingWithDefaults instantiates a new CvrfdocDocumentTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentification

`func (o *CvrfdocDocumentTracking) GetIdentification() CvrfdocDocumentTrackingIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *CvrfdocDocumentTracking) GetIdentificationOk() (*CvrfdocDocumentTrackingIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *CvrfdocDocumentTracking) SetIdentification(v CvrfdocDocumentTrackingIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *CvrfdocDocumentTracking) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetStatus

`func (o *CvrfdocDocumentTracking) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CvrfdocDocumentTracking) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CvrfdocDocumentTracking) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CvrfdocDocumentTracking) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *CvrfdocDocumentTracking) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CvrfdocDocumentTracking) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CvrfdocDocumentTracking) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CvrfdocDocumentTracking) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRevisionHistory

`func (o *CvrfdocDocumentTracking) GetRevisionHistory() []CvrfdocDocumentTrackingRevision`

GetRevisionHistory returns the RevisionHistory field if non-nil, zero value otherwise.

### GetRevisionHistoryOk

`func (o *CvrfdocDocumentTracking) GetRevisionHistoryOk() (*[]CvrfdocDocumentTrackingRevision, bool)`

GetRevisionHistoryOk returns a tuple with the RevisionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistory

`func (o *CvrfdocDocumentTracking) SetRevisionHistory(v []CvrfdocDocumentTrackingRevision)`

SetRevisionHistory sets RevisionHistory field to given value.

### HasRevisionHistory

`func (o *CvrfdocDocumentTracking) HasRevisionHistory() bool`

HasRevisionHistory returns a boolean if a field has been set.

### GetInitialReleaseDate

`func (o *CvrfdocDocumentTracking) GetInitialReleaseDate() time.Time`

GetInitialReleaseDate returns the InitialReleaseDate field if non-nil, zero value otherwise.

### GetInitialReleaseDateOk

`func (o *CvrfdocDocumentTracking) GetInitialReleaseDateOk() (*time.Time, bool)`

GetInitialReleaseDateOk returns a tuple with the InitialReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialReleaseDate

`func (o *CvrfdocDocumentTracking) SetInitialReleaseDate(v time.Time)`

SetInitialReleaseDate sets InitialReleaseDate field to given value.

### HasInitialReleaseDate

`func (o *CvrfdocDocumentTracking) HasInitialReleaseDate() bool`

HasInitialReleaseDate returns a boolean if a field has been set.

### GetCurrentReleaseDate

`func (o *CvrfdocDocumentTracking) GetCurrentReleaseDate() time.Time`

GetCurrentReleaseDate returns the CurrentReleaseDate field if non-nil, zero value otherwise.

### GetCurrentReleaseDateOk

`func (o *CvrfdocDocumentTracking) GetCurrentReleaseDateOk() (*time.Time, bool)`

GetCurrentReleaseDateOk returns a tuple with the CurrentReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReleaseDate

`func (o *CvrfdocDocumentTracking) SetCurrentReleaseDate(v time.Time)`

SetCurrentReleaseDate sets CurrentReleaseDate field to given value.

### HasCurrentReleaseDate

`func (o *CvrfdocDocumentTracking) HasCurrentReleaseDate() bool`

HasCurrentReleaseDate returns a boolean if a field has been set.

### GetGenerator

`func (o *CvrfdocDocumentTracking) GetGenerator() CvrfdocDocumentTrackingGenerator`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *CvrfdocDocumentTracking) GetGeneratorOk() (*CvrfdocDocumentTrackingGenerator, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *CvrfdocDocumentTracking) SetGenerator(v CvrfdocDocumentTrackingGenerator)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *CvrfdocDocumentTracking) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


