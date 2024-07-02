# CatalogV1DocumentSpecSectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title for the section. | 
**Description** | Pointer to **string** | Description for the section. | [optional] 
**Articles** | Pointer to [**[]CatalogV1DocumentSpecSectionsInnerArticlesInner**](CatalogV1DocumentSpecSectionsInnerArticlesInner.md) | Section articles. | [optional] 

## Methods

### NewCatalogV1DocumentSpecSectionsInner

`func NewCatalogV1DocumentSpecSectionsInner(title string, ) *CatalogV1DocumentSpecSectionsInner`

NewCatalogV1DocumentSpecSectionsInner instantiates a new CatalogV1DocumentSpecSectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentSpecSectionsInnerWithDefaults

`func NewCatalogV1DocumentSpecSectionsInnerWithDefaults() *CatalogV1DocumentSpecSectionsInner`

NewCatalogV1DocumentSpecSectionsInnerWithDefaults instantiates a new CatalogV1DocumentSpecSectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CatalogV1DocumentSpecSectionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1DocumentSpecSectionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1DocumentSpecSectionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CatalogV1DocumentSpecSectionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1DocumentSpecSectionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1DocumentSpecSectionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1DocumentSpecSectionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArticles

`func (o *CatalogV1DocumentSpecSectionsInner) GetArticles() []CatalogV1DocumentSpecSectionsInnerArticlesInner`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *CatalogV1DocumentSpecSectionsInner) GetArticlesOk() (*[]CatalogV1DocumentSpecSectionsInnerArticlesInner, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *CatalogV1DocumentSpecSectionsInner) SetArticles(v []CatalogV1DocumentSpecSectionsInnerArticlesInner)`

SetArticles sets Articles field to given value.

### HasArticles

`func (o *CatalogV1DocumentSpecSectionsInner) HasArticles() bool`

HasArticles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


