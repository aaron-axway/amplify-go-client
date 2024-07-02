# DefinitionsV1alpha1CommandLineInterfaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDefinition** | **string** | A reference to a resource definition. | 
**Names** | Pointer to [**DefinitionsV1alpha1CommandLineInterfaceSpecNames**](DefinitionsV1alpha1CommandLineInterfaceSpecNames.md) |  | [optional] 
**Columns** | [**[]DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner**](DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner.md) | Defines an ordered list of the columns and data to be rendered using a table output. | 

## Methods

### NewDefinitionsV1alpha1CommandLineInterfaceSpec

`func NewDefinitionsV1alpha1CommandLineInterfaceSpec(resourceDefinition string, columns []DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner, ) *DefinitionsV1alpha1CommandLineInterfaceSpec`

NewDefinitionsV1alpha1CommandLineInterfaceSpec instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1CommandLineInterfaceSpecWithDefaults

`func NewDefinitionsV1alpha1CommandLineInterfaceSpecWithDefaults() *DefinitionsV1alpha1CommandLineInterfaceSpec`

NewDefinitionsV1alpha1CommandLineInterfaceSpecWithDefaults instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDefinition

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetResourceDefinition() string`

GetResourceDefinition returns the ResourceDefinition field if non-nil, zero value otherwise.

### GetResourceDefinitionOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetResourceDefinitionOk() (*string, bool)`

GetResourceDefinitionOk returns a tuple with the ResourceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDefinition

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) SetResourceDefinition(v string)`

SetResourceDefinition sets ResourceDefinition field to given value.


### GetNames

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetNames() DefinitionsV1alpha1CommandLineInterfaceSpecNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetNamesOk() (*DefinitionsV1alpha1CommandLineInterfaceSpecNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) SetNames(v DefinitionsV1alpha1CommandLineInterfaceSpecNames)`

SetNames sets Names field to given value.

### HasNames

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetColumns

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetColumns() []DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) GetColumnsOk() (*[]DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpec) SetColumns(v []DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


