# ODataServiceDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeAnnotation** | Pointer to [**ODataTypeAnnotation**](ODataTypeAnnotation.md) |  | [optional] 
**EntitySets** | Pointer to [**[]ODataEntitySetInfo**](ODataEntitySetInfo.md) |  | [optional] 
**Singletons** | Pointer to [**[]ODataSingletonInfo**](ODataSingletonInfo.md) |  | [optional] 
**FunctionImports** | Pointer to [**[]ODataFunctionImportInfo**](ODataFunctionImportInfo.md) |  | [optional] 

## Methods

### NewODataServiceDocument

`func NewODataServiceDocument() *ODataServiceDocument`

NewODataServiceDocument instantiates a new ODataServiceDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewODataServiceDocumentWithDefaults

`func NewODataServiceDocumentWithDefaults() *ODataServiceDocument`

NewODataServiceDocumentWithDefaults instantiates a new ODataServiceDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeAnnotation

`func (o *ODataServiceDocument) GetTypeAnnotation() ODataTypeAnnotation`

GetTypeAnnotation returns the TypeAnnotation field if non-nil, zero value otherwise.

### GetTypeAnnotationOk

`func (o *ODataServiceDocument) GetTypeAnnotationOk() (*ODataTypeAnnotation, bool)`

GetTypeAnnotationOk returns a tuple with the TypeAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAnnotation

`func (o *ODataServiceDocument) SetTypeAnnotation(v ODataTypeAnnotation)`

SetTypeAnnotation sets TypeAnnotation field to given value.

### HasTypeAnnotation

`func (o *ODataServiceDocument) HasTypeAnnotation() bool`

HasTypeAnnotation returns a boolean if a field has been set.

### GetEntitySets

`func (o *ODataServiceDocument) GetEntitySets() []ODataEntitySetInfo`

GetEntitySets returns the EntitySets field if non-nil, zero value otherwise.

### GetEntitySetsOk

`func (o *ODataServiceDocument) GetEntitySetsOk() (*[]ODataEntitySetInfo, bool)`

GetEntitySetsOk returns a tuple with the EntitySets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySets

`func (o *ODataServiceDocument) SetEntitySets(v []ODataEntitySetInfo)`

SetEntitySets sets EntitySets field to given value.

### HasEntitySets

`func (o *ODataServiceDocument) HasEntitySets() bool`

HasEntitySets returns a boolean if a field has been set.

### SetEntitySetsNil

`func (o *ODataServiceDocument) SetEntitySetsNil(b bool)`

 SetEntitySetsNil sets the value for EntitySets to be an explicit nil

### UnsetEntitySets
`func (o *ODataServiceDocument) UnsetEntitySets()`

UnsetEntitySets ensures that no value is present for EntitySets, not even an explicit nil
### GetSingletons

`func (o *ODataServiceDocument) GetSingletons() []ODataSingletonInfo`

GetSingletons returns the Singletons field if non-nil, zero value otherwise.

### GetSingletonsOk

`func (o *ODataServiceDocument) GetSingletonsOk() (*[]ODataSingletonInfo, bool)`

GetSingletonsOk returns a tuple with the Singletons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingletons

`func (o *ODataServiceDocument) SetSingletons(v []ODataSingletonInfo)`

SetSingletons sets Singletons field to given value.

### HasSingletons

`func (o *ODataServiceDocument) HasSingletons() bool`

HasSingletons returns a boolean if a field has been set.

### SetSingletonsNil

`func (o *ODataServiceDocument) SetSingletonsNil(b bool)`

 SetSingletonsNil sets the value for Singletons to be an explicit nil

### UnsetSingletons
`func (o *ODataServiceDocument) UnsetSingletons()`

UnsetSingletons ensures that no value is present for Singletons, not even an explicit nil
### GetFunctionImports

`func (o *ODataServiceDocument) GetFunctionImports() []ODataFunctionImportInfo`

GetFunctionImports returns the FunctionImports field if non-nil, zero value otherwise.

### GetFunctionImportsOk

`func (o *ODataServiceDocument) GetFunctionImportsOk() (*[]ODataFunctionImportInfo, bool)`

GetFunctionImportsOk returns a tuple with the FunctionImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionImports

`func (o *ODataServiceDocument) SetFunctionImports(v []ODataFunctionImportInfo)`

SetFunctionImports sets FunctionImports field to given value.

### HasFunctionImports

`func (o *ODataServiceDocument) HasFunctionImports() bool`

HasFunctionImports returns a boolean if a field has been set.

### SetFunctionImportsNil

`func (o *ODataServiceDocument) SetFunctionImportsNil(b bool)`

 SetFunctionImportsNil sets the value for FunctionImports to be an explicit nil

### UnsetFunctionImports
`func (o *ODataServiceDocument) UnsetFunctionImports()`

UnsetFunctionImports ensures that no value is present for FunctionImports, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


