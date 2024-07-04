# IEdmPathExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathSegments** | Pointer to **[]string** |  | [optional] [readonly] 
**Path** | Pointer to **NullableString** |  | [optional] [readonly] 
**ExpressionKind** | Pointer to [**EdmExpressionKind**](EdmExpressionKind.md) |  | [optional] 

## Methods

### NewIEdmPathExpression

`func NewIEdmPathExpression() *IEdmPathExpression`

NewIEdmPathExpression instantiates a new IEdmPathExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmPathExpressionWithDefaults

`func NewIEdmPathExpressionWithDefaults() *IEdmPathExpression`

NewIEdmPathExpressionWithDefaults instantiates a new IEdmPathExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathSegments

`func (o *IEdmPathExpression) GetPathSegments() []string`

GetPathSegments returns the PathSegments field if non-nil, zero value otherwise.

### GetPathSegmentsOk

`func (o *IEdmPathExpression) GetPathSegmentsOk() (*[]string, bool)`

GetPathSegmentsOk returns a tuple with the PathSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSegments

`func (o *IEdmPathExpression) SetPathSegments(v []string)`

SetPathSegments sets PathSegments field to given value.

### HasPathSegments

`func (o *IEdmPathExpression) HasPathSegments() bool`

HasPathSegments returns a boolean if a field has been set.

### SetPathSegmentsNil

`func (o *IEdmPathExpression) SetPathSegmentsNil(b bool)`

 SetPathSegmentsNil sets the value for PathSegments to be an explicit nil

### UnsetPathSegments
`func (o *IEdmPathExpression) UnsetPathSegments()`

UnsetPathSegments ensures that no value is present for PathSegments, not even an explicit nil
### GetPath

`func (o *IEdmPathExpression) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IEdmPathExpression) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IEdmPathExpression) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IEdmPathExpression) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *IEdmPathExpression) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *IEdmPathExpression) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetExpressionKind

`func (o *IEdmPathExpression) GetExpressionKind() EdmExpressionKind`

GetExpressionKind returns the ExpressionKind field if non-nil, zero value otherwise.

### GetExpressionKindOk

`func (o *IEdmPathExpression) GetExpressionKindOk() (*EdmExpressionKind, bool)`

GetExpressionKindOk returns a tuple with the ExpressionKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionKind

`func (o *IEdmPathExpression) SetExpressionKind(v EdmExpressionKind)`

SetExpressionKind sets ExpressionKind field to given value.

### HasExpressionKind

`func (o *IEdmPathExpression) HasExpressionKind() bool`

HasExpressionKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


