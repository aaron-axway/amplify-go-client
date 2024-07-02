# CatalogV1alpha1ConsumerProductVisibilitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]CatalogV1alpha1ProductVisibilitySpecProductsInner**](CatalogV1alpha1ProductVisibilitySpecProductsInner.md) | Defines where the visibility settings apply. | 
**Subjects** | Pointer to [**[]CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner**](CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ConsumerProductVisibilitySpec

`func NewCatalogV1alpha1ConsumerProductVisibilitySpec(products []CatalogV1alpha1ProductVisibilitySpecProductsInner, ) *CatalogV1alpha1ConsumerProductVisibilitySpec`

NewCatalogV1alpha1ConsumerProductVisibilitySpec instantiates a new CatalogV1alpha1ConsumerProductVisibilitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ConsumerProductVisibilitySpecWithDefaults

`func NewCatalogV1alpha1ConsumerProductVisibilitySpecWithDefaults() *CatalogV1alpha1ConsumerProductVisibilitySpec`

NewCatalogV1alpha1ConsumerProductVisibilitySpecWithDefaults instantiates a new CatalogV1alpha1ConsumerProductVisibilitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) GetProducts() []CatalogV1alpha1ProductVisibilitySpecProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) GetProductsOk() (*[]CatalogV1alpha1ProductVisibilitySpecProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) SetProducts(v []CatalogV1alpha1ProductVisibilitySpecProductsInner)`

SetProducts sets Products field to given value.


### GetSubjects

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) GetSubjects() []CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) GetSubjectsOk() (*[]CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) SetSubjects(v []CatalogV1alpha1ConsumerProductVisibilitySpecSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *CatalogV1alpha1ConsumerProductVisibilitySpec) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


