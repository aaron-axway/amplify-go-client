# CatalogV1alpha1StageVisibilitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stages** | [**[]CatalogV1alpha1StageVisibilitySpecStagesInner**](CatalogV1alpha1StageVisibilitySpecStagesInner.md) | Defines where the visibility settings apply. | 
**Exclude** | Pointer to **bool** | Determines if the list of subjects should be excluded from the stage visibility. | [optional] 
**Subjects** | Pointer to [**[]CatalogV1alpha1StageVisibilitySpecSubjectsInner**](CatalogV1alpha1StageVisibilitySpecSubjectsInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1StageVisibilitySpec

`func NewCatalogV1alpha1StageVisibilitySpec(stages []CatalogV1alpha1StageVisibilitySpecStagesInner, ) *CatalogV1alpha1StageVisibilitySpec`

NewCatalogV1alpha1StageVisibilitySpec instantiates a new CatalogV1alpha1StageVisibilitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1StageVisibilitySpecWithDefaults

`func NewCatalogV1alpha1StageVisibilitySpecWithDefaults() *CatalogV1alpha1StageVisibilitySpec`

NewCatalogV1alpha1StageVisibilitySpecWithDefaults instantiates a new CatalogV1alpha1StageVisibilitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStages

`func (o *CatalogV1alpha1StageVisibilitySpec) GetStages() []CatalogV1alpha1StageVisibilitySpecStagesInner`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CatalogV1alpha1StageVisibilitySpec) GetStagesOk() (*[]CatalogV1alpha1StageVisibilitySpecStagesInner, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CatalogV1alpha1StageVisibilitySpec) SetStages(v []CatalogV1alpha1StageVisibilitySpecStagesInner)`

SetStages sets Stages field to given value.


### GetExclude

`func (o *CatalogV1alpha1StageVisibilitySpec) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CatalogV1alpha1StageVisibilitySpec) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CatalogV1alpha1StageVisibilitySpec) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CatalogV1alpha1StageVisibilitySpec) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetSubjects

`func (o *CatalogV1alpha1StageVisibilitySpec) GetSubjects() []CatalogV1alpha1StageVisibilitySpecSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CatalogV1alpha1StageVisibilitySpec) GetSubjectsOk() (*[]CatalogV1alpha1StageVisibilitySpecSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CatalogV1alpha1StageVisibilitySpec) SetSubjects(v []CatalogV1alpha1StageVisibilitySpecSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CatalogV1alpha1StageVisibilitySpec) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


