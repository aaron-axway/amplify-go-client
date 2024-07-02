# ManagementV1alpha1APIServiceRevisionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiService** | **string** |  | 
**Definition** | Pointer to [**ManagementV1alpha1APIServiceRevisionSpecDefinition**](ManagementV1alpha1APIServiceRevisionSpecDefinition.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceRevisionSpec

`func NewManagementV1alpha1APIServiceRevisionSpec(apiService string, ) *ManagementV1alpha1APIServiceRevisionSpec`

NewManagementV1alpha1APIServiceRevisionSpec instantiates a new ManagementV1alpha1APIServiceRevisionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceRevisionSpecWithDefaults

`func NewManagementV1alpha1APIServiceRevisionSpecWithDefaults() *ManagementV1alpha1APIServiceRevisionSpec`

NewManagementV1alpha1APIServiceRevisionSpecWithDefaults instantiates a new ManagementV1alpha1APIServiceRevisionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiService

`func (o *ManagementV1alpha1APIServiceRevisionSpec) GetApiService() string`

GetApiService returns the ApiService field if non-nil, zero value otherwise.

### GetApiServiceOk

`func (o *ManagementV1alpha1APIServiceRevisionSpec) GetApiServiceOk() (*string, bool)`

GetApiServiceOk returns a tuple with the ApiService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiService

`func (o *ManagementV1alpha1APIServiceRevisionSpec) SetApiService(v string)`

SetApiService sets ApiService field to given value.


### GetDefinition

`func (o *ManagementV1alpha1APIServiceRevisionSpec) GetDefinition() ManagementV1alpha1APIServiceRevisionSpecDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ManagementV1alpha1APIServiceRevisionSpec) GetDefinitionOk() (*ManagementV1alpha1APIServiceRevisionSpecDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ManagementV1alpha1APIServiceRevisionSpec) SetDefinition(v ManagementV1alpha1APIServiceRevisionSpecDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ManagementV1alpha1APIServiceRevisionSpec) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


