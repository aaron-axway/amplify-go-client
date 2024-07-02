# CatalogV1alpha1SubscriptionJobSpecPostExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnSuccess** | Pointer to [**[]CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInner**](CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInner.md) | Actions to be executed after the new Subscription was created. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionJobSpecPostExecute

`func NewCatalogV1alpha1SubscriptionJobSpecPostExecute() *CatalogV1alpha1SubscriptionJobSpecPostExecute`

NewCatalogV1alpha1SubscriptionJobSpecPostExecute instantiates a new CatalogV1alpha1SubscriptionJobSpecPostExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionJobSpecPostExecuteWithDefaults

`func NewCatalogV1alpha1SubscriptionJobSpecPostExecuteWithDefaults() *CatalogV1alpha1SubscriptionJobSpecPostExecute`

NewCatalogV1alpha1SubscriptionJobSpecPostExecuteWithDefaults instantiates a new CatalogV1alpha1SubscriptionJobSpecPostExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnSuccess

`func (o *CatalogV1alpha1SubscriptionJobSpecPostExecute) GetOnSuccess() []CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInner`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *CatalogV1alpha1SubscriptionJobSpecPostExecute) GetOnSuccessOk() (*[]CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInner, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *CatalogV1alpha1SubscriptionJobSpecPostExecute) SetOnSuccess(v []CatalogV1alpha1SubscriptionJobSpecPostExecuteOnSuccessInner)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *CatalogV1alpha1SubscriptionJobSpecPostExecute) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


