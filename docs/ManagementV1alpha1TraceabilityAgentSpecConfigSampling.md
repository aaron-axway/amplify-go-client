# ManagementV1alpha1TraceabilityAgentSpecConfigSampling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **float32** | The percentage of transactions that the agent will save the full transaction details for. Valid values(decimals included) are 0 to 10 | [default to 1]
**AllErrors** | **bool** | When set to true will send transactional data for every errored transaction for Business and Consumer Insights | 

## Methods

### NewManagementV1alpha1TraceabilityAgentSpecConfigSampling

`func NewManagementV1alpha1TraceabilityAgentSpecConfigSampling(percentage float32, allErrors bool, ) *ManagementV1alpha1TraceabilityAgentSpecConfigSampling`

NewManagementV1alpha1TraceabilityAgentSpecConfigSampling instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigSampling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentSpecConfigSamplingWithDefaults

`func NewManagementV1alpha1TraceabilityAgentSpecConfigSamplingWithDefaults() *ManagementV1alpha1TraceabilityAgentSpecConfigSampling`

NewManagementV1alpha1TraceabilityAgentSpecConfigSamplingWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfigSampling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetAllErrors

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) GetAllErrors() bool`

GetAllErrors returns the AllErrors field if non-nil, zero value otherwise.

### GetAllErrorsOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) GetAllErrorsOk() (*bool, bool)`

GetAllErrorsOk returns a tuple with the AllErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllErrors

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfigSampling) SetAllErrors(v bool)`

SetAllErrors sets AllErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


