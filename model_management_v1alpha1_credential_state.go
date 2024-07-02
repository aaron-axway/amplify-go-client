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

// checks if the ManagementV1alpha1CredentialState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1CredentialState{}

// ManagementV1alpha1CredentialState Current state of the Credential.
type ManagementV1alpha1CredentialState struct {
	Name string `json:"name"`
	// Additional info on the state.
	Reason *string `json:"reason,omitempty"`
}

type _ManagementV1alpha1CredentialState ManagementV1alpha1CredentialState

// NewManagementV1alpha1CredentialState instantiates a new ManagementV1alpha1CredentialState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1CredentialState(name string) *ManagementV1alpha1CredentialState {
	this := ManagementV1alpha1CredentialState{}
	this.Name = name
	return &this
}

// NewManagementV1alpha1CredentialStateWithDefaults instantiates a new ManagementV1alpha1CredentialState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1CredentialStateWithDefaults() *ManagementV1alpha1CredentialState {
	this := ManagementV1alpha1CredentialState{}
	return &this
}

// GetName returns the Name field value
func (o *ManagementV1alpha1CredentialState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1CredentialState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementV1alpha1CredentialState) SetName(v string) {
	o.Name = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ManagementV1alpha1CredentialState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1CredentialState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ManagementV1alpha1CredentialState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ManagementV1alpha1CredentialState) SetReason(v string) {
	o.Reason = &v
}

func (o ManagementV1alpha1CredentialState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1CredentialState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1CredentialState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varManagementV1alpha1CredentialState := _ManagementV1alpha1CredentialState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1CredentialState)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1CredentialState(varManagementV1alpha1CredentialState)

	return err
}

type NullableManagementV1alpha1CredentialState struct {
	value *ManagementV1alpha1CredentialState
	isSet bool
}

func (v NullableManagementV1alpha1CredentialState) Get() *ManagementV1alpha1CredentialState {
	return v.value
}

func (v *NullableManagementV1alpha1CredentialState) Set(val *ManagementV1alpha1CredentialState) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1CredentialState) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1CredentialState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1CredentialState(val *ManagementV1alpha1CredentialState) *NullableManagementV1alpha1CredentialState {
	return &NullableManagementV1alpha1CredentialState{value: val, isSet: true}
}

func (v NullableManagementV1alpha1CredentialState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1CredentialState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


