# ManagementV1alpha1MeshWorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to **[]string** | References to k8sresources making up the workload. | [optional] 
**Labels** | Pointer to **map[string]string** | Labels from the k8sresources making up the workload. | [optional] 
**Ports** | Pointer to [**[]ManagementV1alpha1MeshWorkloadSpecPortsInner**](ManagementV1alpha1MeshWorkloadSpecPortsInner.md) | Discovered ports on the workload. | [optional] 

## Methods

### NewManagementV1alpha1MeshWorkloadSpec

`func NewManagementV1alpha1MeshWorkloadSpec() *ManagementV1alpha1MeshWorkloadSpec`

NewManagementV1alpha1MeshWorkloadSpec instantiates a new ManagementV1alpha1MeshWorkloadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1MeshWorkloadSpecWithDefaults

`func NewManagementV1alpha1MeshWorkloadSpecWithDefaults() *ManagementV1alpha1MeshWorkloadSpec`

NewManagementV1alpha1MeshWorkloadSpecWithDefaults instantiates a new ManagementV1alpha1MeshWorkloadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ManagementV1alpha1MeshWorkloadSpec) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ManagementV1alpha1MeshWorkloadSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLabels

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ManagementV1alpha1MeshWorkloadSpec) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ManagementV1alpha1MeshWorkloadSpec) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPorts

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetPorts() []ManagementV1alpha1MeshWorkloadSpecPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ManagementV1alpha1MeshWorkloadSpec) GetPortsOk() (*[]ManagementV1alpha1MeshWorkloadSpecPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ManagementV1alpha1MeshWorkloadSpec) SetPorts(v []ManagementV1alpha1MeshWorkloadSpecPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ManagementV1alpha1MeshWorkloadSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


