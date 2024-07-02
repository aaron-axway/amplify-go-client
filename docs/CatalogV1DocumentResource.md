# CatalogV1DocumentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the catalog group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "catalog"]
**ApiVersion** | Pointer to **string** | Resource version is v1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1"]
**Kind** | Pointer to **string** | Resource of kind DocumentResource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "DocumentResource"]
**Name** | Pointer to **string** | Name of the DocumentResource. DocumentResource name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the DocumentResource. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**CatalogV1DocumentResourceSpec**](CatalogV1DocumentResourceSpec.md) |  | 
**Icon** | Pointer to **string** | Image for the file expressed as a Data-Uri (https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs) | [optional] 
**State** | Pointer to [**CatalogV1DocumentResourceState**](CatalogV1DocumentResourceState.md) |  | [optional] 
**References** | Pointer to [**CatalogV1DocumentResourceReferences**](CatalogV1DocumentResourceReferences.md) |  | [optional] 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewCatalogV1DocumentResource

`func NewCatalogV1DocumentResource(spec CatalogV1DocumentResourceSpec, ) *CatalogV1DocumentResource`

NewCatalogV1DocumentResource instantiates a new CatalogV1DocumentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1DocumentResourceWithDefaults

`func NewCatalogV1DocumentResourceWithDefaults() *CatalogV1DocumentResource`

NewCatalogV1DocumentResourceWithDefaults instantiates a new CatalogV1DocumentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogV1DocumentResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogV1DocumentResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogV1DocumentResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CatalogV1DocumentResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *CatalogV1DocumentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CatalogV1DocumentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CatalogV1DocumentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CatalogV1DocumentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1DocumentResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1DocumentResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1DocumentResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogV1DocumentResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CatalogV1DocumentResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1DocumentResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1DocumentResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1DocumentResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *CatalogV1DocumentResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1DocumentResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1DocumentResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogV1DocumentResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogV1DocumentResource) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogV1DocumentResource) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogV1DocumentResource) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogV1DocumentResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *CatalogV1DocumentResource) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CatalogV1DocumentResource) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CatalogV1DocumentResource) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CatalogV1DocumentResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *CatalogV1DocumentResource) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CatalogV1DocumentResource) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CatalogV1DocumentResource) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CatalogV1DocumentResource) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *CatalogV1DocumentResource) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogV1DocumentResource) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogV1DocumentResource) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CatalogV1DocumentResource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *CatalogV1DocumentResource) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogV1DocumentResource) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogV1DocumentResource) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogV1DocumentResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *CatalogV1DocumentResource) GetSpec() CatalogV1DocumentResourceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CatalogV1DocumentResource) GetSpecOk() (*CatalogV1DocumentResourceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CatalogV1DocumentResource) SetSpec(v CatalogV1DocumentResourceSpec)`

SetSpec sets Spec field to given value.


### GetIcon

`func (o *CatalogV1DocumentResource) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CatalogV1DocumentResource) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CatalogV1DocumentResource) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CatalogV1DocumentResource) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetState

`func (o *CatalogV1DocumentResource) GetState() CatalogV1DocumentResourceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1DocumentResource) GetStateOk() (*CatalogV1DocumentResourceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1DocumentResource) SetState(v CatalogV1DocumentResourceState)`

SetState sets State field to given value.

### HasState

`func (o *CatalogV1DocumentResource) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReferences

`func (o *CatalogV1DocumentResource) GetReferences() CatalogV1DocumentResourceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CatalogV1DocumentResource) GetReferencesOk() (*CatalogV1DocumentResourceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CatalogV1DocumentResource) SetReferences(v CatalogV1DocumentResourceReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CatalogV1DocumentResource) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetLanguages

`func (o *CatalogV1DocumentResource) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CatalogV1DocumentResource) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CatalogV1DocumentResource) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CatalogV1DocumentResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *CatalogV1DocumentResource) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *CatalogV1DocumentResource) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *CatalogV1DocumentResource) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *CatalogV1DocumentResource) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *CatalogV1DocumentResource) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *CatalogV1DocumentResource) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *CatalogV1DocumentResource) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *CatalogV1DocumentResource) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *CatalogV1DocumentResource) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *CatalogV1DocumentResource) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *CatalogV1DocumentResource) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *CatalogV1DocumentResource) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *CatalogV1DocumentResource) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *CatalogV1DocumentResource) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *CatalogV1DocumentResource) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *CatalogV1DocumentResource) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *CatalogV1DocumentResource) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CatalogV1DocumentResource) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CatalogV1DocumentResource) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CatalogV1DocumentResource) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


