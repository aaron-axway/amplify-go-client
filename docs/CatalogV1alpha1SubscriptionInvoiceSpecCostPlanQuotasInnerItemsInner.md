# CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the quota invoice item. | [optional] 
**Units** | Pointer to **int32** | Number of billed units. | [optional] 
**ItemCost** | Pointer to **float64** | The cost of a billed unit for the quota or specific quota interval. | [optional] 
**From** | Pointer to **time.Time** | Start time of the invoice item. | [optional] 
**To** | Pointer to **time.Time** | End time of the invoice item. | [optional] 
**Description** | Pointer to **string** | The description of the invoice item for the quota. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner() *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInnerWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInnerWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner`

NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInnerWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetItemCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetItemCost() float64`

GetItemCost returns the ItemCost field if non-nil, zero value otherwise.

### GetItemCostOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetItemCostOk() (*float64, bool)`

GetItemCostOk returns a tuple with the ItemCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetItemCost(v float64)`

SetItemCost sets ItemCost field to given value.

### HasItemCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasItemCost() bool`

HasItemCost returns a boolean if a field has been set.

### GetFrom

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInnerItemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


