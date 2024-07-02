# CatalogV1ProductReleaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the product when the release was generated. | [optional] 
**Version** | **string** | Version of the release. | 
**VersionProperties** | Pointer to [**CatalogV1ProductReleaseSpecVersionProperties**](CatalogV1ProductReleaseSpecVersionProperties.md) |  | [optional] 
**Product** | **string** |  | 
**Assets** | Pointer to [**[]CatalogV1ProductReleaseSpecAssetsInner**](CatalogV1ProductReleaseSpecAssetsInner.md) |  | [optional] 
**ReleaseTag** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** | list of categories for the released product. | [optional] 

## Methods

### NewCatalogV1ProductReleaseSpec

`func NewCatalogV1ProductReleaseSpec(version string, product string, releaseTag string, ) *CatalogV1ProductReleaseSpec`

NewCatalogV1ProductReleaseSpec instantiates a new CatalogV1ProductReleaseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1ProductReleaseSpecWithDefaults

`func NewCatalogV1ProductReleaseSpecWithDefaults() *CatalogV1ProductReleaseSpec`

NewCatalogV1ProductReleaseSpecWithDefaults instantiates a new CatalogV1ProductReleaseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1ProductReleaseSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1ProductReleaseSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1ProductReleaseSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1ProductReleaseSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogV1ProductReleaseSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1ProductReleaseSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1ProductReleaseSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionProperties

`func (o *CatalogV1ProductReleaseSpec) GetVersionProperties() CatalogV1ProductReleaseSpecVersionProperties`

GetVersionProperties returns the VersionProperties field if non-nil, zero value otherwise.

### GetVersionPropertiesOk

`func (o *CatalogV1ProductReleaseSpec) GetVersionPropertiesOk() (*CatalogV1ProductReleaseSpecVersionProperties, bool)`

GetVersionPropertiesOk returns a tuple with the VersionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionProperties

`func (o *CatalogV1ProductReleaseSpec) SetVersionProperties(v CatalogV1ProductReleaseSpecVersionProperties)`

SetVersionProperties sets VersionProperties field to given value.

### HasVersionProperties

`func (o *CatalogV1ProductReleaseSpec) HasVersionProperties() bool`

HasVersionProperties returns a boolean if a field has been set.

### GetProduct

`func (o *CatalogV1ProductReleaseSpec) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogV1ProductReleaseSpec) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CatalogV1ProductReleaseSpec) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetAssets

`func (o *CatalogV1ProductReleaseSpec) GetAssets() []CatalogV1ProductReleaseSpecAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CatalogV1ProductReleaseSpec) GetAssetsOk() (*[]CatalogV1ProductReleaseSpecAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CatalogV1ProductReleaseSpec) SetAssets(v []CatalogV1ProductReleaseSpecAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CatalogV1ProductReleaseSpec) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetReleaseTag

`func (o *CatalogV1ProductReleaseSpec) GetReleaseTag() string`

GetReleaseTag returns the ReleaseTag field if non-nil, zero value otherwise.

### GetReleaseTagOk

`func (o *CatalogV1ProductReleaseSpec) GetReleaseTagOk() (*string, bool)`

GetReleaseTagOk returns a tuple with the ReleaseTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTag

`func (o *CatalogV1ProductReleaseSpec) SetReleaseTag(v string)`

SetReleaseTag sets ReleaseTag field to given value.


### GetState

`func (o *CatalogV1ProductReleaseSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1ProductReleaseSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1ProductReleaseSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CatalogV1ProductReleaseSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCategories

`func (o *CatalogV1ProductReleaseSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CatalogV1ProductReleaseSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CatalogV1ProductReleaseSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CatalogV1ProductReleaseSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


