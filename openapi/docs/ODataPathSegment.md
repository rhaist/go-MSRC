# ODataPathSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdmType** | Pointer to [**IEdmType**](IEdmType.md) |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewODataPathSegment

`func NewODataPathSegment() *ODataPathSegment`

NewODataPathSegment instantiates a new ODataPathSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewODataPathSegmentWithDefaults

`func NewODataPathSegmentWithDefaults() *ODataPathSegment`

NewODataPathSegmentWithDefaults instantiates a new ODataPathSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdmType

`func (o *ODataPathSegment) GetEdmType() IEdmType`

GetEdmType returns the EdmType field if non-nil, zero value otherwise.

### GetEdmTypeOk

`func (o *ODataPathSegment) GetEdmTypeOk() (*IEdmType, bool)`

GetEdmTypeOk returns a tuple with the EdmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdmType

`func (o *ODataPathSegment) SetEdmType(v IEdmType)`

SetEdmType sets EdmType field to given value.

### HasEdmType

`func (o *ODataPathSegment) HasEdmType() bool`

HasEdmType returns a boolean if a field has been set.

### GetIdentifier

`func (o *ODataPathSegment) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ODataPathSegment) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ODataPathSegment) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ODataPathSegment) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *ODataPathSegment) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *ODataPathSegment) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


