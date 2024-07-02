# ManagementV1alpha1Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the management group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "management"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind Credential. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "Credential"]
**Name** | Pointer to **string** | Name of the Credential. Credential name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the Credential. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**ManagementV1alpha1CredentialSpec**](ManagementV1alpha1CredentialSpec.md) |  | 
**References** | Pointer to [**ManagementV1alpha1CredentialReferences**](ManagementV1alpha1CredentialReferences.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | data matching the provisioning schema from CredentialRequestDefinition. | [optional] 
**State** | Pointer to [**ManagementV1alpha1CredentialState**](ManagementV1alpha1CredentialState.md) |  | [optional] 
**Policies** | Pointer to [**ManagementV1alpha1CredentialPolicies**](ManagementV1alpha1CredentialPolicies.md) |  | [optional] 
**Status** | Pointer to [**ManagementV1alpha1CredentialStatus**](ManagementV1alpha1CredentialStatus.md) |  | [optional] 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewManagementV1alpha1Credential

`func NewManagementV1alpha1Credential(spec ManagementV1alpha1CredentialSpec, ) *ManagementV1alpha1Credential`

NewManagementV1alpha1Credential instantiates a new ManagementV1alpha1Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialWithDefaults

`func NewManagementV1alpha1CredentialWithDefaults() *ManagementV1alpha1Credential`

NewManagementV1alpha1CredentialWithDefaults instantiates a new ManagementV1alpha1Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ManagementV1alpha1Credential) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1Credential) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1Credential) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagementV1alpha1Credential) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *ManagementV1alpha1Credential) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ManagementV1alpha1Credential) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ManagementV1alpha1Credential) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ManagementV1alpha1Credential) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1Credential) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1Credential) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1Credential) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ManagementV1alpha1Credential) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1Credential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1Credential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1Credential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1Credential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ManagementV1alpha1Credential) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ManagementV1alpha1Credential) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ManagementV1alpha1Credential) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ManagementV1alpha1Credential) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagementV1alpha1Credential) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagementV1alpha1Credential) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagementV1alpha1Credential) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagementV1alpha1Credential) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *ManagementV1alpha1Credential) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManagementV1alpha1Credential) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManagementV1alpha1Credential) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManagementV1alpha1Credential) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *ManagementV1alpha1Credential) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *ManagementV1alpha1Credential) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *ManagementV1alpha1Credential) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *ManagementV1alpha1Credential) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagementV1alpha1Credential) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagementV1alpha1Credential) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagementV1alpha1Credential) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagementV1alpha1Credential) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *ManagementV1alpha1Credential) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementV1alpha1Credential) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementV1alpha1Credential) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementV1alpha1Credential) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *ManagementV1alpha1Credential) GetSpec() ManagementV1alpha1CredentialSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ManagementV1alpha1Credential) GetSpecOk() (*ManagementV1alpha1CredentialSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ManagementV1alpha1Credential) SetSpec(v ManagementV1alpha1CredentialSpec)`

SetSpec sets Spec field to given value.


### GetReferences

`func (o *ManagementV1alpha1Credential) GetReferences() ManagementV1alpha1CredentialReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ManagementV1alpha1Credential) GetReferencesOk() (*ManagementV1alpha1CredentialReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ManagementV1alpha1Credential) SetReferences(v ManagementV1alpha1CredentialReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ManagementV1alpha1Credential) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetData

`func (o *ManagementV1alpha1Credential) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagementV1alpha1Credential) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagementV1alpha1Credential) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ManagementV1alpha1Credential) HasData() bool`

HasData returns a boolean if a field has been set.

### GetState

`func (o *ManagementV1alpha1Credential) GetState() ManagementV1alpha1CredentialState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagementV1alpha1Credential) GetStateOk() (*ManagementV1alpha1CredentialState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagementV1alpha1Credential) SetState(v ManagementV1alpha1CredentialState)`

SetState sets State field to given value.

### HasState

`func (o *ManagementV1alpha1Credential) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPolicies

`func (o *ManagementV1alpha1Credential) GetPolicies() ManagementV1alpha1CredentialPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ManagementV1alpha1Credential) GetPoliciesOk() (*ManagementV1alpha1CredentialPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ManagementV1alpha1Credential) SetPolicies(v ManagementV1alpha1CredentialPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ManagementV1alpha1Credential) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetStatus

`func (o *ManagementV1alpha1Credential) GetStatus() ManagementV1alpha1CredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagementV1alpha1Credential) GetStatusOk() (*ManagementV1alpha1CredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagementV1alpha1Credential) SetStatus(v ManagementV1alpha1CredentialStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagementV1alpha1Credential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguages

`func (o *ManagementV1alpha1Credential) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ManagementV1alpha1Credential) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ManagementV1alpha1Credential) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ManagementV1alpha1Credential) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *ManagementV1alpha1Credential) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *ManagementV1alpha1Credential) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *ManagementV1alpha1Credential) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *ManagementV1alpha1Credential) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *ManagementV1alpha1Credential) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *ManagementV1alpha1Credential) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *ManagementV1alpha1Credential) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *ManagementV1alpha1Credential) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *ManagementV1alpha1Credential) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *ManagementV1alpha1Credential) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *ManagementV1alpha1Credential) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *ManagementV1alpha1Credential) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *ManagementV1alpha1Credential) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *ManagementV1alpha1Credential) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *ManagementV1alpha1Credential) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *ManagementV1alpha1Credential) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *ManagementV1alpha1Credential) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ManagementV1alpha1Credential) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ManagementV1alpha1Credential) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ManagementV1alpha1Credential) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


