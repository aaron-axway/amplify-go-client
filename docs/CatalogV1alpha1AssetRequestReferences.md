# CatalogV1alpha1AssetRequestReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetRequestDefinition** | Pointer to **string** |  | [optional] 
**AssetRelease** | Pointer to **string** | Reference to Release that got created from this asset request. | [optional] 
**Asset** | Pointer to **string** | Reference to Release that got created from this asset request. | [optional] 
**CredentialRequestDefinitions** | Pointer to **[]string** |  | [optional] 
**AccessRequest** | Pointer to **string** | Reference to Access Request resource | [optional] 

## Methods

### NewCatalogV1alpha1AssetRequestReferences

`func NewCatalogV1alpha1AssetRequestReferences() *CatalogV1alpha1AssetRequestReferences`

NewCatalogV1alpha1AssetRequestReferences instantiates a new CatalogV1alpha1AssetRequestReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetRequestReferencesWithDefaults

`func NewCatalogV1alpha1AssetRequestReferencesWithDefaults() *CatalogV1alpha1AssetRequestReferences`

NewCatalogV1alpha1AssetRequestReferencesWithDefaults instantiates a new CatalogV1alpha1AssetRequestReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetRequestDefinition

`func (o *CatalogV1alpha1AssetRequestReferences) GetAssetRequestDefinition() string`

GetAssetRequestDefinition returns the AssetRequestDefinition field if non-nil, zero value otherwise.

### GetAssetRequestDefinitionOk

`func (o *CatalogV1alpha1AssetRequestReferences) GetAssetRequestDefinitionOk() (*string, bool)`

GetAssetRequestDefinitionOk returns a tuple with the AssetRequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetRequestDefinition

`func (o *CatalogV1alpha1AssetRequestReferences) SetAssetRequestDefinition(v string)`

SetAssetRequestDefinition sets AssetRequestDefinition field to given value.

### HasAssetRequestDefinition

`func (o *CatalogV1alpha1AssetRequestReferences) HasAssetRequestDefinition() bool`

HasAssetRequestDefinition returns a boolean if a field has been set.

### GetAssetRelease

`func (o *CatalogV1alpha1AssetRequestReferences) GetAssetRelease() string`

GetAssetRelease returns the AssetRelease field if non-nil, zero value otherwise.

### GetAssetReleaseOk

`func (o *CatalogV1alpha1AssetRequestReferences) GetAssetReleaseOk() (*string, bool)`

GetAssetReleaseOk returns a tuple with the AssetRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetRelease

`func (o *CatalogV1alpha1AssetRequestReferences) SetAssetRelease(v string)`

SetAssetRelease sets AssetRelease field to given value.

### HasAssetRelease

`func (o *CatalogV1alpha1AssetRequestReferences) HasAssetRelease() bool`

HasAssetRelease returns a boolean if a field has been set.

### GetAsset

`func (o *CatalogV1alpha1AssetRequestReferences) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CatalogV1alpha1AssetRequestReferences) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CatalogV1alpha1AssetRequestReferences) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *CatalogV1alpha1AssetRequestReferences) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetRequestReferences) GetCredentialRequestDefinitions() []string`

GetCredentialRequestDefinitions returns the CredentialRequestDefinitions field if non-nil, zero value otherwise.

### GetCredentialRequestDefinitionsOk

`func (o *CatalogV1alpha1AssetRequestReferences) GetCredentialRequestDefinitionsOk() (*[]string, bool)`

GetCredentialRequestDefinitionsOk returns a tuple with the CredentialRequestDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetRequestReferences) SetCredentialRequestDefinitions(v []string)`

SetCredentialRequestDefinitions sets CredentialRequestDefinitions field to given value.

### HasCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetRequestReferences) HasCredentialRequestDefinitions() bool`

HasCredentialRequestDefinitions returns a boolean if a field has been set.

### GetAccessRequest

`func (o *CatalogV1alpha1AssetRequestReferences) GetAccessRequest() string`

GetAccessRequest returns the AccessRequest field if non-nil, zero value otherwise.

### GetAccessRequestOk

`func (o *CatalogV1alpha1AssetRequestReferences) GetAccessRequestOk() (*string, bool)`

GetAccessRequestOk returns a tuple with the AccessRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequest

`func (o *CatalogV1alpha1AssetRequestReferences) SetAccessRequest(v string)`

SetAccessRequest sets AccessRequest field to given value.

### HasAccessRequest

`func (o *CatalogV1alpha1AssetRequestReferences) HasAccessRequest() bool`

HasAccessRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


