# ManagementV1alpha1AccessRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServiceInstance** | **string** | The name of an APIServiceInstance resource that specifies where the API is deployed. | 
**ManagedApplication** | **string** | The name of an ManagedApplication resource that groups set of credentials. | 
**Data** | **map[string]interface{}** | The value that matches the AccessRequestDefinition schema linked to the referenced APIServiceInstance. | 
**Quota** | Pointer to [**ManagementV1alpha1AccessRequestSpecQuota**](ManagementV1alpha1AccessRequestSpecQuota.md) |  | [optional] 
**AccessRequest** | Pointer to **string** | The AccessRequest from which this resource is being migrated from. Reference must be linked to the same Application and under the same Environment. | [optional] 

## Methods

### NewManagementV1alpha1AccessRequestSpec

`func NewManagementV1alpha1AccessRequestSpec(apiServiceInstance string, managedApplication string, data map[string]interface{}, ) *ManagementV1alpha1AccessRequestSpec`

NewManagementV1alpha1AccessRequestSpec instantiates a new ManagementV1alpha1AccessRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessRequestSpecWithDefaults

`func NewManagementV1alpha1AccessRequestSpecWithDefaults() *ManagementV1alpha1AccessRequestSpec`

NewManagementV1alpha1AccessRequestSpecWithDefaults instantiates a new ManagementV1alpha1AccessRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServiceInstance

`func (o *ManagementV1alpha1AccessRequestSpec) GetApiServiceInstance() string`

GetApiServiceInstance returns the ApiServiceInstance field if non-nil, zero value otherwise.

### GetApiServiceInstanceOk

`func (o *ManagementV1alpha1AccessRequestSpec) GetApiServiceInstanceOk() (*string, bool)`

GetApiServiceInstanceOk returns a tuple with the ApiServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServiceInstance

`func (o *ManagementV1alpha1AccessRequestSpec) SetApiServiceInstance(v string)`

SetApiServiceInstance sets ApiServiceInstance field to given value.


### GetManagedApplication

`func (o *ManagementV1alpha1AccessRequestSpec) GetManagedApplication() string`

GetManagedApplication returns the ManagedApplication field if non-nil, zero value otherwise.

### GetManagedApplicationOk

`func (o *ManagementV1alpha1AccessRequestSpec) GetManagedApplicationOk() (*string, bool)`

GetManagedApplicationOk returns a tuple with the ManagedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedApplication

`func (o *ManagementV1alpha1AccessRequestSpec) SetManagedApplication(v string)`

SetManagedApplication sets ManagedApplication field to given value.


### GetData

`func (o *ManagementV1alpha1AccessRequestSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagementV1alpha1AccessRequestSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagementV1alpha1AccessRequestSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetQuota

`func (o *ManagementV1alpha1AccessRequestSpec) GetQuota() ManagementV1alpha1AccessRequestSpecQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ManagementV1alpha1AccessRequestSpec) GetQuotaOk() (*ManagementV1alpha1AccessRequestSpecQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ManagementV1alpha1AccessRequestSpec) SetQuota(v ManagementV1alpha1AccessRequestSpecQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *ManagementV1alpha1AccessRequestSpec) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetAccessRequest

`func (o *ManagementV1alpha1AccessRequestSpec) GetAccessRequest() string`

GetAccessRequest returns the AccessRequest field if non-nil, zero value otherwise.

### GetAccessRequestOk

`func (o *ManagementV1alpha1AccessRequestSpec) GetAccessRequestOk() (*string, bool)`

GetAccessRequestOk returns a tuple with the AccessRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequest

`func (o *ManagementV1alpha1AccessRequestSpec) SetAccessRequest(v string)`

SetAccessRequest sets AccessRequest field to given value.

### HasAccessRequest

`func (o *ManagementV1alpha1AccessRequestSpec) HasAccessRequest() bool`

HasAccessRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


