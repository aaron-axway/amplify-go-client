# ManagementV1alpha1ResourceHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "management"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind ResourceHook. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "ResourceHook"]
**Name** | Pointer to **string** | Name of the ResourceHook. ResourceHook name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the ResourceHook. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**ManagementV1alpha1ResourceHookSpec**](ManagementV1alpha1ResourceHookSpec.md) |  | 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewManagementV1alpha1ResourceHook

`func NewManagementV1alpha1ResourceHook(spec ManagementV1alpha1ResourceHookSpec, ) *ManagementV1alpha1ResourceHook`

NewManagementV1alpha1ResourceHook instantiates a new ManagementV1alpha1ResourceHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ResourceHookWithDefaults

`func NewManagementV1alpha1ResourceHookWithDefaults() *ManagementV1alpha1ResourceHook`

NewManagementV1alpha1ResourceHookWithDefaults instantiates a new ManagementV1alpha1ResourceHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ManagementV1alpha1ResourceHook) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1ResourceHook) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1ResourceHook) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagementV1alpha1ResourceHook) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *ManagementV1alpha1ResourceHook) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ManagementV1alpha1ResourceHook) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ManagementV1alpha1ResourceHook) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ManagementV1alpha1ResourceHook) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1ResourceHook) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1ResourceHook) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1ResourceHook) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ManagementV1alpha1ResourceHook) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1ResourceHook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1ResourceHook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1ResourceHook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1ResourceHook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ManagementV1alpha1ResourceHook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ManagementV1alpha1ResourceHook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ManagementV1alpha1ResourceHook) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ManagementV1alpha1ResourceHook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagementV1alpha1ResourceHook) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagementV1alpha1ResourceHook) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagementV1alpha1ResourceHook) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagementV1alpha1ResourceHook) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1ResourceHook) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1ResourceHook) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1ResourceHook) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManagementV1alpha1ResourceHook) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *ManagementV1alpha1ResourceHook) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *ManagementV1alpha1ResourceHook) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *ManagementV1alpha1ResourceHook) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *ManagementV1alpha1ResourceHook) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagementV1alpha1ResourceHook) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagementV1alpha1ResourceHook) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagementV1alpha1ResourceHook) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagementV1alpha1ResourceHook) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *ManagementV1alpha1ResourceHook) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementV1alpha1ResourceHook) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementV1alpha1ResourceHook) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementV1alpha1ResourceHook) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *ManagementV1alpha1ResourceHook) GetSpec() ManagementV1alpha1ResourceHookSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ManagementV1alpha1ResourceHook) GetSpecOk() (*ManagementV1alpha1ResourceHookSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ManagementV1alpha1ResourceHook) SetSpec(v ManagementV1alpha1ResourceHookSpec)`

SetSpec sets Spec field to given value.


### GetLanguages

`func (o *ManagementV1alpha1ResourceHook) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ManagementV1alpha1ResourceHook) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ManagementV1alpha1ResourceHook) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *ManagementV1alpha1ResourceHook) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *ManagementV1alpha1ResourceHook) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *ManagementV1alpha1ResourceHook) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *ManagementV1alpha1ResourceHook) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *ManagementV1alpha1ResourceHook) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *ManagementV1alpha1ResourceHook) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *ManagementV1alpha1ResourceHook) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *ManagementV1alpha1ResourceHook) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *ManagementV1alpha1ResourceHook) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *ManagementV1alpha1ResourceHook) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ManagementV1alpha1ResourceHook) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ManagementV1alpha1ResourceHook) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ManagementV1alpha1ResourceHook) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


