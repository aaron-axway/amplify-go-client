# CatalogV1alpha1PublishedProductStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]CatalogV1alpha1SubscriptionStatusReasonsInner**](CatalogV1alpha1SubscriptionStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewCatalogV1alpha1PublishedProductStatus

`func NewCatalogV1alpha1PublishedProductStatus(level string, ) *CatalogV1alpha1PublishedProductStatus`

NewCatalogV1alpha1PublishedProductStatus instantiates a new CatalogV1alpha1PublishedProductStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1PublishedProductStatusWithDefaults

`func NewCatalogV1alpha1PublishedProductStatusWithDefaults() *CatalogV1alpha1PublishedProductStatus`

NewCatalogV1alpha1PublishedProductStatusWithDefaults instantiates a new CatalogV1alpha1PublishedProductStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1PublishedProductStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1PublishedProductStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1PublishedProductStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *CatalogV1alpha1PublishedProductStatus) GetReasons() []CatalogV1alpha1SubscriptionStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CatalogV1alpha1PublishedProductStatus) GetReasonsOk() (*[]CatalogV1alpha1SubscriptionStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CatalogV1alpha1PublishedProductStatus) SetReasons(v []CatalogV1alpha1SubscriptionStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CatalogV1alpha1PublishedProductStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


