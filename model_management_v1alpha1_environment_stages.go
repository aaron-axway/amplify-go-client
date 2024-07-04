/*
AMPLIFY Central API v0.347

APIs to manage AMPLIFY Central configuration resources.

API version: 0.347.0
Contact: support@axway.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amplify

import (
	"encoding/json"
)

// checks if the ManagementV1alpha1EnvironmentStages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1EnvironmentStages{}

// ManagementV1alpha1EnvironmentStages struct for ManagementV1alpha1EnvironmentStages
type ManagementV1alpha1EnvironmentStages struct {
	// The default stage to be assigned to the Environment's APIServiceInstances.
	Default *string `json:"default,omitempty"`
	// The allowed stages that can be set as override on the Environment's APIServiceInstance.
	Allowed []string `json:"allowed,omitempty"`
}

// NewManagementV1alpha1EnvironmentStages instantiates a new ManagementV1alpha1EnvironmentStages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1EnvironmentStages() *ManagementV1alpha1EnvironmentStages {
	this := ManagementV1alpha1EnvironmentStages{}
	return &this
}

// NewManagementV1alpha1EnvironmentStagesWithDefaults instantiates a new ManagementV1alpha1EnvironmentStages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1EnvironmentStagesWithDefaults() *ManagementV1alpha1EnvironmentStages {
	this := ManagementV1alpha1EnvironmentStages{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentStages) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentStages) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentStages) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *ManagementV1alpha1EnvironmentStages) SetDefault(v string) {
	o.Default = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentStages) GetAllowed() []string {
	if o == nil || IsNil(o.Allowed) {
		var ret []string
		return ret
	}
	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentStages) GetAllowedOk() ([]string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentStages) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given []string and assigns it to the Allowed field.
func (o *ManagementV1alpha1EnvironmentStages) SetAllowed(v []string) {
	o.Allowed = v
}

func (o ManagementV1alpha1EnvironmentStages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1EnvironmentStages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1EnvironmentStages struct {
	value *ManagementV1alpha1EnvironmentStages
	isSet bool
}

func (v NullableManagementV1alpha1EnvironmentStages) Get() *ManagementV1alpha1EnvironmentStages {
	return v.value
}

func (v *NullableManagementV1alpha1EnvironmentStages) Set(val *ManagementV1alpha1EnvironmentStages) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1EnvironmentStages) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1EnvironmentStages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1EnvironmentStages(val *ManagementV1alpha1EnvironmentStages) *NullableManagementV1alpha1EnvironmentStages {
	return &NullableManagementV1alpha1EnvironmentStages{value: val, isSet: true}
}

func (v NullableManagementV1alpha1EnvironmentStages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1EnvironmentStages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


