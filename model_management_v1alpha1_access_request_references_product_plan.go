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

// checks if the ManagementV1alpha1AccessRequestReferencesProductPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementV1alpha1AccessRequestReferencesProductPlan{}

// ManagementV1alpha1AccessRequestReferencesProductPlan struct for ManagementV1alpha1AccessRequestReferencesProductPlan
type ManagementV1alpha1AccessRequestReferencesProductPlan struct {
	Kind string `json:"kind"`
	Name *string `json:"name,omitempty"`
}

type _ManagementV1alpha1AccessRequestReferencesProductPlan ManagementV1alpha1AccessRequestReferencesProductPlan

// NewManagementV1alpha1AccessRequestReferencesProductPlan instantiates a new ManagementV1alpha1AccessRequestReferencesProductPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementV1alpha1AccessRequestReferencesProductPlan(kind string) *ManagementV1alpha1AccessRequestReferencesProductPlan {
	this := ManagementV1alpha1AccessRequestReferencesProductPlan{}
	this.Kind = kind
	return &this
}

// NewManagementV1alpha1AccessRequestReferencesProductPlanWithDefaults instantiates a new ManagementV1alpha1AccessRequestReferencesProductPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementV1alpha1AccessRequestReferencesProductPlanWithDefaults() *ManagementV1alpha1AccessRequestReferencesProductPlan {
	this := ManagementV1alpha1AccessRequestReferencesProductPlan{}
	return &this
}

// GetKind returns the Kind field value
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) SetName(v string) {
	o.Name = &v
}

func (o ManagementV1alpha1AccessRequestReferencesProductPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementV1alpha1AccessRequestReferencesProductPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *ManagementV1alpha1AccessRequestReferencesProductPlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
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

	varManagementV1alpha1AccessRequestReferencesProductPlan := _ManagementV1alpha1AccessRequestReferencesProductPlan{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementV1alpha1AccessRequestReferencesProductPlan)

	if err != nil {
		return err
	}

	*o = ManagementV1alpha1AccessRequestReferencesProductPlan(varManagementV1alpha1AccessRequestReferencesProductPlan)

	return err
}

type NullableManagementV1alpha1AccessRequestReferencesProductPlan struct {
	value *ManagementV1alpha1AccessRequestReferencesProductPlan
	isSet bool
}

func (v NullableManagementV1alpha1AccessRequestReferencesProductPlan) Get() *ManagementV1alpha1AccessRequestReferencesProductPlan {
	return v.value
}

func (v *NullableManagementV1alpha1AccessRequestReferencesProductPlan) Set(val *ManagementV1alpha1AccessRequestReferencesProductPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementV1alpha1AccessRequestReferencesProductPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementV1alpha1AccessRequestReferencesProductPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementV1alpha1AccessRequestReferencesProductPlan(val *ManagementV1alpha1AccessRequestReferencesProductPlan) *NullableManagementV1alpha1AccessRequestReferencesProductPlan {
	return &NullableManagementV1alpha1AccessRequestReferencesProductPlan{value: val, isSet: true}
}

func (v NullableManagementV1alpha1AccessRequestReferencesProductPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementV1alpha1AccessRequestReferencesProductPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


