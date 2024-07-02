# CatalogV1alpha1DocumentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]CatalogV1alpha1DocumentStatusReasonsInner**](CatalogV1alpha1DocumentStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewCatalogV1alpha1DocumentStatus

`func NewCatalogV1alpha1DocumentStatus(level string, ) *CatalogV1alpha1DocumentStatus`

NewCatalogV1alpha1DocumentStatus instantiates a new CatalogV1alpha1DocumentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1DocumentStatusWithDefaults

`func NewCatalogV1alpha1DocumentStatusWithDefaults() *CatalogV1alpha1DocumentStatus`

NewCatalogV1alpha1DocumentStatusWithDefaults instantiates a new CatalogV1alpha1DocumentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1DocumentStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1DocumentStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1DocumentStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *CatalogV1alpha1DocumentStatus) GetReasons() []CatalogV1alpha1DocumentStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CatalogV1alpha1DocumentStatus) GetReasonsOk() (*[]CatalogV1alpha1DocumentStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CatalogV1alpha1DocumentStatus) SetReasons(v []CatalogV1alpha1DocumentStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CatalogV1alpha1DocumentStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


