# ManagementV1alpha1ManagedApplicationSpecSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | **string** | public key to be used to encrypt the credentials linked to this Managed Application. | 
**EncryptionAlgorithm** | Pointer to **string** |  | [optional] 
**EncryptionHash** | Pointer to **string** |  | [optional] 

## Methods

### NewManagementV1alpha1ManagedApplicationSpecSecurity

`func NewManagementV1alpha1ManagedApplicationSpecSecurity(encryptionKey string, ) *ManagementV1alpha1ManagedApplicationSpecSecurity`

NewManagementV1alpha1ManagedApplicationSpecSecurity instantiates a new ManagementV1alpha1ManagedApplicationSpecSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1ManagedApplicationSpecSecurityWithDefaults

`func NewManagementV1alpha1ManagedApplicationSpecSecurityWithDefaults() *ManagementV1alpha1ManagedApplicationSpecSecurity`

NewManagementV1alpha1ManagedApplicationSpecSecurityWithDefaults instantiates a new ManagementV1alpha1ManagedApplicationSpecSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.


### GetEncryptionAlgorithm

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetEncryptionHash

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionHash() string`

GetEncryptionHash returns the EncryptionHash field if non-nil, zero value otherwise.

### GetEncryptionHashOk

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) GetEncryptionHashOk() (*string, bool)`

GetEncryptionHashOk returns a tuple with the EncryptionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionHash

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) SetEncryptionHash(v string)`

SetEncryptionHash sets EncryptionHash field to given value.

### HasEncryptionHash

`func (o *ManagementV1alpha1ManagedApplicationSpecSecurity) HasEncryptionHash() bool`

HasEncryptionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


