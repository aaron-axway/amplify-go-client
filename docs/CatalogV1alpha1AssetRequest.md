# CatalogV1alpha1AssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the catalog group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "catalog"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind AssetRequest. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "AssetRequest"]
**Name** | Pointer to **string** | Name of the AssetRequest. AssetRequest name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the AssetRequest. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**CatalogV1alpha1AssetRequestSpec**](CatalogV1alpha1AssetRequestSpec.md) |  | 
**Marketplace** | Pointer to [**CatalogV1alpha1AssetRequestMarketplace**](CatalogV1alpha1AssetRequestMarketplace.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Data matching the provisioning schema from the AssetRequestDefinition | [optional] 
**References** | Pointer to [**CatalogV1alpha1AssetRequestReferences**](CatalogV1alpha1AssetRequestReferences.md) |  | [optional] 
**Approval** | Pointer to [**CatalogV1alpha1AssetRequestApproval**](CatalogV1alpha1AssetRequestApproval.md) |  | [optional] 
**Status** | Pointer to [**CatalogV1alpha1AssetRequestStatus**](CatalogV1alpha1AssetRequestStatus.md) |  | [optional] 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetRequest

`func NewCatalogV1alpha1AssetRequest(spec CatalogV1alpha1AssetRequestSpec, ) *CatalogV1alpha1AssetRequest`

NewCatalogV1alpha1AssetRequest instantiates a new CatalogV1alpha1AssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetRequestWithDefaults

`func NewCatalogV1alpha1AssetRequestWithDefaults() *CatalogV1alpha1AssetRequest`

NewCatalogV1alpha1AssetRequestWithDefaults instantiates a new CatalogV1alpha1AssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogV1alpha1AssetRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogV1alpha1AssetRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogV1alpha1AssetRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CatalogV1alpha1AssetRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *CatalogV1alpha1AssetRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CatalogV1alpha1AssetRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CatalogV1alpha1AssetRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CatalogV1alpha1AssetRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1alpha1AssetRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1alpha1AssetRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1alpha1AssetRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogV1alpha1AssetRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CatalogV1alpha1AssetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1AssetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1AssetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1alpha1AssetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *CatalogV1alpha1AssetRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1alpha1AssetRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1alpha1AssetRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogV1alpha1AssetRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogV1alpha1AssetRequest) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogV1alpha1AssetRequest) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogV1alpha1AssetRequest) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogV1alpha1AssetRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *CatalogV1alpha1AssetRequest) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CatalogV1alpha1AssetRequest) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CatalogV1alpha1AssetRequest) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CatalogV1alpha1AssetRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *CatalogV1alpha1AssetRequest) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CatalogV1alpha1AssetRequest) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CatalogV1alpha1AssetRequest) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CatalogV1alpha1AssetRequest) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *CatalogV1alpha1AssetRequest) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogV1alpha1AssetRequest) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogV1alpha1AssetRequest) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CatalogV1alpha1AssetRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *CatalogV1alpha1AssetRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogV1alpha1AssetRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogV1alpha1AssetRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogV1alpha1AssetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *CatalogV1alpha1AssetRequest) GetSpec() CatalogV1alpha1AssetRequestSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CatalogV1alpha1AssetRequest) GetSpecOk() (*CatalogV1alpha1AssetRequestSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CatalogV1alpha1AssetRequest) SetSpec(v CatalogV1alpha1AssetRequestSpec)`

SetSpec sets Spec field to given value.


### GetMarketplace

`func (o *CatalogV1alpha1AssetRequest) GetMarketplace() CatalogV1alpha1AssetRequestMarketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *CatalogV1alpha1AssetRequest) GetMarketplaceOk() (*CatalogV1alpha1AssetRequestMarketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *CatalogV1alpha1AssetRequest) SetMarketplace(v CatalogV1alpha1AssetRequestMarketplace)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *CatalogV1alpha1AssetRequest) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetData

`func (o *CatalogV1alpha1AssetRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1alpha1AssetRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1alpha1AssetRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CatalogV1alpha1AssetRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetReferences

`func (o *CatalogV1alpha1AssetRequest) GetReferences() CatalogV1alpha1AssetRequestReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CatalogV1alpha1AssetRequest) GetReferencesOk() (*CatalogV1alpha1AssetRequestReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CatalogV1alpha1AssetRequest) SetReferences(v CatalogV1alpha1AssetRequestReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CatalogV1alpha1AssetRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetApproval

`func (o *CatalogV1alpha1AssetRequest) GetApproval() CatalogV1alpha1AssetRequestApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *CatalogV1alpha1AssetRequest) GetApprovalOk() (*CatalogV1alpha1AssetRequestApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *CatalogV1alpha1AssetRequest) SetApproval(v CatalogV1alpha1AssetRequestApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *CatalogV1alpha1AssetRequest) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogV1alpha1AssetRequest) GetStatus() CatalogV1alpha1AssetRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogV1alpha1AssetRequest) GetStatusOk() (*CatalogV1alpha1AssetRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogV1alpha1AssetRequest) SetStatus(v CatalogV1alpha1AssetRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogV1alpha1AssetRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguages

`func (o *CatalogV1alpha1AssetRequest) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CatalogV1alpha1AssetRequest) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CatalogV1alpha1AssetRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *CatalogV1alpha1AssetRequest) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *CatalogV1alpha1AssetRequest) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *CatalogV1alpha1AssetRequest) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *CatalogV1alpha1AssetRequest) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *CatalogV1alpha1AssetRequest) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *CatalogV1alpha1AssetRequest) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *CatalogV1alpha1AssetRequest) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *CatalogV1alpha1AssetRequest) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *CatalogV1alpha1AssetRequest) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *CatalogV1alpha1AssetRequest) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CatalogV1alpha1AssetRequest) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CatalogV1alpha1AssetRequest) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CatalogV1alpha1AssetRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


