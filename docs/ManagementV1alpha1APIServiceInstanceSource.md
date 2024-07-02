# ManagementV1alpha1APIServiceInstanceSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataplaneType** | Pointer to [**ManagementV1alpha1APIServiceSourceDataplaneType**](ManagementV1alpha1APIServiceSourceDataplaneType.md) |  | [optional] 
**References** | Pointer to [**ManagementV1alpha1APIServiceInstanceSourceReferences**](ManagementV1alpha1APIServiceInstanceSourceReferences.md) |  | [optional] 
**Compliance** | Pointer to [**ManagementV1alpha1APIServiceInstanceSourceCompliance**](ManagementV1alpha1APIServiceInstanceSourceCompliance.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceSource

`func NewManagementV1alpha1APIServiceInstanceSource() *ManagementV1alpha1APIServiceInstanceSource`

NewManagementV1alpha1APIServiceInstanceSource instantiates a new ManagementV1alpha1APIServiceInstanceSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceSourceWithDefaults

`func NewManagementV1alpha1APIServiceInstanceSourceWithDefaults() *ManagementV1alpha1APIServiceInstanceSource`

NewManagementV1alpha1APIServiceInstanceSourceWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataplaneType

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetDataplaneType() ManagementV1alpha1APIServiceSourceDataplaneType`

GetDataplaneType returns the DataplaneType field if non-nil, zero value otherwise.

### GetDataplaneTypeOk

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetDataplaneTypeOk() (*ManagementV1alpha1APIServiceSourceDataplaneType, bool)`

GetDataplaneTypeOk returns a tuple with the DataplaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplaneType

`func (o *ManagementV1alpha1APIServiceInstanceSource) SetDataplaneType(v ManagementV1alpha1APIServiceSourceDataplaneType)`

SetDataplaneType sets DataplaneType field to given value.

### HasDataplaneType

`func (o *ManagementV1alpha1APIServiceInstanceSource) HasDataplaneType() bool`

HasDataplaneType returns a boolean if a field has been set.

### GetReferences

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetReferences() ManagementV1alpha1APIServiceInstanceSourceReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetReferencesOk() (*ManagementV1alpha1APIServiceInstanceSourceReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ManagementV1alpha1APIServiceInstanceSource) SetReferences(v ManagementV1alpha1APIServiceInstanceSourceReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ManagementV1alpha1APIServiceInstanceSource) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetCompliance

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetCompliance() ManagementV1alpha1APIServiceInstanceSourceCompliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *ManagementV1alpha1APIServiceInstanceSource) GetComplianceOk() (*ManagementV1alpha1APIServiceInstanceSourceCompliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *ManagementV1alpha1APIServiceInstanceSource) SetCompliance(v ManagementV1alpha1APIServiceInstanceSourceCompliance)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *ManagementV1alpha1APIServiceInstanceSource) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


