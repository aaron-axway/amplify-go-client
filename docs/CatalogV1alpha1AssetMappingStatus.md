# CatalogV1alpha1AssetMappingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**CatalogV1alpha1AssetMappingStatusSource**](CatalogV1alpha1AssetMappingStatusSource.md) |  | [optional] 
**Outputs** | Pointer to [**[]CatalogV1alpha1AssetMappingStatusOutputsInner**](CatalogV1alpha1AssetMappingStatusOutputsInner.md) | Generated catalog resources. | [optional] 

## Methods

### NewCatalogV1alpha1AssetMappingStatus

`func NewCatalogV1alpha1AssetMappingStatus() *CatalogV1alpha1AssetMappingStatus`

NewCatalogV1alpha1AssetMappingStatus instantiates a new CatalogV1alpha1AssetMappingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetMappingStatusWithDefaults

`func NewCatalogV1alpha1AssetMappingStatusWithDefaults() *CatalogV1alpha1AssetMappingStatus`

NewCatalogV1alpha1AssetMappingStatusWithDefaults instantiates a new CatalogV1alpha1AssetMappingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1AssetMappingStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1AssetMappingStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1AssetMappingStatus) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CatalogV1alpha1AssetMappingStatus) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSource

`func (o *CatalogV1alpha1AssetMappingStatus) GetSource() CatalogV1alpha1AssetMappingStatusSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CatalogV1alpha1AssetMappingStatus) GetSourceOk() (*CatalogV1alpha1AssetMappingStatusSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CatalogV1alpha1AssetMappingStatus) SetSource(v CatalogV1alpha1AssetMappingStatusSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CatalogV1alpha1AssetMappingStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOutputs

`func (o *CatalogV1alpha1AssetMappingStatus) GetOutputs() []CatalogV1alpha1AssetMappingStatusOutputsInner`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *CatalogV1alpha1AssetMappingStatus) GetOutputsOk() (*[]CatalogV1alpha1AssetMappingStatusOutputsInner, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *CatalogV1alpha1AssetMappingStatus) SetOutputs(v []CatalogV1alpha1AssetMappingStatusOutputsInner)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *CatalogV1alpha1AssetMappingStatus) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


