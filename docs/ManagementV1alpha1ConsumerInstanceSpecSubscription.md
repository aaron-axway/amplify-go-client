# ManagementV1alpha1ConsumerInstanceSpecSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Defines if subscriptions are allowed on the Catalog Item. | [optional] 
**AutoSubscribe** | Pointer to **bool** | Defines if subscriptions need to be manually approved. | [optional] 
**SubscriptionDefinition** | Pointer to **string** | The name of a ConsumerSubscriptionDefinition kind that defines the schema and possible webhooks to get invoked. | [optional] 

## Methods

### NewManagementV1alpha1ConsumerInstanceSpecSubscription

`func NewManagementV1alpha1ConsumerInstanceSpecSubscription() *ManagementV1alpha1ConsumerInstanceSpecSubscription`

NewManagementV1alpha1ConsumerInstanceSpecSubscription instantiates a new ManagementV1alpha1ConsumerInstanceSpecSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerInstanceSpecSubscriptionWithDefaults

`func NewManagementV1alpha1ConsumerInstanceSpecSubscriptionWithDefaults() *ManagementV1alpha1ConsumerInstanceSpecSubscription`

NewManagementV1alpha1ConsumerInstanceSpecSubscriptionWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceSpecSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoSubscribe

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetAutoSubscribe() bool`

GetAutoSubscribe returns the AutoSubscribe field if non-nil, zero value otherwise.

### GetAutoSubscribeOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetAutoSubscribeOk() (*bool, bool)`

GetAutoSubscribeOk returns a tuple with the AutoSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscribe

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) SetAutoSubscribe(v bool)`

SetAutoSubscribe sets AutoSubscribe field to given value.

### HasAutoSubscribe

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) HasAutoSubscribe() bool`

HasAutoSubscribe returns a boolean if a field has been set.

### GetSubscriptionDefinition

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetSubscriptionDefinition() string`

GetSubscriptionDefinition returns the SubscriptionDefinition field if non-nil, zero value otherwise.

### GetSubscriptionDefinitionOk

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) GetSubscriptionDefinitionOk() (*string, bool)`

GetSubscriptionDefinitionOk returns a tuple with the SubscriptionDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDefinition

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) SetSubscriptionDefinition(v string)`

SetSubscriptionDefinition sets SubscriptionDefinition field to given value.

### HasSubscriptionDefinition

`func (o *ManagementV1alpha1ConsumerInstanceSpecSubscription) HasSubscriptionDefinition() bool`

HasSubscriptionDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


