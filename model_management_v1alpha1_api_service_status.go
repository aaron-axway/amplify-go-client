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
	"fmt"
)

// checks if the ManagementV1alpha1APIServiceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceStatus{}

// ManagementV1alpha1APIServiceStatus struct for ManagementV1alpha1APIServiceStatus
type ManagementV1alpha1APIServiceStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level"`
	// Reasons for the generated status.
	Reasons []ManagementV1alpha1APIServiceStatusReasonsInner `json:"reasons,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagementV1alpha1APIServiceStatus ManagementV1alpha1APIServiceStatus

// NewManagementV1alpha1APIServiceStatus instantiates a new ManagementV1alpha1APIServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceStatus(level string) *ManagementV1alpha1APIServiceStatus {
	this := ManagementV1alpha1APIServiceStatus{}
	this.Level = level
	return &this
}

// NewManagementV1alpha1APIServiceStatusWithDefaults instantiates a new ManagementV1alpha1APIServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceStatusWithDefaults() *ManagementV1alpha1APIServiceStatus {
	this := ManagementV1alpha1APIServiceStatus{}
	return &this
}

// GetLevel returns the Level field value
func (o *ManagementV1alpha1APIServiceStatus) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceStatus) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ManagementV1alpha1APIServiceStatus) SetLevel(v string) {
	o.Level = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceStatus) GetReasons() []ManagementV1alpha1APIServiceStatusReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []ManagementV1alpha1APIServiceStatusReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceStatus) GetReasonsOk() ([]ManagementV1alpha1APIServiceStatusReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceStatus) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ManagementV1alpha1APIServiceStatusReasonsInner and assigns it to the Reasons field.
func (o *ManagementV1alpha1APIServiceStatus) SetReasons(v []ManagementV1alpha1APIServiceStatusReasonsInner) {
	o.Reasons = v
}

func (o ManagementV1alpha1APIServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagementV1alpha1APIServiceStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
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

	varManagementV1alpha1APIServiceStatus := _ManagementV1alpha1APIServiceStatus{}

	err = json.Unmarshal(data, &varManagementV1alpha1APIServiceStatus)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1APIServiceStatus(varManagementV1alpha1APIServiceStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		delete(additionalProperties, "reasons")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagementV1alpha1APIServiceStatus struct {
	value *ManagementV1alpha1APIServiceStatus
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceStatus) Get() *ManagementV1alpha1APIServiceStatus {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceStatus) Set(val *ManagementV1alpha1APIServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceStatus(val *ManagementV1alpha1APIServiceStatus) *NullableManagementV1alpha1APIServiceStatus {
	return &NullableManagementV1alpha1APIServiceStatus{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


