# ManagementV1alpha1DataplaneSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | **string** | public key to be used to encrypt the access data linked to this dataplane. | 
**EncryptionAlgorithm** | Pointer to **string** |  | [optional] 
**EncryptionHash** | Pointer to **string** |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSecurity

`func NewManagementV1alpha1DataplaneSecurity(encryptionKey string, ) *ManagementV1alpha1DataplaneSecurity`

NewManagementV1alpha1DataplaneSecurity instantiates a new ManagementV1alpha1DataplaneSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSecurityWithDefaults

`func NewManagementV1alpha1DataplaneSecurityWithDefaults() *ManagementV1alpha1DataplaneSecurity`

NewManagementV1alpha1DataplaneSecurityWithDefaults instantiates a new ManagementV1alpha1DataplaneSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ManagementV1alpha1DataplaneSecurity) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.


### GetEncryptionAlgorithm

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *ManagementV1alpha1DataplaneSecurity) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *ManagementV1alpha1DataplaneSecurity) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetEncryptionHash

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionHash() string`

GetEncryptionHash returns the EncryptionHash field if non-nil, zero value otherwise.

### GetEncryptionHashOk

`func (o *ManagementV1alpha1DataplaneSecurity) GetEncryptionHashOk() (*string, bool)`

GetEncryptionHashOk returns a tuple with the EncryptionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionHash

`func (o *ManagementV1alpha1DataplaneSecurity) SetEncryptionHash(v string)`

SetEncryptionHash sets EncryptionHash field to given value.

### HasEncryptionHash

`func (o *ManagementV1alpha1DataplaneSecurity) HasEncryptionHash() bool`

HasEncryptionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


