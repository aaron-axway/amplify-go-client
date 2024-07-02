# ManagementV1alpha1WatchTopicSpecFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | Value for the group of the resource. Use \&quot;*\&quot; for any. | 
**Scope** | Pointer to [**ManagementV1alpha1WatchTopicSpecFiltersInnerScope**](ManagementV1alpha1WatchTopicSpecFiltersInnerScope.md) |  | [optional] 
**Kind** | **string** | Value for the Kind of the resource. Use \&quot;*\&quot; for any. | 
**Name** | **string** | Name of the resource. Use \&quot;*\&quot; for any. | 
**Type** | **[]string** |  | 

## Methods

### NewManagementV1alpha1WatchTopicSpecFiltersInner

`func NewManagementV1alpha1WatchTopicSpecFiltersInner(group string, kind string, name string, type_ []string, ) *ManagementV1alpha1WatchTopicSpecFiltersInner`

NewManagementV1alpha1WatchTopicSpecFiltersInner instantiates a new ManagementV1alpha1WatchTopicSpecFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1WatchTopicSpecFiltersInnerWithDefaults

`func NewManagementV1alpha1WatchTopicSpecFiltersInnerWithDefaults() *ManagementV1alpha1WatchTopicSpecFiltersInner`

NewManagementV1alpha1WatchTopicSpecFiltersInnerWithDefaults instantiates a new ManagementV1alpha1WatchTopicSpecFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetScope

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetScope() ManagementV1alpha1WatchTopicSpecFiltersInnerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetScopeOk() (*ManagementV1alpha1WatchTopicSpecFiltersInnerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) SetScope(v ManagementV1alpha1WatchTopicSpecFiltersInnerScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1WatchTopicSpecFiltersInner) SetType(v []string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


