# ManagementV1alpha1AccessRequestSpecQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The limit of the allowed quota for the access request. | 
**Interval** | **string** |  | 

## Methods

### NewManagementV1alpha1AccessRequestSpecQuota

`func NewManagementV1alpha1AccessRequestSpecQuota(limit int32, interval string, ) *ManagementV1alpha1AccessRequestSpecQuota`

NewManagementV1alpha1AccessRequestSpecQuota instantiates a new ManagementV1alpha1AccessRequestSpecQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1AccessRequestSpecQuotaWithDefaults

`func NewManagementV1alpha1AccessRequestSpecQuotaWithDefaults() *ManagementV1alpha1AccessRequestSpecQuota`

NewManagementV1alpha1AccessRequestSpecQuotaWithDefaults instantiates a new ManagementV1alpha1AccessRequestSpecQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ManagementV1alpha1AccessRequestSpecQuota) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ManagementV1alpha1AccessRequestSpecQuota) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ManagementV1alpha1AccessRequestSpecQuota) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetInterval

`func (o *ManagementV1alpha1AccessRequestSpecQuota) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ManagementV1alpha1AccessRequestSpecQuota) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ManagementV1alpha1AccessRequestSpecQuota) SetInterval(v string)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


