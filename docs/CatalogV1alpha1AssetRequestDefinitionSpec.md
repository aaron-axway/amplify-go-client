# CatalogV1alpha1AssetRequestDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **map[string]interface{}** | JSON Schema draft \\#7 for defining the AssetRequest properties needed to get access to an AssetResource. | 
**Provision** | Pointer to [**CatalogV1alpha1AssetRequestDefinitionSpecProvision**](CatalogV1alpha1AssetRequestDefinitionSpecProvision.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetRequestDefinitionSpec

`func NewCatalogV1alpha1AssetRequestDefinitionSpec(schema map[string]interface{}, ) *CatalogV1alpha1AssetRequestDefinitionSpec`

NewCatalogV1alpha1AssetRequestDefinitionSpec instantiates a new CatalogV1alpha1AssetRequestDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetRequestDefinitionSpecWithDefaults

`func NewCatalogV1alpha1AssetRequestDefinitionSpecWithDefaults() *CatalogV1alpha1AssetRequestDefinitionSpec`

NewCatalogV1alpha1AssetRequestDefinitionSpecWithDefaults instantiates a new CatalogV1alpha1AssetRequestDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetProvision

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) GetProvision() CatalogV1alpha1AssetRequestDefinitionSpecProvision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) GetProvisionOk() (*CatalogV1alpha1AssetRequestDefinitionSpecProvision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) SetProvision(v CatalogV1alpha1AssetRequestDefinitionSpecProvision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *CatalogV1alpha1AssetRequestDefinitionSpec) HasProvision() bool`

HasProvision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


