# CatalogV1QuotaSpecLimitTypeTiered

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "tiered"]
**Tiers** | [**[]CatalogV1QuotaSpecLimitTypeTieredTiersInner**](CatalogV1QuotaSpecLimitTypeTieredTiersInner.md) |  | 

## Methods

### NewCatalogV1QuotaSpecLimitTypeTiered

`func NewCatalogV1QuotaSpecLimitTypeTiered(type_ string, tiers []CatalogV1QuotaSpecLimitTypeTieredTiersInner, ) *CatalogV1QuotaSpecLimitTypeTiered`

NewCatalogV1QuotaSpecLimitTypeTiered instantiates a new CatalogV1QuotaSpecLimitTypeTiered object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1QuotaSpecLimitTypeTieredWithDefaults

`func NewCatalogV1QuotaSpecLimitTypeTieredWithDefaults() *CatalogV1QuotaSpecLimitTypeTiered`

NewCatalogV1QuotaSpecLimitTypeTieredWithDefaults instantiates a new CatalogV1QuotaSpecLimitTypeTiered object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1QuotaSpecLimitTypeTiered) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1QuotaSpecLimitTypeTiered) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1QuotaSpecLimitTypeTiered) SetType(v string)`

SetType sets Type field to given value.


### GetTiers

`func (o *CatalogV1QuotaSpecLimitTypeTiered) GetTiers() []CatalogV1QuotaSpecLimitTypeTieredTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CatalogV1QuotaSpecLimitTypeTiered) GetTiersOk() (*[]CatalogV1QuotaSpecLimitTypeTieredTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CatalogV1QuotaSpecLimitTypeTiered) SetTiers(v []CatalogV1QuotaSpecLimitTypeTieredTiersInner)`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


