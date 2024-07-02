# CatalogV1alpha1CredentialSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialRequestDefinition** | **string** | Reference to Credential Request Definition resource | 
**Data** | **map[string]interface{}** | data matching the credential request definition schema. | 
**State** | Pointer to [**CatalogV1alpha1CredentialSpecState**](CatalogV1alpha1CredentialSpecState.md) |  | [optional] 

## Methods

### NewCatalogV1alpha1CredentialSpec

`func NewCatalogV1alpha1CredentialSpec(credentialRequestDefinition string, data map[string]interface{}, ) *CatalogV1alpha1CredentialSpec`

NewCatalogV1alpha1CredentialSpec instantiates a new CatalogV1alpha1CredentialSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1CredentialSpecWithDefaults

`func NewCatalogV1alpha1CredentialSpecWithDefaults() *CatalogV1alpha1CredentialSpec`

NewCatalogV1alpha1CredentialSpecWithDefaults instantiates a new CatalogV1alpha1CredentialSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialRequestDefinition

`func (o *CatalogV1alpha1CredentialSpec) GetCredentialRequestDefinition() string`

GetCredentialRequestDefinition returns the CredentialRequestDefinition field if non-nil, zero value otherwise.

### GetCredentialRequestDefinitionOk

`func (o *CatalogV1alpha1CredentialSpec) GetCredentialRequestDefinitionOk() (*string, bool)`

GetCredentialRequestDefinitionOk returns a tuple with the CredentialRequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRequestDefinition

`func (o *CatalogV1alpha1CredentialSpec) SetCredentialRequestDefinition(v string)`

SetCredentialRequestDefinition sets CredentialRequestDefinition field to given value.


### GetData

`func (o *CatalogV1alpha1CredentialSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogV1alpha1CredentialSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogV1alpha1CredentialSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetState

`func (o *CatalogV1alpha1CredentialSpec) GetState() CatalogV1alpha1CredentialSpecState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CatalogV1alpha1CredentialSpec) GetStateOk() (*CatalogV1alpha1CredentialSpecState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CatalogV1alpha1CredentialSpec) SetState(v CatalogV1alpha1CredentialSpecState)`

SetState sets State field to given value.

### HasState

`func (o *CatalogV1alpha1CredentialSpec) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

