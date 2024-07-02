# CatalogV1alpha1QuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of the quota. | [optional] 
**Unit** | **string** | The unit used to measure the access to the resource. | 
**Pricing** | [**CatalogV1alpha1QuotaSpecPricing**](CatalogV1alpha1QuotaSpecPricing.md) |  | 
**Resources** | [**[]CatalogV1alpha1QuotaSpecResourcesInner**](CatalogV1alpha1QuotaSpecResourcesInner.md) | The resources that the access is being granted to. | 

## Methods

### NewCatalogV1alpha1QuotaSpec

`func NewCatalogV1alpha1QuotaSpec(unit string, pricing CatalogV1alpha1QuotaSpecPricing, resources []CatalogV1alpha1QuotaSpecResourcesInner, ) *CatalogV1alpha1QuotaSpec`

NewCatalogV1alpha1QuotaSpec instantiates a new CatalogV1alpha1QuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1QuotaSpecWithDefaults

`func NewCatalogV1alpha1QuotaSpecWithDefaults() *CatalogV1alpha1QuotaSpec`

NewCatalogV1alpha1QuotaSpecWithDefaults instantiates a new CatalogV1alpha1QuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1QuotaSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1QuotaSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1QuotaSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1QuotaSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *CatalogV1alpha1QuotaSpec) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CatalogV1alpha1QuotaSpec) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CatalogV1alpha1QuotaSpec) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetPricing

`func (o *CatalogV1alpha1QuotaSpec) GetPricing() CatalogV1alpha1QuotaSpecPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *CatalogV1alpha1QuotaSpec) GetPricingOk() (*CatalogV1alpha1QuotaSpecPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *CatalogV1alpha1QuotaSpec) SetPricing(v CatalogV1alpha1QuotaSpecPricing)`

SetPricing sets Pricing field to given value.


### GetResources

`func (o *CatalogV1alpha1QuotaSpec) GetResources() []CatalogV1alpha1QuotaSpecResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CatalogV1alpha1QuotaSpec) GetResourcesOk() (*[]CatalogV1alpha1QuotaSpecResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CatalogV1alpha1QuotaSpec) SetResources(v []CatalogV1alpha1QuotaSpecResourcesInner)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


