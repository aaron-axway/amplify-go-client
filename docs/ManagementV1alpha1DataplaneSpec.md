# ManagementV1alpha1DataplaneSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The dataplane type that this agent connects to | 
**Config** | Pointer to [**ManagementV1alpha1DataplaneSpecConfig**](ManagementV1alpha1DataplaneSpecConfig.md) |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpec

`func NewManagementV1alpha1DataplaneSpec(type_ string, ) *ManagementV1alpha1DataplaneSpec`

NewManagementV1alpha1DataplaneSpec instantiates a new ManagementV1alpha1DataplaneSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecWithDefaults

`func NewManagementV1alpha1DataplaneSpecWithDefaults() *ManagementV1alpha1DataplaneSpec`

NewManagementV1alpha1DataplaneSpecWithDefaults instantiates a new ManagementV1alpha1DataplaneSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpec) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *ManagementV1alpha1DataplaneSpec) GetConfig() ManagementV1alpha1DataplaneSpecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ManagementV1alpha1DataplaneSpec) GetConfigOk() (*ManagementV1alpha1DataplaneSpecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ManagementV1alpha1DataplaneSpec) SetConfig(v ManagementV1alpha1DataplaneSpecConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ManagementV1alpha1DataplaneSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


