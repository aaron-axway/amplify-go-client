# CatalogV1DocumentResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Document description. | [optional] 
**Version** | **string** | Version of the DocumentResource. | 
**Usage** | [**CatalogV1DocumentResourceSpecUsage**](CatalogV1DocumentResourceSpecUsage.md) |  | 
**Data** | [**CatalogV1DocumentResourceSpecData**](CatalogV1DocumentResourceSpecData.md) |  | 

## Methods

### NewCatalogV1DocumentResourceSpec

`func NewCatalogV1DocumentResourceSpec(version string, usage CatalogV1DocumentResourceSpecUsage, data CatalogV1DocumentResourceSpecData, ) *CatalogV1DocumentResourceSpec`

NewCatalogV1DocumentResourceSpec instantiates a new CatalogV1DocumentResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentResourceSpecWithDefaults

`func NewCatalogV1DocumentResourceSpecWithDefaults() *CatalogV1DocumentResourceSpec`

NewCatalogV1DocumentResourceSpecWithDefaults instantiates a new CatalogV1DocumentResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1DocumentResourceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1DocumentResourceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1DocumentResourceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1DocumentResourceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogV1DocumentResourceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1DocumentResourceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1DocumentResourceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetUsage

`func (o *CatalogV1DocumentResourceSpec) GetUsage() CatalogV1DocumentResourceSpecUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CatalogV1DocumentResourceSpec) GetUsageOk() (*CatalogV1DocumentResourceSpecUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CatalogV1DocumentResourceSpec) SetUsage(v CatalogV1DocumentResourceSpecUsage)`

SetUsage sets Usage field to given value.


### GetData

`func (o *CatalogV1DocumentResourceSpec) GetData() CatalogV1DocumentResourceSpecData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1DocumentResourceSpec) GetDataOk() (*CatalogV1DocumentResourceSpecData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1DocumentResourceSpec) SetData(v CatalogV1DocumentResourceSpecData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


