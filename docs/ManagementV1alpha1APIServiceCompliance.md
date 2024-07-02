# ManagementV1alpha1APIServiceCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Design** | Pointer to [**ManagementV1alpha1APIServiceComplianceLintingStatus**](ManagementV1alpha1APIServiceComplianceLintingStatus.md) |  | [optional] 
**Security** | Pointer to [**ManagementV1alpha1APIServiceComplianceLintingStatus**](ManagementV1alpha1APIServiceComplianceLintingStatus.md) |  | [optional] 
**Runtime** | Pointer to [**ManagementV1alpha1APIServiceComplianceRuntimeStatus**](ManagementV1alpha1APIServiceComplianceRuntimeStatus.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceCompliance

`func NewManagementV1alpha1APIServiceCompliance() *ManagementV1alpha1APIServiceCompliance`

NewManagementV1alpha1APIServiceCompliance instantiates a new ManagementV1alpha1APIServiceCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceComplianceWithDefaults

`func NewManagementV1alpha1APIServiceComplianceWithDefaults() *ManagementV1alpha1APIServiceCompliance`

NewManagementV1alpha1APIServiceComplianceWithDefaults instantiates a new ManagementV1alpha1APIServiceCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesign

`func (o *ManagementV1alpha1APIServiceCompliance) GetDesign() ManagementV1alpha1APIServiceComplianceLintingStatus`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *ManagementV1alpha1APIServiceCompliance) GetDesignOk() (*ManagementV1alpha1APIServiceComplianceLintingStatus, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *ManagementV1alpha1APIServiceCompliance) SetDesign(v ManagementV1alpha1APIServiceComplianceLintingStatus)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *ManagementV1alpha1APIServiceCompliance) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetSecurity

`func (o *ManagementV1alpha1APIServiceCompliance) GetSecurity() ManagementV1alpha1APIServiceComplianceLintingStatus`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ManagementV1alpha1APIServiceCompliance) GetSecurityOk() (*ManagementV1alpha1APIServiceComplianceLintingStatus, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ManagementV1alpha1APIServiceCompliance) SetSecurity(v ManagementV1alpha1APIServiceComplianceLintingStatus)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ManagementV1alpha1APIServiceCompliance) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetRuntime

`func (o *ManagementV1alpha1APIServiceCompliance) GetRuntime() ManagementV1alpha1APIServiceComplianceRuntimeStatus`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ManagementV1alpha1APIServiceCompliance) GetRuntimeOk() (*ManagementV1alpha1APIServiceComplianceRuntimeStatus, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ManagementV1alpha1APIServiceCompliance) SetRuntime(v ManagementV1alpha1APIServiceComplianceRuntimeStatus)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ManagementV1alpha1APIServiceCompliance) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


