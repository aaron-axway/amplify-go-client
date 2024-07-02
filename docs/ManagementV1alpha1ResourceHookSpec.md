# ManagementV1alpha1ResourceHookSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | [**[]ManagementV1alpha1ResourceHookSpecTriggersInner**](ManagementV1alpha1ResourceHookSpecTriggersInner.md) |  | 
**Webhooks** | **[]string** | List of Webhook kind resource names that dictates what webhooks will be invoked on matching triggers. | 

## Methods

### NewManagementV1alpha1ResourceHookSpec

`func NewManagementV1alpha1ResourceHookSpec(triggers []ManagementV1alpha1ResourceHookSpecTriggersInner, webhooks []string, ) *ManagementV1alpha1ResourceHookSpec`

NewManagementV1alpha1ResourceHookSpec instantiates a new ManagementV1alpha1ResourceHookSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ResourceHookSpecWithDefaults

`func NewManagementV1alpha1ResourceHookSpecWithDefaults() *ManagementV1alpha1ResourceHookSpec`

NewManagementV1alpha1ResourceHookSpecWithDefaults instantiates a new ManagementV1alpha1ResourceHookSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *ManagementV1alpha1ResourceHookSpec) GetTriggers() []ManagementV1alpha1ResourceHookSpecTriggersInner`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ManagementV1alpha1ResourceHookSpec) GetTriggersOk() (*[]ManagementV1alpha1ResourceHookSpecTriggersInner, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ManagementV1alpha1ResourceHookSpec) SetTriggers(v []ManagementV1alpha1ResourceHookSpecTriggersInner)`

SetTriggers sets Triggers field to given value.


### GetWebhooks

`func (o *ManagementV1alpha1ResourceHookSpec) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *ManagementV1alpha1ResourceHookSpec) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *ManagementV1alpha1ResourceHookSpec) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


