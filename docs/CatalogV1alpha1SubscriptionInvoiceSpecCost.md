# CatalogV1alpha1SubscriptionInvoiceSpecCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float64** | The computed cost of the entire invoice, including plan and quota items costs. | [optional] 
**Plan** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceSpecCostPlan**](CatalogV1alpha1SubscriptionInvoiceSpecCostPlan.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpecCost

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCost() *CatalogV1alpha1SubscriptionInvoiceSpecCost`

NewCatalogV1alpha1SubscriptionInvoiceSpecCost instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecCostWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCostWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpecCost`

NewCatalogV1alpha1SubscriptionInvoiceSpecCostWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPlan

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) GetPlan() CatalogV1alpha1SubscriptionInvoiceSpecCostPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) GetPlanOk() (*CatalogV1alpha1SubscriptionInvoiceSpecCostPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) SetPlan(v CatalogV1alpha1SubscriptionInvoiceSpecCostPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCost) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


