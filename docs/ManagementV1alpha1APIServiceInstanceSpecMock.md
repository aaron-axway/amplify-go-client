# ManagementV1alpha1APIServiceInstanceSpecMock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Assigned to the mock server&#39;s URL base path. Must be unique for the organization. | 
**UseLatestRevision** | Pointer to [**ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision**](ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision.md) |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceInstanceSpecMock

`func NewManagementV1alpha1APIServiceInstanceSpecMock(name string, ) *ManagementV1alpha1APIServiceInstanceSpecMock`

NewManagementV1alpha1APIServiceInstanceSpecMock instantiates a new ManagementV1alpha1APIServiceInstanceSpecMock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceInstanceSpecMockWithDefaults

`func NewManagementV1alpha1APIServiceInstanceSpecMockWithDefaults() *ManagementV1alpha1APIServiceInstanceSpecMock`

NewManagementV1alpha1APIServiceInstanceSpecMockWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSpecMock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) SetName(v string)`

SetName sets Name field to given value.


### GetUseLatestRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetUseLatestRevision() ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision`

GetUseLatestRevision returns the UseLatestRevision field if non-nil, zero value otherwise.

### GetUseLatestRevisionOk

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) GetUseLatestRevisionOk() (*ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision, bool)`

GetUseLatestRevisionOk returns a tuple with the UseLatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLatestRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) SetUseLatestRevision(v ManagementV1alpha1APIServiceInstanceSpecMockUseLatestRevision)`

SetUseLatestRevision sets UseLatestRevision field to given value.

### HasUseLatestRevision

`func (o *ManagementV1alpha1APIServiceInstanceSpecMock) HasUseLatestRevision() bool`

HasUseLatestRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


