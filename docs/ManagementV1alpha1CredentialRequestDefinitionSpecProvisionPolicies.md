# ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Renewable** | Pointer to **bool** | Defines on if Credentials using this definitions can be renewed. | [optional] 
**Suspendable** | Pointer to **bool** | Defines on if Credentials can be suspended. | [optional] 
**Expiry** | Pointer to [**ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry**](ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry.md) |  | [optional] 

## Methods

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies() *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesWithDefaults

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesWithDefaults() *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesWithDefaults instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenewable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.

### HasRenewable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.

### GetSuspendable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetSuspendable() bool`

GetSuspendable returns the Suspendable field if non-nil, zero value otherwise.

### GetSuspendableOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetSuspendableOk() (*bool, bool)`

GetSuspendableOk returns a tuple with the Suspendable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) SetSuspendable(v bool)`

SetSuspendable sets Suspendable field to given value.

### HasSuspendable

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) HasSuspendable() bool`

HasSuspendable returns a boolean if a field has been set.

### GetExpiry

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetExpiry() ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) GetExpiryOk() (*ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) SetExpiry(v ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


