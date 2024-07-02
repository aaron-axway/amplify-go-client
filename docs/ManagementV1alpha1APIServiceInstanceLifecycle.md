# ManagementV1alpha1APIServiceInstanceLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | Pointer to **string** | Overrides the Environment&#39;s default stage. Value must be one of Environment&#39;s allowed stages. | [optional] 
**ReleaseState** | Pointer to [**ManagementV1alpha1APIServiceInstanceLifecycleReleaseState**](ManagementV1alpha1APIServiceInstanceLifecycleReleaseState.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceLifecycle

`func NewManagementV1alpha1APIServiceInstanceLifecycle() *ManagementV1alpha1APIServiceInstanceLifecycle`

NewManagementV1alpha1APIServiceInstanceLifecycle instantiates a new ManagementV1alpha1APIServiceInstanceLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceLifecycleWithDefaults

`func NewManagementV1alpha1APIServiceInstanceLifecycleWithDefaults() *ManagementV1alpha1APIServiceInstanceLifecycle`

NewManagementV1alpha1APIServiceInstanceLifecycleWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetReleaseState

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) GetReleaseState() ManagementV1alpha1APIServiceInstanceLifecycleReleaseState`

GetReleaseState returns the ReleaseState field if non-nil, zero value otherwise.

### GetReleaseStateOk

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) GetReleaseStateOk() (*ManagementV1alpha1APIServiceInstanceLifecycleReleaseState, bool)`

GetReleaseStateOk returns a tuple with the ReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseState

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) SetReleaseState(v ManagementV1alpha1APIServiceInstanceLifecycleReleaseState)`

SetReleaseState sets ReleaseState field to given value.

### HasReleaseState

`func (o *ManagementV1alpha1APIServiceInstanceLifecycle) HasReleaseState() bool`

HasReleaseState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


