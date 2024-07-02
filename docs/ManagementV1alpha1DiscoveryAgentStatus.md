# ManagementV1alpha1DiscoveryAgentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version name for the agent revision. | [optional] 
**LatestAvailableVersion** | Pointer to **string** | Latest available version for the agent revision. | [optional] 
**State** | Pointer to **string** | Agent status:  * running - Passed all health checks.  Up and running  * stopped - Agent is not running  * failed - Failed health checks  * unhealthy - Agent is running with health check failure  | [optional] 
**PreviousState** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** | A way to communicate details about the current status by the agent | [optional] 
**LastActivityTime** | Pointer to **time.Time** | The last updated event timestamp provided by the agent | [optional] 
**SdkVersion** | Pointer to **string** | Version name for the SDK revision. | [optional] 

## Methods

### NewManagementV1alpha1DiscoveryAgentStatus

`func NewManagementV1alpha1DiscoveryAgentStatus() *ManagementV1alpha1DiscoveryAgentStatus`

NewManagementV1alpha1DiscoveryAgentStatus instantiates a new ManagementV1alpha1DiscoveryAgentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DiscoveryAgentStatusWithDefaults

`func NewManagementV1alpha1DiscoveryAgentStatusWithDefaults() *ManagementV1alpha1DiscoveryAgentStatus`

NewManagementV1alpha1DiscoveryAgentStatusWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLatestAvailableVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetLatestAvailableVersion() string`

GetLatestAvailableVersion returns the LatestAvailableVersion field if non-nil, zero value otherwise.

### GetLatestAvailableVersionOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetLatestAvailableVersionOk() (*string, bool)`

GetLatestAvailableVersionOk returns a tuple with the LatestAvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAvailableVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetLatestAvailableVersion(v string)`

SetLatestAvailableVersion sets LatestAvailableVersion field to given value.

### HasLatestAvailableVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasLatestAvailableVersion() bool`

HasLatestAvailableVersion returns a boolean if a field has been set.

### GetState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPreviousState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetPreviousState() string`

GetPreviousState returns the PreviousState field if non-nil, zero value otherwise.

### GetPreviousStateOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetPreviousStateOk() (*string, bool)`

GetPreviousStateOk returns a tuple with the PreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetPreviousState(v string)`

SetPreviousState sets PreviousState field to given value.

### HasPreviousState

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasPreviousState() bool`

HasPreviousState returns a boolean if a field has been set.

### GetMessage

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastActivityTime

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetLastActivityTime() time.Time`

GetLastActivityTime returns the LastActivityTime field if non-nil, zero value otherwise.

### GetLastActivityTimeOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetLastActivityTimeOk() (*time.Time, bool)`

GetLastActivityTimeOk returns a tuple with the LastActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityTime

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetLastActivityTime(v time.Time)`

SetLastActivityTime sets LastActivityTime field to given value.

### HasLastActivityTime

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasLastActivityTime() bool`

HasLastActivityTime returns a boolean if a field has been set.

### GetSdkVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ManagementV1alpha1DiscoveryAgentStatus) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ManagementV1alpha1DiscoveryAgentStatus) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


