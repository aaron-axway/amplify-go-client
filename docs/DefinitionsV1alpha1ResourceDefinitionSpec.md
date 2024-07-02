# DefinitionsV1alpha1ResourceDefinitionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plural** | **string** | Value used in the endpoint path for accessing the resource. | 
**Kind** | **string** | Defines the kind of the resource. The server infers this from the endpoint the client submits the request to. | 
**Scope** | Pointer to [**DefinitionsV1alpha1ResourceDefinitionSpecScope**](DefinitionsV1alpha1ResourceDefinitionSpecScope.md) |  | [optional] 
**QueryParams** | Pointer to **[]string** |  | [optional] 
**SubResources** | Pointer to [**DefinitionsV1alpha1ResourceDefinitionSpecSubResources**](DefinitionsV1alpha1ResourceDefinitionSpecSubResources.md) |  | [optional] 
**References** | Pointer to [**DefinitionsV1alpha1ResourceDefinitionSpecReferences**](DefinitionsV1alpha1ResourceDefinitionSpecReferences.md) |  | [optional] 

## Methods

### NewDefinitionsV1alpha1ResourceDefinitionSpec

`func NewDefinitionsV1alpha1ResourceDefinitionSpec(plural string, kind string, ) *DefinitionsV1alpha1ResourceDefinitionSpec`

NewDefinitionsV1alpha1ResourceDefinitionSpec instantiates a new DefinitionsV1alpha1ResourceDefinitionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1ResourceDefinitionSpecWithDefaults

`func NewDefinitionsV1alpha1ResourceDefinitionSpecWithDefaults() *DefinitionsV1alpha1ResourceDefinitionSpec`

NewDefinitionsV1alpha1ResourceDefinitionSpecWithDefaults instantiates a new DefinitionsV1alpha1ResourceDefinitionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlural

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetPluralOk() (*string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlural

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetPlural(v string)`

SetPlural sets Plural field to given value.


### GetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetScope

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetScope() DefinitionsV1alpha1ResourceDefinitionSpecScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetScopeOk() (*DefinitionsV1alpha1ResourceDefinitionSpecScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetScope(v DefinitionsV1alpha1ResourceDefinitionSpecScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetQueryParams

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetQueryParams() []string`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetQueryParamsOk() (*[]string, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetQueryParams(v []string)`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetSubResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetSubResources() DefinitionsV1alpha1ResourceDefinitionSpecSubResources`

GetSubResources returns the SubResources field if non-nil, zero value otherwise.

### GetSubResourcesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetSubResourcesOk() (*DefinitionsV1alpha1ResourceDefinitionSpecSubResources, bool)`

GetSubResourcesOk returns a tuple with the SubResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetSubResources(v DefinitionsV1alpha1ResourceDefinitionSpecSubResources)`

SetSubResources sets SubResources field to given value.

### HasSubResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) HasSubResources() bool`

HasSubResources returns a boolean if a field has been set.

### GetReferences

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetReferences() DefinitionsV1alpha1ResourceDefinitionSpecReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) GetReferencesOk() (*DefinitionsV1alpha1ResourceDefinitionSpecReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) SetReferences(v DefinitionsV1alpha1ResourceDefinitionSpecReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *DefinitionsV1alpha1ResourceDefinitionSpec) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


