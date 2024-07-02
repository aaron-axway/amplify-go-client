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

// checks if the ManagementV1APISpecLintingJobResultDetailsInnerLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1APISpecLintingJobResultDetailsInnerLocation{}

// ManagementV1APISpecLintingJobResultDetailsInnerLocation The location of the linting result.
type ManagementV1APISpecLintingJobResultDetailsInnerLocation struct {
	// The 1-based line number in the API Specification File.
	Line int32 `json:"line"`
	// The 1-based character number in the API Specification File.
	Character int32 `json:"character"`
}

type _ManagementV1APISpecLintingJobResultDetailsInnerLocation ManagementV1APISpecLintingJobResultDetailsInnerLocation

// NewManagementV1APISpecLintingJobResultDetailsInnerLocation instantiates a new ManagementV1APISpecLintingJobResultDetailsInnerLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1APISpecLintingJobResultDetailsInnerLocation(line int32, character int32) *ManagementV1APISpecLintingJobResultDetailsInnerLocation {
	this := ManagementV1APISpecLintingJobResultDetailsInnerLocation{}
	this.Line = line
	this.Character = character
	return &this
}

// NewManagementV1APISpecLintingJobResultDetailsInnerLocationWithDefaults instantiates a new ManagementV1APISpecLintingJobResultDetailsInnerLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1APISpecLintingJobResultDetailsInnerLocationWithDefaults() *ManagementV1APISpecLintingJobResultDetailsInnerLocation {
	this := ManagementV1APISpecLintingJobResultDetailsInnerLocation{}
	return &this
}

// GetLine returns the Line field value
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) GetLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) GetLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) SetLine(v int32) {
	o.Line = v
}

// GetCharacter returns the Character field value
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) GetCharacter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Character
}

// GetCharacterOk returns a tuple with the Character field value
// and a boolean to check if the value has been set.
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) GetCharacterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Character, true
}

// SetCharacter sets field value
func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) SetCharacter(v int32) {
	o.Character = v
}

func (o ManagementV1APISpecLintingJobResultDetailsInnerLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1APISpecLintingJobResultDetailsInnerLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["line"] = o.Line
	toSerialize["character"] = o.Character
	return toSerialize, nil
}

func (o *ManagementV1APISpecLintingJobResultDetailsInnerLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"line",
		"character",
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

	varManagementV1APISpecLintingJobResultDetailsInnerLocation := _ManagementV1APISpecLintingJobResultDetailsInnerLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1APISpecLintingJobResultDetailsInnerLocation)

	if err != nil {
		return err
	}

	*o = ManagementV1APISpecLintingJobResultDetailsInnerLocation(varManagementV1APISpecLintingJobResultDetailsInnerLocation)

	return err
}

type NullableManagementV1APISpecLintingJobResultDetailsInnerLocation struct {
	value *ManagementV1APISpecLintingJobResultDetailsInnerLocation
	isSet bool
}

func (v NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) Get() *ManagementV1APISpecLintingJobResultDetailsInnerLocation {
	return v.value
}

func (v *NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) Set(val *ManagementV1APISpecLintingJobResultDetailsInnerLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1APISpecLintingJobResultDetailsInnerLocation(val *ManagementV1APISpecLintingJobResultDetailsInnerLocation) *NullableManagementV1APISpecLintingJobResultDetailsInnerLocation {
	return &NullableManagementV1APISpecLintingJobResultDetailsInnerLocation{value: val, isSet: true}
}

func (v NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1APISpecLintingJobResultDetailsInnerLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


