# CatalogV1DocumentResourceReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceSettings** | Pointer to [**[]CatalogV1DocumentResourceReferencesMarketplaceSettingsInner**](CatalogV1DocumentResourceReferencesMarketplaceSettingsInner.md) | The marketplaces this DocumentResource is being used in as part of the marketplace settings. | [optional] 
**PlatformSettings** | Pointer to [**CatalogV1DocumentResourceReferencesPlatformSettings**](CatalogV1DocumentResourceReferencesPlatformSettings.md) |  | [optional] 

## Methods

### NewCatalogV1DocumentResourceReferences

`func NewCatalogV1DocumentResourceReferences() *CatalogV1DocumentResourceReferences`

NewCatalogV1DocumentResourceReferences instantiates a new CatalogV1DocumentResourceReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentResourceReferencesWithDefaults

`func NewCatalogV1DocumentResourceReferencesWithDefaults() *CatalogV1DocumentResourceReferences`

NewCatalogV1DocumentResourceReferencesWithDefaults instantiates a new CatalogV1DocumentResourceReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceSettings

`func (o *CatalogV1DocumentResourceReferences) GetMarketplaceSettings() []CatalogV1DocumentResourceReferencesMarketplaceSettingsInner`

GetMarketplaceSettings returns the MarketplaceSettings field if non-nil, zero value otherwise.

### GetMarketplaceSettingsOk

`func (o *CatalogV1DocumentResourceReferences) GetMarketplaceSettingsOk() (*[]CatalogV1DocumentResourceReferencesMarketplaceSettingsInner, bool)`

GetMarketplaceSettingsOk returns a tuple with the MarketplaceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceSettings

`func (o *CatalogV1DocumentResourceReferences) SetMarketplaceSettings(v []CatalogV1DocumentResourceReferencesMarketplaceSettingsInner)`

SetMarketplaceSettings sets MarketplaceSettings field to given value.

### HasMarketplaceSettings

`func (o *CatalogV1DocumentResourceReferences) HasMarketplaceSettings() bool`

HasMarketplaceSettings returns a boolean if a field has been set.

### GetPlatformSettings

`func (o *CatalogV1DocumentResourceReferences) GetPlatformSettings() CatalogV1DocumentResourceReferencesPlatformSettings`

GetPlatformSettings returns the PlatformSettings field if non-nil, zero value otherwise.

### GetPlatformSettingsOk

`func (o *CatalogV1DocumentResourceReferences) GetPlatformSettingsOk() (*CatalogV1DocumentResourceReferencesPlatformSettings, bool)`

GetPlatformSettingsOk returns a tuple with the PlatformSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSettings

`func (o *CatalogV1DocumentResourceReferences) SetPlatformSettings(v CatalogV1DocumentResourceReferencesPlatformSettings)`

SetPlatformSettings sets PlatformSettings field to given value.

### HasPlatformSettings

`func (o *CatalogV1DocumentResourceReferences) HasPlatformSettings() bool`

HasPlatformSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


