# CatalogV1alpha1SubscriptionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Reason** | Pointer to **string** | Additional info on the state. | [optional] 
**When** | Pointer to [**CatalogV1alpha1SubscriptionStateWhen**](CatalogV1alpha1SubscriptionStateWhen.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionState

`func NewCatalogV1alpha1SubscriptionState(name string, ) *CatalogV1alpha1SubscriptionState`

NewCatalogV1alpha1SubscriptionState instantiates a new CatalogV1alpha1SubscriptionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionStateWithDefaults

`func NewCatalogV1alpha1SubscriptionStateWithDefaults() *CatalogV1alpha1SubscriptionState`

NewCatalogV1alpha1SubscriptionStateWithDefaults instantiates a new CatalogV1alpha1SubscriptionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogV1alpha1SubscriptionState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1SubscriptionState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1SubscriptionState) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *CatalogV1alpha1SubscriptionState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CatalogV1alpha1SubscriptionState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CatalogV1alpha1SubscriptionState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CatalogV1alpha1SubscriptionState) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetWhen

`func (o *CatalogV1alpha1SubscriptionState) GetWhen() CatalogV1alpha1SubscriptionStateWhen`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *CatalogV1alpha1SubscriptionState) GetWhenOk() (*CatalogV1alpha1SubscriptionStateWhen, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *CatalogV1alpha1SubscriptionState) SetWhen(v CatalogV1alpha1SubscriptionStateWhen)`

SetWhen sets When field to given value.

### HasWhen

`func (o *CatalogV1alpha1SubscriptionState) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


