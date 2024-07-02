# DefinitionsV1alpha1CommandLineInterfaceSpecNames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plural** | **string** | Defines the name used to access resources in this group. Also provided as default in the autocomplete for listing commands. | 
**Singular** | **string** | Defines the name used to access a resource in this group. Also provided as default in the autocomplete for single resource access commands. | 
**ShortNames** | **[]string** | Defines the short names that the cli can use to fetch a resource in the group. | 
**ShortNamesAlias** | Pointer to **[]string** | Defines the short names alias that the cli can use to fetch a resource in the group. | [optional] 

## Methods

### NewDefinitionsV1alpha1CommandLineInterfaceSpecNames

`func NewDefinitionsV1alpha1CommandLineInterfaceSpecNames(plural string, singular string, shortNames []string, ) *DefinitionsV1alpha1CommandLineInterfaceSpecNames`

NewDefinitionsV1alpha1CommandLineInterfaceSpecNames instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecNames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1CommandLineInterfaceSpecNamesWithDefaults

`func NewDefinitionsV1alpha1CommandLineInterfaceSpecNamesWithDefaults() *DefinitionsV1alpha1CommandLineInterfaceSpecNames`

NewDefinitionsV1alpha1CommandLineInterfaceSpecNamesWithDefaults instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecNames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlural

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetPlural() string`

GetPlural returns the Plural field if non-nil, zero value otherwise.

### GetPluralOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetPluralOk() (*string, bool)`

GetPluralOk returns a tuple with the Plural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlural

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) SetPlural(v string)`

SetPlural sets Plural field to given value.


### GetSingular

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetSingular() string`

GetSingular returns the Singular field if non-nil, zero value otherwise.

### GetSingularOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetSingularOk() (*string, bool)`

GetSingularOk returns a tuple with the Singular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingular

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) SetSingular(v string)`

SetSingular sets Singular field to given value.


### GetShortNames

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetShortNames() []string`

GetShortNames returns the ShortNames field if non-nil, zero value otherwise.

### GetShortNamesOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetShortNamesOk() (*[]string, bool)`

GetShortNamesOk returns a tuple with the ShortNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNames

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) SetShortNames(v []string)`

SetShortNames sets ShortNames field to given value.


### GetShortNamesAlias

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetShortNamesAlias() []string`

GetShortNamesAlias returns the ShortNamesAlias field if non-nil, zero value otherwise.

### GetShortNamesAliasOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) GetShortNamesAliasOk() (*[]string, bool)`

GetShortNamesAliasOk returns a tuple with the ShortNamesAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNamesAlias

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) SetShortNamesAlias(v []string)`

SetShortNamesAlias sets ShortNamesAlias field to given value.

### HasShortNamesAlias

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecNames) HasShortNamesAlias() bool`

HasShortNamesAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


