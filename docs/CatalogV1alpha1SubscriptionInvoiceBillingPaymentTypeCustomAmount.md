# CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency for the billed amount. | 
**Total** | **float64** | Total amount after discounts and taxes. | 
**Due** | **float64** | Final amount due at this time for this invoice. | 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount(currency string, total float64, due float64, ) *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmountWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmountWithDefaults() *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount`

NewCatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmountWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTotal

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) SetTotal(v float64)`

SetTotal sets Total field to given value.


### GetDue

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetDue() float64`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) GetDueOk() (*float64, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *CatalogV1alpha1SubscriptionInvoiceBillingPaymentTypeCustomAmount) SetDue(v float64)`

SetDue sets Due field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


