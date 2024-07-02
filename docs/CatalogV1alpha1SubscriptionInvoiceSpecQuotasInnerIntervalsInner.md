# CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **time.Time** | The start of the interval. | 
**Units** | **int32** | Number of consumed units in the interval. | 
**PreviousInvoiceItem** | Pointer to **bool** | In case the item is from a prior invoice for the same quota. Prior invoice items are included if the first item is not for the complete quota interval. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner

`func NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner(from time.Time, units int32, ) *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInnerWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInnerWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInnerWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) SetUnits(v int32)`

SetUnits sets Units field to given value.


### GetPreviousInvoiceItem

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetPreviousInvoiceItem() bool`

GetPreviousInvoiceItem returns the PreviousInvoiceItem field if non-nil, zero value otherwise.

### GetPreviousInvoiceItemOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) GetPreviousInvoiceItemOk() (*bool, bool)`

GetPreviousInvoiceItemOk returns a tuple with the PreviousInvoiceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousInvoiceItem

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) SetPreviousInvoiceItem(v bool)`

SetPreviousInvoiceItem sets PreviousInvoiceItem field to given value.

### HasPreviousInvoiceItem

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecQuotasInnerIntervalsInner) HasPreviousInvoiceItem() bool`

HasPreviousInvoiceItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


