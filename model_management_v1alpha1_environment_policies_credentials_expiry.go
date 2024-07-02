/*
AMPLIFY Central API v0.347

APIs to manage AMPLIFY Central configuration resources.

API version: 0.347.0
Contact: support@axway.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry{}

// ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry Expiry properties for Credentials generated in the scoped Environment.
type ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry struct {
	// The number of days after the Credentials are considered to be expired.
	Period *int32 `json:"period,omitempty"`
	// The action taken when the Credential expires.
	Action *string `json:"action,omitempty"`
	Notifications *CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications `json:"notifications,omitempty"`
}

// NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiry instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiry() *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry {
	this := ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry{}
	return &this
}

// NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiryWithDefaults instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1EnvironmentPoliciesCredentialsExpiryWithDefaults() *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry {
	this := ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetPeriod() int32 {
	if o == nil || IsNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetPeriod(v int32) {
	o.Period = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetAction(v string) {
	o.Action = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetNotifications() CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) GetNotificationsOk() (*CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications and assigns it to the Notifications field.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) SetNotifications(v CatalogV1alpha1AuthorizationProfilePoliciesCredentialsExpiryNotifications) {
	o.Notifications = &v
}

func (o ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry struct {
	value *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry
	isSet bool
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) Get() *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry {
	return v.value
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) Set(val *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry(val *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) *NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry {
	return &NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry{value: val, isSet: true}
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

