# ManagementV1alpha1K8SResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDiscovery** | Pointer to **string** |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**OwnerReferences** | Pointer to [**[]ManagementV1alpha1K8SResourceSpecOwnerReferencesInner**](ManagementV1alpha1K8SResourceSpecOwnerReferencesInner.md) |  | [optional] 
**ResourceSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourceStatus** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManagementV1alpha1K8SResourceSpec

`func NewManagementV1alpha1K8SResourceSpec() *ManagementV1alpha1K8SResourceSpec`

NewManagementV1alpha1K8SResourceSpec instantiates a new ManagementV1alpha1K8SResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1K8SResourceSpecWithDefaults

`func NewManagementV1alpha1K8SResourceSpecWithDefaults() *ManagementV1alpha1K8SResourceSpec`

NewManagementV1alpha1K8SResourceSpecWithDefaults instantiates a new ManagementV1alpha1K8SResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDiscovery

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceDiscovery() string`

GetResourceDiscovery returns the ResourceDiscovery field if non-nil, zero value otherwise.

### GetResourceDiscoveryOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceDiscoveryOk() (*string, bool)`

GetResourceDiscoveryOk returns a tuple with the ResourceDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDiscovery

`func (o *ManagementV1alpha1K8SResourceSpec) SetResourceDiscovery(v string)`

SetResourceDiscovery sets ResourceDiscovery field to given value.

### HasResourceDiscovery

`func (o *ManagementV1alpha1K8SResourceSpec) HasResourceDiscovery() bool`

HasResourceDiscovery returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ManagementV1alpha1K8SResourceSpec) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ManagementV1alpha1K8SResourceSpec) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *ManagementV1alpha1K8SResourceSpec) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ManagementV1alpha1K8SResourceSpec) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ManagementV1alpha1K8SResourceSpec) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetOwnerReferences

`func (o *ManagementV1alpha1K8SResourceSpec) GetOwnerReferences() []ManagementV1alpha1K8SResourceSpecOwnerReferencesInner`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetOwnerReferencesOk() (*[]ManagementV1alpha1K8SResourceSpecOwnerReferencesInner, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *ManagementV1alpha1K8SResourceSpec) SetOwnerReferences(v []ManagementV1alpha1K8SResourceSpecOwnerReferencesInner)`

SetOwnerReferences sets OwnerReferences field to given value.

### HasOwnerReferences

`func (o *ManagementV1alpha1K8SResourceSpec) HasOwnerReferences() bool`

HasOwnerReferences returns a boolean if a field has been set.

### GetResourceSpec

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceSpec() map[string]interface{}`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceSpecOk() (*map[string]interface{}, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *ManagementV1alpha1K8SResourceSpec) SetResourceSpec(v map[string]interface{})`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *ManagementV1alpha1K8SResourceSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetResourceStatus

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceStatus() map[string]interface{}`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *ManagementV1alpha1K8SResourceSpec) GetResourceStatusOk() (*map[string]interface{}, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *ManagementV1alpha1K8SResourceSpec) SetResourceStatus(v map[string]interface{})`

SetResourceStatus sets ResourceStatus field to given value.

### HasResourceStatus

`func (o *ManagementV1alpha1K8SResourceSpec) HasResourceStatus() bool`

HasResourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


