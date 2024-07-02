# CatalogV1alpha1SubscriptionJobSpecPlanMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ProductPlan** | **string** | The Product Plan to which to migrate the Subscription. | 
**Reason** | Pointer to **string** | Description on why the subscription was migrated. | [optional] 
**ProductPlanJob** | Pointer to **string** | The Product Plan Job which triggered the migration. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionJobSpecPlanMigration

`func NewCatalogV1alpha1SubscriptionJobSpecPlanMigration(type_ string, productPlan string, ) *CatalogV1alpha1SubscriptionJobSpecPlanMigration`

NewCatalogV1alpha1SubscriptionJobSpecPlanMigration instantiates a new CatalogV1alpha1SubscriptionJobSpecPlanMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionJobSpecPlanMigrationWithDefaults

`func NewCatalogV1alpha1SubscriptionJobSpecPlanMigrationWithDefaults() *CatalogV1alpha1SubscriptionJobSpecPlanMigration`

NewCatalogV1alpha1SubscriptionJobSpecPlanMigrationWithDefaults instantiates a new CatalogV1alpha1SubscriptionJobSpecPlanMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) SetType(v string)`

SetType sets Type field to given value.


### GetProductPlan

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetProductPlan() string`

GetProductPlan returns the ProductPlan field if non-nil, zero value otherwise.

### GetProductPlanOk

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetProductPlanOk() (*string, bool)`

GetProductPlanOk returns a tuple with the ProductPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlan

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) SetProductPlan(v string)`

SetProductPlan sets ProductPlan field to given value.


### GetReason

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetProductPlanJob() string`

GetProductPlanJob returns the ProductPlanJob field if non-nil, zero value otherwise.

### GetProductPlanJobOk

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) GetProductPlanJobOk() (*string, bool)`

GetProductPlanJobOk returns a tuple with the ProductPlanJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) SetProductPlanJob(v string)`

SetProductPlanJob sets ProductPlanJob field to given value.

### HasProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecPlanMigration) HasProductPlanJob() bool`

HasProductPlanJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


