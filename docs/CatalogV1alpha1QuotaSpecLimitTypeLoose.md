# CatalogV1alpha1QuotaSpecLimitTypeLoose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | **int32** | The limit of the unit that is provided. | 
**Overages** | [**CatalogV1QuotaSpecLimitTypeLooseOverages**](CatalogV1QuotaSpecLimitTypeLooseOverages.md) |  | 

## Methods

### NewCatalogV1alpha1QuotaSpecLimitTypeLoose

`func NewCatalogV1alpha1QuotaSpecLimitTypeLoose(type_ string, value int32, overages CatalogV1QuotaSpecLimitTypeLooseOverages, ) *CatalogV1alpha1QuotaSpecLimitTypeLoose`

NewCatalogV1alpha1QuotaSpecLimitTypeLoose instantiates a new CatalogV1alpha1QuotaSpecLimitTypeLoose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1QuotaSpecLimitTypeLooseWithDefaults

`func NewCatalogV1alpha1QuotaSpecLimitTypeLooseWithDefaults() *CatalogV1alpha1QuotaSpecLimitTypeLoose`

NewCatalogV1alpha1QuotaSpecLimitTypeLooseWithDefaults instantiates a new CatalogV1alpha1QuotaSpecLimitTypeLoose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) SetValue(v int32)`

SetValue sets Value field to given value.


### GetOverages

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetOverages() CatalogV1QuotaSpecLimitTypeLooseOverages`

GetOverages returns the Overages field if non-nil, zero value otherwise.

### GetOveragesOk

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) GetOveragesOk() (*CatalogV1QuotaSpecLimitTypeLooseOverages, bool)`

GetOveragesOk returns a tuple with the Overages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverages

`func (o *CatalogV1alpha1QuotaSpecLimitTypeLoose) SetOverages(v CatalogV1QuotaSpecLimitTypeLooseOverages)`

SetOverages sets Overages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


