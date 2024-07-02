# ManagementV1alpha1AssetMappingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMappingTemplate** | Pointer to **string** | Reference to the executed AssetMappingTemplate. | [optional] 
**Inputs** | [**CatalogV1alpha1AssetMappingSpecInputs**](CatalogV1alpha1AssetMappingSpecInputs.md) |  | 

## Methods

### NewManagementV1alpha1AssetMappingSpec

`func NewManagementV1alpha1AssetMappingSpec(inputs CatalogV1alpha1AssetMappingSpecInputs, ) *ManagementV1alpha1AssetMappingSpec`

NewManagementV1alpha1AssetMappingSpec instantiates a new ManagementV1alpha1AssetMappingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AssetMappingSpecWithDefaults

`func NewManagementV1alpha1AssetMappingSpecWithDefaults() *ManagementV1alpha1AssetMappingSpec`

NewManagementV1alpha1AssetMappingSpecWithDefaults instantiates a new ManagementV1alpha1AssetMappingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMappingTemplate

`func (o *ManagementV1alpha1AssetMappingSpec) GetAssetMappingTemplate() string`

GetAssetMappingTemplate returns the AssetMappingTemplate field if non-nil, zero value otherwise.

### GetAssetMappingTemplateOk

`func (o *ManagementV1alpha1AssetMappingSpec) GetAssetMappingTemplateOk() (*string, bool)`

GetAssetMappingTemplateOk returns a tuple with the AssetMappingTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMappingTemplate

`func (o *ManagementV1alpha1AssetMappingSpec) SetAssetMappingTemplate(v string)`

SetAssetMappingTemplate sets AssetMappingTemplate field to given value.

### HasAssetMappingTemplate

`func (o *ManagementV1alpha1AssetMappingSpec) HasAssetMappingTemplate() bool`

HasAssetMappingTemplate returns a boolean if a field has been set.

### GetInputs

`func (o *ManagementV1alpha1AssetMappingSpec) GetInputs() CatalogV1alpha1AssetMappingSpecInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ManagementV1alpha1AssetMappingSpec) GetInputsOk() (*CatalogV1alpha1AssetMappingSpecInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ManagementV1alpha1AssetMappingSpec) SetInputs(v CatalogV1alpha1AssetMappingSpecInputs)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


