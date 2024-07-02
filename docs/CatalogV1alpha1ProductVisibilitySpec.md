# CatalogV1alpha1ProductVisibilitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]CatalogV1alpha1ProductVisibilitySpecProductsInner**](CatalogV1alpha1ProductVisibilitySpecProductsInner.md) | Defines where the visibility settings apply. | 
**Exclude** | Pointer to **bool** | Determines if the list of subjects should be excluded from the product visibility. | [optional] 
**Subjects** | Pointer to [**[]CatalogV1alpha1ProductVisibilitySpecSubjectsInner**](CatalogV1alpha1ProductVisibilitySpecSubjectsInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ProductVisibilitySpec

`func NewCatalogV1alpha1ProductVisibilitySpec(products []CatalogV1alpha1ProductVisibilitySpecProductsInner, ) *CatalogV1alpha1ProductVisibilitySpec`

NewCatalogV1alpha1ProductVisibilitySpec instantiates a new CatalogV1alpha1ProductVisibilitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ProductVisibilitySpecWithDefaults

`func NewCatalogV1alpha1ProductVisibilitySpecWithDefaults() *CatalogV1alpha1ProductVisibilitySpec`

NewCatalogV1alpha1ProductVisibilitySpecWithDefaults instantiates a new CatalogV1alpha1ProductVisibilitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetProducts() []CatalogV1alpha1ProductVisibilitySpecProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetProductsOk() (*[]CatalogV1alpha1ProductVisibilitySpecProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CatalogV1alpha1ProductVisibilitySpec) SetProducts(v []CatalogV1alpha1ProductVisibilitySpecProductsInner)`

SetProducts sets Products field to given value.


### GetExclude

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CatalogV1alpha1ProductVisibilitySpec) SetExclude(v bool)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CatalogV1alpha1ProductVisibilitySpec) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetSubjects

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetSubjects() []CatalogV1alpha1ProductVisibilitySpecSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CatalogV1alpha1ProductVisibilitySpec) GetSubjectsOk() (*[]CatalogV1alpha1ProductVisibilitySpecSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CatalogV1alpha1ProductVisibilitySpec) SetSubjects(v []CatalogV1alpha1ProductVisibilitySpecSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CatalogV1alpha1ProductVisibilitySpec) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


