# CatalogV1alpha1QuotaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]CatalogV1alpha1QuotaStatusReasonsInner**](CatalogV1alpha1QuotaStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewCatalogV1alpha1QuotaStatus

`func NewCatalogV1alpha1QuotaStatus(level string, ) *CatalogV1alpha1QuotaStatus`

NewCatalogV1alpha1QuotaStatus instantiates a new CatalogV1alpha1QuotaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1QuotaStatusWithDefaults

`func NewCatalogV1alpha1QuotaStatusWithDefaults() *CatalogV1alpha1QuotaStatus`

NewCatalogV1alpha1QuotaStatusWithDefaults instantiates a new CatalogV1alpha1QuotaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1QuotaStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1QuotaStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1QuotaStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *CatalogV1alpha1QuotaStatus) GetReasons() []CatalogV1alpha1QuotaStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CatalogV1alpha1QuotaStatus) GetReasonsOk() (*[]CatalogV1alpha1QuotaStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CatalogV1alpha1QuotaStatus) SetReasons(v []CatalogV1alpha1QuotaStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CatalogV1alpha1QuotaStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


