# ManagementV1alpha1DiscoveryAgentSpecConfigOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the owner. Defaults to team if not present. | [optional] [default to "team"]
**Id** | **string** | Id of the owner of the resource. | 

## Methods

### NewManagementV1alpha1DiscoveryAgentSpecConfigOwner

`func NewManagementV1alpha1DiscoveryAgentSpecConfigOwner(id string, ) *ManagementV1alpha1DiscoveryAgentSpecConfigOwner`

NewManagementV1alpha1DiscoveryAgentSpecConfigOwner instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfigOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DiscoveryAgentSpecConfigOwnerWithDefaults

`func NewManagementV1alpha1DiscoveryAgentSpecConfigOwnerWithDefaults() *ManagementV1alpha1DiscoveryAgentSpecConfigOwner`

NewManagementV1alpha1DiscoveryAgentSpecConfigOwnerWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentSpecConfigOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementV1alpha1DiscoveryAgentSpecConfigOwner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


