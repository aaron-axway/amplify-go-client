# ManagementV1alpha1TraceabilityAgentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataplaneType** | **string** | The dataplane type that this agent connects to | 
**Config** | [**ManagementV1alpha1TraceabilityAgentSpecConfig**](ManagementV1alpha1TraceabilityAgentSpecConfig.md) |  | 
**Logging** | Pointer to [**ManagementV1alpha1DiscoveryAgentSpecLogging**](ManagementV1alpha1DiscoveryAgentSpecLogging.md) |  | [optional] 

## Methods

### NewManagementV1alpha1TraceabilityAgentSpec

`func NewManagementV1alpha1TraceabilityAgentSpec(dataplaneType string, config ManagementV1alpha1TraceabilityAgentSpecConfig, ) *ManagementV1alpha1TraceabilityAgentSpec`

NewManagementV1alpha1TraceabilityAgentSpec instantiates a new ManagementV1alpha1TraceabilityAgentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentSpecWithDefaults

`func NewManagementV1alpha1TraceabilityAgentSpecWithDefaults() *ManagementV1alpha1TraceabilityAgentSpec`

NewManagementV1alpha1TraceabilityAgentSpecWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataplaneType

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetDataplaneType() string`

GetDataplaneType returns the DataplaneType field if non-nil, zero value otherwise.

### GetDataplaneTypeOk

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetDataplaneTypeOk() (*string, bool)`

GetDataplaneTypeOk returns a tuple with the DataplaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplaneType

`func (o *ManagementV1alpha1TraceabilityAgentSpec) SetDataplaneType(v string)`

SetDataplaneType sets DataplaneType field to given value.


### GetConfig

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetConfig() ManagementV1alpha1TraceabilityAgentSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetConfigOk() (*ManagementV1alpha1TraceabilityAgentSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ManagementV1alpha1TraceabilityAgentSpec) SetConfig(v ManagementV1alpha1TraceabilityAgentSpecConfig)`

SetConfig sets Config field to given value.


### GetLogging

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetLogging() ManagementV1alpha1DiscoveryAgentSpecLogging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *ManagementV1alpha1TraceabilityAgentSpec) GetLoggingOk() (*ManagementV1alpha1DiscoveryAgentSpecLogging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *ManagementV1alpha1TraceabilityAgentSpec) SetLogging(v ManagementV1alpha1DiscoveryAgentSpecLogging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *ManagementV1alpha1TraceabilityAgentSpec) HasLogging() bool`

HasLogging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


