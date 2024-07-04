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

// checks if the ManagementV1BatchJobSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1BatchJobSpec{}

// ManagementV1BatchJobSpec struct for ManagementV1BatchJobSpec
type ManagementV1BatchJobSpec struct {
	// Unique string ID assigned by a controller indicating what action this job is performing.
	Action string `json:"action"`
	References []map[string]interface{} `json:"references,omitempty"`
}

type _ManagementV1BatchJobSpec ManagementV1BatchJobSpec

// NewManagementV1BatchJobSpec instantiates a new ManagementV1BatchJobSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1BatchJobSpec(action string) *ManagementV1BatchJobSpec {
	this := ManagementV1BatchJobSpec{}
	this.Action = action
	return &this
}

// NewManagementV1BatchJobSpecWithDefaults instantiates a new ManagementV1BatchJobSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1BatchJobSpecWithDefaults() *ManagementV1BatchJobSpec {
	this := ManagementV1BatchJobSpec{}
	return &this
}

// GetAction returns the Action field value
func (o *ManagementV1BatchJobSpec) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ManagementV1BatchJobSpec) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ManagementV1BatchJobSpec) SetAction(v string) {
	o.Action = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ManagementV1BatchJobSpec) GetReferences() []map[string]interface{} {
	if o == nil || IsNil(o.References) {
		var ret []map[string]interface{}
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1BatchJobSpec) GetReferencesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.References) {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ManagementV1BatchJobSpec) HasReferences() bool {
	if o != nil && !IsNil(o.References) {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []map[string]interface{} and assigns it to the References field.
func (o *ManagementV1BatchJobSpec) SetReferences(v []map[string]interface{}) {
	o.References = v
}

func (o ManagementV1BatchJobSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1BatchJobSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.References) {
		toSerialize["references"] = o.References
	}
	return toSerialize, nil
}

func (o *ManagementV1BatchJobSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varManagementV1BatchJobSpec := _ManagementV1BatchJobSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1BatchJobSpec)

	if err != nil {
		return err
	}

	*o = ManagementV1BatchJobSpec(varManagementV1BatchJobSpec)

	return err
}

type NullableManagementV1BatchJobSpec struct {
	value *ManagementV1BatchJobSpec
	isSet bool
}

func (v NullableManagementV1BatchJobSpec) Get() *ManagementV1BatchJobSpec {
	return v.value
}

func (v *NullableManagementV1BatchJobSpec) Set(val *ManagementV1BatchJobSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1BatchJobSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1BatchJobSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1BatchJobSpec(val *ManagementV1BatchJobSpec) *NullableManagementV1BatchJobSpec {
	return &NullableManagementV1BatchJobSpec{value: val, isSet: true}
}

func (v NullableManagementV1BatchJobSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1BatchJobSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


