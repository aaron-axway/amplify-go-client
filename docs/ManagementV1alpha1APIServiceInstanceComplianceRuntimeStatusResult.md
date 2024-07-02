# ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Time when the update occurred. | [optional] 
**Grade** | Pointer to **string** | Grade result from the results summary. | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult

`func NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult() *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult`

NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult instantiates a new ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResultWithDefaults

`func NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResultWithDefaults() *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult`

NewManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResultWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGrade

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatusResult) HasGrade() bool`

HasGrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


