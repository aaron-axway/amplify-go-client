# ManagementV1alpha1DataplaneSpecAzure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**TenantId** | **string** | The tenantId is a Microsoft Entra ID entity that typically encompasses an organization | 
**ResourceGroup** | **string** | The resource group holds related resources for any Azure solution | 
**SubscriptionId** | **string** | The subscriptionId is the ID given to the subscription tied to the tenant | 
**ApimServiceName** | **string** | The name of the azure API management | 
**EventHubName** | Pointer to **string** | The event hub processes and stores events, data, or telemetry produced by distributed software or devices | [optional] 
**EventHubNamespace** | Pointer to **string** | The event hub namespace is a management container for event hubs or topics | [optional] 
**EventHubConsumerGroup** | Pointer to **string** | Consumer groups enable consuming applications to each have a separate view of the event stream. They read the stream independently at their own pace and with their own offsets. | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpecAzure

`func NewManagementV1alpha1DataplaneSpecAzure(tenantId string, resourceGroup string, subscriptionId string, apimServiceName string, ) *ManagementV1alpha1DataplaneSpecAzure`

NewManagementV1alpha1DataplaneSpecAzure instantiates a new ManagementV1alpha1DataplaneSpecAzure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecAzureWithDefaults

`func NewManagementV1alpha1DataplaneSpecAzureWithDefaults() *ManagementV1alpha1DataplaneSpecAzure`

NewManagementV1alpha1DataplaneSpecAzureWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecAzure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagementV1alpha1DataplaneSpecAzure) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTenantId

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetResourceGroup

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSubscriptionId

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetApimServiceName

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetApimServiceName() string`

GetApimServiceName returns the ApimServiceName field if non-nil, zero value otherwise.

### GetApimServiceNameOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetApimServiceNameOk() (*string, bool)`

GetApimServiceNameOk returns a tuple with the ApimServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApimServiceName

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetApimServiceName(v string)`

SetApimServiceName sets ApimServiceName field to given value.


### GetEventHubName

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubName() string`

GetEventHubName returns the EventHubName field if non-nil, zero value otherwise.

### GetEventHubNameOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNameOk() (*string, bool)`

GetEventHubNameOk returns a tuple with the EventHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubName

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubName(v string)`

SetEventHubName sets EventHubName field to given value.

### HasEventHubName

`func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubName() bool`

HasEventHubName returns a boolean if a field has been set.

### GetEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNamespace() string`

GetEventHubNamespace returns the EventHubNamespace field if non-nil, zero value otherwise.

### GetEventHubNamespaceOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubNamespaceOk() (*string, bool)`

GetEventHubNamespaceOk returns a tuple with the EventHubNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubNamespace(v string)`

SetEventHubNamespace sets EventHubNamespace field to given value.

### HasEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubNamespace() bool`

HasEventHubNamespace returns a boolean if a field has been set.

### GetEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubConsumerGroup() string`

GetEventHubConsumerGroup returns the EventHubConsumerGroup field if non-nil, zero value otherwise.

### GetEventHubConsumerGroupOk

`func (o *ManagementV1alpha1DataplaneSpecAzure) GetEventHubConsumerGroupOk() (*string, bool)`

GetEventHubConsumerGroupOk returns a tuple with the EventHubConsumerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecAzure) SetEventHubConsumerGroup(v string)`

SetEventHubConsumerGroup sets EventHubConsumerGroup field to given value.

### HasEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecAzure) HasEventHubConsumerGroup() bool`

HasEventHubConsumerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


