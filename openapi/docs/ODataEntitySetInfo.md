# ODataEntitySetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeAnnotation** | Pointer to [**ODataTypeAnnotation**](ODataTypeAnnotation.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewODataEntitySetInfo

`func NewODataEntitySetInfo() *ODataEntitySetInfo`

NewODataEntitySetInfo instantiates a new ODataEntitySetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewODataEntitySetInfoWithDefaults

`func NewODataEntitySetInfoWithDefaults() *ODataEntitySetInfo`

NewODataEntitySetInfoWithDefaults instantiates a new ODataEntitySetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeAnnotation

`func (o *ODataEntitySetInfo) GetTypeAnnotation() ODataTypeAnnotation`

GetTypeAnnotation returns the TypeAnnotation field if non-nil, zero value otherwise.

### GetTypeAnnotationOk

`func (o *ODataEntitySetInfo) GetTypeAnnotationOk() (*ODataTypeAnnotation, bool)`

GetTypeAnnotationOk returns a tuple with the TypeAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAnnotation

`func (o *ODataEntitySetInfo) SetTypeAnnotation(v ODataTypeAnnotation)`

SetTypeAnnotation sets TypeAnnotation field to given value.

### HasTypeAnnotation

`func (o *ODataEntitySetInfo) HasTypeAnnotation() bool`

HasTypeAnnotation returns a boolean if a field has been set.

### GetUrl

`func (o *ODataEntitySetInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ODataEntitySetInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ODataEntitySetInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ODataEntitySetInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ODataEntitySetInfo) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ODataEntitySetInfo) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetName

`func (o *ODataEntitySetInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ODataEntitySetInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ODataEntitySetInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ODataEntitySetInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ODataEntitySetInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ODataEntitySetInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *ODataEntitySetInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ODataEntitySetInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ODataEntitySetInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ODataEntitySetInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ODataEntitySetInfo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ODataEntitySetInfo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


