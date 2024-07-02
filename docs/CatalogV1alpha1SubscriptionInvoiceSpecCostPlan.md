# CatalogV1alpha1SubscriptionInvoiceSpecCostPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Product Plan to which to migrate the Subscription. | [optional] 
**Cost** | Pointer to **float64** | The cost of the plan when the invoice was generated. | [optional] 
**Setup** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceSpecCostPlanSetup**](CatalogV1alpha1SubscriptionInvoiceSpecCostPlanSetup.md) |  | [optional] 
**Quotas** | Pointer to [**[]CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInner**](CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInner.md) | Quotas for which there is cost associated. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlan

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlan() *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan`

NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlan instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCostPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanWithDefaults() *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan`

NewCatalogV1alpha1SubscriptionInvoiceSpecCostPlanWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoiceSpecCostPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSetup

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetSetup() CatalogV1alpha1SubscriptionInvoiceSpecCostPlanSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetSetupOk() (*CatalogV1alpha1SubscriptionInvoiceSpecCostPlanSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) SetSetup(v CatalogV1alpha1SubscriptionInvoiceSpecCostPlanSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetQuotas() []CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInner`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) GetQuotasOk() (*[]CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInner, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) SetQuotas(v []CatalogV1alpha1SubscriptionInvoiceSpecCostPlanQuotasInner)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *CatalogV1alpha1SubscriptionInvoiceSpecCostPlan) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


