# CatalogV1alpha1ProductReleaseStatusPending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | Time when the change occured. | 
**Detail** | **string** | Reason for being in Pending. | 
**Meta** | Pointer to [**CatalogV1ProductStatusReasonsInnerMeta**](CatalogV1ProductStatusReasonsInnerMeta.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductReleaseStatusPending

`func NewCatalogV1alpha1ProductReleaseStatusPending(type_ string, timestamp time.Time, detail string, ) *CatalogV1alpha1ProductReleaseStatusPending`

NewCatalogV1alpha1ProductReleaseStatusPending instantiates a new CatalogV1alpha1ProductReleaseStatusPending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductReleaseStatusPendingWithDefaults

`func NewCatalogV1alpha1ProductReleaseStatusPendingWithDefaults() *CatalogV1alpha1ProductReleaseStatusPending`

NewCatalogV1alpha1ProductReleaseStatusPendingWithDefaults instantiates a new CatalogV1alpha1ProductReleaseStatusPending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1ProductReleaseStatusPending) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CatalogV1alpha1ProductReleaseStatusPending) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDetail

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CatalogV1alpha1ProductReleaseStatusPending) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetMeta

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetMeta() CatalogV1ProductStatusReasonsInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CatalogV1alpha1ProductReleaseStatusPending) GetMetaOk() (*CatalogV1ProductStatusReasonsInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CatalogV1alpha1ProductReleaseStatusPending) SetMeta(v CatalogV1ProductStatusReasonsInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CatalogV1alpha1ProductReleaseStatusPending) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


