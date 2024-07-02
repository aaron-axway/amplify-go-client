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

// checks if the ManagementV1alpha1EnvironmentPoliciesCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1EnvironmentPoliciesCredentials{}

// ManagementV1alpha1EnvironmentPoliciesCredentials Defines the policies for the Environment's Credentials
type ManagementV1alpha1EnvironmentPoliciesCredentials struct {
	Expiry *ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry `json:"expiry,omitempty"`
}

// NewManagementV1alpha1EnvironmentPoliciesCredentials instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1EnvironmentPoliciesCredentials() *ManagementV1alpha1EnvironmentPoliciesCredentials {
	this := ManagementV1alpha1EnvironmentPoliciesCredentials{}
	return &this
}

// NewManagementV1alpha1EnvironmentPoliciesCredentialsWithDefaults instantiates a new ManagementV1alpha1EnvironmentPoliciesCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1EnvironmentPoliciesCredentialsWithDefaults() *ManagementV1alpha1EnvironmentPoliciesCredentials {
	this := ManagementV1alpha1EnvironmentPoliciesCredentials{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentials) GetExpiry() ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry {
	if o == nil || IsNil(o.Expiry) {
		var ret ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentials) GetExpiryOk() (*ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentials) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry and assigns it to the Expiry field.
func (o *ManagementV1alpha1EnvironmentPoliciesCredentials) SetExpiry(v ManagementV1alpha1EnvironmentPoliciesCredentialsExpiry) {
	o.Expiry = &v
}

func (o ManagementV1alpha1EnvironmentPoliciesCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1EnvironmentPoliciesCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1EnvironmentPoliciesCredentials struct {
	value *ManagementV1alpha1EnvironmentPoliciesCredentials
	isSet bool
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentials) Get() *ManagementV1alpha1EnvironmentPoliciesCredentials {
	return v.value
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentials) Set(val *ManagementV1alpha1EnvironmentPoliciesCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1EnvironmentPoliciesCredentials(val *ManagementV1alpha1EnvironmentPoliciesCredentials) *NullableManagementV1alpha1EnvironmentPoliciesCredentials {
	return &NullableManagementV1alpha1EnvironmentPoliciesCredentials{value: val, isSet: true}
}

func (v NullableManagementV1alpha1EnvironmentPoliciesCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1EnvironmentPoliciesCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

