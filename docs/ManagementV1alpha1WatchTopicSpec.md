# ManagementV1alpha1WatchTopicSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the WatchTopic | [optional] 
**Filters** | [**[]ManagementV1alpha1WatchTopicSpecFiltersInner**](ManagementV1alpha1WatchTopicSpecFiltersInner.md) |  | 

## Methods

### NewManagementV1alpha1WatchTopicSpec

`func NewManagementV1alpha1WatchTopicSpec(filters []ManagementV1alpha1WatchTopicSpecFiltersInner, ) *ManagementV1alpha1WatchTopicSpec`

NewManagementV1alpha1WatchTopicSpec instantiates a new ManagementV1alpha1WatchTopicSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1WatchTopicSpecWithDefaults

`func NewManagementV1alpha1WatchTopicSpecWithDefaults() *ManagementV1alpha1WatchTopicSpec`

NewManagementV1alpha1WatchTopicSpecWithDefaults instantiates a new ManagementV1alpha1WatchTopicSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManagementV1alpha1WatchTopicSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagementV1alpha1WatchTopicSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagementV1alpha1WatchTopicSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagementV1alpha1WatchTopicSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *ManagementV1alpha1WatchTopicSpec) GetFilters() []ManagementV1alpha1WatchTopicSpecFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ManagementV1alpha1WatchTopicSpec) GetFiltersOk() (*[]ManagementV1alpha1WatchTopicSpecFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ManagementV1alpha1WatchTopicSpec) SetFilters(v []ManagementV1alpha1WatchTopicSpecFiltersInner)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


