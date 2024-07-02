# CatalogV1alpha1ApplicationSpecSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | **string** | public key to be used to encrypt the credentials linked to this Application. | 
**EncryptionAlgorithm** | Pointer to **string** |  | [optional] 
**EncryptionHash** | Pointer to **string** |  | [optional] 

## Methods

### NewCatalogV1alpha1ApplicationSpecSecurity

`func NewCatalogV1alpha1ApplicationSpecSecurity(encryptionKey string, ) *CatalogV1alpha1ApplicationSpecSecurity`

NewCatalogV1alpha1ApplicationSpecSecurity instantiates a new CatalogV1alpha1ApplicationSpecSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogV1alpha1ApplicationSpecSecurityWithDefaults

`func NewCatalogV1alpha1ApplicationSpecSecurityWithDefaults() *CatalogV1alpha1ApplicationSpecSecurity`

NewCatalogV1alpha1ApplicationSpecSecurityWithDefaults instantiates a new CatalogV1alpha1ApplicationSpecSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionKey

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *CatalogV1alpha1ApplicationSpecSecurity) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.


### GetEncryptionAlgorithm

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CatalogV1alpha1ApplicationSpecSecurity) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *CatalogV1alpha1ApplicationSpecSecurity) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetEncryptionHash

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionHash() string`

GetEncryptionHash returns the EncryptionHash field if non-nil, zero value otherwise.

### GetEncryptionHashOk

`func (o *CatalogV1alpha1ApplicationSpecSecurity) GetEncryptionHashOk() (*string, bool)`

GetEncryptionHashOk returns a tuple with the EncryptionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionHash

`func (o *CatalogV1alpha1ApplicationSpecSecurity) SetEncryptionHash(v string)`

SetEncryptionHash sets EncryptionHash field to given value.

### HasEncryptionHash

`func (o *CatalogV1alpha1ApplicationSpecSecurity) HasEncryptionHash() bool`

HasEncryptionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


