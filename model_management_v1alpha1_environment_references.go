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

// checks if the ManagementV1alpha1EnvironmentReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1EnvironmentReferences{}

// ManagementV1alpha1EnvironmentReferences struct for ManagementV1alpha1EnvironmentReferences
type ManagementV1alpha1EnvironmentReferences struct {
	// The list of referenced managed Environments
	ManagedEnvironments []string `json:"managedEnvironments,omitempty"`
}

// NewManagementV1alpha1EnvironmentReferences instantiates a new ManagementV1alpha1EnvironmentReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1EnvironmentReferences() *ManagementV1alpha1EnvironmentReferences {
	this := ManagementV1alpha1EnvironmentReferences{}
	return &this
}

// NewManagementV1alpha1EnvironmentReferencesWithDefaults instantiates a new ManagementV1alpha1EnvironmentReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1EnvironmentReferencesWithDefaults() *ManagementV1alpha1EnvironmentReferences {
	this := ManagementV1alpha1EnvironmentReferences{}
	return &this
}

// GetManagedEnvironments returns the ManagedEnvironments field value if set, zero value otherwise.
func (o *ManagementV1alpha1EnvironmentReferences) GetManagedEnvironments() []string {
	if o == nil || IsNil(o.ManagedEnvironments) {
		var ret []string
		return ret
	}
	return o.ManagedEnvironments
}

// GetManagedEnvironmentsOk returns a tuple with the ManagedEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1EnvironmentReferences) GetManagedEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedEnvironments) {
		return nil, false
	}
	return o.ManagedEnvironments, true
}

// HasManagedEnvironments returns a boolean if a field has been set.
func (o *ManagementV1alpha1EnvironmentReferences) HasManagedEnvironments() bool {
	if o != nil && !IsNil(o.ManagedEnvironments) {
		return true
	}

	return false
}

// SetManagedEnvironments gets a reference to the given []string and assigns it to the ManagedEnvironments field.
func (o *ManagementV1alpha1EnvironmentReferences) SetManagedEnvironments(v []string) {
	o.ManagedEnvironments = v
}

func (o ManagementV1alpha1EnvironmentReferences) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1EnvironmentReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedEnvironments) {
		toSerialize["managedEnvironments"] = o.ManagedEnvironments
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1EnvironmentReferences struct {
	value *ManagementV1alpha1EnvironmentReferences
	isSet bool
}

func (v NullableManagementV1alpha1EnvironmentReferences) Get() *ManagementV1alpha1EnvironmentReferences {
	return v.value
}

func (v *NullableManagementV1alpha1EnvironmentReferences) Set(val *ManagementV1alpha1EnvironmentReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1EnvironmentReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1EnvironmentReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1EnvironmentReferences(val *ManagementV1alpha1EnvironmentReferences) *NullableManagementV1alpha1EnvironmentReferences {
	return &NullableManagementV1alpha1EnvironmentReferences{value: val, isSet: true}
}

func (v NullableManagementV1alpha1EnvironmentReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1EnvironmentReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

