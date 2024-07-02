# CatalogV1alpha1AssetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of the asset. | [optional] 
**Type** | **string** |  | 
**Categories** | Pointer to **[]string** | list of categories for the asset. | [optional] 
**AutoRelease** | Pointer to [**CatalogV1alpha1AssetSpecAutoRelease**](CatalogV1alpha1AssetSpecAutoRelease.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetSpec

`func NewCatalogV1alpha1AssetSpec(type_ string, ) *CatalogV1alpha1AssetSpec`

NewCatalogV1alpha1AssetSpec instantiates a new CatalogV1alpha1AssetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetSpecWithDefaults

`func NewCatalogV1alpha1AssetSpecWithDefaults() *CatalogV1alpha1AssetSpec`

NewCatalogV1alpha1AssetSpecWithDefaults instantiates a new CatalogV1alpha1AssetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1AssetSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1AssetSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1AssetSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1AssetSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CatalogV1alpha1AssetSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1AssetSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1AssetSpec) SetType(v string)`

SetType sets Type field to given value.


### GetCategories

`func (o *CatalogV1alpha1AssetSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CatalogV1alpha1AssetSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CatalogV1alpha1AssetSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CatalogV1alpha1AssetSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAutoRelease

`func (o *CatalogV1alpha1AssetSpec) GetAutoRelease() CatalogV1alpha1AssetSpecAutoRelease`

GetAutoRelease returns the AutoRelease field if non-nil, zero value otherwise.

### GetAutoReleaseOk

`func (o *CatalogV1alpha1AssetSpec) GetAutoReleaseOk() (*CatalogV1alpha1AssetSpecAutoRelease, bool)`

GetAutoReleaseOk returns a tuple with the AutoRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRelease

`func (o *CatalogV1alpha1AssetSpec) SetAutoRelease(v CatalogV1alpha1AssetSpecAutoRelease)`

SetAutoRelease sets AutoRelease field to given value.

### HasAutoRelease

`func (o *CatalogV1alpha1AssetSpec) HasAutoRelease() bool`

HasAutoRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


