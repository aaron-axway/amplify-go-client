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

// checks if the ManagementV1alpha1DataplaneSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1DataplaneSpec{}

// ManagementV1alpha1DataplaneSpec struct for ManagementV1alpha1DataplaneSpec
type ManagementV1alpha1DataplaneSpec struct {
	// The dataplane type that this agent connects to
	Type string `json:"type"`
	Config *ManagementV1alpha1DataplaneSpecConfig `json:"config,omitempty"`
}

type _ManagementV1alpha1DataplaneSpec ManagementV1alpha1DataplaneSpec

// NewManagementV1alpha1DataplaneSpec instantiates a new ManagementV1alpha1DataplaneSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1DataplaneSpec(type_ string) *ManagementV1alpha1DataplaneSpec {
	this := ManagementV1alpha1DataplaneSpec{}
	this.Type = type_
	return &this
}

// NewManagementV1alpha1DataplaneSpecWithDefaults instantiates a new ManagementV1alpha1DataplaneSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1DataplaneSpecWithDefaults() *ManagementV1alpha1DataplaneSpec {
	this := ManagementV1alpha1DataplaneSpec{}
	return &this
}

// GetType returns the Type field value
func (o *ManagementV1alpha1DataplaneSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManagementV1alpha1DataplaneSpec) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ManagementV1alpha1DataplaneSpec) GetConfig() ManagementV1alpha1DataplaneSpecConfig {
	if o == nil || IsNil(o.Config) {
		var ret ManagementV1alpha1DataplaneSpecConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DataplaneSpec) GetConfigOk() (*ManagementV1alpha1DataplaneSpecConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ManagementV1alpha1DataplaneSpec) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ManagementV1alpha1DataplaneSpecConfig and assigns it to the Config field.
func (o *ManagementV1alpha1DataplaneSpec) SetConfig(v ManagementV1alpha1DataplaneSpecConfig) {
	o.Config = &v
}

func (o ManagementV1alpha1DataplaneSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1DataplaneSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1DataplaneSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varManagementV1alpha1DataplaneSpec := _ManagementV1alpha1DataplaneSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1DataplaneSpec)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1DataplaneSpec(varManagementV1alpha1DataplaneSpec)

	return err
}

type NullableManagementV1alpha1DataplaneSpec struct {
	value *ManagementV1alpha1DataplaneSpec
	isSet bool
}

func (v NullableManagementV1alpha1DataplaneSpec) Get() *ManagementV1alpha1DataplaneSpec {
	return v.value
}

func (v *NullableManagementV1alpha1DataplaneSpec) Set(val *ManagementV1alpha1DataplaneSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1DataplaneSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1DataplaneSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1DataplaneSpec(val *ManagementV1alpha1DataplaneSpec) *NullableManagementV1alpha1DataplaneSpec {
	return &NullableManagementV1alpha1DataplaneSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1DataplaneSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1DataplaneSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

