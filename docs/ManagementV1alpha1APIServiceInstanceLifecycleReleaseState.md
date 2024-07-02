# ManagementV1alpha1APIServiceInstanceLifecycleReleaseState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Current release state of the API endpoint(s) such as stable or deprecated. | 
**Message** | Pointer to **string** | Optional info to be associated with the current state. If state is deprecated, then this is intended to indicate when the servers will become decommissioned.  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceLifecycleReleaseState

`func NewManagementV1alpha1APIServiceInstanceLifecycleReleaseState(name string, ) *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState`

NewManagementV1alpha1APIServiceInstanceLifecycleReleaseState instantiates a new ManagementV1alpha1APIServiceInstanceLifecycleReleaseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceLifecycleReleaseStateWithDefaults

`func NewManagementV1alpha1APIServiceInstanceLifecycleReleaseStateWithDefaults() *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState`

NewManagementV1alpha1APIServiceInstanceLifecycleReleaseStateWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceLifecycleReleaseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManagementV1alpha1APIServiceInstanceLifecycleReleaseState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


