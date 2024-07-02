# ManagementV1alpha1CredentialSpecState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Reason** | Pointer to **string** | Additional info on the desired state. | [optional] 
**Rotate** | Pointer to **bool** | Defines if credential needs to be rotated. | [optional] 

## Methods

### NewManagementV1alpha1CredentialSpecState

`func NewManagementV1alpha1CredentialSpecState(name string, ) *ManagementV1alpha1CredentialSpecState`

NewManagementV1alpha1CredentialSpecState instantiates a new ManagementV1alpha1CredentialSpecState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1CredentialSpecStateWithDefaults

`func NewManagementV1alpha1CredentialSpecStateWithDefaults() *ManagementV1alpha1CredentialSpecState`

NewManagementV1alpha1CredentialSpecStateWithDefaults instantiates a new ManagementV1alpha1CredentialSpecState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagementV1alpha1CredentialSpecState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementV1alpha1CredentialSpecState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementV1alpha1CredentialSpecState) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *ManagementV1alpha1CredentialSpecState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ManagementV1alpha1CredentialSpecState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ManagementV1alpha1CredentialSpecState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ManagementV1alpha1CredentialSpecState) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRotate

`func (o *ManagementV1alpha1CredentialSpecState) GetRotate() bool`

GetRotate returns the Rotate field if non-nil, zero value otherwise.

### GetRotateOk

`func (o *ManagementV1alpha1CredentialSpecState) GetRotateOk() (*bool, bool)`

GetRotateOk returns a tuple with the Rotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotate

`func (o *ManagementV1alpha1CredentialSpecState) SetRotate(v bool)`

SetRotate sets Rotate field to given value.

### HasRotate

`func (o *ManagementV1alpha1CredentialSpecState) HasRotate() bool`

HasRotate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


