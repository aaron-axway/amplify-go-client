# CatalogV1ProductReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentResources** | Pointer to **[]string** | Computed DocumentResources used in all Product&#39;s Documents articles. | [optional] 
**Assets** | Pointer to [**[]CatalogV1ProductReferencesAssetsInner**](CatalogV1ProductReferencesAssetsInner.md) | Computed latest AssetReleases for each Asset reference in the Product. | [optional] 
**Marketplaces** | Pointer to [**[]CatalogV1ProductReferencesMarketplacesInner**](CatalogV1ProductReferencesMarketplacesInner.md) | The marketplaces this product has been published to. | [optional] 

## Methods

### NewCatalogV1ProductReferences

`func NewCatalogV1ProductReferences() *CatalogV1ProductReferences`

NewCatalogV1ProductReferences instantiates a new CatalogV1ProductReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1ProductReferencesWithDefaults

`func NewCatalogV1ProductReferencesWithDefaults() *CatalogV1ProductReferences`

NewCatalogV1ProductReferencesWithDefaults instantiates a new CatalogV1ProductReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentResources

`func (o *CatalogV1ProductReferences) GetDocumentResources() []string`

GetDocumentResources returns the DocumentResources field if non-nil, zero value otherwise.

### GetDocumentResourcesOk

`func (o *CatalogV1ProductReferences) GetDocumentResourcesOk() (*[]string, bool)`

GetDocumentResourcesOk returns a tuple with the DocumentResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentResources

`func (o *CatalogV1ProductReferences) SetDocumentResources(v []string)`

SetDocumentResources sets DocumentResources field to given value.

### HasDocumentResources

`func (o *CatalogV1ProductReferences) HasDocumentResources() bool`

HasDocumentResources returns a boolean if a field has been set.

### GetAssets

`func (o *CatalogV1ProductReferences) GetAssets() []CatalogV1ProductReferencesAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CatalogV1ProductReferences) GetAssetsOk() (*[]CatalogV1ProductReferencesAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CatalogV1ProductReferences) SetAssets(v []CatalogV1ProductReferencesAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CatalogV1ProductReferences) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetMarketplaces

`func (o *CatalogV1ProductReferences) GetMarketplaces() []CatalogV1ProductReferencesMarketplacesInner`

GetMarketplaces returns the Marketplaces field if non-nil, zero value otherwise.

### GetMarketplacesOk

`func (o *CatalogV1ProductReferences) GetMarketplacesOk() (*[]CatalogV1ProductReferencesMarketplacesInner, bool)`

GetMarketplacesOk returns a tuple with the Marketplaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaces

`func (o *CatalogV1ProductReferences) SetMarketplaces(v []CatalogV1ProductReferencesMarketplacesInner)`

SetMarketplaces sets Marketplaces field to given value.

### HasMarketplaces

`func (o *CatalogV1ProductReferences) HasMarketplaces() bool`

HasMarketplaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


