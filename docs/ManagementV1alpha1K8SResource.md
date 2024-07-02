# ManagementV1alpha1K8SResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "management"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind K8SResource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "K8SResource"]
**Name** | Pointer to **string** | Name of the K8SResource. K8SResource name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the K8SResource. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**ManagementV1alpha1K8SResourceSpec**](ManagementV1alpha1K8SResourceSpec.md) |  | 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewManagementV1alpha1K8SResource

`func NewManagementV1alpha1K8SResource(spec ManagementV1alpha1K8SResourceSpec, ) *ManagementV1alpha1K8SResource`

NewManagementV1alpha1K8SResource instantiates a new ManagementV1alpha1K8SResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1K8SResourceWithDefaults

`func NewManagementV1alpha1K8SResourceWithDefaults() *ManagementV1alpha1K8SResource`

NewManagementV1alpha1K8SResourceWithDefaults instantiates a new ManagementV1alpha1K8SResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ManagementV1alpha1K8SResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1K8SResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1K8SResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagementV1alpha1K8SResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *ManagementV1alpha1K8SResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ManagementV1alpha1K8SResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ManagementV1alpha1K8SResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ManagementV1alpha1K8SResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1K8SResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1K8SResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1K8SResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ManagementV1alpha1K8SResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1K8SResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1K8SResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1K8SResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1K8SResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ManagementV1alpha1K8SResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ManagementV1alpha1K8SResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ManagementV1alpha1K8SResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ManagementV1alpha1K8SResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagementV1alpha1K8SResource) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagementV1alpha1K8SResource) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagementV1alpha1K8SResource) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagementV1alpha1K8SResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1K8SResource) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1K8SResource) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1K8SResource) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManagementV1alpha1K8SResource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *ManagementV1alpha1K8SResource) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *ManagementV1alpha1K8SResource) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *ManagementV1alpha1K8SResource) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *ManagementV1alpha1K8SResource) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagementV1alpha1K8SResource) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagementV1alpha1K8SResource) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagementV1alpha1K8SResource) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagementV1alpha1K8SResource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *ManagementV1alpha1K8SResource) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementV1alpha1K8SResource) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementV1alpha1K8SResource) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementV1alpha1K8SResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *ManagementV1alpha1K8SResource) GetSpec() ManagementV1alpha1K8SResourceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ManagementV1alpha1K8SResource) GetSpecOk() (*ManagementV1alpha1K8SResourceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ManagementV1alpha1K8SResource) SetSpec(v ManagementV1alpha1K8SResourceSpec)`

SetSpec sets Spec field to given value.


### GetLanguages

`func (o *ManagementV1alpha1K8SResource) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ManagementV1alpha1K8SResource) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ManagementV1alpha1K8SResource) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ManagementV1alpha1K8SResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *ManagementV1alpha1K8SResource) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *ManagementV1alpha1K8SResource) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *ManagementV1alpha1K8SResource) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *ManagementV1alpha1K8SResource) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *ManagementV1alpha1K8SResource) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *ManagementV1alpha1K8SResource) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *ManagementV1alpha1K8SResource) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *ManagementV1alpha1K8SResource) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *ManagementV1alpha1K8SResource) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *ManagementV1alpha1K8SResource) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *ManagementV1alpha1K8SResource) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *ManagementV1alpha1K8SResource) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *ManagementV1alpha1K8SResource) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *ManagementV1alpha1K8SResource) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *ManagementV1alpha1K8SResource) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *ManagementV1alpha1K8SResource) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *ManagementV1alpha1K8SResource) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ManagementV1alpha1K8SResource) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ManagementV1alpha1K8SResource) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ManagementV1alpha1K8SResource) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


