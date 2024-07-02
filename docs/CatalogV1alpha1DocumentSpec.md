# CatalogV1alpha1DocumentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Document description. | [optional] 
**Rank** | Pointer to **float32** | Rank of document. | [optional] [default to 0]
**Sections** | Pointer to [**[]CatalogV1alpha1DocumentSpecSectionsInner**](CatalogV1alpha1DocumentSpecSectionsInner.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1DocumentSpec

`func NewCatalogV1alpha1DocumentSpec() *CatalogV1alpha1DocumentSpec`

NewCatalogV1alpha1DocumentSpec instantiates a new CatalogV1alpha1DocumentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1DocumentSpecWithDefaults

`func NewCatalogV1alpha1DocumentSpecWithDefaults() *CatalogV1alpha1DocumentSpec`

NewCatalogV1alpha1DocumentSpecWithDefaults instantiates a new CatalogV1alpha1DocumentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1DocumentSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1DocumentSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1DocumentSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1DocumentSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRank

`func (o *CatalogV1alpha1DocumentSpec) GetRank() float32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *CatalogV1alpha1DocumentSpec) GetRankOk() (*float32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *CatalogV1alpha1DocumentSpec) SetRank(v float32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *CatalogV1alpha1DocumentSpec) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSections

`func (o *CatalogV1alpha1DocumentSpec) GetSections() []CatalogV1alpha1DocumentSpecSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CatalogV1alpha1DocumentSpec) GetSectionsOk() (*[]CatalogV1alpha1DocumentSpecSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CatalogV1alpha1DocumentSpec) SetSections(v []CatalogV1alpha1DocumentSpecSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *CatalogV1alpha1DocumentSpec) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


