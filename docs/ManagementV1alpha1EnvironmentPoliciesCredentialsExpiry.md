# ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** | The number of days after the Credentials are considered to be expired. | [optional] 
**Action** | Pointer to **string** | The action taken when the Credential expires. | [optional] 
**Notifications** | Pointer to [**CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications**](CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications.md) |  | [optional] 

## Methods

### NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiry

`func NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiry() *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry`

NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiry instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiryWithDefaults

`func NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiryWithDefaults() *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry`

NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiryWithDefaults instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetAction

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetNotifications

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetNotifications() CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetNotificationsOk() (*CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetNotifications(v CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


