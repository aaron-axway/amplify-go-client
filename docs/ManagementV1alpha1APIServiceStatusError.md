# ManagementV1alpha1APIServiceStatusError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | **string** | Details of the error. | 
**Timestamp** | **time.Time** | Time when the change occurred. | 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewManagementV1alpha1APIServiceStatusError

`func NewManagementV1alpha1APIServiceStatusError(type_ string, detail string, timestamp time.Time, ) *ManagementV1alpha1APIServiceStatusError`

NewManagementV1alpha1APIServiceStatusError instantiates a new ManagementV1alpha1APIServiceStatusError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1APIServiceStatusErrorWithDefaults

`func NewManagementV1alpha1APIServiceStatusErrorWithDefaults() *ManagementV1alpha1APIServiceStatusError`

NewManagementV1alpha1APIServiceStatusErrorWithDefaults instantiates a new ManagementV1alpha1APIServiceStatusError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ManagementV1alpha1APIServiceStatusError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementV1alpha1APIServiceStatusError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementV1alpha1APIServiceStatusError) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *ManagementV1alpha1APIServiceStatusError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ManagementV1alpha1APIServiceStatusError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ManagementV1alpha1APIServiceStatusError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetTimestamp

`func (o *ManagementV1alpha1APIServiceStatusError) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ManagementV1alpha1APIServiceStatusError) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ManagementV1alpha1APIServiceStatusError) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMeta

`func (o *ManagementV1alpha1APIServiceStatusError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ManagementV1alpha1APIServiceStatusError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ManagementV1alpha1APIServiceStatusError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ManagementV1alpha1APIServiceStatusError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


