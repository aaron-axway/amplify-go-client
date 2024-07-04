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
	"bytes"
	"fmt"
)

// checks if the ManagementV1alpha1TraceabilityAgentSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1TraceabilityAgentSpec{}

// ManagementV1alpha1TraceabilityAgentSpec struct for ManagementV1alpha1TraceabilityAgentSpec
type ManagementV1alpha1TraceabilityAgentSpec struct {
	// The dataplane type that this agent connects to
	DataplaneType string `json:"dataplaneType"`
	Config ManagementV1alpha1TraceabilityAgentSpecConfig `json:"config"`
	Logging *ManagementV1alpha1DiscoveryAgentSpecLogging `json:"logging,omitempty"`
}

type _ManagementV1alpha1TraceabilityAgentSpec ManagementV1alpha1TraceabilityAgentSpec

// NewManagementV1alpha1TraceabilityAgentSpec instantiates a new ManagementV1alpha1TraceabilityAgentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1TraceabilityAgentSpec(dataplaneType string, config ManagementV1alpha1TraceabilityAgentSpecConfig) *ManagementV1alpha1TraceabilityAgentSpec {
	this := ManagementV1alpha1TraceabilityAgentSpec{}
	this.DataplaneType = dataplaneType
	this.Config = config
	return &this
}

// NewManagementV1alpha1TraceabilityAgentSpecWithDefaults instantiates a new ManagementV1alpha1TraceabilityAgentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1TraceabilityAgentSpecWithDefaults() *ManagementV1alpha1TraceabilityAgentSpec {
	this := ManagementV1alpha1TraceabilityAgentSpec{}
	return &this
}

// GetDataplaneType returns the DataplaneType field value
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetDataplaneType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataplaneType
}

// GetDataplaneTypeOk returns a tuple with the DataplaneType field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetDataplaneTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataplaneType, true
}

// SetDataplaneType sets field value
func (o *ManagementV1alpha1TraceabilityAgentSpec) SetDataplaneType(v string) {
	o.DataplaneType = v
}

// GetConfig returns the Config field value
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetConfig() ManagementV1alpha1TraceabilityAgentSpecConfig {
	if o == nil {
		var ret ManagementV1alpha1TraceabilityAgentSpecConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetConfigOk() (*ManagementV1alpha1TraceabilityAgentSpecConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ManagementV1alpha1TraceabilityAgentSpec) SetConfig(v ManagementV1alpha1TraceabilityAgentSpecConfig) {
	o.Config = v
}

// GetLogging returns the Logging field value if set, zero value otherwise.
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetLogging() ManagementV1alpha1DiscoveryAgentSpecLogging {
	if o == nil || IsNil(o.Logging) {
		var ret ManagementV1alpha1DiscoveryAgentSpecLogging
		return ret
	}
	return *o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpec) GetLoggingOk() (*ManagementV1alpha1DiscoveryAgentSpecLogging, bool) {
	if o == nil || IsNil(o.Logging) {
		return nil, false
	}
	return o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *ManagementV1alpha1TraceabilityAgentSpec) HasLogging() bool {
	if o != nil && !IsNil(o.Logging) {
		return true
	}

	return false
}

// SetLogging gets a reference to the given ManagementV1alpha1DiscoveryAgentSpecLogging and assigns it to the Logging field.
func (o *ManagementV1alpha1TraceabilityAgentSpec) SetLogging(v ManagementV1alpha1DiscoveryAgentSpecLogging) {
	o.Logging = &v
}

func (o ManagementV1alpha1TraceabilityAgentSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1TraceabilityAgentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataplaneType"] = o.DataplaneType
	toSerialize["config"] = o.Config
	if !IsNil(o.Logging) {
		toSerialize["logging"] = o.Logging
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1TraceabilityAgentSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataplaneType",
		"config",
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

	varManagementV1alpha1TraceabilityAgentSpec := _ManagementV1alpha1TraceabilityAgentSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1TraceabilityAgentSpec)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1TraceabilityAgentSpec(varManagementV1alpha1TraceabilityAgentSpec)

	return err
}

type NullableManagementV1alpha1TraceabilityAgentSpec struct {
	value *ManagementV1alpha1TraceabilityAgentSpec
	isSet bool
}

func (v NullableManagementV1alpha1TraceabilityAgentSpec) Get() *ManagementV1alpha1TraceabilityAgentSpec {
	return v.value
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpec) Set(val *ManagementV1alpha1TraceabilityAgentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1TraceabilityAgentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1TraceabilityAgentSpec(val *ManagementV1alpha1TraceabilityAgentSpec) *NullableManagementV1alpha1TraceabilityAgentSpec {
	return &NullableManagementV1alpha1TraceabilityAgentSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1TraceabilityAgentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1TraceabilityAgentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


