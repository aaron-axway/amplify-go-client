# CatalogV1alpha1ProductSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Product. | [optional] 
**SupportContact** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Assets** | Pointer to [**[]CatalogV1ProductSpecAssetsInner**](CatalogV1ProductSpecAssetsInner.md) | Defines all the Assets that the Product will be built from. | [optional] 
**AutoRelease** | Pointer to [**CatalogV1ProductSpecAutoRelease**](CatalogV1ProductSpecAutoRelease.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductSpec

`func NewCatalogV1alpha1ProductSpec() *CatalogV1alpha1ProductSpec`

NewCatalogV1alpha1ProductSpec instantiates a new CatalogV1alpha1ProductSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductSpecWithDefaults

`func NewCatalogV1alpha1ProductSpecWithDefaults() *CatalogV1alpha1ProductSpec`

NewCatalogV1alpha1ProductSpecWithDefaults instantiates a new CatalogV1alpha1ProductSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1ProductSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1ProductSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1ProductSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1ProductSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSupportContact

`func (o *CatalogV1alpha1ProductSpec) GetSupportContact() string`

GetSupportContact returns the SupportContact field if non-nil, zero value otherwise.

### GetSupportContactOk

`func (o *CatalogV1alpha1ProductSpec) GetSupportContactOk() (*string, bool)`

GetSupportContactOk returns a tuple with the SupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContact

`func (o *CatalogV1alpha1ProductSpec) SetSupportContact(v string)`

SetSupportContact sets SupportContact field to given value.

### HasSupportContact

`func (o *CatalogV1alpha1ProductSpec) HasSupportContact() bool`

HasSupportContact returns a boolean if a field has been set.

### GetCategories

`func (o *CatalogV1alpha1ProductSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CatalogV1alpha1ProductSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CatalogV1alpha1ProductSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CatalogV1alpha1ProductSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAssets

`func (o *CatalogV1alpha1ProductSpec) GetAssets() []CatalogV1ProductSpecAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CatalogV1alpha1ProductSpec) GetAssetsOk() (*[]CatalogV1ProductSpecAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CatalogV1alpha1ProductSpec) SetAssets(v []CatalogV1ProductSpecAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CatalogV1alpha1ProductSpec) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAutoRelease

`func (o *CatalogV1alpha1ProductSpec) GetAutoRelease() CatalogV1ProductSpecAutoRelease`

GetAutoRelease returns the AutoRelease field if non-nil, zero value otherwise.

### GetAutoReleaseOk

`func (o *CatalogV1alpha1ProductSpec) GetAutoReleaseOk() (*CatalogV1ProductSpecAutoRelease, bool)`

GetAutoReleaseOk returns a tuple with the AutoRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRelease

`func (o *CatalogV1alpha1ProductSpec) SetAutoRelease(v CatalogV1ProductSpecAutoRelease)`

SetAutoRelease sets AutoRelease field to given value.

### HasAutoRelease

`func (o *CatalogV1alpha1ProductSpec) HasAutoRelease() bool`

HasAutoRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


