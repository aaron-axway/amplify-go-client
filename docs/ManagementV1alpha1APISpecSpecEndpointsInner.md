# ManagementV1alpha1APISpecSpecEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the endpoint. | [optional] 
**Host** | Pointer to **string** | The host of the pod &lt;namespace&gt;.&lt;pod-name&gt; | [optional] 
**Port** | Pointer to **string** | The port of the pod on which this spec was discovered. | [optional] 
**Path** | Pointer to **string** | The path on which this spec was discovered. | [optional] 
**Labels** | Pointer to **map[string]string** | The labels for the pod this spec was discovered on. | [optional] 

## Methods

### NewManagementV1alpha1APISpecSpecEndpointsInner

`func NewManagementV1alpha1APISpecSpecEndpointsInner() *ManagementV1alpha1APISpecSpecEndpointsInner`

NewManagementV1alpha1APISpecSpecEndpointsInner instantiates a new ManagementV1alpha1APISpecSpecEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APISpecSpecEndpointsInnerWithDefaults

`func NewManagementV1alpha1APISpecSpecEndpointsInnerWithDefaults() *ManagementV1alpha1APISpecSpecEndpointsInner`

NewManagementV1alpha1APISpecSpecEndpointsInnerWithDefaults instantiates a new ManagementV1alpha1APISpecSpecEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHost

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLabels

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ManagementV1alpha1APISpecSpecEndpointsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


