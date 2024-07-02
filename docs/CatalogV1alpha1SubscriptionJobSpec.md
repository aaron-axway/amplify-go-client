# CatalogV1alpha1SubscriptionJobSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**CatalogV1alpha1SubscriptionJobSpecAction**](CatalogV1alpha1SubscriptionJobSpecAction.md) |  | 
**When** | Pointer to [**CatalogV1alpha1SubscriptionJobSpecWhen**](CatalogV1alpha1SubscriptionJobSpecWhen.md) |  | [optional] 
**PostExecute** | Pointer to [**CatalogV1alpha1SubscriptionJobSpecPostExecute**](CatalogV1alpha1SubscriptionJobSpecPostExecute.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionJobSpec

`func NewCatalogV1alpha1SubscriptionJobSpec(action CatalogV1alpha1SubscriptionJobSpecAction, ) *CatalogV1alpha1SubscriptionJobSpec`

NewCatalogV1alpha1SubscriptionJobSpec instantiates a new CatalogV1alpha1SubscriptionJobSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionJobSpecWithDefaults

`func NewCatalogV1alpha1SubscriptionJobSpecWithDefaults() *CatalogV1alpha1SubscriptionJobSpec`

NewCatalogV1alpha1SubscriptionJobSpecWithDefaults instantiates a new CatalogV1alpha1SubscriptionJobSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetAction() CatalogV1alpha1SubscriptionJobSpecAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetActionOk() (*CatalogV1alpha1SubscriptionJobSpecAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CatalogV1alpha1SubscriptionJobSpec) SetAction(v CatalogV1alpha1SubscriptionJobSpecAction)`

SetAction sets Action field to given value.


### GetWhen

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetWhen() CatalogV1alpha1SubscriptionJobSpecWhen`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetWhenOk() (*CatalogV1alpha1SubscriptionJobSpecWhen, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *CatalogV1alpha1SubscriptionJobSpec) SetWhen(v CatalogV1alpha1SubscriptionJobSpecWhen)`

SetWhen sets When field to given value.

### HasWhen

`func (o *CatalogV1alpha1SubscriptionJobSpec) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPostExecute

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetPostExecute() CatalogV1alpha1SubscriptionJobSpecPostExecute`

GetPostExecute returns the PostExecute field if non-nil, zero value otherwise.

### GetPostExecuteOk

`func (o *CatalogV1alpha1SubscriptionJobSpec) GetPostExecuteOk() (*CatalogV1alpha1SubscriptionJobSpecPostExecute, bool)`

GetPostExecuteOk returns a tuple with the PostExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExecute

`func (o *CatalogV1alpha1SubscriptionJobSpec) SetPostExecute(v CatalogV1alpha1SubscriptionJobSpecPostExecute)`

SetPostExecute sets PostExecute field to given value.

### HasPostExecute

`func (o *CatalogV1alpha1SubscriptionJobSpec) HasPostExecute() bool`

HasPostExecute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


