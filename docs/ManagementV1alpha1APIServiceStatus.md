# ManagementV1alpha1APIServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The current status level, indicating progress towards consistency. | 
**Reasons** | Pointer to [**[]ManagementV1alpha1APIServiceStatusReasonsInner**](ManagementV1alpha1APIServiceStatusReasonsInner.md) | Reasons for the generated status. | [optional] 

## Methods

### NewManagementV1alpha1APIServiceStatus

`func NewManagementV1alpha1APIServiceStatus(level string, ) *ManagementV1alpha1APIServiceStatus`

NewManagementV1alpha1APIServiceStatus instantiates a new ManagementV1alpha1APIServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceStatusWithDefaults

`func NewManagementV1alpha1APIServiceStatusWithDefaults() *ManagementV1alpha1APIServiceStatus`

NewManagementV1alpha1APIServiceStatusWithDefaults instantiates a new ManagementV1alpha1APIServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ManagementV1alpha1APIServiceStatus) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1APIServiceStatus) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1APIServiceStatus) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetReasons

`func (o *ManagementV1alpha1APIServiceStatus) GetReasons() []ManagementV1alpha1APIServiceStatusReasonsInner`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ManagementV1alpha1APIServiceStatus) GetReasonsOk() (*[]ManagementV1alpha1APIServiceStatusReasonsInner, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ManagementV1alpha1APIServiceStatus) SetReasons(v []ManagementV1alpha1APIServiceStatusReasonsInner)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ManagementV1alpha1APIServiceStatus) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


