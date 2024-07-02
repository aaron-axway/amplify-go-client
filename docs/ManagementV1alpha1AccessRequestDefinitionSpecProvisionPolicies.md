# ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transferable** | Pointer to **bool** | Defines on if AccessRequests using this definition can be moved from their current Subscription to a new one. This update implies that quotas and references are updated depending on the new Subscription&#39;s plan. | [optional] 

## Methods

### NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies

`func NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies() *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies`

NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies instantiates a new ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPoliciesWithDefaults

`func NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPoliciesWithDefaults() *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies`

NewManagementV1alpha1AccessRequestDefinitionSpecProvisionPoliciesWithDefaults instantiates a new ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferable

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies) GetTransferable() bool`

GetTransferable returns the Transferable field if non-nil, zero value otherwise.

### GetTransferableOk

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies) GetTransferableOk() (*bool, bool)`

GetTransferableOk returns a tuple with the Transferable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferable

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies) SetTransferable(v bool)`

SetTransferable sets Transferable field to given value.

### HasTransferable

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies) HasTransferable() bool`

HasTransferable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


