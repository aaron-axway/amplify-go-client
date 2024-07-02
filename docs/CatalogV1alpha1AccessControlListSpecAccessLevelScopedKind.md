# CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | Resource level at which access is being granted. | [optional] 
**Kind** | **string** | The kind of scoped resources to provide access to or \&quot;*\&quot; for all kinds. | 
**AllowCreate** | Pointer to **bool** | Set true to allow users to create scoped resources of the specified kind. | [optional] 
**AllowDelete** | Pointer to **bool** | Set true to allow users to delete scoped resources of the specified kind. | [optional] 
**AllowWrite** | Pointer to **bool** | Set true to allow users to update scoped resources of the specified kind. | [optional] 

## Methods

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKind

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKind(kind string, ) *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKind instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKindWithDefaults

`func NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKindWithDefaults() *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind`

NewCatalogV1alpha1AccessControlListSpecAccessLevelScopedKindWithDefaults instantiates a new CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetKind

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAllowCreate

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowCreate() bool`

GetAllowCreate returns the AllowCreate field if non-nil, zero value otherwise.

### GetAllowCreateOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowCreateOk() (*bool, bool)`

GetAllowCreateOk returns a tuple with the AllowCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreate

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) SetAllowCreate(v bool)`

SetAllowCreate sets AllowCreate field to given value.

### HasAllowCreate

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) HasAllowCreate() bool`

HasAllowCreate returns a boolean if a field has been set.

### GetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowDelete() bool`

GetAllowDelete returns the AllowDelete field if non-nil, zero value otherwise.

### GetAllowDeleteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowDeleteOk() (*bool, bool)`

GetAllowDeleteOk returns a tuple with the AllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) SetAllowDelete(v bool)`

SetAllowDelete sets AllowDelete field to given value.

### HasAllowDelete

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) HasAllowDelete() bool`

HasAllowDelete returns a boolean if a field has been set.

### GetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *CatalogV1alpha1AccessControlListSpecAccessLevelScopedKind) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


