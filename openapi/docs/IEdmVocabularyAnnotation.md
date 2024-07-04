# IEdmVocabularyAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualifier** | Pointer to **NullableString** |  | [optional] [readonly] 
**Term** | Pointer to [**IEdmTerm**](IEdmTerm.md) |  | [optional] 
**Target** | Pointer to **map[string]interface{}** |  | [optional] 
**Value** | Pointer to [**IEdmExpression**](IEdmExpression.md) |  | [optional] 

## Methods

### NewIEdmVocabularyAnnotation

`func NewIEdmVocabularyAnnotation() *IEdmVocabularyAnnotation`

NewIEdmVocabularyAnnotation instantiates a new IEdmVocabularyAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIEdmVocabularyAnnotationWithDefaults

`func NewIEdmVocabularyAnnotationWithDefaults() *IEdmVocabularyAnnotation`

NewIEdmVocabularyAnnotationWithDefaults instantiates a new IEdmVocabularyAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualifier

`func (o *IEdmVocabularyAnnotation) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *IEdmVocabularyAnnotation) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *IEdmVocabularyAnnotation) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *IEdmVocabularyAnnotation) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### SetQualifierNil

`func (o *IEdmVocabularyAnnotation) SetQualifierNil(b bool)`

 SetQualifierNil sets the value for Qualifier to be an explicit nil

### UnsetQualifier
`func (o *IEdmVocabularyAnnotation) UnsetQualifier()`

UnsetQualifier ensures that no value is present for Qualifier, not even an explicit nil
### GetTerm

`func (o *IEdmVocabularyAnnotation) GetTerm() IEdmTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *IEdmVocabularyAnnotation) GetTermOk() (*IEdmTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *IEdmVocabularyAnnotation) SetTerm(v IEdmTerm)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *IEdmVocabularyAnnotation) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTarget

`func (o *IEdmVocabularyAnnotation) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IEdmVocabularyAnnotation) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IEdmVocabularyAnnotation) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.

### HasTarget

`func (o *IEdmVocabularyAnnotation) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetValue

`func (o *IEdmVocabularyAnnotation) GetValue() IEdmExpression`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IEdmVocabularyAnnotation) GetValueOk() (*IEdmExpression, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IEdmVocabularyAnnotation) SetValue(v IEdmExpression)`

SetValue sets Value field to given value.

### HasValue

`func (o *IEdmVocabularyAnnotation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


