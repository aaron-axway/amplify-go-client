# ManagementV1alpha1DiscoveryAgentDataplane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **string** | Defines the interval that the dataplane will be accessed (30m &#x3D; 30 minutes, 5h 5m &#x3D; 5 hours and 5 mins, 2d &#x3D; 2 days). 30m minimum | [optional] 
**QueueDiscovery** | Pointer to **bool** | Queues this agent to run a discovery process. Defaults to false | [optional] 

## Methods

### NewManagementV1alpha1DiscoveryAgentDataplane

`func NewManagementV1alpha1DiscoveryAgentDataplane() *ManagementV1alpha1DiscoveryAgentDataplane`

NewManagementV1alpha1DiscoveryAgentDataplane instantiates a new ManagementV1alpha1DiscoveryAgentDataplane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DiscoveryAgentDataplaneWithDefaults

`func NewManagementV1alpha1DiscoveryAgentDataplaneWithDefaults() *ManagementV1alpha1DiscoveryAgentDataplane`

NewManagementV1alpha1DiscoveryAgentDataplaneWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentDataplane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFrequency

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetQueueDiscovery

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetQueueDiscovery() bool`

GetQueueDiscovery returns the QueueDiscovery field if non-nil, zero value otherwise.

### GetQueueDiscoveryOk

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetQueueDiscoveryOk() (*bool, bool)`

GetQueueDiscoveryOk returns a tuple with the QueueDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDiscovery

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetQueueDiscovery(v bool)`

SetQueueDiscovery sets QueueDiscovery field to given value.

### HasQueueDiscovery

`func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasQueueDiscovery() bool`

HasQueueDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


