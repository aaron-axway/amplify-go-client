# ManagementV1alpha1APISpecLintingJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**[]ManagementV1APISpecLintingJobResultDetailsInner**](ManagementV1APISpecLintingJobResultDetailsInner.md) | The API Specification Linting Result details. | 
**Summary** | [**ManagementV1APISpecLintingJobResultSummary**](ManagementV1APISpecLintingJobResultSummary.md) |  | 
**ApiSpecLintingRulesetRevision** | Pointer to **int32** | Reference to the APISpecLintingRuleset revision | [optional] 
**DetailsThresholdExceeded** | Pointer to **bool** | Set the value to true if the linting result details count has reached the threshold | [optional] 

## Methods

### NewManagementV1alpha1APISpecLintingJobResult

`func NewManagementV1alpha1APISpecLintingJobResult(details []ManagementV1APISpecLintingJobResultDetailsInner, summary ManagementV1APISpecLintingJobResultSummary, ) *ManagementV1alpha1APISpecLintingJobResult`

NewManagementV1alpha1APISpecLintingJobResult instantiates a new ManagementV1alpha1APISpecLintingJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APISpecLintingJobResultWithDefaults

`func NewManagementV1alpha1APISpecLintingJobResultWithDefaults() *ManagementV1alpha1APISpecLintingJobResult`

NewManagementV1alpha1APISpecLintingJobResultWithDefaults instantiates a new ManagementV1alpha1APISpecLintingJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetDetails() []ManagementV1APISpecLintingJobResultDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetDetailsOk() (*[]ManagementV1APISpecLintingJobResultDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManagementV1alpha1APISpecLintingJobResult) SetDetails(v []ManagementV1APISpecLintingJobResultDetailsInner)`

SetDetails sets Details field to given value.


### GetSummary

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetSummary() ManagementV1APISpecLintingJobResultSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetSummaryOk() (*ManagementV1APISpecLintingJobResultSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ManagementV1alpha1APISpecLintingJobResult) SetSummary(v ManagementV1APISpecLintingJobResultSummary)`

SetSummary sets Summary field to given value.


### GetApiSpecLintingRulesetRevision

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetApiSpecLintingRulesetRevision() int32`

GetApiSpecLintingRulesetRevision returns the ApiSpecLintingRulesetRevision field if non-nil, zero value otherwise.

### GetApiSpecLintingRulesetRevisionOk

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetApiSpecLintingRulesetRevisionOk() (*int32, bool)`

GetApiSpecLintingRulesetRevisionOk returns a tuple with the ApiSpecLintingRulesetRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSpecLintingRulesetRevision

`func (o *ManagementV1alpha1APISpecLintingJobResult) SetApiSpecLintingRulesetRevision(v int32)`

SetApiSpecLintingRulesetRevision sets ApiSpecLintingRulesetRevision field to given value.

### HasApiSpecLintingRulesetRevision

`func (o *ManagementV1alpha1APISpecLintingJobResult) HasApiSpecLintingRulesetRevision() bool`

HasApiSpecLintingRulesetRevision returns a boolean if a field has been set.

### GetDetailsThresholdExceeded

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetDetailsThresholdExceeded() bool`

GetDetailsThresholdExceeded returns the DetailsThresholdExceeded field if non-nil, zero value otherwise.

### GetDetailsThresholdExceededOk

`func (o *ManagementV1alpha1APISpecLintingJobResult) GetDetailsThresholdExceededOk() (*bool, bool)`

GetDetailsThresholdExceededOk returns a tuple with the DetailsThresholdExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsThresholdExceeded

`func (o *ManagementV1alpha1APISpecLintingJobResult) SetDetailsThresholdExceeded(v bool)`

SetDetailsThresholdExceeded sets DetailsThresholdExceeded field to given value.

### HasDetailsThresholdExceeded

`func (o *ManagementV1alpha1APISpecLintingJobResult) HasDetailsThresholdExceeded() bool`

HasDetailsThresholdExceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


