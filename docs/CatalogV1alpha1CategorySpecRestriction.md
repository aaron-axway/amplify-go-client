# CatalogV1alpha1CategorySpecRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Marketplace** | Pointer to [**CatalogV1CategorySpecProductRestrictionMarketplace**](CatalogV1CategorySpecProductRestrictionMarketplace.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1CategorySpecRestriction

`func NewCatalogV1alpha1CategorySpecRestriction(type_ string, ) *CatalogV1alpha1CategorySpecRestriction`

NewCatalogV1alpha1CategorySpecRestriction instantiates a new CatalogV1alpha1CategorySpecRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1CategorySpecRestrictionWithDefaults

`func NewCatalogV1alpha1CategorySpecRestrictionWithDefaults() *CatalogV1alpha1CategorySpecRestriction`

NewCatalogV1alpha1CategorySpecRestrictionWithDefaults instantiates a new CatalogV1alpha1CategorySpecRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1CategorySpecRestriction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1CategorySpecRestriction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1CategorySpecRestriction) SetType(v string)`

SetType sets Type field to given value.


### GetMarketplace

`func (o *CatalogV1alpha1CategorySpecRestriction) GetMarketplace() CatalogV1CategorySpecProductRestrictionMarketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *CatalogV1alpha1CategorySpecRestriction) GetMarketplaceOk() (*CatalogV1CategorySpecProductRestrictionMarketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *CatalogV1alpha1CategorySpecRestriction) SetMarketplace(v CatalogV1CategorySpecProductRestrictionMarketplace)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *CatalogV1alpha1CategorySpecRestriction) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


