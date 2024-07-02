# CatalogV1alpha1AccessControlListSpecAccessLevelScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Resource level at which access is being granted. | [optional] 
**AllowCreateScoped** | Pointer to **bool** | Set true to allow users to create scoped resources. | [optional] 
**AllowDelete** | Pointer to **bool** | Set true to allow users to delete the unscoped resource. | [optional] 
**AllowWrite** | Pointer to **bool** | Set true to allow users to update the unscoped resource. | [optional] 

## Methods

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScope

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScope() *CatalogV1alpha1AccessControlListSpecAccessLevelScope`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScope instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScopeWithDefaults

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopeWithDefaults() *CatalogV1alpha1AccessControlListSpecAccessLevelScope`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScopeWithDefaults instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetAllowCreateScoped

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowCreateScoped() bool`

GetAllowCreateScoped returns the AllowCreateScoped field if non-nil, zero value otherwise.

### GetAllowCreateScopedOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowCreateScopedOk() (*bool, bool)`

GetAllowCreateScopedOk returns a tuple with the AllowCreateScoped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreateScoped

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowCreateScoped(v bool)`

SetAllowCreateScoped sets AllowCreateScoped field to given value.

### HasAllowCreateScoped

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowCreateScoped() bool`

HasAllowCreateScoped returns a boolean if a field has been set.

### GetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScope) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


