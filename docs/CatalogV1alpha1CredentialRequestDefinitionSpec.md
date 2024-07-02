# CatalogV1alpha1CredentialRequestDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **map[string]interface{}** | JSON Schema draft \\#7 for describing the fields needed to provision Credentials of that type. | 
**Provision** | [**ManagementV1alpha1CredentialRequestDefinitionSpecProvision**](ManagementV1alpha1CredentialRequestDefinitionSpecProvision.md) |  | 

## Methods

### NewCatalogV1alpha1CredentialRequestDefinitionSpec

`func NewCatalogV1alpha1CredentialRequestDefinitionSpec(schema map[string]interface{}, provision ManagementV1alpha1CredentialRequestDefinitionSpecProvision, ) *CatalogV1alpha1CredentialRequestDefinitionSpec`

NewCatalogV1alpha1CredentialRequestDefinitionSpec instantiates a new CatalogV1alpha1CredentialRequestDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1CredentialRequestDefinitionSpecWithDefaults

`func NewCatalogV1alpha1CredentialRequestDefinitionSpecWithDefaults() *CatalogV1alpha1CredentialRequestDefinitionSpec`

NewCatalogV1alpha1CredentialRequestDefinitionSpecWithDefaults instantiates a new CatalogV1alpha1CredentialRequestDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetProvision

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetProvision() ManagementV1alpha1CredentialRequestDefinitionSpecProvision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) GetProvisionOk() (*ManagementV1alpha1CredentialRequestDefinitionSpecProvision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *CatalogV1alpha1CredentialRequestDefinitionSpec) SetProvision(v ManagementV1alpha1CredentialRequestDefinitionSpecProvision)`

SetProvision sets Provision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


