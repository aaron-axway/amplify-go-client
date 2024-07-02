# ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to **[]string** |  | [optional] 
**Subset** | Pointer to **map[string]string** | A subset is the set of labels one or more workloads has outside the labels in the service selector. | [optional] 
**Specs** | Pointer to **[]string** | Array of references to apispecs discovered by the discovery agent.  | [optional] 

## Methods

### NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner

`func NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner() *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner`

NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner instantiates a new ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInnerWithDefaults

`func NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInnerWithDefaults() *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner`

NewManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInnerWithDefaults instantiates a new ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetSubset

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetSubset() map[string]string`

GetSubset returns the Subset field if non-nil, zero value otherwise.

### GetSubsetOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetSubsetOk() (*map[string]string, bool)`

GetSubsetOk returns a tuple with the Subset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubset

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) SetSubset(v map[string]string)`

SetSubset sets Subset field to given value.

### HasSubset

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) HasSubset() bool`

HasSubset returns a boolean if a field has been set.

### GetSpecs

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetSpecs() []string`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) GetSpecsOk() (*[]string, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) SetSpecs(v []string)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


