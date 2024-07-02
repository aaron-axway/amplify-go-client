# ManagementV1alpha1DataplaneSecretStatusSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | Time when the change occured. | 
**Detail** | **string** | message of the result | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManagementV1alpha1DataplaneSecretStatusSuccess

`func NewManagementV1alpha1DataplaneSecretStatusSuccess(type_ string, timestamp time.Time, detail string, ) *ManagementV1alpha1DataplaneSecretStatusSuccess`

NewManagementV1alpha1DataplaneSecretStatusSuccess instantiates a new ManagementV1alpha1DataplaneSecretStatusSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1DataplaneSecretStatusSuccessWithDefaults

`func NewManagementV1alpha1DataplaneSecretStatusSuccessWithDefaults() *ManagementV1alpha1DataplaneSecretStatusSuccess`

NewManagementV1alpha1DataplaneSecretStatusSuccessWithDefaults instantiates a new ManagementV1alpha1DataplaneSecretStatusSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetDetail

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ManagementV1alpha1DataplaneSecretStatusSuccess) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


