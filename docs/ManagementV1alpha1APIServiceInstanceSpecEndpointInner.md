# ManagementV1alpha1APIServiceInstanceSpecEndpointInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | **string** |  | 
**Routing** | Pointer to [**ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting**](ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceSpecEndpointInner

`func NewManagementV1alpha1APIServiceInstanceSpecEndpointInner(host string, protocol string, ) *ManagementV1alpha1APIServiceInstanceSpecEndpointInner`

NewManagementV1alpha1APIServiceInstanceSpecEndpointInner instantiates a new ManagementV1alpha1APIServiceInstanceSpecEndpointInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerWithDefaults

`func NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerWithDefaults() *ManagementV1alpha1APIServiceInstanceSpecEndpointInner`

NewManagementV1alpha1APIServiceInstanceSpecEndpointInnerWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpecEndpointInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetRouting

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetRouting() ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) GetRoutingOk() (*ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) SetRouting(v ManagementV1alpha1APIServiceInstanceSpecEndpointInnerRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *ManagementV1alpha1APIServiceInstanceSpecEndpointInner) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


