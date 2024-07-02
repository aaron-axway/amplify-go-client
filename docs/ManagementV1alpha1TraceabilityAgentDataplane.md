# ManagementV1alpha1TraceabilityAgentDataplane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to **string** | Defines the interval that the dataplane will be accessed (30m &#x3D; 30 minutes, 5h 5m &#x3D; 5 hours and 5 mins, 2d &#x3D; 2 days). 30m minimum | [optional] 
**QueueTrafficCollection** | Pointer to **bool** | Queues this agent to run a traffic collection process. Defaults to false | [optional] 

## Methods

### NewManagementV1alpha1TraceabilityAgentDataplane

`func NewManagementV1alpha1TraceabilityAgentDataplane() *ManagementV1alpha1TraceabilityAgentDataplane`

NewManagementV1alpha1TraceabilityAgentDataplane instantiates a new ManagementV1alpha1TraceabilityAgentDataplane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1TraceabilityAgentDataplaneWithDefaults

`func NewManagementV1alpha1TraceabilityAgentDataplaneWithDefaults() *ManagementV1alpha1TraceabilityAgentDataplane`

NewManagementV1alpha1TraceabilityAgentDataplaneWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentDataplane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFrequency

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetQueueTrafficCollection

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetQueueTrafficCollection() bool`

GetQueueTrafficCollection returns the QueueTrafficCollection field if non-nil, zero value otherwise.

### GetQueueTrafficCollectionOk

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) GetQueueTrafficCollectionOk() (*bool, bool)`

GetQueueTrafficCollectionOk returns a tuple with the QueueTrafficCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTrafficCollection

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) SetQueueTrafficCollection(v bool)`

SetQueueTrafficCollection sets QueueTrafficCollection field to given value.

### HasQueueTrafficCollection

`func (o *ManagementV1alpha1TraceabilityAgentDataplane) HasQueueTrafficCollection() bool`

HasQueueTrafficCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


