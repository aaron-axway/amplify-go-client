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

// checks if the ManagementV1alpha1APIServiceInstanceSourceCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceInstanceSourceCompliance{}

// ManagementV1alpha1APIServiceInstanceSourceCompliance struct for ManagementV1alpha1APIServiceInstanceSourceCompliance
type ManagementV1alpha1APIServiceInstanceSourceCompliance struct {
	Runtime *ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus `json:"runtime,omitempty"`
}

// NewManagementV1alpha1APIServiceInstanceSourceCompliance instantiates a new ManagementV1alpha1APIServiceInstanceSourceCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceInstanceSourceCompliance() *ManagementV1alpha1APIServiceInstanceSourceCompliance {
	this := ManagementV1alpha1APIServiceInstanceSourceCompliance{}
	return &this
}

// NewManagementV1alpha1APIServiceInstanceSourceComplianceWithDefaults instantiates a new ManagementV1alpha1APIServiceInstanceSourceCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceInstanceSourceComplianceWithDefaults() *ManagementV1alpha1APIServiceInstanceSourceCompliance {
	this := ManagementV1alpha1APIServiceInstanceSourceCompliance{}
	return &this
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceInstanceSourceCompliance) GetRuntime() ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus {
	if o == nil || IsNil(o.Runtime) {
		var ret ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceCompliance) GetRuntimeOk() (*ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceInstanceSourceCompliance) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus and assigns it to the Runtime field.
func (o *ManagementV1alpha1APIServiceInstanceSourceCompliance) SetRuntime(v ManagementV1alpha1APIServiceInstanceSourceRuntimeStatus) {
	o.Runtime = &v
}

func (o ManagementV1alpha1APIServiceInstanceSourceCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceInstanceSourceCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1APIServiceInstanceSourceCompliance struct {
	value *ManagementV1alpha1APIServiceInstanceSourceCompliance
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceCompliance) Get() *ManagementV1alpha1APIServiceInstanceSourceCompliance {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceCompliance) Set(val *ManagementV1alpha1APIServiceInstanceSourceCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceInstanceSourceCompliance(val *ManagementV1alpha1APIServiceInstanceSourceCompliance) *NullableManagementV1alpha1APIServiceInstanceSourceCompliance {
	return &NullableManagementV1alpha1APIServiceInstanceSourceCompliance{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceInstanceSourceCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceInstanceSourceCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


