# CatalogV1DocumentResourceSpecBinaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Content** | **string** | Base64 encoded value of the file. | 
**FileName** | Pointer to **string** | The name of the file. | [optional] 
**FileType** | **string** | The type of the resource, example: pdf, markdown | 
**ContentType** | **string** | The content type | 

## Methods

### NewCatalogV1DocumentResourceSpecBinaryData

`func NewCatalogV1DocumentResourceSpecBinaryData(type_ string, content string, fileType string, contentType string, ) *CatalogV1DocumentResourceSpecBinaryData`

NewCatalogV1DocumentResourceSpecBinaryData instantiates a new CatalogV1DocumentResourceSpecBinaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentResourceSpecBinaryDataWithDefaults

`func NewCatalogV1DocumentResourceSpecBinaryDataWithDefaults() *CatalogV1DocumentResourceSpecBinaryData`

NewCatalogV1DocumentResourceSpecBinaryDataWithDefaults instantiates a new CatalogV1DocumentResourceSpecBinaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1DocumentResourceSpecBinaryData) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CatalogV1DocumentResourceSpecBinaryData) SetContent(v string)`

SetContent sets Content field to given value.


### GetFileName

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CatalogV1DocumentResourceSpecBinaryData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CatalogV1DocumentResourceSpecBinaryData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *CatalogV1DocumentResourceSpecBinaryData) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetContentType

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1DocumentResourceSpecBinaryData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1DocumentResourceSpecBinaryData) SetContentType(v string)`

SetContentType sets ContentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


