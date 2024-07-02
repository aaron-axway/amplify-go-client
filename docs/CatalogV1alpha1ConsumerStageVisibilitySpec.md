# CatalogV1alpha1ConsumerStageVisibilitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]CatalogV1alpha1StageVisibilitySpecStagesInner**](CatalogV1alpha1StageVisibilitySpecStagesInner.md) | Defines where the visibility settings apply. | 
**Subjects** | Pointer to [**[]CatalogV1alpha1ConsumerStageVisibilitySpecSubjectsInner**](CatalogV1alpha1ConsumerStageVisibilitySpecSubjectsInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ConsumerStageVisibilitySpec

`func NewCatalogV1alpha1ConsumerStageVisibilitySpec(stages []CatalogV1alpha1StageVisibilitySpecStagesInner, ) *CatalogV1alpha1ConsumerStageVisibilitySpec`

NewCatalogV1alpha1ConsumerStageVisibilitySpec instantiates a new CatalogV1alpha1ConsumerStageVisibilitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ConsumerStageVisibilitySpecWithDefaults

`func NewCatalogV1alpha1ConsumerStageVisibilitySpecWithDefaults() *CatalogV1alpha1ConsumerStageVisibilitySpec`

NewCatalogV1alpha1ConsumerStageVisibilitySpecWithDefaults instantiates a new CatalogV1alpha1ConsumerStageVisibilitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) GetStages() []CatalogV1alpha1StageVisibilitySpecStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) GetStagesOk() (*[]CatalogV1alpha1StageVisibilitySpecStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) SetStages(v []CatalogV1alpha1StageVisibilitySpecStagesInner)`

SetStages sets Stages field to given value.


### GetSubjects

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) GetSubjects() []CatalogV1alpha1ConsumerStageVisibilitySpecSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) GetSubjectsOk() (*[]CatalogV1alpha1ConsumerStageVisibilitySpecSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) SetSubjects(v []CatalogV1alpha1ConsumerStageVisibilitySpecSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CatalogV1alpha1ConsumerStageVisibilitySpec) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


