# CatalogV1alpha1ProductPlanSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** |  | 
**Description** | Pointer to **string** | description of the Plan. | [optional] 
**Type** | **string** | The type of the Plan. | 
**Billing** | Pointer to [**CatalogV1alpha1ProductPlanSpecBilling**](CatalogV1alpha1ProductPlanSpecBilling.md) |  | [optional] 
**Features** | Pointer to [**[]CatalogV1alpha1ProductPlanSpecFeaturesInner**](CatalogV1alpha1ProductPlanSpecFeaturesInner.md) | Defines all features supported by the Plan. | [optional] 
**Subscription** | Pointer to [**CatalogV1alpha1ProductPlanSpecSubscription**](CatalogV1alpha1ProductPlanSpecSubscription.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductPlanSpec

`func NewCatalogV1alpha1ProductPlanSpec(product string, type_ string, ) *CatalogV1alpha1ProductPlanSpec`

NewCatalogV1alpha1ProductPlanSpec instantiates a new CatalogV1alpha1ProductPlanSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductPlanSpecWithDefaults

`func NewCatalogV1alpha1ProductPlanSpecWithDefaults() *CatalogV1alpha1ProductPlanSpec`

NewCatalogV1alpha1ProductPlanSpecWithDefaults instantiates a new CatalogV1alpha1ProductPlanSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *CatalogV1alpha1ProductPlanSpec) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CatalogV1alpha1ProductPlanSpec) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetDescription

`func (o *CatalogV1alpha1ProductPlanSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1ProductPlanSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1ProductPlanSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CatalogV1alpha1ProductPlanSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1ProductPlanSpec) SetType(v string)`

SetType sets Type field to given value.


### GetBilling

`func (o *CatalogV1alpha1ProductPlanSpec) GetBilling() CatalogV1alpha1ProductPlanSpecBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetBillingOk() (*CatalogV1alpha1ProductPlanSpecBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *CatalogV1alpha1ProductPlanSpec) SetBilling(v CatalogV1alpha1ProductPlanSpecBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *CatalogV1alpha1ProductPlanSpec) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetFeatures

`func (o *CatalogV1alpha1ProductPlanSpec) GetFeatures() []CatalogV1alpha1ProductPlanSpecFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetFeaturesOk() (*[]CatalogV1alpha1ProductPlanSpecFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CatalogV1alpha1ProductPlanSpec) SetFeatures(v []CatalogV1alpha1ProductPlanSpecFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CatalogV1alpha1ProductPlanSpec) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSubscription

`func (o *CatalogV1alpha1ProductPlanSpec) GetSubscription() CatalogV1alpha1ProductPlanSpecSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CatalogV1alpha1ProductPlanSpec) GetSubscriptionOk() (*CatalogV1alpha1ProductPlanSpecSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CatalogV1alpha1ProductPlanSpec) SetSubscription(v CatalogV1alpha1ProductPlanSpecSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CatalogV1alpha1ProductPlanSpec) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


