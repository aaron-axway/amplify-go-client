# ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **int32** | The number of days after the Credentials are considered to be expired. | 
**Actions** | Pointer to [**[]ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryActionsInner**](ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryActionsInner.md) | The actions taken when the Credentials expire. | [optional] 

## Methods

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry(period int32, ) *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryWithDefaults

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryWithDefaults() *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryWithDefaults instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetActions

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) GetActions() []ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) GetActionsOk() (*[]ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) SetActions(v []ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiryActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPoliciesExpiry) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


