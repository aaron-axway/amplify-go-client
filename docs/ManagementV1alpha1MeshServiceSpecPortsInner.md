# ManagementV1alpha1MeshServiceSpecPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Endpoints** | Pointer to [**[]ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner**](ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner.md) |  | [optional] 

## Methods

### NewManagementV1alpha1MeshServiceSpecPortsInner

`func NewManagementV1alpha1MeshServiceSpecPortsInner() *ManagementV1alpha1MeshServiceSpecPortsInner`

NewManagementV1alpha1MeshServiceSpecPortsInner instantiates a new ManagementV1alpha1MeshServiceSpecPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1MeshServiceSpecPortsInnerWithDefaults

`func NewManagementV1alpha1MeshServiceSpecPortsInnerWithDefaults() *ManagementV1alpha1MeshServiceSpecPortsInner`

NewManagementV1alpha1MeshServiceSpecPortsInnerWithDefaults instantiates a new ManagementV1alpha1MeshServiceSpecPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEndpoints

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetEndpoints() []ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) GetEndpointsOk() (*[]ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) SetEndpoints(v []ManagementV1alpha1MeshServiceSpecPortsInnerEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ManagementV1alpha1MeshServiceSpecPortsInner) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


