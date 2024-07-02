# CatalogV1alpha1SubscriptionInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Resource belongs to the catalog group. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "catalog"]
**ApiVersion** | Pointer to **string** | Resource version is v1alpha1. The version defines the structure of the resource. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "v1alpha1"]
**Kind** | Pointer to **string** | Resource of kind SubscriptionInvoice. Cannot be updated. The server infers this from the endpoint the client submits the request to. | [optional] [readonly] [default to "SubscriptionInvoice"]
**Name** | Pointer to **string** | Name of the SubscriptionInvoice. SubscriptionInvoice name is unique and cannot be updated. The characters allowed in names are: digits (0-9), lower case letters (a-z), -, and . | [optional] 
**Title** | Pointer to **string** | The friendly name of the SubscriptionInvoice. | [optional] 
**Metadata** | Pointer to [**ApiV1Metadata**](ApiV1Metadata.md) |  | [optional] 
**Owner** | Pointer to [**ApiV1Owner**](ApiV1Owner.md) |  | [optional] 
**Finalizers** | Pointer to [**[]ApiV1Finalizer**](ApiV1Finalizer.md) | List of finalizers | [optional] 
**Attributes** | Pointer to **map[string]string** | Custom attributes added to objects. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. | [optional] 
**Spec** | [**CatalogV1alpha1SubscriptionInvoiceSpec**](CatalogV1alpha1SubscriptionInvoiceSpec.md) |  | 
**Marketplace** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceMarketplace**](CatalogV1alpha1SubscriptionInvoiceMarketplace.md) |  | [optional] 
**State** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceState**](CatalogV1alpha1SubscriptionInvoiceState.md) |  | [optional] 
**Billing** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceBilling**](CatalogV1alpha1SubscriptionInvoiceBilling.md) |  | [optional] 
**Status** | Pointer to [**CatalogV1alpha1SubscriptionInvoiceStatus**](CatalogV1alpha1SubscriptionInvoiceStatus.md) |  | [optional] 
**Languages** | Pointer to [**ApiV1Languages**](ApiV1Languages.md) |  | [optional] 
**LanguagesEnUs** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesPtBr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesDeDe** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**LanguagesFrFr** | Pointer to [**ApiV1Language**](ApiV1Language.md) |  | [optional] 
**Embedded** | Pointer to [**ApiV1Embedded**](ApiV1Embedded.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1SubscriptionInvoice

`func NewCatalogV1alpha1SubscriptionInvoice(spec CatalogV1alpha1SubscriptionInvoiceSpec, ) *CatalogV1alpha1SubscriptionInvoice`

NewCatalogV1alpha1SubscriptionInvoice instantiates a new CatalogV1alpha1SubscriptionInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1SubscriptionInvoiceWithDefaults

`func NewCatalogV1alpha1SubscriptionInvoiceWithDefaults() *CatalogV1alpha1SubscriptionInvoice`

NewCatalogV1alpha1SubscriptionInvoiceWithDefaults instantiates a new CatalogV1alpha1SubscriptionInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogV1alpha1SubscriptionInvoice) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogV1alpha1SubscriptionInvoice) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CatalogV1alpha1SubscriptionInvoice) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetApiVersion

`func (o *CatalogV1alpha1SubscriptionInvoice) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CatalogV1alpha1SubscriptionInvoice) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CatalogV1alpha1SubscriptionInvoice) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1alpha1SubscriptionInvoice) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1alpha1SubscriptionInvoice) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogV1alpha1SubscriptionInvoice) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CatalogV1alpha1SubscriptionInvoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1SubscriptionInvoice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogV1alpha1SubscriptionInvoice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *CatalogV1alpha1SubscriptionInvoice) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogV1alpha1SubscriptionInvoice) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CatalogV1alpha1SubscriptionInvoice) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogV1alpha1SubscriptionInvoice) GetMetadata() ApiV1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetMetadataOk() (*ApiV1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogV1alpha1SubscriptionInvoice) SetMetadata(v ApiV1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogV1alpha1SubscriptionInvoice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *CatalogV1alpha1SubscriptionInvoice) GetOwner() ApiV1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetOwnerOk() (*ApiV1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CatalogV1alpha1SubscriptionInvoice) SetOwner(v ApiV1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CatalogV1alpha1SubscriptionInvoice) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFinalizers

`func (o *CatalogV1alpha1SubscriptionInvoice) GetFinalizers() []ApiV1Finalizer`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetFinalizersOk() (*[]ApiV1Finalizer, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *CatalogV1alpha1SubscriptionInvoice) SetFinalizers(v []ApiV1Finalizer)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *CatalogV1alpha1SubscriptionInvoice) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetAttributes

`func (o *CatalogV1alpha1SubscriptionInvoice) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogV1alpha1SubscriptionInvoice) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CatalogV1alpha1SubscriptionInvoice) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *CatalogV1alpha1SubscriptionInvoice) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogV1alpha1SubscriptionInvoice) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogV1alpha1SubscriptionInvoice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSpec

`func (o *CatalogV1alpha1SubscriptionInvoice) GetSpec() CatalogV1alpha1SubscriptionInvoiceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetSpecOk() (*CatalogV1alpha1SubscriptionInvoiceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CatalogV1alpha1SubscriptionInvoice) SetSpec(v CatalogV1alpha1SubscriptionInvoiceSpec)`

SetSpec sets Spec field to given value.


### GetMarketplace

`func (o *CatalogV1alpha1SubscriptionInvoice) GetMarketplace() CatalogV1alpha1SubscriptionInvoiceMarketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetMarketplaceOk() (*CatalogV1alpha1SubscriptionInvoiceMarketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *CatalogV1alpha1SubscriptionInvoice) SetMarketplace(v CatalogV1alpha1SubscriptionInvoiceMarketplace)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *CatalogV1alpha1SubscriptionInvoice) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetState

`func (o *CatalogV1alpha1SubscriptionInvoice) GetState() CatalogV1alpha1SubscriptionInvoiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetStateOk() (*CatalogV1alpha1SubscriptionInvoiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1alpha1SubscriptionInvoice) SetState(v CatalogV1alpha1SubscriptionInvoiceState)`

SetState sets State field to given value.

### HasState

`func (o *CatalogV1alpha1SubscriptionInvoice) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBilling

`func (o *CatalogV1alpha1SubscriptionInvoice) GetBilling() CatalogV1alpha1SubscriptionInvoiceBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetBillingOk() (*CatalogV1alpha1SubscriptionInvoiceBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *CatalogV1alpha1SubscriptionInvoice) SetBilling(v CatalogV1alpha1SubscriptionInvoiceBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *CatalogV1alpha1SubscriptionInvoice) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogV1alpha1SubscriptionInvoice) GetStatus() CatalogV1alpha1SubscriptionInvoiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetStatusOk() (*CatalogV1alpha1SubscriptionInvoiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogV1alpha1SubscriptionInvoice) SetStatus(v CatalogV1alpha1SubscriptionInvoiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogV1alpha1SubscriptionInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLanguages

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguages() ApiV1Languages`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesOk() (*ApiV1Languages, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CatalogV1alpha1SubscriptionInvoice) SetLanguages(v ApiV1Languages)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CatalogV1alpha1SubscriptionInvoice) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLanguagesEnUs

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesEnUs() ApiV1Language`

GetLanguagesEnUs returns the LanguagesEnUs field if non-nil, zero value otherwise.

### GetLanguagesEnUsOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesEnUsOk() (*ApiV1Language, bool)`

GetLanguagesEnUsOk returns a tuple with the LanguagesEnUs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesEnUs

`func (o *CatalogV1alpha1SubscriptionInvoice) SetLanguagesEnUs(v ApiV1Language)`

SetLanguagesEnUs sets LanguagesEnUs field to given value.

### HasLanguagesEnUs

`func (o *CatalogV1alpha1SubscriptionInvoice) HasLanguagesEnUs() bool`

HasLanguagesEnUs returns a boolean if a field has been set.

### GetLanguagesPtBr

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesPtBr() ApiV1Language`

GetLanguagesPtBr returns the LanguagesPtBr field if non-nil, zero value otherwise.

### GetLanguagesPtBrOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesPtBrOk() (*ApiV1Language, bool)`

GetLanguagesPtBrOk returns a tuple with the LanguagesPtBr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesPtBr

`func (o *CatalogV1alpha1SubscriptionInvoice) SetLanguagesPtBr(v ApiV1Language)`

SetLanguagesPtBr sets LanguagesPtBr field to given value.

### HasLanguagesPtBr

`func (o *CatalogV1alpha1SubscriptionInvoice) HasLanguagesPtBr() bool`

HasLanguagesPtBr returns a boolean if a field has been set.

### GetLanguagesDeDe

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesDeDe() ApiV1Language`

GetLanguagesDeDe returns the LanguagesDeDe field if non-nil, zero value otherwise.

### GetLanguagesDeDeOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesDeDeOk() (*ApiV1Language, bool)`

GetLanguagesDeDeOk returns a tuple with the LanguagesDeDe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesDeDe

`func (o *CatalogV1alpha1SubscriptionInvoice) SetLanguagesDeDe(v ApiV1Language)`

SetLanguagesDeDe sets LanguagesDeDe field to given value.

### HasLanguagesDeDe

`func (o *CatalogV1alpha1SubscriptionInvoice) HasLanguagesDeDe() bool`

HasLanguagesDeDe returns a boolean if a field has been set.

### GetLanguagesFrFr

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesFrFr() ApiV1Language`

GetLanguagesFrFr returns the LanguagesFrFr field if non-nil, zero value otherwise.

### GetLanguagesFrFrOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetLanguagesFrFrOk() (*ApiV1Language, bool)`

GetLanguagesFrFrOk returns a tuple with the LanguagesFrFr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesFrFr

`func (o *CatalogV1alpha1SubscriptionInvoice) SetLanguagesFrFr(v ApiV1Language)`

SetLanguagesFrFr sets LanguagesFrFr field to given value.

### HasLanguagesFrFr

`func (o *CatalogV1alpha1SubscriptionInvoice) HasLanguagesFrFr() bool`

HasLanguagesFrFr returns a boolean if a field has been set.

### GetEmbedded

`func (o *CatalogV1alpha1SubscriptionInvoice) GetEmbedded() ApiV1Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CatalogV1alpha1SubscriptionInvoice) GetEmbeddedOk() (*ApiV1Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CatalogV1alpha1SubscriptionInvoice) SetEmbedded(v ApiV1Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CatalogV1alpha1SubscriptionInvoice) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


