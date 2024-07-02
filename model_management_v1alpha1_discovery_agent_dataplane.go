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

// checks if the ManagementV1alpha1DiscoveryAgentDataplane type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1DiscoveryAgentDataplane{}

// ManagementV1alpha1DiscoveryAgentDataplane struct for ManagementV1alpha1DiscoveryAgentDataplane
type ManagementV1alpha1DiscoveryAgentDataplane struct {
	Name *string `json:"name,omitempty"`
	// Defines the interval that the dataplane will be accessed (30m = 30 minutes, 5h 5m = 5 hours and 5 mins, 2d = 2 days). 30m minimum
	Frequency *string `json:"frequency,omitempty" validate:"regexp=^(\\\\d*[d])?(\\\\d*[h])?(\\\\d*[m])?$"`
	// Queues this agent to run a discovery process. Defaults to false
	QueueDiscovery *bool `json:"queueDiscovery,omitempty"`
}

// NewManagementV1alpha1DiscoveryAgentDataplane instantiates a new ManagementV1alpha1DiscoveryAgentDataplane object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1DiscoveryAgentDataplane() *ManagementV1alpha1DiscoveryAgentDataplane {
	this := ManagementV1alpha1DiscoveryAgentDataplane{}
	return &this
}

// NewManagementV1alpha1DiscoveryAgentDataplaneWithDefaults instantiates a new ManagementV1alpha1DiscoveryAgentDataplane object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1DiscoveryAgentDataplaneWithDefaults() *ManagementV1alpha1DiscoveryAgentDataplane {
	this := ManagementV1alpha1DiscoveryAgentDataplane{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetName(v string) {
	o.Name = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetFrequency(v string) {
	o.Frequency = &v
}

// GetQueueDiscovery returns the QueueDiscovery field value if set, zero value otherwise.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetQueueDiscovery() bool {
	if o == nil || IsNil(o.QueueDiscovery) {
		var ret bool
		return ret
	}
	return *o.QueueDiscovery
}

// GetQueueDiscoveryOk returns a tuple with the QueueDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) GetQueueDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.QueueDiscovery) {
		return nil, false
	}
	return o.QueueDiscovery, true
}

// HasQueueDiscovery returns a boolean if a field has been set.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) HasQueueDiscovery() bool {
	if o != nil && !IsNil(o.QueueDiscovery) {
		return true
	}

	return false
}

// SetQueueDiscovery gets a reference to the given bool and assigns it to the QueueDiscovery field.
func (o *ManagementV1alpha1DiscoveryAgentDataplane) SetQueueDiscovery(v bool) {
	o.QueueDiscovery = &v
}

func (o ManagementV1alpha1DiscoveryAgentDataplane) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1DiscoveryAgentDataplane) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.QueueDiscovery) {
		toSerialize["queueDiscovery"] = o.QueueDiscovery
	}
	return toSerialize, nil
}

type NullableManagementV1alpha1DiscoveryAgentDataplane struct {
	value *ManagementV1alpha1DiscoveryAgentDataplane
	isSet bool
}

func (v NullableManagementV1alpha1DiscoveryAgentDataplane) Get() *ManagementV1alpha1DiscoveryAgentDataplane {
	return v.value
}

func (v *NullableManagementV1alpha1DiscoveryAgentDataplane) Set(val *ManagementV1alpha1DiscoveryAgentDataplane) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1DiscoveryAgentDataplane) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1DiscoveryAgentDataplane) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1DiscoveryAgentDataplane(val *ManagementV1alpha1DiscoveryAgentDataplane) *NullableManagementV1alpha1DiscoveryAgentDataplane {
	return &NullableManagementV1alpha1DiscoveryAgentDataplane{value: val, isSet: true}
}

func (v NullableManagementV1alpha1DiscoveryAgentDataplane) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1DiscoveryAgentDataplane) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

