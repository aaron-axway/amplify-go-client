# CatalogV1alpha1ReleaseTagStatusSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | Time when the change occurred. | 
**Detail** | **string** | message of the result | 
**Meta** | Pointer to [**CatalogV1alpha1ReleaseTagStatusSuccessMeta**](CatalogV1alpha1ReleaseTagStatusSuccessMeta.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ReleaseTagStatusSuccess

`func NewCatalogV1alpha1ReleaseTagStatusSuccess(type_ string, timestamp time.Time, detail string, ) *CatalogV1alpha1ReleaseTagStatusSuccess`

NewCatalogV1alpha1ReleaseTagStatusSuccess instantiates a new CatalogV1alpha1ReleaseTagStatusSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ReleaseTagStatusSuccessWithDefaults

`func NewCatalogV1alpha1ReleaseTagStatusSuccessWithDefaults() *CatalogV1alpha1ReleaseTagStatusSuccess`

NewCatalogV1alpha1ReleaseTagStatusSuccessWithDefaults instantiates a new CatalogV1alpha1ReleaseTagStatusSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDetail

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetMeta

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetMeta() CatalogV1alpha1ReleaseTagStatusSuccessMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) GetMetaOk() (*CatalogV1alpha1ReleaseTagStatusSuccessMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) SetMeta(v CatalogV1alpha1ReleaseTagStatusSuccessMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CatalogV1alpha1ReleaseTagStatusSuccess) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


