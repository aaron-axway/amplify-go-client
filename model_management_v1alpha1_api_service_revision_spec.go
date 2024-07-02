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

// checks if the ManagementV1alpha1APIServiceRevisionSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1APIServiceRevisionSpec{}

// ManagementV1alpha1APIServiceRevisionSpec struct for ManagementV1alpha1APIServiceRevisionSpec
type ManagementV1alpha1APIServiceRevisionSpec struct {
	ApiService string `json:"apiService"`
	Definition *ManagementV1alpha1APIServiceRevisionSpecDefinition `json:"definition,omitempty"`
}

type _ManagementV1alpha1APIServiceRevisionSpec ManagementV1alpha1APIServiceRevisionSpec

// NewManagementV1alpha1APIServiceRevisionSpec instantiates a new ManagementV1alpha1APIServiceRevisionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1APIServiceRevisionSpec(apiService string) *ManagementV1alpha1APIServiceRevisionSpec {
	this := ManagementV1alpha1APIServiceRevisionSpec{}
	this.ApiService = apiService
	return &this
}

// NewManagementV1alpha1APIServiceRevisionSpecWithDefaults instantiates a new ManagementV1alpha1APIServiceRevisionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1APIServiceRevisionSpecWithDefaults() *ManagementV1alpha1APIServiceRevisionSpec {
	this := ManagementV1alpha1APIServiceRevisionSpec{}
	return &this
}

// GetApiService returns the ApiService field value
func (o *ManagementV1alpha1APIServiceRevisionSpec) GetApiService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiService
}

// GetApiServiceOk returns a tuple with the ApiService field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceRevisionSpec) GetApiServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiService, true
}

// SetApiService sets field value
func (o *ManagementV1alpha1APIServiceRevisionSpec) SetApiService(v string) {
	o.ApiService = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *ManagementV1alpha1APIServiceRevisionSpec) GetDefinition() ManagementV1alpha1APIServiceRevisionSpecDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret ManagementV1alpha1APIServiceRevisionSpecDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1APIServiceRevisionSpec) GetDefinitionOk() (*ManagementV1alpha1APIServiceRevisionSpecDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ManagementV1alpha1APIServiceRevisionSpec) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given ManagementV1alpha1APIServiceRevisionSpecDefinition and assigns it to the Definition field.
func (o *ManagementV1alpha1APIServiceRevisionSpec) SetDefinition(v ManagementV1alpha1APIServiceRevisionSpecDefinition) {
	o.Definition = &v
}

func (o ManagementV1alpha1APIServiceRevisionSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1APIServiceRevisionSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiService"] = o.ApiService
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1APIServiceRevisionSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiService",
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

	varManagementV1alpha1APIServiceRevisionSpec := _ManagementV1alpha1APIServiceRevisionSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1APIServiceRevisionSpec)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1APIServiceRevisionSpec(varManagementV1alpha1APIServiceRevisionSpec)

	return err
}

type NullableManagementV1alpha1APIServiceRevisionSpec struct {
	value *ManagementV1alpha1APIServiceRevisionSpec
	isSet bool
}

func (v NullableManagementV1alpha1APIServiceRevisionSpec) Get() *ManagementV1alpha1APIServiceRevisionSpec {
	return v.value
}

func (v *NullableManagementV1alpha1APIServiceRevisionSpec) Set(val *ManagementV1alpha1APIServiceRevisionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1APIServiceRevisionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1APIServiceRevisionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1APIServiceRevisionSpec(val *ManagementV1alpha1APIServiceRevisionSpec) *NullableManagementV1alpha1APIServiceRevisionSpec {
	return &NullableManagementV1alpha1APIServiceRevisionSpec{value: val, isSet: true}
}

func (v NullableManagementV1alpha1APIServiceRevisionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1APIServiceRevisionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


