# ManagementV1alpha1APIServiceComplianceRuntimeStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | Pointer to **float32** | The average risk score in the runtime compliance result. | [optional] 
**Grade** | Pointer to **string** | Grade result from the results summary. | [optional] 

## Methods

### NewManagementV1alpha1APIServiceComplianceRuntimeStatusResult

`func NewManagementV1alpha1APIServiceComplianceRuntimeStatusResult() *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult`

NewManagementV1alpha1APIServiceComplianceRuntimeStatusResult instantiates a new ManagementV1alpha1APIServiceComplianceRuntimeStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceComplianceRuntimeStatusResultWithDefaults

`func NewManagementV1alpha1APIServiceComplianceRuntimeStatusResultWithDefaults() *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult`

NewManagementV1alpha1APIServiceComplianceRuntimeStatusResultWithDefaults instantiates a new ManagementV1alpha1APIServiceComplianceRuntimeStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScore

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetGrade

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *ManagementV1alpha1APIServiceComplianceRuntimeStatusResult) HasGrade() bool`

HasGrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


