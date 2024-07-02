# DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the column for the resource. | 
**Type** | **string** | The type of the column. | 
**JsonPath** | **string** | The JSONPath used to fetch data for the column starting from the root. | 
**Description** | Pointer to **string** | The description of the column data. | [optional] 

## Methods

### NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner

`func NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner(name string, type_ string, jsonPath string, ) *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner`

NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInnerWithDefaults

`func NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInnerWithDefaults() *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner`

NewDefinitionsV1alpha1CommandLineInterfaceSpecColumnsInnerWithDefaults instantiates a new DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetType(v string)`

SetType sets Type field to given value.


### GetJsonPath

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.


### GetDescription

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefinitionsV1alpha1CommandLineInterfaceSpecColumnsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


