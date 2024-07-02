# DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Defines the kind of the resource. | [optional] 
**Kind** | Pointer to **string** | Defines the kind of the resource. | [optional] 
**ScopeKind** | Pointer to **string** | Defines the scope kind of the resource. Omit for unscoped resources. | [optional] 
**From** | Pointer to [**DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerFrom**](DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerFrom.md) |  | [optional] 
**Types** | Pointer to **[]string** | The type of the reference. | [optional] 

## Methods

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner() *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerWithDefaults

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerWithDefaults() *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerWithDefaults instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetScopeKind() string`

GetScopeKind returns the ScopeKind field if non-nil, zero value otherwise.

### GetScopeKindOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetScopeKindOk() (*string, bool)`

GetScopeKindOk returns a tuple with the ScopeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) SetScopeKind(v string)`

SetScopeKind sets ScopeKind field to given value.

### HasScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) HasScopeKind() bool`

HasScopeKind returns a boolean if a field has been set.

### GetFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetFrom() DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetFromOk() (*DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) SetFrom(v DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInnerFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


