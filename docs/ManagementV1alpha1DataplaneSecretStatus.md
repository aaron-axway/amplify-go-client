# ManagementV1alpha1DataplaneSecretStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]ManagementV1alpha1DataplaneSecretStatusReasonsInner**](ManagementV1alpha1DataplaneSecretStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSecretStatus

`func NewManagementV1alpha1DataplaneSecretStatus(level string, ) *ManagementV1alpha1DataplaneSecretStatus`

NewManagementV1alpha1DataplaneSecretStatus instantiates a new ManagementV1alpha1DataplaneSecretStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSecretStatusWithDefaults

`func NewManagementV1alpha1DataplaneSecretStatusWithDefaults() *ManagementV1alpha1DataplaneSecretStatus`

NewManagementV1alpha1DataplaneSecretStatusWithDefaults instantiates a new ManagementV1alpha1DataplaneSecretStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ManagementV1alpha1DataplaneSecretStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1DataplaneSecretStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1DataplaneSecretStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *ManagementV1alpha1DataplaneSecretStatus) GetReasons() []ManagementV1alpha1DataplaneSecretStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ManagementV1alpha1DataplaneSecretStatus) GetReasonsOk() (*[]ManagementV1alpha1DataplaneSecretStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ManagementV1alpha1DataplaneSecretStatus) SetReasons(v []ManagementV1alpha1DataplaneSecretStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ManagementV1alpha1DataplaneSecretStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


