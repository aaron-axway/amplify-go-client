# CatalogV1ResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Resource description. | [optional] 
**Version** | Pointer to **string** | Version of the Resource. | [optional] 
**FileType** | **string** | The type of the resource, example: pdf | 
**ContentType** | **string** | The content type | 
**Data** | [**CatalogV1ResourceSpecData**](CatalogV1ResourceSpecData.md) |  | 

## Methods

### NewCatalogV1ResourceSpec

`func NewCatalogV1ResourceSpec(fileType string, contentType string, data CatalogV1ResourceSpecData, ) *CatalogV1ResourceSpec`

NewCatalogV1ResourceSpec instantiates a new CatalogV1ResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1ResourceSpecWithDefaults

`func NewCatalogV1ResourceSpecWithDefaults() *CatalogV1ResourceSpec`

NewCatalogV1ResourceSpecWithDefaults instantiates a new CatalogV1ResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1ResourceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1ResourceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1ResourceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1ResourceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogV1ResourceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1ResourceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1ResourceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogV1ResourceSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFileType

`func (o *CatalogV1ResourceSpec) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *CatalogV1ResourceSpec) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *CatalogV1ResourceSpec) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetContentType

`func (o *CatalogV1ResourceSpec) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1ResourceSpec) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1ResourceSpec) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetData

`func (o *CatalogV1ResourceSpec) GetData() CatalogV1ResourceSpecData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1ResourceSpec) GetDataOk() (*CatalogV1ResourceSpecData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1ResourceSpec) SetData(v CatalogV1ResourceSpecData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


