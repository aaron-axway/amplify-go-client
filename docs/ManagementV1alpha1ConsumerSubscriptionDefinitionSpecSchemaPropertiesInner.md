# ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The Catalog Item subscription spec key. For API type subscriptions, the key needs to be &#39;profile&#39;. | 
**Value** | **map[string]interface{}** | JSON schema to specify data needed from consumers when subscriptions are created. | 

## Methods

### NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner

`func NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner(key string, value map[string]interface{}, ) *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner`

NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner instantiates a new ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInnerWithDefaults

`func NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInnerWithDefaults() *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner`

NewManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInnerWithDefaults instantiates a new ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ManagementV1alpha1ConsumerSubscriptionDefinitionSpecSchemaPropertiesInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


