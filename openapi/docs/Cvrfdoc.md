# Cvrfdoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTitle** | Pointer to [**CvrfdocDocumentTitle**](cvrfdocDocumentTitle.md) |  | [optional] 
**DocumentType** | Pointer to [**CvrfdocDocumentType**](cvrfdocDocumentType.md) |  | [optional] 
**DocumentPublisher** | Pointer to [**CvrfdocDocumentPublisher**](cvrfdocDocumentPublisher.md) |  | [optional] 
**DocumentTracking** | Pointer to [**CvrfdocDocumentTracking**](cvrfdocDocumentTracking.md) |  | [optional] 
**DocumentNotes** | Pointer to [**[]CvrfdocNote**](CvrfdocNote.md) |  | [optional] 
**DocumentDistribution** | Pointer to [**CvrfdocDocumentDistribution**](cvrfdocDocumentDistribution.md) |  | [optional] 
**AggregateSeverity** | Pointer to [**CvrfdocAggregateSeverity**](cvrfdocAggregateSeverity.md) |  | [optional] 
**DocumentReferences** | Pointer to [**[]CvrfdocReference**](CvrfdocReference.md) |  | [optional] 
**Acknowledgments** | Pointer to [**[]CvrfdocAcknowledgment**](CvrfdocAcknowledgment.md) |  | [optional] 
**ProductTree** | Pointer to [**ProductTree**](ProductTree.md) |  | [optional] 
**Vulnerability** | Pointer to [**[]Vulnerability**](Vulnerability.md) |  | [optional] 

## Methods

### NewCvrfdoc

`func NewCvrfdoc() *Cvrfdoc`

NewCvrfdoc instantiates a new Cvrfdoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvrfdocWithDefaults

`func NewCvrfdocWithDefaults() *Cvrfdoc`

NewCvrfdocWithDefaults instantiates a new Cvrfdoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTitle

`func (o *Cvrfdoc) GetDocumentTitle() CvrfdocDocumentTitle`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *Cvrfdoc) GetDocumentTitleOk() (*CvrfdocDocumentTitle, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *Cvrfdoc) SetDocumentTitle(v CvrfdocDocumentTitle)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *Cvrfdoc) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetDocumentType

`func (o *Cvrfdoc) GetDocumentType() CvrfdocDocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *Cvrfdoc) GetDocumentTypeOk() (*CvrfdocDocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *Cvrfdoc) SetDocumentType(v CvrfdocDocumentType)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *Cvrfdoc) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentPublisher

`func (o *Cvrfdoc) GetDocumentPublisher() CvrfdocDocumentPublisher`

GetDocumentPublisher returns the DocumentPublisher field if non-nil, zero value otherwise.

### GetDocumentPublisherOk

`func (o *Cvrfdoc) GetDocumentPublisherOk() (*CvrfdocDocumentPublisher, bool)`

GetDocumentPublisherOk returns a tuple with the DocumentPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentPublisher

`func (o *Cvrfdoc) SetDocumentPublisher(v CvrfdocDocumentPublisher)`

SetDocumentPublisher sets DocumentPublisher field to given value.

### HasDocumentPublisher

`func (o *Cvrfdoc) HasDocumentPublisher() bool`

HasDocumentPublisher returns a boolean if a field has been set.

### GetDocumentTracking

`func (o *Cvrfdoc) GetDocumentTracking() CvrfdocDocumentTracking`

GetDocumentTracking returns the DocumentTracking field if non-nil, zero value otherwise.

### GetDocumentTrackingOk

`func (o *Cvrfdoc) GetDocumentTrackingOk() (*CvrfdocDocumentTracking, bool)`

GetDocumentTrackingOk returns a tuple with the DocumentTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTracking

`func (o *Cvrfdoc) SetDocumentTracking(v CvrfdocDocumentTracking)`

SetDocumentTracking sets DocumentTracking field to given value.

### HasDocumentTracking

`func (o *Cvrfdoc) HasDocumentTracking() bool`

HasDocumentTracking returns a boolean if a field has been set.

### GetDocumentNotes

`func (o *Cvrfdoc) GetDocumentNotes() []CvrfdocNote`

GetDocumentNotes returns the DocumentNotes field if non-nil, zero value otherwise.

### GetDocumentNotesOk

`func (o *Cvrfdoc) GetDocumentNotesOk() (*[]CvrfdocNote, bool)`

GetDocumentNotesOk returns a tuple with the DocumentNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNotes

`func (o *Cvrfdoc) SetDocumentNotes(v []CvrfdocNote)`

SetDocumentNotes sets DocumentNotes field to given value.

### HasDocumentNotes

`func (o *Cvrfdoc) HasDocumentNotes() bool`

HasDocumentNotes returns a boolean if a field has been set.

### GetDocumentDistribution

`func (o *Cvrfdoc) GetDocumentDistribution() CvrfdocDocumentDistribution`

GetDocumentDistribution returns the DocumentDistribution field if non-nil, zero value otherwise.

### GetDocumentDistributionOk

`func (o *Cvrfdoc) GetDocumentDistributionOk() (*CvrfdocDocumentDistribution, bool)`

GetDocumentDistributionOk returns a tuple with the DocumentDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDistribution

`func (o *Cvrfdoc) SetDocumentDistribution(v CvrfdocDocumentDistribution)`

SetDocumentDistribution sets DocumentDistribution field to given value.

### HasDocumentDistribution

`func (o *Cvrfdoc) HasDocumentDistribution() bool`

HasDocumentDistribution returns a boolean if a field has been set.

### GetAggregateSeverity

`func (o *Cvrfdoc) GetAggregateSeverity() CvrfdocAggregateSeverity`

GetAggregateSeverity returns the AggregateSeverity field if non-nil, zero value otherwise.

### GetAggregateSeverityOk

`func (o *Cvrfdoc) GetAggregateSeverityOk() (*CvrfdocAggregateSeverity, bool)`

GetAggregateSeverityOk returns a tuple with the AggregateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateSeverity

`func (o *Cvrfdoc) SetAggregateSeverity(v CvrfdocAggregateSeverity)`

SetAggregateSeverity sets AggregateSeverity field to given value.

### HasAggregateSeverity

`func (o *Cvrfdoc) HasAggregateSeverity() bool`

HasAggregateSeverity returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *Cvrfdoc) GetDocumentReferences() []CvrfdocReference`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *Cvrfdoc) GetDocumentReferencesOk() (*[]CvrfdocReference, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *Cvrfdoc) SetDocumentReferences(v []CvrfdocReference)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *Cvrfdoc) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetAcknowledgments

`func (o *Cvrfdoc) GetAcknowledgments() []CvrfdocAcknowledgment`

GetAcknowledgments returns the Acknowledgments field if non-nil, zero value otherwise.

### GetAcknowledgmentsOk

`func (o *Cvrfdoc) GetAcknowledgmentsOk() (*[]CvrfdocAcknowledgment, bool)`

GetAcknowledgmentsOk returns a tuple with the Acknowledgments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgments

`func (o *Cvrfdoc) SetAcknowledgments(v []CvrfdocAcknowledgment)`

SetAcknowledgments sets Acknowledgments field to given value.

### HasAcknowledgments

`func (o *Cvrfdoc) HasAcknowledgments() bool`

HasAcknowledgments returns a boolean if a field has been set.

### GetProductTree

`func (o *Cvrfdoc) GetProductTree() ProductTree`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *Cvrfdoc) GetProductTreeOk() (*ProductTree, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *Cvrfdoc) SetProductTree(v ProductTree)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *Cvrfdoc) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetVulnerability

`func (o *Cvrfdoc) GetVulnerability() []Vulnerability`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *Cvrfdoc) GetVulnerabilityOk() (*[]Vulnerability, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *Cvrfdoc) SetVulnerability(v []Vulnerability)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *Cvrfdoc) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


