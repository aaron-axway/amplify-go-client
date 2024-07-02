# CatalogV1ProductSpecAutoRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the generated release tag. | [optional] 
**ReleaseType** | **string** |  | 
**ReleaseVersionProperties** | Pointer to [**CatalogV1ProductSpecAutoReleaseReleaseVersionProperties**](CatalogV1ProductSpecAutoReleaseReleaseVersionProperties.md) |  | [optional] 
**RequiresInitialActivation** | Pointer to **bool** | Set true to suspend auto-release until product state changes to active or deprecated. This property will be automatically removed once activated.  | [optional] 
**PreviousReleases** | Pointer to [**CatalogV1ProductSpecAutoReleasePreviousReleases**](CatalogV1ProductSpecAutoReleasePreviousReleases.md) |  | [optional] 

## Methods

### NewCatalogV1ProductSpecAutoRelease

`func NewCatalogV1ProductSpecAutoRelease(releaseType string, ) *CatalogV1ProductSpecAutoRelease`

NewCatalogV1ProductSpecAutoRelease instantiates a new CatalogV1ProductSpecAutoRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1ProductSpecAutoReleaseWithDefaults

`func NewCatalogV1ProductSpecAutoReleaseWithDefaults() *CatalogV1ProductSpecAutoRelease`

NewCatalogV1ProductSpecAutoReleaseWithDefaults instantiates a new CatalogV1ProductSpecAutoRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CatalogV1ProductSpecAutoRelease) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogV1ProductSpecAutoRelease) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogV1ProductSpecAutoRelease) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogV1ProductSpecAutoRelease) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReleaseType

`func (o *CatalogV1ProductSpecAutoRelease) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *CatalogV1ProductSpecAutoRelease) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *CatalogV1ProductSpecAutoRelease) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.


### GetReleaseVersionProperties

`func (o *CatalogV1ProductSpecAutoRelease) GetReleaseVersionProperties() CatalogV1ProductSpecAutoReleaseReleaseVersionProperties`

GetReleaseVersionProperties returns the ReleaseVersionProperties field if non-nil, zero value otherwise.

### GetReleaseVersionPropertiesOk

`func (o *CatalogV1ProductSpecAutoRelease) GetReleaseVersionPropertiesOk() (*CatalogV1ProductSpecAutoReleaseReleaseVersionProperties, bool)`

GetReleaseVersionPropertiesOk returns a tuple with the ReleaseVersionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersionProperties

`func (o *CatalogV1ProductSpecAutoRelease) SetReleaseVersionProperties(v CatalogV1ProductSpecAutoReleaseReleaseVersionProperties)`

SetReleaseVersionProperties sets ReleaseVersionProperties field to given value.

### HasReleaseVersionProperties

`func (o *CatalogV1ProductSpecAutoRelease) HasReleaseVersionProperties() bool`

HasReleaseVersionProperties returns a boolean if a field has been set.

### GetRequiresInitialActivation

`func (o *CatalogV1ProductSpecAutoRelease) GetRequiresInitialActivation() bool`

GetRequiresInitialActivation returns the RequiresInitialActivation field if non-nil, zero value otherwise.

### GetRequiresInitialActivationOk

`func (o *CatalogV1ProductSpecAutoRelease) GetRequiresInitialActivationOk() (*bool, bool)`

GetRequiresInitialActivationOk returns a tuple with the RequiresInitialActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresInitialActivation

`func (o *CatalogV1ProductSpecAutoRelease) SetRequiresInitialActivation(v bool)`

SetRequiresInitialActivation sets RequiresInitialActivation field to given value.

### HasRequiresInitialActivation

`func (o *CatalogV1ProductSpecAutoRelease) HasRequiresInitialActivation() bool`

HasRequiresInitialActivation returns a boolean if a field has been set.

### GetPreviousReleases

`func (o *CatalogV1ProductSpecAutoRelease) GetPreviousReleases() CatalogV1ProductSpecAutoReleasePreviousReleases`

GetPreviousReleases returns the PreviousReleases field if non-nil, zero value otherwise.

### GetPreviousReleasesOk

`func (o *CatalogV1ProductSpecAutoRelease) GetPreviousReleasesOk() (*CatalogV1ProductSpecAutoReleasePreviousReleases, bool)`

GetPreviousReleasesOk returns a tuple with the PreviousReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousReleases

`func (o *CatalogV1ProductSpecAutoRelease) SetPreviousReleases(v CatalogV1ProductSpecAutoReleasePreviousReleases)`

SetPreviousReleases sets PreviousReleases field to given value.

### HasPreviousReleases

`func (o *CatalogV1ProductSpecAutoRelease) HasPreviousReleases() bool`

HasPreviousReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


