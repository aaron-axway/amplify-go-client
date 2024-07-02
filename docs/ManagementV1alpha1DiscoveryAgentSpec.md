# ManagementV1alpha1DiscoveryAgentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataplaneType** | **string** | The dataplane type that this agent connects to | 
**Config** | [**ManagementV1alpha1DiscoveryAgentSpecConfig**](ManagementV1alpha1DiscoveryAgentSpecConfig.md) |  | 
**Logging** | Pointer to [**ManagementV1alpha1DiscoveryAgentSpecLogging**](ManagementV1alpha1DiscoveryAgentSpecLogging.md) |  | [optional] 

## Methods

### NewManagementV1alpha1DiscoveryAgentSpec

`func NewManagementV1alpha1DiscoveryAgentSpec(dataplaneType string, config ManagementV1alpha1DiscoveryAgentSpecConfig, ) *ManagementV1alpha1DiscoveryAgentSpec`

NewManagementV1alpha1DiscoveryAgentSpec instantiates a new ManagementV1alpha1DiscoveryAgentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DiscoveryAgentSpecWithDefaults

`func NewManagementV1alpha1DiscoveryAgentSpecWithDefaults() *ManagementV1alpha1DiscoveryAgentSpec`

NewManagementV1alpha1DiscoveryAgentSpecWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataplaneType

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetDataplaneType() string`

GetDataplaneType returns the DataplaneType field if non-nil, zero value otherwise.

### GetDataplaneTypeOk

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetDataplaneTypeOk() (*string, bool)`

GetDataplaneTypeOk returns a tuple with the DataplaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplaneType

`func (o *ManagementV1alpha1DiscoveryAgentSpec) SetDataplaneType(v string)`

SetDataplaneType sets DataplaneType field to given value.


### GetConfig

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetConfig() ManagementV1alpha1DiscoveryAgentSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetConfigOk() (*ManagementV1alpha1DiscoveryAgentSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ManagementV1alpha1DiscoveryAgentSpec) SetConfig(v ManagementV1alpha1DiscoveryAgentSpecConfig)`

SetConfig sets Config field to given value.


### GetLogging

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetLogging() ManagementV1alpha1DiscoveryAgentSpecLogging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *ManagementV1alpha1DiscoveryAgentSpec) GetLoggingOk() (*ManagementV1alpha1DiscoveryAgentSpecLogging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *ManagementV1alpha1DiscoveryAgentSpec) SetLogging(v ManagementV1alpha1DiscoveryAgentSpecLogging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *ManagementV1alpha1DiscoveryAgentSpec) HasLogging() bool`

HasLogging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


