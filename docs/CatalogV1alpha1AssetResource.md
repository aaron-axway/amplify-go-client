# CatalogV1alpha1AssetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the catalog group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "catalog"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind AssetResource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "AssetResource"]
**Name** | Pointer to **string** | Name of the AssetResource. AssetResource name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the AssetResource. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**CatalogV1alpha1AssetResourceSpec**](CatalogV1alpha1AssetResourceSpec.md) |  | 
**References** | Pointer to [**CatalogV1alpha1AssetResourceReferences**](CatalogV1alpha1AssetResourceReferences.md) |  | [optional] 
**Releasehash** | Pointer to **string** | Hash of particular AssetResource properties which would require a new asset release if changed. | [optional] 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1AssetResource

`func NewCatalogV1alpha1AssetResource(spec CatalogV1alpha1AssetResourceSpec, ) *CatalogV1alpha1AssetResource`

NewCatalogV1alpha1AssetResource instantiates a new CatalogV1alpha1AssetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AssetResourceWithDefaults

`func NewCatalogV1alpha1AssetResourceWithDefaults() *CatalogV1alpha1AssetResource`

NewCatalogV1alpha1AssetResourceWithDefaults instantiates a new CatalogV1alpha1AssetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogV1alpha1AssetResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogV1alpha1AssetResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogV1alpha1AssetResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CatalogV1alpha1AssetResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *CatalogV1alpha1AssetResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CatalogV1alpha1AssetResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CatalogV1alpha1AssetResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CatalogV1alpha1AssetResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1alpha1AssetResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1alpha1AssetResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1alpha1AssetResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogV1alpha1AssetResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CatalogV1alpha1AssetResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1AssetResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1AssetResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1alpha1AssetResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *CatalogV1alpha1AssetResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1alpha1AssetResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1alpha1AssetResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogV1alpha1AssetResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogV1alpha1AssetResource) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogV1alpha1AssetResource) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogV1alpha1AssetResource) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogV1alpha1AssetResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *CatalogV1alpha1AssetResource) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CatalogV1alpha1AssetResource) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CatalogV1alpha1AssetResource) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CatalogV1alpha1AssetResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *CatalogV1alpha1AssetResource) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CatalogV1alpha1AssetResource) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CatalogV1alpha1AssetResource) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CatalogV1alpha1AssetResource) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *CatalogV1alpha1AssetResource) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogV1alpha1AssetResource) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogV1alpha1AssetResource) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CatalogV1alpha1AssetResource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *CatalogV1alpha1AssetResource) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogV1alpha1AssetResource) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogV1alpha1AssetResource) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogV1alpha1AssetResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *CatalogV1alpha1AssetResource) GetSpec() CatalogV1alpha1AssetResourceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CatalogV1alpha1AssetResource) GetSpecOk() (*CatalogV1alpha1AssetResourceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CatalogV1alpha1AssetResource) SetSpec(v CatalogV1alpha1AssetResourceSpec)`

SetSpec sets Spec field to given value.


### GetReferences

`func (o *CatalogV1alpha1AssetResource) GetReferences() CatalogV1alpha1AssetResourceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CatalogV1alpha1AssetResource) GetReferencesOk() (*CatalogV1alpha1AssetResourceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CatalogV1alpha1AssetResource) SetReferences(v CatalogV1alpha1AssetResourceReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CatalogV1alpha1AssetResource) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetReleasehash

`func (o *CatalogV1alpha1AssetResource) GetReleasehash() string`

GetReleasehash returns the Releasehash field if non-nil, zero value otherwise.

### GetReleasehashOk

`func (o *CatalogV1alpha1AssetResource) GetReleasehashOk() (*string, bool)`

GetReleasehashOk returns a tuple with the Releasehash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasehash

`func (o *CatalogV1alpha1AssetResource) SetReleasehash(v string)`

SetReleasehash sets Releasehash field to given value.

### HasReleasehash

`func (o *CatalogV1alpha1AssetResource) HasReleasehash() bool`

HasReleasehash returns a boolean if a field has been set.

### GetLanguages

`func (o *CatalogV1alpha1AssetResource) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CatalogV1alpha1AssetResource) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CatalogV1alpha1AssetResource) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CatalogV1alpha1AssetResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *CatalogV1alpha1AssetResource) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *CatalogV1alpha1AssetResource) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *CatalogV1alpha1AssetResource) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *CatalogV1alpha1AssetResource) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *CatalogV1alpha1AssetResource) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *CatalogV1alpha1AssetResource) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *CatalogV1alpha1AssetResource) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *CatalogV1alpha1AssetResource) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *CatalogV1alpha1AssetResource) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *CatalogV1alpha1AssetResource) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *CatalogV1alpha1AssetResource) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *CatalogV1alpha1AssetResource) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *CatalogV1alpha1AssetResource) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *CatalogV1alpha1AssetResource) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *CatalogV1alpha1AssetResource) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *CatalogV1alpha1AssetResource) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *CatalogV1alpha1AssetResource) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CatalogV1alpha1AssetResource) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CatalogV1alpha1AssetResource) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CatalogV1alpha1AssetResource) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


