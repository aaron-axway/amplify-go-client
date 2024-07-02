# CatalogV1alpha1ProductPlanJobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**CatalogV1alpha1ProductPlanJobSpecAction**](CatalogV1alpha1ProductPlanJobSpecAction.md) |  | 
**When** | Pointer to [**CatalogV1alpha1SubscriptionJobSpecWhen**](CatalogV1alpha1SubscriptionJobSpecWhen.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductPlanJobSpec

`func NewCatalogV1alpha1ProductPlanJobSpec(action CatalogV1alpha1ProductPlanJobSpecAction, ) *CatalogV1alpha1ProductPlanJobSpec`

NewCatalogV1alpha1ProductPlanJobSpec instantiates a new CatalogV1alpha1ProductPlanJobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductPlanJobSpecWithDefaults

`func NewCatalogV1alpha1ProductPlanJobSpecWithDefaults() *CatalogV1alpha1ProductPlanJobSpec`

NewCatalogV1alpha1ProductPlanJobSpecWithDefaults instantiates a new CatalogV1alpha1ProductPlanJobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CatalogV1alpha1ProductPlanJobSpec) GetAction() CatalogV1alpha1ProductPlanJobSpecAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CatalogV1alpha1ProductPlanJobSpec) GetActionOk() (*CatalogV1alpha1ProductPlanJobSpecAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CatalogV1alpha1ProductPlanJobSpec) SetAction(v CatalogV1alpha1ProductPlanJobSpecAction)`

SetAction sets Action field to given value.


### GetWhen

`func (o *CatalogV1alpha1ProductPlanJobSpec) GetWhen() CatalogV1alpha1SubscriptionJobSpecWhen`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *CatalogV1alpha1ProductPlanJobSpec) GetWhenOk() (*CatalogV1alpha1SubscriptionJobSpecWhen, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *CatalogV1alpha1ProductPlanJobSpec) SetWhen(v CatalogV1alpha1SubscriptionJobSpecWhen)`

SetWhen sets When field to given value.

### HasWhen

`func (o *CatalogV1alpha1ProductPlanJobSpec) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


