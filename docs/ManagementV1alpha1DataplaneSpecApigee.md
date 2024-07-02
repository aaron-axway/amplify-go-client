# ManagementV1alpha1DataplaneSpecApigee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ProjectId** | **string** | The Project ID on GCP that Apigee is configured in | 
**DeveloperEmail** | **string** | The Developer that will own all Apigee Applications created by the agent | 
**Mode** | Pointer to **string** | The discovery mode that the Apigee agents should use | [optional] 
**Environment** | Pointer to **string** | The environment for which the metrics are gathered | [optional] 
**MetricsFilter** | Pointer to [**ManagementV1alpha1DataplaneSpecApigeeMetricsFilter**](ManagementV1alpha1DataplaneSpecApigeeMetricsFilter.md) |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSpecApigee

`func NewManagementV1alpha1DataplaneSpecApigee(type_ string, projectId string, developerEmail string, ) *ManagementV1alpha1DataplaneSpecApigee`

NewManagementV1alpha1DataplaneSpecApigee instantiates a new ManagementV1alpha1DataplaneSpecApigee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSpecApigeeWithDefaults

`func NewManagementV1alpha1DataplaneSpecApigeeWithDefaults() *ManagementV1alpha1DataplaneSpecApigee`

NewManagementV1alpha1DataplaneSpecApigeeWithDefaults instantiates a new ManagementV1alpha1DataplaneSpecApigee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetType(v string)`

SetType sets Type field to given value.


### GetProjectId

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDeveloperEmail

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetDeveloperEmail() string`

GetDeveloperEmail returns the DeveloperEmail field if non-nil, zero value otherwise.

### GetDeveloperEmailOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetDeveloperEmailOk() (*string, bool)`

GetDeveloperEmailOk returns a tuple with the DeveloperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperEmail

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetDeveloperEmail(v string)`

SetDeveloperEmail sets DeveloperEmail field to given value.


### GetMode

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ManagementV1alpha1DataplaneSpecApigee) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEnvironment

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ManagementV1alpha1DataplaneSpecApigee) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetMetricsFilter() ManagementV1alpha1DataplaneSpecApigeeMetricsFilter`

GetMetricsFilter returns the MetricsFilter field if non-nil, zero value otherwise.

### GetMetricsFilterOk

`func (o *ManagementV1alpha1DataplaneSpecApigee) GetMetricsFilterOk() (*ManagementV1alpha1DataplaneSpecApigeeMetricsFilter, bool)`

GetMetricsFilterOk returns a tuple with the MetricsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecApigee) SetMetricsFilter(v ManagementV1alpha1DataplaneSpecApigeeMetricsFilter)`

SetMetricsFilter sets MetricsFilter field to given value.

### HasMetricsFilter

`func (o *ManagementV1alpha1DataplaneSpecApigee) HasMetricsFilter() bool`

HasMetricsFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


