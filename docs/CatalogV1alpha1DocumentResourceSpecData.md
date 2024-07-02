# CatalogV1alpha1DocumentResourceSpecData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Content** | **string** | Base64 encoded value of the file. | 
**FileName** | Pointer to **string** | The name of the file. | [optional] 
**FileType** | **string** | The type of the resource, example: pdf, markdown | 
**ContentType** | **string** | The content type | 
**Value** | **string** | URL value. | 

## Methods

### NewCatalogV1alpha1DocumentResourceSpecData

`func NewCatalogV1alpha1DocumentResourceSpecData(type_ string, content string, fileType string, contentType string, value string, ) *CatalogV1alpha1DocumentResourceSpecData`

NewCatalogV1alpha1DocumentResourceSpecData instantiates a new CatalogV1alpha1DocumentResourceSpecData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1DocumentResourceSpecDataWithDefaults

`func NewCatalogV1alpha1DocumentResourceSpecDataWithDefaults() *CatalogV1alpha1DocumentResourceSpecData`

NewCatalogV1alpha1DocumentResourceSpecDataWithDefaults instantiates a new CatalogV1alpha1DocumentResourceSpecData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetContent(v string)`

SetContent sets Content field to given value.


### GetFileName

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CatalogV1alpha1DocumentResourceSpecData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetContentType

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetValue

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CatalogV1alpha1DocumentResourceSpecData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CatalogV1alpha1DocumentResourceSpecData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


