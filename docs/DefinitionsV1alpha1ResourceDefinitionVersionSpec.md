# DefinitionsV1alpha1ResourceDefinitionVersionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDefinition** | **string** | A reference to a resource definition. | 
**Name** | **string** | The name of the version to be used in the resource definition. | 
**Served** | **bool** | Defines if this version is currently served by the server. | 
**Storage** | **bool** | Defines if this version is used when storing the resource instance details. | 
**Spec** | Pointer to **map[string]interface{}** | Defines the strucure of the spec for this resource version. | [optional] 

## Methods

### NewDefinitionsV1alpha1ResourceDefinitionVersionSpec

`func NewDefinitionsV1alpha1ResourceDefinitionVersionSpec(resourceDefinition string, name string, served bool, storage bool, ) *DefinitionsV1alpha1ResourceDefinitionVersionSpec`

NewDefinitionsV1alpha1ResourceDefinitionVersionSpec instantiates a new DefinitionsV1alpha1ResourceDefinitionVersionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionsV1alpha1ResourceDefinitionVersionSpecWithDefaults

`func NewDefinitionsV1alpha1ResourceDefinitionVersionSpecWithDefaults() *DefinitionsV1alpha1ResourceDefinitionVersionSpec`

NewDefinitionsV1alpha1ResourceDefinitionVersionSpecWithDefaults instantiates a new DefinitionsV1alpha1ResourceDefinitionVersionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDefinition

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetResourceDefinition() string`

GetResourceDefinition returns the ResourceDefinition field if non-nil, zero value otherwise.

### GetResourceDefinitionOk

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetResourceDefinitionOk() (*string, bool)`

GetResourceDefinitionOk returns a tuple with the ResourceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDefinition

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) SetResourceDefinition(v string)`

SetResourceDefinition sets ResourceDefinition field to given value.


### GetName

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) SetName(v string)`

SetName sets Name field to given value.


### GetServed

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetServed() bool`

GetServed returns the Served field if non-nil, zero value otherwise.

### GetServedOk

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetServedOk() (*bool, bool)`

GetServedOk returns a tuple with the Served field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServed

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) SetServed(v bool)`

SetServed sets Served field to given value.


### GetStorage

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetStorage() bool`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetStorageOk() (*bool, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) SetStorage(v bool)`

SetStorage sets Storage field to given value.


### GetSpec

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DefinitionsV1alpha1ResourceDefinitionVersionSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


