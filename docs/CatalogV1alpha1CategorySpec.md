# CatalogV1alpha1CategorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Markdown representing the category description. | [optional] 
**ParentCategory** | Pointer to **string** | Defines a parent category reference. Write access needed on the parent category to allow referencing it. | [optional] 
**Restriction** | Pointer to [**CatalogV1alpha1CategorySpecRestriction**](CatalogV1alpha1CategorySpecRestriction.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1CategorySpec

`func NewCatalogV1alpha1CategorySpec() *CatalogV1alpha1CategorySpec`

NewCatalogV1alpha1CategorySpec instantiates a new CatalogV1alpha1CategorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1CategorySpecWithDefaults

`func NewCatalogV1alpha1CategorySpecWithDefaults() *CatalogV1alpha1CategorySpec`

NewCatalogV1alpha1CategorySpecWithDefaults instantiates a new CatalogV1alpha1CategorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1CategorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1CategorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1CategorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1CategorySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentCategory

`func (o *CatalogV1alpha1CategorySpec) GetParentCategory() string`

GetParentCategory returns the ParentCategory field if non-nil, zero value otherwise.

### GetParentCategoryOk

`func (o *CatalogV1alpha1CategorySpec) GetParentCategoryOk() (*string, bool)`

GetParentCategoryOk returns a tuple with the ParentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategory

`func (o *CatalogV1alpha1CategorySpec) SetParentCategory(v string)`

SetParentCategory sets ParentCategory field to given value.

### HasParentCategory

`func (o *CatalogV1alpha1CategorySpec) HasParentCategory() bool`

HasParentCategory returns a boolean if a field has been set.

### GetRestriction

`func (o *CatalogV1alpha1CategorySpec) GetRestriction() CatalogV1alpha1CategorySpecRestriction`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *CatalogV1alpha1CategorySpec) GetRestrictionOk() (*CatalogV1alpha1CategorySpecRestriction, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *CatalogV1alpha1CategorySpec) SetRestriction(v CatalogV1alpha1CategorySpecRestriction)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *CatalogV1alpha1CategorySpec) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


