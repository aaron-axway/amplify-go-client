# ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Resource level at which access is being granted. | [optional] 
**Kind** | **string** | The kind of scoped resources to provide access. | 
**Name** | **string** | The name of the scoped resource to provide access to. | 
**AllowDelete** | Pointer to **bool** | Set true to allow users to delete the referenced scoped resource. | [optional] 
**AllowWrite** | Pointer to **bool** | Set true to allow users to update the referenced scoped resource. | [optional] 

## Methods

### NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResource

`func NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResource(kind string, name string, ) *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource`

NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResource instantiates a new ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults

`func NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults() *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource`

NewManagementV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults instantiates a new ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) SetName(v string)`

SetName sets Name field to given value.


### GetAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecAccessLevelScopedResource) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


