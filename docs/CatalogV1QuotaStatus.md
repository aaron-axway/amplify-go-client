# CatalogV1QuotaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]CatalogV1QuotaStatusReasonsInner**](CatalogV1QuotaStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewCatalogV1QuotaStatus

`func NewCatalogV1QuotaStatus(level string, ) *CatalogV1QuotaStatus`

NewCatalogV1QuotaStatus instantiates a new CatalogV1QuotaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1QuotaStatusWithDefaults

`func NewCatalogV1QuotaStatusWithDefaults() *CatalogV1QuotaStatus`

NewCatalogV1QuotaStatusWithDefaults instantiates a new CatalogV1QuotaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1QuotaStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1QuotaStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1QuotaStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *CatalogV1QuotaStatus) GetReasons() []CatalogV1QuotaStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CatalogV1QuotaStatus) GetReasonsOk() (*[]CatalogV1QuotaStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CatalogV1QuotaStatus) SetReasons(v []CatalogV1QuotaStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CatalogV1QuotaStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


