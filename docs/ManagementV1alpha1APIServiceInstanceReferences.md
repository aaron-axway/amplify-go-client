# ManagementV1alpha1APIServiceInstanceReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiService** | Pointer to **string** | Reference to Amplify Central APIService | [optional] 
**Stage** | Pointer to **string** | Computes the APIServiceInstance stage based on the Environment default stage or any override provided in the APIServiceInstance spec.stage property. | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceReferences

`func NewManagementV1alpha1APIServiceInstanceReferences() *ManagementV1alpha1APIServiceInstanceReferences`

NewManagementV1alpha1APIServiceInstanceReferences instantiates a new ManagementV1alpha1APIServiceInstanceReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceReferencesWithDefaults

`func NewManagementV1alpha1APIServiceInstanceReferencesWithDefaults() *ManagementV1alpha1APIServiceInstanceReferences`

NewManagementV1alpha1APIServiceInstanceReferencesWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiService

`func (o *ManagementV1alpha1APIServiceInstanceReferences) GetApiService() string`

GetApiService returns the ApiService field if non-nil, zero value otherwise.

### GetApiServiceOk

`func (o *ManagementV1alpha1APIServiceInstanceReferences) GetApiServiceOk() (*string, bool)`

GetApiServiceOk returns a tuple with the ApiService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiService

`func (o *ManagementV1alpha1APIServiceInstanceReferences) SetApiService(v string)`

SetApiService sets ApiService field to given value.

### HasApiService

`func (o *ManagementV1alpha1APIServiceInstanceReferences) HasApiService() bool`

HasApiService returns a boolean if a field has been set.

### GetStage

`func (o *ManagementV1alpha1APIServiceInstanceReferences) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ManagementV1alpha1APIServiceInstanceReferences) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ManagementV1alpha1APIServiceInstanceReferences) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ManagementV1alpha1APIServiceInstanceReferences) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


