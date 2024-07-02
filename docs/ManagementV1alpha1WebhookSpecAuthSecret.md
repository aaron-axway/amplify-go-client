# ManagementV1alpha1WebhookSpecAuthSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Secret name to be used as a reference. If the secret is removed, the reference gets removed and the webhook invocation will be done with no Authorization header.  | [optional] 
**Key** | Pointer to **string** | Key to be used from the referenced secret. | [optional] 

## Methods

### NewManagementV1alpha1WebhookSpecAuthSecret

`func NewManagementV1alpha1WebhookSpecAuthSecret() *ManagementV1alpha1WebhookSpecAuthSecret`

NewManagementV1alpha1WebhookSpecAuthSecret instantiates a new ManagementV1alpha1WebhookSpecAuthSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1WebhookSpecAuthSecretWithDefaults

`func NewManagementV1alpha1WebhookSpecAuthSecretWithDefaults() *ManagementV1alpha1WebhookSpecAuthSecret`

NewManagementV1alpha1WebhookSpecAuthSecretWithDefaults instantiates a new ManagementV1alpha1WebhookSpecAuthSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ManagementV1alpha1WebhookSpecAuthSecret) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


