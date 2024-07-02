# ManagementV1alpha1ConsumerInstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Maps to the name of the Catalog Item. If not provided, the resource title will be used. | [optional] 
**ApiServiceInstance** | Pointer to **string** | The name of an APIServiceInstance resource that specifies where the API is deployed. | [optional] 
**OwningTeam** | Pointer to **string** | Name of the team that owns the Catalog Item. If not provided, the Default team will be used. | [optional] 
**Description** | Pointer to **string** | Maps to the description of the Catalog Item. Defaults to the API service description. | [optional] 
**Visibility** | Pointer to **string** | The visibility of the Catalog Item:  * PUBLIC - If Catalog Item is in PUBLISHED state, it will be visible to the entire organization.  * RESTRICTED - If Catalog Item is in PUBLISHED state, it will be visible to the owning team and teams part of the Catalog Item Access Control List.  | [optional] 
**Version** | Pointer to **string** | Version name for the revision. | [optional] 
**State** | Pointer to **string** | Catalog Item state:  * UNPUBLISHED - Only developers in the owning team will be able to access the Catalog Item.  * PUBLISHED - Developers and Consumers in the owning team and any additional team in the ACL will be able to access the Catalog Item.  | [optional] 
**Status** | Pointer to **string** | A way to communicate the status of the service to consumers. Examples: DRAFT, BETA, GA | [optional] 
**Tags** | Pointer to **[]string** | List of tags to be set on the Catalog Item. Max allowed tags for the Catalog Item is 80. | [optional] 
**Icon** | Pointer to [**ManagementV1alpha1ConsumerInstanceSpecIcon**](ManagementV1alpha1ConsumerInstanceSpecIcon.md) |  | [optional] 
**Documentation** | Pointer to **string** | Markdown documentation for the Catalog Item documentation. | [optional] 
**UnstructuredDataProperties** | Pointer to [**ManagementV1alpha1ConsumerInstanceSpecUnstructuredDataProperties**](ManagementV1alpha1ConsumerInstanceSpecUnstructuredDataProperties.md) |  | [optional] 
**AdditionalDataProperties** | Pointer to [**ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties**](ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties.md) |  | [optional] 
**Subscription** | Pointer to [**ManagementV1alpha1ConsumerInstanceSpecSubscription**](ManagementV1alpha1ConsumerInstanceSpecSubscription.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 

## Methods

### NewManagementV1alpha1ConsumerInstanceSpec

`func NewManagementV1alpha1ConsumerInstanceSpec() *ManagementV1alpha1ConsumerInstanceSpec`

NewManagementV1alpha1ConsumerInstanceSpec instantiates a new ManagementV1alpha1ConsumerInstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerInstanceSpecWithDefaults

`func NewManagementV1alpha1ConsumerInstanceSpecWithDefaults() *ManagementV1alpha1ConsumerInstanceSpec`

NewManagementV1alpha1ConsumerInstanceSpecWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiServiceInstance

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetApiServiceInstance() string`

GetApiServiceInstance returns the ApiServiceInstance field if non-nil, zero value otherwise.

### GetApiServiceInstanceOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetApiServiceInstanceOk() (*string, bool)`

GetApiServiceInstanceOk returns a tuple with the ApiServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServiceInstance

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetApiServiceInstance(v string)`

SetApiServiceInstance sets ApiServiceInstance field to given value.

### HasApiServiceInstance

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasApiServiceInstance() bool`

HasApiServiceInstance returns a boolean if a field has been set.

### GetOwningTeam

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetOwningTeam() string`

GetOwningTeam returns the OwningTeam field if non-nil, zero value otherwise.

### GetOwningTeamOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetOwningTeamOk() (*string, bool)`

GetOwningTeamOk returns a tuple with the OwningTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningTeam

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetOwningTeam(v string)`

SetOwningTeam sets OwningTeam field to given value.

### HasOwningTeam

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasOwningTeam() bool`

HasOwningTeam returns a boolean if a field has been set.

### GetDescription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetVersion

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetState

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIcon

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetIcon() ManagementV1alpha1ConsumerInstanceSpecIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetIconOk() (*ManagementV1alpha1ConsumerInstanceSpecIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetIcon(v ManagementV1alpha1ConsumerInstanceSpecIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetDocumentation

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetUnstructuredDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetUnstructuredDataProperties() ManagementV1alpha1ConsumerInstanceSpecUnstructuredDataProperties`

GetUnstructuredDataProperties returns the UnstructuredDataProperties field if non-nil, zero value otherwise.

### GetUnstructuredDataPropertiesOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetUnstructuredDataPropertiesOk() (*ManagementV1alpha1ConsumerInstanceSpecUnstructuredDataProperties, bool)`

GetUnstructuredDataPropertiesOk returns a tuple with the UnstructuredDataProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstructuredDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetUnstructuredDataProperties(v ManagementV1alpha1ConsumerInstanceSpecUnstructuredDataProperties)`

SetUnstructuredDataProperties sets UnstructuredDataProperties field to given value.

### HasUnstructuredDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasUnstructuredDataProperties() bool`

HasUnstructuredDataProperties returns a boolean if a field has been set.

### GetAdditionalDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetAdditionalDataProperties() ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties`

GetAdditionalDataProperties returns the AdditionalDataProperties field if non-nil, zero value otherwise.

### GetAdditionalDataPropertiesOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetAdditionalDataPropertiesOk() (*ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties, bool)`

GetAdditionalDataPropertiesOk returns a tuple with the AdditionalDataProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetAdditionalDataProperties(v ManagementV1alpha1ConsumerInstanceSpecAdditionalDataProperties)`

SetAdditionalDataProperties sets AdditionalDataProperties field to given value.

### HasAdditionalDataProperties

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasAdditionalDataProperties() bool`

HasAdditionalDataProperties returns a boolean if a field has been set.

### GetSubscription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetSubscription() ManagementV1alpha1ConsumerInstanceSpecSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetSubscriptionOk() (*ManagementV1alpha1ConsumerInstanceSpecSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetSubscription(v ManagementV1alpha1ConsumerInstanceSpecSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetCategories

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ManagementV1alpha1ConsumerInstanceSpec) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ManagementV1alpha1ConsumerInstanceSpec) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ManagementV1alpha1ConsumerInstanceSpec) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


