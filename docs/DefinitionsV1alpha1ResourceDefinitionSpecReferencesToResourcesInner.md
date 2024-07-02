# DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Defines the kind of the resource. | [optional] 
**Kind** | Pointer to **string** | Defines the kind of the resource. | [optional] 
**ScopeKind** | Pointer to **string** | Defines the scope kind of the resource. Omit for unscoped resources. | [optional] 
**Types** | Pointer to **[]string** | The type of the reference. | [optional] 
**From** | Pointer to [**DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerFrom**](DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerFrom.md) |  | [optional] 

## Methods

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner() *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerWithDefaults

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerWithDefaults() *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerWithDefaults instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetScopeKind() string`

GetScopeKind returns the ScopeKind field if non-nil, zero value otherwise.

### GetScopeKindOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetScopeKindOk() (*string, bool)`

GetScopeKindOk returns a tuple with the ScopeKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) SetScopeKind(v string)`

SetScopeKind sets ScopeKind field to given value.

### HasScopeKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) HasScopeKind() bool`

HasScopeKind returns a boolean if a field has been set.

### GetTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetFrom() DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) GetFromOk() (*DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) SetFrom(v DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInnerFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


