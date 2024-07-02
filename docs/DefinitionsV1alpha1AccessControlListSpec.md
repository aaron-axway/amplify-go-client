# DefinitionsV1alpha1AccessControlListSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]DefinitionsV1alpha1AccessControlListSpecRulesInner**](DefinitionsV1alpha1AccessControlListSpecRulesInner.md) |  | [optional] 
**Subjects** | Pointer to [**[]CatalogV1alpha1AccessControlListSpecSubjectsInner**](CatalogV1alpha1AccessControlListSpecSubjectsInner.md) |  | [optional] 

## Methods

### NewDefinitionsV1alpha1AccessControlListSpec

`func NewDefinitionsV1alpha1AccessControlListSpec() *DefinitionsV1alpha1AccessControlListSpec`

NewDefinitionsV1alpha1AccessControlListSpec instantiates a new DefinitionsV1alpha1AccessControlListSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1AccessControlListSpecWithDefaults

`func NewDefinitionsV1alpha1AccessControlListSpecWithDefaults() *DefinitionsV1alpha1AccessControlListSpec`

NewDefinitionsV1alpha1AccessControlListSpecWithDefaults instantiates a new DefinitionsV1alpha1AccessControlListSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *DefinitionsV1alpha1AccessControlListSpec) GetRules() []DefinitionsV1alpha1AccessControlListSpecRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DefinitionsV1alpha1AccessControlListSpec) GetRulesOk() (*[]DefinitionsV1alpha1AccessControlListSpecRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DefinitionsV1alpha1AccessControlListSpec) SetRules(v []DefinitionsV1alpha1AccessControlListSpecRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DefinitionsV1alpha1AccessControlListSpec) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSubjects

`func (o *DefinitionsV1alpha1AccessControlListSpec) GetSubjects() []CatalogV1alpha1AccessControlListSpecSubjectsInner`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *DefinitionsV1alpha1AccessControlListSpec) GetSubjectsOk() (*[]CatalogV1alpha1AccessControlListSpecSubjectsInner, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *DefinitionsV1alpha1AccessControlListSpec) SetSubjects(v []CatalogV1alpha1AccessControlListSpecSubjectsInner)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *DefinitionsV1alpha1AccessControlListSpec) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


