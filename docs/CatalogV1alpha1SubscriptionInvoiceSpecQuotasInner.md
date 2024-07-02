# CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **int32** | Number of consumed units. | [optional] 
**Intervals** | Pointer to [**[]CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner**](CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner.md) | Defined for quotas with overages per specific time period. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInner

`func NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInner() *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInner instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetIntervals

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetIntervals() []CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) GetIntervalsOk() (*[]CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) SetIntervals(v []CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInner) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


