# ManagementV1alpha1EnvironmentStages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** | The default stage to be assigned to the Environment&#39;s APIServiceInstances. | [optional] 
**Allowed** | Pointer to **[]string** | The allowed stages that can be set as override on the Environment&#39;s APIServiceInstance. | [optional] 

## Methods

### NewManagementV1alpha1EnvironmentStages

`func NewManagementV1alpha1EnvironmentStages() *ManagementV1alpha1EnvironmentStages`

NewManagementV1alpha1EnvironmentStages instantiates a new ManagementV1alpha1EnvironmentStages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1EnvironmentStagesWithDefaults

`func NewManagementV1alpha1EnvironmentStagesWithDefaults() *ManagementV1alpha1EnvironmentStages`

NewManagementV1alpha1EnvironmentStagesWithDefaults instantiates a new ManagementV1alpha1EnvironmentStages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ManagementV1alpha1EnvironmentStages) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ManagementV1alpha1EnvironmentStages) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ManagementV1alpha1EnvironmentStages) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ManagementV1alpha1EnvironmentStages) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetAllowed

`func (o *ManagementV1alpha1EnvironmentStages) GetAllowed() []string`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ManagementV1alpha1EnvironmentStages) GetAllowedOk() (*[]string, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ManagementV1alpha1EnvironmentStages) SetAllowed(v []string)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ManagementV1alpha1EnvironmentStages) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


