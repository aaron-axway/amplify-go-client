# CatalogV1alpha1AssetResourceSpecSourceReleaseState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Current release state of the source such as stable or deprecated. | 
**Message** | Pointer to **string** | Optional info to be associated with the current state. If state is deprecated, then this is intended to indicate when the source will become archived/decommissioned.  | [optional] 

## Methods

### NewCatalogV1alpha1AssetResourceSpecSourceReleaseState

`func NewCatalogV1alpha1AssetResourceSpecSourceReleaseState(name string, ) *CatalogV1alpha1AssetResourceSpecSourceReleaseState`

NewCatalogV1alpha1AssetResourceSpecSourceReleaseState instantiates a new CatalogV1alpha1AssetResourceSpecSourceReleaseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetResourceSpecSourceReleaseStateWithDefaults

`func NewCatalogV1alpha1AssetResourceSpecSourceReleaseStateWithDefaults() *CatalogV1alpha1AssetResourceSpecSourceReleaseState`

NewCatalogV1alpha1AssetResourceSpecSourceReleaseStateWithDefaults instantiates a new CatalogV1alpha1AssetResourceSpecSourceReleaseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CatalogV1alpha1AssetResourceSpecSourceReleaseState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


