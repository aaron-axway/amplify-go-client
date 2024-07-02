# DefinitionsV1alpha1ResourceDefinitionSpecReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToResources** | Pointer to [**[]DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner**](DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner.md) | A list of resources that the current resources has references to. | [optional] 
**FromResources** | Pointer to [**[]DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner**](DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner.md) | A list of resources that the current resources is beging referenced from. | [optional] 

## Methods

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferences

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferences() *DefinitionsV1alpha1ResourceDefinitionSpecReferences`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferences instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesWithDefaults

`func NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesWithDefaults() *DefinitionsV1alpha1ResourceDefinitionSpecReferences`

NewDefinitionsV1alpha1ResourceDefinitionSpecReferencesWithDefaults instantiates a new DefinitionsV1alpha1ResourceDefinitionSpecReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) GetToResources() []DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner`

GetToResources returns the ToResources field if non-nil, zero value otherwise.

### GetToResourcesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) GetToResourcesOk() (*[]DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner, bool)`

GetToResourcesOk returns a tuple with the ToResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) SetToResources(v []DefinitionsV1alpha1ResourceDefinitionSpecReferencesToResourcesInner)`

SetToResources sets ToResources field to given value.

### HasToResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) HasToResources() bool`

HasToResources returns a boolean if a field has been set.

### GetFromResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) GetFromResources() []DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner`

GetFromResources returns the FromResources field if non-nil, zero value otherwise.

### GetFromResourcesOk

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) GetFromResourcesOk() (*[]DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner, bool)`

GetFromResourcesOk returns a tuple with the FromResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) SetFromResources(v []DefinitionsV1alpha1ResourceDefinitionSpecReferencesFromResourcesInner)`

SetFromResources sets FromResources field to given value.

### HasFromResources

`func (o *DefinitionsV1alpha1ResourceDefinitionSpecReferences) HasFromResources() bool`

HasFromResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


