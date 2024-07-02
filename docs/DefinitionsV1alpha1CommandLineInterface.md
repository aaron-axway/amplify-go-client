# DefinitionsV1alpha1CommandLineInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the definitions group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "definitions"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind CommandLineInterface. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "CommandLineInterface"]
**Name** | Pointer to **string** | Name of the CommandLineInterface. CommandLineInterface name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the CommandLineInterface. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**DefinitionsV1alpha1CommandLineInterfaceSpec**](DefinitionsV1alpha1CommandLineInterfaceSpec.md) |  | 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewDefinitionsV1alpha1CommandLineInterface

`func NewDefinitionsV1alpha1CommandLineInterface(spec DefinitionsV1alpha1CommandLineInterfaceSpec, ) *DefinitionsV1alpha1CommandLineInterface`

NewDefinitionsV1alpha1CommandLineInterface instantiates a new DefinitionsV1alpha1CommandLineInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1CommandLineInterfaceWithDefaults

`func NewDefinitionsV1alpha1CommandLineInterfaceWithDefaults() *DefinitionsV1alpha1CommandLineInterface`

NewDefinitionsV1alpha1CommandLineInterfaceWithDefaults instantiates a new DefinitionsV1alpha1CommandLineInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DefinitionsV1alpha1CommandLineInterface) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DefinitionsV1alpha1CommandLineInterface) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DefinitionsV1alpha1CommandLineInterface) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *DefinitionsV1alpha1CommandLineInterface) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DefinitionsV1alpha1CommandLineInterface) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DefinitionsV1alpha1CommandLineInterface) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *DefinitionsV1alpha1CommandLineInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DefinitionsV1alpha1CommandLineInterface) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DefinitionsV1alpha1CommandLineInterface) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DefinitionsV1alpha1CommandLineInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefinitionsV1alpha1CommandLineInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefinitionsV1alpha1CommandLineInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *DefinitionsV1alpha1CommandLineInterface) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DefinitionsV1alpha1CommandLineInterface) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DefinitionsV1alpha1CommandLineInterface) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *DefinitionsV1alpha1CommandLineInterface) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DefinitionsV1alpha1CommandLineInterface) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DefinitionsV1alpha1CommandLineInterface) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *DefinitionsV1alpha1CommandLineInterface) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DefinitionsV1alpha1CommandLineInterface) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DefinitionsV1alpha1CommandLineInterface) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *DefinitionsV1alpha1CommandLineInterface) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *DefinitionsV1alpha1CommandLineInterface) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *DefinitionsV1alpha1CommandLineInterface) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *DefinitionsV1alpha1CommandLineInterface) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DefinitionsV1alpha1CommandLineInterface) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DefinitionsV1alpha1CommandLineInterface) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *DefinitionsV1alpha1CommandLineInterface) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DefinitionsV1alpha1CommandLineInterface) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DefinitionsV1alpha1CommandLineInterface) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *DefinitionsV1alpha1CommandLineInterface) GetSpec() DefinitionsV1alpha1CommandLineInterfaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetSpecOk() (*DefinitionsV1alpha1CommandLineInterfaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DefinitionsV1alpha1CommandLineInterface) SetSpec(v DefinitionsV1alpha1CommandLineInterfaceSpec)`

SetSpec sets Spec field to given value.


### GetLanguages

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *DefinitionsV1alpha1CommandLineInterface) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *DefinitionsV1alpha1CommandLineInterface) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *DefinitionsV1alpha1CommandLineInterface) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *DefinitionsV1alpha1CommandLineInterface) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *DefinitionsV1alpha1CommandLineInterface) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *DefinitionsV1alpha1CommandLineInterface) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *DefinitionsV1alpha1CommandLineInterface) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *DefinitionsV1alpha1CommandLineInterface) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *DefinitionsV1alpha1CommandLineInterface) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *DefinitionsV1alpha1CommandLineInterface) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *DefinitionsV1alpha1CommandLineInterface) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DefinitionsV1alpha1CommandLineInterface) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DefinitionsV1alpha1CommandLineInterface) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DefinitionsV1alpha1CommandLineInterface) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


