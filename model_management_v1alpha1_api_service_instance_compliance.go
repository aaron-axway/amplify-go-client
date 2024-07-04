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

// checks if the ManagementV1alpha1APIServiceInstanceCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceCompliance{}

// ManagementV1alpha1APIServiceInstanceCompliance struct for ManagementV1alpha1APIServiceInstanceCompliance
type ManagementV1alpha1APIServiceInstanceCompliance struct {
	Runtime *ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus `json:"runtime,omitempty"`
}

// NewManagementV1alpha1APIServiceInstanceCompliance instantiates a new ManagementV1alpha1APIServiceInstanceCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceCompliance() *ManagementV1alpha1APIServiceInstanceCompliance {
	this := ManagementV1alpha1APIServiceInstanceCompliance{}
	return &this
}

// NewManagementV1alpha1APIServiceInstanceComplianceWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceComplianceWithDefaults() *ManagementV1alpha1APIServiceInstanceCompliance {
	this := ManagementV1alpha1APIServiceInstanceCompliance{}
	return &this
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceCompliance) GetRuntime() ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus {
	if o == nil || IsNil(o.Runtime) {
		var ret ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceCompliance) GetRuntimeOk() (*ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceCompliance) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus and assigns it to the Runtime field.
func (o *ManagementV1alpha1APIServiceInstanceCompliance) SetRuntime(v ManagementV1alpha1APIServiceInstanceComplianceRuntimeStatus) {
	o.Runtime = &v
}

func (o ManagementV1alpha1APIServiceInstanceCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APIServiceInstanceCompliance struct {
	value *ManagementV1alpha1APIServiceInstanceCompliance
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceCompliance) Get() *ManagementV1alpha1APIServiceInstanceCompliance {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceCompliance) Set(val *ManagementV1alpha1APIServiceInstanceCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceCompliance(val *ManagementV1alpha1APIServiceInstanceCompliance) *NullableManagementV1alpha1APIServiceInstanceCompliance {
	return &NullableManagementV1alpha1APIServiceInstanceCompliance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


