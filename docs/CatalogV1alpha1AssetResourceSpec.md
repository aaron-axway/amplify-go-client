# CatalogV1alpha1AssetResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | Pointer to **string** | The Stage this Asset Resource is deployed on. | [optional] 
**AssetRequestDefinition** | Pointer to **string** |  | [optional] 
**CredentialRequestDefinitions** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | 
**ContentType** | Pointer to **string** | content-type of the spec. | [optional] 
**Version** | Pointer to **string** | The version of referenced resource. | [optional] 
**Definition** | **string** | Base64 encoded value of the api specification. | 
**Status** | **string** | Resource availability | 
**AccessInfo** | Pointer to [**[]CatalogV1alpha1AssetResourceSpecAccessInfoInner**](CatalogV1alpha1AssetResourceSpecAccessInfoInner.md) | information to access the definition. | [optional] 
**SourceReleaseState** | Pointer to [**CatalogV1alpha1AssetResourceSpecSourceReleaseState**](CatalogV1alpha1AssetResourceSpecSourceReleaseState.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetResourceSpec

`func NewCatalogV1alpha1AssetResourceSpec(type_ string, definition string, status string, ) *CatalogV1alpha1AssetResourceSpec`

NewCatalogV1alpha1AssetResourceSpec instantiates a new CatalogV1alpha1AssetResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetResourceSpecWithDefaults

`func NewCatalogV1alpha1AssetResourceSpecWithDefaults() *CatalogV1alpha1AssetResourceSpec`

NewCatalogV1alpha1AssetResourceSpecWithDefaults instantiates a new CatalogV1alpha1AssetResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *CatalogV1alpha1AssetResourceSpec) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *CatalogV1alpha1AssetResourceSpec) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *CatalogV1alpha1AssetResourceSpec) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetAssetRequestDefinition

`func (o *CatalogV1alpha1AssetResourceSpec) GetAssetRequestDefinition() string`

GetAssetRequestDefinition returns the AssetRequestDefinition field if non-nil, zero value otherwise.

### GetAssetRequestDefinitionOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetAssetRequestDefinitionOk() (*string, bool)`

GetAssetRequestDefinitionOk returns a tuple with the AssetRequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetRequestDefinition

`func (o *CatalogV1alpha1AssetResourceSpec) SetAssetRequestDefinition(v string)`

SetAssetRequestDefinition sets AssetRequestDefinition field to given value.

### HasAssetRequestDefinition

`func (o *CatalogV1alpha1AssetResourceSpec) HasAssetRequestDefinition() bool`

HasAssetRequestDefinition returns a boolean if a field has been set.

### GetCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetResourceSpec) GetCredentialRequestDefinitions() []string`

GetCredentialRequestDefinitions returns the CredentialRequestDefinitions field if non-nil, zero value otherwise.

### GetCredentialRequestDefinitionsOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetCredentialRequestDefinitionsOk() (*[]string, bool)`

GetCredentialRequestDefinitionsOk returns a tuple with the CredentialRequestDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetResourceSpec) SetCredentialRequestDefinitions(v []string)`

SetCredentialRequestDefinitions sets CredentialRequestDefinitions field to given value.

### HasCredentialRequestDefinitions

`func (o *CatalogV1alpha1AssetResourceSpec) HasCredentialRequestDefinitions() bool`

HasCredentialRequestDefinitions returns a boolean if a field has been set.

### GetType

`func (o *CatalogV1alpha1AssetResourceSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogV1alpha1AssetResourceSpec) SetType(v string)`

SetType sets Type field to given value.


### GetContentType

`func (o *CatalogV1alpha1AssetResourceSpec) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CatalogV1alpha1AssetResourceSpec) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CatalogV1alpha1AssetResourceSpec) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetVersion

`func (o *CatalogV1alpha1AssetResourceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogV1alpha1AssetResourceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogV1alpha1AssetResourceSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefinition

`func (o *CatalogV1alpha1AssetResourceSpec) GetDefinition() string`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetDefinitionOk() (*string, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CatalogV1alpha1AssetResourceSpec) SetDefinition(v string)`

SetDefinition sets Definition field to given value.


### GetStatus

`func (o *CatalogV1alpha1AssetResourceSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogV1alpha1AssetResourceSpec) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccessInfo

`func (o *CatalogV1alpha1AssetResourceSpec) GetAccessInfo() []CatalogV1alpha1AssetResourceSpecAccessInfoInner`

GetAccessInfo returns the AccessInfo field if non-nil, zero value otherwise.

### GetAccessInfoOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetAccessInfoOk() (*[]CatalogV1alpha1AssetResourceSpecAccessInfoInner, bool)`

GetAccessInfoOk returns a tuple with the AccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInfo

`func (o *CatalogV1alpha1AssetResourceSpec) SetAccessInfo(v []CatalogV1alpha1AssetResourceSpecAccessInfoInner)`

SetAccessInfo sets AccessInfo field to given value.

### HasAccessInfo

`func (o *CatalogV1alpha1AssetResourceSpec) HasAccessInfo() bool`

HasAccessInfo returns a boolean if a field has been set.

### GetSourceReleaseState

`func (o *CatalogV1alpha1AssetResourceSpec) GetSourceReleaseState() CatalogV1alpha1AssetResourceSpecSourceReleaseState`

GetSourceReleaseState returns the SourceReleaseState field if non-nil, zero value otherwise.

### GetSourceReleaseStateOk

`func (o *CatalogV1alpha1AssetResourceSpec) GetSourceReleaseStateOk() (*CatalogV1alpha1AssetResourceSpecSourceReleaseState, bool)`

GetSourceReleaseStateOk returns a tuple with the SourceReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReleaseState

`func (o *CatalogV1alpha1AssetResourceSpec) SetSourceReleaseState(v CatalogV1alpha1AssetResourceSpecSourceReleaseState)`

SetSourceReleaseState sets SourceReleaseState field to given value.

### HasSourceReleaseState

`func (o *CatalogV1alpha1AssetResourceSpec) HasSourceReleaseState() bool`

HasSourceReleaseState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


