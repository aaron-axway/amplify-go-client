# ManagementV1alpha1MeshServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to **map[string]string** | The labels used by the service to match workloads it exposes. | [optional] 
**Resource** | Pointer to **string** | Reference to the K8SResource derived from the Kubernetes Service. | [optional] 
**Workloads** | Pointer to **[]string** | References to workloads exposed by the service. | [optional] 
**Ports** | Pointer to [**[]ManagementV1alpha1MeshServiceSpecPortsInner**](ManagementV1alpha1MeshServiceSpecPortsInner.md) | Details per port. | [optional] 

## Methods

### NewManagementV1alpha1MeshServiceSpec

`func NewManagementV1alpha1MeshServiceSpec() *ManagementV1alpha1MeshServiceSpec`

NewManagementV1alpha1MeshServiceSpec instantiates a new ManagementV1alpha1MeshServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1MeshServiceSpecWithDefaults

`func NewManagementV1alpha1MeshServiceSpecWithDefaults() *ManagementV1alpha1MeshServiceSpec`

NewManagementV1alpha1MeshServiceSpecWithDefaults instantiates a new ManagementV1alpha1MeshServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *ManagementV1alpha1MeshServiceSpec) GetSelector() map[string]string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ManagementV1alpha1MeshServiceSpec) GetSelectorOk() (*map[string]string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ManagementV1alpha1MeshServiceSpec) SetSelector(v map[string]string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ManagementV1alpha1MeshServiceSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetResource

`func (o *ManagementV1alpha1MeshServiceSpec) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ManagementV1alpha1MeshServiceSpec) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ManagementV1alpha1MeshServiceSpec) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ManagementV1alpha1MeshServiceSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWorkloads

`func (o *ManagementV1alpha1MeshServiceSpec) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *ManagementV1alpha1MeshServiceSpec) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *ManagementV1alpha1MeshServiceSpec) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *ManagementV1alpha1MeshServiceSpec) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetPorts

`func (o *ManagementV1alpha1MeshServiceSpec) GetPorts() []ManagementV1alpha1MeshServiceSpecPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ManagementV1alpha1MeshServiceSpec) GetPortsOk() (*[]ManagementV1alpha1MeshServiceSpecPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ManagementV1alpha1MeshServiceSpec) SetPorts(v []ManagementV1alpha1MeshServiceSpecPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ManagementV1alpha1MeshServiceSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


