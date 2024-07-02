# CatalogV1alpha1ProductPlanSpecBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Price** | Pointer to **float64** | The base price for the plan. | [optional] [default to 0]
**Cycle** | Pointer to **string** | The billing cycle type. | [optional] [default to "recurring"]
**Interval** | **string** |  | 
**Setup** | Pointer to [**CatalogV1alpha1ProductPlanSpecBillingSetup**](CatalogV1alpha1ProductPlanSpecBillingSetup.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductPlanSpecBilling

`func NewCatalogV1alpha1ProductPlanSpecBilling(currency string, interval string, ) *CatalogV1alpha1ProductPlanSpecBilling`

NewCatalogV1alpha1ProductPlanSpecBilling instantiates a new CatalogV1alpha1ProductPlanSpecBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductPlanSpecBillingWithDefaults

`func NewCatalogV1alpha1ProductPlanSpecBillingWithDefaults() *CatalogV1alpha1ProductPlanSpecBilling`

NewCatalogV1alpha1ProductPlanSpecBillingWithDefaults instantiates a new CatalogV1alpha1ProductPlanSpecBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CatalogV1alpha1ProductPlanSpecBilling) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPrice

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogV1alpha1ProductPlanSpecBilling) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogV1alpha1ProductPlanSpecBilling) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCycle

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetCycle() string`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetCycleOk() (*string, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *CatalogV1alpha1ProductPlanSpecBilling) SetCycle(v string)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *CatalogV1alpha1ProductPlanSpecBilling) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetInterval

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CatalogV1alpha1ProductPlanSpecBilling) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetSetup

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetSetup() CatalogV1alpha1ProductPlanSpecBillingSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *CatalogV1alpha1ProductPlanSpecBilling) GetSetupOk() (*CatalogV1alpha1ProductPlanSpecBillingSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *CatalogV1alpha1ProductPlanSpecBilling) SetSetup(v CatalogV1alpha1ProductPlanSpecBillingSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *CatalogV1alpha1ProductPlanSpecBilling) HasSetup() bool`

HasSetup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


