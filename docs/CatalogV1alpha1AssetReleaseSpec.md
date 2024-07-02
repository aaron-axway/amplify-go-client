# CatalogV1alpha1AssetReleaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the asset release. | [optional] 
**Type** | **string** |  | 
**Version** | **string** | version of the asset release. | 
**Asset** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** | list of categories for the released asset. | [optional] 

## Methods

### NewCatalogV1alpha1AssetReleaseSpec

`func NewCatalogV1alpha1AssetReleaseSpec(type_ string, version string, asset string, ) *CatalogV1alpha1AssetReleaseSpec`

NewCatalogV1alpha1AssetReleaseSpec instantiates a new CatalogV1alpha1AssetReleaseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetReleaseSpecWithDefaults

`func NewCatalogV1alpha1AssetReleaseSpecWithDefaults() *CatalogV1alpha1AssetReleaseSpec`

NewCatalogV1alpha1AssetReleaseSpecWithDefaults instantiates a new CatalogV1alpha1AssetReleaseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1AssetReleaseSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1AssetReleaseSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1AssetReleaseSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CatalogV1alpha1AssetReleaseSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1AssetReleaseSpec) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *CatalogV1alpha1AssetReleaseSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1alpha1AssetReleaseSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetAsset

`func (o *CatalogV1alpha1AssetReleaseSpec) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CatalogV1alpha1AssetReleaseSpec) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetState

`func (o *CatalogV1alpha1AssetReleaseSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1alpha1AssetReleaseSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CatalogV1alpha1AssetReleaseSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCategories

`func (o *CatalogV1alpha1AssetReleaseSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CatalogV1alpha1AssetReleaseSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CatalogV1alpha1AssetReleaseSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CatalogV1alpha1AssetReleaseSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


