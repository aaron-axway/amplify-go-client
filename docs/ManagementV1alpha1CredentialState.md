# ManagementV1alpha1CredentialState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Reason** | Pointer to **string** | Additional info on the state. | [optional] 

## Methods

### NewManagementV1alpha1CredentialState

`func NewManagementV1alpha1CredentialState(name string, ) *ManagementV1alpha1CredentialState`

NewManagementV1alpha1CredentialState instantiates a new ManagementV1alpha1CredentialState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialStateWithDefaults

`func NewManagementV1alpha1CredentialStateWithDefaults() *ManagementV1alpha1CredentialState`

NewManagementV1alpha1CredentialStateWithDefaults instantiates a new ManagementV1alpha1CredentialState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1CredentialState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1CredentialState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1CredentialState) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *ManagementV1alpha1CredentialState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ManagementV1alpha1CredentialState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ManagementV1alpha1CredentialState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ManagementV1alpha1CredentialState) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


