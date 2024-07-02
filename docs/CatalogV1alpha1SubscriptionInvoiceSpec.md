# CatalogV1alpha1SubscriptionInvoiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceSpecCost**](CatalogV1alpha1SubscriptionInvoiceSpecCost.md) |  | [optional] 
**Period** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceSpecPeriod**](CatalogV1alpha1SubscriptionInvoiceSpecPeriod.md) |  | [optional] 
**Quotas** | Pointer to [**[]CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner**](CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpec

`func NewCatalogV1alpha1SubscriptionInvoiceSpec() *CatalogV1alpha1SubscriptionInvoiceSpec`

NewCatalogV1alpha1SubscriptionInvoiceSpec instantiates a new CatalogV1alpha1SubscriptionInvoiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpec`

NewCatalogV1alpha1SubscriptionInvoiceSpecWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetCost() CatalogV1alpha1SubscriptionInvoiceSpecCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetCostOk() (*CatalogV1alpha1SubscriptionInvoiceSpecCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) SetCost(v CatalogV1alpha1SubscriptionInvoiceSpecCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPeriod

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetPeriod() CatalogV1alpha1SubscriptionInvoiceSpecPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetPeriodOk() (*CatalogV1alpha1SubscriptionInvoiceSpecPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) SetPeriod(v CatalogV1alpha1SubscriptionInvoiceSpecPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetQuotas() []CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) GetQuotasOk() (*[]CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) SetQuotas(v []CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpec) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


