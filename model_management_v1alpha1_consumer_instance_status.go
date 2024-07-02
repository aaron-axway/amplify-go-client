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
	"bytes"
	"fmt"
)

// checks if the ManagementV1alpha1ConsumerInstanceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ConsumerInstanceStatus{}

// ManagementV1alpha1ConsumerInstanceStatus struct for ManagementV1alpha1ConsumerInstanceStatus
type ManagementV1alpha1ConsumerInstanceStatus struct {
	Phase ManagementV1alpha1ConsumerInstanceStatusPhase `json:"phase"`
}

type _ManagementV1alpha1ConsumerInstanceStatus ManagementV1alpha1ConsumerInstanceStatus

// NewManagementV1alpha1ConsumerInstanceStatus instantiates a new ManagementV1alpha1ConsumerInstanceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ConsumerInstanceStatus(phase ManagementV1alpha1ConsumerInstanceStatusPhase) *ManagementV1alpha1ConsumerInstanceStatus {
	this := ManagementV1alpha1ConsumerInstanceStatus{}
	this.Phase = phase
	return &this
}

// NewManagementV1alpha1ConsumerInstanceStatusWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ConsumerInstanceStatusWithDefaults() *ManagementV1alpha1ConsumerInstanceStatus {
	this := ManagementV1alpha1ConsumerInstanceStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *ManagementV1alpha1ConsumerInstanceStatus) GetPhase() ManagementV1alpha1ConsumerInstanceStatusPhase {
	if o == nil {
		var ret ManagementV1alpha1ConsumerInstanceStatusPhase
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatus) GetPhaseOk() (*ManagementV1alpha1ConsumerInstanceStatusPhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *ManagementV1alpha1ConsumerInstanceStatus) SetPhase(v ManagementV1alpha1ConsumerInstanceStatusPhase) {
	o.Phase = v
}

func (o ManagementV1alpha1ConsumerInstanceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ConsumerInstanceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phase"] = o.Phase
	return toSerialize, nil
}

func (o *ManagementV1alpha1ConsumerInstanceStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phase",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varManagementV1alpha1ConsumerInstanceStatus := _ManagementV1alpha1ConsumerInstanceStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1ConsumerInstanceStatus)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1ConsumerInstanceStatus(varManagementV1alpha1ConsumerInstanceStatus)

	return err
}

type NullableManagementV1alpha1ConsumerInstanceStatus struct {
	value *ManagementV1alpha1ConsumerInstanceStatus
	isSet bool
}

func (v NullableManagementV1alpha1ConsumerInstanceStatus) Get() *ManagementV1alpha1ConsumerInstanceStatus {
	return v.value
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatus) Set(val *ManagementV1alpha1ConsumerInstanceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ConsumerInstanceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ConsumerInstanceStatus(val *ManagementV1alpha1ConsumerInstanceStatus) *NullableManagementV1alpha1ConsumerInstanceStatus {
	return &NullableManagementV1alpha1ConsumerInstanceStatus{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ConsumerInstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


