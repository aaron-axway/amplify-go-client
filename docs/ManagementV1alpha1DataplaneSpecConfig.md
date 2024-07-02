# ManagementV1alpha1DataplaneSpecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AccessLogARN** | Pointer to **string** | The ARN of the cloud watch log resource that AWS API Gateway will be configured to send API Access data to | [optional] 
**FullTransactionLogging** | Pointer to **bool** | If true, the discovery agent will enable full transaction logging for discovered API stages | [optional] 
**ProjectId** | **string** | The Project ID on GCP that Apigee is configured in | 
**DeveloperEmail** | **string** | The Developer that will own all Apigee Applications created by the agent | 
**Mode** | Pointer to **string** | The discovery mode that the Apigee agents should use | [optional] 
**Environment** | Pointer to **string** | The environment for which the metrics are gathered | [optional] 
**MetricsFilter** | Pointer to [**ManagementV1alpha1DataplaneSpecApigeeMetricsFilter**](ManagementV1alpha1DataplaneSpecApigeeMetricsFilter.md) |  | [optional] 
**Name** | **string** | The repository name used for api specs discovery | 
**OwnerName** | **string** | The repository owner Name | 
**Filter** | Pointer to [**ManagementV1alpha1DataplaneSpecSwaggerHubFilter**](ManagementV1alpha1DataplaneSpecSwaggerHubFilter.md) |  | [optional] 
**TenantId** | **string** | The tenantId is a Microsoft Entra ID entity that typically encompasses an organization | 
**ResourceGroup** | **string** | The resource group holds related resources for any Azure solution | 
**SubscriptionId** | **string** | The subscriptionId is the ID given to the subscription tied to the tenant | 
**ApimServiceName** | **string** | The name of the azure API management | 
**EventHubName** | Pointer to **string** | The event hub processes and stores events, data, or telemetry produced by distributed software or devices | [optional] 
**EventHubNamespace** | Pointer to **string** | The event hub namespace is a management container for event hubs or topics | [optional] 
**EventHubConsumerGroup** | Pointer to **string** | Consumer groups enable consuming applications to each have a separate view of the event stream. They read the stream independently at their own pace and with their own offsets. | [optional] 
**Owner** | **string** | The owner of the organization that is used to discovery the API Specs. | 

## Methods

### NewManagementV1alpha1DataplaneSpecConfig

`func NewManagementV1alpha1DataplaneSpecConfig(type_ string, projectId string, developerEmail string, name string, ownerName string, tenantId string, resourceGroup string, subscriptionId string, apimServiceName string, owner string, ) *ManagementV1alpha1DataplaneSpecConfig`

NewManagementV1alpha1DataplaneSpecConfig instantiates a new ManagementV1alpha1DataplaneSpecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecConfigWithDefaults

`func NewManagementV1alpha1DataplaneSpecConfigWithDefaults() *ManagementV1alpha1DataplaneSpecConfig`

NewManagementV1alpha1DataplaneSpecConfigWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetType(v string)`

SetType sets Type field to given value.


### GetAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetAccessLogARN() string`

GetAccessLogARN returns the AccessLogARN field if non-nil, zero value otherwise.

### GetAccessLogARNOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetAccessLogARNOk() (*string, bool)`

GetAccessLogARNOk returns a tuple with the AccessLogARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetAccessLogARN(v string)`

SetAccessLogARN sets AccessLogARN field to given value.

### HasAccessLogARN

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasAccessLogARN() bool`

HasAccessLogARN returns a boolean if a field has been set.

### GetFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetFullTransactionLogging() bool`

GetFullTransactionLogging returns the FullTransactionLogging field if non-nil, zero value otherwise.

### GetFullTransactionLoggingOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetFullTransactionLoggingOk() (*bool, bool)`

GetFullTransactionLoggingOk returns a tuple with the FullTransactionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetFullTransactionLogging(v bool)`

SetFullTransactionLogging sets FullTransactionLogging field to given value.

### HasFullTransactionLogging

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasFullTransactionLogging() bool`

HasFullTransactionLogging returns a boolean if a field has been set.

### GetProjectId

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDeveloperEmail

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetDeveloperEmail() string`

GetDeveloperEmail returns the DeveloperEmail field if non-nil, zero value otherwise.

### GetDeveloperEmailOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetDeveloperEmailOk() (*string, bool)`

GetDeveloperEmailOk returns a tuple with the DeveloperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperEmail

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetDeveloperEmail(v string)`

SetDeveloperEmail sets DeveloperEmail field to given value.


### GetMode

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEnvironment

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetMetricsFilter() ManagementV1alpha1DataplaneSpecApigeeMetricsFilter`

GetMetricsFilter returns the MetricsFilter field if non-nil, zero value otherwise.

### GetMetricsFilterOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetMetricsFilterOk() (*ManagementV1alpha1DataplaneSpecApigeeMetricsFilter, bool)`

GetMetricsFilterOk returns a tuple with the MetricsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetMetricsFilter(v ManagementV1alpha1DataplaneSpecApigeeMetricsFilter)`

SetMetricsFilter sets MetricsFilter field to given value.

### HasMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasMetricsFilter() bool`

HasMetricsFilter returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerName

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.


### GetFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetFilter() ManagementV1alpha1DataplaneSpecSwaggerHubFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetFilterOk() (*ManagementV1alpha1DataplaneSpecSwaggerHubFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetFilter(v ManagementV1alpha1DataplaneSpecSwaggerHubFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetTenantId

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetResourceGroup

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSubscriptionId

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetApimServiceName

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetApimServiceName() string`

GetApimServiceName returns the ApimServiceName field if non-nil, zero value otherwise.

### GetApimServiceNameOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetApimServiceNameOk() (*string, bool)`

GetApimServiceNameOk returns a tuple with the ApimServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApimServiceName

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetApimServiceName(v string)`

SetApimServiceName sets ApimServiceName field to given value.


### GetEventHubName

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubName() string`

GetEventHubName returns the EventHubName field if non-nil, zero value otherwise.

### GetEventHubNameOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubNameOk() (*string, bool)`

GetEventHubNameOk returns a tuple with the EventHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubName

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetEventHubName(v string)`

SetEventHubName sets EventHubName field to given value.

### HasEventHubName

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasEventHubName() bool`

HasEventHubName returns a boolean if a field has been set.

### GetEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubNamespace() string`

GetEventHubNamespace returns the EventHubNamespace field if non-nil, zero value otherwise.

### GetEventHubNamespaceOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubNamespaceOk() (*string, bool)`

GetEventHubNamespaceOk returns a tuple with the EventHubNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetEventHubNamespace(v string)`

SetEventHubNamespace sets EventHubNamespace field to given value.

### HasEventHubNamespace

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasEventHubNamespace() bool`

HasEventHubNamespace returns a boolean if a field has been set.

### GetEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubConsumerGroup() string`

GetEventHubConsumerGroup returns the EventHubConsumerGroup field if non-nil, zero value otherwise.

### GetEventHubConsumerGroupOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetEventHubConsumerGroupOk() (*string, bool)`

GetEventHubConsumerGroupOk returns a tuple with the EventHubConsumerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetEventHubConsumerGroup(v string)`

SetEventHubConsumerGroup sets EventHubConsumerGroup field to given value.

### HasEventHubConsumerGroup

`func (o *ManagementV1alpha1DataplaneSpecConfig) HasEventHubConsumerGroup() bool`

HasEventHubConsumerGroup returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1DataplaneSpecConfig) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1DataplaneSpecConfig) SetOwner(v string)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


