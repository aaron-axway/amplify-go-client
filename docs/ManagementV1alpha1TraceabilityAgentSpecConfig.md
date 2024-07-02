# ManagementV1alpha1TraceabilityAgentSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwningTeam** | Pointer to **string** | Name of the team that owns the catalog item created by agent. If not provided, the default team will be used. | [optional] 
**ProcessHeaders** | Pointer to **bool** | Configures the agent to include request and response headers in captured traffic. Defaults to true | [optional] 
**Redaction** | Pointer to [**ManagementV1alpha1TraceabilityAgentSpecConfigRedaction**](ManagementV1alpha1TraceabilityAgentSpecConfigRedaction.md) |  | [optional] 
**Sampling** | Pointer to [**ManagementV1alpha1TraceabilityAgentSpecConfigSampling**](ManagementV1alpha1TraceabilityAgentSpecConfigSampling.md) |  | [optional] 
**Owner** | Pointer to [**ManagementV1alpha1DiscoveryAgentSpecConfigOwner**](ManagementV1alpha1DiscoveryAgentSpecConfigOwner.md) |  | [optional] 

## Methods

### NewManagementV1alpha1TraceabilityAgentSpecConfig

`func NewManagementV1alpha1TraceabilityAgentSpecConfig() *ManagementV1alpha1TraceabilityAgentSpecConfig`

NewManagementV1alpha1TraceabilityAgentSpecConfig instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentSpecConfigWithDefaults

`func NewManagementV1alpha1TraceabilityAgentSpecConfigWithDefaults() *ManagementV1alpha1TraceabilityAgentSpecConfig`

NewManagementV1alpha1TraceabilityAgentSpecConfigWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwningTeam

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetOwningTeam() string`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetOwningTeamOk() (*string, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) SetOwningTeam(v string)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetProcessHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetProcessHeaders() bool`

GetProcessHeaders returns the ProcessHeaders field if non-nil, zero value otherwise.

### GetProcessHeadersOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetProcessHeadersOk() (*bool, bool)`

GetProcessHeadersOk returns a tuple with the ProcessHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) SetProcessHeaders(v bool)`

SetProcessHeaders sets ProcessHeaders field to given value.

### HasProcessHeaders

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) HasProcessHeaders() bool`

HasProcessHeaders returns a boolean if a field has been set.

### GetRedaction

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetRedaction() ManagementV1alpha1TraceabilityAgentSpecConfigRedaction`

GetRedaction returns the Redaction field if non-nil, zero value otherwise.

### GetRedactionOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetRedactionOk() (*ManagementV1alpha1TraceabilityAgentSpecConfigRedaction, bool)`

GetRedactionOk returns a tuple with the Redaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedaction

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) SetRedaction(v ManagementV1alpha1TraceabilityAgentSpecConfigRedaction)`

SetRedaction sets Redaction field to given value.

### HasRedaction

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) HasRedaction() bool`

HasRedaction returns a boolean if a field has been set.

### GetSampling

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetSampling() ManagementV1alpha1TraceabilityAgentSpecConfigSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetSamplingOk() (*ManagementV1alpha1TraceabilityAgentSpecConfigSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) SetSampling(v ManagementV1alpha1TraceabilityAgentSpecConfigSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetOwner() ManagementV1alpha1DiscoveryAgentSpecConfigOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) GetOwnerOk() (*ManagementV1alpha1DiscoveryAgentSpecConfigOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) SetOwner(v ManagementV1alpha1DiscoveryAgentSpecConfigOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManagementV1alpha1TraceabilityAgentSpecConfig) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


