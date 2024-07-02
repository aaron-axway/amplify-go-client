# CatalogV1alpha1ProductPlanSpecSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | Pointer to **string** | Defines properties required from a consumer to subscribe to the plan. | [optional] 
**Interval** | Pointer to [**CatalogV1alpha1ProductPlanSpecSubscriptionInterval**](CatalogV1alpha1ProductPlanSpecSubscriptionInterval.md) |  | [optional] 
**Renewal** | Pointer to **string** |  | [optional] 
**Approval** | Pointer to **string** |  | [optional] 
**Cycles** | Pointer to **int32** | Optional number of cycles after which the subscription will be archived. Cycles start once the subscription has been approved. | [optional] 

## Methods

### NewCatalogV1alpha1ProductPlanSpecSubscription

`func NewCatalogV1alpha1ProductPlanSpecSubscription() *CatalogV1alpha1ProductPlanSpecSubscription`

NewCatalogV1alpha1ProductPlanSpecSubscription instantiates a new CatalogV1alpha1ProductPlanSpecSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductPlanSpecSubscriptionWithDefaults

`func NewCatalogV1alpha1ProductPlanSpecSubscriptionWithDefaults() *CatalogV1alpha1ProductPlanSpecSubscription`

NewCatalogV1alpha1ProductPlanSpecSubscriptionWithDefaults instantiates a new CatalogV1alpha1ProductPlanSpecSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetDefinition(v string)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetInterval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetInterval() CatalogV1alpha1ProductPlanSpecSubscriptionInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetIntervalOk() (*CatalogV1alpha1ProductPlanSpecSubscriptionInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetInterval(v CatalogV1alpha1ProductPlanSpecSubscriptionInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRenewal

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetRenewal() string`

GetRenewal returns the Renewal field if non-nil, zero value otherwise.

### GetRenewalOk

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetRenewalOk() (*string, bool)`

GetRenewalOk returns a tuple with the Renewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewal

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetRenewal(v string)`

SetRenewal sets Renewal field to given value.

### HasRenewal

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasRenewal() bool`

HasRenewal returns a boolean if a field has been set.

### GetApproval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetApproval() string`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetApprovalOk() (*string, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetApproval(v string)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetCycles

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetCycles() int32`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) GetCyclesOk() (*int32, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) SetCycles(v int32)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *CatalogV1alpha1ProductPlanSpecSubscription) HasCycles() bool`

HasCycles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


