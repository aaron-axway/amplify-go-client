# ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Time when the update occurred. | [optional] 
**RiskScore** | Pointer to **float32** | The average risk score in the runtime compliance result. | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult

`func NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult() *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult`

NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult instantiates a new ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResultWithDefaults

`func NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResultWithDefaults() *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult`

NewManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResultWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRiskScore

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatusResult) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


