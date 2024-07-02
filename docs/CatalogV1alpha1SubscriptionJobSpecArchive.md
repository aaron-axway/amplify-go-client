# CatalogV1alpha1SubscriptionJobSpecArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Reason** | Pointer to **string** | Description on why the subscription was archived. | [optional] 
**ProductPlanJob** | Pointer to **string** | The Product Plan Job which triggered the archival. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionJobSpecArchive

`func NewCatalogV1alpha1SubscriptionJobSpecArchive(type_ string, ) *CatalogV1alpha1SubscriptionJobSpecArchive`

NewCatalogV1alpha1SubscriptionJobSpecArchive instantiates a new CatalogV1alpha1SubscriptionJobSpecArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionJobSpecArchiveWithDefaults

`func NewCatalogV1alpha1SubscriptionJobSpecArchiveWithDefaults() *CatalogV1alpha1SubscriptionJobSpecArchive`

NewCatalogV1alpha1SubscriptionJobSpecArchiveWithDefaults instantiates a new CatalogV1alpha1SubscriptionJobSpecArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetProductPlanJob() string`

GetProductPlanJob returns the ProductPlanJob field if non-nil, zero value otherwise.

### GetProductPlanJobOk

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) GetProductPlanJobOk() (*string, bool)`

GetProductPlanJobOk returns a tuple with the ProductPlanJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) SetProductPlanJob(v string)`

SetProductPlanJob sets ProductPlanJob field to given value.

### HasProductPlanJob

`func (o *CatalogV1alpha1SubscriptionJobSpecArchive) HasProductPlanJob() bool`

HasProductPlanJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


