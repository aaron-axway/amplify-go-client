# ManagementV1alpha1APISpecSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecDiscoveryRef** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**ManagementV1alpha1APISpecSpecDefinition**](ManagementV1alpha1APISpecSpecDefinition.md) |  | [optional] 
**Endpoints** | Pointer to [**[]ManagementV1alpha1APISpecSpecEndpointsInner**](ManagementV1alpha1APISpecSpecEndpointsInner.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APISpecSpec

`func NewManagementV1alpha1APISpecSpec() *ManagementV1alpha1APISpecSpec`

NewManagementV1alpha1APISpecSpec instantiates a new ManagementV1alpha1APISpecSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APISpecSpecWithDefaults

`func NewManagementV1alpha1APISpecSpecWithDefaults() *ManagementV1alpha1APISpecSpec`

NewManagementV1alpha1APISpecSpecWithDefaults instantiates a new ManagementV1alpha1APISpecSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecDiscoveryRef

`func (o *ManagementV1alpha1APISpecSpec) GetSpecDiscoveryRef() string`

GetSpecDiscoveryRef returns the SpecDiscoveryRef field if non-nil, zero value otherwise.

### GetSpecDiscoveryRefOk

`func (o *ManagementV1alpha1APISpecSpec) GetSpecDiscoveryRefOk() (*string, bool)`

GetSpecDiscoveryRefOk returns a tuple with the SpecDiscoveryRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecDiscoveryRef

`func (o *ManagementV1alpha1APISpecSpec) SetSpecDiscoveryRef(v string)`

SetSpecDiscoveryRef sets SpecDiscoveryRef field to given value.

### HasSpecDiscoveryRef

`func (o *ManagementV1alpha1APISpecSpec) HasSpecDiscoveryRef() bool`

HasSpecDiscoveryRef returns a boolean if a field has been set.

### GetDefinition

`func (o *ManagementV1alpha1APISpecSpec) GetDefinition() ManagementV1alpha1APISpecSpecDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ManagementV1alpha1APISpecSpec) GetDefinitionOk() (*ManagementV1alpha1APISpecSpecDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ManagementV1alpha1APISpecSpec) SetDefinition(v ManagementV1alpha1APISpecSpecDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ManagementV1alpha1APISpecSpec) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetEndpoints

`func (o *ManagementV1alpha1APISpecSpec) GetEndpoints() []ManagementV1alpha1APISpecSpecEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ManagementV1alpha1APISpecSpec) GetEndpointsOk() (*[]ManagementV1alpha1APISpecSpecEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ManagementV1alpha1APISpecSpec) SetEndpoints(v []ManagementV1alpha1APISpecSpecEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ManagementV1alpha1APISpecSpec) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


