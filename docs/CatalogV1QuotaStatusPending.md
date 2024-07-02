# CatalogV1QuotaStatusPending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **string** | Details of the Pending status. | 
**Timestamp** | **time.Time** | Time when the resource moved to Pending. | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCatalogV1QuotaStatusPending

`func NewCatalogV1QuotaStatusPending(type_ string, detail string, timestamp time.Time, ) *CatalogV1QuotaStatusPending`

NewCatalogV1QuotaStatusPending instantiates a new CatalogV1QuotaStatusPending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1QuotaStatusPendingWithDefaults

`func NewCatalogV1QuotaStatusPendingWithDefaults() *CatalogV1QuotaStatusPending`

NewCatalogV1QuotaStatusPendingWithDefaults instantiates a new CatalogV1QuotaStatusPending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1QuotaStatusPending) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1QuotaStatusPending) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1QuotaStatusPending) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *CatalogV1QuotaStatusPending) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CatalogV1QuotaStatusPending) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CatalogV1QuotaStatusPending) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetTimestamp

`func (o *CatalogV1QuotaStatusPending) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CatalogV1QuotaStatusPending) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CatalogV1QuotaStatusPending) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMeta

`func (o *CatalogV1QuotaStatusPending) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CatalogV1QuotaStatusPending) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CatalogV1QuotaStatusPending) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CatalogV1QuotaStatusPending) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

