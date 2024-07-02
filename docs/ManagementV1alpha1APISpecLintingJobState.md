# ManagementV1alpha1APISpecLintingJobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The current state, indicating progress towards consistency. | 
**Message** | Pointer to **string** | Details of the state. | [optional] 
**Timestamp** | Pointer to **time.Time** | Time when the update occurred. | [optional] 

## Methods

### NewManagementV1alpha1APISpecLintingJobState

`func NewManagementV1alpha1APISpecLintingJobState(name string, ) *ManagementV1alpha1APISpecLintingJobState`

NewManagementV1alpha1APISpecLintingJobState instantiates a new ManagementV1alpha1APISpecLintingJobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APISpecLintingJobStateWithDefaults

`func NewManagementV1alpha1APISpecLintingJobStateWithDefaults() *ManagementV1alpha1APISpecLintingJobState`

NewManagementV1alpha1APISpecLintingJobStateWithDefaults instantiates a new ManagementV1alpha1APISpecLintingJobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1APISpecLintingJobState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1APISpecLintingJobState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1APISpecLintingJobState) SetName(v string)`

SetName sets Name field to given value.


### GetMessage

`func (o *ManagementV1alpha1APISpecLintingJobState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManagementV1alpha1APISpecLintingJobState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManagementV1alpha1APISpecLintingJobState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManagementV1alpha1APISpecLintingJobState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *ManagementV1alpha1APISpecLintingJobState) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1APISpecLintingJobState) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1APISpecLintingJobState) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ManagementV1alpha1APISpecLintingJobState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


