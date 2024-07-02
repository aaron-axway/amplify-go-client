# ManagementV1alpha1DataplaneSecretStatusError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **string** | Details of the error. | 
**Timestamp** | **time.Time** | Time when the update occurred. | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSecretStatusError

`func NewManagementV1alpha1DataplaneSecretStatusError(type_ string, detail string, timestamp time.Time, ) *ManagementV1alpha1DataplaneSecretStatusError`

NewManagementV1alpha1DataplaneSecretStatusError instantiates a new ManagementV1alpha1DataplaneSecretStatusError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSecretStatusErrorWithDefaults

`func NewManagementV1alpha1DataplaneSecretStatusErrorWithDefaults() *ManagementV1alpha1DataplaneSecretStatusError`

NewManagementV1alpha1DataplaneSecretStatusErrorWithDefaults instantiates a new ManagementV1alpha1DataplaneSecretStatusError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSecretStatusError) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ManagementV1alpha1DataplaneSecretStatusError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetTimestamp

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1DataplaneSecretStatusError) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ManagementV1alpha1DataplaneSecretStatusError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


