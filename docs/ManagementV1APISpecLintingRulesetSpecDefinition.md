# ManagementV1APISpecLintingRulesetSpecDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LintingType** | **string** | The type of the API Specification Linting Ruleset. | 
**Value** | **string** | Base64 encoded value of the API Specification Linting Ruleset. | 
**ContentType** | **string** | The MIME type of property \&quot;value\&quot; such as \&quot;application/json\&quot; or \&quot;application/yaml\&quot;. | 
**FileName** | Pointer to **string** | File name and extension of the APISpecLintingRuleset such as \&quot;my-file.json\&quot;. | [optional] 
**IsDefault** | Pointer to **bool** | Defines if the APISpecLintingRuleset is default. | [optional] 
**ApiTypes** | **[]string** | Indicates which API specification types can be linted by this ruleset. | 

## Methods

### NewManagementV1APISpecLintingRulesetSpecDefinition

`func NewManagementV1APISpecLintingRulesetSpecDefinition(lintingType string, value string, contentType string, apiTypes []string, ) *ManagementV1APISpecLintingRulesetSpecDefinition`

NewManagementV1APISpecLintingRulesetSpecDefinition instantiates a new ManagementV1APISpecLintingRulesetSpecDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1APISpecLintingRulesetSpecDefinitionWithDefaults

`func NewManagementV1APISpecLintingRulesetSpecDefinitionWithDefaults() *ManagementV1APISpecLintingRulesetSpecDefinition`

NewManagementV1APISpecLintingRulesetSpecDefinitionWithDefaults instantiates a new ManagementV1APISpecLintingRulesetSpecDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLintingType

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetLintingType() string`

GetLintingType returns the LintingType field if non-nil, zero value otherwise.

### GetLintingTypeOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetLintingTypeOk() (*string, bool)`

GetLintingTypeOk returns a tuple with the LintingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLintingType

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetLintingType(v string)`

SetLintingType sets LintingType field to given value.


### GetValue

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetValue(v string)`

SetValue sets Value field to given value.


### GetContentType

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetFileName

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIsDefault

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetApiTypes

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetApiTypes() []string`

GetApiTypes returns the ApiTypes field if non-nil, zero value otherwise.

### GetApiTypesOk

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) GetApiTypesOk() (*[]string, bool)`

GetApiTypesOk returns a tuple with the ApiTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTypes

`func (o *ManagementV1APISpecLintingRulesetSpecDefinition) SetApiTypes(v []string)`

SetApiTypes sets ApiTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


