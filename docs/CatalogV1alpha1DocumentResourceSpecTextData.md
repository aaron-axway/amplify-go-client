# CatalogV1alpha1DocumentResourceSpecTextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Content** | **string** | Resource content. | 
**FileName** | Pointer to **string** | The name of the file. | [optional] 
**FileType** | **string** | The type of the resource, example: pdf | 
**ContentType** | **string** | The content type | 

## Methods

### NewCatalogV1alpha1DocumentResourceSpecTextData

`func NewCatalogV1alpha1DocumentResourceSpecTextData(type_ string, content string, fileType string, contentType string, ) *CatalogV1alpha1DocumentResourceSpecTextData`

NewCatalogV1alpha1DocumentResourceSpecTextData instantiates a new CatalogV1alpha1DocumentResourceSpecTextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1DocumentResourceSpecTextDataWithDefaults

`func NewCatalogV1alpha1DocumentResourceSpecTextDataWithDefaults() *CatalogV1alpha1DocumentResourceSpecTextData`

NewCatalogV1alpha1DocumentResourceSpecTextDataWithDefaults instantiates a new CatalogV1alpha1DocumentResourceSpecTextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) SetContent(v string)`

SetContent sets Content field to given value.


### GetFileName

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetContentType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1alpha1DocumentResourceSpecTextData) SetContentType(v string)`

SetContentType sets ContentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


