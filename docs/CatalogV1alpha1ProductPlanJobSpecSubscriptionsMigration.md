# CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ProductPlan** | **string** | The Product Plan to which to migrate the existing Product Subscriptions. | 
**Reason** | Pointer to **string** | Description on why the subscriptions was migrated. | [optional] 
**Filters** | Pointer to [**[]CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationFiltersInner**](CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationFiltersInner.md) | Filter the subscriptions that are wanted to be migrated. | [optional] 

## Methods

### NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration

`func NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration(type_ string, productPlan string, ) *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration`

NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration instantiates a new CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationWithDefaults

`func NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationWithDefaults() *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration`

NewCatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationWithDefaults instantiates a new CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) SetType(v string)`

SetType sets Type field to given value.


### GetProductPlan

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetProductPlan() string`

GetProductPlan returns the ProductPlan field if non-nil, zero value otherwise.

### GetProductPlanOk

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetProductPlanOk() (*string, bool)`

GetProductPlanOk returns a tuple with the ProductPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlan

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) SetProductPlan(v string)`

SetProductPlan sets ProductPlan field to given value.


### GetReason

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetFilters

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetFilters() []CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) GetFiltersOk() (*[]CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) SetFilters(v []CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigrationFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CatalogV1alpha1ProductPlanJobSpecSubscriptionsMigration) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


