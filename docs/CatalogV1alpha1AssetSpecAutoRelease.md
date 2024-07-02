# CatalogV1alpha1AssetSpecAutoRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseType** | **string** |  | 
**RequiresInitialActivation** | Pointer to **bool** | Set true to suspend auto-release until asset state changes to active or deprecated. This property will be automatically removed once activated.  | [optional] 
**PreviousReleases** | Pointer to [**CatalogV1alpha1AssetSpecAutoReleasePreviousReleases**](CatalogV1alpha1AssetSpecAutoReleasePreviousReleases.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetSpecAutoRelease

`func NewCatalogV1alpha1AssetSpecAutoRelease(releaseType string, ) *CatalogV1alpha1AssetSpecAutoRelease`

NewCatalogV1alpha1AssetSpecAutoRelease instantiates a new CatalogV1alpha1AssetSpecAutoRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetSpecAutoReleaseWithDefaults

`func NewCatalogV1alpha1AssetSpecAutoReleaseWithDefaults() *CatalogV1alpha1AssetSpecAutoRelease`

NewCatalogV1alpha1AssetSpecAutoReleaseWithDefaults instantiates a new CatalogV1alpha1AssetSpecAutoRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseType

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetReleaseType() string`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### GetReleaseTypeOk

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetReleaseTypeOk() (*string, bool)`

GetReleaseTypeOk returns a tuple with the ReleaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseType

`func (o *CatalogV1alpha1AssetSpecAutoRelease) SetReleaseType(v string)`

SetReleaseType sets ReleaseType field to given value.


### GetRequiresInitialActivation

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetRequiresInitialActivation() bool`

GetRequiresInitialActivation returns the RequiresInitialActivation field if non-nil, zero value otherwise.

### GetRequiresInitialActivationOk

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetRequiresInitialActivationOk() (*bool, bool)`

GetRequiresInitialActivationOk returns a tuple with the RequiresInitialActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresInitialActivation

`func (o *CatalogV1alpha1AssetSpecAutoRelease) SetRequiresInitialActivation(v bool)`

SetRequiresInitialActivation sets RequiresInitialActivation field to given value.

### HasRequiresInitialActivation

`func (o *CatalogV1alpha1AssetSpecAutoRelease) HasRequiresInitialActivation() bool`

HasRequiresInitialActivation returns a boolean if a field has been set.

### GetPreviousReleases

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetPreviousReleases() CatalogV1alpha1AssetSpecAutoReleasePreviousReleases`

GetPreviousReleases returns the PreviousReleases field if non-nil, zero value otherwise.

### GetPreviousReleasesOk

`func (o *CatalogV1alpha1AssetSpecAutoRelease) GetPreviousReleasesOk() (*CatalogV1alpha1AssetSpecAutoReleasePreviousReleases, bool)`

GetPreviousReleasesOk returns a tuple with the PreviousReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousReleases

`func (o *CatalogV1alpha1AssetSpecAutoRelease) SetPreviousReleases(v CatalogV1alpha1AssetSpecAutoReleasePreviousReleases)`

SetPreviousReleases sets PreviousReleases field to given value.

### HasPreviousReleases

`func (o *CatalogV1alpha1AssetSpecAutoRelease) HasPreviousReleases() bool`

HasPreviousReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


