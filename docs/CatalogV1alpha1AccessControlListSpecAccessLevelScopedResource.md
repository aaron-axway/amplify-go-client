# CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Resource level at which access is being granted. | [optional] 
**Kind** | **string** | The kind of scoped resources to provide access. | 
**Name** | **string** | The name of the scoped resource to provide access to. | 
**AllowDelete** | Pointer to **bool** | Set true to allow users to delete the referenced scoped resource. | [optional] 
**AllowWrite** | Pointer to **bool** | Set true to allow users to update the referenced scoped resource. | [optional] 

## Methods

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResource

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResource(kind string, name string, ) *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResource instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults() *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedResourceWithDefaults instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) SetName(v string)`

SetName sets Name field to given value.


### GetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedResource) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


