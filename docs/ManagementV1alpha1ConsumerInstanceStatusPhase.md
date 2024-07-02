# ManagementV1alpha1ConsumerInstanceStatusPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Level** | **string** | The criticality of the last update | 
**TransitionTime** | **time.Time** | Time of the current phase. | 
**Message** | Pointer to **string** | Description of the phase. | [optional] 

## Methods

### NewManagementV1alpha1ConsumerInstanceStatusPhase

`func NewManagementV1alpha1ConsumerInstanceStatusPhase(name string, level string, transitionTime time.Time, ) *ManagementV1alpha1ConsumerInstanceStatusPhase`

NewManagementV1alpha1ConsumerInstanceStatusPhase instantiates a new ManagementV1alpha1ConsumerInstanceStatusPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ConsumerInstanceStatusPhaseWithDefaults

`func NewManagementV1alpha1ConsumerInstanceStatusPhaseWithDefaults() *ManagementV1alpha1ConsumerInstanceStatusPhase`

NewManagementV1alpha1ConsumerInstanceStatusPhaseWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceStatusPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetTransitionTime

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetTransitionTime() time.Time`

GetTransitionTime returns the TransitionTime field if non-nil, zero value otherwise.

### GetTransitionTimeOk

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetTransitionTimeOk() (*time.Time, bool)`

GetTransitionTimeOk returns a tuple with the TransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionTime

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetTransitionTime(v time.Time)`

SetTransitionTime sets TransitionTime field to given value.


### GetMessage

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


