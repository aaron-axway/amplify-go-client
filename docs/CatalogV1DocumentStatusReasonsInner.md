# CatalogV1DocumentStatusReasonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **string** | message of the result | 
**Timestamp** | **time.Time** | Time when the change occured. | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCatalogV1DocumentStatusReasonsInner

`func NewCatalogV1DocumentStatusReasonsInner(type_ string, detail string, timestamp time.Time, ) *CatalogV1DocumentStatusReasonsInner`

NewCatalogV1DocumentStatusReasonsInner instantiates a new CatalogV1DocumentStatusReasonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentStatusReasonsInnerWithDefaults

`func NewCatalogV1DocumentStatusReasonsInnerWithDefaults() *CatalogV1DocumentStatusReasonsInner`

NewCatalogV1DocumentStatusReasonsInnerWithDefaults instantiates a new CatalogV1DocumentStatusReasonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1DocumentStatusReasonsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1DocumentStatusReasonsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1DocumentStatusReasonsInner) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *CatalogV1DocumentStatusReasonsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CatalogV1DocumentStatusReasonsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CatalogV1DocumentStatusReasonsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetTimestamp

`func (o *CatalogV1DocumentStatusReasonsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CatalogV1DocumentStatusReasonsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CatalogV1DocumentStatusReasonsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMeta

`func (o *CatalogV1DocumentStatusReasonsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CatalogV1DocumentStatusReasonsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CatalogV1DocumentStatusReasonsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CatalogV1DocumentStatusReasonsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


