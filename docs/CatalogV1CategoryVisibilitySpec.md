# CatalogV1CategoryVisibilitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**Featured** | **bool** | Defines if the Category is featured in the marketplace. | 

## Methods

### NewCatalogV1CategoryVisibilitySpec

`func NewCatalogV1CategoryVisibilitySpec(category string, featured bool, ) *CatalogV1CategoryVisibilitySpec`

NewCatalogV1CategoryVisibilitySpec instantiates a new CatalogV1CategoryVisibilitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1CategoryVisibilitySpecWithDefaults

`func NewCatalogV1CategoryVisibilitySpecWithDefaults() *CatalogV1CategoryVisibilitySpec`

NewCatalogV1CategoryVisibilitySpecWithDefaults instantiates a new CatalogV1CategoryVisibilitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CatalogV1CategoryVisibilitySpec) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CatalogV1CategoryVisibilitySpec) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CatalogV1CategoryVisibilitySpec) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetFeatured

`func (o *CatalogV1CategoryVisibilitySpec) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CatalogV1CategoryVisibilitySpec) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CatalogV1CategoryVisibilitySpec) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


