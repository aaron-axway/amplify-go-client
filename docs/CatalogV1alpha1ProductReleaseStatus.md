# CatalogV1alpha1ProductReleaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | The current status level, indicating progress towards consistency. | [optional] 
**Reasons** | Pointer to [**[]CatalogV1alpha1ProductReleaseStatusReasonsInner**](CatalogV1alpha1ProductReleaseStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewCatalogV1alpha1ProductReleaseStatus

`func NewCatalogV1alpha1ProductReleaseStatus() *CatalogV1alpha1ProductReleaseStatus`

NewCatalogV1alpha1ProductReleaseStatus instantiates a new CatalogV1alpha1ProductReleaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductReleaseStatusWithDefaults

`func NewCatalogV1alpha1ProductReleaseStatusWithDefaults() *CatalogV1alpha1ProductReleaseStatus`

NewCatalogV1alpha1ProductReleaseStatusWithDefaults instantiates a new CatalogV1alpha1ProductReleaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1ProductReleaseStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1ProductReleaseStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1ProductReleaseStatus) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CatalogV1alpha1ProductReleaseStatus) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetReasons

`func (o *CatalogV1alpha1ProductReleaseStatus) GetReasons() []CatalogV1alpha1ProductReleaseStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CatalogV1alpha1ProductReleaseStatus) GetReasonsOk() (*[]CatalogV1alpha1ProductReleaseStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CatalogV1alpha1ProductReleaseStatus) SetReasons(v []CatalogV1alpha1ProductReleaseStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *CatalogV1alpha1ProductReleaseStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


