# CatalogV1alpha1ReleaseTagSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Release Tag. | [optional] 
**ReleaseType** | **string** |  | 
**ReleaseVersionProperties** | Pointer to [**CatalogV1alpha1ReleaseTagSpecReleaseVersionProperties**](CatalogV1alpha1ReleaseTagSpecReleaseVersionProperties.md) |  | [optional] 
**PreviousReleases** | Pointer to [**CatalogV1alpha1ReleaseTagSpecPreviousReleases**](CatalogV1alpha1ReleaseTagSpecPreviousReleases.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1ReleaseTagSpec

`func NewCatalogV1alpha1ReleaseTagSpec(releaseType string, ) *CatalogV1alpha1ReleaseTagSpec`

NewCatalogV1alpha1ReleaseTagSpec instantiates a new CatalogV1alpha1ReleaseTagSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ReleaseTagSpecWithDefaults

`func NewCatalogV1alpha1ReleaseTagSpecWithDefaults() *CatalogV1alpha1ReleaseTagSpec`

NewCatalogV1alpha1ReleaseTagSpecWithDefaults instantiates a new CatalogV1alpha1ReleaseTagSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1alpha1ReleaseTagSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1alpha1ReleaseTagSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1alpha1ReleaseTagSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1alpha1ReleaseTagSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReleaseType

`func (o *CatalogV1alpha1ReleaseTagSpec) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *CatalogV1alpha1ReleaseTagSpec) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *CatalogV1alpha1ReleaseTagSpec) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.


### GetReleaseVersionProperties

`func (o *CatalogV1alpha1ReleaseTagSpec) GetReleaseVersionProperties() CatalogV1alpha1ReleaseTagSpecReleaseVersionProperties`

GetReleaseVersionProperties returns the ReleaseVersionProperties field if non-nil, zero value otherwise.

### GetReleaseVersionPropertiesOk

`func (o *CatalogV1alpha1ReleaseTagSpec) GetReleaseVersionPropertiesOk() (*CatalogV1alpha1ReleaseTagSpecReleaseVersionProperties, bool)`

GetReleaseVersionPropertiesOk returns a tuple with the ReleaseVersionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionProperties

`func (o *CatalogV1alpha1ReleaseTagSpec) SetReleaseVersionProperties(v CatalogV1alpha1ReleaseTagSpecReleaseVersionProperties)`

SetReleaseVersionProperties sets ReleaseVersionProperties field to given value.

### HasReleaseVersionProperties

`func (o *CatalogV1alpha1ReleaseTagSpec) HasReleaseVersionProperties() bool`

HasReleaseVersionProperties returns a boolean if a field has been set.

### GetPreviousReleases

`func (o *CatalogV1alpha1ReleaseTagSpec) GetPreviousReleases() CatalogV1alpha1ReleaseTagSpecPreviousReleases`

GetPreviousReleases returns the PreviousReleases field if non-nil, zero value otherwise.

### GetPreviousReleasesOk

`func (o *CatalogV1alpha1ReleaseTagSpec) GetPreviousReleasesOk() (*CatalogV1alpha1ReleaseTagSpecPreviousReleases, bool)`

GetPreviousReleasesOk returns a tuple with the PreviousReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousReleases

`func (o *CatalogV1alpha1ReleaseTagSpec) SetPreviousReleases(v CatalogV1alpha1ReleaseTagSpecPreviousReleases)`

SetPreviousReleases sets PreviousReleases field to given value.

### HasPreviousReleases

`func (o *CatalogV1alpha1ReleaseTagSpec) HasPreviousReleases() bool`

HasPreviousReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


