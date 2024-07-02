# CatalogV1alpha1QuotaSpecPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Interval** | **string** |  | 
**Limit** | [**CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit**](CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit.md) |  | 
**Cost** | **float64** | The cost per unit of usage. | 

## Methods

### NewCatalogV1alpha1QuotaSpecPricing

`func NewCatalogV1alpha1QuotaSpecPricing(type_ string, interval string, limit CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit, cost float64, ) *CatalogV1alpha1QuotaSpecPricing`

NewCatalogV1alpha1QuotaSpecPricing instantiates a new CatalogV1alpha1QuotaSpecPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1QuotaSpecPricingWithDefaults

`func NewCatalogV1alpha1QuotaSpecPricingWithDefaults() *CatalogV1alpha1QuotaSpecPricing`

NewCatalogV1alpha1QuotaSpecPricingWithDefaults instantiates a new CatalogV1alpha1QuotaSpecPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1QuotaSpecPricing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1QuotaSpecPricing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1QuotaSpecPricing) SetType(v string)`

SetType sets Type field to given value.


### GetInterval

`func (o *CatalogV1alpha1QuotaSpecPricing) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CatalogV1alpha1QuotaSpecPricing) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CatalogV1alpha1QuotaSpecPricing) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetLimit

`func (o *CatalogV1alpha1QuotaSpecPricing) GetLimit() CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CatalogV1alpha1QuotaSpecPricing) GetLimitOk() (*CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CatalogV1alpha1QuotaSpecPricing) SetLimit(v CatalogV1alpha1QuotaSpecGraduatedPricingTypeLimit)`

SetLimit sets Limit field to given value.


### GetCost

`func (o *CatalogV1alpha1QuotaSpecPricing) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogV1alpha1QuotaSpecPricing) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogV1alpha1QuotaSpecPricing) SetCost(v float64)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


