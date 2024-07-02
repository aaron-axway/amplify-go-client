# ManagementV1alpha1APIServiceRevisionComplianceLintingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the compliance job. | [optional] 
**ApiSpecLintingJob** | Pointer to **string** |  | [optional] 
**Ruleset** | Pointer to **string** | Ruleset logical name. | [optional] 
**RulesetFileName** | Pointer to **string** | File name of the APISpecLintingRuleset. | [optional] 
**Result** | Pointer to [**ManagementV1alpha1APIServiceRevisionComplianceLintingStatusResult**](ManagementV1alpha1APIServiceRevisionComplianceLintingStatusResult.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceRevisionComplianceLintingStatus

`func NewManagementV1alpha1APIServiceRevisionComplianceLintingStatus() *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus`

NewManagementV1alpha1APIServiceRevisionComplianceLintingStatus instantiates a new ManagementV1alpha1APIServiceRevisionComplianceLintingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceRevisionComplianceLintingStatusWithDefaults

`func NewManagementV1alpha1APIServiceRevisionComplianceLintingStatusWithDefaults() *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus`

NewManagementV1alpha1APIServiceRevisionComplianceLintingStatusWithDefaults instantiates a new ManagementV1alpha1APIServiceRevisionComplianceLintingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiSpecLintingJob

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetApiSpecLintingJob() string`

GetApiSpecLintingJob returns the ApiSpecLintingJob field if non-nil, zero value otherwise.

### GetApiSpecLintingJobOk

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetApiSpecLintingJobOk() (*string, bool)`

GetApiSpecLintingJobOk returns a tuple with the ApiSpecLintingJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSpecLintingJob

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) SetApiSpecLintingJob(v string)`

SetApiSpecLintingJob sets ApiSpecLintingJob field to given value.

### HasApiSpecLintingJob

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) HasApiSpecLintingJob() bool`

HasApiSpecLintingJob returns a boolean if a field has been set.

### GetRuleset

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetRuleset() string`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetRulesetOk() (*string, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) SetRuleset(v string)`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetRulesetFileName

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetRulesetFileName() string`

GetRulesetFileName returns the RulesetFileName field if non-nil, zero value otherwise.

### GetRulesetFileNameOk

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetRulesetFileNameOk() (*string, bool)`

GetRulesetFileNameOk returns a tuple with the RulesetFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetFileName

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) SetRulesetFileName(v string)`

SetRulesetFileName sets RulesetFileName field to given value.

### HasRulesetFileName

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) HasRulesetFileName() bool`

HasRulesetFileName returns a boolean if a field has been set.

### GetResult

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetResult() ManagementV1alpha1APIServiceRevisionComplianceLintingStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) GetResultOk() (*ManagementV1alpha1APIServiceRevisionComplianceLintingStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) SetResult(v ManagementV1alpha1APIServiceRevisionComplianceLintingStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ManagementV1alpha1APIServiceRevisionComplianceLintingStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


