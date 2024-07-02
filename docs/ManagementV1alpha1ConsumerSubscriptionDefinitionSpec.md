# ManagementV1alpha1ConsumerSubscriptionDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhooks** | Pointer to **[]string** | List of Webhook kind resource names that dictates what webhooks will be invoked on subscription changes. | [optional] 
**Schema** | Pointer to [**ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchema**](ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchema.md) |  | [optional] 

## Methods

### NewManagementV1alpha1ConsumerSubscriptionDefinitionSpec

`func NewManagementV1alpha1ConsumerSubscriptionDefinitionSpec() *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec`

NewManagementV1alpha1ConsumerSubscriptionDefinitionSpec instantiates a new ManagementV1alpha1ConsumerSubscriptionDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecWithDefaults

`func NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecWithDefaults() *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec`

NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecWithDefaults instantiates a new ManagementV1alpha1ConsumerSubscriptionDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhooks

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetSchema

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) GetSchema() ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) GetSchemaOk() (*ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) SetSchema(v ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpec) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


