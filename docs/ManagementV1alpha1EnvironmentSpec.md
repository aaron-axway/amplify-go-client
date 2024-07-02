# ManagementV1alpha1EnvironmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]string** |  | [optional] 
**Production** | Pointer to **bool** | Production environment will be used to perform production processing or connect to a non-Axway gateway. This usage will count against your entitled quota. | [optional] 
**AxwayManaged** | Pointer to **bool** | Axway Managed environment is hosted in the Axway Managed Cloud. | [optional] 
**Icon** | Pointer to [**ManagementV1alpha1EnvironmentSpecIcon**](ManagementV1alpha1EnvironmentSpecIcon.md) |  | [optional] 
**Compliance** | Pointer to [**ManagementV1alpha1EnvironmentSpecCompliance**](ManagementV1alpha1EnvironmentSpecCompliance.md) |  | [optional] 

## Methods

### NewManagementV1alpha1EnvironmentSpec

`func NewManagementV1alpha1EnvironmentSpec() *ManagementV1alpha1EnvironmentSpec`

NewManagementV1alpha1EnvironmentSpec instantiates a new ManagementV1alpha1EnvironmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1EnvironmentSpecWithDefaults

`func NewManagementV1alpha1EnvironmentSpecWithDefaults() *ManagementV1alpha1EnvironmentSpec`

NewManagementV1alpha1EnvironmentSpecWithDefaults instantiates a new ManagementV1alpha1EnvironmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManagementV1alpha1EnvironmentSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagementV1alpha1EnvironmentSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagementV1alpha1EnvironmentSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariables

`func (o *ManagementV1alpha1EnvironmentSpec) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ManagementV1alpha1EnvironmentSpec) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ManagementV1alpha1EnvironmentSpec) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetProduction

`func (o *ManagementV1alpha1EnvironmentSpec) GetProduction() bool`

GetProduction returns the Production field if non-nil, zero value otherwise.

### GetProductionOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetProductionOk() (*bool, bool)`

GetProductionOk returns a tuple with the Production field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduction

`func (o *ManagementV1alpha1EnvironmentSpec) SetProduction(v bool)`

SetProduction sets Production field to given value.

### HasProduction

`func (o *ManagementV1alpha1EnvironmentSpec) HasProduction() bool`

HasProduction returns a boolean if a field has been set.

### GetAxwayManaged

`func (o *ManagementV1alpha1EnvironmentSpec) GetAxwayManaged() bool`

GetAxwayManaged returns the AxwayManaged field if non-nil, zero value otherwise.

### GetAxwayManagedOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetAxwayManagedOk() (*bool, bool)`

GetAxwayManagedOk returns a tuple with the AxwayManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxwayManaged

`func (o *ManagementV1alpha1EnvironmentSpec) SetAxwayManaged(v bool)`

SetAxwayManaged sets AxwayManaged field to given value.

### HasAxwayManaged

`func (o *ManagementV1alpha1EnvironmentSpec) HasAxwayManaged() bool`

HasAxwayManaged returns a boolean if a field has been set.

### GetIcon

`func (o *ManagementV1alpha1EnvironmentSpec) GetIcon() ManagementV1alpha1EnvironmentSpecIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetIconOk() (*ManagementV1alpha1EnvironmentSpecIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ManagementV1alpha1EnvironmentSpec) SetIcon(v ManagementV1alpha1EnvironmentSpecIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ManagementV1alpha1EnvironmentSpec) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetCompliance

`func (o *ManagementV1alpha1EnvironmentSpec) GetCompliance() ManagementV1alpha1EnvironmentSpecCompliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *ManagementV1alpha1EnvironmentSpec) GetComplianceOk() (*ManagementV1alpha1EnvironmentSpecCompliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *ManagementV1alpha1EnvironmentSpec) SetCompliance(v ManagementV1alpha1EnvironmentSpecCompliance)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *ManagementV1alpha1EnvironmentSpec) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


