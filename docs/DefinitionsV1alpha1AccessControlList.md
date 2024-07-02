# DefinitionsV1alpha1AccessControlList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the definitions group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "definitions"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind AccessControlList. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "AccessControlList"]
**Name** | Pointer to **string** | Name of the AccessControlList. AccessControlList name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the AccessControlList. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**DefinitionsV1alpha1AccessControlListSpec**](DefinitionsV1alpha1AccessControlListSpec.md) |  | 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewDefinitionsV1alpha1AccessControlList

`func NewDefinitionsV1alpha1AccessControlList(spec DefinitionsV1alpha1AccessControlListSpec, ) *DefinitionsV1alpha1AccessControlList`

NewDefinitionsV1alpha1AccessControlList instantiates a new DefinitionsV1alpha1AccessControlList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1AccessControlListWithDefaults

`func NewDefinitionsV1alpha1AccessControlListWithDefaults() *DefinitionsV1alpha1AccessControlList`

NewDefinitionsV1alpha1AccessControlListWithDefaults instantiates a new DefinitionsV1alpha1AccessControlList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DefinitionsV1alpha1AccessControlList) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DefinitionsV1alpha1AccessControlList) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DefinitionsV1alpha1AccessControlList) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DefinitionsV1alpha1AccessControlList) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *DefinitionsV1alpha1AccessControlList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DefinitionsV1alpha1AccessControlList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DefinitionsV1alpha1AccessControlList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DefinitionsV1alpha1AccessControlList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *DefinitionsV1alpha1AccessControlList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DefinitionsV1alpha1AccessControlList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DefinitionsV1alpha1AccessControlList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DefinitionsV1alpha1AccessControlList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DefinitionsV1alpha1AccessControlList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefinitionsV1alpha1AccessControlList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefinitionsV1alpha1AccessControlList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefinitionsV1alpha1AccessControlList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *DefinitionsV1alpha1AccessControlList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DefinitionsV1alpha1AccessControlList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DefinitionsV1alpha1AccessControlList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DefinitionsV1alpha1AccessControlList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *DefinitionsV1alpha1AccessControlList) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DefinitionsV1alpha1AccessControlList) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DefinitionsV1alpha1AccessControlList) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DefinitionsV1alpha1AccessControlList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *DefinitionsV1alpha1AccessControlList) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DefinitionsV1alpha1AccessControlList) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DefinitionsV1alpha1AccessControlList) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DefinitionsV1alpha1AccessControlList) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *DefinitionsV1alpha1AccessControlList) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *DefinitionsV1alpha1AccessControlList) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *DefinitionsV1alpha1AccessControlList) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *DefinitionsV1alpha1AccessControlList) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *DefinitionsV1alpha1AccessControlList) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DefinitionsV1alpha1AccessControlList) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DefinitionsV1alpha1AccessControlList) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DefinitionsV1alpha1AccessControlList) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *DefinitionsV1alpha1AccessControlList) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DefinitionsV1alpha1AccessControlList) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DefinitionsV1alpha1AccessControlList) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DefinitionsV1alpha1AccessControlList) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *DefinitionsV1alpha1AccessControlList) GetSpec() DefinitionsV1alpha1AccessControlListSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DefinitionsV1alpha1AccessControlList) GetSpecOk() (*DefinitionsV1alpha1AccessControlListSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DefinitionsV1alpha1AccessControlList) SetSpec(v DefinitionsV1alpha1AccessControlListSpec)`

SetSpec sets Spec field to given value.


### GetLanguages

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *DefinitionsV1alpha1AccessControlList) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *DefinitionsV1alpha1AccessControlList) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *DefinitionsV1alpha1AccessControlList) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *DefinitionsV1alpha1AccessControlList) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *DefinitionsV1alpha1AccessControlList) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *DefinitionsV1alpha1AccessControlList) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *DefinitionsV1alpha1AccessControlList) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *DefinitionsV1alpha1AccessControlList) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *DefinitionsV1alpha1AccessControlList) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *DefinitionsV1alpha1AccessControlList) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *DefinitionsV1alpha1AccessControlList) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *DefinitionsV1alpha1AccessControlList) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DefinitionsV1alpha1AccessControlList) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DefinitionsV1alpha1AccessControlList) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DefinitionsV1alpha1AccessControlList) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


