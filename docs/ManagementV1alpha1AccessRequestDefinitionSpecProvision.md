# ManagementV1alpha1AccessRequestDefinitionSpecProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **map[string]interface{}** | JSON Schema draft \\#7 for describing the data to be sent back after access has been provisioned. | [optional] 
**Policies** | Pointer to [**ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies**](ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies.md) |  | [optional] 

## Methods

### NewManagementV1alpha1AccessRequestDefinitionSpecProvision

`func NewManagementV1alpha1AccessRequestDefinitionSpecProvision() *ManagementV1alpha1AccessRequestDefinitionSpecProvision`

NewManagementV1alpha1AccessRequestDefinitionSpecProvision instantiates a new ManagementV1alpha1AccessRequestDefinitionSpecProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessRequestDefinitionSpecProvisionWithDefaults

`func NewManagementV1alpha1AccessRequestDefinitionSpecProvisionWithDefaults() *ManagementV1alpha1AccessRequestDefinitionSpecProvision`

NewManagementV1alpha1AccessRequestDefinitionSpecProvisionWithDefaults instantiates a new ManagementV1alpha1AccessRequestDefinitionSpecProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetPolicies

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) GetPolicies() ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) GetPoliciesOk() (*ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) SetPolicies(v ManagementV1alpha1AccessRequestDefinitionSpecProvisionPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ManagementV1alpha1AccessRequestDefinitionSpecProvision) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


