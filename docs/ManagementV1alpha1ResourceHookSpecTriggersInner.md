# ManagementV1alpha1ResourceHookSpecTriggersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | Value for the group of the resource. Use \&quot;*\&quot; for any. | 
**Scope** | Pointer to [**ManagementV1alpha1ResourceHookSpecTriggersInnerScope**](ManagementV1alpha1ResourceHookSpecTriggersInnerScope.md) |  | [optional] 
**Kind** | **string** | Value for the Kind of the resource. Use \&quot;*\&quot; for any. | 
**Name** | **string** | Name of the resource. Use \&quot;*\&quot; for any. | 
**Type** | **[]string** |  | 

## Methods

### NewManagementV1alpha1ResourceHookSpecTriggersInner

`func NewManagementV1alpha1ResourceHookSpecTriggersInner(group string, kind string, name string, type_ []string, ) *ManagementV1alpha1ResourceHookSpecTriggersInner`

NewManagementV1alpha1ResourceHookSpecTriggersInner instantiates a new ManagementV1alpha1ResourceHookSpecTriggersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ResourceHookSpecTriggersInnerWithDefaults

`func NewManagementV1alpha1ResourceHookSpecTriggersInnerWithDefaults() *ManagementV1alpha1ResourceHookSpecTriggersInner`

NewManagementV1alpha1ResourceHookSpecTriggersInnerWithDefaults instantiates a new ManagementV1alpha1ResourceHookSpecTriggersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetScope

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetScope() ManagementV1alpha1ResourceHookSpecTriggersInnerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetScopeOk() (*ManagementV1alpha1ResourceHookSpecTriggersInnerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetScope(v ManagementV1alpha1ResourceHookSpecTriggersInnerScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1ResourceHookSpecTriggersInner) SetType(v []string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


