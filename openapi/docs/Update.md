# Update

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Alias** | **string** |  | 
**DocumentTitle** | **string** |  | 
**Severity** | Pointer to **NullableString** |  | [optional] 
**InitialReleaseDate** | **time.Time** |  | 
**CurrentReleaseDate** | **time.Time** |  | 
**CvrfUrl** | **string** |  | 

## Methods

### NewUpdate

`func NewUpdate(id string, alias string, documentTitle string, initialReleaseDate time.Time, currentReleaseDate time.Time, cvrfUrl string, ) *Update`

NewUpdate instantiates a new Update object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWithDefaults

`func NewUpdateWithDefaults() *Update`

NewUpdateWithDefaults instantiates a new Update object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Update) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Update) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Update) SetId(v string)`

SetId sets Id field to given value.


### GetAlias

`func (o *Update) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Update) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Update) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetDocumentTitle

`func (o *Update) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *Update) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *Update) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.


### GetSeverity

`func (o *Update) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Update) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Update) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Update) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *Update) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *Update) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetInitialReleaseDate

`func (o *Update) GetInitialReleaseDate() time.Time`

GetInitialReleaseDate returns the InitialReleaseDate field if non-nil, zero value otherwise.

### GetInitialReleaseDateOk

`func (o *Update) GetInitialReleaseDateOk() (*time.Time, bool)`

GetInitialReleaseDateOk returns a tuple with the InitialReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialReleaseDate

`func (o *Update) SetInitialReleaseDate(v time.Time)`

SetInitialReleaseDate sets InitialReleaseDate field to given value.


### GetCurrentReleaseDate

`func (o *Update) GetCurrentReleaseDate() time.Time`

GetCurrentReleaseDate returns the CurrentReleaseDate field if non-nil, zero value otherwise.

### GetCurrentReleaseDateOk

`func (o *Update) GetCurrentReleaseDateOk() (*time.Time, bool)`

GetCurrentReleaseDateOk returns a tuple with the CurrentReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReleaseDate

`func (o *Update) SetCurrentReleaseDate(v time.Time)`

SetCurrentReleaseDate sets CurrentReleaseDate field to given value.


### GetCvrfUrl

`func (o *Update) GetCvrfUrl() string`

GetCvrfUrl returns the CvrfUrl field if non-nil, zero value otherwise.

### GetCvrfUrlOk

`func (o *Update) GetCvrfUrlOk() (*string, bool)`

GetCvrfUrlOk returns a tuple with the CvrfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvrfUrl

`func (o *Update) SetCvrfUrl(v string)`

SetCvrfUrl sets CvrfUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


