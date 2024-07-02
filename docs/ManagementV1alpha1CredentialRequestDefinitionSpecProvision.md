# ManagementV1alpha1CredentialRequestDefinitionSpecProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **map[string]interface{}** | JSON Schema draft \\#7 for describing the generated credentials format. | 
**Policies** | Pointer to [**ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies**](ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies.md) |  | [optional] 

## Methods

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvision

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvision(schema map[string]interface{}, ) *ManagementV1alpha1CredentialRequestDefinitionSpecProvision`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvision instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionWithDefaults

`func NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionWithDefaults() *ManagementV1alpha1CredentialRequestDefinitionSpecProvision`

NewManagementV1alpha1CredentialRequestDefinitionSpecProvisionWithDefaults instantiates a new ManagementV1alpha1CredentialRequestDefinitionSpecProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetPolicies

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) GetPolicies() ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) GetPoliciesOk() (*ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) SetPolicies(v ManagementV1alpha1CredentialRequestDefinitionSpecProvisionPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ManagementV1alpha1CredentialRequestDefinitionSpecProvision) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


