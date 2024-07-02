# CatalogV1alpha1ResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Resource description. | [optional] 
**Version** | Pointer to **string** | Version of the Resource. | [optional] 
**FileType** | **string** | The type of the resource, example: pdf | 
**ContentType** | **string** | The content type | 
**Data** | [**CatalogV1alpha1ResourceSpecData**](CatalogV1alpha1ResourceSpecData.md) |  | 

## Methods

### NewCatalogV1alpha1ResourceSpec

`func NewCatalogV1alpha1ResourceSpec(fileType string, contentType string, data CatalogV1alpha1ResourceSpecData, ) *CatalogV1alpha1ResourceSpec`

NewCatalogV1alpha1ResourceSpec instantiates a new CatalogV1alpha1ResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ResourceSpecWithDefaults

`func NewCatalogV1alpha1ResourceSpecWithDefaults() *CatalogV1alpha1ResourceSpec`

NewCatalogV1alpha1ResourceSpecWithDefaults instantiates a new CatalogV1alpha1ResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1ResourceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1ResourceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1ResourceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1ResourceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogV1alpha1ResourceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1alpha1ResourceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1alpha1ResourceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogV1alpha1ResourceSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFileType

`func (o *CatalogV1alpha1ResourceSpec) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *CatalogV1alpha1ResourceSpec) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *CatalogV1alpha1ResourceSpec) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetContentType

`func (o *CatalogV1alpha1ResourceSpec) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1alpha1ResourceSpec) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1alpha1ResourceSpec) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetData

`func (o *CatalogV1alpha1ResourceSpec) GetData() CatalogV1alpha1ResourceSpecData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1alpha1ResourceSpec) GetDataOk() (*CatalogV1alpha1ResourceSpecData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1alpha1ResourceSpec) SetData(v CatalogV1alpha1ResourceSpecData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


