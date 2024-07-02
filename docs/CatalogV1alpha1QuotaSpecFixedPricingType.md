# CatalogV1alpha1QuotaSpecFixedPricingType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Interval** | **string** |  | 
**Limit** | [**CatalogV1alpha1QuotaSpecFixedPricingTypeLimit**](CatalogV1alpha1QuotaSpecFixedPricingTypeLimit.md) |  | 

## Methods

### NewCatalogV1alpha1QuotaSpecFixedPricingType

`func NewCatalogV1alpha1QuotaSpecFixedPricingType(type_ string, interval string, limit CatalogV1alpha1QuotaSpecFixedPricingTypeLimit, ) *CatalogV1alpha1QuotaSpecFixedPricingType`

NewCatalogV1alpha1QuotaSpecFixedPricingType instantiates a new CatalogV1alpha1QuotaSpecFixedPricingType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1QuotaSpecFixedPricingTypeWithDefaults

`func NewCatalogV1alpha1QuotaSpecFixedPricingTypeWithDefaults() *CatalogV1alpha1QuotaSpecFixedPricingType`

NewCatalogV1alpha1QuotaSpecFixedPricingTypeWithDefaults instantiates a new CatalogV1alpha1QuotaSpecFixedPricingType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) SetType(v string)`

SetType sets Type field to given value.


### GetInterval

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetLimit

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetLimit() CatalogV1alpha1QuotaSpecFixedPricingTypeLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) GetLimitOk() (*CatalogV1alpha1QuotaSpecFixedPricingTypeLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CatalogV1alpha1QuotaSpecFixedPricingType) SetLimit(v CatalogV1alpha1QuotaSpecFixedPricingTypeLimit)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


