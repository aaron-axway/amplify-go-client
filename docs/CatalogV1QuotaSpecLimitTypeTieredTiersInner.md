# CatalogV1QuotaSpecLimitTypeTieredTiersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **int32** | The limit start of the tier provided. | 
**To** | **int32** | The limit end of the tier provided. | 
**Cost** | **float64** | The tier price. | 
**FlatFee** | **float64** | The flat fee for the tier. | 

## Methods

### NewCatalogV1QuotaSpecLimitTypeTieredTiersInner

`func NewCatalogV1QuotaSpecLimitTypeTieredTiersInner(from int32, to int32, cost float64, flatFee float64, ) *CatalogV1QuotaSpecLimitTypeTieredTiersInner`

NewCatalogV1QuotaSpecLimitTypeTieredTiersInner instantiates a new CatalogV1QuotaSpecLimitTypeTieredTiersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1QuotaSpecLimitTypeTieredTiersInnerWithDefaults

`func NewCatalogV1QuotaSpecLimitTypeTieredTiersInnerWithDefaults() *CatalogV1QuotaSpecLimitTypeTieredTiersInner`

NewCatalogV1QuotaSpecLimitTypeTieredTiersInnerWithDefaults instantiates a new CatalogV1QuotaSpecLimitTypeTieredTiersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) SetTo(v int32)`

SetTo sets To field to given value.


### GetCost

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) SetCost(v float64)`

SetCost sets Cost field to given value.


### GetFlatFee

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetFlatFee() float64`

GetFlatFee returns the FlatFee field if non-nil, zero value otherwise.

### GetFlatFeeOk

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) GetFlatFeeOk() (*float64, bool)`

GetFlatFeeOk returns a tuple with the FlatFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatFee

`func (o *CatalogV1QuotaSpecLimitTypeTieredTiersInner) SetFlatFee(v float64)`

SetFlatFee sets FlatFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


