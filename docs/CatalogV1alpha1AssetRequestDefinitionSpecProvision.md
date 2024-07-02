# CatalogV1alpha1AssetRequestDefinitionSpecProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **map[string]interface{}** | JSON Schema draft \\#7 for describing the data to be sent back after access has been provisioned. | [optional] 
**Policies** | Pointer to [**CatalogV1alpha1AssetRequestDefinitionSpecProvisionPolicies**](CatalogV1alpha1AssetRequestDefinitionSpecProvisionPolicies.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetRequestDefinitionSpecProvision

`func NewCatalogV1alpha1AssetRequestDefinitionSpecProvision() *CatalogV1alpha1AssetRequestDefinitionSpecProvision`

NewCatalogV1alpha1AssetRequestDefinitionSpecProvision instantiates a new CatalogV1alpha1AssetRequestDefinitionSpecProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetRequestDefinitionSpecProvisionWithDefaults

`func NewCatalogV1alpha1AssetRequestDefinitionSpecProvisionWithDefaults() *CatalogV1alpha1AssetRequestDefinitionSpecProvision`

NewCatalogV1alpha1AssetRequestDefinitionSpecProvisionWithDefaults instantiates a new CatalogV1alpha1AssetRequestDefinitionSpecProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetPolicies

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) GetPolicies() CatalogV1alpha1AssetRequestDefinitionSpecProvisionPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) GetPoliciesOk() (*CatalogV1alpha1AssetRequestDefinitionSpecProvisionPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) SetPolicies(v CatalogV1alpha1AssetRequestDefinitionSpecProvisionPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CatalogV1alpha1AssetRequestDefinitionSpecProvision) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


