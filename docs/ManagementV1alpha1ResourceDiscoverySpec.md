# ManagementV1alpha1ResourceDiscoverySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**NamespaceFilter** | Pointer to [**ManagementV1alpha1ResourceDiscoverySpecNamespaceFilter**](ManagementV1alpha1ResourceDiscoverySpecNamespaceFilter.md) |  | [optional] 
**ResourceFilter** | Pointer to [**ManagementV1alpha1ResourceDiscoverySpecResourceFilter**](ManagementV1alpha1ResourceDiscoverySpecResourceFilter.md) |  | [optional] 
**KeepSpecFields** | Pointer to **[]string** |  | [optional] 
**KeepStatusFields** | Pointer to **[]string** |  | [optional] 
**IgnoreLabels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ExtraAttributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementV1alpha1ResourceDiscoverySpec

`func NewManagementV1alpha1ResourceDiscoverySpec() *ManagementV1alpha1ResourceDiscoverySpec`

NewManagementV1alpha1ResourceDiscoverySpec instantiates a new ManagementV1alpha1ResourceDiscoverySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ResourceDiscoverySpecWithDefaults

`func NewManagementV1alpha1ResourceDiscoverySpecWithDefaults() *ManagementV1alpha1ResourceDiscoverySpec`

NewManagementV1alpha1ResourceDiscoverySpecWithDefaults instantiates a new ManagementV1alpha1ResourceDiscoverySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetGroup

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetNamespaceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetNamespaceFilter() ManagementV1alpha1ResourceDiscoverySpecNamespaceFilter`

GetNamespaceFilter returns the NamespaceFilter field if non-nil, zero value otherwise.

### GetNamespaceFilterOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetNamespaceFilterOk() (*ManagementV1alpha1ResourceDiscoverySpecNamespaceFilter, bool)`

GetNamespaceFilterOk returns a tuple with the NamespaceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetNamespaceFilter(v ManagementV1alpha1ResourceDiscoverySpecNamespaceFilter)`

SetNamespaceFilter sets NamespaceFilter field to given value.

### HasNamespaceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasNamespaceFilter() bool`

HasNamespaceFilter returns a boolean if a field has been set.

### GetResourceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetResourceFilter() ManagementV1alpha1ResourceDiscoverySpecResourceFilter`

GetResourceFilter returns the ResourceFilter field if non-nil, zero value otherwise.

### GetResourceFilterOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetResourceFilterOk() (*ManagementV1alpha1ResourceDiscoverySpecResourceFilter, bool)`

GetResourceFilterOk returns a tuple with the ResourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetResourceFilter(v ManagementV1alpha1ResourceDiscoverySpecResourceFilter)`

SetResourceFilter sets ResourceFilter field to given value.

### HasResourceFilter

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasResourceFilter() bool`

HasResourceFilter returns a boolean if a field has been set.

### GetKeepSpecFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKeepSpecFields() []string`

GetKeepSpecFields returns the KeepSpecFields field if non-nil, zero value otherwise.

### GetKeepSpecFieldsOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKeepSpecFieldsOk() (*[]string, bool)`

GetKeepSpecFieldsOk returns a tuple with the KeepSpecFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepSpecFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetKeepSpecFields(v []string)`

SetKeepSpecFields sets KeepSpecFields field to given value.

### HasKeepSpecFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasKeepSpecFields() bool`

HasKeepSpecFields returns a boolean if a field has been set.

### GetKeepStatusFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKeepStatusFields() []string`

GetKeepStatusFields returns the KeepStatusFields field if non-nil, zero value otherwise.

### GetKeepStatusFieldsOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetKeepStatusFieldsOk() (*[]string, bool)`

GetKeepStatusFieldsOk returns a tuple with the KeepStatusFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepStatusFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetKeepStatusFields(v []string)`

SetKeepStatusFields sets KeepStatusFields field to given value.

### HasKeepStatusFields

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasKeepStatusFields() bool`

HasKeepStatusFields returns a boolean if a field has been set.

### GetIgnoreLabels

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetIgnoreLabels() []string`

GetIgnoreLabels returns the IgnoreLabels field if non-nil, zero value otherwise.

### GetIgnoreLabelsOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetIgnoreLabelsOk() (*[]string, bool)`

GetIgnoreLabelsOk returns a tuple with the IgnoreLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreLabels

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetIgnoreLabels(v []string)`

SetIgnoreLabels sets IgnoreLabels field to given value.

### HasIgnoreLabels

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasIgnoreLabels() bool`

HasIgnoreLabels returns a boolean if a field has been set.

### GetTags

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExtraAttributes

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetExtraAttributes() map[string]string`

GetExtraAttributes returns the ExtraAttributes field if non-nil, zero value otherwise.

### GetExtraAttributesOk

`func (o *ManagementV1alpha1ResourceDiscoverySpec) GetExtraAttributesOk() (*map[string]string, bool)`

GetExtraAttributesOk returns a tuple with the ExtraAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAttributes

`func (o *ManagementV1alpha1ResourceDiscoverySpec) SetExtraAttributes(v map[string]string)`

SetExtraAttributes sets ExtraAttributes field to given value.

### HasExtraAttributes

`func (o *ManagementV1alpha1ResourceDiscoverySpec) HasExtraAttributes() bool`

HasExtraAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


