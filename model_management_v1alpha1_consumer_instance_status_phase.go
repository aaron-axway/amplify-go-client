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
	"time"
	"bytes"
	"fmt"
)

// checks if the ManagementV1alpha1ConsumerInstanceStatusPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1ConsumerInstanceStatusPhase{}

// ManagementV1alpha1ConsumerInstanceStatusPhase struct for ManagementV1alpha1ConsumerInstanceStatusPhase
type ManagementV1alpha1ConsumerInstanceStatusPhase struct {
	Name string `json:"name"`
	// The criticality of the last update
	Level string `json:"level"`
	// Time of the current phase.
	TransitionTime time.Time `json:"transitionTime"`
	// Description of the phase.
	Message *string `json:"message,omitempty"`
}

type _ManagementV1alpha1ConsumerInstanceStatusPhase ManagementV1alpha1ConsumerInstanceStatusPhase

// NewManagementV1alpha1ConsumerInstanceStatusPhase instantiates a new ManagementV1alpha1ConsumerInstanceStatusPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1ConsumerInstanceStatusPhase(name string, level string, transitionTime time.Time) *ManagementV1alpha1ConsumerInstanceStatusPhase {
	this := ManagementV1alpha1ConsumerInstanceStatusPhase{}
	this.Name = name
	this.Level = level
	this.TransitionTime = transitionTime
	return &this
}

// NewManagementV1alpha1ConsumerInstanceStatusPhaseWithDefaults instantiates a new ManagementV1alpha1ConsumerInstanceStatusPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1ConsumerInstanceStatusPhaseWithDefaults() *ManagementV1alpha1ConsumerInstanceStatusPhase {
	this := ManagementV1alpha1ConsumerInstanceStatusPhase{}
	return &this
}

// GetName returns the Name field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetName(v string) {
	o.Name = v
}

// GetLevel returns the Level field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetLevel(v string) {
	o.Level = v
}

// GetTransitionTime returns the TransitionTime field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetTransitionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransitionTime
}

// GetTransitionTimeOk returns a tuple with the TransitionTime field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetTransitionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransitionTime, true
}

// SetTransitionTime sets field value
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetTransitionTime(v time.Time) {
	o.TransitionTime = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) SetMessage(v string) {
	o.Message = &v
}

func (o ManagementV1alpha1ConsumerInstanceStatusPhase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1ConsumerInstanceStatusPhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["level"] = o.Level
	toSerialize["transitionTime"] = o.TransitionTime
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1ConsumerInstanceStatusPhase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"level",
		"transitionTime",
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

	varManagementV1alpha1ConsumerInstanceStatusPhase := _ManagementV1alpha1ConsumerInstanceStatusPhase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1ConsumerInstanceStatusPhase)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1ConsumerInstanceStatusPhase(varManagementV1alpha1ConsumerInstanceStatusPhase)

	return err
}

type NullableManagementV1alpha1ConsumerInstanceStatusPhase struct {
	value *ManagementV1alpha1ConsumerInstanceStatusPhase
	isSet bool
}

func (v NullableManagementV1alpha1ConsumerInstanceStatusPhase) Get() *ManagementV1alpha1ConsumerInstanceStatusPhase {
	return v.value
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatusPhase) Set(val *ManagementV1alpha1ConsumerInstanceStatusPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1ConsumerInstanceStatusPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatusPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1ConsumerInstanceStatusPhase(val *ManagementV1alpha1ConsumerInstanceStatusPhase) *NullableManagementV1alpha1ConsumerInstanceStatusPhase {
	return &NullableManagementV1alpha1ConsumerInstanceStatusPhase{value: val, isSet: true}
}

func (v NullableManagementV1alpha1ConsumerInstanceStatusPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1ConsumerInstanceStatusPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


