# CatalogV1alpha1SubscriptionApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Reason** | Pointer to **string** | Reason for the state. | [optional] 
**UserId** | Pointer to **string** | Id of the user that approved or reject the subscription. | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionApproval

`func NewCatalogV1alpha1SubscriptionApproval(state string, ) *CatalogV1alpha1SubscriptionApproval`

NewCatalogV1alpha1SubscriptionApproval instantiates a new CatalogV1alpha1SubscriptionApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionApprovalWithDefaults

`func NewCatalogV1alpha1SubscriptionApprovalWithDefaults() *CatalogV1alpha1SubscriptionApproval`

NewCatalogV1alpha1SubscriptionApprovalWithDefaults instantiates a new CatalogV1alpha1SubscriptionApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CatalogV1alpha1SubscriptionApproval) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1alpha1SubscriptionApproval) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1alpha1SubscriptionApproval) SetState(v string)`

SetState sets State field to given value.


### GetReason

`func (o *CatalogV1alpha1SubscriptionApproval) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CatalogV1alpha1SubscriptionApproval) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CatalogV1alpha1SubscriptionApproval) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CatalogV1alpha1SubscriptionApproval) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUserId

`func (o *CatalogV1alpha1SubscriptionApproval) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CatalogV1alpha1SubscriptionApproval) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CatalogV1alpha1SubscriptionApproval) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CatalogV1alpha1SubscriptionApproval) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


