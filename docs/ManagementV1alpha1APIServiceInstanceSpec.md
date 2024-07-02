# ManagementV1alpha1APIServiceInstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServiceRevision** | Pointer to **string** |  | [optional] 
**AccessRequestDefinition** | Pointer to **string** |  | [optional] 
**CredentialRequestDefinitions** | Pointer to **[]string** |  | [optional] 
**Endpoint** | Pointer to [**[]ManagementV1alpha1APIServiceInstanceSpecEndpointInner**](ManagementV1alpha1APIServiceInstanceSpecEndpointInner.md) | A list of locations where the api is deployed. If \&quot;mock\&quot; property is set, then the mock endpoints will be assigned by the system.  | [optional] 
**Mock** | Pointer to [**ManagementV1alpha1APIServiceInstanceSpecMock**](ManagementV1alpha1APIServiceInstanceSpecMock.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceSpec

`func NewManagementV1alpha1APIServiceInstanceSpec() *ManagementV1alpha1APIServiceInstanceSpec`

NewManagementV1alpha1APIServiceInstanceSpec instantiates a new ManagementV1alpha1APIServiceInstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceSpecWithDefaults

`func NewManagementV1alpha1APIServiceInstanceSpecWithDefaults() *ManagementV1alpha1APIServiceInstanceSpec`

NewManagementV1alpha1APIServiceInstanceSpecWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServiceRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetApiServiceRevision() string`

GetApiServiceRevision returns the ApiServiceRevision field if non-nil, zero value otherwise.

### GetApiServiceRevisionOk

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetApiServiceRevisionOk() (*string, bool)`

GetApiServiceRevisionOk returns a tuple with the ApiServiceRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServiceRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpec) SetApiServiceRevision(v string)`

SetApiServiceRevision sets ApiServiceRevision field to given value.

### HasApiServiceRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpec) HasApiServiceRevision() bool`

HasApiServiceRevision returns a boolean if a field has been set.

### GetAccessRequestDefinition

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetAccessRequestDefinition() string`

GetAccessRequestDefinition returns the AccessRequestDefinition field if non-nil, zero value otherwise.

### GetAccessRequestDefinitionOk

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetAccessRequestDefinitionOk() (*string, bool)`

GetAccessRequestDefinitionOk returns a tuple with the AccessRequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestDefinition

`func (o *ManagementV1alpha1APIServiceInstanceSpec) SetAccessRequestDefinition(v string)`

SetAccessRequestDefinition sets AccessRequestDefinition field to given value.

### HasAccessRequestDefinition

`func (o *ManagementV1alpha1APIServiceInstanceSpec) HasAccessRequestDefinition() bool`

HasAccessRequestDefinition returns a boolean if a field has been set.

### GetCredentialRequestDefinitions

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetCredentialRequestDefinitions() []string`

GetCredentialRequestDefinitions returns the CredentialRequestDefinitions field if non-nil, zero value otherwise.

### GetCredentialRequestDefinitionsOk

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetCredentialRequestDefinitionsOk() (*[]string, bool)`

GetCredentialRequestDefinitionsOk returns a tuple with the CredentialRequestDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRequestDefinitions

`func (o *ManagementV1alpha1APIServiceInstanceSpec) SetCredentialRequestDefinitions(v []string)`

SetCredentialRequestDefinitions sets CredentialRequestDefinitions field to given value.

### HasCredentialRequestDefinitions

`func (o *ManagementV1alpha1APIServiceInstanceSpec) HasCredentialRequestDefinitions() bool`

HasCredentialRequestDefinitions returns a boolean if a field has been set.

### GetEndpoint

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetEndpoint() []ManagementV1alpha1APIServiceInstanceSpecEndpointInner`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetEndpointOk() (*[]ManagementV1alpha1APIServiceInstanceSpecEndpointInner, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ManagementV1alpha1APIServiceInstanceSpec) SetEndpoint(v []ManagementV1alpha1APIServiceInstanceSpecEndpointInner)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ManagementV1alpha1APIServiceInstanceSpec) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetMock

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetMock() ManagementV1alpha1APIServiceInstanceSpecMock`

GetMock returns the Mock field if non-nil, zero value otherwise.

### GetMockOk

`func (o *ManagementV1alpha1APIServiceInstanceSpec) GetMockOk() (*ManagementV1alpha1APIServiceInstanceSpecMock, bool)`

GetMockOk returns a tuple with the Mock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMock

`func (o *ManagementV1alpha1APIServiceInstanceSpec) SetMock(v ManagementV1alpha1APIServiceInstanceSpecMock)`

SetMock sets Mock field to given value.

### HasMock

`func (o *ManagementV1alpha1APIServiceInstanceSpec) HasMock() bool`

HasMock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


