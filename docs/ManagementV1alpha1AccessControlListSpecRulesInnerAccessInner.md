# ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Resource level at which access is being granted. | [optional] 
**AllowCreateScoped** | Pointer to **bool** | Set true to allow users to create scoped resources. | [optional] 
**AllowDelete** | Pointer to **bool** | Set true to allow users to delete the referenced scoped resource. | [optional] 
**AllowWrite** | Pointer to **bool** | Set true to allow users to update the referenced scoped resource. | [optional] 
**Kind** | **string** | The kind of scoped resources to provide access. | 
**AllowCreate** | Pointer to **bool** | Set true to allow users to create scoped resources of the specified kind. | [optional] 
**Name** | **string** | The name of the scoped resource to provide access to. | 

## Methods

### NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInner

`func NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInner(kind string, name string, ) *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner`

NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInner instantiates a new ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInnerWithDefaults

`func NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInnerWithDefaults() *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner`

NewManagementV1alpha1AccessControlListSpecRulesInnerAccessInnerWithDefaults instantiates a new ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAllowCreateScoped

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowCreateScoped() bool`

GetAllowCreateScoped returns the AllowCreateScoped field if non-nil, zero value otherwise.

### GetAllowCreateScopedOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowCreateScopedOk() (*bool, bool)`

GetAllowCreateScopedOk returns a tuple with the AllowCreateScoped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreateScoped

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetAllowCreateScoped(v bool)`

SetAllowCreateScoped sets AllowCreateScoped field to given value.

### HasAllowCreateScoped

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) HasAllowCreateScoped() bool`

HasAllowCreateScoped returns a boolean if a field has been set.

### GetAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### GetKind

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAllowCreate

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowCreate() bool`

GetAllowCreate returns the AllowCreate field if non-nil, zero value otherwise.

### GetAllowCreateOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetAllowCreateOk() (*bool, bool)`

GetAllowCreateOk returns a tuple with the AllowCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreate

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetAllowCreate(v bool)`

SetAllowCreate sets AllowCreate field to given value.

### HasAllowCreate

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) HasAllowCreate() bool`

HasAllowCreate returns a boolean if a field has been set.

### GetName

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1AccessControlListSpecRulesInnerAccessInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


