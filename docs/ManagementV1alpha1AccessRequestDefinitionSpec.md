# ManagementV1alpha1AccessRequestDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **map[string]interface{}** | JSON Schema draft \\#7 for defining the AccessRequest properties needed to get access to an APIServiceInstance. | 
**Provision** | Pointer to [**ManagementV1alpha1AccessRequestDefinitionSpecProvision**](ManagementV1alpha1AccessRequestDefinitionSpecProvision.md) |  | [optional] 

## Methods

### NewManagementV1alpha1AccessRequestDefinitionSpec

`func NewManagementV1alpha1AccessRequestDefinitionSpec(schema map[string]interface{}, ) *ManagementV1alpha1AccessRequestDefinitionSpec`

NewManagementV1alpha1AccessRequestDefinitionSpec instantiates a new ManagementV1alpha1AccessRequestDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessRequestDefinitionSpecWithDefaults

`func NewManagementV1alpha1AccessRequestDefinitionSpecWithDefaults() *ManagementV1alpha1AccessRequestDefinitionSpec`

NewManagementV1alpha1AccessRequestDefinitionSpecWithDefaults instantiates a new ManagementV1alpha1AccessRequestDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetProvision

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) GetProvision() ManagementV1alpha1AccessRequestDefinitionSpecProvision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) GetProvisionOk() (*ManagementV1alpha1AccessRequestDefinitionSpecProvision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) SetProvision(v ManagementV1alpha1AccessRequestDefinitionSpecProvision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *ManagementV1alpha1AccessRequestDefinitionSpec) HasProvision() bool`

HasProvision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


